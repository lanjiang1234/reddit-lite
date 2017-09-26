package models

import (
	"time"

	"github.com/revel/revel"
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

//Validate post model
func (post *Post) Validate(v *revel.Validation) {
	v.Check(post.Title,
		revel.Required{},
		revel.MinSize{10},
		revel.MaxSize{100},
	)

	v.Check(post.Content,
		revel.Required{},
		revel.MinSize{20},
		revel.MaxSize{255},
	)
}
