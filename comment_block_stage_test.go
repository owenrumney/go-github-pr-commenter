package go_github_pr_commenter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type commentBlockTest struct {
	t *testing.T
	c *CommentBlock
}

func (t *commentBlockTest) a_new_comment_block_for_file_with_comments_(filename, comment string) {
	t.c = NewCommentBlock(filename, comment, 1, 10)
}

func (t *commentBlockTest) the_no_start_or_end_line_are_applied() {
	// do nothing
}

func (t *commentBlockTest) has_filename(expected string) *commentBlockTest {
	assert.Equal(t.t, t.c.fileName, expected)
	return t
}

func (t *commentBlockTest) has_comment(expected string) *commentBlockTest {
	assert.Equal(t.t, t.c.comment, expected)
	return t
}

func (t *commentBlockTest) has_start_line(expected int) *commentBlockTest {
	assert.Equal(t.t, t.c.startLine, expected)
	return t
}

func (t *commentBlockTest) has_end_line(expected int) *commentBlockTest {
	assert.Equal(t.t, t.c.endLine, expected)
	return t
}

func (t *commentBlockTest) and() *commentBlockTest {
	return t
}

func newCommentBlockTest(t *testing.T) (*commentBlockTest, *commentBlockTest, *commentBlockTest) {
	cbt := &commentBlockTest{
		t: t,
	}
	return cbt, cbt, cbt
}
