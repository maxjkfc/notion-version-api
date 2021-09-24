package model

import "time"

// Service -
type Service struct {
	Name          string // Name 服務名稱
	Tag           string
	Commit        string // Commit id
	CommitMessage string
	LastModifly   string // LastModifly 最後更新者
	Date          time.Time
}
