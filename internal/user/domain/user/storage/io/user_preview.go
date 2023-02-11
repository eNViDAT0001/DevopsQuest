package io

import "github.com/eNViDAT0001/GolangAdventure/internal/user/entities"

type UserPreview struct {
	ID       uint
	Username string
	Email    string
	Phone    string
	Type     entities.UserType
}
