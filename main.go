package main

import (
	"context"
	"dpai/handler"
	"dpai/repository"
	"dpai/router"
	"dpai/usecase"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	tsRepo := repository.NewTwoSumRepository()
	tsUseCase := usecase.NewTwoSumUseCase(tsRepo)
	tsHandler := handler.NewTwoSumHandler(tsUseCase)

	threeSumUseCase := usecase.NewThreeSumUseCase(tsRepo)
	threeSumHandler := handler.NewThreeSumHandler(threeSumUseCase)

	routerOpt := router.Options{
		TwoSumHandler:   tsHandler,
		ThreeSumHandler: threeSumHandler,
	}

	r := router.NewRouter(routerOpt)

	srv := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")

}
