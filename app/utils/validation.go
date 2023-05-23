package utils

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/go-playground/validator/v10"
)


type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateRequest(FindEmail interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	validate.RegisterValidation("in", func(fl validator.FieldLevel) bool {
		allowedValues := strings.Split(fl.Param(), "+")
		return RuleIn(allowedValues, fl.Field().String())
	})
	err := validate.Struct(FindEmail)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {

			textErr := getValueValidation(err.Tag())
			textAttribute := strings.ReplaceAll(textErr, ":attribute", "Parameter")
			textOther := strings.ReplaceAll(textAttribute, ":other", err.Field())
			textFormat := strings.ReplaceAll(textOther, ":format", err.Field())
			textDate := strings.ReplaceAll(textFormat, ":date", err.Field())
			value := textDate
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = value
			errors = append(errors, &element)
		}
	}
	return errors
}


func getValueValidation(name string, lang ...string) string {
	// Baca file JSON
	language := ""
	if len(lang) > 0  {
		language = "en";
	}
	
	var nameFile = ""
	if language == "en" {
		nameFile = "./app/utils/validation/en.json"
	}else if language == "in"{
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