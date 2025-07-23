package model

import (
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github/SamedArslan28/statix/services"
)

type menuOption int
type speedTestDoneMsg string

const (
	optionSpeedTest menuOption = iota
	optionKernelInformation
	optionCpuInformation
)

type appState struct {
	cursor     int
	choices    []string
	selected   *menuOption
	outputText string
	loading    bool
	spinner    spinner.Model
}

func InitialModel() tea.Model {
	s := spinner.New()
	s.Spinner = spinner.Dot

	return appState{
		choices: []string{"🚀 Speed Test", "🧠 Kernel Info", "💻 CPU Information"},
		spinner: s,
	}
}

func (a appState) Init() tea.Cmd {
	return nil
}

func (a appState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	if a.loading {
		var cmd tea.Cmd
		a.spinner, cmd = a.spinner.Update(msg)
		cmds = append(cmds, cmd)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return a, tea.Quit

		case "up", "k":
			if !a.loading && a.cursor > 0 {
				a.cursor--
			}
		case "down", "j":
			if !a.loading && a.cursor < len(a.choices)-1 {
				a.cursor++
			}
		case "enter":
			if a.outputText != "" {
				a.outputText = ""
				a.selected = nil
				return a, nil
			}

			choice := menuOption(a.cursor)
			a.selected = &choice

			switch choice {
			case optionSpeedTest:
				a.loading = true
				cmds = append(cmds, a.spinner.Tick, runSpeedTestCmd())
			case optionKernelInformation:
				a.outputText = services.GetPlatformInfo()
			case optionCpuInformation:
				a.outputText = services.GetCpuInfo()
			}
		}

	case speedTestDoneMsg:
		a.loading = false
		a.outputText = string(msg)
	}

	return a, tea.Batch(cmds...)
}
func (a appState) View() string {
	if a.loading {
		return fmt.Sprintf("\n\n%s Running speed test...\n", a.spinner.View())
	}

	if a.outputText != "" {
		return fmt.Sprintf("%s\n\nPress [Enter] to return to menu.", a.outputText)
	}

	s := "📊 Select an option:\n\n"
	for i, choice := range a.choices {
		cursor := " "
		if a.cursor == i {
			cursor = "➤"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	s += "\n↑/↓ to move, [Enter] to select, [q] to quit"
	return s
}

func runSpeedTestCmd() tea.Cmd {
	return func() tea.Msg {
		result := services.RunSpeedTest()
		return speedTestDoneMsg(result)
	}
}
