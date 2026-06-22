package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/task-tracker/internal/db/postgres"
	"github.com/task-tracker/internal/handler"
	"github.com/task-tracker/internal/repository"
	"github.com/task-tracker/internal/service"
)

func main() {
	var cfg postgres.Config
	if err := envconfig.Process("", &cfg); err != nil {
		fmt.Printf("failed to process env vars: %v", err)
	}

	pool, err := postgres.NewPool(&cfg)
	if err != nil {
		fmt.Printf("failed to connect to database: %v", err)
		os.Exit(1)
	}
	defer pool.Close()

	taskListRepo := repository.NewTaskListRepository(pool)
	taskListService := service.NewTaskListService(taskListRepo)
	taskListHandler := handler.NewTaskListHandler(taskListService)
	mux := http.NewServeMux()

	mux.HandleFunc("GET /debug/info", handler.HandleDebugInfo)
	mux.HandleFunc("POST /v1/lists", taskListHandler.CreateList)
	mux.HandleFunc("GET /v1/lists/{list_id}", taskListHandler.GetList)

	fmt.Println("server started at localhost:8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		_ = fmt.Errorf("failed to start server: %v", err)
	}
}
