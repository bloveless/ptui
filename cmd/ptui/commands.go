package main

import (
	tea "charm.land/bubbletea/v2"

	"github.com/bloveless/ptui/pangolin"
)

type (
	orgsLoaded struct {
		orgs pangolin.ListOrgsResponse
	}
	orgsLoadFailed struct {
		err error
	}
	orgSelected struct {
		id string
	}
)

type EventGenerator struct {
	logger Logger
	api    pangolin.API
}

func NewEventGenerator(logger Logger, api pangolin.API) EventGenerator {
	return EventGenerator{
		logger: logger,
		api:    api,
	}
}

func (e EventGenerator) LoadOrgs() tea.Msg {
	e.logger.Logf("listing orgs")
	orgs, err := e.api.ListOrgs()
	if err != nil {
		return orgsLoadFailed{err}
	}
	e.logger.Logf("orgs loaded: %v", orgs)
	return orgsLoaded{orgs}
}

func (e EventGenerator) OrgSelected(id string) tea.Msg {
	return orgSelected{id}
}
