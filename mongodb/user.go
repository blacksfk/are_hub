package mongodb

import (
	"context"

	"github.com/blacksfk/are_hub"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Implements are_hub.UserRepo.
type User struct {
	collection
}

// Create a users collection in the database.
func NewUserCollection(client *mongo.Client, db string) User {
	return User{collection{client, db, "users"}}
}

// Get all users without their respective password and key.
func (u User) All(ctx context.Context) ([]are_hub.User, error) {
	var users []are_hub.User

	// hide the password and key
	projection := bson.M{"password": 0, "key": 0}
	cursor, e := u.get().Find(
		ctx, bson.M{}, options.Find().SetProjection(projection))

	if e != nil {
		return nil, e
	}

	return users, cursor.All(ctx, &users)
}

func (u User) FindID(ctx context.Context, id string) (*are_hub.User, error) {
	user := &are_hub.User{}

	return user, u.findID(ctx, id, user)
}

func (u User) DeleteID(ctx context.Context, id string) (*are_hub.User, error) {
	var user *are_hub.User

	return user, u.deleteID(ctx, id, user)
}
