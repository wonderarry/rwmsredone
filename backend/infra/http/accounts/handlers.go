package accountsapi

import (
	"net/http"

	"github.com/wonderarry/rwmsredone/infra/http/httputils"
	"github.com/wonderarry/rwmsredone/internal/app/accounts"
	"github.com/wonderarry/rwmsredone/internal/domain"
)

type Handlers struct {
	Svc accounts.Service
}

func New(svc accounts.Service) *Handlers { return &Handlers{Svc: svc} }

// RegisterLocalReq is the request payload for local registration.
type RegisterLocalReq struct {
	Login          string `json:"login" example:"jdoe"`
	Password       string `json:"password" example:"secret123"`
	FirstName      string `json:"first_name" example:"John"`
	MiddleName     string `json:"middle_name" example:"Q"`
	LastName       string `json:"last_name" example:"Doe"`
	GroupNumber    string `json:"group_number" example:"CS-101"`
	GrantCanCreate bool   `json:"grant_can_create" example:"true"`
}

// RegisterLocal godoc
// @Summary      Register a local account
// @Description  Create a new account with username/password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      RegisterLocalReq  true  "registration payload"
// @Success      200      {object}  map[string]string  "account_id"
// @Failure      400      {object}  map[string]string
// @Router       /api/auth/register-local [post]
func (h *Handlers) RegisterLocal(w http.ResponseWriter, r *http.Request) {
	var req RegisterLocalReq
	if err := httputils.DecodeJSON(r, &req); err != nil {
		httputils.ErrorJSON(w, 400, err)
		return
	}
	roles := make([]domain.GlobalRole, 0, len(domain.AllGlobalRoles))
	if req.GrantCanCreate {
		roles = append(roles, domain.RoleCanCreateProjects)
	}

	id, err := h.Svc.RegisterLocal(r.Context(), accounts.RegisterLocal{
		Login:      req.Login,
		Password:   req.Password,
		FirstName:  req.FirstName,
		MiddleName: req.MiddleName,
		LastName:   req.LastName,
		Group:      req.GroupNumber,
		Roles:      roles,
	})
	if err != nil {
		httputils.ErrorJSON(w, 400, err)
		return
	}
	httputils.WriteJSON(w, 200, map[string]string{"account_id": string(id)})
}

// LoginLocalReq is the request payload for local login.
type LoginLocalReq struct {
	Login    string `json:"login" example:"jdoe"`
	Password string `json:"password" example:"secret123"`
}

// LoginLocal godoc
// @Summary      Login with local credentials
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      LoginLocalReq  true  "login payload"
// @Success      200      {object}  map[string]string  "token"
// @Failure      400      {object}  map[string]string
// @Failure      401      {object}  map[string]string
// @Router       /api/auth/login-local [post]
func (h *Handlers) LoginLocal(w http.ResponseWriter, r *http.Request) {
	// slog.Info("osdfoishfois")
	var req LoginLocalReq
	if err := httputils.DecodeJSON(r, &req); err != nil {
		httputils.ErrorJSON(w, 400, err)
		return
	}
	tok, err := h.Svc.LoginLocal(r.Context(), req.Login, req.Password)
	if err != nil {
		httputils.ErrorJSON(w, 401, err)
		return
	}
	httputils.WriteJSON(w, 200, map[string]string{"token": tok.AccessToken})
}

// GetMe godoc
// @Summary      Get current user profile
// @Tags         accounts
// @Security     BearerAuth
// @Produce      json
// @Success      200  {object}  domain.Account
// @Failure      400  {object}  map[string]string
// @Router       /api/accounts/me [get]
func (h *Handlers) GetMe(w http.ResponseWriter, r *http.Request) {
	actor := httputils.ActorIDFrom(r.Context())
	me, err := h.Svc.GetMe(r.Context(), actor)
	if err != nil {
		httputils.ErrorJSON(w, 400, err)
		return
	}
	httputils.WriteJSON(w, 200, me)
}
