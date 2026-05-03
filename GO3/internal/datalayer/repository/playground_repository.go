package repository

import (
	"fmt"
	"os"
	"s21_rogue/internal/app/mappers"
	"s21_rogue/internal/domain/aggregates"
	"s21_rogue/internal/domain/dto"
	"s21_rogue/internal/domain/utils"

	"github.com/mailru/easyjson"
	"go.uber.org/zap"
)

type PlaygroundRepository interface {
	Save(playground *aggregates.Playground, noCache bool) error
	Get(noCache bool) (*aggregates.Playground, error)
}

type playgroundJsonCacheRepository struct {
	filePath string
	cache    *aggregates.Playground
	log      utils.Logger
}

func NewPlaygroundRepositoryJson(filePath string,
	logger utils.Logger) PlaygroundRepository {

	return &playgroundJsonCacheRepository{
		filePath: filePath,
		log:      logger}
}

func (r *playgroundJsonCacheRepository) Get(noCache bool) (*aggregates.Playground, error) {
	if noCache {
		r.log.Info("Try to load dungeon from file")

		file, err := os.Open(r.filePath)
		if err != nil {
			r.log.Error("Unable to open file", zap.String("filepath", r.filePath))
			return nil, err
		}
		defer file.Close()

		playgroundDTO := dto.PlaygroundDTO{}
		err = easyjson.UnmarshalFromReader(file, &playgroundDTO)
		if err != nil {
			r.log.Error("Incorrect unmarshalling dungeon data from json", zap.Error(err))
			return nil, err
		}

		playground, err := mappers.MapFromPlaygroundDTO(&playgroundDTO)
		if err != nil {
			return nil, err
		}

		r.cache = playground
		return playground, nil
	}

	if r.cache == nil {
		return nil, fmt.Errorf("Cache is nil")
	}
	return r.cache, nil
}

func (r *playgroundJsonCacheRepository) Save(playground *aggregates.Playground, noCache bool) error {
	if noCache {
		r.log.Info("Try to save dungeon")
		file, err := os.Create(r.filePath)
		if err != nil {
			r.log.Error("Incorrect creating file", zap.String("filepath", r.filePath))
			return err
		}
		defer file.Close()

		playgroundDTO := mappers.MapToPlaygroundDTO(playground)

		_, err = easyjson.MarshalToWriter(playgroundDTO, file)
		if err != nil {
			r.log.Error("Incorrect dungeon marshall to json")
			return err
		}
	}
	r.cache = playground
	playground = nil
	return nil
}
