package main

import (
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_ "github.com/go-sql-driver/mysql"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/controller"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/controller/auth"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/controller/handler"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/infrastructure/repository"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/service"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/usecase"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/pkg/genid"
)

var (
	secretKey   = []byte(os.Getenv("SECRET_KEY"))
	port        = "8080"
	redisURI    = "redis://user:@localhost:6379/0"
	mongodbURI  = "mongodb://root:password@localhost:27017"
	mongodbName = "redditclone"
	mysqlURI    = "root:password@tcp(localhost:3306)/redditclone?"
	idLen       = 20
)

func main() {
	ctx := context.Background()

	connect, errMongoConnect := mongo.Connect(ctx, options.Client().ApplyURI(mongodbURI))
	if errMongoConnect != nil {
		log.Fatalf("mongo.Connect(ctx, options.Client().ApplyURI(mongodbURI)) failed: %v", errMongoConnect)
	}
	collection := connect.Database(mongodbName).Collection("posts")

	postRepositoryMongoDB := repository.NewPostRepositoryMongoDB(connect, collection)
	log.Println("||INFO|| postRepositoryMongoDB was created")

	mysqlURI += "&charset=utf8"
	mysqlURI += "&interpolateParams=true"
	dbMySQL, errMySQLOpen := sql.Open("mysql", mysqlURI)
	if errMySQLOpen != nil {
		log.Fatalf(" sql.Open(\"mysql\", mysqlURI)failed: %v", errMySQLOpen)
	}
	dbMySQL.SetMaxOpenConns(10)

	errPing := dbMySQL.PingContext(ctx)
	if errPing != nil {
		log.Fatalf("dbMySql.Ping() failed: %v", errPing)
	}

	userRepositoryMySQL := repository.NewUserRepositoryMySQL(dbMySQL)
	log.Println("||INFO|| userRepositoryMySQL was created")

	service := service.NewService(postRepositoryMongoDB, userRepositoryMySQL)
	log.Println("||INFO|| service was created")

	genID := genid.NewIDGeneratorFacade(ctx, genid.NewIDGenerator(ctx, uint64(idLen)), genid.NewIDTracker(ctx))
	useCase := usecase.NewUseCase(service, genID)
	log.Println("||INFO|| useCase was created")

	template, err := template.New("name").ParseFiles("../../static/html/index.html")
	if err != nil {
		log.Fatalf("template.New(\"name\").ParseFiles(\"index.html\") failed: %v", err)
	}
	log.Println("||INFO|| template was created")

	c, errDialURL := redis.DialURL(redisURI)
	if errDialURL != nil {
		log.Fatalf("redis.DialURL(redisURL) failed: %v", err)
	}
	defer c.Close()
	log.Println("||INFO|| sessionRepositoryRedis was created")

	authManager := auth.NewAuthManager(secretKey, func(token *jwt.Token) (interface{}, error) {
		method, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok || method.Alg() != "HS256" {
			return nil, fmt.Errorf("bad sign method")
		}
		return secretKey, nil
	}, c)
	log.Println("||INFO|| authManager was created")
	handler := handler.NewHandler(useCase, template, authManager)
	log.Println("||INFO|| handler was created")

	router := mux.NewRouter()
	log.Println("||INFO|| router was created")

	contr := controller.NewController(handler, router)
	contr.RegisterRoutes()
	mux := contr.UseMiddleware()
	log.Println("||INFO|| controller was created")

	log.Printf("||INFO|| start listen and serve port :%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("||FATAL|| failed start server on port :%s", port)
	}

}
