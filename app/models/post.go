package models

import (
	"time"
)

//Post stores the post data
type Post struct {
	PostId    int
	Title     string
	Content   string
	CreatedOn time.Time
	UpVote    int
	DownVote  int
}