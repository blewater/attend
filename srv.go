package main

import (
	"context"

	// "github.com/blewater/attend/gcpfunc"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/blewater/attend/gcpfunc"
	"github.com/joho/godotenv"
)

const (
	shutdownSecondsAllowance = 15
	defaultHTTPPort          = 8080
	viberKifisiaKey          = "VIBER_BOT_KEY"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	viberKey := os.Getenv(viberKifisiaKey)
	workflow(logger, defaultHTTPPort, viberKey)
	// ctx := context.Background()
	// // if err := funcframework.RegisterHTTPFunctionContext(ctx, "/", gcpfunc.HelloWorld); err != nil {
	// // 	log.Fatalf("funcframework.RegisterHTTPFunctionContext: %v\n", err)
	// // }
	// if err := funcframework.RegisterHTTPFunctionContext(ctx, "/", gcpfunc.Inquire); err != nil {
	// 	log.Fatalf("funcframework.RegisterHTTPFunctionContext: %v\n", err)
	// }
	// // Use PORT environment variable, or default to 8080.
	// port := "8080"
	// if envPort := os.Getenv("PORT"); envPort != "" {
	// 	port = envPort
	// }
	// if err := funcframework.Start(port); err != nil {
	// 	log.Fatalf("funcframework.Start: %v\n", err)
	// }
}

// Performs steps for launching the web server.
func workflow(logger *log.Logger, port int, botKey string) {
	httpServer := getHTTPServer(logger, port, botKey)

	setupTerminateSignal(logger, httpServer, port)

	launchHTTPListener(logger, httpServer, port)
}

// getHTTPServer constructs an HTTP listening server with 2 request handlers
func getHTTPServer(logger *log.Logger, port int, botKey string) *http.Server {
	viber := gcpfunc.Viber{Key: botKey}

	router := http.NewServeMux()

	// mute favicon requests
	router.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
	})

	// handler echoes the Path component of the request URL r.
	router.HandleFunc("/", viber.Inquire) // each request calls handler

	return &http.Server{
		Addr:     fmt.Sprintf(":%d", port),
		Handler:  router,
		ErrorLog: logger,
	}
}

// setupTerminateSignal connects the os.Interrupt signal to a quit channel to
// start teardown.
func setupTerminateSignal(logger *log.Logger, httpServer *http.Server, port int) {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)

	go httpServerShutdown(logger, httpServer, port, quit)
}

// Final step in launching an http server: Start accepting requests.
func launchHTTPListener(logger *log.Logger, httpServer *http.Server, port int) {
	logger.Printf("Http server at container port %d listening...\n", port)

	err := httpServer.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Can't launch http listener at container %d...\n", port)
	}
}

// httpServerShutdown handles the termination signal by shutting down the http server
// by closing connections and forcing shutdown if needed: "shutdownSecondsAllowance" max allowance.
func httpServerShutdown(logger *log.Logger, httpServer *http.Server, port int, quit <-chan os.Signal) {
	<-quit
	logger.Printf("Http server at container port %d is shutting down...\n", port)

	// Allow
	ctx, cancel := context.WithTimeout(context.Background(), shutdownSecondsAllowance*time.Second)
	defer cancel()

	err := httpServer.Shutdown(ctx)
	if err != nil {
		logger.Fatalf("Could not shutdown the server @ %d. Error: %v\n", port, err)
	}
}
