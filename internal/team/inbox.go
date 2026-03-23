package team

import (
	"context"
	"encoding/json"
)

type InboxMessage struct {
	From_     string `json:"from"`
	Text      string `json:"text"`
	Timestamp string `json:"timestamp"`
	Read      bool   `json:"read"`
	Summary   string `json:"summary,omitempty"`
	Color     string `json:"color,omitempty"`
}

type InboxService struct {
	caller *PythonCaller
}

func NewInboxService() *InboxService {
	return &InboxService{
		caller: NewPythonCaller(),
	}
}

func (s *InboxService) ReadInbox(ctx context.Context, teamName, agentName string, unreadOnly bool) ([]InboxMessage, error) {
	result, err := s.caller.Call(ctx, "claude_teams.messaging", "read_inbox", map[string]any{
		"team_name":    teamName,
		"agent_name":   agentName,
		"unread_only":  unreadOnly,
		"mark_as_read": false,
	})
	if err != nil {
		return nil, err
	}

	var messages []InboxMessage
	if err := json.Unmarshal(result, &messages); err != nil {
		return nil, err
	}
	return messages, nil
}

func (s *InboxService) SendMessage(ctx context.Context, teamName, from, to, content, summary string) error {
	return s.caller.CallSimple(ctx, "claude_teams.messaging", "send_message", map[string]any{
		"team_name": teamName,
		"from":      from,
		"to":        to,
		"content":   content,
		"summary":   summary,
	})
}

func (s *InboxService) Broadcast(ctx context.Context, teamName, from, content, summary string) error {
	return s.caller.CallSimple(ctx, "claude_teams.messaging", "send_broadcast", map[string]any{
		"team_name": teamName,
		"from":      from,
		"content":   content,
		"summary":   summary,
	})
}
