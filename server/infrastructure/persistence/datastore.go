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
		log.Errorf(ctx, "could not find: %v", err)
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

func exist(r *http.Request, kind, excludeId, key string, value interface{}) (bool, error) {
	c, err := count(r, kind, excludeId, key, value)

	if err != nil {
		return false, err
	}

	return c > 0, nil
}

func count(r *http.Request, kind, excludeId, key string, value interface{}) (int, error) {
	keys := make([]*datastore.Key, 0)
	ctx := appengine.NewContext(r)

	emptyTime := time.Time{}
	it := datastore.NewQuery(kind).
		Filter(equal(key), value).
		Filter("deleted_at =", emptyTime).
		KeysOnly().
		Run(ctx)

	var k *datastore.Key
	var err error
	for err == nil {
		k, err = it.Next(nil)

		if err != nil {
			break
		}

		keys = append(keys, k)
	}

	if includeKey(keys, key) {
		return len(keys) - 1, nil
	}

	return len(keys), nil
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

func physicalDestroy(r *http.Request, kind string, id string) error {
	ctx := appengine.NewContext(r)

	key := datastore.NewKey(ctx, kind, id, 0, nil)
	err := datastore.Delete(ctx, key)

	if err != nil {
		log.Errorf(ctx, "could not delete: %v", err)
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

func includeKey(keys []*datastore.Key, key string) bool {
	for _, k := range keys {
		if k.StringID() == key {
			return true
		}
	}

	return false
}
