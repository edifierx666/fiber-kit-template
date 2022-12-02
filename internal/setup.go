package app

import (
  "fiber-kit-template/internal/api"
  "fiber-kit-template/internal/boot"
  "fiber-kit-template/internal/g"
  "fiber-kit-template/internal/middleware"
  "fmt"
  "github.com/bytedance/sonic"
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/swagger"
  "go.uber.org/zap"
)

func Run() {
  boot.Init()
  app := fiber.New(fiber.Config{
    Prefork:           g.Cfg().Server.Prefork,
    ServerHeader:      g.Cfg().Appname,
    Immutable:         true,
    ErrorHandler:      middleware.CustomErrorHandle,
    AppName:           g.Cfg().Appname,
    EnablePrintRoutes: true,
    ColorScheme: fiber.Colors{
      Yellow: fiber.DefaultColors.Reset,
    },
    JSONEncoder: sonic.Marshal,
    JSONDecoder: sonic.Unmarshal,
  })

  // 跟路由
  apiGroup := app.Group("/api")

  // 注册swagger
  app.Get("/swagger/*", swagger.HandlerDefault) // default

  // 注册中间件
  middleware.Middleware(app)

  // 注册api
  api.RegisterApi(apiGroup)

  // 启动成功
  app.Hooks().OnListen(func() error {
    g.GLogger.Info("服务器启动成功", zap.Any("配置", g.Cfg()))
    return nil
  })
  // 启动服务器
  if err := app.Listen(fmt.Sprintf("%v:%v", g.Cfg().Server.Address, g.Cfg().Server.Port)); err != nil {
    g.GLogger.Fatal("服务启动错误", zap.Error(err))
  }
}
