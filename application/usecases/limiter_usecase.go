package usecases

import "log"

type LimiterUseCaseInterface interface {
	ValidRateLimiter(string) string
}

type limiterUseCase struct{}

func NewLimiterUseCase() LimiterUseCaseInterface {
	return &limiterUseCase{}
}

func (u *limiterUseCase) ValidRateLimiter(parameter string) string {

	log.Printf("Parameter %s", parameter)

	return "Hello from UseCase"
}
