package utils

import (
	"regexp"
	"strconv"
	"time"
	"unicode"
)

func Password(password string) bool {
	var (
		upp, low, num, sym bool
		tot                uint8
	)
 
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return false
		}
	}
 
	if !upp || !low || !num || !sym || tot < 8 {
		return false
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

func Alpha(value string) bool {
	var regexAlpha = "^[a-zA-Z]+$"
	chekAlpha, _ := regexp.MatchString(regexAlpha, value)
	if !chekAlpha {
		return false
	}
	return true
}

func AlphaNum(value string) bool {
	var regexAlphaNum = "^[a-zA-Z0-9]+$"
	chekAlphaNum, _ := regexp.MatchString(regexAlphaNum, value)
	if !chekAlphaNum {
		return false
	}
	return true
}

func AlphaDash(value string) bool {
	var regexAlphaDash = "^[a-zA-Z0-9_-]+$"
	chekAlphaDash, _ := regexp.MatchString(regexAlphaDash, value)
	if !chekAlphaDash {
		return false
	}
	return true
}

func RuleIn(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func Size(param string, value interface{}) bool {
	size, _ := strconv.Atoi(param)
	switch value.(type) {
	case string:
		if len(value.(string)) > size {
			return false
		}
	case int32:
		if value.(int32) > int32(size) {
			return false
		}
	default:		
		return false
	}

	return true
}

func DateFormat(param string, value string) bool {
	format := ""
	switch param {
	case "Y-m-d":
		format = "2002-05-02"
		
	case "Y-m-d H:i:s":
		format = "2002-05-02 14:25:03"
		
	case "Y-m-d H:i":
		format = "2002-05-02 14:25"
	case "H:i":
		format = "14:25"
	case "H:i:s":
		format = "14:25:09"
	default :
		format = ""
	}
   _, err := time.Parse(format, value)
   return err == nil
}


