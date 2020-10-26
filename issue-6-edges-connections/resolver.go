// Generated with https://github.com/web-ridge/gqlgen-sqlboiler.
package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/rs/zerolog/log"
	fm "github.com/web-ridge/gqlgen-sqlboiler-examples/issue-6-edges-connections/graphql_models"
	. "github.com/web-ridge/gqlgen-sqlboiler-examples/issue-6-edges-connections/helpers"
	dm "github.com/web-ridge/gqlgen-sqlboiler-examples/issue-6-edges-connections/models"
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

func getIDFromCursor(id string) interface{} {
	splitID := strings.SplitN(id, "-", 2)
	if len(splitID) != 2 {
		return 0
	}
	return splitID[1]
}

func getReverseLimitAndCursor(forward *fm.ConnectionForwardPagination, backward *fm.ConnectionBackwardPagination) (*string, string, string, int) {
	if forward != nil {
		return forward.After, "<=", "DESC", zeroOrMore(forward.First + 1)
	}
	if backward != nil {
		return backward.Before, ">=", "ASC", zeroOrMore(backward.Last + 1)
	}
	return nil, "", "", 0
}

func zeroOrMore(limit int) int {
	if limit < 0 {
		return 0
	}
	return limit
}

func getLimitAndCursor(forward *fm.ConnectionForwardPagination, backward *fm.ConnectionBackwardPagination) (*string, string, string, int) {
	if forward != nil {
		return forward.After, ">", "ASC", zeroOrMore(forward.First + 1)
	}
	if backward != nil {
		return backward.Before, "<", "DESC", zeroOrMore(backward.Last + 1)
	}
	return nil, "", "", 0
}

func UserPaginationMods(pagination fm.ConnectionPagination, ordering []*fm.UserOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	var mods []qm.QueryMod
	cursor, comparisonSign, direction, limit := getLimitAndCursor(pagination.Forward, pagination.Backward)
	if cursor != nil {
		mods = append(mods, FromUserCursor(*cursor, comparisonSign)...)
	}

	mods = append(mods, ToUserSortMods(ordering, direction)...)
	mods = append(mods, qm.Limit(limit))
	return mods, nil
}

func ToUserSortMods(ordering []*fm.UserOrdering, sortDirection string) []qm.QueryMod {
	var a []qm.QueryMod

	for _, column := range ordering {
		switch column.Sort {
		case fm.UserSortFirstName:
			a = append(a, qm.OrderBy(dm.UserColumns.FirstName+" "+sortDirection))
		case fm.UserSortLastName:
			a = append(a, qm.OrderBy(dm.UserColumns.LastName+" "+sortDirection))
		case fm.UserSortEmail:
			a = append(a, qm.OrderBy(dm.UserColumns.Email+" "+sortDirection))
		}
	}
	a = append(a, qm.OrderBy(dm.UserColumns.ID+" "+sortDirection))
	return a
}

const (
	Separator1 = ",____,"
	Separator2 = ":"
)

func ToUserCursor(ordering []*fm.UserOrdering, user *fm.User) string {
	var a []string

	for _, column := range ordering {
		switch column.Sort {
		case fm.UserSortID:
			a = append(a, fmt.Sprintf("%v%v%v", fm.UserSortID, Separator2, user.ID))
		case fm.UserSortFirstName:
			a = append(a, fmt.Sprintf("%v%v%v", fm.UserSortFirstName, Separator2, user.FirstName))
		case fm.UserSortLastName:
			a = append(a, fmt.Sprintf("%v%v%v", fm.UserSortLastName, Separator2, user.LastName))
		case fm.UserSortEmail:
			a = append(a, fmt.Sprintf("%v%v%v", fm.UserSortEmail, Separator2, user.Email))
		}
	}
	if len(a) == 0 {
		a = append(a, fmt.Sprintf("%v%v%v", fm.UserSortID, Separator2, user.ID))
	}

	return strings.Join(a, Separator1)
}

