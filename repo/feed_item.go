package repo

import (
	"context"
	"graph/ent"
)

type FeedItem struct {
	Db *ent.Client
}

func NewFeedItemRepo(client *ent.Client) *FeedItem {
	return &FeedItem{
		Db: client,
	}
}

func (f *FeedItem) Create(owner *ent.User, content string) (err error) {
	_, err = f.Db.FeedItem.Create().
		SetUserOwner(owner).
		SetContents(content).
		Save(context.Background())
	return err
}