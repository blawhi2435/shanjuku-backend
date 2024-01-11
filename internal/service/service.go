package service

type Service struct {
	GinService      *GinService
	PostgresService *PostgresService
}

func ProvideService(g *GinService, p *PostgresService) *Service {
	return &Service{GinService: g, PostgresService: p}
}
