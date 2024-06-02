package easy

import "strings"

func isValid(s string) bool {
	ss := strings.Split(s, "")
	check := make([]string, 0)
	for i := range ss {
		if ss[i] == "{" || ss[i] == "[" || ss[i] == "(" {
			check = append(check, ss[i])
		} else {
			if len(check) <= 0 {
				return false
			}
			switch ss[i] {
			case "}":
				if check[len(check)-1] != "{" {
					return false
				}
			case "]":
				if check[len(check)-1] != "[" {
					return false
				}
			case ")":
				if check[len(check)-1] != "(" {
					return false
				}
			default:
				return false
			}
			check = check[:len(check)-1]
		}
	}
	if len(check) != 0 {
		return false
	}
	return true
}
