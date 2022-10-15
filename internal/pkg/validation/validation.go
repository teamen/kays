package validation

import (
	"reflect"
	"strings"

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

		v.RegisterTranslation("eqfield", trans, func(ut ut.Translator) error {
			err := ut.Add("eqfield", "{0}必须等于{1}", true)
			return err
		}, func(ut ut.Translator, fe validator.FieldError) string {
			feParam := fe.Param()
			if feParamAlias, ok := translationDict[fe.Param()]; ok {
				feParam = feParamAlias
			}
			t, err := ut.T(fe.Tag(), fe.Field(), feParam)
			if err != nil {
				// fmt.Printf("警告: 翻译字段错误: %#v", fe)
				return fe.(error).Error()
			}
			return t
		})
	}
}

func ParseValidationErrors(validationErrors validator.ValidationErrors, request interface{}) ([]APIError, error) {
	out := make([]APIError, len(validationErrors))
	for i, e := range validationErrors {
		// 获取原来的标签 - json
		fieldName := e.StructField()
		// fmt.Println(e.Param())
		t := reflect.TypeOf(request)
		field, _ := t.FieldByName(fieldName)
		j := strings.Split(field.Tag.Get("json"), ",")[0]
		out[i] = APIError{j, e.Translate(trans)}
	}
	return out, nil
}
