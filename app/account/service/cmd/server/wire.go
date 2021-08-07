// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/xun1009/xun_go_blog/app/account/service/internal/biz"
	"github.com/xun1009/xun_go_blog/app/account/service/internal/conf"
	"github.com/xun1009/xun_go_blog/app/account/service/internal/data"
	"github.com/xun1009/xun_go_blog/app/account/service/internal/server"
	"github.com/xun1009/xun_go_blog/app/account/service/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
