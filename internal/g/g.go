package g

import (
  "github.com/edifierx666/goproject-kit/os/klog"
  "gorm.io/gorm"
)

const AppName = "pet-blind"

// 日志

var (
  GLoggerCfg *klog.LoggerCfg
  GLogger    *klog.Logger
)

// 配置

var (
  GSCfg = &AppConfig{}
)

// 数据库

var (
  GDb *gorm.DB
)

func DB() *gorm.DB {
  return GDb
}

func Cfg() *AppConfig {
  return GSCfg
}

func Log() *klog.Logger {
  return GLogger
}
