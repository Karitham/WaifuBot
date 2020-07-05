package disc

import (
	"bot/database"
	"bot/query"
	"fmt"

	"github.com/andersfylling/disgord"
)

type fav struct {
	ID    interface{} `bson:"ID"`
	Name  string      `bson:"Name"`
	Image string      `bson:"Image"`
}

func favourite(data *disgord.MessageCreate, args []string) {
	if len(args) > 0 {
		resp, err := query.CharSearch(args)
		if err != nil {
			fmt.Println(err)
		}
		database.SetFavourite(database.FavouriteStruct{
			UserID: data.Message.Author.ID,
			Favourite: fav{
				ID:    resp.Character.ID,
				Name:  resp.Character.Name.Full,
				Image: resp.Character.Image.Large,
			},
		})
	}
}
