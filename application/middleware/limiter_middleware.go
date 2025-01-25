package middleware

import (
	"log"
	"rate-limiter/application/usecases"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

// RateLimiterConfig Configuração personalizada para o rate limiter
type RateLimiterConfig struct {
	Limit        int
	Window       time.Duration
	ErrorMessage string
}

func RateLimiterMiddleware(config RateLimiterConfig) fiber.Handler {
	start := time.Now()

	// Mapa para armazenar contadores de requisições
	requestCounts := make(map[string]int)

	return func(c *fiber.Ctx) error {

		// Identifica o IP do cliente
		clientIP := c.IP()
		parameter := clientIP

		headers := c.GetReqHeaders()

		token := headers["Token"][0]
		if !strings.EqualFold(token, "") {
			log.Printf("Token informado: %s", token)
			parameter = token
		}

		limiterUseCase := usecases.NewLimiterUseCase()
		resp := limiterUseCase.ValidRateLimiter(parameter)
		log.Printf("Resposta usecase %s", resp)

		// Incrementa o contador de requisições
		requestCounts[clientIP]++

		// Verifica se o limite foi excedido
		if requestCounts[clientIP] > config.Limit {
			log.Printf("IP %s excedeu o limite de %d requisições.", clientIP, config.Limit)
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"message": config.ErrorMessage,
			})
		}

		// Registra os detalhes da requisição
		log.Printf("[%s] %s | Status: %d | Request Time: %s",
			c.Method(),
			c.Path(),
			c.Response().StatusCode(),
			time.Since(start),
		)

		// Reseta o contador após o período da janela
		go func(ip string) {
			time.Sleep(config.Window)
			delete(requestCounts, ip)
		}(clientIP)

		// Continua o fluxo padrão da chamada (controller)
		return c.Next()
	}
}