func FromUserCursor(v string, comparisonSign string) []qm.QueryMod {
	var a []qm.QueryMod

	for _, columnAndValue := range strings.Split(v, Separator1) {
		s := strings.SplitN(columnAndValue, Separator2, 2)
		if len(s) != 2 {
			continue
		}
		column := fm.UserSort(s[0])
		value := s[1]
		switch column {
		case fm.UserSortID:
			a = append(a, qm.Where(dm.UserColumns.ID+" "+comparisonSign+" ?", getIDFromCursor(value)))
		case fm.UserSortFirstName:
			a = append(a, qm.Where(dm.UserColumns.FirstName+" "+comparisonSign+" ?", value))
		case fm.UserSortLastName:
			a = append(a, qm.Where(dm.UserColumns.LastName+" "+comparisonSign+" ?", value))
		case fm.UserSortEmail:
			a = append(a, qm.Where(dm.UserColumns.Email+" "+comparisonSign+" ?", value))
		}
	}
	return a
}

func ToUserConnection(ctx context.Context, db *sql.DB, originalMods []qm.QueryMod, pagination fm.ConnectionPagination, ordering []*fm.UserOrdering) (*fm.UserConnection, error) {
	paginationMods, err := UserPaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := dm.Users(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}

	// reverse query to check hasPrevious / hasNext
	reverseMods := originalMods

	cursor, comparisonSign, sortDirection, limit := getReverseLimitAndCursor(pagination.Forward, pagination.Backward)
	if cursor != nil {
		reverseMods = append(reverseMods, FromUserCursor(*cursor, comparisonSign)...)
	}
	reverseMods = append(reverseMods, ToUserSortMods(ordering, sortDirection)...)
	reverseMods = append(reverseMods, qm.Limit(1))

	var hasNextPage, hasPreviousPage bool
	var startCursor, endCursor *string
	if cursor != nil {
		reverseCount, err := dm.Users(reverseMods...).Count(ctx, db)
		if err != nil {
			return nil, err
		}

		if reverseCount > 0 {
			if pagination.Backward != nil {
				hasNextPage = true
			}
			if pagination.Forward != nil {
				hasPreviousPage = true
			}
		}
	}

	maxLength := limit - 1
	lowestLength := math.Min(float64(len(a)), float64(maxLength))
	edges := make([]*fm.UserEdge, 0, int(lowestLength))

	switch {
	case pagination.Backward != nil:
		// If no less last+1 results are returned, set hasPreviousPage: true, otherwise I set it to false.
		if len(a) == limit {
			hasPreviousPage = true
		}

		// If the last argument is provided, reverse the order of the results
		for i := len(a) - 1; i >= 0; i-- {
			isLast := i == maxLength
			if isLast {
				continue
			}
			n := UserToGraphQL(a[i])
			edges = append(edges, &fm.UserEdge{
				Cursor: ToUserCursor(ordering, n),
				Node:   n,
			})
		}

	case pagination.Forward != nil:

		// If no less than first+1 results are returned,  set hasNextPage: true, otherwise set it to false.
		if len(a) == limit {
			hasNextPage = true
		}
		for i, row := range a {
			isLast := i == maxLength
			if isLast {
				break
			}
			n := UserToGraphQL(row)
			edges = append(edges, &fm.UserEdge{
				Cursor: ToUserCursor(ordering, n),
				Node:   n,
			})
		}
	}

	fmt.Println("edges length", len(edges))
	if len(edges) >= 2 {
		s, e := edges[0].Cursor, edges[len(edges)-1].Cursor
		startCursor = &s
		endCursor = &e
	} else if len(edges) == 1 {
		c := edges[0].Cursor
		startCursor = &c
		endCursor = &c
	}

	return &fm.UserConnection{
		Edges: edges,
		PageInfo: &fm.PageInfo{
			HasNextPage:     hasNextPage,
			HasPreviousPage: hasPreviousPage,
			StartCursor:     startCursor,
			EndCursor:       endCursor,
		},
	}, nil
}

func (r *queryResolver) Users(ctx context.Context, pagination fm.ConnectionPagination, ordering []*fm.UserOrdering, filter *fm.UserFilter) (*fm.UserConnection, error) {
	mods := GetUserPreloadMods(ctx)
	mods = append(mods, UserFilterToMods(filter)...)
	connection, err := ToUserConnection(ctx, r.db, mods, pagination, ordering)
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

func (r *Resolver) Query() fm.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
