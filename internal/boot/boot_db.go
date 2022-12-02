package boot

import (
  "context"
  "fiber-kit-template/internal/g"
  "time"

  "go.uber.org/zap"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "gorm.io/gorm/logger"
  "gorm.io/gorm/utils"
)

type LoggerImpl struct {
  LogLevel logger.LogLevel
  sugar    *zap.SugaredLogger
}

func (l *LoggerImpl) LogMode(level logger.LogLevel) logger.Interface {
  newlogger := *l
  newlogger.LogLevel = level
  return &newlogger
}

func (l *LoggerImpl) Info(ctx context.Context, s string, i ...interface{}) {
  l.sugar.Info(s, i)
}

func (l *LoggerImpl) Warn(ctx context.Context, s string, i ...interface{}) {
  l.sugar.Warn(s, i)
}

func (l *LoggerImpl) Error(ctx context.Context, s string, i ...interface{}) {
  l.sugar.Error(s, i)
}

func (l *LoggerImpl) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
  elapsed := time.Since(begin)
  sql, rows := fc()
  if rows == -1 {
    l.sugar.Infof("%s\n[%.3fms] [rows:%v] %s", utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
  } else {
    l.sugar.Infof("%s\n[%.3fms] [rows:%v] %s", utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
  }
}

func DB() {
  dbCfg := g.Cfg().GetDBCfg()
  l := &LoggerImpl{
    LogLevel: logger.Info,
    sugar:    g.GLogger.Sugar(),
  }
  db, err := gorm.Open(mysql.Open(dbCfg.Link), &gorm.Config{
    Logger: l,
  })
  if err != nil {
    g.GLogger.Fatal("数据库初始化失败", zap.Error(err))
  }
  g.GDb = db
  if dbCfg.Debug {
    db.Debug()
  }
  g.GLogger.Info("数据库初始化成功")
}
