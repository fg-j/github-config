package internal

import (
	"encoding/json"
	"fmt"
	"time"
)

type RepositoryContainer struct {
	Repository Repository
	Error      error
}

type Repository struct {
	Name  string `json:"full_name"`
	URL   string `json:"url"`
	Owner struct {
		Login string `json:"login"`
	} `json:"owner"`
}

type Organization struct {
	Name string
}

//go:generate faux --interface Clock --output fakes/clock.go
type Clock interface {
	Now() time.Time
}

func (o *Organization) GetRepos(client APIClient) ([]Repository, error) {
	body, err := client.Get(fmt.Sprintf("orgs/%s/repos", o.Name), "per_page=100")
	if err != nil {
		return nil, fmt.Errorf("failed getting org repos: %s", err)
	}

	repos := []Repository{}
	err = json.Unmarshal(body, &repos)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal response: %s\n  : %s", string(body), err)
	}
	return repos, nil
}

func (r *Repository) GetRecentIssues(client Client, clock Clock) ([]Issue, error) {
	timeString := clock.Now().UTC().Add(-30 * 24 * time.Hour).Format(time.RFC3339)

	body, err := client.Get(fmt.Sprintf("/repos/%s/issues", r.Name),
		"per_page=100",
		fmt.Sprintf("since=%s", timeString))
	if err != nil {
		panic(err)
	}

	issues := []Issue{}
	err = json.Unmarshal(body, &issues)
	if err != nil {
		panic(err)
	}
	return issues, nil
}

func (r *Repository) GetFirstContactTimes(client APIClient, output chan TimeContainer) {

}
