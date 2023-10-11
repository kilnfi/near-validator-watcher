package watcher

import (
	"io"
	"time"
)

type Config struct {
	TrackedAccounts []string
	RefreshRate     time.Duration
	Writer          io.Writer
}
