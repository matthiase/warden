package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/birdbox/authnz"
	"github.com/birdbox/authnz/config"
	"github.com/birdbox/authnz/routes"
)

func main() {
	cfg := config.ReadEnv()

	app, err := authnz.NewApplication(cfg)
	if err != nil {
		log.Fatalf("Could not set up application: %v", err)
	}

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
	}

	//dbUser, dbPassword, dbName :=
	//    os.Getenv("POSTGRES_USER"),
	//    os.Getenv("POSTGRES_PASSWORD"),
	//    os.Getenv("POSTGRES_DB")
	//database, err := db.Initialize(dbUser, dbPassword, dbName)
	//if err != nil {
	//    log.Fatalf("Could not set up database: %v", err)
	//}
	//defer database.Conn.Close()

	httpHandler := routes.NewHandler(app)
	server := &http.Server{
		Handler: httpHandler,
	}
	go func() {
		server.Serve(listener)
	}()
	defer Stop(server)
	log.Printf("Started server on %s", addr)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(fmt.Sprint(<-ch))
	log.Println("Stopping API server")
}

func Stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Could not shut down server correctly: %v\n", err)
		os.Exit(1)
	}
}
