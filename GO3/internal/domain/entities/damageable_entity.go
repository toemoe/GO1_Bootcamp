package entities

import (
	"container/list"
	"math"
	"math/rand"
	"s21_rogue/internal/domain/commands"
)

type DamageableEntitiyBehaviour interface {
	TakeDamage(cmd commands.AttackCommand)
	GetAttackCommand() commands.AttackCommand
	IsTranquilize() bool
	SetTranquilize(bool)
	GetAgile() int
	AppendAgile(int)
	BoostAgile(int, int)
	GetAgileBoosts() *list.List
	GetStrength() int
	SetWeaponStrength(int)
	AppendStrength(int)
	BoostStrength(int, int)
	GetStrengthBoosts() *list.List
	GetMaxHealth() int
	AppendMaxHealth(int)
	BoostMaxHealth(int, int)
	GetMaxHealthBoosts() *list.List
	GetHealth() int
	AppendHealth(int)
}

type BoostTimes struct {
	BoostValue, CountSteps int
}

type damageableEntitiy struct {
	CurrentHealth      int
	MaxHealth          int
	Agile              int
	Strength           int
	WeaponStrength     int
	isTranquilize      bool
	takeDamageFn       func(cmd commands.AttackCommand)
	getAttackCommandFn func() commands.AttackCommand
	AgileBoost         *list.List //BoostTimes
	MaxHealthBoost     *list.List //BoostTimes
	StrengthBoost      *list.List //BoostTimes
}

func NewSimpleDamageableEntity(health, maxHealth, agile, strength int) DamageableEntitiyBehaviour {
	de := &damageableEntitiy{CurrentHealth: health,
		MaxHealth:      maxHealth,
		Agile:          agile,
		Strength:       strength,
		WeaponStrength: 0,
		AgileBoost:     list.New(),
		MaxHealthBoost: list.New(),
		StrengthBoost:  list.New()}
	de.takeDamageFn = de.DefaultTakeDamage
	de.getAttackCommandFn = de.DefaultGetAttackCommand
	return de
}

func NewVampireDamageableEntity(health, agile, strength int) DamageableEntitiyBehaviour {
	de := &vampireDamageableEntitiy{damageableEntitiy: &damageableEntitiy{CurrentHealth: health,
		MaxHealth:      health,
		Agile:          agile,
		Strength:       strength,
		AgileBoost:     list.New(),
		MaxHealthBoost: list.New(),
		StrengthBoost:  list.New()}}
	de.takeDamageFn = de.vampireTakeDamage
	de.getAttackCommandFn = de.vampireGetAttackCommand
	return de
}

func NewOrgeDamageableEntity(health, agile, strength int) DamageableEntitiyBehaviour {
	de := &orgeDamageableEntitiy{damageableEntitiy: &damageableEntitiy{CurrentHealth: health,
		MaxHealth:      health,
		Agile:          agile,
		Strength:       strength,
		AgileBoost:     list.New(),
		MaxHealthBoost: list.New(),
		StrengthBoost:  list.New()}}
	de.takeDamageFn = de.DefaultTakeDamage
	de.getAttackCommandFn = de.orgeGetAttackCommand
	return de
}

func NewSnakeDamageableEntity(health, agile, strength int) DamageableEntitiyBehaviour {
	de := &damageableEntitiy{CurrentHealth: health,
		MaxHealth:      health,
		Agile:          agile,
		Strength:       strength,
		AgileBoost:     list.New(),
		MaxHealthBoost: list.New(),
		StrengthBoost:  list.New()}
	de.takeDamageFn = de.DefaultTakeDamage
	de.getAttackCommandFn = de.snakeGetAttackCommand
	return de
}

func (de *damageableEntitiy) TakeDamage(cmd commands.AttackCommand) {
	de.takeDamageFn(cmd)
}

func (de *damageableEntitiy) DefaultTakeDamage(cmd commands.AttackCommand) {
	hitChange := math.Min(95, float64(de.Agile-cmd.AttackAgile)*5)
	if float64(rand.Intn(100)) > hitChange || cmd.WithoutMiss {
		de.CurrentHealth -= cmd.Damage
		newMaxHealth := max(de.MaxHealth-cmd.StealMaxHealth, 10)
		de.MaxHealth = newMaxHealth
	}
}

func (de *damageableEntitiy) GetAttackCommand() commands.AttackCommand {
	return de.getAttackCommandFn()
}

