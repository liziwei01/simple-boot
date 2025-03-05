/*
 * @Author: liziwei01
 * @Date: 2022-03-03 16:04:06
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-10-30 16:42:00
 * @Description: 读取配置文件, 初始化路由
 */
package bootstrap

import (
	"context"
	"log"

	"simple-boot/library/env"

	"github.com/gin-gonic/gin"
)

const (
	appConfPath = "./conf/app.toml"
)

// AppServer struct.
type AppServer struct {
	Handler *gin.Engine
	Ctx     context.Context
	Config  *Config
	Cancel  context.CancelFunc
}

// Setup 准备.
func Setup() (*AppServer, error) {
	appServer := &AppServer{}
	var (
		err error
	)
	appServer.Config, err = ParserAppConfig(appConfPath)
	if err != nil {
		return nil, err
	}
	env.Default = appServer.Config.Env
	appServer.Ctx, appServer.Cancel = context.WithCancel(context.Background())
	appServer.Handler = InitHandler(appServer)

	return appServer, nil
}

// Start 启动http服务器.
func (appServer *AppServer) Start() {
	defer appServer.Cancel()
	app := NewApp(appServer.Ctx, appServer.Config, appServer.Handler)
	log.Fatalln("server exit:", app.Start())
}

// Start 启动https服务器.
func (appServer *AppServer) StartTLS() {
	defer appServer.Cancel()
	app := NewApp(appServer.Ctx, appServer.Config, appServer.Handler)
	log.Fatalln("server exit:", app.StartTLS())
}
