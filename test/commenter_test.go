package test

import "testing"

func Test_not_using_token_causes_error(t *testing.T) {

	given, when, then := newCommenterTest(t)

	given.the_pull_request_(1).
		for_owner_("owenrumney").
		in_repo_("go-github-pr-commenter")

	when.a_new_commenter_is_created()

	then.there_is_an_error()
}

func Test_can_connect_to_pull_request_with_token(t *testing.T) {

	given, when, then := newCommenterTest(t)

	given.the_pull_request_(1).
		for_owner_("owenrumney").
		in_repo_("go-github-pr-commenter").
		using_token_from_environment()

	when.a_new_commenter_is_created()

	then.there_is_no_errors()
}

func Test_can_connect_to_pull_request_with_token_but_it_doesn_exist(t *testing.T) {

	given, when, then := newCommenterTest(t)

	given.the_pull_request_(-1).
		for_owner_("owenrumney").
		in_repo_("go-github-pr-commenter").
		using_token_from_environment()

	when.a_new_commenter_is_created()

	then.there_is_an_error().
		and().the_error_is_pr_not_exist()
}

func Test_can_add_a_single_line_comment(t *testing.T) {
	given, when, then := newCommenterTest(t)

	given.the_pull_request_(1).
		for_owner_("owenrumney").
		in_repo_("go-github-pr-commenter").
		using_token_from_environment()

	when.a_new_commenter_is_created().
		and().a_single_line_comment_is_created()
	then.there_is_no_errors().
		and().the_single_line_comment_has_been_written()
}

func Test_add_a_single_line_comment_that_is_on_a_line_that_isnt_in_pr(t *testing.T) {
	given, when, then := newCommenterTest(t)

	given.the_pull_request_(1).
		for_owner_("owenrumney").
		in_repo_("go-github-pr-commenter").
		using_token_from_environment()

	when.a_new_commenter_is_created().
		and().a_single_line_comment_is_created_that_isnt_valid()
	then.there_is_an_error().
		and().the_error_is_comment_is_invalid()
}

func Test_add_a_single_line_comment_on_a_file_that_that_isnt_in_pr(t *testing.T) {
	given, when, then := newCommenterTest(t)

	given.the_pull_request_(1).
		for_owner_("owenrumney").
		in_repo_("go-github-pr-commenter").
		using_token_from_environment()

	when.a_new_commenter_is_created().
		and().a_single_line_comment_is_created_that_isnt_valid_file()
	then.there_is_an_error().
		and().the_error_is_comment_is_invalid()
}

func Test_can_add_a_single_line_comment_second_one_has_an_error(t *testing.T) {
	given, when, then := newCommenterTest(t)

	given.the_pull_request_(1).
		for_owner_("owenrumney").
		in_repo_("go-github-pr-commenter").
		using_token_from_environment()

	when.a_new_commenter_is_created().
		and().a_single_line_comment_is_created_with_duplicate()
	then.there_is_no_errors().
		and().the_single_line_comment_with_duplicate_has_been_written()

	when.a_new_commenter_is_created().
		and().a_single_line_comment_is_created_with_duplicate()
	then.there_is_an_error().
		and().the_error_is_comment_already_written()

}

func Test_can_add_a_multi_line_comment(t *testing.T) {
	given, when, then := newCommenterTest(t)

	given.the_pull_request_(1).
		for_owner_("owenrumney").
		in_repo_("go-github-pr-commenter").
		using_token_from_environment()

	when.a_new_commenter_is_created().
		and().a_multi_line_comment_is_created()
	then.there_is_no_errors().and().
		the_multi_line_comment_has_been_written()
}
