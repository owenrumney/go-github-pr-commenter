package commenter

import (
	"fmt"
	"strings"
)

type commitFileInfo struct {
	FileName     string
	hunkInfos    []*hunkInfo
	sha          string
	likelyBinary bool
}

type hunkInfo struct {
	hunkStart int
	hunkEnd   int
}

func (hi hunkInfo) isLineInHunk(line int) bool {
	return line >= hi.hunkStart && line <= hi.hunkEnd
}

func getCommitFileInfo(ghConnector *connector) ([]*commitFileInfo, error) {

	prFiles, err := ghConnector.getFilesForPr()
	if err != nil {
		return nil, err
	}

	var (
		errs            []string
		commitFileInfos []*commitFileInfo
	)

	for _, file := range prFiles {
		info, err := getCommitInfo(file)
		if err != nil {
			errs = append(errs, err.Error())
			continue
		}
		commitFileInfos = append(commitFileInfos, info)
	}
	if len(errs) > 0 {
		return nil, fmt.Errorf("there were errors processing the PR files.\n%s", strings.Join(errs, "\n"))
	}
	return commitFileInfos, nil
}

func (cfi *commitFileInfo) getHunkInfo(line int) *hunkInfo {
	for _, hunkInfo := range cfi.hunkInfos {
		if hunkInfo.isLineInHunk(line) {
			return hunkInfo
		}
	}
	return nil
}

func (cfi *commitFileInfo) isLineInChange(line int) bool {
	return cfi.getHunkInfo(line) != nil
}

func (cfi commitFileInfo) calculatePosition(line int) *int {
	hi := cfi.getHunkInfo(line)
	position := line - hi.hunkStart
	return &position
}

func (cfi commitFileInfo) isBinary() bool {
	return cfi.likelyBinary
}

func (cfi commitFileInfo) isResolvable() bool {
	return cfi.isBinary() && len(cfi.hunkInfos) == 0
}
