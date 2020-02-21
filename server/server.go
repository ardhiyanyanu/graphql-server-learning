package main

import (
	"net/http"
	"time"

	usermanagement "github.com/alterra/graphql-server/usermanagement"

	"github.com/99designs/gqlgen/handler"
	graphql_server "github.com/alterra/graphql-server"
	"github.com/alterra/graphql-server/graphql/resolver"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.GraphQL(graphql_server.NewExecutableSchema(graphql_server.Config{Resolvers: &resolver.Resolver{}}),
		handler.WebsocketKeepAliveDuration(20*time.Second),
		handler.WebsocketUpgrader(websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// simulate some private data
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
	// Setting up Gin
	r := gin.Default()

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "OPTION"},
		AllowHeaders: []string{"*"},
	}))

	// graphQL := r.Group("/query")
	// graphQL.Use(middlewares.IsLoggedIn())
	// graphQL.Use(middlewares.GinContextToContextMiddleware())
	r.POST("/query", graphqlHandler())
	r.GET("/query", graphqlHandler())

	r.GET("/", playgroundHandler())

	user := r.Group("/user")
	usermanagement.GetRoute(user)

	r.Run()
}
