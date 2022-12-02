package middleware

import (
  "fiber-kit-template/internal/consts"
  "fiber-kit-template/internal/g"
  "fiber-kit-template/internal/model/resp"
  "fmt"
  "time"

  "github.com/edifierx666/fiber-kit/middleware"
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/cors"
  recover2 "github.com/gofiber/fiber/v2/middleware/recover"
  "github.com/gofiber/fiber/v2/middleware/requestid"
  "go.uber.org/zap"
)

func Middleware(app *fiber.App) {
  app.Use(
    recover2.New(recover2.Config{
      EnableStackTrace: true,
    }),
    cors.New(),
    requestid.New(),
    BaseLogger,
  )
}
func JwtMiddleware() fiber.Handler {
  jwtMiddleWare := middleware.JWT(
    &middleware.JWTConfig{
      ErrorHandlerWithContext: func(err error, c *fiber.Ctx) error {
        return resp.FailWithMessage(err.Error(), c)
      },
      TokenLookup: "header:token",
      ContextKey:  consts.CtxJWTKey,
      SigningKey:  []byte(g.Cfg().Jwt.Key),
    },
  )
  return jwtMiddleWare
}
func CustomErrorHandle(c *fiber.Ctx, err error) error {
  if e, ok := err.(*fiber.Error); ok {
    return resp.FailWithDetailed(e.Code, nil, e.Message, c)
  }
  if respE, ok := err.(resp.IErrorCode); ok {
    return resp.FailWithDetailed(respE.StatusCode(), respE.Data(), respE.Error(), c)
  }
  return resp.FailWithMessage(err.Error(), c)
}

func BaseLogger(c *fiber.Ctx) error {
  now := time.Now()
  err := c.Next()
  uri := c.Request().URI()
  // 日志打印
  zapFields := []zap.Field{
    zap.Any("URI", string(uri.Path())),
    zap.String("QUERY", string(uri.QueryString())),
  }
  rid := c.Locals("requestid")
  if err != nil {
    zapFields = append(zapFields, zap.Error(err))
    formFile, hasFile := c.MultipartForm()
    if hasFile == nil {
      zapFields = append(zapFields, zap.Any("file", formFile.File))
    } else {
      zapFields = append(zapFields, zap.Any("BODY", string(c.Body())))
    }
    g.GLogger.Error(
      fmt.Sprintf("[%v] [%v]", rid, c.Method()),
      zapFields...,
    )
  } else {
    formFile, hasFile := c.MultipartForm()
    if hasFile == nil {
      zapFields = append(zapFields, zap.Any("file", formFile.File))
    } else {
      zapFields = append(zapFields, zap.Any("BODY", string(c.Body())))
    }
    zapFields = append(zapFields, zap.Any("IP", c.IP()),
      zap.String("USER-AGENT", string(c.Request().Header.UserAgent())),
      zap.String("耗时", fmt.Sprintf("%v", time.Since(now))),
    )
    g.GLogger.Info(
      fmt.Sprintf("[%v] [%v]", rid, c.Method()),
      zapFields...,
    )
  }
  // 日志打印end
  return err
}
