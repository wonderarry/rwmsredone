package httpapi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	middleware2 "github.com/go-chi/chi/v5/middleware"

	httpswagger "github.com/swaggo/http-swagger"
	accapi "github.com/wonderarry/rwmsredone/infra/http/accounts"
	"github.com/wonderarry/rwmsredone/infra/http/middleware"
	procapi "github.com/wonderarry/rwmsredone/infra/http/processes"
	projapi "github.com/wonderarry/rwmsredone/infra/http/projects"
	"github.com/wonderarry/rwmsredone/internal/app"
	"github.com/wonderarry/rwmsredone/internal/app/contract"
)

type Server struct {
	Accounts  *accapi.Handlers
	Projects  *projapi.Handlers
	Processes *procapi.Handlers
	Tokens    contract.TokenIssuer
}

func New(svcs *app.Services, tokens contract.TokenIssuer) *Server {
	return &Server{
		Accounts:  accapi.New(svcs.Accounts),
		Projects:  projapi.New(svcs.Projects),
		Processes: procapi.New(svcs.Processes),
		Tokens:    tokens,
	}
}

func (s *Server) Routes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware2.RequestID, middleware2.RealIP, middleware2.Recoverer)

	// public
	r.Route("/swagger", func(r chi.Router) {
		r.Get("/*", httpswagger.WrapHandler)
	})

	r.Route("/api/auth", func(r chi.Router) {
		r.Post("/register-local", s.Accounts.RegisterLocal)
		r.Post("/login-local", s.Accounts.LoginLocal)
	})

	// protected
	r.Group(func(r chi.Router) {
		r.Use(middleware.RequireAuth(s.Tokens))

		r.Route("/api/accounts", func(r chi.Router) {
			r.Get("/me", s.Accounts.GetMe)
		})

		r.Route("/api/projects", func(r chi.Router) {
			r.Post("/", s.Projects.Create)
			r.Get("/", s.Projects.ListMine)
			r.Put("/{id}", s.Projects.EditMeta)
			r.Post("/{id}/members", s.Projects.AddMember)
			r.Delete("/{id}/members", s.Projects.RemoveMember)
		})

		r.Route("/api/processes", func(r chi.Router) {
			r.Post("/", s.Processes.Create)
			r.Post("/{pid}/members", s.Processes.AddMember)
			r.Delete("/{pid}/members", s.Processes.RemoveMember)
			r.Post("/{pid}/approvals", s.Processes.RecordApproval)
			// r.Get("/{pid}", s.Processes.Get)
			// r.Get("/{pid}/graph", s.Processes.GetGraph)
			// r.Get("/{pid}/approvals/{stage}", s.Processes.ListApprovals)
		})
	})

	return r
}
