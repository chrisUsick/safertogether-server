package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var filepath string = "resources/messages/"

var bundle *i18n.Bundle

func init() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.LoadMessageFile("resources/messages/en.toml")
}
func NewLocalizer(lang string, accept string) func(string) string {
	return func(message string) string {
		localizer := i18n.NewLocalizer(bundle, lang, accept)
		localizeConfig := i18n.LocalizeConfig{
			MessageID: message,
		}
		msg, err := localizer.Localize(&localizeConfig)
		if err != nil {
			fmt.Printf("Error! %v\n", err)
		}
		return msg
	}
}

func NewLocalizerFromContext(c *gin.Context) func(string) string {
	lang := c.Request.FormValue("lang")
	accept := c.GetHeader("Accept-Language")
	return NewLocalizer(lang, accept)
}
