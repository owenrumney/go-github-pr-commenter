package go_github_pr_commenter

type CommentBlock struct {
	FileName  string
	StartLine int
	EndLine   int
	Comment   string
}

func (cb *CommentBlock) CalculatePosition() *int {
	position := cb.StartLine - cb.CommitFileInfo.hunkStart
	return &position
}
