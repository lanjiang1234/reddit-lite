package controllers

import (
	"github.com/lanjiang/reddit-lite/app/models"
	"github.com/lanjiang/reddit-lite/app/routes"
	"github.com/lanjiang/reddit-lite/app/services"

	"github.com/revel/revel"
)

type Posts struct {
	*revel.Controller
}

//view post detail
func (c Posts) View(id int) revel.Result {
	post := services.GetPostById(id)
	return c.Render(post)
}

//list top 20 topic sorted by vote and datetime
func (c Posts) List() revel.Result {
	posts := services.GetAllPosts(20)
	return c.Render(posts)
}

//create new post
func (c Posts) New() revel.Result {
	return c.Render()
}

//save new post 
func (c Posts) SavePost(post models.Post) revel.Result {
	post.Validate(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Posts.New())
	}
	services.AddPost(&post)
	return c.Redirect(routes.Posts.List())
}

//upvote or downvote post
func (c Posts) Vote(id int, voteType int) revel.Result {
	post := services.GetPostById(id)
	services.VotePost(post.PostId, voteType)

	data := make(map[string]interface{})
	data["id"] = post.PostId
	data["upvote"] = post.UpVote
	data["downvote"] = post.DownVote

	return c.RenderJSON(data)
}