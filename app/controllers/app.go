package controllers

import (
	"github.com/lanjiang/reddit-lite/app/routes"
	"github.com/revel/revel"
)

//App controller
type App struct {
	*revel.Controller
}

//Index navigates to post listing page
func (c App) Index() revel.Result {
	return c.Redirect(routes.Posts.List())
}
