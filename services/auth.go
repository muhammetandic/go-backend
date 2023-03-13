package services

import (
	"fmt"

	"github.com/muhammetandic/go-backend/main/db/model"
)

func Login(info model.Auth) error {
	if info.Username == "m_andic@hotmail.com" && info.Password == "deneme" {
		return nil
	}
	return fmt.Errorf("login incorrect")
}
