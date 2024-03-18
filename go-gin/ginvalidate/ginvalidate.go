package ginvalidate

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales"
	en2 "github.com/go-playground/locales/en"
	zh2 "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"sync"
)

type BaseController struct {
}

func (b *BaseController) Validate(ctx *gin.Context, param any) bool {
	local := ctx.DefaultQuery("local", "zh")
	if err := ctx.ShouldBind(param); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok || len(errs) == 0 {
			panic(err)
		}
		tans, ok := uni.GetTranslator(local) // 获取转换的实例
		if !ok {
			fmt.Println("获取语言包实例失败!")
		}
		// errMsg := errs.Translate(tans) // 返回全部错误(map)
		errMsg := errs[0].Translate(tans) // 返回第一个错误
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400, "msg": errMsg, "data": "",
		})
		return false
	}
	return true
}

var (
	uni        *ut.UniversalTranslator
	validate   *validator.Validate
	utLangList []locales.Translator
	_once      sync.Once
)

func init() {
	lazyInit()
}

// 注册语言包
func lazyInit() {
	_once.Do(func() {
		utLangList = []locales.Translator{
			zh2.New(),
			en2.New(),
		}
		uni = ut.New(utLangList[0], utLangList...)
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			validate = v
		} else {
			validate = validator.New()
		}
		registerLang := func() error {
			var translations func(*validator.Validate, ut.Translator) (err error)
			for _, val := range utLangList {
				tans, ok := uni.GetTranslator(val.Locale())
				if !ok {
					return errors.New("获取语言包实例错误")
				}
				switch val.Locale() {
				case "en":
					translations = enTranslations.RegisterDefaultTranslations
				case "zh":
					translations = zhTranslations.RegisterDefaultTranslations
				default:
					translations = enTranslations.RegisterDefaultTranslations
				}
				err := translations(validate, tans)
				if err != nil {
					return err
				}
			}
			return nil
		}
		err := registerLang()
		if err != nil {
			panic(err)
		}
	})
}
