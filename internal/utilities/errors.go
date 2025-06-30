package utilities

import "fmt"

type CommentInsertionError struct {
	Err error
}

func (c *CommentInsertionError) Error() string {
	return fmt.Sprintf("CommentInsertionError-> %v", c.Err)
}

type ListFilesError struct {
	Err error
}

func (l *ListFilesError) Error() string {
	return fmt.Sprintf("ListFilesError-> %v", l.Err)
}

type IncludePathError struct {
	Err error
}

func (i *IncludePathError) Error() string {
	return fmt.Sprintf("IncludePathError-> %v", i.Err)
}
