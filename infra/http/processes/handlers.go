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

type createReq struct {
	ProjectID   domain.ProjectID   `json:"project_id"`
	TemplateKey domain.TemplateKey `json:"template_key"`
	Name        string             `json:"name"`
}

func (h *Handlers) Create(w http.ResponseWriter, r *http.Request) {
	var req createReq
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

type memberReq struct {
	AccountID domain.AccountID   `json:"account_id"`
	Role      domain.ProcessRole `json:"role"`
}

func (h *Handlers) AddMember(w http.ResponseWriter, r *http.Request) {
	pid := domain.ProcessID(chi.URLParam(r, "pid"))
	var req memberReq
	if err := httputils.DecodeJSON(r, &req); err != nil {
		httputils.ErrorJSON(w, 400, err)
		return
	}
	err := h.Svc.AddMember(r.Context(), processes.AddProcessMember{
		ActorID: httputils.ActorIDFrom(r.Context()), ProcessID: pid,
		AccountID: req.AccountID, Role: req.Role,
	})
	if err != nil {
		httputils.ErrorJSON(w, 403, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handlers) RemoveMember(w http.ResponseWriter, r *http.Request) {
	pid := domain.ProcessID(chi.URLParam(r, "pid"))
	var req memberReq
	if err := httputils.DecodeJSON(r, &req); err != nil {
		httputils.ErrorJSON(w, 400, err)
		return
	}
	err := h.Svc.RemoveMember(r.Context(), processes.RemoveProcessMember{
		ActorID: httputils.ActorIDFrom(r.Context()), ProcessID: pid,
		AccountID: req.AccountID, Role: req.Role,
	})
	if err != nil {
		httputils.ErrorJSON(w, 403, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

type approvalReq struct {
	Decision  domain.Decision    `json:"decision"` // "approve" | "reject"
	Comment   string             `json:"comment"`
	ActorRole domain.ProcessRole `json:"actor_role"`
}

func (h *Handlers) RecordApproval(w http.ResponseWriter, r *http.Request) {
	pid := domain.ProcessID(chi.URLParam(r, "pid"))
	var req approvalReq
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
