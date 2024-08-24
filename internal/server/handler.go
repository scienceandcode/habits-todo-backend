package server

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/scienceandcode/habits-todo-backend/pkg/common"
)

var ginLambda *ginadapter.GinLambda

func Handle(key string, app *gin.Engine) {
	handlers := map[string]func(*gin.Engine){
		"default": func(app *gin.Engine) {
			err := app.Run(":" + common.GetEnv("HTTP_SERVER_PORT"))
			if err != nil {
				log.Fatal(err.Error())
			}
		},
		"lambda": func(app *gin.Engine) {
			ginLambda = ginadapter.New(app)
			lambda.Start(func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
				return ginLambda.ProxyWithContext(ctx, req)
			})
		},
	}

	handler, exists := handlers[key]

	if !exists {
		log.Fatalf("(%s) HTTP server handler key does not exist", key)
	}

	handler(app)
}
