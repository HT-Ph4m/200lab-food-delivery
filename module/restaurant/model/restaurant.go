package restaurantmodel

import (
	"200lab-project-1/common"
	"errors"
	"strings"
)

type RestaurantType string

const (
	TypeNormal  RestaurantType = "normal"
	TypePremium RestaurantType = "premium"
	EntityName                 = "Restaurant"
)

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name;"` //tag
	Addr            string         `json:"addr" gorm:"column:addr;"` //tag
	Type            RestaurantType `json:"type" gorm:"column:type;"` //tag
}

func (data *Restaurant) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)
}

var (
	ErrNameIsEmpty = errors.New("Name cannot be empty")
)

func (Restaurant) TableName() string { return "restaurants" }

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;"` //tag
	Addr            string `json:"addr" gorm:"column:addr;"` //tag
}

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)

	if data.Name == "" {
		return ErrNameIsEmpty
	}
	return nil
}

func (data *RestaurantCreate) Mask(isAdminOrOner bool) {
	data.GenUID(common.DbTypeRestaurant)
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

type RestaurantUpdate struct {
	common.SQLModel `json:",inline"`
	Name            *string `json:"name" gorm:"column:name;"` //tag
	Addr            *string `json:"addr" gorm:"column:addr;"` //tag
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }
