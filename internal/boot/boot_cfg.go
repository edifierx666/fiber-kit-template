package boot

import (
  "fiber-kit-template/internal/g"
  "log"

  "github.com/edifierx666/goproject-kit/os/kcfg"
  "github.com/spf13/pflag"
  "go.uber.org/zap"
)

func KCfg() {
  k := kcfg.New()
  s := pflag.String("config", "", "提供文件地址")
  st := pflag.String("ctype", "", "提供文件类型默认yaml")
  pflag.Parse()

  if *s != "" {
    k.SetConfigPath(*s)
  }
  if *st != "" {
    k.SetConfigType(*st)
  }
  k.SetDefault("appname", g.AppName)
  k.SetDefault("server.port", 8080)
  k.SetDefault("server.address", "0.0.0.0")
  k.SetDefault("server.prefork", false)

  cfg := g.Cfg()
  k.ChangeFn = func(isWatch bool) {
    if err := k.Unmarshal(&cfg); err != nil {
      g.GLogger.Error("解析配置文件出错", zap.Error(err))
    }
    if isWatch {
      g.GLogger.Info("配置文件改变:", zap.Any("config", &cfg))
    }
  }

  if _, err := k.Load(); err != nil {
    log.Printf("配置文件读取失败,%v", err)
  }
}
