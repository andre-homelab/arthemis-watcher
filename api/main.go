// @title           Arthemis Watcher API
// @version         1.0
// @description     API service for Arthemis Watcher.
// @termsOfService  http://swagger.io/terms/

// @license.name  MIT
// @license.url   http://opensource.org/licenses/MIT

// @host      localhost:8082
// @BasePath  /
package main

import (
	"log/slog"
	"net/http"
	"os"

	"arthemis-watcher/internal/database"
	"arthemis-watcher/internal/env"
	"arthemis-watcher/internal/handlers"

	_ "arthemis-watcher/docs"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	r := chi.NewRouter()
	r.Use(chiMiddleware.Logger)

	textHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	logger := slog.New(textHandler)

	db, err := database.ConnectToDatabase(logger)
	if err != nil {
		logger.Error("Error initializing database!")
	}

	healthHandler := handlers.HealthHandler(logger, db)
	r.Get("/health", healthHandler.HealthCheck)

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("doc.json"),
	))

	auditHandler := handlers.AuditHandler(logger, db)
	r.Post("/audit/log", auditHandler.CreateAuditLog)

	port := env.GetEnv("PORT", "8082")
	logger.Info("Server started!")
	logger.Info("http://localhost:" + port)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		logger.Error("HTTP routing error", "error", err)
	}
}
