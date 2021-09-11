package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	db "github.com/wildanpurnomo/go-persistent-chat/server/db"
	gqlSchema "github.com/wildanpurnomo/go-persistent-chat/server/gql/schema"
	authLibs "github.com/wildanpurnomo/go-persistent-chat/server/libs/auth"
	corsLibs "github.com/wildanpurnomo/go-persistent-chat/server/libs/cors"
	environmentLibs "github.com/wildanpurnomo/go-persistent-chat/server/libs/environment"
)

const (
	GracefulShutdownTimeout = 5 * time.Second
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	if err := db.MigrateUp(); err != nil {
		log.Fatal(err)
	}

	databaseClient, err := db.CreateNewDatabaseClient()
	if err != nil {
		log.Fatal(err)
	}

	db.InitiateAppRepository()
	db.AppRepository.InjectDatabaseClient(databaseClient)
	defer db.AppRepository.StopDatabaseClientConnection()

	if environmentLibs.IsRunningOnProductionEnvironment() {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	router.Use(corsLibs.CORSMiddleware())
	router.Use(authLibs.AuthMiddleware())

	gqlRequestHandler, err := gqlSchema.CreateGQLRequestHandler()
	if err != nil {
		log.Fatal(err)
	}

	router.GET("/gql", gqlRequestHandler)
	router.POST("/gql", gqlRequestHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	signalNotifContext, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	log.Println("Server is starting: http://localhost:8080/gql")
	<-signalNotifContext.Done()

	log.Println("Shutting down gracefully")

	timeoutContext, cancel := context.WithTimeout(context.Background(), GracefulShutdownTimeout)
	defer cancel()
	if err := server.Shutdown(timeoutContext); err != nil {
		log.Fatal("Server is forcefully shut down")
	}

	log.Println("Server is exiting")
}
