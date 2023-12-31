package app

import (
	"context"
	"fmt"
	"github.com/huerni/gmitex/pkg/etcd"
	"github.com/huerni/miao-gateway/internal/config"
	"os"
	"os/signal"
	"syscall"
)

type GwServer struct {
	Cfg *config.Config
}

func NewGmServer(c *config.Config) *GwServer {

	return &GwServer{
		Cfg: c,
	}
}

func (g *GwServer) Start(ctx context.Context) {

	g.RegisterComponents(ctx)

}

func (g *GwServer) RegisterComponents(ctx context.Context) {
	if g.Cfg.Etcd.HasConfig() {
		err := etcd.PutWithInfo(context.Background(), g.Cfg.Etcd.Hosts, &etcd.ServerInfo{
			ServerKey: g.Cfg.Etcd.Key,
			Data:      nil,
		})

		if err != nil {
			fmt.Println(err)
		}
	}

	if g.Cfg.Mysql.HasConfig() {
		err := etcd.PutWithInfo(context.Background(), g.Cfg.Etcd.Hosts, &etcd.ServerInfo{
			ServerKey: fmt.Sprintf("%s-%s", g.Cfg.Prefix, "mysql-master"),
			Data:      map[string]string{"dsn": g.Cfg.Mysql.DSN},
		})

		if err != nil {
			fmt.Println(err)
		}
	}

	if g.Cfg.MysqlSalve.HasConfig() {
		err := etcd.PutWithInfo(context.Background(), g.Cfg.Etcd.Hosts, &etcd.ServerInfo{
			ServerKey: fmt.Sprintf("%s-%s", g.Cfg.Prefix, "mysql-slave"),
			Data:      map[string]string{"dsn": g.Cfg.MysqlSalve.DSN},
		})

		if err != nil {
			fmt.Println(err)
		}
	}

}

func (g *GwServer) WaitForShutdown(ctx context.Context) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
