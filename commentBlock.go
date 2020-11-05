package go_github_pr_commenter

type CommentBlock struct {
	fileName  string
	startLine int
	endLine   int
	comment   string
}

func NewCommentBlock(filename, comment string) *CommentBlock {
	return &CommentBlock{
		fileName:  filename,
		comment:   comment,
		startLine: -1,
		endLine:   -1,
	}
}

func (c *CommentBlock) WithStartLine(lineNo int) *CommentBlock {
	c.startLine = lineNo
	return c
}

func (c *CommentBlock) WithEndLine(lineNo int) *CommentBlock {
	c.endLine = lineNo
	return c
}
