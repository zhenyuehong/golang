package persist

import (
	"golang/carwler/engine"
	"golang/carwler/persist"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) SaveItem(item engine.Item, result *string) error {
	err := persist.SaveItem(s.Client, s.Index, item)
	log.Printf("item %v saved.", item)
	if err != nil {
		*result = "ok"
	} else {
		log.Printf("error saving item %v: %v", item, err)
	}
	return err
}
