package test

import (
	"github.com/owenrumney/go-github-pr-commenter/commenter"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

type commenterTest struct {
	t         *testing.T
	token     string
	owner     string
	repo      string
	prNo      int
	commenter *commenter.Commenter
	err       error
}

func newCommenterTest(t *testing.T) (*commenterTest, *commenterTest, *commenterTest) {
	ct := &commenterTest{
		t: t,
	}

	return ct, ct, ct
}

func (ct *commenterTest) the_pull_request_(prNo int) *commenterTest {
	ct.prNo = prNo
	return ct
}

func (ct *commenterTest) for_owner_(owner string) *commenterTest {
	ct.owner = owner
	return ct
}

func (ct *commenterTest) in_repo_(repo string) *commenterTest {
	ct.repo = repo
	return ct
}

func (ct *commenterTest) using_token_(token string) *commenterTest {
	ct.token = token
	return ct
}

func (ct *commenterTest) using_token_from_environment() *commenterTest {
	ct.token = os.Getenv("GITHUB_TOKEN")
	return ct
}

func (ct *commenterTest) a_new_commenter_is_created() *commenterTest {
	c, err := commenter.NewCommenter(ct.token, ct.owner, ct.repo, ct.prNo)
	if err != nil {
		ct.err = err
	}
	ct.commenter = c
	return ct
}

func (ct *commenterTest) and() *commenterTest {
	return ct
}

func (ct *commenterTest) there_is_no_errors() *commenterTest {
	assert.True(ct.t, ct.err == nil)
	return ct
}

func (ct *commenterTest) there_is_an_error() *commenterTest {
	assert.True(ct.t, ct.err != nil)
	return ct
}

func (ct *commenterTest) a_single_line_comment_is_created() {
	err := ct.commenter.WriteLineComment("commitFileInfo.go", "This is awesome", 7)
	ct.err = err
}

func (ct *commenterTest) a_multi_line_comment_is_created() {
	err := ct.commenter.WriteMultiLineComment("connector.go", "Is this the best way", 9, 14)
	ct.err = err
}

func (ct *commenterTest) the_single_line_comment_has_been_written() {
	assert.True(ct.t, checkForComment("commitFileInfo.go", "This is awesome", 7))
}

func (ct *commenterTest) the_multi_line_comment_has_been_written() {
	assert.True(ct.t, checkForComment("connector.go", "Is this the best way", 14))
}

func (ct *commenterTest) the_error_is_pr_not_exist() {
	existError := ct.err.(commenter.PrDoesNotExistError)
	assert.NotNil(ct.t, existError)
	assert.Equal(ct.t, ct.err.Error(), "PR number [-1] not found for owenrumney/go-github-pr-commenter")
}

func (ct *commenterTest) a_single_line_comment_is_created_with_duplicate() {
	err := ct.commenter.WriteLineComment("commitFileInfo.go", "This is going to be duped", 7)
	ct.err = err
}

func (ct *commenterTest) the_single_line_comment_with_duplicate_has_been_written() {
	assert.True(ct.t, checkForComment("commitFileInfo.go", "This is going to be duped", 7))
}

func (ct *commenterTest) the_error_is_comment_already_written() {
	existError := ct.err.(commenter.CommentAlreadyWrittenError)
	assert.NotNil(ct.t, existError)
	assert.Equal(ct.t, ct.err.Error(), "The file [commitFileInfo.go] already has the comment written [This is going to be duped]")
}

func (ct *commenterTest) the_error_is_comment_is_invalid() {
	existError := ct.err.(commenter.CommentNotValidError)
	assert.NotNil(ct.t, existError)
	assert.True(ct.t, strings.HasPrefix(ct.err.Error(), "There is nothing to comment on at line [100] in file"))
}

func (ct *commenterTest) a_single_line_comment_is_created_that_isnt_valid() {
	err := ct.commenter.WriteLineComment("commitFileInfo.go", "This is going to be duped", 100)
	ct.err = err
}

func (ct *commenterTest) a_single_line_comment_is_created_that_isnt_valid_file() {
	err := ct.commenter.WriteLineComment("madeup_file.go", "This is going to be duped", 100)
	ct.err = err
}
