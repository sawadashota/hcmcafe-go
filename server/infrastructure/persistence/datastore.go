package persistence

import (
	"net/http"

	"fmt"

	"time"

	"context"

	"github.com/mjibson/goon"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

type Entity interface {
	Delete()
	UpdateStamp()
	IsDeleted() bool
}

type Projector func(entities interface{}) (int, error)

type datastoreRepository struct {
	kind string
}

func getAll(r *http.Request, kind string, limit, page int) Projector {
	g := goon.NewGoon(r)

	emptyTime := time.Time{}
	q := datastore.NewQuery(kind).
		Filter("deleted_at =", emptyTime).
		Limit(limit).
		Offset(pageToOffset(page, limit))

	return func(entities interface{}) (int, error) {
		_, err := g.GetAll(q, entities)

		if err != nil {
			errorLog(g.Context, "%v", err)
			return 0, err
		}

		count, err := q.Count(g.Context)

		if err != nil {
			errorLog(g.Context, "%v", err)
			return 0, err
		}
		return count, nil
	}
}

func get(r *http.Request, kind, key string, value interface{}, es []Entity, limit, page int) error {
	ctx := appengine.NewContext(r)

	emptyTime := time.Time{}
	_, err := datastore.NewQuery(kind).
		Filter(equal(key), value).
		Filter("deleted_at =", emptyTime).
		Limit(limit).
		Offset((page*limit)+1).
		GetAll(ctx, es)

	if err != nil {
		errorLog(ctx, "could not get: %v", err)
		return err
	}

	return nil

}

func find(r *http.Request, id string, e Entity) error {

	g := goon.NewGoon(r)
	err := g.Get(e)

	if err != nil {
		errorLog(g.Context, "could not find: %v", err)
		return err
	}

	if e.IsDeleted() {
		return fmt.Errorf("id: %s is alredy deleted", id)
	}

	return nil
}

func first(r *http.Request, key string, value interface{}, e Entity) error {
	g := goon.NewGoon(r)

	emptyTime := time.Time{}
	q := datastore.NewQuery(g.Kind(e)).
		Filter(equal(key), value).
		Filter("deleted_at =", emptyTime)

	it := g.Run(q)

	_, err := it.Next(e)

	if err != nil {
		errorLog(g.Context, "could not get: %v", err)
		return err
	}

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
	g := goon.NewGoon(r)

	emptyTime := time.Time{}
	q := datastore.NewQuery(kind).
		Filter(equal(key), value).
		Filter("deleted_at =", emptyTime).
		KeysOnly()

	keys, err := g.GetAll(q, nil)

	if err != nil {
		return 0, err
	}

	if includeKey(keys, excludeId) {
		return len(keys) - 1, nil
	}

	return len(keys), nil
}

func put(r *http.Request, e Entity) error {

	e.UpdateStamp()

	g := goon.NewGoon(r)
	_, err := g.Put(e)

	if err != nil {
		errorLog(g.Context, "could not put into datastore: %v", err)
		return err
	}

	return nil
}

func destroy(r *http.Request, e Entity) error {
	g := goon.NewGoon(r)

	return g.RunInTransaction(func(tg *goon.Goon) error {
		err := g.Get(e)

		if err != nil {
			errorLog(tg.Context, "%v", err)
			return err
		}

		e.Delete()

		_, err = g.Put(e)

		return err
	}, nil)
}

// physicalDestroy is deprecated
func physicalDestroy(r *http.Request, kind string, id string) error {
	ctx := appengine.NewContext(r)

	key := datastore.NewKey(ctx, kind, id, 0, nil)
	err := datastore.Delete(ctx, key)

	if err != nil {
		errorLog(ctx, "could not delete: %v", err)
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

func errorLog(ctx context.Context, format string, args ...interface{}) {
	log.Errorf(ctx, format, args)
}

func pageToOffset(page, limit int) int {
	return (page - 1) * limit
}
