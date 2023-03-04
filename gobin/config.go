package gobin

import (
	"fmt"
	"strings"
	"time"
)

type Config struct {
	DevMode         bool             `cfg:"dev_mode"`
	Debug           bool             `cfg:"debug"`
	ListenAddr      string           `cfg:"listen_addr"`
	Database        DatabaseConfig   `cfg:"database"`
	MaxDocumentSize int              `cfg:"max_document_size"`
	RateLimit       *RateLimitConfig `cfg:"rate_limit"`
	JWTSecret       string           `cfg:"jwt_secret"`
}

func (c Config) String() string {
	return fmt.Sprintf("\n DevMode: %t\n Debug: %t\n ListenAddr: %s\n Database: %s\n MaxDocumentSize: %d\n Rate Limit: %s\n JWTSecret: %s\n", c.DevMode, c.Debug, c.ListenAddr, c.Database, c.MaxDocumentSize, c.RateLimit, strings.Repeat("*", len(c.JWTSecret)))
}

type DatabaseConfig struct {
	Type            string        `cfg:"type"`
	Debug           bool          `cfg:"debug"`
	ExpireAfter     time.Duration `cfg:"expire_after"`
	CleanupInterval time.Duration `cfg:"cleanup_interval"`

	// SQLite
	Path string `cfg:"path"`

	// PostgreSQL
	Host     string `cfg:"host"`
	Port     int    `cfg:"port"`
	Username string `cfg:"username"`
	Password string `cfg:"password"`
	Database string `cfg:"database"`
	SSLMode  string `cfg:"ssl_mode"`
}

func (c DatabaseConfig) String() string {
	str := fmt.Sprintf("\n  Type: %s\n  Debug: %t\n  ExpireAfter: %s\n  CleanupInterval: %s\n  ", c.Type, c.Debug, c.ExpireAfter, c.CleanupInterval)
	switch c.Type {
	case "postgres":
		str += fmt.Sprintf("Host: %s\n  Port: %d\n  Username: %s\n  Password: %s\n  Database: %s\n  SSLMode: %s", c.Host, c.Port, c.Username, strings.Repeat("*", len(c.Password)), c.Database, c.SSLMode)
	case "sqlite":
		str += fmt.Sprintf("Path: %s", c.Path)
	default:
		str += "Invalid database type!"
	}
	return str
}

func (c DatabaseConfig) PostgresDataSourceName() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", c.Host, c.Port, c.Username, c.Password, c.Database, c.SSLMode)
}

type RateLimitConfig struct {
	Requests int           `cfg:"requests"`
	Duration time.Duration `cfg:"duration"`
}

func (c RateLimitConfig) String() string {
	return fmt.Sprintf("\n  Requests: %d\n  Duration: %s", c.Requests, c.Duration)
}
