package controllers

import (
	"github.com/revel/revel"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	name := "Justin Beckwith"
	context := appengine.NewContext(c.Request.Request)
	log.Infof(context, "Serving the front page.")
	return c.Render(name)
}
