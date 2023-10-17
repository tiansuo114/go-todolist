package main

import (
	"api-gateway/config"
	"api-gateway/discovery"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc/resolver"
)

func main() {
	config.InitConfig()

	// 服务发现
	etcdAddress := []string{viper.GetString("etcd.address")}
	etcdRegister := discovery.NewResolver(etcdAddress, logrus.New())
	resolver.Register(etcdRegister)

}
