package model

type Power struct {
	ActivePower int `gorm:"active_power" json:"active_power"`
	PowerInput  int `gorm:"power_input" json:"power_input"`
}
