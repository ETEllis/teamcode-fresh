package team

import (
	"context"
	"encoding/json"
	"time"
)

type TeamContext struct {
	TeamName         string           `json:"teamName"`
	Charter          string           `json:"charter"`
	Goals            []Goal           `json:"goals"`
	Roles            map[string]Role  `json:"roles"`
	WorkingAgreement WorkingAgreement `json:"workingAgreement"`
	CreatedAt        int64            `json:"createdAt"`
	UpdatedAt        int64            `json:"updatedAt"`
}

type Goal struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   int64  `json:"createdAt,omitempty"`
}

type Role struct {
	Name         string `json:"name"`
	Responsible  string `json:"responsible"`
	CurrentFocus string `json:"currentFocus,omitempty"`
	Agent        string `json:"agent,omitempty"`
}

type WorkingAgreement struct {
	CommitMessageFormat string   `json:"commitMessageFormat"`
	MaxWIP              int      `json:"maxWip"`
	HandoffRequires     []string `json:"handoffRequires"`
	ReviewRequired      bool     `json:"reviewRequired"`
	ApprovalRequiredFor []string `json:"approvalRequiredFor,omitempty"`
}

type TeamContextService struct {
	caller *PythonCaller
}

func NewTeamContextService() *TeamContextService {
	return &TeamContextService{
		caller: NewPythonCaller(),
	}
}

func (s *TeamContextService) CreateContext(ctx context.Context, teamName, charter string, roles map[string]Role) (*TeamContext, error) {
	rolesMap := make(map[string]any)
	for k, v := range roles {
		rolesMap[k] = v
	}

	result, err := s.caller.Call(ctx, "claude_teams.team_context", "create_team_context", map[string]any{
		"team_name": teamName,
		"charter":   charter,
		"roles":     rolesMap,
	})
	if err != nil {
		return nil, err
	}

	var tc TeamContext
	if err := json.Unmarshal(result, &tc); err != nil {
		return nil, err
	}
	return &tc, nil
}

func (s *TeamContextService) ReadContext(ctx context.Context, teamName string) (*TeamContext, error) {
	result, err := s.caller.Call(ctx, "claude_teams.team_context", "read_team_context", map[string]any{
		"team_name": teamName,
	})
	if err != nil {
		return nil, err
	}

	var tc TeamContext
	if err := json.Unmarshal(result, &tc); err != nil {
		return nil, err
	}
	return &tc, nil
}

func (s *TeamContextService) UpdateContext(ctx context.Context, teamName string, updates map[string]any) (*TeamContext, error) {
	updates["team_name"] = teamName
	result, err := s.caller.Call(ctx, "claude_teams.team_context", "update_team_context", updates)
	if err != nil {
		return nil, err
	}

	var tc TeamContext
	if err := json.Unmarshal(result, &tc); err != nil {
		return nil, err
	}
	return &tc, nil
}

func (s *TeamContextService) AddRole(ctx context.Context, teamName, roleName string, role Role) (*TeamContext, error) {
	roleMap := map[string]any{
		"name":        role.Name,
		"responsible": role.Responsible,
	}
	if role.CurrentFocus != "" {
		roleMap["current_focus"] = role.CurrentFocus
	}
	if role.Agent != "" {
		roleMap["agent"] = role.Agent
	}

	result, err := s.caller.Call(ctx, "claude_teams.team_context", "add_role", map[string]any{
		"team_name": teamName,
		"role_name": roleName,
		"role":      roleMap,
	})
	if err != nil {
		return nil, err
	}

	var tc TeamContext
	if err := json.Unmarshal(result, &tc); err != nil {
		return nil, err
	}
	return &tc, nil
}

func (s *TeamContextService) AssignAgentToRole(ctx context.Context, teamName, roleName, agentName string) (*TeamContext, error) {
	result, err := s.caller.Call(ctx, "claude_teams.team_context", "assign_agent_to_role", map[string]any{
		"team_name":  teamName,
		"role_name":  roleName,
		"agent_name": agentName,
	})
	if err != nil {
		return nil, err
	}

	var tc TeamContext
	if err := json.Unmarshal(result, &tc); err != nil {
		return nil, err
	}
	return &tc, nil
}

func (s *TeamContextService) AddGoal(ctx context.Context, teamName string, goal Goal) (*TeamContext, error) {
	goalMap := map[string]any{
		"id":          goal.ID,
		"description": goal.Description,
		"status":      goal.Status,
	}
	if goal.CreatedAt == 0 {
		goalMap["created_at"] = int(time.Now().UnixMilli())
	}

	result, err := s.caller.Call(ctx, "claude_teams.team_context", "add_goal", map[string]any{
		"team_name": teamName,
		"goal":      goalMap,
	})
	if err != nil {
		return nil, err
	}

	var tc TeamContext
	if err := json.Unmarshal(result, &tc); err != nil {
		return nil, err
	}
	return &tc, nil
}

func (s *TeamContextService) UpdateGoalStatus(ctx context.Context, teamName, goalID, status string) (*TeamContext, error) {
	result, err := s.caller.Call(ctx, "claude_teams.team_context", "update_goal_status", map[string]any{
		"team_name": teamName,
		"goal_id":   goalID,
		"status":    status,
	})
	if err != nil {
		return nil, err
	}

	var tc TeamContext
	if err := json.Unmarshal(result, &tc); err != nil {
		return nil, err
	}
	return &tc, nil
}
