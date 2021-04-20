package test

import "testing"

func Test_not_using_token_causes_error(t *testing.T) {
	given, when, then := newCommenterTest(t)

	given.thePullRequest(1).
		forOwner("owenrumney").
		inRepo("go-github-pr-commenter")

	when.aNewCommenterIsCreated()

	then.thereIsAnError()
}

func Test_can_connect_to_pull_request_with_token(t *testing.T) {
	given, when, then := newCommenterTest(t)

	given.thePullRequest(1).
		forOwner("owenrumney").
		inRepo("go-github-pr-commenter").
		usingTokenFromEnvironment()

	when.aNewCommenterIsCreated()

	then.thereIsNoErrors()
}

func Test_can_connect_to_pull_request_with_token_but_it_doesn_exist(t *testing.T) {

	given, when, then := newCommenterTest(t)

	given.thePullRequest(-1).
		forOwner("owenrumney").
		inRepo("go-github-pr-commenter").
		usingTokenFromEnvironment()

	when.aNewCommenterIsCreated()

	then.thereIsAnError().
		and().theErrorIsPrNotExist()
}

func Test_can_add_a_single_line_comment(t *testing.T) {
	given, when, then := newCommenterTest(t)

	given.thePullRequest(1).
		forOwner("owenrumney").
		inRepo("go-github-pr-commenter").
		usingTokenFromEnvironment()

	when.aNewCommenterIsCreated().
		and().aSingleLineCommentIsCreated()

	then.thereIsNoErrors().
		and().theSingleLineCommentHasBeenWritten()
}

func Test_add_a_single_line_comment_that_is_on_a_line_that_isnt_in_pr(t *testing.T) {
	given, when, then := newCommenterTest(t)

	given.thePullRequest(5).
		forOwner("owenrumney").
		inRepo("go-github-pr-commenter").
		usingTokenFromEnvironment()

	when.aNewCommenterIsCreated().
		and().aSingleLineCommentIsCreatedThatIsntValid()

	then.thereIsAnError().
		and().theErrorIsCommentIsInvalid()
}

func Test_add_a_single_line_comment_on_a_file_that_that_isnt_in_pr(t *testing.T) {
	given, when, then := newCommenterTest(t)

	given.thePullRequest(1).
		forOwner("owenrumney").
		inRepo("go-github-pr-commenter").
		usingTokenFromEnvironment()

	when.aNewCommenterIsCreated().
		and().aSingleLineCommentIsCreatedThatIsntValidFile()
	then.thereIsAnError().
		and().theErrorIsCommentIsInvalid()
}

func Test_add_a_general_comment(t *testing.T) {
	given, when, then := newCommenterTest(t)

	given.thePullRequest(1).
		forOwner("owenrumney").
		inRepo("go-github-pr-commenter").
		usingTokenFromEnvironment()

	when.aNewCommenterIsCreated().
		and().aNewGeneralCommentIsCreated("test comment")

	then.thereIsNoErrors()
}

func Test_can_add_a_single_line_comment_second_one_has_an_error(t *testing.T) {
	given, when, then := newCommenterTest(t)

	given.thePullRequest(1).
		forOwner("owenrumney").
		inRepo("go-github-pr-commenter").
		usingTokenFromEnvironment()

	when.aNewCommenterIsCreated().
		and().aSingleLineCommentIsCreatedWithDuplicate()

	then.thereIsNoErrors().
		and().theSingleLineCommentWithDuplicateHasBeenWritten()

	then.aNewCommenterIsCreated().
		and().aSingleLineCommentIsCreatedWithDuplicate()

	then.thereIsAnError().
		and().theErrorIsCommentAlreadyWritten()

}

func Test_can_add_a_multi_line_comment(t *testing.T) {
	given, when, then := newCommenterTest(t)

	given.thePullRequest(1).
		forOwner("owenrumney").
		inRepo("go-github-pr-commenter").
		usingTokenFromEnvironment()

	when.aNewCommenterIsCreated().
		and().aMultiLineCommentIsCreated()

	then.thereIsNoErrors().and().
		theMultiLineCommentHasBeenWritten()
}
