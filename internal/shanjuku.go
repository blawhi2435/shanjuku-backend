package internal

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/blawhi2435/shanjuku-backend/graph"
	"github.com/blawhi2435/shanjuku-backend/graph/directive"
	"github.com/blawhi2435/shanjuku-backend/internal/service"
	"github.com/gin-gonic/gin"
)


func Init() {
	svc, err := service.InitService()
	if err != nil {
		panic(err)
	}

	r := svc.GinService.Engine
	r.POST("/graphql", graphqlHandler(svc))
	r.GET("/", playgroundHandler())
	r.Run()
}

func graphqlHandler(svc *service.Service) gin.HandlerFunc {

	myResolver := InitResolver(svc.PostgresService.DB, svc.LoggerService)
	
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	graphConfig := graph.Config{Resolvers: myResolver}
	graphConfig.Directives.Binding = directive.Binding
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graphConfig))
	h.Use(extension.FixedComplexityLimit(20))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}