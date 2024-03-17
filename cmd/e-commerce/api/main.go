package main

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/olivere/elastic/v7"
	slogfiber "github.com/samber/slog-fiber"
	"github.com/streadway/amqp"
	"github.com/volkankocaali/e-commorce-go/cmd/e-commerce/routes"
	"github.com/volkankocaali/e-commorce-go/pkg/config"
	"github.com/volkankocaali/e-commorce-go/pkg/database"
	"github.com/volkankocaali/e-commorce-go/pkg/seed"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"log"
	"log/slog"
	"os"
)

type App struct {
	FiberApp   *fiber.App
	DB         *gorm.DB
	Config     *config.Config
	Redis      *redis.Client
	Mongo      *mongo.Client
	Elastic    *elastic.Client
	AmqpClient *amqp.Connection
	LogFile    *os.File // Log dosyasını tutacak yeni alan
}

func main() {
	app := App{}

	app.initialize()
	app.FiberApp = app.initFiber(app.Config)

	// seed database
	seed.CreateAndSeed()

	// setup routes
	routes.SetupRoutes(app.FiberApp, app.Config, app.DB, app.AmqpClient, app.Mongo)

	// run fiber app
	app.run()

}

func (app *App) initialize() {
	cfg := config.LoadConfig()
	mysql, _ := database.NewMysqlDB(*cfg)
	redisClient, _ := database.NewRedisClient(*cfg)
	mongoClient, _ := database.NewMongoClient(*cfg)
	elasticClient, _ := database.NewElasticClient(*cfg)
	amqpClient, _ := database.NewAMQPClient(*cfg)

	app.Config = cfg
	app.DB = mysql
	app.Redis = redisClient
	app.Mongo = mongoClient
	app.Elastic = elasticClient
	app.AmqpClient = amqpClient
}

func (app *App) initFiber(cfg *config.Config) *fiber.App {
	fiberApp := fiber.New(fiber.Config{
		AppName:     cfg.AppName,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// Alternatif olarak, özelleştirilmiş bir format ile log dosyasına loglama yapabilirsiniz
	file, err := os.OpenFile("fiber_logs.json", os.O_RDWR|os.O_SYNC|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Log dosyası açılamadı: %v", err)
	}
	app.LogFile = file

	// slog için bir logger oluştur
	logger := slog.New(slog.NewJSONHandler(file, nil))

	// Fiber uygulamasına slog-fiber middleware'ini ekleyin
	fiberApp.Use(slogfiber.New(logger))
	fiberApp.Use(recover.New())

	return fiberApp
}

func (app *App) run() {
	defer app.LogFile.Close()
	log.Printf("Starting %s app, server on port %s", app.Config.AppName, app.Config.AppPort)
	if err := app.FiberApp.Listen(":" + app.Config.AppPort); err != nil {
		log.Fatalf("Fiber uygulaması başlatılırken hata: %v", err)
	}
}
