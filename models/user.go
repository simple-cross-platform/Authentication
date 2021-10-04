/*
 * This file is part of the Authentication distribution (https://github.com/simple-cross-platform/Authentication).
 * Copyright (c) 2021 Aimer Neige.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, version 3.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 * General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthenticationUser struct {
	gorm.Model
	Account  string
	Password string
	Email    string
}

type AuthenticationUserDto struct {
	ID       uint   `json:"id"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (a AuthenticationUser) ToDto() (dto AuthenticationUserDto) {
	dto.ID = a.ID
	dto.Account = a.Account
	dto.Password = a.Password
	dto.Email = a.Email
	return
}

func (a AuthenticationUser) CheckPassword(password string) (ret bool) {
	ret = false
	if err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password)); err == nil {
		ret = true
	}
	return
}

func (a *AuthenticationUser) GeneratePassword(password string) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err == nil {
		(*a).Password = string(hashPassword)
		return nil
	}
	return err
}
