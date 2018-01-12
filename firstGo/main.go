// https://iris-go.com/v10/start
//go run main.go ; Open localhost:8080 in browser
package main

import (
  "github.com/kataras/iris"
  "github.com/kataras/iris/middleware/logger"
  "github.com/kataras/iris/middleware/recover"
)

func main() {
  app := iris.New()
  app.Logger().SetLevel("debug")

  // recover from any http-relative panics
  app.Use(recover.New())
  // log the requests to the terminal
  app.Use(logger.New())

  // Resource: http://localhost:8080
  app.Handle("GET", "/", func(ctx iris.Context) {
    ctx.HTML("<p>There we go!</p>")
  })

  // same as app.Handle("GET", "/ping", [...])
  app.Get("/ping", func(ctx iris.Context) {
    ctx.WriteString("pong")
  }) //in React util.js, GET request to "/ping" w/o payload

  app.Get("/hello", func(ctx iris.Context) {
    ctx.JSON(iris.Map{"message": "Hello from Iris!"})
  })

  app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
