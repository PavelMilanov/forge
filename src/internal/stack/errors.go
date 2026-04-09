package stack

import "errors"

var (
	// ErrNotImplemented помечает точки архитектуры, которые пока оставлены на этап интеграции.
	ErrNotImplemented = errors.New("stack architecture skeleton: implementation is not wired yet")
)
