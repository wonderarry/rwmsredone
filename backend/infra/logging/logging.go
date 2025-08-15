package logging

import (
	"io"
	"log"
	"log/slog"
	"os"
	"strings"
	"time"
)

type Options struct {
	Level     string // "debug"|"info"|"warn"|"error"
	JSON      bool   // true => JSON, false => text
	AddSource bool   // include file:line (handy in dev)
	Out       io.Writer
	Service   string // e.g. "rwms"
	Env       string // e.g. os.Getenv("RWMS_ENV")
}

func Init(opts Options) *slog.Logger {
	if opts.Out == nil {
		opts.Out = os.Stdout
	}
	lvl := parseLevel(opts.Level)

	hopts := &slog.HandlerOptions{
		Level:     lvl,
		AddSource: opts.AddSource,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.TimeKey:

				a.Key = "ts"
				a.Value = slog.StringValue(a.Value.Time().UTC().Format(time.RFC3339Nano))
			case slog.LevelKey:
				if lv, ok := a.Value.Any().(slog.Level); ok {
					a.Key = "lvl"
					a.Value = slog.StringValue(strings.ToLower(lv.String()))
				}
			}
			return a
		},
	}

	var h slog.Handler
	if opts.JSON {
		h = slog.NewJSONHandler(opts.Out, hopts)
	} else {
		h = slog.NewTextHandler(opts.Out, hopts)
	}

	base := slog.New(h).
		With("service", coalesce(opts.Service, "rwms")).
		With("env", coalesce(opts.Env, getenvOr("RWMS_ENV", "local")))

	slog.SetDefault(base)

	ll := slog.NewLogLogger(h, slog.LevelInfo)
	log.SetFlags(0)
	log.SetOutput(ll.Writer())

	return base
}

func parseLevel(s string) slog.Level {
	switch strings.ToLower(s) {
	case "debug":
		return slog.LevelDebug
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func getenvOr(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
func coalesce(v, def string) string {
	if strings.TrimSpace(v) == "" {
		return def
	}
	return v
}
