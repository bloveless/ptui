package main

import (
	tea "charm.land/bubbletea/v2"
)

type model struct {
	logger         Logger
	active         tea.Model
	eventGenerator EventGenerator
}

type modelOption func(m *modelArgs)

func WithLogger(l Logger) modelOption {
	return func(m *modelArgs) {
		m.logger = l
	}
}

type modelArgs struct {
	logger Logger
}

func initialModel(eg EventGenerator, opts ...modelOption) model {
	ma := modelArgs{}
	for _, o := range opts {
		o(&ma)
	}
	m := model{
		// Where to write log lines
		logger: ma.logger,

		active: NewOrgsView(ma.logger, eg),

		// The pangolin api client
		eventGenerator: eg,
	}
	return m
}

func (m model) Init() tea.Cmd {
	return m.active.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.logger.Logf("received message: %v", msg)
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	var activeCmd tea.Cmd
	m.active, activeCmd = m.active.Update(msg)

	return m, tea.Batch(activeCmd)
}

func (m model) View() tea.View {
	v := m.active.View()
	v.AltScreen = true
	return v
}
