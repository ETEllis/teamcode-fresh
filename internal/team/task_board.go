package team

import (
	"context"
	"encoding/json"
)

type TaskBoard struct {
	TeamName    string              `json:"teamName"`
	Columns     map[string][]string `json:"columns"`
	Assignments map[string]string   `json:"assignments"`
	Constraints map[string]any      `json:"constraints"`
}

type TaskBoardService struct {
	caller *PythonCaller
}

func NewTaskBoardService() *TaskBoardService {
	return &TaskBoardService{
		caller: NewPythonCaller(),
	}
}

func (s *TaskBoardService) CreateBoard(ctx context.Context, teamName string) (*TaskBoard, error) {
	result, err := s.caller.Call(ctx, "claude_teams.task_board", "create_task_board", map[string]any{
		"team_name": teamName,
	})
	if err != nil {
		return nil, err
	}

	var board TaskBoard
	if err := json.Unmarshal(result, &board); err != nil {
		return nil, err
	}
	return &board, nil
}

func (s *TaskBoardService) ReadBoard(ctx context.Context, teamName string) (*TaskBoard, error) {
	result, err := s.caller.Call(ctx, "claude_teams.task_board", "read_task_board", map[string]any{
		"team_name": teamName,
	})
	if err != nil {
		return nil, err
	}

	var board TaskBoard
	if err := json.Unmarshal(result, &board); err != nil {
		return nil, err
	}
	return &board, nil
}

func (s *TaskBoardService) MoveTask(ctx context.Context, teamName, taskID, fromColumn, toColumn string, agent string) (*TaskBoard, error) {
	kwargs := map[string]any{
		"team_name":   teamName,
		"task_id":     taskID,
		"from_column": fromColumn,
		"to_column":   toColumn,
	}
	if agent != "" {
		kwargs["agent"] = agent
	}

	result, err := s.caller.Call(ctx, "claude_teams.task_board", "move_task", kwargs)
	if err != nil {
		return nil, err
	}

	var board TaskBoard
	if err := json.Unmarshal(result, &board); err != nil {
		return nil, err
	}
	return &board, nil
}

func (s *TaskBoardService) AddTaskToColumn(ctx context.Context, teamName, taskID, column string) (*TaskBoard, error) {
	result, err := s.caller.Call(ctx, "claude_teams.task_board", "add_task_to_column", map[string]any{
		"team_name": teamName,
		"task_id":   taskID,
		"column":    column,
	})
	if err != nil {
		return nil, err
	}

	var board TaskBoard
	if err := json.Unmarshal(result, &board); err != nil {
		return nil, err
	}
	return &board, nil
}

func (s *TaskBoardService) AssignTask(ctx context.Context, teamName, taskID, agent string) (*TaskBoard, error) {
	result, err := s.caller.Call(ctx, "claude_teams.task_board", "assign_task", map[string]any{
		"team_name": teamName,
		"task_id":   taskID,
		"agent":     agent,
	})
	if err != nil {
		return nil, err
	}

	var board TaskBoard
	if err := json.Unmarshal(result, &board); err != nil {
		return nil, err
	}
	return &board, nil
}

func (s *TaskBoardService) GetTaskLocation(ctx context.Context, teamName, taskID string) (string, error) {
	result, err := s.caller.Call(ctx, "claude_teams.task_board", "get_task_location", map[string]any{
		"team_name": teamName,
		"task_id":   taskID,
	})
	if err != nil {
		return "", err
	}

	// Result might be nil if not found
	if string(result) == "null" {
		return "", nil
	}

	var location string
	if err := json.Unmarshal(result, &location); err != nil {
		return "", err
	}
	return location, nil
}

func (s *TaskBoardService) GetBoardSummary(ctx context.Context, teamName string) (map[string]int, error) {
	result, err := s.caller.Call(ctx, "claude_teams.task_board", "get_board_summary", map[string]any{
		"team_name": teamName,
	})
	if err != nil {
		return nil, err
	}

	var summary map[string]int
	if err := json.Unmarshal(result, &summary); err != nil {
		return nil, err
	}
	return summary, nil
}
