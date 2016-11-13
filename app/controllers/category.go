package controllers

import (
    "github.com/revel/revel"
    "github.com/john-sharp/toursite/app"
    "github.com/john-sharp/toursite/app/models"
    "log"
)

type Category struct {
	*revel.Controller
}

func getTourCatalogueEntries(list *[]models.TourCatalogueEntry, categoryId int) error {
    query := "select t.title Title, u.username Operator, t.description Description from tours t inner join users u on t.operator = u.id inner join category_tour_links ctl on ctl.tour_id = t.id where ctl.category_id = ?"

    _, err := app.Dbm.Select(list, query, categoryId)

    return err
}

func (c Category) Browse(category string) revel.Result {
    categoryEntry := GetCategoryEntryFromUrl(category)
    var tourCatalogueEntries []models.TourCatalogueEntry
    err := getTourCatalogueEntries(&tourCatalogueEntries, categoryEntry.Id)

    if err != nil {
        log.Fatalln(err)
    }

    return c.Render(categoryEntry, tourCatalogueEntries)
}
