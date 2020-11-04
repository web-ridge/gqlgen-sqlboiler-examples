// Generated with https://github.com/web-ridge/gqlgen-sqlboiler.
package main

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/web-ridge/utils-go/boilergql"

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

func UserSortDirection(ordering []*fm.UserOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func UserCursorType(ordering []*fm.UserOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func UserCursorMods(ordering []*fm.UserOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if UserCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromUserCursor(*cursor, sign)
		}
		return FromOffsetCursor(*cursor)
	}
	return nil
}

func UserPaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*fm.UserOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := UserSortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, UserCursorMods(ordering, cursor, sign)...)
	mods = append(mods, UserSortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func UserPaginationMods(pagination boilergql.ConnectionPagination, ordering []*fm.UserOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	reverse := pagination.Backward != nil
	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	_, mods := UserPaginationModsBase(pagination, ordering, reverse, limit)
	return mods, nil
}

func UserSortMods(ordering []*fm.UserOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, column := range ordering {
		var dbColumn string
		switch column.Sort {
		case fm.UserSortID:
			dbColumn, handledID = dm.UserColumns.ID, true
		case fm.UserSortAge:
			dbColumn = dm.UserColumns.Age
		case fm.UserSortFirstName:
			dbColumn = dm.UserColumns.FirstName
		case fm.UserSortLastName:
			dbColumn = dm.UserColumns.LastName
		case fm.UserSortEmail:
			dbColumn = dm.UserColumns.Email
		}
		if dbColumn != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				dbColumn,
				boilergql.GetDirection(column.Direction, reverse),
			)))
		}

	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			dm.UserColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func ToUserCursorSwitch(ordering []*fm.UserOrdering, user *fm.User, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToUserCursor(ordering, user)
	}
	return ""
}

func ToUserCursor(ordering []*fm.UserOrdering, m *fm.User) string {
	var a []string

	var handledID bool
	for _, column := range ordering {
		var col fm.UserSort
		var value interface{}
		switch column.Sort {
		case fm.UserSortID:
			col, value, handledID = fm.UserSortID, m.ID, true
		case fm.UserSortAge:
			col, value = fm.UserSortAge, m.Age
		case fm.UserSortFirstName:
			col, value = fm.UserSortFirstName, m.FirstName
		case fm.UserSortLastName:
			col, value = fm.UserSortLastName, m.LastName
		case fm.UserSortEmail:
			col, value = fm.UserSortEmail, m.Email
		}
		if col != "" {
			a = append(a, boilergql.ToCursorValue(string(col), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(fm.UserSortID), m.ID))
	}

	return boilergql.CursorValuesToString(a)
}

func FromOffsetCursor(cursor string) []qm.QueryMod {
	offset, _ := strconv.Atoi(cursor)

	if offset > 0 {
		return []qm.QueryMod{
			qm.Offset(offset),
		}
	}
	return nil
}

func FromUserCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	// b, _ := base64.StdEncoding.DecodeString(cursor)

	var columns []string
	var values []interface{}

	// https://www.slideshare.net/MarkusWinand/p2d2-pagination-done-the-postgresql-way
	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := boilergql.FromCursorValue(cursorValue)
		switch fm.UserSort(column) {
		case fm.UserSortID:
			columns = append(columns, dm.UserColumns.ID)
			values = append(values, boilergql.GetIDFromCursor(value))
		case fm.UserSortAge:
			columns = append(columns, dm.UserColumns.Age)
			values = append(values, value)
		case fm.UserSortFirstName:
			columns = append(columns, dm.UserColumns.FirstName)
			values = append(values, value)
		case fm.UserSortLastName:
			columns = append(columns, dm.UserColumns.LastName)
			values = append(values, value)
		case fm.UserSortEmail:
			columns = append(columns, dm.UserColumns.Email)
			values = append(values, value)
		}
	}

	if len(columns) > 0 {
		return []qm.QueryMod{
			qm.Where(boilergql.GetCursorWhere(comparisonSign, columns, values), values...),
		}
	}
	return nil
}

func UserReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*fm.UserOrdering,
) (bool, error) {
	// reverse query to check hasPrevious / hasNext
	reverse := pagination.Forward != nil
	cursor, reverseMods := UserPaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := UserCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return dm.Users(reverseMods...).Count(ctx, db)
	})
}

func UserEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*fm.UserOrdering) func(*dm.User, int) *fm.UserEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), UserCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *dm.User, i int) *fm.UserEdge {
		n := UserToGraphQL(m)
		return &fm.UserEdge{
			Cursor: ToUserCursorSwitch(ordering, n, cursorType, offset, i),
			Node:   n,
		}
	}
}

func UserStartEndCursor(edges []*fm.UserEdge) (*string, *string) {
	var startCursor, endCursor *string
	if len(edges) >= 2 {
		s, e := edges[0].Cursor, edges[len(edges)-1].Cursor
		startCursor = &s
		endCursor = &e
	} else if len(edges) == 1 {
		c := edges[0].Cursor
		startCursor = &c
		endCursor = &c
	}
	return startCursor, endCursor
}

func UserConnection(
	ctx context.Context,
	db *sql.DB,
	originalMods []qm.QueryMod,
	pagination boilergql.ConnectionPagination,
	ordering []*fm.UserOrdering,
) (*fm.UserConnection, error) {
	paginationMods, err := UserPaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := UserReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := dm.Users(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*fm.UserEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := UserEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := UserStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
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
