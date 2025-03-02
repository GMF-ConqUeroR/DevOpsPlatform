package start

import (
	"cmdb/conf"
	"cmdb/logger"
	"cmdb/protocol"
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/infraboard/mcube/ioc"
	"github.com/spf13/cobra"

	_ "cmdb/apps"
)

var (
	configType string
	configFile string
)

var CMD = &cobra.Command{
	Use:   "start",
	Short: "cmdb Center",
	Long: `
		用户资产管理模块。
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// 加载全局配置
		switch configType {
		case "env":
			_, err := conf.LoadConfigFromEnv()
			cobra.CheckErr(err)
		default:
			_, err := conf.LoadConfigFromToml(configFile)
			cobra.CheckErr(err)
		}

		// 启动 http API
		httpServer := protocol.NewHttp()
		go func() {
			httpError := httpServer.Start()
			if httpError != nil {
				cobra.CheckErr(httpError)
			}
		}()
		// 启动 GRPC API
		grpcServer := protocol.NewGrpc()
		go func() {
			grpcError := grpcServer.Start()
			if grpcError != nil {
				cobra.CheckErr(grpcError)
			}
		}()

		err := ioc.InitIocObject()
		cobra.CheckErr(err)

		// 负责阻塞住主进行(Loop Os Signal)
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)

		s := <-ch
		logger.L().Debug().Msgf("received os signal: %s, exit ...", s)
		ctx := context.Background()
		httpServer.Stop(ctx)
		grpcServer.Stop(ctx)

		conf.C().MySQL.Close()
	},
}

func init() {
	CMD.Flags().StringVarP(&configType, "config-type", "t", "file", "程序加载配置的方式")
	CMD.Flags().StringVarP(&configFile, "config-file", "f", "etc/config.toml", "配置文件的路径")
}
