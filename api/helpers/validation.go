package helpers

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
)

type (
	StdRespError struct {
		Message string `json:"message"`
		Field   string `json:"field"`
	}
	StdResp struct {
		Code    string         `json:"code,omitempty"`
		Message string         `json:"message,omitempty"`
		Errors  []StdRespError `json:"errors,omitempty"`
	}
)

func BindJSON(c *gin.Context, obj interface{}) error {
	err := c.ShouldBindJSON(obj)
	if err == nil {
		return nil
	}
	elem := reflect.TypeOf(obj).Elem()
	switch err.(type) {
	case validator.ValidationErrors:
		errors := make([]StdRespError, 0)
		for _, fieldErr := range err.(validator.ValidationErrors) {
			text := ""
			field, ok := elem.FieldByName(fieldErr.Field)
			if !ok {
				continue
			}
			tagValue := field.Tag.Get("json")
			if tagValue == "" {
				tagValue = fieldErr.Field
			}
			switch fieldErr.Tag {
			case "required":
				text = "Cannot be blank"
			case "email":
				text = "Is not a valid email address"
			case "min":
				if fieldErr.Type.Name() == "string" {
					i, _ := strconv.Atoi(vMethodArg(field, "min"))
					text = fmt.Sprintf("Should contain at least %d characters", i)
				}
			case "eqfield":
				rel := ""
				if arg := vMethodArg(field, "eqfield"); arg != "" {
					if field, ok := elem.FieldByName(arg); ok {
						rel = field.Tag.Get("json")
					}
				}
				text = fmt.Sprintf("Must be equal to \"%s\"", rel)
			case "lte":
				text = fmt.Sprintf("Must be less than or equal to \"%s\"", vMethodArg(field, "lte"))
			case "gte":
				text = fmt.Sprintf("Must be greater than or equal to \"%s\"", vMethodArg(field, "gte"))
			}
			if text != "" {
				errors = append(errors, StdRespError{text, tagValue})
			}
		}
		c.JSON(http.StatusBadRequest, StdResp{Errors: errors})
		return fmt.Errorf("Input data %+v are not valid", obj)
	default:
		c.JSON(http.StatusBadRequest, StdResp{Code: "PARSE_REQUEST_BODY", Message: err.Error()})
		return err
	}
}

func vMethodArg(field reflect.StructField, vMethod string) string {
	tagValue := field.Tag.Get("binding")
	i, l := 0, len(vMethod)+1
	if i = strings.Index(tagValue, vMethod+"="); i == -1 {
		return ""
	}
	if j := strings.Index(tagValue[i+l:], ","); j != -1 {
		return tagValue[i+l : i+l+j]
	}
	return tagValue[i+l:]
}

func InEnums(str string, enums []string) bool {
	for _, enum := range enums {
		if enum == str {
			return true
		}
	}
	return false
}
