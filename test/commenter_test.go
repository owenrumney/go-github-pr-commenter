package test

import "testing"

func Test_comment_is_created(t *testing.T) {

	given, when, then := newCommenterTest(t)

	given.the_pull_request_(1).
		for_owner_("owenrumney").
		in_repo_("go-github-pr-commenter").
		using_token_("GITHUBTOKEN")
	when.a_new_commenter_is_created()
	then.there_is_no_errors()
}
