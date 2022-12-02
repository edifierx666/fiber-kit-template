package api

import (
  "fiber-kit-template/internal/controller"
  "fiber-kit-template/internal/middleware"
  "github.com/gofiber/fiber/v2"
)

func RegisterApi(rootGroup fiber.Router) {
  // normal 路由
  normal := rootGroup.Group("/")

  {
    normal.Get("/", controller.Hello)
  }
  // jwt校验 路由
  rootGroup.Group("/app", middleware.JwtMiddleware())

}
