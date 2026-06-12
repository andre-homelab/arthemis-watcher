package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type GlobalParams struct {
	logger *slog.Logger
	db     *gorm.DB
}

func HealthHandler(logger *slog.Logger, db *gorm.DB) *GlobalParams {
	return &GlobalParams{logger, db}
}

// HealthCheck evaluates the status of the server and database
// @Summary      Health Check
// @Description  Evaluates the status of the server and database
// @Tags         health
// @Produce      json
// @Success      200  {object}  map[string]string "Successful response containing status, db connection, and timestamp"
// @Router       /health [get]
func (g *GlobalParams) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	dbIsUp := "available"

	if g.db == nil {
		dbIsUp = "unavailable"
	}

	if err := json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
		"db":     dbIsUp,
		"time":   time.Now().UTC().Format(time.RFC3339),
	}); err != nil {
		g.logger.Error("Error in the health response: ", "error", err)
	}
}
