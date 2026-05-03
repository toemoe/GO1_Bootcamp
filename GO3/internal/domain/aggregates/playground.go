package aggregates

import (
	"container/list"
	"math"
	"math/rand"
	"s21_rogue/internal/domain/constants"
	"s21_rogue/internal/domain/entities"
	"s21_rogue/internal/domain/events"
	"s21_rogue/internal/domain/utils"
	"s21_rogue/internal/domain/value_objects"
	"slices"

	"github.com/google/uuid"
)

type Playground struct {
	Id        uuid.UUID
	Dungeon   *entities.Dungeon
	Character *entities.Character
	Backpack  *entities.Backpack
	Monsters  *list.List // Monster
	Foods     map[value_objects.Position]entities.Food
	Scrolls   map[value_objects.Position]entities.Scroll
	Elixirs   map[value_objects.Position]entities.Elixir
	Weapons   map[value_objects.Position]*entities.Weapon
	Treasures map[value_objects.Position]entities.Treasure
}

func NewPlayground(id uuid.UUID,
	dungeon *entities.Dungeon,
	character *entities.Character,
	backpack *entities.Backpack,
	monsters *list.List,
	foods map[value_objects.Position]entities.Food,
	scrolls map[value_objects.Position]entities.Scroll,
	elixirs map[value_objects.Position]entities.Elixir,
	weapons map[value_objects.Position]*entities.Weapon) *Playground {

	return &Playground{
		Id:        id,
		Dungeon:   dungeon,
		Character: character,
		Backpack:  backpack,
		Monsters:  monsters,
		Foods:     foods,
		Scrolls:   scrolls,
		Elixirs:   elixirs,
		Weapons:   weapons,
		Treasures: make(map[value_objects.Position]entities.Treasure)}
}

func (p *Playground) AttackOrUpdateMonstersPosition() (ev []events.Event) {
	ev = make([]events.Event, 0)
	p.updateChasing()

	sortedMonsters := utils.GetSortedMonsterByStepPriority(p.Monsters, p.Character)

	for i := range sortedMonsters {
		monster := sortedMonsters[i]
		if monster.CanAttack(p.Character.Position) {
			attackCommand := monster.GetAttackCommand()
			prevHealth := p.Character.GetHealth()
			p.Character.TakeDamage(attackCommand)
			if prevHealth != p.Character.GetHealth() {
				ev = append(ev, events.NewEvent(events.MonsterAttacked))
			}
		} else if monster.IsChasing() {
			nextPos := utils.GetNextMonsterPosition(p.Dungeon, p.Character, monster)
			if !p.isMonsterPosition(nextPos) {
				monster.SetPosition(nextPos)
			}
		} else {
			monsterRoomI, monsterRoomJ, _ := p.Dungeon.SearchRoomIndexByPos(monster.GetPosition().X, monster.GetPosition().Y)
			monsterRoom := p.Dungeon.Rooms[monsterRoomI][monsterRoomJ]
			busyPos := make([]*value_objects.Position, 0)
			busyPos = append(busyPos, p.Character.Position)
			busyPos = append(busyPos, p.Dungeon.Portal)
			monsterInRoom := p.getMonstersInRoomByPos(monster.GetPosition())

			for e := monsterInRoom.Front(); e != nil; e = e.Next() {
				mRoom := e.Value.(entities.Monster)
				if monster != mRoom {
					busyPos = append(busyPos, mRoom.GetPosition())
				}
			}

			for fPos := range p.Foods {
				busyPos = append(busyPos, &fPos)
			}
			for sPos := range p.Scrolls {
				busyPos = append(busyPos, &sPos)
			}
			for ePos := range p.Elixirs {
				busyPos = append(busyPos, &ePos)
			}
			for wPos := range p.Weapons {
				busyPos = append(busyPos, &wPos)
			}
			for tPos := range p.Treasures {
				busyPos = append(busyPos, &tPos)
			}

			monster.Walk(&monsterRoom.TopLeft, &monsterRoom.BotRight, busyPos)
		}
	}
	return
}

