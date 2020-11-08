package test

import (
	"github.com/owenrumney/go-github-pr-commenter/commenter"
	"testing"
)

type commenterTest struct {
	t         *testing.T
	token     string
	owner     string
	repo      string
	prNo      int
	commenter *commenter.Commenter
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

func (ct *commenterTest) a_new_commenter_is_created() {
	c, err := commenter.NewCommenter(ct.token, ct.owner, ct.repo, ct.prNo)
	if err != nil {
		panic(err)
	}
	ct.commenter = c
}

func (ct *commenterTest) there_is_no_errors() {

}
