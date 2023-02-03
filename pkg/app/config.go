package app

import "time"

type Config struct {
	TrackedAccounts []string
	RefreshRate     time.Duration
}
