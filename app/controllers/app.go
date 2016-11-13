package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

type CategoryEntry struct {
    Name string
    Description string
    ImagePath string
    Url string
    Id int
}

var (
    cats = []CategoryEntry{
        {Name: "Historical",
         Description: "Tours steeped in history with knowledgable guides explaining the significance of locations easily overlooked",
         ImagePath: "historical.png",
         Url: "historical",
         Id: 1},
        {Name: "Ghost Tours",
         Description: "Get frightened out of your wits on spooky guided walks through atmospheric streets with murky pasts...",
         ImagePath: "ghost_tours.png",
         Url: "ghost_tours",
         Id: 2},
        {Name: "Famous People",
         Description: "Discover places which the great and the good called home, worked or played and get closer to your heroes",
         ImagePath: "famous_people.png",
         Url: "famous_people",
         Id: 3},
        {Name: "Architecture",
         Description: "Get introduced to the most awe-inspiring buildings and structures your location has to offer",
         ImagePath: "architecture.png",
         Url: "architecture",
         Id: 4},
        {Name: "Film and TV",
         Description: "Become star-struck as you're guided around locations familiar and famous from featuring on the big or small screen",
         ImagePath: "film_and_tv.png",
         Url: "film_and_tv",
         Id: 5},
        {Name: "Literature",
         Description: "Visit places that inspired literature, places that featured in literature and places where the great authors lived and worked",
         ImagePath: "literature.png",
         Url: "literature",
         Id: 6}}
     )

func GetCategoryEntryFromUrl(targetUrl string) CategoryEntry {
    for _, category := range cats {
        if category.Url == targetUrl {
            return category
        }
    }

    return cats[0]
}

func (c App) Index() revel.Result {
	return c.Render(cats)
}
