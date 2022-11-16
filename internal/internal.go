/*
TISAX Evaluation : mardown to HTML
Copyright (C) 2022 Fabrice THIROUX

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/
*/

package internal

func GetMaturityIcon(maturityLevel int64) string {
	switch maturityLevel {
	case 3:
		return "\u2705"
	case 4:
		return "\u2705\u16ED"
	case 5:
		return "\u2705\u16ED\u16ED"
	default:
		return "\u274C"
	}
}
