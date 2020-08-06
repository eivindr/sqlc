// Code generated by sqlc. DO NOT EDIT.

package booktest

import (
	"time"
)

type Author struct {
	AuthorID int32
	Name     string
}

type Book struct {
	BookID    int32
	AuthorID  int32
	Isbn      string
	BookType  string
	Title     string
	Yr        int32
	Available time.Time
	Tags      string
}