func (de *damageableEntitiy) DefaultGetAttackCommand() commands.AttackCommand {
	removeElems := make([]*list.Element, 0)
	for e := de.AgileBoost.Front(); e != nil; e = e.Next() {
		boost := e.Value.(BoostTimes)
		boost.CountSteps--
		if boost.CountSteps <= 0 {
			removeElems = append(removeElems, e)
			de.Agile = max(1, de.Agile-boost.BoostValue)
		}
	}
	for i := range removeElems {
		de.AgileBoost.Remove(removeElems[i])
	}

	removeElems = make([]*list.Element, 0)
	for e := de.MaxHealthBoost.Front(); e != nil; e = e.Next() {
		boost := e.Value.(BoostTimes)
		boost.CountSteps--
		if boost.CountSteps <= 0 {
			removeElems = append(removeElems, e)
			de.MaxHealth = max(1, de.MaxHealth-boost.BoostValue)
			if de.CurrentHealth > de.MaxHealth {
				de.CurrentHealth = de.MaxHealth
			}
		}
	}
	for i := range removeElems {
		de.MaxHealthBoost.Remove(removeElems[i])
	}

	removeElems = make([]*list.Element, 0)
	for e := de.StrengthBoost.Front(); e != nil; e = e.Next() {
		boost := e.Value.(BoostTimes)
		boost.CountSteps--
		if boost.CountSteps <= 0 {
			removeElems = append(removeElems, e)
			de.Strength = max(1, de.Strength-boost.BoostValue)
		}
	}
	for i := range removeElems {
		de.StrengthBoost.Remove(removeElems[i])
	}

	return commands.AttackCommand{Damage: de.Strength + de.WeaponStrength,
		AttackAgile:    de.Agile,
		StealMaxHealth: 0,
		WithoutMiss:    false,
		IsTranquilize:  false}
}

func (de *damageableEntitiy) snakeGetAttackCommand() commands.AttackCommand {
	isTranqualize := false
	if rand.Intn(10) == 0 {
		isTranqualize = true
	}
	return commands.AttackCommand{Damage: de.Strength, AttackAgile: de.Agile, StealMaxHealth: 0, WithoutMiss: false, IsTranquilize: isTranqualize}
}

func (de *damageableEntitiy) GetHealth() int {
	return de.CurrentHealth
}

func (de *damageableEntitiy) AppendHealth(health int) {
	de.CurrentHealth = min(de.MaxHealth, de.CurrentHealth+health)
}

func (de *damageableEntitiy) IsTranquilize() bool {
	return de.isTranquilize
}

func (de *damageableEntitiy) SetTranquilize(tranqualize bool) {
	de.isTranquilize = tranqualize
}

func (de *damageableEntitiy) GetAgile() int {
	return de.Agile
}

func (de *damageableEntitiy) AppendAgile(agile int) {
	de.Agile += agile
}

func (de *damageableEntitiy) BoostAgile(agile, boostSteps int) {
	de.Agile += agile
	de.AgileBoost.PushBack(BoostTimes{BoostValue: agile, CountSteps: boostSteps})
}

func (de *damageableEntitiy) GetAgileBoosts() *list.List {
	return de.AgileBoost
}

func (de *damageableEntitiy) GetStrength() int {
	return de.Strength + de.WeaponStrength
}

func (de *damageableEntitiy) SetWeaponStrength(weaponStrength int) {
	de.WeaponStrength = weaponStrength
}

func (de *damageableEntitiy) AppendStrength(strength int) {
	de.Strength += strength
}

func (de *damageableEntitiy) BoostStrength(strength, boostSteps int) {
	de.Strength += strength
	de.StrengthBoost.PushBack(BoostTimes{BoostValue: strength, CountSteps: boostSteps})
}

func (de *damageableEntitiy) GetStrengthBoosts() *list.List {
	return de.StrengthBoost
}

func (de *damageableEntitiy) GetMaxHealth() int {
	return de.MaxHealth
}

func (de *damageableEntitiy) AppendMaxHealth(maxHealth int) {
	newMaxHealth := de.MaxHealth + maxHealth
	de.CurrentHealth += maxHealth
	if newMaxHealth < de.CurrentHealth {
		de.CurrentHealth = newMaxHealth
	}
	de.MaxHealth = newMaxHealth
}

func (de *damageableEntitiy) BoostMaxHealth(maxHealth, boostSteps int) {
	de.AppendMaxHealth(maxHealth)
	de.MaxHealthBoost.PushBack(BoostTimes{BoostValue: maxHealth, CountSteps: boostSteps})
}

func (de *damageableEntitiy) GetMaxHealthBoosts() *list.List {
	return de.MaxHealthBoost
}

type vampireDamageableEntitiy struct {
	*damageableEntitiy
	FirstHitDone bool
}

func (v *vampireDamageableEntitiy) vampireTakeDamage(cmd commands.AttackCommand) {
	if v.FirstHitDone {
		v.DefaultTakeDamage(cmd)
	}
	v.FirstHitDone = true
}

func (v *vampireDamageableEntitiy) vampireGetAttackCommand() commands.AttackCommand {
	return commands.AttackCommand{Damage: v.Strength, AttackAgile: v.Agile, StealMaxHealth: 2, WithoutMiss: false, IsTranquilize: false}
}

type orgeDamageableEntitiy struct {
	*damageableEntitiy
	isRestLastStep bool
}

func (v *orgeDamageableEntitiy) orgeGetAttackCommand() commands.AttackCommand {
	var cmd commands.AttackCommand
	if v.isRestLastStep {
		cmd = commands.AttackCommand{Damage: v.Strength, AttackAgile: v.Agile, StealMaxHealth: 0, WithoutMiss: true, IsTranquilize: false}
	} else {
		cmd = commands.AttackCommand{Damage: 0, AttackAgile: 0, StealMaxHealth: 0, WithoutMiss: false, IsTranquilize: false}
	}
	v.isRestLastStep = !v.isRestLastStep
	return cmd
}
