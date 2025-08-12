package domain

import "time"

type Event interface {
	Topic() string
	OccuredAt() time.Time
}

type baseEvent struct {
	at time.Time
}

func (b baseEvent) OccuredAt() time.Time {
	return b.at
}

type ProjectCreated struct {
	baseEvent
	ProjectID ProjectID
	Name      string
	By        AccountID
}

func (ProjectCreated) Topic() string {
	return "ProjectCreated.v1"
}

type ProjectMemberAdded struct {
	baseEvent
	ProjectID ProjectID
	AccountID AccountID
	AddedBy   AccountID
	RoleKey   string
}

func (ProjectMemberAdded) Topic() string {
	return "ProjectMemberAdded.v1"
}

type ProcessCreated struct {
	baseEvent
	ProcessID   ProcessID
	ProjectID   ProjectID
	TemplateKey TemplateKey
	Name        string
	By          AccountID
}

func (ProcessCreated) Topic() string {
	return "ProcessCreated.v1"
}

type ProcessMemberAdded struct {
	baseEvent
	ProcessID ProcessID
	AccountID AccountID
	RoleKey   string
	AddedBy   AccountID
}

func (ProcessMemberAdded) Topic() string {
	return "ProcessMemberAdded.v1"
}

type ApprovalRecorded struct {
	baseEvent
	ProcessID ProcessID
	StageKey  StageKey
	ByAccount AccountID
	ByRole    string
	Decision  Decision
}

func (ApprovalRecorded) Topic() string {
	return "ApprovalRecorded.v1"
}

type StageAdvanced struct {
	baseEvent
	ProcessID ProcessID
	From      StageKey
	To        StageKey
	ByAccount AccountID
	Reason    string
}

func (StageAdvanced) Topic() string {
	return "StageAdvanced.v1"
}

type ProcessFinalized struct {
	baseEvent
	ProcessID ProcessID
}

func (ProcessFinalized) Topic() string {
	return "ProcessFinalized.v1"
}
