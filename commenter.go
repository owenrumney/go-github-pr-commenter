package go_github_pr_commenter

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
	"regexp"
	"strconv"
	"strings"
)

type commenter struct {
	connector        *github.Client
	owner            string
	repo             string
	prNumber         int
	existingComments []*existingComment
	files            []*CommitFileInfo
}

type existingComment struct {
	filename *string
	comment  *string
}

var patchRegex *regexp.Regexp
var commitRefRegex *regexp.Regexp

// NewCommenter creates a commenter for updating PR with comments
func NewCommenter(token, owner, repo string, prNumber int) (*commenter, error) {
	regex, err := regexp.Compile("^@@.*\\+(\\d+),(\\d+).+?@@")
	if err != nil {
		return nil, err
	}
	patchRegex = regex

	regex, err = regexp.Compile(".+ref=(.+)")
	if err != nil {
		return nil, err
	}
	commitRefRegex = regex

	if len(token) == 0 {
		return nil, errors.New("the INPUT_GITHUB_TOKEN has not been set")
	}

	client := createClient(token)

	c := &commenter{
		connector: client,
		owner:     owner,
		repo:      repo,
		prNumber:  prNumber,
	}

	err = c.loadPr()
	if err != nil {
		return nil, err
	}
	return c, nil
}

// GetCommitFileInfo get file info for files in the commit
func (gc *commenter) getCommitFileInfo() error {
	prFiles, err := gc.getFilesForPr()
	if err != nil {
		return err
	}
	var errs []string
	for _, file := range prFiles {
		info, err := getCommitInfo(file)
		if err != nil {
			errs = append(errs, err.Error())
			continue
		}
		gc.files = append(gc.files, info)
	}
	if len(errs) > 0 {
		return errors.New(fmt.Sprintf("there were errors processing the PR files.\n%s", strings.Join(errs, "\n")))
	}
	return nil
}

func (gc *commenter) WriteComment(block *CommentBlock) error {
	connector := gc.connector
	ctx := context.Background()

	var _, _, err = connector.PullRequests.CreateComment(ctx, gc.owner, gc.repo, gc.prNumber, buildComment(block))
	if err != nil {
		return err
	}

	return nil
}

func buildComment(block *CommentBlock) *github.PullRequestComment {
	comment := &github.PullRequestComment{
		Line:     &block.StartLine,
		Path:     &block.CommitFileInfo.FileName,
		CommitID: &block.CommitFileInfo.sha,
		Body:     &block.Comment,
		Position: block.CalculatePosition(),
	}
	if block.StartLine != block.EndLine {
		comment.StartLine = &block.StartLine
		comment.Line = &block.EndLine
	}
	return comment
}

func (gc *commenter) getFilesForPr() ([]*github.CommitFile, error) {
	connector := gc.connector

	files, _, err := connector.PullRequests.ListFiles(context.Background(), gc.owner, gc.repo, gc.prNumber, nil)
	if err != nil {
		return nil, err
	}
	var commitFiles []*github.CommitFile
	for _, file := range files {
		if *file.Status != "deleted" {
			commitFiles = append(commitFiles, file)
		}
	}
	return commitFiles, nil
}

func (gc *commenter) getExistingComments() error {
	connector := gc.connector
	ctx := context.Background()

	comments, _, err := connector.PullRequests.ListComments(ctx, gc.owner, gc.repo, gc.prNumber, &github.PullRequestListCommentsOptions{})
	if err != nil {
		return err
	}
	for _, comment := range comments {
		gc.existingComments = append(gc.existingComments, &existingComment{
			filename: comment.Path,
			comment:  comment.Body,
		})
	}
	return nil
}

func (gc *commenter) loadPr() error {
	err := gc.getCommitFileInfo()
	if err != nil {
		return err
	}

	err = gc.getExistingComments()
	if err != nil {
		return err
	}

	return nil
}

func createClient(token string) *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}

func getCommitInfo(file *github.CommitFile) (*CommitFileInfo, error) {
	groups := patchRegex.FindAllStringSubmatch(file.GetPatch(), -1)
	if len(groups) < 1 {
		return nil, errors.New("the patch details could not be resolved")
	}
	hunkStart, _ := strconv.Atoi(groups[0][1])
	hunkEnd, _ := strconv.Atoi(groups[0][2])

	shaGroups := commitRefRegex.FindAllStringSubmatch(file.GetContentsURL(), -1)
	if len(shaGroups) < 1 {
		return nil, errors.New("the sha details could not be resolved")
	}
	sha := shaGroups[0][1]

	return &CommitFileInfo{
		FileName:  *file.Filename,
		hunkStart: hunkStart,
		hunkEnd:   hunkStart + (hunkEnd - 1),
		sha:       sha,
	}, nil
}

func (gc *commenter) CheckCommentRelevant(filename string, line int) bool {
	for _, file := range gc.files {
		if file.FileName == filename {
			if line > file.hunkStart && line < file.hunkEnd {
				return true
			}
		}
	}
	return false
}
