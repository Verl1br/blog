package repository

import (
	"context"
	"fmt"

	"github.com/dhevve/blog/internal/model"
	"github.com/jmoiron/sqlx"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/sirupsen/logrus"
)

type NewsFeedRepository struct {
	db     *sqlx.DB
	driver neo4j.DriverWithContext
}

func NewNewsFeedRepository(db *sqlx.DB, driver neo4j.DriverWithContext) *NewsFeedRepository {
	return &NewsFeedRepository{
		db:     db,
		driver: driver,
	}
}

func (r *NewsFeedRepository) GetNews(id int) ([]model.Post, error) {
	var posts []model.Post
	var friendsPosts []model.Post

	res := make([]int64, 0)
	ctx := context.Background()

	readFriendById := `
			Match(:Friends{user_id: $id})<-[:FRIEND]-(m)
			RETURN m.user_id;`

	session := r.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	_, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx, readFriendById, map[string]any{
			"id": id,
		})
		if err != nil {
			logrus.Errorf("TX ERROR: %s", err.Error())
			return nil, err
		}
		for result.Next(ctx) {
			res = append(res, result.Record().Values[0].(int64))
		}

		return nil, result.Err()
	})

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1 ORDER BY published_at DESC LIMIT 3", postTable)
	for _, v := range res {
		if err := r.db.Select(&friendsPosts, query, v); err != nil {
			logrus.Fatal(err.Error())
		}
		posts = append(posts, friendsPosts...)
	}

	return posts, err
}
