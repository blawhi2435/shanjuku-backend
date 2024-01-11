package service

type Service struct {
	GinService      *GinService
	PostgresService *PostgresService
	LoggerService   *LoggerService
}

func ProvideService(g *GinService, p *PostgresService, l *LoggerService) *Service {
	return &Service{GinService: g,
		PostgresService: p,
		LoggerService:   l,
	}
}
