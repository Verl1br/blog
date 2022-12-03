package repository

import (
	"context"
	"fmt"

	"github.com/dhevve/blog/internal/model"
	"github.com/jmoiron/sqlx"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/sirupsen/logrus"
)

type FriendRepository struct {
	db     *sqlx.DB
	driver neo4j.DriverWithContext
}

func NewFriendRepository(db *sqlx.DB, driver neo4j.DriverWithContext) *FriendRepository {
	return &FriendRepository{
		db:     db,
		driver: driver,
	}
}

func (r *FriendRepository) GetFriends(id int) []model.User {
	var friends []model.User
	var friend model.User

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
	if err != nil {
		logrus.Errorf("ANOTHER ERROR: %s", err.Error())
	}

	query := fmt.Sprintf("SELECT first_name FROM %s WHERE id = $1", usersTable)
	for _, v := range res {
		r.db.Get(&friend, query, v)
		friends = append(friends, friend)
	}

	return friends
}

func (r *FriendRepository) CreateFriends(myId, friendId int) error {
	res := make([]int64, 0)
	readFriendById := `
			Match(f:Friends{user_id: $id})
			RETURN f.user_id;`

	createRelationshipBetweenPeopleQuery := `
			MATCH (f1:Friends { user_id: $id1 })
			MATCH (f2:Friends { user_id: $id2 })
			CREATE (f1)-[:FRIEND]->(f2)
			RETURN f1, f2`

	createFriend := `
			CREATE (f:Friends { user_id: $id })
			RETURN f.user_id
	`

	ctx := context.Background()

	session := r.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	_, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(ctx, readFriendById, map[string]any{
			"id": myId,
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
	if err != nil {
		logrus.Errorf("ANOTHER ERROR: %s", err.Error())
	}

	if len(res) == 0 {
		_, err := session.ExecuteWrite(ctx,
			func(tx neo4j.ManagedTransaction) (any, error) {
				result, err := tx.Run(ctx, createFriend, map[string]any{
					"id": myId,
				})
				if err != nil {
					return nil, err
				}

				return result.Collect(ctx)
			})
		if err != nil {
			panic(err)
		}
	}

	records, err := session.ExecuteWrite(ctx,
		func(tx neo4j.ManagedTransaction) (any, error) {
			result, err := tx.Run(ctx, createRelationshipBetweenPeopleQuery, map[string]any{
				"id1": myId,
				"id2": friendId,
			})
			if err != nil {
				return nil, err
			}

			return result.Collect(ctx)
		})
	for _, record := range records.([]*neo4j.Record) {
		firstPerson := record.Values[0].(neo4j.Node)
		fmt.Printf("First: '%d'\n", firstPerson.Props["user_id"].(int64))
		secondPerson := record.Values[1].(neo4j.Node)
		fmt.Printf("Second: '%d'\n", secondPerson.Props["user_id"].(int64))
	}
	return err
}
