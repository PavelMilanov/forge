package stack

import "time"

// Runtime определяет платформу развертывания.
type Runtime string

const (
	RuntimeCompose    Runtime = "compose"
	RuntimeSwarm      Runtime = "swarm"
	RuntimeKubernetes Runtime = "kubernetes"
)

// TargetRef описывает разрешенную целевую среду для конкретного runtime.
type TargetRef struct {
	ID       string
	Name     string
	Runtime  Runtime
	Metadata map[string]string
}

// PlanInput описывает входные параметры генерации манифеста.
type PlanInput struct {
	Project    string
	Template   string
	Runtime    Runtime
	OutputPath string
}

// PlanResult содержит метаданные сгенерированного артефакта.
type PlanResult struct {
	Runtime     Runtime
	OutputPath  string
	ContentSHA  string
	GeneratedAt time.Time
}

// ApplyInput описывает запрос на применение подготовленного манифеста.
type ApplyInput struct {
	Runtime      Runtime
	TargetName   string
	WorkloadName string
	ManifestFile string
	Prune        bool
	PullImage    bool
	AccessGroups []string
}

// ApplyResult содержит итог выполнения операции применения.
type ApplyResult struct {
	Runtime    Runtime
	TargetID   string
	WorkloadID string
	Action     string
}
