package main

import (
	commenter "github.com/owenrumney/go-github-pr-commenter"
	"os"
)

func main() {

	token := os.Getenv("GITHUB_TOKEN")

	c, err := commenter.NewCommenter(token, "tfsec", "tfsec-example-project", 8)
	if err != nil {
		panic(err)
	}

	if c.CheckCommentRelevant(".travis.yml", 5) {
		c.WriteComment(&commenter.CommentBlock{
			CommitFileInfo: ,
			StartLine:      0,
			EndLine:        0,
			Comment:        "",
		})
	}
}
