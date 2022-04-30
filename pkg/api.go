package api

import (
	"github.com/allanpk716/ChineseSubFinder/internal/logic/file_downloader"
	"github.com/allanpk716/ChineseSubFinder/internal/logic/sub_supplier/zimuku"
	"github.com/allanpk716/ChineseSubFinder/internal/pkg/log_helper"
	"github.com/allanpk716/ChineseSubFinder/internal/pkg/settings"
)

type Zimuku struct {
	*zimuku.Supplier
}

func (c *Zimuku) NewZimuku() {
	c.Supplier = zimuku.NewSupplier(file_downloader.NewFileDownloader(
		settings.GetSettings(),
		log_helper.GetLogger(),
	))
}
