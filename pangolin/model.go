package pangolin

import "time"

type ListOrgsResponse struct {
	Data    OrgData
	Success bool
	Error   bool
	Message string
	Status  int32
}

type OrgData struct {
	Orgs []Org
}

type Org struct {
	ID        string `json:"orgId"`
	Name      string
	CreatedAt time.Time
}
