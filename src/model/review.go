package model

import (
	"context"
	"google.golang.org/appengine/datastore"
	"time"
)

// Review is the review of wine
type Review struct {
	Score   int
	Color   string
	Aroma   string
	Taste   string
	Reviewd time.Time
	Created time.Time
	Updated time.Time
	Account string
	WineID  int
}

func GetReviews(c context.Context) ([]Review, error) {
	var reviews []Review
	q := datastore.NewQuery("Review")

	if _, err := q.GetAll(c, &reviews); err != nil {
		return nil, err
	}

	return reviews, nil
}

func InsertReview(c context.Context, r Review) error {
	key := datastore.NewIncompleteKey(c, "Review", nil)
	_, err := datastore.Put(c, key, &r)
	return err
}
