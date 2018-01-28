package main

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

func main() {
	server := iris.New()

	server.Logger().SetLevel("debug")
	mvc.Configure(server.Party("/#"), rootMVC)

	server.Run(iris.Addr(":8080"))
}

// types & funcs used w/in rootMVC
type rootController struct {
	Logger  LoggerService
	Session *sessions.Session
}

type LoggerService interface {
	Log(string)
}

func (receiver *prefixedLogger) Log(msg string) {
	fmt.Printf("%s: %s\n", receiver.prefix, msg)
}

type prefixedLogger struct {
	prefix string
}

// prob. best to "separate concerns"
func (c *rootController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/custom", "Custom")
}

func (c *rootController) AfterActivation(a mvc.AfterActivation) {
	if a.Singleton() {
		// WHAT IS A SINGLETON? prob. in Tour of Go.
		panic("rootController should be stateless, a request-scoped, we have a 'Session' which depends on the context.")
		// HOW DO I BREAK THIS STRING UP INTO MULTIPLE LINES?
	}
}

func (c *rootController) Get() string {
	counter := c.Session.Increment("count", 1)
	body := fmt.Sprintf(
		"Hello from rootController\nTotal visits from you: %d", counter)

	c.Logger.Log(body)

	return body
}

func (c *rootController) Custom() string {
	return "custom"
}

type rootSubController struct {
	Session *sessions.Session
}

func (c *rootSubController) Get() string {
	counter, _ := c.Session.GetIntDefault("count", 1)
	return fmt.Sprintf(
		"Hello from rootSubController.\nRead-only visits count: %d",
		counter)
}

/*
If a controller's fields (or even its functions) expect an interface
  but a struct value is binded, then it will check
  if that struct value implements the interface.

If true, then it will add this to the available bindings
  before the server runs.
*/

func rootMVC(server *mvc.Application) {
	// You can use normal middlewares at MVC apps of course.
	server.Router.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Path: %s", ctx.Path())
		ctx.Next()
	})

	// Register dependencies which will be binding to the controller(s),
	// can be either a function which accepts an iris.Context and returns a single value (dynamic binding)
	// or a static struct value (service).
	server.Register(
		sessions.New(sessions.Config{}).Start,
		&prefixedLogger{prefix: "DEV"},
	)

	// GET: http://localhost:8080/#
	// GET: http://localhost:8080/#/custom
	server.Handle(new(rootController))

	// All dependencies of the parent *mvc.Application
	// are cloned to this new child;
	// thefore it has access to the same session as well.
	// GET: http://localhost:8080/#/sub
	server.Party("/sub").
		Handle(new(rootSubController))
}

// go run main.go
