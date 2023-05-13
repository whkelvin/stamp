package helpers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type GithubUser struct {
	Username string `json:"login"`
}

type IGithubTokenValidator interface {
	ValidateGithubToken(token string) (*GithubUser, error)
}

type GithubTokenValidator struct {
}

func (validator *GithubTokenValidator) ValidateGithubToken(token string) (*GithubUser, error) {
	// see https://docs.github.com/en/rest/users/users?apiVersion=2022-11-28#get-the-authenticated-user
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("Github authentication failed.")
	}

	bytes, err := io.ReadAll(res.Body)

	var user GithubUser
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		return nil, errors.New("Unable to parse Github response.")
	}

	return &user, nil
}
