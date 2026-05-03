package entities

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

type Item struct {
	id    uuid.UUID
	Label string
}

func (i Item) IsSelected() bool {
	return false
}

func (i Item) GetUUID() uuid.UUID {
	return i.id
}

func (i Item) GetLabel() string {
	return i.Label
}

type Food struct {
	Item
	Health int
}

func NewFood(id string, label string, health int) (*Food, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return &Food{Item: Item{id: uuid, Label: label}, Health: health}, nil
}

func GenerateFood() Food {
	id := uuid.New()
	foods := []Food{
		{Health: 18, Item: Item{Label: "Banana"}},
		{Health: 22, Item: Item{Label: "Sandwich"}},
		{Health: 35, Item: Item{Label: "Steak"}},
		{Health: 28, Item: Item{Label: "Yogurt"}},
		{Health: 32, Item: Item{Label: "Health Elixir"}},
	}
	item := foods[rand.Intn(len(foods))]
	item.id = id
	return item
}

type BoostType int

const (
	AgileBoost BoostType = iota
	StrengthBoost
	MaxHealthBoost
)

type Scroll struct {
	Item
	BoostType BoostType
	Value     int
}

func NewScroll(id string, label string, boostType string, value int) (*Scroll, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	var bt BoostType
	switch boostType {
	case "agile":
		bt = AgileBoost
	case "strength":
		bt = StrengthBoost
	case "max_health":
		bt = MaxHealthBoost
	default:
		return nil, fmt.Errorf("Incorrect boost")
	}

	return &Scroll{Item: Item{id: uuid, Label: label}, BoostType: bt, Value: value}, nil
}

func GenerateScroll() Scroll {
	id := uuid.New()
	scrolls := []Scroll{
		{Value: 3, BoostType: AgileBoost, Item: Item{Label: "Swift Wind's Embrace"}},
		{Value: 6, BoostType: AgileBoost, Item: Item{Label: "Dancer’s Grace"}},
		{Value: 8, BoostType: AgileBoost, Item: Item{Label: "Shadow Stride"}},
		{Value: 10, BoostType: AgileBoost, Item: Item{Label: "Quickstep Parchment"}},
		{Value: 12, BoostType: AgileBoost, Item: Item{Label: "Acrobat's Blessing"}},

		{Value: 3, BoostType: StrengthBoost, Item: Item{Label: "Titan's Might"}},
		{Value: 6, BoostType: StrengthBoost, Item: Item{Label: "Warrior's Fury"}},
		{Value: 8, BoostType: StrengthBoost, Item: Item{Label: "Bear Claw Power"}},
		{Value: 10, BoostType: StrengthBoost, Item: Item{Label: "Ironfist Scroll"}},
		{Value: 12, BoostType: StrengthBoost, Item: Item{Label: "Colossus’ Vigor"}},

		{Value: 7, BoostType: MaxHealthBoost, Item: Item{Label: "Healer's Boon"}},
		{Value: 12, BoostType: MaxHealthBoost, Item: Item{Label: "Lifeforce Renewal"}},
		{Value: 15, BoostType: MaxHealthBoost, Item: Item{Label: "Stamina Surge"}},
		{Value: 18, BoostType: MaxHealthBoost, Item: Item{Label: "Endurance Amulet"}},
		{Value: 22, BoostType: MaxHealthBoost, Item: Item{Label: "Vitality Infusion"}},
	}
	item := scrolls[rand.Intn(len(scrolls))]
	item.id = id
	return item
}

type Elixir struct {
	Scroll
	CountSteps int
}

func NewElixir(id string, label string, boostType string, value int, countSteps int) (*Elixir, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	var bt BoostType
	switch boostType {
	case "agile":
		bt = AgileBoost
	case "strength":
		bt = StrengthBoost
	case "max_health":
		bt = MaxHealthBoost
	default:
		return nil, fmt.Errorf("Incorrect boost")
	}

	return &Elixir{Scroll: Scroll{Item: Item{id: uuid, Label: label}, BoostType: bt, Value: value}, CountSteps: countSteps}, nil
}

