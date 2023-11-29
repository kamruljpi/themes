package separation

import (
	"io/ioutil"

	"github.com/kamruljpi/go-admin/modules/config"
	adminTemplate "github.com/kamruljpi/go-admin/template"
	"github.com/kamruljpi/go-admin/template/components"
	"github.com/kamruljpi/go-admin/template/types"
	"github.com/kamruljpi/themes/adminlte/resource"
	"github.com/kamruljpi/themes/common"
)

type Theme struct {
	ThemeName string
	components.Base
	*common.BaseTheme
}

var Adminlte = Theme{
	ThemeName: "adminlte_sep",
	Base: components.Base{
		Attribute: types.Attribute{
			TemplateList: common.SepTemplateList,
			Separation:   true,
		},
	},
	BaseTheme: &common.BaseTheme{
		AssetPaths:   resource.AssetPaths,
		TemplateList: common.SepTemplateList,
		Separation:   true,
	},
}

func init() {
	adminTemplate.Add("adminlte_sep", &Adminlte)
}

func Get() *Theme {
	return &Adminlte
}

func (t *Theme) Name() string {
	return t.ThemeName
}

func (t *Theme) GetTmplList() map[string]string {
	return common.SepTemplateList
}

func (t *Theme) GetAsset(path string) ([]byte, error) {
	return ioutil.ReadFile(config.GetAssetRootPath() + path)
}

func (t *Theme) GetAssetList() []string {
	return resource.AssetsList
}
