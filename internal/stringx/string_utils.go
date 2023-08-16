package stringx

import (
	"strings"
)

func CamelToSnake(s string) string {
	var sb strings.Builder
	for i := 0; i < len(s); i++ {
		c := s[i]
		if IsUpper(c) {
			if i > 0 {
				sb.WriteByte('_')
			}
			sb.WriteByte(ToLower(c))
			j := i + 1
			for ; j < len(s); j++ {
				nc := s[j]
				if !IsUpper(nc) {
					j--
					break
				}
			}
			j--
			if j > i {
				for k := i + 1; k <= j; k++ {
					sb.WriteByte(ToLower(s[k]))
				}
				i = j
			}
		} else {
			sb.WriteByte(c)
		}
	}

	return sb.String()
}

// IsUpper return if v is a upper case ascii char.
func IsUpper(v byte) bool {
	return 'A' <= v && v <= 'Z'
}

// IsLower return if v is a lower case ascii char.
func IsLower(v byte) bool {
	return 'a' <= v && v <= 'z'
}

// ToLower convert ascii char to upper case
// Return v if v is not a upper case ascii character.
func ToLower(v byte) byte {
	if IsUpper(v) {
		return v - 'A' + 'a'
	}
	return v
}

func ToUpper(v byte) byte {
	if IsLower(v) {
		return v - 'a' + 'A'
	}
	return v
}

// Capitalize return str with first char of ascii str upper case.
func Capitalize(str string) string {
	if str == "" {
		return str
	}
	if !IsLower(str[0]) {
		return str
	}
	bytes := []byte(str)
	bytes[0] = ToUpper(str[0])
	return string(bytes)
}

// DeCapitalize return str with first char of ascii str lower case.
func DeCapitalize(str string) string {
	if str == "" {
		return str
	}
	if !IsUpper(str[0]) {
		return str
	}
	bytes := []byte(str)
	bytes[0] = ToLower(str[0])
	return string(bytes)
}
