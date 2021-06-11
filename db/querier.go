package db

import "context"

type Querier interface {
	Tx(func(Querier) error) error

	CreateUser(context.Context, int64) error
	UpdateUser(ctx context.Context, user User) error

	InsertChar(ctx context.Context, arg InsertCharParams) error
	GiveChar(ctx context.Context, arg GiveCharParams) (Character, error)

	GetChar(context.Context, GetCharParams) (Character, error)
	GetChars(ctx context.Context, userID int64) ([]Character, error)

	GetProfile(context.Context, int64) (Profile, error)
}
