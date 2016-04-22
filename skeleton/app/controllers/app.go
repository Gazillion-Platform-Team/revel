package controllers

import "github.com/Gazillion-Platform-Team/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}
