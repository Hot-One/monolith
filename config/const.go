package config

import "time"

const (
	DebugMode   = "debug"
	TestMode    = "test"
	ReleaseMode = "release"

	AccessTokenExpiresInTime  time.Duration = 1 * 60 * 24 * time.Minute
	RefreshTokenExpiresInTime time.Duration = 30 * 24 * 60 * time.Minute
)
