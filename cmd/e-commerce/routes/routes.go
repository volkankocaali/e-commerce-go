package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/streadway/amqp"
	"github.com/volkankocaali/e-commorce-go/cmd/e-commerce/middleware"
	"github.com/volkankocaali/e-commorce-go/pkg/api/handler"
	"github.com/volkankocaali/e-commorce-go/pkg/config"
	"github.com/volkankocaali/e-commorce-go/pkg/helper"
	"github.com/volkankocaali/e-commorce-go/pkg/parser"
	"github.com/volkankocaali/e-commorce-go/pkg/repository"
	"github.com/volkankocaali/e-commorce-go/pkg/usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

/*
rd *redis.Client,
el *elastic.Client,
*/
func SetupRoutes(
	app *fiber.App,
	cfg *config.Config,
	db *gorm.DB,
	amqpConn *amqp.Connection,
	mg *mongo.Client,
) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Accept,Authorization,Content-Type,X-CSRF-TOKEN",
		ExposeHeaders:    "Link",
		AllowCredentials: true,
		MaxAge:           300,
	}))

	h := helper.NewHelper(cfg)

	// Parser
	productParser := parser.NewProductParser()

	// Repository
	userRepository := repository.NewUserRepository(db)
	walletRepository := repository.NewWalletRepository(db)
	productRepository := repository.NewProductRepository(db)
	categoriesRepository := repository.NewCategoriesRepository(db)

	// Usecase
	userUseCase := usecase.NewUserUseCase(userRepository, walletRepository, cfg, h)
	productUseCase := usecase.NewProductUseCase(productRepository, categoriesRepository, cfg, h, productParser)

	// Handler
	userHandler := handler.NewUserHandler(userUseCase)
	productHandler := handler.NewProductHandler(productUseCase)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Post("/sign-up", userHandler.SignUp)
	v1.Post("/login", userHandler.Login)

	// Jwt Middleware for Client User
	v1.Use(middleware.UserAuthMiddleware)

	// Log Processor Middleware
	v1.Use(func(ctx *fiber.Ctx) error {
		return middleware.LogProcessorMiddleware(ctx, amqpConn)
	})

	// list products for user
	v1.Get("/products", productHandler.ListProduct)
	v1.Get("/product/:id", productHandler.GetProduct)

	// Category
	// -- get category
	// -- get category products

	// add to cart
	// wishlist-add

	// User

	// -- get user details
	// -- get user address
	// -- add user address
	// -- get user reference link

	// Orders

	// -- get orders
	// -- get order details
	// -- cancel order
	// -- return order

	//
	v1.Get("/test-middleware", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "test middleware",
		})
	})
}
