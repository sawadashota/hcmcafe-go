package persistence

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

type Entity interface {
	Delete()
	UpdateStamp()
}

type datastoreRepository struct {
	kind string
}

func put(r *http.Request, kind string, e Entity) error {
	ctx := appengine.NewContext(r)
	e.UpdateStamp()

	_, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, kind, nil), e)

	if err != nil {
		log.Errorf(ctx, "could not put into datastore: %v", err)
		return err
	}

	return nil
}

//func delete(r *http.Request, kind string, id string) error {
//	return nil
//}
