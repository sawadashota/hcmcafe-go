package persistence

import (
	"net/http"

	"github.com/sawadashota/hcmcafe/server/domain/entity"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

const (
	TypeCafe = "cafe"
)

// StoreCafe cafe to database
func StoreCafe(c *entity.Cafe, r *http.Request) error {
	ctx := appengine.NewContext(r)

	_, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, TypeCafe, nil), c)

	if err != nil {
		log.Errorf(ctx, "could not put into datastore: %v", err)
		return err
	}

	return nil
}
