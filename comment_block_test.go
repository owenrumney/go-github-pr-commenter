package go_github_pr_commenter

import (
	"testing"
)

func Test_comment_block_created_ok(t *testing.T) {
	given, when, then := newCommentBlockTest(t)

	given.a_new_comment_block_for_file_with_comments_("test.file", "this code is really broken")

	when.the_no_start_or_end_line_are_applied()

	then.has_filename("test.file").
		and().has_comment("this code is really broken").
		and().has_start_line(-1).
		and().has_end_line(-1)
}
