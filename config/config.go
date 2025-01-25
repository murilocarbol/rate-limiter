package config

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"rate-limiter/application/controllers"
	"rate-limiter/application/middleware"
	"time"
)

type Configure struct {
	TOKEN string `mapstructure:"TOKEN"`
}

func Initialize() {

	// Configuração das variáveis de ambiente
	_, err := LoadConfig(".")
	if err != nil {
		panic(err)
	}

	app := fiber.New()

	// Configuração do middleware
	rateLimiterConfig := middleware.RateLimiterConfig{
		Limit:        5,                // Limite de requisições
		Window:       10 * time.Second, // Janela de tempo para resetar
		ErrorMessage: "you have reached the maximum number of requests or actions allowed within a certain time frame.",
	}

	// Aplica o middleware com parâmetros personalizados
	app.Use(middleware.RateLimiterMiddleware(rateLimiterConfig))

	setRoutes(app)

	// Inicializa o servidor
	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}

func LoadConfig(path string) (*Configure, error) {
	var cfg *Configure
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath("./")
	viper.SetConfigFile("config.env")
	viper.AutomaticEnv()

	fmt.Println("Loading config from path:", path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, err
}

func setRoutes(app *fiber.App) {

	// Controllers
	rateLimiterController := controllers.NewRateLimiterController()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "API is running",
		})
	})

	app.Get("/", rateLimiterController.GetController)
}
