package github

import "testing"
import "fmt"

func TestRepository(t *testing.T) {
    repos := Repository{"jiraffeinc", "smama"}
    pulls := repos.PullRequests()
    for _, pull := range pulls {
        user := pull.User
        fmt.Println(*user.Login)
    }
}

