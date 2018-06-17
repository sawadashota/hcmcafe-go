package persistence

import (
	"net/http"

	"fmt"

	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

type Entity interface {
	GetId() string
	SetId(id string)
	Delete()
	UpdateStamp()
	IsDeleted() bool
}

type datastoreRepository struct {
	kind string
}

func find(r *http.Request, kind string, id string, e Entity) error {
	ctx := appengine.NewContext(r)

	key := datastore.NewKey(ctx, kind, id, 0, nil)
	err := datastore.Get(ctx, key, e)

	if err != nil {
		log.Errorf(ctx, "could not get: %v", err)
		return err
	}

	if e.IsDeleted() {
		return fmt.Errorf("id: %s is alredy deleted", id)
	}

	e.SetId(key.StringID())

	return nil
}

func first(r *http.Request, kind string, key string, value interface{}, e Entity) error {
	ctx := appengine.NewContext(r)

	emptyTime := time.Time{}
	it := datastore.NewQuery(kind).
		Filter(equal(key), value).
		Filter("deleted_at =", emptyTime).
		Run(ctx)

	k, err := it.Next(e)

	if err != nil {
		log.Errorf(ctx, "could not get: %v", err)
		return err
	}

	e.SetId(k.StringID())

	return nil
}

func put(r *http.Request, kind string, e Entity) error {
	ctx := appengine.NewContext(r)
	e.UpdateStamp()

	key := datastore.NewKey(ctx, kind, e.GetId(), 0, nil)
	_, err := datastore.Put(ctx, key, e)

	if err != nil {
		log.Errorf(ctx, "could not put into datastore: %v", err)
		return err
	}

	return nil
}

func destroy(r *http.Request, kind string, id string, e Entity) error {

	ctx := appengine.NewContext(r)

	key := datastore.NewKey(ctx, kind, id, 0, nil)
	err := datastore.Get(ctx, key, e)

	if err != nil {
		log.Errorf(ctx, "could not get: %v", err)
		return err
	}

	e.Delete()

	_, err = datastore.Put(ctx, key, e)

	if err != nil {
		log.Errorf(ctx, "could not put into datastore: %v", err)
		return err
	}

	return nil
}

func equal(key string) string {
	return withOperator(key, "=")
}

func withOperator(key, operator string) string {
	return fmt.Sprintf("%s %s", key, operator)
}
