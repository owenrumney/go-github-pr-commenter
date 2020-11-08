package test

import (
	"context"
	"fmt"
	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
	"os"
	"testing"
)

var client *github.Client

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	// clear down all the comments, disgusting
	token := os.Getenv("GITHUB_TOKEN")
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)

	client = github.NewClient(tc)

	comments, _, err := client.PullRequests.ListComments(ctx, "owenrumney", "go-github-pr-commenter", 1, nil)
	if err != nil {
		panic(err)
	}
	for _, comment := range comments {
		_, err := client.PullRequests.DeleteComment(ctx, "owenrumney", "go-github-pr-commenter", *comment.ID)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Deleted existing comment on test PR: %d\n", comment.ID)
	}

	fmt.Printf("\033[1;36m%s\033[0m", "> Setup completed\n")
}

func teardown() {
	// Do something here.

	fmt.Printf("\033[1;36m%s\033[0m", "> Teardown completed")
	fmt.Printf("\n")
}

func checkForComment(file, commentString string, line int) bool {
	ctx := context.Background()

	comments, _, err := client.PullRequests.ListComments(ctx, "owenrumney", "go-github-pr-commenter", 1, nil)
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		if found := func(c *github.PullRequestComment) bool {
			return *c.Path == file && *c.Body == commentString && *c.Line == line

		}(comment); found {
			return true
		}
	}
	return false
}
