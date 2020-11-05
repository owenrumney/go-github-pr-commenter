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

type githubConnector struct {
	client   *github.Client
	owner    string
	repo     string
	prNumber int
}

type commenter struct {
	connector *githubConnector
}

type CommitFileInfo struct {
	fileName  string
	hunkStart int
	hunkEnd   int
	sha       string
}

var patchRegex *regexp.Regexp
var commitRefRegex *regexp.Regexp

func New(owner, repo, token string, prNumber int) (*commenter, error) {
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

	return &commenter{
		connector: &githubConnector{
			client:   client,
			owner:    owner,
			repo:     repo,
			prNumber: prNumber,
		},
	}, nil
}

func (gc *commenter) GetCommitFileInfo() ([]*CommitFileInfo, error) {
	prFiles, err := gc.getFilesForPr()
	if err != nil {
		return nil, err
	}
	var commitFileInfos []*CommitFileInfo
	var errs []string
	for _, file := range prFiles {
		info, err := getCommitInfo(file)
		if err != nil {
			errs = append(errs, err.Error())
			continue
		}
		commitFileInfos = append(commitFileInfos, info)
	}
	if len(errs) > 0 {
		return commitFileInfos, errors.New(fmt.Sprintf("there were errors processing the PR files.\n%s", strings.Join(errs, "\n")))
	}
	return commitFileInfos, nil
}

func createClient(token string) *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}

func (gc *commenter) getFilesForPr() ([]*github.CommitFile, error) {
	connector := gc.connector
	ctx := context.Background()

	files, _, err := connector.client.PullRequests.ListFiles(ctx, connector.owner, connector.repo, connector.prNumber, nil)
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

func (gc *commenter) getExistingComments() ([]string, error) {
	connector := gc.connector
	ctx := context.Background()

	var bodies []string
	comments, _, err := connector.client.PullRequests.ListComments(ctx, connector.owner, connector.repo, connector.prNumber, &github.PullRequestListCommentsOptions{})
	if err != nil {
		return nil, err
	}
	for _, comment := range comments {
		bodies = append(bodies, comment.GetBody())
	}
	return bodies, nil
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
		fileName:  *file.Filename,
		hunkStart: hunkStart,
		hunkEnd:   hunkEnd,
		sha:       sha,
	}, nil
}
