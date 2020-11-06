package go_github_pr_commenter

type CommitFileInfo struct {
	FileName  string
	hunkStart int
	hunkEnd   int
	sha       string
}

func (cfi CommitFileInfo) CommentRequired(filename string, startLine int) bool {
	return filename == cfi.FileName && startLine > cfi.hunkStart && startLine < cfi.hunkEnd
}
