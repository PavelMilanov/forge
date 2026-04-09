package prepare

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	stack "github.com/PavelMilanov/forge/internal/stack"
)

// Source загружает исходный контент шаблона/проекта для рендера.
type Source interface {
	Load(ctx context.Context, in stack.PlanInput) (string, error)
}

// Renderer подставляет переменные в исходный контент и возвращает итоговый манифест.
type Renderer interface {
	Render(ctx context.Context, raw string, in stack.PlanInput) (string, error)
}

// Validator проверяет итоговый манифест перед записью.
type Validator interface {
	Validate(ctx context.Context, runtime stack.Runtime, content string) error
}

// Writer сохраняет сгенерированный контент.
type Writer interface {
	Write(ctx context.Context, path string, content string) error
}

// Service оркестрирует этап подготовки.
type Service struct {
	source    Source
	renderer  Renderer
	validator Validator
	writer    Writer
	now       func() time.Time
}

// NewService создает сервис подготовки с внедренными зависимостями.
func NewService(source Source, renderer Renderer, validator Validator, writer Writer) *Service {
	return &Service{
		source:    source,
		renderer:  renderer,
		validator: validator,
		writer:    writer,
		now:       time.Now,
	}
}

// Prepare формирует файл манифеста и метаданные.
func (s *Service) Prepare(ctx context.Context, in stack.PlanInput) (stack.PlanResult, error) {
	if strings.TrimSpace(in.Project) == "" {
		return stack.PlanResult{}, fmt.Errorf("project is required")
	}
	if strings.TrimSpace(in.OutputPath) == "" {
		return stack.PlanResult{}, fmt.Errorf("output path is required")
	}
	if in.Runtime == "" {
		return stack.PlanResult{}, fmt.Errorf("runtime is required")
	}
	if s.source == nil || s.renderer == nil || s.validator == nil || s.writer == nil {
		return stack.PlanResult{}, stack.ErrNotImplemented
	}

	raw, err := s.source.Load(ctx, in)
	if err != nil {
		return stack.PlanResult{}, err
	}

	content, err := s.renderer.Render(ctx, raw, in)
	if err != nil {
		return stack.PlanResult{}, err
	}

	if err := s.validator.Validate(ctx, in.Runtime, content); err != nil {
		return stack.PlanResult{}, err
	}

	if err := s.writer.Write(ctx, in.OutputPath, content); err != nil {
		return stack.PlanResult{}, err
	}

	sum := sha256.Sum256([]byte(content))
	return stack.PlanResult{
		Runtime:     in.Runtime,
		OutputPath:  in.OutputPath,
		ContentSHA:  hex.EncodeToString(sum[:]),
		GeneratedAt: s.now().UTC(),
	}, nil
}
