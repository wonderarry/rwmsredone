package main

/* swagger setup */
// @title           RWMS API
// @version         0.1
// @description     Research Workflow Management System
// @BasePath        /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description  Use "Bearer <JWT>"

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/wonderarry/rwmsredone/docs"

	"github.com/wonderarry/rwmsredone/infra/config"
	dbuow "github.com/wonderarry/rwmsredone/infra/db"
	"github.com/wonderarry/rwmsredone/infra/db/templates"
	httpapi "github.com/wonderarry/rwmsredone/infra/http"
	"github.com/wonderarry/rwmsredone/infra/logging"
	"github.com/wonderarry/rwmsredone/infra/security"
	"github.com/wonderarry/rwmsredone/internal/app"
)

type idGen struct{}

func (idGen) NewID() string { return security.NewULID() }

func main() {
	ctx := context.Background()
	cfg, err := config.Load()
	must(err)

	pool, err := pgxpool.New(ctx, cfg.DB.DSN)
	must(err)
	defer pool.Close()

	uow := dbuow.NewUoW(pool)
	hasher := security.NewBcrypt(0)
	tokens := security.NewJWTIssuer(cfg.Security.JWTSecret, "rwms")
	tpls := templates.NewFSProvider(cfg.TemplatesDir)
	ids := idGen{}

	svcs, err := app.NewServices(app.Deps{
		UoW:            uow,
		PasswordHasher: hasher,
		TokenIssuer:    tokens,
		Templates:      tpls,
		IDGen:          ids,
	})
	must(err)

	server := &http.Server{
		Addr:         cfg.HTTP.Addr,
		Handler:      httpapi.New(&svcs, tokens).Routes(),
		ReadTimeout:  cfg.HTTP.ReadTimeout,
		WriteTimeout: cfg.HTTP.WriteTimeout,
		IdleTimeout:  cfg.HTTP.IdleTimeout,
	}

	logging.Init(logging.Options{
		Level:     cfg.LogLevel,                // RWMS_LOG_LEVEL: debug/info/warn/error
		JSON:      os.Getenv("RWMS_ENV") != "", // JSON in docker/k8s, text locally
		AddSource: os.Getenv("RWMS_ENV") == "", // file:line in dev
		Service:   "rwms",
		Env:       os.Getenv("RWMS_ENV"),
	})

	slog.Info("starting", "addr", cfg.HTTP.Addr)
	log.Fatal(server.ListenAndServe())
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
