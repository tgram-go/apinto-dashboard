package initialize

import (
	_ "embed"
	"github.com/eolinker/apinto-dashboard/store"
)

var (
	//go:embed navigation.yml
	navigationContent []byte
)

func initNavigation(db store.IDB) {
	// 初始化导航
}
