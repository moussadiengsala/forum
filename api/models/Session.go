package models

import "time"

type Session struct {
	Token          string `json:"token"`
	ExpirationDate time.Time `json:"expiration"`
	UserID         string `json:"user_id"`
	CreationDate   time.Time `json:"creation_date"`
}
