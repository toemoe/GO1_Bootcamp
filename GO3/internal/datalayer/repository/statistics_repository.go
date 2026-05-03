package repository

import (
	"os"
	"s21_rogue/internal/app/mappers"
	"s21_rogue/internal/domain/dto"
	"s21_rogue/internal/domain/entities"
	"s21_rogue/internal/domain/utils"
	"sort"

	"github.com/google/uuid"
	"github.com/mailru/easyjson"
	"go.uber.org/zap"
)

type StatisticsRepository interface {
	Save(statistics *entities.Statistics) error
	FindAll() []*entities.Statistics
	GetLevelById(uuid.UUID) int
	FindLastStats() map[uuid.UUID]*entities.Statistics
	FindLeaderOrderByTreasure(limit int) []*entities.Statistics
}

func NewStatisticsJsonRepository(filePath string, log utils.Logger) StatisticsRepository {
	return &statisticsJsonCacheRepository{
		filePath: filePath,
		log:      log}
}

type statisticsJsonCacheRepository struct {
	cache    *entities.Statistics
	filePath string
	log      utils.Logger
}

func (r *statisticsJsonCacheRepository) FindAll() []*entities.Statistics {
	r.log.Info("Try to load statistics list from file")

	file, err := os.Open(r.filePath)
	if err != nil {
		r.log.Warn("Unable to open file", zap.String("filepath", r.filePath))
		return make([]*entities.Statistics, 0)
	}
	defer file.Close()

	statisticsSliceDTO := dto.StatisticsSliceDTO{}
	err = easyjson.UnmarshalFromReader(file, &statisticsSliceDTO)
	if err != nil {
		r.log.Warn("Incorrect unmarshalling statistics data from json", zap.Error(err))
		return make([]*entities.Statistics, 0)
	}

	res := make([]*entities.Statistics, 0, len(statisticsSliceDTO.Statistics))
	for i := range statisticsSliceDTO.Statistics {
		s, err := mappers.MapFromStatisticsDTO(&statisticsSliceDTO.Statistics[i])
		if err != nil {
			r.log.Warn("Incorrect mapping statistics data from json", zap.Error(err))
			return make([]*entities.Statistics, 0)
		}
		res = append(res, s)
	}
	return res
}

func (r *statisticsJsonCacheRepository) GetLevelById(id uuid.UUID) int {
	if r.cache != nil && r.cache.Id == id {
		return r.cache.Level
	}
	s, ok := r.FindLastStats()[id]
	if !ok {
		return 1
	}
	r.cache = s
	return r.cache.Level
}

func (r *statisticsJsonCacheRepository) Save(statistics *entities.Statistics) error {
	r.log.Info("Try to save statistics")
	l := r.FindAll()

	resDTO := make([]dto.StatisticsDTO, 0, len(l)+1)

	for _, v := range l {
		res := mappers.MapToStatisticsDTO(v)
		resDTO = append(resDTO, *res)
	}

	lastStat := mappers.MapToStatisticsDTO(statistics)
	resDTO = append(resDTO, *lastStat)

	file, err := os.Create(r.filePath)
	if err != nil {
		r.log.Error("Incorrect creating file", zap.String("filepath", r.filePath))
		return err
	}
	defer file.Close()

	_, err = easyjson.MarshalToWriter(dto.StatisticsSliceDTO{Statistics: resDTO}, file)
	if err != nil {
		r.log.Error("Incorrect statistics marshall to json")
		return err
	}
	r.cache = statistics
	return nil
}

func (r *statisticsJsonCacheRepository) FindLastStats() map[uuid.UUID]*entities.Statistics {
	allStats := r.FindAll()

	resultMap := make(map[uuid.UUID]*entities.Statistics)

	for _, stat := range allStats {
		currentStat, exists := resultMap[stat.Id]
		if !exists || currentStat.Level < stat.Level {
			resultMap[stat.Id] = stat
		}
	}
	return resultMap
}

func (r *statisticsJsonCacheRepository) FindLeaderOrderByTreasure(limit int) []*entities.Statistics {
	lastStats := r.FindLastStats()

	res := make([]*entities.Statistics, 0, len(lastStats))
	for _, v := range lastStats {
		res = append(res, v)
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].TreasureFound > res[j].TreasureFound
	})

	if limit < len(res) {
		return res[:limit]
	}
	return res
}
