package models

import (
	"time"
)

type AuditLog struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Timestamp  time.Time `json:"timestamp"`
	Method     string    `json:"method"`
	Path       string    `json:"path"`
	StatusCode int       `json:"status_code"`
	Duration   int64     `json:"duration_ms"`
	UserID     string    `json:"user_id,omitempty"`
	ReqBody    string    `json:"req_body,omitempty"`
}
