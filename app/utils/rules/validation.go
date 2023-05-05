package rules

import "regexp"

func Password(password string) bool {
	tests := []string{".{7,}", "[a-z]", "[A-Z]", "[0-9]", "[^\\d\\w]"}
	for _, test := range tests {
		t, _ := regexp.MatchString(test, password)
		if !t {
			return false
		}
	}
	return true
}

func Nik(nik string) bool {
	var RegexNik = "^(1[1-9]|21|[37][1-6]|5[1-3]|6[1-5]|[89][12])\\d{2}\\d{2}([04][1-9]|[1256][0-9]|[37][01])(0[1-9]|1[0-2])\\d{2}\\d{4}$"
	chekNik, _ := regexp.MatchString(RegexNik, nik)
	if !chekNik {
		return false
	}
	return true
}