// Generated with https://github.com/web-ridge/gqlgen-sqlboiler.
package main

import (
	"context"
	"database/sql"
	"errors"

	"github.com/web-ridge/utils-go/boilergql"

	"github.com/rs/zerolog/log"
	fm "github.com/web-ridge/gqlgen-sqlboiler-examples/issue-6-edges-connections/graphql_models"
	. "github.com/web-ridge/gqlgen-sqlboiler-examples/issue-6-edges-connections/helpers"
	dm "github.com/web-ridge/gqlgen-sqlboiler-examples/issue-6-edges-connections/models"
)

type Resolver struct {
	db *sql.DB
}

const publicUserSingleError = "could not get user"

func (r *queryResolver) User(ctx context.Context, id string) (*fm.User, error) {
	dbID := UserID(id)

	mods := GetUserPreloadMods(ctx)
	mods = append(mods, dm.UserWhere.ID.EQ(dbID))

	m, err := dm.Users(mods...).One(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicUserSingleError)
		return nil, errors.New(publicUserSingleError)
	}
	return UserToGraphQL(m), nil
}

const publicUserListError = "could not list users"

func (r *queryResolver) Users(
	ctx context.Context,
	pagination boilergql.ConnectionPagination,
	ordering []*fm.UserOrdering,
	filter *fm.UserFilter,
) (*fm.UserConnection, error) {
	mods := GetUserPreloadMods(ctx)
	mods = append(mods, UserFilterToMods(filter)...)
	connection, err := UserConnection(ctx, r.db, mods, pagination, ordering)
	if err != nil {
		log.Error().Err(err).Msg(publicUserListError)
		return nil, errors.New(publicUserListError)
	}
	return connection, nil
}

func (r *queryResolver) Node(ctx context.Context, id string) (fm.Node, error) {
	model := boilergql.GetModelFromCursor(id)
	switch model {
	case "User":
		return r.User(ctx, id)
	default:
		return nil, errors.New("could not find corresponding model for id")
	}
}

func (r *Resolver) Query() fm.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
