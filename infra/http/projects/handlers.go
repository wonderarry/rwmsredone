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

// CreateProjectReq is the request payload for creating a project.
type CreateProjectReq struct {
	Name        string `json:"name"        example:"Thesis Project"`
	Theme       string `json:"theme"       example:"Machine Learning"`
	Description string `json:"description" example:"Exploring X with Y"`
}

// Create godoc
// @Summary      Create a project
// @Tags         projects
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        request  body      CreateProjectReq  true  "create project payload"
// @Success      200      {object}  map[string]string  "project_id"
// @Failure      400      {object}  map[string]string
// @Failure      403      {object}  map[string]string
// @Router       /api/projects/ [post]
func (h *Handlers) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateProjectReq
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

// ProjectBrief is a lightweight response item for listing projects.
type ProjectBrief struct {
	ID          domain.ProjectID `json:"id"          swaggertype:"string" example:"prj_123"`
	Name        string           `json:"name"        example:"Thesis Project"`
	Theme       string           `json:"theme"       example:"Machine Learning"`
	Description string           `json:"description" example:"Exploring X with Y"`
}

// ListMine godoc
// @Summary      List projects owned/by the current user
// @Tags         projects
// @Security     BearerAuth
// @Produce      json
// @Success      200  {array}  ProjectBrief
// @Failure      400  {object} map[string]string
// @Router       /api/projects/ [get]
func (h *Handlers) ListMine(w http.ResponseWriter, r *http.Request) {
	items, err := h.Svc.ListMyProjects(r.Context(), httputils.ActorIDFrom(r.Context()))
	if err != nil {
		httputils.ErrorJSON(w, 400, err)
		return
	}
	httputils.WriteJSON(w, 200, items)
}

// EditProjectReq is the request payload for editing project metadata.
type EditProjectReq struct {
	Name        string `json:"name"        example:"Thesis Project v2"`
	Theme       string `json:"theme"       example:"Deep Learning"`
	Description string `json:"description" example:"Updated description"`
}

// EditMeta godoc
// @Summary      Edit project metadata
// @Tags         projects
// @Security     BearerAuth
// @Accept       json
// @Param        id       path   string          true  "Project ID"
// @Param        request  body   EditProjectReq  true  "edit project payload"
// @Success      204      {string} string        "No Content"
// @Failure      400      {object} map[string]string
// @Failure      403      {object} map[string]string
// @Router       /api/projects/{id} [put]
func (h *Handlers) EditMeta(w http.ResponseWriter, r *http.Request) {
	id := domain.ProjectID(chi.URLParam(r, "id"))
	var req EditProjectReq
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

// MemberReq is the request payload for adding/removing a project member.
type MemberReq struct {
	AccountID domain.AccountID   `json:"account_id" swaggertype:"string" example:"acc_456"`
	Role      domain.ProjectRole `json:"role"       swaggertype:"string" example:"Contributor"`
}

// AddMember godoc
// @Summary      Add a member to the project
// @Tags         projects
// @Security     BearerAuth
// @Accept       json
// @Param        id       path      string    true  "Project ID"
// @Param        request  body      MemberReq true  "member payload"
// @Success      204      {string}  string    "No Content"
// @Failure      400      {object}  map[string]string
// @Failure      403      {object}  map[string]string
// @Router       /api/projects/{id}/members [post]
func (h *Handlers) AddMember(w http.ResponseWriter, r *http.Request) {
	id := domain.ProjectID(chi.URLParam(r, "id"))
	var req MemberReq
	if err := httputils.DecodeJSON(r, &req); err != nil {
		httputils.ErrorJSON(w, 400, err)
		return
	}
	err := h.Svc.AddProjectMember(r.Context(), projects.AddProjectMember{
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

// RemoveMember godoc
// @Summary      Remove a member from the project
// @Tags         projects
// @Security     BearerAuth
// @Accept       json
// @Param        id       path      string    true  "Project ID"
// @Param        request  body      MemberReq true  "member payload"
// @Success      204      {string}  string    "No Content"
// @Failure      400      {object}  map[string]string
// @Failure      403      {object}  map[string]string
// @Router       /api/projects/{id}/members [delete]
func (h *Handlers) RemoveMember(w http.ResponseWriter, r *http.Request) {
	id := domain.ProjectID(chi.URLParam(r, "id"))
	var req MemberReq
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
