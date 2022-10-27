package i18n

import (
	"embed"
	"github.com/aesoper101/kratos-utils/pkg/middleware/localize"
	"github.com/google/wire"
)

var (
	//go:embed locales/locale.*.toml
	fs embed.FS

	ProviderSet = wire.NewSet(NewI18nBundle)
)

func NewI18nBundle() (*localize.I18nBundle, error) {
	return localize.NewBundleFromEmbedFs(fs, "locales/locale.*.toml")
}
