package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"arthemis-watcher/internal/models"

	"gorm.io/gorm"
)

type AuditParams struct {
	logger *slog.Logger
	db     *gorm.DB
}

func AuditHandler(logger *slog.Logger, db *gorm.DB) *AuditParams {
	return &AuditParams{logger, db}
}

// CreateAuditLog receives an audit log entry and saves it to the database
// @Summary      Create Audit Log
// @Description  Receives an audit log entry and saves it to the database
// @Tags         audit
// @Accept       json
// @Produce      json
// @Param        entry  body      models.AuditLog  true  "Audit Log entry payload"
// @Success      201    {object}  models.AuditLog        "Created audit log entry"
// @Failure      400    {object}  map[string]string      "Invalid request body"
// @Failure      500    {object}  map[string]string      "Failed to save to database"
// @Router       /audit/log [post]
func (a *AuditParams) CreateAuditLog(w http.ResponseWriter, r *http.Request) {
	var entry models.AuditLog
	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		a.logger.Error("watcher: failed to decode audit log body", "err", err)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	if err := a.db.Create(&entry).Error; err != nil {
		a.logger.Error("watcher: failed to save audit log to DB", "err", err)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "Failed to save to database"})
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(entry); err != nil {
		a.logger.Error("watcher: failed to encode response", "err", err)
	}
}
