package main

import (
	"net/http"
	"time"

	"github.com/99designs/gqlgen/handler"
	graphql_server "github.com/alterra/graphql-server"
	"github.com/alterra/graphql-server/resolver"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.GraphQL(graphql_server.NewExecutableSchema(graphql_server.Config{Resolvers: &resolver.Resolver{}}),
		// handler.WebsocketKeepAliveDuration(19*time.Second),
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

	// r.Use(gin.BasicAuth(gin.Accounts{
	// 	"foo":    "bar",
	// 	"austin": "1234",
	// 	"lena":   "hello2",
	// 	"manu":   "4321",
	// }))

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "OPTION"},
		AllowHeaders: []string{"*"},
	}))

	r.POST("/query", graphqlHandler())
	r.GET("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run()
}
