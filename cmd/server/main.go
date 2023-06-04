package main

import (
	"context"
	"github.com/MicBun/protobuf-golang-todo/internal"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	serveCtx, cancelServeCtx := context.WithCancel(context.Background())
	log.Println("Initialize server...")

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Server failed to load env file: %v", err)
	}

	app, initAppErr := internal.InitApp(serveCtx)
	if initAppErr != nil {
		log.Fatal(initAppErr)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(
		sig,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	go func() {
		<-sig

		var cleanTimeoutDuration time.Duration = 30
		cleanCtx, cancelCleanCtx := context.WithTimeout(serveCtx, cleanTimeoutDuration*time.Second)

		go func() {
			<-cleanCtx.Done()

			if cleanCtx.Err() == context.DeadlineExceeded {
				log.Fatal("Graceful shutdown timed out... Forcing exit now...")
			}
		}()

		log.Println("Gracefully shutdown server...")
		cleanAppErr := app.Clean()
		if cleanAppErr != nil {
			log.Fatal(cleanAppErr)
		}

		cancelCleanCtx()
		cancelServeCtx()
	}()

	log.Println("Start serving...")
	err := app.Serve()
	if err != nil {
		log.Fatal(err)
	}

	<-serveCtx.Done()
}
