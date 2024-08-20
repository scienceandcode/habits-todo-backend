package service

type HealthService struct{}

func (*HealthService) GetHealthMessage() string {
	return "Welcome to HabitsTodo API!"
}

func NewHealthService() *HealthService {
	return &HealthService{}
}
