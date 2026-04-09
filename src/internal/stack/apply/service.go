package deploy

import (
	"context"
	"fmt"
	"strings"

	stack "github.com/PavelMilanov/forge/internal/stack"
)

// TargetResolver преобразует логическое имя цели (cluster/context/endpoint) в ссылку на runtime-цель.
type TargetResolver interface {
	ResolveTarget(ctx context.Context, runtime stack.Runtime, targetName string) (stack.TargetRef, error)
}

// WorkloadRepository предоставляет поиск существующих workload по имени.
type WorkloadRepository interface {
	FindByName(ctx context.Context, runtime stack.Runtime, target stack.TargetRef, workloadName string) (workloadID string, exists bool, err error)
}

// WorkloadApplier выполняет runtime-специфичные операции create/update.
type WorkloadApplier interface {
	Create(ctx context.Context, runtime stack.Runtime, target stack.TargetRef, in stack.ApplyInput) (workloadID string, action string, err error)
	Update(ctx context.Context, runtime stack.Runtime, target stack.TargetRef, workloadID string, in stack.ApplyInput) (action string, err error)
}

// AccessManager назначает права после применения.
type AccessManager interface {
	GrantManagementAccess(ctx context.Context, runtime stack.Runtime, target stack.TargetRef, workloadID string, groups []string) error
}

// Service оркестрирует этап применения подготовленного манифеста.
type Service struct {
	targetResolver TargetResolver
	repository     WorkloadRepository
	applier        WorkloadApplier
	access         AccessManager
}

// NewService создает сервис применения.
func NewService(targetResolver TargetResolver, repository WorkloadRepository, applier WorkloadApplier, access AccessManager) *Service {
	return &Service{
		targetResolver: targetResolver,
		repository:     repository,
		applier:        applier,
		access:         access,
	}
}

// Deploy применяет подготовленный манифест в режиме upsert.
func (s *Service) Deploy(ctx context.Context, in stack.ApplyInput) (stack.ApplyResult, error) {
	if in.Runtime == "" {
		return stack.ApplyResult{}, fmt.Errorf("runtime is required")
	}
	if strings.TrimSpace(in.TargetName) == "" {
		return stack.ApplyResult{}, fmt.Errorf("target name is required")
	}
	if strings.TrimSpace(in.WorkloadName) == "" {
		return stack.ApplyResult{}, fmt.Errorf("workload name is required")
	}
	if strings.TrimSpace(in.ManifestFile) == "" {
		return stack.ApplyResult{}, fmt.Errorf("manifest file is required")
	}
	if s.targetResolver == nil || s.repository == nil || s.applier == nil {
		return stack.ApplyResult{}, stack.ErrNotImplemented
	}

	target, err := s.targetResolver.ResolveTarget(ctx, in.Runtime, in.TargetName)
	if err != nil {
		return stack.ApplyResult{}, err
	}

	existingID, exists, err := s.repository.FindByName(ctx, in.Runtime, target, in.WorkloadName)
	if err != nil {
		return stack.ApplyResult{}, err
	}

	result := stack.ApplyResult{
		Runtime:  in.Runtime,
		TargetID: target.ID,
	}

	if exists {
		result.WorkloadID = existingID
		result.Action, err = s.applier.Update(ctx, in.Runtime, target, existingID, in)
	} else {
		result.WorkloadID, result.Action, err = s.applier.Create(ctx, in.Runtime, target, in)
	}
	if err != nil {
		return stack.ApplyResult{}, err
	}

	if len(in.AccessGroups) > 0 {
		if s.access == nil {
			return stack.ApplyResult{}, stack.ErrNotImplemented
		}
		if err := s.access.GrantManagementAccess(ctx, in.Runtime, target, result.WorkloadID, in.AccessGroups); err != nil {
			return stack.ApplyResult{}, err
		}
	}
	return result, nil
}
