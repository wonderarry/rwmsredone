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

type registerLocalReq struct {
	Login          string `json:"login"`
	Password       string `json:"password"`
	FirstName      string `json:"first_name"`
	MiddleName     string `json:"middle_name"`
	LastName       string `json:"last_name"`
	GroupNumber    string `json:"group_number"`
	GrantCanCreate bool   `json:"grant_can_create"`
}

func (h *Handlers) RegisterLocal(w http.ResponseWriter, r *http.Request) {
	var req registerLocalReq
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

type loginLocalReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h *Handlers) LoginLocal(w http.ResponseWriter, r *http.Request) {
	var req loginLocalReq
	if err := httputils.DecodeJSON(r, &req); err != nil {
		httputils.ErrorJSON(w, 400, err)
		return
	}
	tok, err := h.Svc.LoginLocal(r.Context(), req.Login, req.Password)
	if err != nil {
		httputils.ErrorJSON(w, 401, err)
		return
	}
	httputils.WriteJSON(w, 200, map[string]string{"token": tok})
}

func (h *Handlers) GetMe(w http.ResponseWriter, r *http.Request) {
	actor := httputils.ActorIDFrom(r.Context())
	me, err := h.Svc.GetMe(r.Context(), actor)
	if err != nil {
		httputils.ErrorJSON(w, 400, err)
		return
	}
	httputils.WriteJSON(w, 200, me)
}
