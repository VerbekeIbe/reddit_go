package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/verbekeibe/reddit-backend/graph"
	"github.com/verbekeibe/reddit-backend/graph/generated"
	"github.com/verbekeibe/reddit-backend/graph/model"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"

	"github.com/joho/godotenv"
)

func getEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
  
	return os.Getenv(key)
  }

  var database *gorm.DB;

func initDB() {
	var err error
	connectionString := "root:" + getEnvVariable("DB_PASSWORD") +"@tcp(" + getEnvVariable("DB_PORT")+ ")/reddit_go?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(connectionString)
	database, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		fmt.Println((err))
		panic("Failed to connect to database")
	}

	

	database.AutoMigrate( &model.User{},&model.Community{}, &model.Post{}, &model.Comment{}, &model.UserCommunity{})

}

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	initDB()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{ DB: database,}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
