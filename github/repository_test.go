package github

import (
    "testing"
    "github.com/apibillme/vcr"
)

func TestRepository(t *testing.T) {
    vcr.Start("my_feature", nil)
    defer vcr.Stop()
    repos := Repository{"jiraffeinc", "smama"}
    pulls := repos.PullRequests()
    count := make(map[string]int)
    for _, pull := range pulls {
        user := pull.User
        count[*user.Login] ++
    }
    expected := 3
    if (count["yalab"] != expected) {
        t.Errorf("got: %v want: %v", count["yalab"], expected)
    }
}

