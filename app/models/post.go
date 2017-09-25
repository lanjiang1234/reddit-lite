package models

import (
	"github.com/revel/revel"
	"time"
)

type Post struct {
	PostId  int
	Title   string
	Content string
	CreatedOn time.Time
	UpVote 	int
	DownVote int
}

//validate post model
func (post *Post) Validate(v *revel.Validation) {
	v.Check(post.Title,
		revel.Required{},
		revel.MinSize{1},
		revel.MaxSize{100},
	)

	v.Check(post.Content,
		revel.Required{},
		revel.MinSize{1},
		revel.MaxSize{200},
	)
}