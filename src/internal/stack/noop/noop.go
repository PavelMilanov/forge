package noop

import (
	"context"

	stack "github.com/PavelMilanov/forge/internal/stack"
)

// PrepareSource — заглушка зависимости для этапа подготовки.
type PrepareSource struct{}

func (PrepareSource) Load(ctx context.Context, in stack.PlanInput) (string, error) {
	return "", stack.ErrNotImplemented
}

// Renderer — заглушка зависимости для этапа подготовки.
type Renderer struct{}

func (Renderer) Render(ctx context.Context, raw string, in stack.PlanInput) (string, error) {
	return "", stack.ErrNotImplemented
}

// Validator — заглушка зависимости для этапа подготовки.
type Validator struct{}

func (Validator) Validate(ctx context.Context, runtime stack.Runtime, content string) error {
	return stack.ErrNotImplemented
}

// Writer — заглушка зависимости для этапа подготовки.
type Writer struct{}

func (Writer) Write(ctx context.Context, path string, content string) error {
	return stack.ErrNotImplemented
}

// TargetResolver — заглушка зависимости для этапа применения.
type TargetResolver struct{}

func (TargetResolver) ResolveTarget(ctx context.Context, runtime stack.Runtime, targetName string) (stack.TargetRef, error) {
	return stack.TargetRef{}, stack.ErrNotImplemented
}

// WorkloadRepository — заглушка зависимости для этапа применения.
type WorkloadRepository struct{}

func (WorkloadRepository) FindByName(ctx context.Context, runtime stack.Runtime, target stack.TargetRef, workloadName string) (string, bool, error) {
	return "", false, stack.ErrNotImplemented
}

// WorkloadApplier — заглушка зависимости для этапа применения.
type WorkloadApplier struct{}

func (WorkloadApplier) Create(ctx context.Context, runtime stack.Runtime, target stack.TargetRef, in stack.ApplyInput) (string, string, error) {
	return "", "", stack.ErrNotImplemented
}

func (WorkloadApplier) Update(ctx context.Context, runtime stack.Runtime, target stack.TargetRef, workloadID string, in stack.ApplyInput) (string, error) {
	return "", stack.ErrNotImplemented
}

// AccessManager — заглушка зависимости для этапа применения.
type AccessManager struct{}

func (AccessManager) GrantManagementAccess(ctx context.Context, runtime stack.Runtime, target stack.TargetRef, workloadID string, groups []string) error {
	return stack.ErrNotImplemented
}