func (p *Playground) AttackOrUpdateCharacterPosition(nextCharacterPos *value_objects.Position) (ev []events.Event) {
	ev = make([]events.Event, 0)
	if p.Dungeon.InDungeon(nextCharacterPos) {
		if p.isMonsterPosition(nextCharacterPos) {
			attackCmd := p.Character.GetAttackCommand()
			monster := p.getMonsterFromPosition(nextCharacterPos)

			prevHealth := monster.GetHealth()
			monster.TakeDamage(attackCmd)
			if monster.GetHealth() != prevHealth {
				ev = append(ev, events.NewEvent(events.CharacterAttacked))
			}

			evs := p.deleteKilledMonster()
			ev = append(ev, evs...)
		} else {
			food, foodExist := p.Foods[*nextCharacterPos]
			if foodExist && p.Backpack.AppendFood(food) {
				delete(p.Foods, *nextCharacterPos)
			}

			scroll, scrollExist := p.Scrolls[*nextCharacterPos]
			if scrollExist && p.Backpack.AppendScroll(scroll) {
				delete(p.Scrolls, *nextCharacterPos)
			}

			elixir, elixirExist := p.Elixirs[*nextCharacterPos]
			if elixirExist && p.Backpack.AppendElixir(elixir) {
				delete(p.Elixirs, *nextCharacterPos)
			}

			weapon, weaponExist := p.Weapons[*nextCharacterPos]
			if weaponExist && p.Backpack.AppendWeapon(weapon) {
				delete(p.Weapons, *nextCharacterPos)
			}

			treasure, treasureExist := p.Treasures[*nextCharacterPos]
			if treasureExist {
				p.Backpack.AppendTreasure(treasure)
				delete(p.Treasures, *nextCharacterPos)
				for range treasure.Score {
					ev = append(ev, events.NewEvent(events.TreasureFoundEvent))
				}
			}

			p.Character.SetPosition(nextCharacterPos)
			p.Dungeon.UpdateVisitedState(nextCharacterPos)
			ev = append(ev, events.NewEvent(events.CharacterStepped))
		}
	}
	return ev
}

func (p *Playground) IsPortalPosition(pos *value_objects.Position) bool {
	return pos.IsEqual(p.Dungeon.Portal)
}

func (p *Playground) UseItemByUUID(id uuid.UUID) (ev events.Event) {
	ev = events.NewEvent(events.NothingEvent)
	for e := p.Backpack.Foods.Front(); e != nil; e = e.Next() {
		food := e.Value.(entities.Food)
		if food.GetUUID() == id {
			p.Character.AppendHealth(food.Health)
			p.Backpack.Foods.Remove(e)
			ev = events.NewEvent(events.FoodEatenEvent)
			return
		}
	}

	for e := p.Backpack.Scrolls.Front(); e != nil; e = e.Next() {
		scroll := e.Value.(entities.Scroll)
		if scroll.GetUUID() == id {
			switch scroll.BoostType {
			case entities.AgileBoost:
				p.Character.AppendAgile(scroll.Value)
			case entities.MaxHealthBoost:
				p.Character.AppendMaxHealth(scroll.Value)
			case entities.StrengthBoost:
				p.Character.AppendStrength(scroll.Value)
			}

			p.Backpack.Scrolls.Remove(e)
			ev = events.NewEvent(events.ScrollsReadEvent)
			return
		}
	}

	for e := p.Backpack.Elixirs.Front(); e != nil; e = e.Next() {
		elixir := e.Value.(entities.Elixir)
		if elixir.GetUUID() == id {
			switch elixir.BoostType {
			case entities.AgileBoost:
				p.Character.BoostAgile(elixir.Value, elixir.CountSteps)
			case entities.MaxHealthBoost:
				p.Character.BoostMaxHealth(elixir.Value, elixir.CountSteps)
			case entities.StrengthBoost:
				p.Character.BoostStrength(elixir.Value, elixir.CountSteps)
			}

			p.Backpack.Elixirs.Remove(e)
			ev = events.NewEvent(events.ElixirsDrunkEvent)
			return
		}
	}

	isSelectedWeapon := false
	for e := p.Backpack.Weapons.Front(); e != nil && !isSelectedWeapon; e = e.Next() {
		weapon := e.Value.(*entities.Weapon)
		if weapon.GetUUID() == id {
			isSelectedWeapon = true
			p.Character.SetWeaponStrength(weapon.Strength)
			weapon.Selected = true
		}
	}
	if isSelectedWeapon {
		for e := p.Backpack.Weapons.Front(); e != nil; e = e.Next() {
			weapon := e.Value.(*entities.Weapon)
			if weapon.GetUUID() != id && weapon.Selected {
				weapon.Selected = false
				p.dropWeapon(e)
			}
		}
	}
	return
}

