package main

import "fmt"
import "github.com/yalab/github-count/github"

func main() {
	repos := github.Repository{"jiraffeinc", "smama"}
	pulls := repos.PullRequests()
	count := make(map[string]int)
	for _, pull := range pulls {
		user := pull.User
		count[*user.Login]++
	}
	fmt.Println(count)
}
