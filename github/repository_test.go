package github

import "testing"
import "fmt"

func TestRepository(t *testing.T) {
    repos := Repository{"jiraffeinc", "smama"}
    pulls := repos.PullRequests()
    count := make(map[string]int)
    for _, pull := range pulls {
        user := pull.User
        count[*user.Login] ++
    }
    fmt.Println(count)
}

