package mongodb

import (
	"context"

	"github.com/blacksfk/are_hub"
	"go.mongodb.org/mongo-driver/mongo"
)

// Implements are_hub.UserRepo.
type User struct {
	collection
}

// Create a users collection in the database.
func NewUserCollection(client *mongo.Client, db string) User {
	return User{collection{client, db, "users"}}
}

func (u User) All(ctx context.Context) ([]are_hub.User, error) {
	var users []are_hub.User

	return users, u.all(ctx, &users)
}

func (u User) FindID(ctx context.Context, id string) (*are_hub.User, error) {
	var user *are_hub.User

	return user, u.findID(ctx, id, user)
}

func (u User) DeleteID(ctx context.Context, id string) (*are_hub.User, error) {
	var user *are_hub.User

	return user, u.deleteID(ctx, id, user)
}
