// Generated with https://github.com/web-ridge/gqlgen-sqlboiler.
package main

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/rs/zerolog/log"
	fm "github.com/web-ridge/gqlgen-sqlboiler-examples/issue-6-edges-connections/graphql_models"
	. "github.com/web-ridge/gqlgen-sqlboiler-examples/issue-6-edges-connections/helpers"
	dm "github.com/web-ridge/gqlgen-sqlboiler-examples/issue-6-edges-connections/models"
	"github.com/web-ridge/utils-go/boilergql"
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

func (r *queryResolver) Users(ctx context.Context, pagination boilergql.ConnectionPagination, ordering []*fm.UserOrdering, filter *fm.UserFilter) (*fm.UserConnection, error) {
	mods := GetUserPreloadMods(ctx)

	// TODO: use these if no connection is used
	// return UsersToGraphQL(a), nil

	mods = append(mods, UserFilterToMods(filter)...)

	a, err := dm.Users(mods...).All(ctx, r.db)
	if err != nil {
		log.Error().Err(err).Msg(publicUserListError)
		return nil, errors.New(publicUserListError)
	}

	edges := make([]*fm.UserEdge, len(a))
	for i, row := range a {
		edges[i] = &fm.UserEdge{
			Cursor: "",
			Node:   UserToGraphQL(row),
		}
	}

	return &fm.UserConnection{
		Edges: edges,
		PageInfo: &fm.PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
			StartCursor:     nil,
			EndCursor:       nil,
		},
	}, nil
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

func (r *Resolver) Query() fm.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
