package tests

import (
	"github.com/lanjiang/reddit-lite/app/models"
	"github.com/lanjiang/reddit-lite/app/services"
	"github.com/revel/revel/testing"
)

type AppTest struct {
	testing.TestSuite
}

func (t *AppTest) Before() {
	println("Set up")
}

func (t *AppTest) TestIndexPage() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *AppTest) TestAddPostService() {
	post := models.Post{Title: "this is a test post", Content: "this is a test post"}
	id := postService.Add(&post)

	t.Assert(id >= 0)
	postService.Delete(&post)
}

func (t *AppTest) TestGetPostService() {
	post := models.Post{Title: "this is a test post", Content: "this is a test post"}
	id := postService.Add(&post)
	_, p := postService.GetByID(id)

	t.Assert(p != nil)
	postService.Delete(&post)
}

func (t *AppTest) TestVotePostService() {
	post := models.Post{Title: "", Content: ""}
	id := postService.Add(&post)
	postService.Vote(id, 0)
	postService.Vote(id, 1)
	_, p := postService.GetByID(id)

	t.Assert(p.UpVote == 1 && p.DownVote == 1)
	postService.Delete(&post)
}

func (t *AppTest) After() {
	println("Tear down")
}
