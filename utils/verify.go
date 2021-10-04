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

package utils

import (
	"regexp"
)

const (
	EMAIL_VERIFY_PATTERN = `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
)

// VerifyEmailFormat verify email format
func VerifyEmailFormat(email string) bool {
	if email == "" {
		return false
	}
	reg := regexp.MustCompile(EMAIL_VERIFY_PATTERN)
	return reg.MatchString(email)
}
