package main

import app "fiber-kit-template/internal"

// Run @title  Swagger Example API
// @version 1.0
// @description
// @BasePath                   /api
// @securityDefinitions.apikey ApiKeyAuth
// @in                         header
// @name                       token
func main() {
  app.Run()
}
