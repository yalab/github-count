package github

import (
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"os"
	"time"
)

// Repository
type Repository struct {
	Owner string
	Name  string
}

// PullRequests
func (repos *Repository) PullRequests(since time.Time) []*github.PullRequest {
	fmt.Println(since)
	client := getClient()
	opt := &github.PullRequestListOptions{"closed", "", "", "created", "desc", github.ListOptions{}}
	var pulls []*github.PullRequest
	_pulls, _, err := client.PullRequests.List(oauth2.NoContext, repos.Owner, repos.Name, opt)
	pulls = append(pulls, _pulls...)
	if err != nil {
		panic("Cannot get pull requests")
	}
	return pulls
}

func getClient() *github.Client {
	fmt.Println("getClient")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_ACCESS_TOKEN")},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)
	return client
}
