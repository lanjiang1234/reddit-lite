package controllers

import (
	"github.com/lanjiang/reddit-lite/app/models"
	"github.com/lanjiang/reddit-lite/app/routes"
	"github.com/lanjiang/reddit-lite/app/services"

	"github.com/revel/revel"
)

//Posts Controller
type Posts struct {
	*revel.Controller
}

//View post detail
func (c Posts) View(id int) revel.Result {
	_, post := postService.GetByID(id)
	if post == nil {
		return c.NotFound("Post %d does not exist", id)
	}
	return c.Render(post)
}

//List top 20 topic sorted by vote and datetime
func (c Posts) List() revel.Result {
	posts := postService.GetAll(20)
	return c.Render(posts)
}

//New post page
func (c Posts) New() revel.Result {
	return c.Render()
}

//Save new post
func (c Posts) Save(post models.Post) revel.Result {
	post.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Posts.New())
	}
	postService.Add(&post)
	return c.Redirect(routes.Posts.List())
}

//Vote a post (upvote or downvote)
func (c Posts) Vote(id int, voteType int) revel.Result {
	_, post := postService.GetByID(id)
	data := make(map[string]interface{})
	if post == nil {
		data["id"] = post.PostId
		data["error"] = "Post does not exists"
	} else {
		postService.Vote(post.PostId, voteType)
		data["id"] = post.PostId
		data["upvote"] = post.UpVote
		data["downvote"] = post.DownVote
	}
	return c.RenderJSON(data)
}
