package github

import (
	"github.com/apibillme/vcr"
	"time"
	"testing"
)

func TestRepository(t *testing.T) {
	vcr.Start("my_feature", nil)
	defer vcr.Stop()
	repos := Repository{"jiraffeinc", "smama"}
	since := time.Now().Add(-24 * time.Hour * 7)
	pulls := repos.PullRequests(since)
	count := make(map[string]int)
	for _, pull := range pulls {
		user := pull.User
		count[*user.Login]++
	}
	expected := 3
	if count["yalab"] != expected {
		t.Errorf("got: %v want: %v", count["yalab"], expected)
	}
}
