package validation

import (
	"reflect"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

type APIError struct {
	Param   string
	Message string
}

func RegisterTranslations() {

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zh := zh.New()
		uni := ut.New(zh, zh)
		trans, _ = uni.GetTranslator("zh")
		zh_translations.RegisterDefaultTranslations(v, trans)

		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			return field.Tag.Get("comment")
		})
	}
}

func ParseValidationErrors(validationErrors validator.ValidationErrors, request interface{}) ([]APIError, error) {
	out := make([]APIError, len(validationErrors))
	for i, e := range validationErrors {
		// 获取原来的标签 - json
		fieldName := e.StructField()
		t := reflect.TypeOf(request)
		field, _ := t.FieldByName(fieldName)
		j := field.Tag.Get("json")
		out[i] = APIError{j, e.Translate(trans)}
	}
	return out, nil
}
