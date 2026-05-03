package dto

import (
	"fmt"
	"s21_rogue/internal/domain/value_objects"
)

type ItemBehaviour interface {
	GetUUID() string
	GetLabel() string
	ToString() string
	IsSelected() bool
}

type FoodDTO struct {
	UUID   string `json:"uuid"`
	Label  string `json:"label"`
	Health int    `json:"health"`
}

func (i FoodDTO) IsSelected() bool {
	return false
}

func (i FoodDTO) GetUUID() string {
	return i.UUID
}

func (i FoodDTO) GetLabel() string {
	return i.Label
}

func (f FoodDTO) ToString() string {
	return fmt.Sprintf("%v {restores health - %v}", f.Label, f.Health)
}

type ScrollDTO struct {
	UUID      string `json:"uuid"`
	Label     string `json:"label"`
	BoostType string `json:"boost_type"`
	Value     int    `json:"value"`
}

func (i ScrollDTO) IsSelected() bool {
	return false
}

func (i ScrollDTO) GetUUID() string {
	return i.UUID
}

func (i ScrollDTO) GetLabel() string {
	return i.Label
}

func (i ScrollDTO) ToString() string {
	return fmt.Sprintf("%v {%v boost - %v}", i.Label, i.BoostType, i.Value)
}

type ElixirDTO struct {
	UUID       string `json:"uuid"`
	Label      string `json:"label"`
	BoostType  string `json:"boost_type"`
	Value      int    `json:"value"`
	CountSteps int    `json:"steps"`
}

func (i ElixirDTO) IsSelected() bool {
	return false
}

func (i ElixirDTO) GetUUID() string {
	return i.UUID
}

func (i ElixirDTO) GetLabel() string {
	return i.Label
}

func (i ElixirDTO) ToString() string {
	return fmt.Sprintf("%v {%v boost - %v}", i.Label, i.BoostType, i.Value)
}

type WeaponDTO struct {
	UUID     string `json:"uuid"`
	Label    string `json:"label"`
	Strength int    `json:"strength"`
	Selected bool   `json:"selected"`
}

func (i WeaponDTO) IsSelected() bool {
	return i.Selected
}

func (i WeaponDTO) GetUUID() string {
	return i.UUID
}

func (i WeaponDTO) GetLabel() string {
	return i.Label
}

func (i WeaponDTO) ToString() string {
	return fmt.Sprintf("%v {boost strength - %v}", i.Label, i.Strength)
}

type TreasureDTO struct {
	UUID  string `json:"uuid"`
	Label string `json:"label"`
	Score int    `json:"score"`
}

func (i TreasureDTO) IsSelected() bool {
	return false
}

func (i TreasureDTO) GetUUID() string {
	return i.UUID
}

func (i TreasureDTO) GetLabel() string {
	return i.Label
}

func (i TreasureDTO) ToString() string {
	return fmt.Sprintf("%v - %v score", i.Label, i.Score)
}

type FoodPositionDTO struct {
	FoodDTO  FoodDTO                `json:"food"`
	Position value_objects.Position `json:"position"`
}

type ScrollPositionDTO struct {
	ScrollDTO ScrollDTO              `json:"scroll"`
	Position  value_objects.Position `json:"position"`
}

type ElixirPositionDTO struct {
	ElixirDTO ElixirDTO              `json:"elixir"`
	Position  value_objects.Position `json:"position"`
}

type WeaponPositionDTO struct {
	WeaponDTO WeaponDTO              `json:"weapon"`
	Position  value_objects.Position `json:"position"`
}

type TreasurePositionDTO struct {
	TreasureDTO TreasureDTO            `json:"treasure"`
	Position    value_objects.Position `json:"position"`
}
