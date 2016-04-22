package controllers

import "github.com/Gazillion-Platform-Team/revel"

type Application struct {
	*revel.Controller
}

func (c Application) Index() revel.Result {
	return c.Render()
}
