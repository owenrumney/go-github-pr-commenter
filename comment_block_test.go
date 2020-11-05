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

func Test_comment_block_created_ok_with_a_start_line(t *testing.T) {
	given, when, then := newCommentBlockTest(t)

	given.a_new_comment_block_for_file_with_comments_("test.file", "this code is really broken")

	when.the_start_line_is_set_to_(10)

	then.has_filename("test.file").
		and().has_comment("this code is really broken").
		and().has_start_line(10).
		and().has_end_line(-1)
}

func Test_comment_block_created_ok_with_an_end_line(t *testing.T) {
	given, when, then := newCommentBlockTest(t)

	given.a_new_comment_block_for_file_with_comments_("test.file", "this code is really broken")

	when.the_end_line_is_set_to_(10)

	then.has_filename("test.file").
		and().has_comment("this code is really broken").
		and().has_start_line(-1).
		and().has_end_line(10)
}

func Test_comment_block_created_ok_with_a_start_and_an_end_line(t *testing.T) {
	given, when, then := newCommentBlockTest(t)

	given.a_new_comment_block_for_file_with_comments_("test.file", "this code is really broken")

	when.the_start_line_is_set_to_(10).
		and().the_end_line_is_set_to_(10)

	then.has_filename("test.file").
		and().has_comment("this code is really broken").
		and().has_start_line(10).
		and().has_end_line(10)
}
