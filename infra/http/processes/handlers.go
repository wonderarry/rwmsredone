package processesapi

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/wonderarry/rwmsredone/infra/http/httputils"
	"github.com/wonderarry/rwmsredone/internal/app/processes"
	"github.com/wonderarry/rwmsredone/internal/domain"
)

type Handlers struct{ Svc processes.Service }

func New(svc processes.Service) *Handlers { return &Handlers{Svc: svc} }

// CreateReq is the request payload to create a process.
type CreateReq struct {
	ProjectID   domain.ProjectID   `json:"project_id"   swaggertype:"string" example:"prj_123"`
	TemplateKey domain.TemplateKey `json:"template_key" swaggertype:"string" example:"thesis-review"`
	Name        string             `json:"name"         example:"Thesis Approval – Spring"`
}

// Create godoc
// @Summary      Create a new process
// @Tags         processes
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        request  body      CreateReq  true  "process creation payload"
// @Success      200      {object}  map[string]string  "process_id"
// @Failure      400      {object}  map[string]string
// @Failure      403      {object}  map[string]string
// @Router       /api/processes/ [post]
func (h *Handlers) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateReq
	if err := httputils.DecodeJSON(r, &req); err != nil {
		httputils.ErrorJSON(w, 400, err)
		return
	}
	id, err := h.Svc.CreateProcess(r.Context(), processes.CreateProcess{
		ActorID:     httputils.ActorIDFrom(r.Context()),
		ProjectID:   req.ProjectID,
		TemplateKey: req.TemplateKey,
		Name:        req.Name,
	})
	if err != nil {
		httputils.ErrorJSON(w, 403, err)
		return
	}
	httputils.WriteJSON(w, 200, map[string]string{"process_id": string(id)})
}

// MemberReq is the request payload to add or remove a process member.
type MemberReq struct {
	AccountID domain.AccountID   `json:"account_id" swaggertype:"string" example:"acc_456"`
	Role      domain.ProcessRole `json:"role"       swaggertype:"string" example:"Reviewer"`
}

// AddMember godoc
// @Summary      Add a member to the process
// @Tags         processes
// @Security     BearerAuth
// @Accept       json
// @Param        pid      path      string    true  "Process ID"
// @Param        request  body      MemberReq true  "member payload"
// @Success      204      {string}  string    "No Content"
// @Failure      400      {object}  map[string]string
// @Failure      403      {object}  map[string]string
// @Router       /api/processes/{pid}/members [post]
func (h *Handlers) AddMember(w http.ResponseWriter, r *http.Request) {
	pid := domain.ProcessID(chi.URLParam(r, "pid"))
	var req MemberReq
	if err := httputils.DecodeJSON(r, &req); err != nil {
		httputils.ErrorJSON(w, 400, err)
		return
	}
	err := h.Svc.AddMember(r.Context(), processes.AddProcessMember{
		ActorID:   httputils.ActorIDFrom(r.Context()),
		ProcessID: pid,
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
// @Summary      Remove a member from the process
// @Tags         processes
// @Security     BearerAuth
// @Accept       json
// @Param        pid      path      string    true  "Process ID"
// @Param        request  body      MemberReq true  "member payload"
// @Success      204      {string}  string    "No Content"
// @Failure      400      {object}  map[string]string
// @Failure      403      {object}  map[string]string
// @Router       /api/processes/{pid}/members [delete]
func (h *Handlers) RemoveMember(w http.ResponseWriter, r *http.Request) {
	pid := domain.ProcessID(chi.URLParam(r, "pid"))
	var req MemberReq
	if err := httputils.DecodeJSON(r, &req); err != nil {
		httputils.ErrorJSON(w, 400, err)
		return
	}
	err := h.Svc.RemoveMember(r.Context(), processes.RemoveProcessMember{
		ActorID:   httputils.ActorIDFrom(r.Context()),
		ProcessID: pid,
		AccountID: req.AccountID,
		Role:      req.Role,
	})
	if err != nil {
		httputils.ErrorJSON(w, 403, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// ApprovalReq is the request payload to record an approval/decision.
type ApprovalReq struct {
	Decision  domain.Decision    `json:"decision"   swaggertype:"string" enums:"approve,reject" example:"approve"`
	Comment   string             `json:"comment"    example:"Looks good"`
	ActorRole domain.ProcessRole `json:"actor_role" swaggertype:"string" example:"Supervisor"`
}

// RecordApproval godoc
// @Summary      Record an approval/decision for a process
// @Tags         processes
// @Security     BearerAuth
// @Accept       json
// @Param        pid      path      string       true  "Process ID"
// @Param        request  body      ApprovalReq  true  "approval payload"
// @Success      204      {string}  string       "No Content"
// @Failure      400      {object}  map[string]string
// @Failure      403      {object}  map[string]string
// @Router       /api/processes/{pid}/approvals [post]
func (h *Handlers) RecordApproval(w http.ResponseWriter, r *http.Request) {
	pid := domain.ProcessID(chi.URLParam(r, "pid"))
	var req ApprovalReq
	if err := httputils.DecodeJSON(r, &req); err != nil {
		httputils.ErrorJSON(w, 400, err)
		return
	}
	err := h.Svc.RecordApproval(r.Context(), processes.RecordApproval{
		ActorID:   httputils.ActorIDFrom(r.Context()),
		ProcessID: pid,
		ActorRole: req.ActorRole,
		Decision:  req.Decision,
		Comment:   req.Comment,
	})
	if err != nil {
		httputils.ErrorJSON(w, 403, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// TODO: GetProcess, GetProcessGraph, ListApprovals
