package projectsapi

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/wonderarry/rwmsredone/infra/http/httputils"
	"github.com/wonderarry/rwmsredone/internal/app/projects"
	"github.com/wonderarry/rwmsredone/internal/domain"
)

type Handlers struct{ Svc projects.Service }

func New(svc projects.Service) *Handlers { return &Handlers{Svc: svc} }

type createProjectReq struct {
	Name        string `json:"name"`
	Theme       string `json:"theme"`
	Description string `json:"description"`
}

func (h *Handlers) Create(w http.ResponseWriter, r *http.Request) {
	var req createProjectReq
	if err := httputils.DecodeJSON(r, &req); err != nil {
		httputils.ErrorJSON(w, 400, err)
		return
	}
	id, err := h.Svc.CreateProject(r.Context(), projects.CreateProject{
		ActorID:     httputils.ActorIDFrom(r.Context()),
		Name:        req.Name,
		Theme:       req.Theme,
		Description: req.Description,
	})
	if err != nil {
		httputils.ErrorJSON(w, 403, err)
		return
	}
	httputils.WriteJSON(w, 200, map[string]string{"project_id": string(id)})
}

func (h *Handlers) ListMine(w http.ResponseWriter, r *http.Request) {
	items, err := h.Svc.ListMyProjects(r.Context(), httputils.ActorIDFrom(r.Context()))
	if err != nil {
		httputils.ErrorJSON(w, 400, err)
		return
	}
	httputils.WriteJSON(w, 200, items)
}

type editProjectReq struct{ Name, Theme, Description string }

func (h *Handlers) EditMeta(w http.ResponseWriter, r *http.Request) {
	id := domain.ProjectID(chi.URLParam(r, "id"))
	var req editProjectReq
	if err := httputils.DecodeJSON(r, &req); err != nil {
		httputils.ErrorJSON(w, 400, err)
		return
	}
	err := h.Svc.EditProjectMeta(r.Context(), projects.EditProjectMeta{
		ActorID:     httputils.ActorIDFrom(r.Context()),
		ProjectID:   id,
		Name:        req.Name,
		Theme:       req.Theme,
		Description: req.Description,
	})
	if err != nil {
		httputils.ErrorJSON(w, 403, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

type memberReq struct {
	AccountID domain.AccountID   `json:"account_id"`
	Role      domain.ProjectRole `json:"role"`
}

func (h *Handlers) AddMember(w http.ResponseWriter, r *http.Request) {
	id := domain.ProjectID(chi.URLParam(r, "id"))
	var req memberReq
	if err := httputils.DecodeJSON(r, &req); err != nil {
		httputils.ErrorJSON(w, 400, err)
		return
	}
	err := h.Svc.AddProjectMember(r.Context(), projects.AddProjectMember{
		ActorID: httputils.ActorIDFrom(r.Context()), ProjectID: id,
		AccountID: req.AccountID, Role: req.Role,
	})
	if err != nil {
		httputils.ErrorJSON(w, 403, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handlers) RemoveMember(w http.ResponseWriter, r *http.Request) {
	id := domain.ProjectID(chi.URLParam(r, "id"))
	var req memberReq
	if err := httputils.DecodeJSON(r, &req); err != nil {
		httputils.ErrorJSON(w, 400, err)
		return
	}
	err := h.Svc.RemoveProjectMember(r.Context(), projects.RemoveProjectMember{
		ActorID:   httputils.ActorIDFrom(r.Context()),
		ProjectID: id,
		AccountID: req.AccountID,
		Role:      req.Role,
	})
	if err != nil {
		httputils.ErrorJSON(w, 403, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
