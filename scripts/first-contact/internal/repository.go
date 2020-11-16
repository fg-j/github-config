package internal

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
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
		return nil, fmt.Errorf("getting recent issues: %s", err)
	}

	issues := []Issue{}
	err = json.Unmarshal(body, &issues)
	if err != nil {
		return nil, fmt.Errorf("getting recent issues: could not unmarshal JSON '%s' : %s", string(body), err)
	}
	return issues, nil
}

func (r *Repository) GetFirstContactTimes(client Client, issues []CommentGetter, clock Clock, output chan TimeContainer) {
	for _, issue := range issues {
		if strings.Contains(issue.GetUserLogin(), "bot") {
			continue
		}

		comment, err := issue.GetFirstReply()

		if err != nil {
			output <- TimeContainer{Error: fmt.Errorf("error getting first contact times for repo %s: %s", r.Name, err)}
			close(output)
			return
		}
		// TODO: feels a little weird converting to string and then back again
		// TODO: decide if we actually want to include issues without comments on them
		if comment.CreatedAt == "" {
			comment.CreatedAt = clock.Now().UTC().Format(time.RFC3339)
		}

		replyCreated, err := time.Parse(time.RFC3339, comment.CreatedAt)
		if err != nil {
			panic(err)
		}
		issueCreated, err := time.Parse(time.RFC3339, issue.GetCreatedAt())
		if err != nil {
			panic(err)
		}
		replyTime := math.Round(replyCreated.Sub(issueCreated).Minutes())
		output <- TimeContainer{Time: replyTime, Error: nil}
	}
	close(output)
}
