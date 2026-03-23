package team

import (
	"context"
	"encoding/json"
)

type Handoff struct {
	ID                 string   `json:"id"`
	TaskID             string   `json:"taskId"`
	FromAgent          string   `json:"fromAgent"`
	ToAgent            string   `json:"toAgent"`
	Status             string   `json:"status"`
	WorkSummary        string   `json:"workSummary"`
	Artifacts          []string `json:"artifacts"`
	AcceptanceCriteria []string `json:"acceptanceCriteria"`
	CreatedAt          int64    `json:"createdAt"`
	AcceptedAt         *int64   `json:"acceptedAt,omitempty"`
	AcceptedBy         string   `json:"acceptedBy,omitempty"`
	RejectionReason    string   `json:"rejectionReason,omitempty"`
}

type HandoffService struct {
	caller *PythonCaller
}

func NewHandoffService() *HandoffService {
	return &HandoffService{
		caller: NewPythonCaller(),
	}
}

func (s *HandoffService) Create(ctx context.Context, teamName, taskID, fromAgent, toAgent, workSummary string, artifacts, criteria []string) (*Handoff, error) {
	kwargs := map[string]any{
		"team_name":  teamName,
		"task_id":    taskID,
		"from_agent": fromAgent,
		"to_agent":   toAgent,
	}
	if workSummary != "" {
		kwargs["work_summary"] = workSummary
	}
	if len(artifacts) > 0 {
		kwargs["artifacts"] = artifacts
	}
	if len(criteria) > 0 {
		kwargs["acceptance_criteria"] = criteria
	}

	result, err := s.caller.Call(ctx, "claude_teams.handoff", "create_handoff", kwargs)
	if err != nil {
		return nil, err
	}

	var h Handoff
	if err := json.Unmarshal(result, &h); err != nil {
		return nil, err
	}
	return &h, nil
}

func (s *HandoffService) Read(ctx context.Context, teamName, handoffID string) (*Handoff, error) {
	result, err := s.caller.Call(ctx, "claude_teams.handoff", "read_handoff", map[string]any{
		"team_name":  teamName,
		"handoff_id": handoffID,
	})
	if err != nil {
		return nil, err
	}

	var h Handoff
	if err := json.Unmarshal(result, &h); err != nil {
		return nil, err
	}
	return &h, nil
}

func (s *HandoffService) Accept(ctx context.Context, teamName, handoffID, acceptedBy string) (*Handoff, error) {
	result, err := s.caller.Call(ctx, "claude_teams.handoff", "accept_handoff", map[string]any{
		"team_name":   teamName,
		"handoff_id":  handoffID,
		"accepted_by": acceptedBy,
	})
	if err != nil {
		return nil, err
	}

	var h Handoff
	if err := json.Unmarshal(result, &h); err != nil {
		return nil, err
	}
	return &h, nil
}

func (s *HandoffService) Reject(ctx context.Context, teamName, handoffID, rejectedBy, reason string) (*Handoff, error) {
	result, err := s.caller.Call(ctx, "claude_teams.handoff", "reject_handoff", map[string]any{
		"team_name":   teamName,
		"handoff_id":  handoffID,
		"rejected_by": rejectedBy,
		"reason":      reason,
	})
	if err != nil {
		return nil, err
	}

	var h Handoff
	if err := json.Unmarshal(result, &h); err != nil {
		return nil, err
	}
	return &h, nil
}

func (s *HandoffService) List(ctx context.Context, teamName string, status, toAgent string) ([]Handoff, error) {
	kwargs := map[string]any{
		"team_name": teamName,
	}
	if status != "" {
		kwargs["status"] = status
	}
	if toAgent != "" {
		kwargs["to_agent"] = toAgent
	}

	result, err := s.caller.Call(ctx, "claude_teams.handoff", "list_handoffs", kwargs)
	if err != nil {
		return nil, err
	}

	var handoffs []Handoff
	if err := json.Unmarshal(result, &handoffs); err != nil {
		return nil, err
	}
	return handoffs, nil
}

func (s *HandoffService) GetPendingForAgent(ctx context.Context, teamName, agent string) ([]Handoff, error) {
	result, err := s.caller.Call(ctx, "claude_teams.handoff", "get_pending_handoffs_for_agent", map[string]any{
		"team_name": teamName,
		"agent":     agent,
	})
	if err != nil {
		return nil, err
	}

	var handoffs []Handoff
	if err := json.Unmarshal(result, &handoffs); err != nil {
		return nil, err
	}
	return handoffs, nil
}
