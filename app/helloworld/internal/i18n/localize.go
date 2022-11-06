package i18n

import (
	"embed"
	"github.com/aesoper101/kratos-utils/pkg/middleware/translator"
	ut "github.com/go-playground/universal-translator"
	"github.com/google/wire"
)

var (
	//go:embed locales/*
	fs embed.FS

	ProviderSet = wire.NewSet(NewTranslator)
)

func NewTranslator() (*ut.UniversalTranslator, error) {
	return translator.NewTranslatorFromFs(fs, "locales")
}