func (p *Playground) deleteKilledMonster() (ev []events.Event) {
	ev = make([]events.Event, 0)
	var toRemove []*list.Element
	for e := p.Monsters.Front(); e != nil; e = e.Next() {
		monster := e.Value.(entities.Monster)
		if monster.GetHealth() <= 0 {
			if rand.Intn(2) != 0 {
				costTreasure := entities.ZombieCostTreasure
				switch monster.GetMonsterType() {
				case entities.ZombieType:
					costTreasure = entities.ZombieCostTreasure
				case entities.VampireType:
					costTreasure = entities.VampireCostTreasure
				case entities.GhostType:
					costTreasure = entities.GhostCostTreasure
				case entities.OgreType:
					costTreasure = entities.OgreCostTreasure
				case entities.SnakeType:
					costTreasure = entities.SnakeCostTreasure
				}
				p.Treasures[*monster.GetPosition()] = entities.GenerateTreasure(costTreasure)
			}
			toRemove = append(toRemove, e)
			ev = append(ev, events.NewEvent(events.MonsterDeadEvent))
		}
	}
	for _, elem := range toRemove {
		p.Monsters.Remove(elem)
	}
	return
}

func (p *Playground) isMonsterPosition(position *value_objects.Position) bool {
	for e := p.Monsters.Front(); e != nil; e = e.Next() {
		monster := e.Value.(entities.Monster)
		if monster.GetPosition().IsEqual(position) {
			return true
		}
	}
	return false
}

func (p *Playground) getMonsterFromPosition(position *value_objects.Position) entities.Monster {
	for e := p.Monsters.Front(); e != nil; e = e.Next() {
		monster := e.Value.(entities.Monster)
		if monster.GetPosition().IsEqual(position) {
			return monster
		}
	}
	return nil
}

func (p *Playground) getMonstersInCharacterRoom() *list.List {
	return p.getMonstersInRoomByPos(p.Character.Position)
}

func (p *Playground) getMonstersInRoomByPos(pos *value_objects.Position) *list.List {
	var searchRoom *entities.Room
	isFound := false
	for i := 0; i < constants.RoomsPerSide && !isFound; i++ {
		for j := 0; j < constants.RoomsPerSide && !isFound; j++ {
			if p.Dungeon.Rooms[i][j].InRoom(pos.X, pos.Y) {
				searchRoom = &p.Dungeon.Rooms[i][j]
				isFound = true
			}
		}
	}

	res := list.New()
	if isFound {
		for e := p.Monsters.Front(); e != nil; e = e.Next() {
			monster := e.Value.(entities.Monster)
			if searchRoom.InRoom(monster.GetPosition().X, monster.GetPosition().Y) {
				res.PushBack(monster)
			}
		}
	}
	return res
}

func (p *Playground) updateChasing() {
	monstersInCharacterRoom := p.getMonstersInCharacterRoom()
	for e := monstersInCharacterRoom.Front(); e != nil; e = e.Next() {
		monster := e.Value.(entities.Monster)
		dist := math.Sqrt(math.Pow(float64(monster.GetPosition().X-p.Character.Position.X), 2) + math.Pow(float64(monster.GetPosition().Y-p.Character.Position.Y), 2))
		if dist < float64(monster.GetHostility()) {
			monster.SetChasing(true)
		}
	}
}

func (p *Playground) dropWeapon(weapon *list.Element) {
	freePosAroundCharacter := make([]*value_objects.Position, 0)

	otherBusyPos := make([]*value_objects.Position, 0)

	for e := p.Monsters.Front(); e != nil; e = e.Next() {
		monster := e.Value.(entities.Monster)
		otherBusyPos = append(otherBusyPos, monster.GetPosition())
	}

	for p := range p.Foods {
		otherBusyPos = append(otherBusyPos, &p)
	}

	for p := range p.Scrolls {
		otherBusyPos = append(otherBusyPos, &p)
	}

	for p := range p.Elixirs {
		otherBusyPos = append(otherBusyPos, &p)
	}

	for p := range p.Weapons {
		otherBusyPos = append(otherBusyPos, &p)
	}

	for p := range p.Treasures {
		otherBusyPos = append(otherBusyPos, &p)
	}
	otherBusyPos = append(otherBusyPos, p.Dungeon.Portal)

	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			dropPos := value_objects.NewPosition(p.Character.Position.X+x, p.Character.Position.Y+y)
			if !(x == 0 && y == 0) &&
				p.Dungeon.InDungeon(dropPos) &&
				!slices.Contains(otherBusyPos, dropPos) {
				freePosAroundCharacter = append(freePosAroundCharacter, dropPos)
			}
		}
	}

	if len(freePosAroundCharacter) != 0 {
		randPos := freePosAroundCharacter[rand.Intn(len(freePosAroundCharacter))]
		p.Weapons[*randPos] = weapon.Value.(*entities.Weapon)
		p.Backpack.Weapons.Remove(weapon)
	}
}
