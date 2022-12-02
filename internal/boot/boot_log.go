package boot

import (
  "fiber-kit-template/internal/g"
  "github.com/edifierx666/goproject-kit/os/klog"
)

func Logger() {
  cfg := klog.NewLoggerCfg()
  g.GLoggerCfg = cfg
  if g.Cfg().Log.File {
    cfg.SetLogInFile(true)
  }
  g.GLogger = klog.New(g.GLoggerCfg).Build()
}
