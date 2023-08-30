package main

import (
	"fmt"
	"github.com/eolinker/apinto-dashboard/config"
	grpcservice "github.com/eolinker/apinto-dashboard/grpc-service"
	apintomodule "github.com/eolinker/apinto-dashboard/module"
	"github.com/eolinker/apinto-dashboard/modules/grpc-service/service"
	"github.com/eolinker/apinto-dashboard/modules/module-plugin/embed_registry"
	"github.com/eolinker/apinto-dashboard/modules/notice"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"

	"github.com/eolinker/apinto-dashboard/modules/core"
	"github.com/eolinker/apinto-dashboard/modules/plugin/plugin_timer"

	"github.com/eolinker/apinto-dashboard/app/apserver/version"
	_ "github.com/eolinker/apinto-dashboard/report"
	"github.com/eolinker/eosc/common/bean"
	"github.com/eolinker/eosc/log"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:     "apserver",
		HelpName: "apserver",
		Usage:    "apinto dashboard",

		Version:     version.Version,
		Description: "",
		Commands:    []*cli.Command{version.Build()},
		Flags:       nil,
		Action: func(context *cli.Context) error {
			run()
			return nil
		},
	}
	_ = app.Run(os.Args)
}

func run() {
	config.ReadConfig()
	initLog()
	gin.SetMode(gin.ReleaseMode)

	var coreService core.ICore
	bean.Autowired(&coreService)
	var front core.EngineCreate = new(Front)
	bean.Injection(&front)
	config.InitDb()
	config.InitRedis()
	err := bean.Check()
	if err != nil {
		log.Fatal(err)
	}

	// 执行内置插件初始化
	err = embed_registry.InitEmbedPlugins()
	if err != nil {
		log.Fatal(err)
	}
	_ = coreService.ReloadModule()
	go plugin_timer.ExtenderTimer()

	//初始化通知渠道驱动
	initNoticeChannelDriver()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GetPort()))
	if err != nil {
		panic(err)
	}
	// Create a cmux.
	m := cmux.New(listener)
	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldPrefixSendSettings("content-type", "application/grpc"))

	httpL := m.Match(cmux.HTTP1Fast())

	httpServer := &http.Server{Handler: coreService}
	grpcServer := grpc.NewServer()

	grpcservice.RegisterGetConsoleInfoServer(grpcServer, service.NewConsoleInfoService())
	grpcservice.RegisterNoticeSendServer(grpcServer, service.NewNoticeSendService())

	console := newConsoleServer(httpServer, grpcServer)
	go func() {
		err := httpServer.Serve(httpL)
		if err != nil {
			log.Error("listen httpServer error: ", err)
		}
	}()
	go func() {
		err := grpcServer.Serve(grpcL)
		if err != nil {
			log.Error("listen grpcServer error: ", err)
		}
	}()
	go func() {
		err := m.Serve()
		if err != nil {
			log.Error("server close: ", err)
			return
		}
	}()
	err = console.Wait()
	if err != nil {
		log.Fatal(err)
	}

}

type Front struct {
}

func (f *Front) CreateEngine() *gin.Engine {
	engine := gin.Default()
	engine.Use(apintomodule.SetRepeatReader)
	return engine
}

func initNoticeChannelDriver() {
	var noticeChannelService notice.INoticeChannelService
	bean.Autowired(&noticeChannelService)
	err := noticeChannelService.InitChannelDriver()
	if err != nil {
		panic(err)
	}
}
