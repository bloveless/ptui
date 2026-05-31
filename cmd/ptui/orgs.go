package main

import (
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"github.com/bloveless/ptui/pangolin"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, desc, id string
}

func (i item) ID() string          { return i.id }
func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type OrgsView struct {
	logger         Logger
	eventGenerator EventGenerator
	orgsErr        error
	orgs           []pangolin.Org
	list           list.Model
}

func NewOrgsView(logger Logger, eg EventGenerator) OrgsView {
	l := list.New(nil, list.NewDefaultDelegate(), 100, 2)
	l.Title = "Organizations"
	return OrgsView{
		logger:         logger,
		eventGenerator: eg,
		list:           l,
	}
}

func (o OrgsView) Init() tea.Cmd {
	return o.eventGenerator.LoadOrgs
}

func (o OrgsView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		o.list.SetSize(msg.Width-h, msg.Height-v)
	case tea.KeyPressMsg:
		switch msg.String() {
		case "enter":
			si := o.list.SelectedItem().(item)
			o.logger.Logf("enter pressed, org: %s id: %s", si.title, si.id)
			return o, func() tea.Msg { return o.eventGenerator.OrgSelected(si.id) }
		}
	case orgsLoaded:
		items := []list.Item{}
		for _, o := range msg.orgs.Data.Orgs {
			items = append(items, item{title: o.Name, desc: o.ID, id: o.ID})
		}
		o.list.SetItems(items)
	case orgsLoadFailed:
		o.orgsErr = msg.err
	}

	return o, nil
}

func (o OrgsView) View() tea.View {
	return tea.NewView(docStyle.Render(o.list.View()))
}
