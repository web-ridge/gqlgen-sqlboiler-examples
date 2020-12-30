package main

import (
	"context"
	"errors"
	"strings"

	"github.com/web-ridge/utils-go/boilergql/v3"

	"database/sql"

	"github.com/rs/zerolog/log"

	. "github.com/web-ridge/gqlgen-sqlboiler-examples/issue-6-edges-connections/helpers"

	dm "github.com/web-ridge/gqlgen-sqlboiler-examples/issue-6-edges-connections/models"

	fm "github.com/web-ridge/gqlgen-sqlboiler-examples/issue-6-edges-connections/graphql_models"
)

type Resolver struct {
	db *sql.DB
}

const inputKey = "input"

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

func (r *queryResolver) Users(ctx context.Context, first int, after *string, ordering []*fm.UserOrdering, filter *fm.UserFilter) (*fm.UserConnection, error) {
	mods := GetUserNodePreloadMods(ctx)

	mods = append(mods, UserFilterToMods(filter)...)
	connection, err := UserConnection(ctx, r.db, mods, boilergql.NewForwardPagination(first, after), ordering)
	if err != nil {
		log.Error().Err(err).Msg(publicUserListError)
		return nil, errors.New(publicUserListError)
	}
	return connection, nil
}

func (r *queryResolver) Node(ctx context.Context, globalGraphID string) (fm.Node, error) {
	splitID := strings.SplitN(globalGraphID, "-", 1)
	if len(splitID) != 2 {
		return nil, errors.New("could not parse id")
	}

	model := splitID[0]
	switch model {
	case "User":
		return r.User(ctx, globalGraphID)
	default:
		return nil, errors.New("could not find corresponding model for id")
	}
}

func (r *Resolver) Query() fm.QueryResolver	{ return &queryResolver{r} }

type queryResolver struct{ *Resolver }
