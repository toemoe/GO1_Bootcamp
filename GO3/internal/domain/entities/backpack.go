package entities

import "container/list"

const MaxLenIntems = 9

type Backpack struct {
	Foods     *list.List //Food
	Scrolls   *list.List //Scroll
	Elixirs   *list.List //Elixir
	Weapons   *list.List //*Weapon
	Treasures *list.List //Treasure
}

func NewBackpack() *Backpack {
	return &Backpack{Foods: list.New(),
		Scrolls:   list.New(),
		Elixirs:   list.New(),
		Weapons:   list.New(),
		Treasures: list.New()}
}

func (bp *Backpack) AppendFood(food Food) bool {
	if bp.Foods.Len() < MaxLenIntems {
		bp.Foods.PushBack(food)
		return true
	}
	return false
}

func (bp *Backpack) AppendScroll(scroll Scroll) bool {
	if bp.Scrolls.Len() < MaxLenIntems {
		bp.Scrolls.PushBack(scroll)
		return true
	}
	return false
}

func (bp *Backpack) AppendElixir(elixir Elixir) bool {
	if bp.Elixirs.Len() < MaxLenIntems {
		bp.Elixirs.PushBack(elixir)
		return true
	}
	return false
}

func (bp *Backpack) AppendWeapon(weapon *Weapon) bool {
	if bp.Weapons.Len() < MaxLenIntems {
		bp.Weapons.PushBack(weapon)
		return true
	}
	return false
}

func (bp *Backpack) AppendTreasure(treasure Treasure) {
	bp.Treasures.PushBack(treasure)
}