func GenerateElixir() Elixir {
	id := uuid.New()
	elixirs := []Elixir{
		{Scroll: Scroll{Value: 3, BoostType: AgileBoost, Item: Item{Label: "Fleetfoot Formula"}}, CountSteps: 5},
		{Scroll: Scroll{Value: 6, BoostType: AgileBoost, Item: Item{Label: "Nimble Step Brew"}}, CountSteps: 8},
		{Scroll: Scroll{Value: 8, BoostType: AgileBoost, Item: Item{Label: "Lightning Reflexes Draft"}}, CountSteps: 10},
		{Scroll: Scroll{Value: 10, BoostType: AgileBoost, Item: Item{Label: "Shadow Dash Concoction"}}, CountSteps: 12},
		{Scroll: Scroll{Value: 12, BoostType: AgileBoost, Item: Item{Label: "Catlike Agility Potion"}}, CountSteps: 15},

		{Scroll: Scroll{Value: 3, BoostType: StrengthBoost, Item: Item{Label: "Brute Force Decoction"}}, CountSteps: 6},
		{Scroll: Scroll{Value: 6, BoostType: StrengthBoost, Item: Item{Label: "Muscle Growth Mix"}}, CountSteps: 9},
		{Scroll: Scroll{Value: 8, BoostType: StrengthBoost, Item: Item{Label: "Powerful Hammer Broth"}}, CountSteps: 11},
		{Scroll: Scroll{Value: 10, BoostType: StrengthBoost, Item: Item{Label: "Battle Hardened Extract"}}, CountSteps: 13},
		{Scroll: Scroll{Value: 12, BoostType: StrengthBoost, Item: Item{Label: "Unbreakable Sinew Solution"}}, CountSteps: 16},

		{Scroll: Scroll{Value: 7, BoostType: MaxHealthBoost, Item: Item{Label: "Health Restoration Draught"}}, CountSteps: 7},
		{Scroll: Scroll{Value: 12, BoostType: MaxHealthBoost, Item: Item{Label: "Bloodstone Elixir"}}, CountSteps: 10},
		{Scroll: Scroll{Value: 15, BoostType: MaxHealthBoost, Item: Item{Label: "Ancient Vitality Syrup"}}, CountSteps: 12},
		{Scroll: Scroll{Value: 18, BoostType: MaxHealthBoost, Item: Item{Label: "Golden Wellness Brew"}}, CountSteps: 14},
		{Scroll: Scroll{Value: 22, BoostType: MaxHealthBoost, Item: Item{Label: "Dragonheart Regeneration Potion"}}, CountSteps: 18},
	}
	item := elixirs[rand.Intn(len(elixirs))]
	item.id = id
	return item
}

type Weapon struct {
	Item
	Strength int
	Selected bool
}

func NewWeapon(id string, label string, strength int, selected bool) (*Weapon, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return &Weapon{Item: Item{id: uuid, Label: label}, Strength: strength, Selected: selected}, nil
}

func GenerateWeapon() *Weapon {
	id := uuid.New()
	weapons := []Weapon{
		{Item: Item{Label: "Rusty Dagger"}, Strength: 3, Selected: false},
		{Item: Item{Label: "Short Sword"}, Strength: 5, Selected: false},
		{Item: Item{Label: "Hunting Bow"}, Strength: 7, Selected: false},
		{Item: Item{Label: "Broad Axe"}, Strength: 9, Selected: false},
		{Item: Item{Label: "Warhammer"}, Strength: 11, Selected: false},
		{Item: Item{Label: "Legendary Blade"}, Strength: 15, Selected: false},
	}
	item := weapons[rand.Intn(len(weapons))]
	item.id = id
	return &item
}

func (w *Weapon) IsSelected() bool {
	return w.Selected
}

type Treasure struct {
	Item
	Score int
}

func NewTreasure(id string, label string, score int) (*Treasure, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return &Treasure{Item: Item{id: uuid, Label: label}, Score: score}, nil
}

type MonsterCostTreasure int

const (
	ZombieCostTreasure MonsterCostTreasure = iota + 1
	VampireCostTreasure
	GhostCostTreasure
	OgreCostTreasure
	SnakeCostTreasure
)

func GenerateTreasure(cost MonsterCostTreasure) Treasure {
	id := uuid.New()
	score := int(cost)
	treasure := []Treasure{
		{Score: score, Item: Item{Label: "Philosopher's Stone"}},
		{Score: score, Item: Item{Label: "Ancient Warrior's Amulet"}},
		{Score: score, Item: Item{Label: "Ring of Eternity"}},
		{Score: score, Item: Item{Label: "Silver Dagger of Justice"}},
		{Score: score, Item: Item{Label: "Shield of Courage"}},
	}
	item := treasure[rand.Intn(len(treasure))]
	item.id = id
	return item
}
