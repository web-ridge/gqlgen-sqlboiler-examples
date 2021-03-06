package helpers

import (
	"context"
	"errors"

	"github.com/web-ridge/gqlgen-sqlboiler-examples/issue-6-edges-connections/graphql_models"
	"github.com/web-ridge/gqlgen-sqlboiler-examples/issue-6-edges-connections/models"
	"github.com/web-ridge/utils-go/boilergql/v3"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"database/sql"
)

var UserSortColumn = map[graphql_models.UserSort]string{
	graphql_models.UserSortID:		models.UserColumns.ID,
	graphql_models.UserSortFirstName:	models.UserColumns.FirstName,
	graphql_models.UserSortLastName:	models.UserColumns.LastName,
	graphql_models.UserSortAge:		models.UserColumns.Age,
	graphql_models.UserSortEmail:		models.UserColumns.Email,
}

func UserSortValueFromCursorValue(cursorValue string) (string, interface{}) {
	key, value := boilergql.FromCursorValue(cursorValue)
	column := UserSortColumn[graphql_models.UserSort(key)]

	if graphql_models.UserSort(key) == graphql_models.UserSortID {
		return column, boilergql.GetIDFromCursor(value)
	}

	return column, boilergql.StringToInterface(value)
}

func UserSortCursorValue(sort graphql_models.UserSort, m *graphql_models.User) interface{} {
	switch sort {
	case graphql_models.UserSortID:
		return m.ID
	case graphql_models.UserSortFirstName:
		return m.FirstName
	case graphql_models.UserSortLastName:
		return m.LastName
	case graphql_models.UserSortAge:
		return m.Age
	case graphql_models.UserSortEmail:
		return m.Email
	}
	return nil
}

func UserSortDirection(ordering []*graphql_models.UserOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func FromUserCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	var columns []string
	var values []interface{}

	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := UserSortValueFromCursorValue(cursorValue)
		if column != "" && value != nil {
			columns = append(columns, column)
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

func ToUserCursor(ordering []*graphql_models.UserOrdering, m *graphql_models.User) string {
	var a []string
	var handledID bool

	for _, order := range ordering {
		if order.Sort == graphql_models.UserSortID {
			handledID = true
		}
		value := UserSortCursorValue(order.Sort, m)
		if value != nil {
			a = append(a, boilergql.ToCursorValue(string(order.Sort), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(graphql_models.UserSortID), m.ID))
	}

	return boilergql.CursorValuesToString(a)
}

func UserCursorType(ordering []*graphql_models.UserOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func UserCursorMods(ordering []*graphql_models.UserOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if UserCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromUserCursor(*cursor, sign)
		}
		return boilergql.FromOffsetCursor(*cursor)
	}
	return nil
}

func UserSortMods(ordering []*graphql_models.UserOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, order := range ordering {
		if order.Sort == graphql_models.UserSortID {
			handledID = true
		}
		column := UserSortColumn[order.Sort]
		if column != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				column,
				boilergql.GetDirection(order.Direction, reverse),
			)))
		}
	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			models.UserColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func UserPaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*graphql_models.UserOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := UserSortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, UserCursorMods(ordering, cursor, sign)...)
	mods = append(mods, UserSortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func UserPaginationMods(pagination boilergql.ConnectionPagination, ordering []*graphql_models.UserOrdering) ([]qm.QueryMod, error) {
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

func ToUserCursorSwitch(ordering []*graphql_models.UserOrdering, m *graphql_models.User, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToUserCursor(ordering, m)
	}
	return ""
}

func UserReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.UserOrdering,
) (bool, error) {
	reverse := pagination.Forward != nil
	cursor, reverseMods := UserPaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := UserCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return models.Users(reverseMods...).Count(ctx, db)
	})
}

func UserEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*graphql_models.UserOrdering) func(*models.User, int) *graphql_models.UserEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), UserCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *models.User, i int) *graphql_models.UserEdge {
		n := UserToGraphQL(m)
		return &graphql_models.UserEdge{
			Cursor:	ToUserCursorSwitch(ordering, n, cursorType, offset, i),
			Node:	n,
		}
	}
}

func UserStartEndCursor(edges []*graphql_models.UserEdge) (*string, *string) {
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
	ordering []*graphql_models.UserOrdering,
) (*graphql_models.UserConnection, error) {
	paginationMods, err := UserPaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := UserReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := models.Users(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*graphql_models.UserEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := UserEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := UserStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
	return &graphql_models.UserConnection{
		Edges:	edges,
		PageInfo: &graphql_models.PageInfo{
			HasNextPage:		hasNextPage,
			HasPreviousPage:	hasPreviousPage,
			StartCursor:		startCursor,
			EndCursor:		endCursor,
		},
	}, nil
}
