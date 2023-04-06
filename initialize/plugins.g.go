package initialize

import (
	"embed"
	"github.com/eolinker/apinto-dashboard/store"
)

var (
	//go:embed plugins
	pluginDir embed.FS
)

func InitPlugins(idb store.IDB) {
	// todo 初始化插件
	pluginDir.ReadDir(".")
}
