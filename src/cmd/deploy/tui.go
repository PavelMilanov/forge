package deploy

import (
	"fmt"
	"time"

	"github.com/PavelMilanov/forge/agent"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type statusMsg string
type doneMsg struct{}
type errMsg struct{ err error }

type model struct {
	spinner spinner.Model
	status  string
	final   string
	err     error
	done    bool
	// Зависимости
	cfg      agent.Agent
	stack    string
	template string
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		func() tea.Msg { return statusMsg("Подготовка к деплою...") },
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case statusMsg:
		m.status = string(msg)
		if m.status == "Подготовка к деплою..." {
			return m, func() tea.Msg {
				time.Sleep(500 * time.Millisecond) // Пауза для красоты
				return statusMsg("Отправка запроса в Portainer API...")
			}
		}
		if m.status == "Отправка запроса в Portainer API..." {
			return m, m.DeployTask
		}

	case doneMsg:
		m.done = true
		m.final = fmt.Sprintf(
			"Стек [%s] успешно развернут",
			m.stack,
		)
		return m, tea.Quit

	case errMsg:
		m.err = msg.err
		return m, tea.Quit

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf(
			" %s %s",
			m.spinner.Style.Render("✖"),
			lipgloss.NewStyle().
				Foreground(lipgloss.Color("196")).
				Bold(true).
				Render(m.err.Error()),
		)
	}

	if m.done {
		return fmt.Sprintf(
			" %s",
			m.spinner.Style.Render("✔"),
		)
	}

	return fmt.Sprintf(
		" %s %s",
		m.spinner.View(),
		m.status,
	)
}

func (m model) DeployTask() tea.Msg {
	var err error
	if m.template != "" {
		err = m.cfg.CreateStack(m.stack, m.template)
	} else {
		err = m.cfg.DeployStack(m.stack)
	}

	if err != nil {
		return errMsg{err}
	}
	return doneMsg{}
}
