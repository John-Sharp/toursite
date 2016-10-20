package controllers

import "github.com/revel/revel"

type Category struct {
	*revel.Controller
}

func (c Category) Browse(category string) revel.Result {
    categoryEntry := GetCategoryEntryFromUrl(category)
    return c.Render(categoryEntry)
}
