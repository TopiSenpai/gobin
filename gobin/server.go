package gobin

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/go-chi/httprate"
	"github.com/go-jose/go-jose/v3"
)

type ExecuteTemplateFunc func(wr io.Writer, name string, data any) error

func NewServer(version string, cfg Config, db *DB, signer jose.Signer, assets http.FileSystem, tmpl ExecuteTemplateFunc) *Server {
	s := &Server{
		version: version,
		cfg:     cfg,
		db:      db,
		signer:  signer,
		assets:  assets,
		tmpl:    tmpl,
	}

	if cfg.RateLimit != nil && cfg.RateLimit.Requests > 0 && cfg.RateLimit.Duration > 0 {
		s.rateLimitHandler = httprate.NewRateLimiter(
			cfg.RateLimit.Requests,
			cfg.RateLimit.Duration,
			httprate.WithLimitHandler(s.rateLimit),
			httprate.WithKeyFuncs(
				httprate.KeyByIP,
				httprate.KeyByEndpoint,
			),
		).Handler
	}

	return s
}

type Server struct {
	version          string
	cfg              Config
	db               *DB
	signer           jose.Signer
	assets           http.FileSystem
	tmpl             ExecuteTemplateFunc
	rateLimitHandler func(http.Handler) http.Handler
}

func (s *Server) Start() {
	if err := http.ListenAndServe(s.cfg.ListenAddr, s.Routes()); err != nil {
		log.Fatalln("Error while listening:", err)
	}
}

func FormatBuildVersion(version string, commit string, buildTime string) string {
	if len(commit) > 7 {
		commit = commit[:7]
	}

	buildTimeStr := "unknown"
	if buildTime != "unknown" {
		parsedTime, _ := time.Parse(time.RFC3339, buildTime)
		if !parsedTime.IsZero() {
			buildTimeStr = parsedTime.Format(time.ANSIC)
		}
	}
	return fmt.Sprintf("Go Version: %s\nVersion: %s\nCommit: %s\nBuild Time: %s\nOS/Arch: %s/%s\n", runtime.Version(), version, commit, buildTimeStr, runtime.GOOS, runtime.GOARCH)
}
