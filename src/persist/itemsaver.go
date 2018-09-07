package persist

import (
"log"
"context"
"engine"
_"errors"
"github.com/olivere/elastic"
)

func ItemSaver(index string, index_type string) (chan engine.Item, error){
	client, err := elastic.NewClient(elastic.SetURL("http://192.168.187.185:9200"))
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		for {
			item := <- out
			log.Printf("Item Saver: got item %v",item)
			err := save(client, index, index_type, item)
			if err != nil {
				log.Printf("item saver error")
			}
		}
	}()
	return out, nil
}
	
func save(client *elastic.Client, index string, index_type string,  item engine.Item) error {
	/*
	if item.Type == "" {
		return errors.New("must supply Type")
	}
	*/

	indexService := client.Index().Index(index).
		Type(index_type).Id(item.Id).BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}	

	_, err := indexService.Do(context.Background())

	if err != nil {
		return err
	}
	return nil
}
