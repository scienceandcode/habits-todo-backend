package service

type HealthService struct{}

func (*HealthService) GetHealthMessage() string {
	return "OK"
}

func NewHealthService() *HealthService {
	return &HealthService{}
}
