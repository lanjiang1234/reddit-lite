package controllers

import (
	"github.com/lanjiang/reddit-lite/app/routes"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

//navigate to post list page
func (c App) Index() revel.Result {
	return c.Redirect(routes.Posts.List())
}
