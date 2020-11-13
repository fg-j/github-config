package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type RepositoryContainer struct {
	Repository Repository
	Error      error
}

type Repository struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Owner struct {
		Login string `json:"login"`
	} `json:"owner"`
}

type Organization struct {
	Name string
}

func (o *Organization) GetRepos(client APIClient) ([]Repository, error) {
	response, err := client.Get(fmt.Sprintf("orgs/%s/repos", o.Name), "per_page=100")
	if err != nil {
		return nil, fmt.Errorf("failed getting org repos: %s", err)
	}

	body, _ := ioutil.ReadAll(response.Body)

	repos := []Repository{}
	err = json.Unmarshal(body, &repos)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal response: %s\n  : %s", string(body), err)
	}
	return repos, nil
}

func (r *Repository) GetRecentIssues(client APIClient) ([]Issue, error) {
	return nil, nil
}

func (r *Repository) GetFirstContactTimes(client APIClient, output chan TimeContainer) {

}
