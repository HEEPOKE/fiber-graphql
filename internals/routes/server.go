package routes

import (
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/HEEPOKE/fiber-graphql/internals/app/resolvers"
	"github.com/HEEPOKE/fiber-graphql/internals/domains/generated"
	"github.com/HEEPOKE/fiber-graphql/pkg/configs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"gorm.io/gorm"
)

type Server struct {
	fib *fiber.App
	db  *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	version := fmt.Sprintf("App v%s", configs.Cfg.VERSION)
	app := fiber.New(fiber.Config{
		ServerHeader:             "Fiber",
		AppName:                  version,
		CaseSensitive:            true,
		StrictRouting:            true,
		EnablePrintRoutes:        true,
		EnableSplittingOnParsers: true,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PATCH, DELETE",
	}))
	app.Use(helmet.New())
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format:     "[${ip}]:${port} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Bangkok",
	}))

	return &Server{
		fib: app,
		db:  db,
	}
}

func adaptHTTPHandler(h http.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		handler := fasthttpadaptor.NewFastHTTPHandler(h)
		handler(c.Context())
		return nil
	}
}

func (s *Server) Init() *fiber.App {
	basicAuthMiddleware := basicauth.Config{
		Users: map[string]string{
			"admin": "        ",
		},
	}

	apis := s.fib.Group("/apis")
	apis.Get("/monitor", basicauth.New(basicAuthMiddleware), monitor.New(monitor.Config{Title: "Monitor Page"}))

	gqlHandler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}}))
	playgroundHandler := playground.Handler("GraphQL playground", "/apis/query")

	apis.Get("/playground", adaptHTTPHandler(playgroundHandler))
	apis.Post("/query", adaptHTTPHandler(gqlHandler))

	return s.fib
}
