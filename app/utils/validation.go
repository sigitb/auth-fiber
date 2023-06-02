package utils

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/go-playground/validator/v10"
)


type ErrorResponse struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

var validate = validator.New()

func ValidateRequest(request interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	validate.RegisterValidation("in", func(fl validator.FieldLevel) bool {
		allowedValues := strings.Split(fl.Param(), "+")
		return RuleIn(allowedValues, fl.Field().String())
	})
	validate.RegisterValidation("size", func(fl validator.FieldLevel) bool {
		return Size(fl.Param(), fl.Field().String())
	})
	validate.RegisterValidation("date_format", func(fl validator.FieldLevel) bool {
		return Size(fl.Param(), fl.Field().String())
	})
	validate.RegisterValidation("password-custom", func(fl validator.FieldLevel) bool {
		return Password(fl.Field().String())
	})
	validate.RegisterValidation("nik", func(fl validator.FieldLevel) bool {
		return Nik(fl.Field().String())
	})
	validate.RegisterValidation("alpha", func(fl validator.FieldLevel) bool {
		return Alpha(fl.Field().String())
	})
	validate.RegisterValidation("alpha_num", func(fl validator.FieldLevel) bool {
		return AlphaNum(fl.Field().String())
	})
	validate.RegisterValidation("alpha_dash", func(fl validator.FieldLevel) bool {
		return AlphaDash(fl.Field().String())
	})
	err := validate.Struct(request)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			textErr := getValueValidation(err.Tag())
			value := getReplaceValueValidation(textErr, err.StructField(), err.Param())
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = value
			errors = append(errors, &element)
		}
	}
	return errors
}

func getReplaceValueValidation(value string, attr string , field string) string {
	textAttribute := strings.ReplaceAll(value, ":attribute", attr)
	textOther := strings.ReplaceAll(textAttribute, ":other", field)
	textFormat := strings.ReplaceAll(textOther, ":format", field)
	textDate := strings.ReplaceAll(textFormat, ":date", field)
	textSize := strings.ReplaceAll(textDate, ":size", field)
	return textSize
}


func getValueValidation(name string, lang ...string) string {
	// Baca file JSON
	language := ""
	if len(lang) == 0  {
		language = "en";
	}
	
	var nameFile = ""
	if language == "en" {
		nameFile = "./app/utils/validation/en.json"
	}else if language == "id"{
		nameFile = "./app/utils/validation/id.json"
	}


	data, err := ioutil.ReadFile(nameFile)
	if err != nil {
		Log("Failed write file:"+err.Error(), "err", "validation")
		return ""
	}

	// Dekode JSON ke dalam struktur data
	var jsonData map[string]interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		Log("Failed JSON:"+err.Error(), "err", "validation")
		return ""
	}

	// Mengambil nilai berdasarkan kunci (key)
	value := jsonData[name]

	return value.(string)
}