package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"transactions-routine/handlers"
	"transactions-routine/repository"
	"transactions-routine/services"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize repositories
	accountRepo := repository.NewInMemoryAccountRepository()
	transactionRepo := repository.NewInMemoryTransactionRepository()

	// Initialize services
	accountService := services.NewAccountService(accountRepo)
	transactionService := services.NewTransactionService(accountRepo, transactionRepo)

	// Initialize handlers
	accountHandler := handlers.NewAccountHandler(accountService)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	// routes
	router := mux.NewRouter()
	router.HandleFunc("/accounts", accountHandler.CreateAccount).Methods("POST")
	router.HandleFunc("/accounts/{accountId}", accountHandler.GetAccount).Methods("GET")
	router.HandleFunc("/transactions", transactionHandler.CreateTransaction).Methods("POST")

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Channel to listen for OS interrupts
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Run the server
	go func() {
		log.Println("Server running on port 8080")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	// Wait for shutdown signal
	<-stop
	log.Println("Shutting down server...")

	// context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// gracefully shutdown the server
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
