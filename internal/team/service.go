package team

import (
	"context"
)

type Service struct {
	Context *TeamContextService
	Board   *TaskBoardService
	Handoff *HandoffService
	Inbox   *InboxService
}

func NewService() *Service {
	return &Service{
		Context: NewTeamContextService(),
		Board:   NewTaskBoardService(),
		Handoff: NewHandoffService(),
		Inbox:   NewInboxService(),
	}
}

// EnsureDirectory ensures the team directory structure exists
func (s *Service) EnsureDirectory(ctx context.Context, teamName string) error {
	// Create team context to initialize directory structure
	_, err := s.Context.CreateContext(ctx, teamName, "", nil)
	if err != nil {
		// Ignore if already exists
		return nil
	}
	return nil
}
