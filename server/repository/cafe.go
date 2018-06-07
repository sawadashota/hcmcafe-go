package repository

import (
	"google.golang.org/appengine"
	"net/http"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"github.com/sawadashota/hcmcafe/server/domain"
)

const (
	TypeCafe = "cafe"
)

// StoreCafe cafe to database
func  StoreCafe(c *domain.Cafe, r *http.Request) error {
	ctx := appengine.NewContext(r)

	_, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, TypeCafe, nil), c)

	if err != nil {
		log.Errorf(ctx, "could not put into datastore: %v", err)
		return err
	}

	return nil
}
