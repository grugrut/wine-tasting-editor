package model

import (
	"context"
	"google.golang.org/appengine/datastore"
	"time"
)

// Wine model
type Wine struct {
	Name    string
	Vintage int
	What    []string
	Who     string
	Where   string
	Type    string
	URL     string
	Created time.Time
	Updated time.Time
}

// GetWines returns wine mode list
func GetWines(c context.Context) ([]Wine, error) {
	q := datastore.NewQuery("Wine").Limit(10)
	wines := make([]Wine, 0, 10)
	if _, err := q.GetAll(c, &wines); err != nil {
		return nil, err
	}

	return wines, nil
}

// InsertWine insert wine to datastore
func InsertWine(c context.Context, wine Wine) error {
	key := datastore.NewIncompleteKey(c, "Wine", nil)
	_, err := datastore.Put(c, key, &wine)
	return err
}
