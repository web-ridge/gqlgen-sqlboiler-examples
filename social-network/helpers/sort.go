package helpers

import (
	"context"
	"errors"

	"github.com/web-ridge/gqlgen-sqlboiler-examples/social-network/graphql_models"
	"github.com/web-ridge/gqlgen-sqlboiler-examples/social-network/models"
	"github.com/web-ridge/utils-go/boilergql/v3"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"database/sql"
)

var CommentLikeSortColumn = map[graphql_models.CommentLikeSort]string{
	graphql_models.CommentLikeSortID:		models.CommentLikeColumns.ID,
	graphql_models.CommentLikeSortLikeType:		models.CommentLikeColumns.LikeType,
	graphql_models.CommentLikeSortCreatedAt:	models.CommentLikeColumns.CreatedAt,
}

func CommentLikeSortValueFromCursorValue(cursorValue string) (string, interface{}) {
	key, value := boilergql.FromCursorValue(cursorValue)
	column := CommentLikeSortColumn[graphql_models.CommentLikeSort(key)]

	if graphql_models.CommentLikeSort(key) == graphql_models.CommentLikeSortID {
		return column, boilergql.GetIDFromCursor(value)
	}

	return column, boilergql.StringToInterface(value)
}

func CommentLikeSortCursorValue(sort graphql_models.CommentLikeSort, m *graphql_models.CommentLike) interface{} {
	switch sort {
	case graphql_models.CommentLikeSortID:
		return m.ID
	case graphql_models.CommentLikeSortLikeType:
		return m.LikeType
	case graphql_models.CommentLikeSortCreatedAt:
		return m.CreatedAt
	}
	return nil
}

func CommentLikeSortDirection(ordering []*graphql_models.CommentLikeOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func FromCommentLikeCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	var columns []string
	var values []interface{}

	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := CommentLikeSortValueFromCursorValue(cursorValue)
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

func ToCommentLikeCursor(ordering []*graphql_models.CommentLikeOrdering, m *graphql_models.CommentLike) string {
	var a []string
	var handledID bool

	for _, order := range ordering {
		if order.Sort == graphql_models.CommentLikeSortID {
			handledID = true
		}
		value := CommentLikeSortCursorValue(order.Sort, m)
		if value != nil {
			a = append(a, boilergql.ToCursorValue(string(order.Sort), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(graphql_models.CommentLikeSortID), m.ID))
	}

	return boilergql.CursorValuesToString(a)
}

func CommentLikeCursorType(ordering []*graphql_models.CommentLikeOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func CommentLikeCursorMods(ordering []*graphql_models.CommentLikeOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if CommentLikeCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromCommentLikeCursor(*cursor, sign)
		}
		return boilergql.FromOffsetCursor(*cursor)
	}
	return nil
}

func CommentLikeSortMods(ordering []*graphql_models.CommentLikeOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, order := range ordering {
		if order.Sort == graphql_models.CommentLikeSortID {
			handledID = true
		}
		column := CommentLikeSortColumn[order.Sort]
		if column != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				column,
				boilergql.GetDirection(order.Direction, reverse),
			)))
		}
	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			models.CommentLikeColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func CommentLikePaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*graphql_models.CommentLikeOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := CommentLikeSortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, CommentLikeCursorMods(ordering, cursor, sign)...)
	mods = append(mods, CommentLikeSortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func CommentLikePaginationMods(pagination boilergql.ConnectionPagination, ordering []*graphql_models.CommentLikeOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	reverse := pagination.Backward != nil
	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	_, mods := CommentLikePaginationModsBase(pagination, ordering, reverse, limit)
	return mods, nil
}

func ToCommentLikeCursorSwitch(ordering []*graphql_models.CommentLikeOrdering, m *graphql_models.CommentLike, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToCommentLikeCursor(ordering, m)
	}
	return ""
}

func CommentLikeReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.CommentLikeOrdering,
) (bool, error) {
	reverse := pagination.Forward != nil
	cursor, reverseMods := CommentLikePaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := CommentLikeCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return models.CommentLikes(reverseMods...).Count(ctx, db)
	})
}

func CommentLikeEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*graphql_models.CommentLikeOrdering) func(*models.CommentLike, int) *graphql_models.CommentLikeEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), CommentLikeCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *models.CommentLike, i int) *graphql_models.CommentLikeEdge {
		n := CommentLikeToGraphQL(m)
		return &graphql_models.CommentLikeEdge{
			Cursor:	ToCommentLikeCursorSwitch(ordering, n, cursorType, offset, i),
			Node:	n,
		}
	}
}

func CommentLikeStartEndCursor(edges []*graphql_models.CommentLikeEdge) (*string, *string) {
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

func CommentLikeConnection(
	ctx context.Context,
	db *sql.DB,
	originalMods []qm.QueryMod,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.CommentLikeOrdering,
) (*graphql_models.CommentLikeConnection, error) {
	paginationMods, err := CommentLikePaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := CommentLikeReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := models.CommentLikes(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*graphql_models.CommentLikeEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := CommentLikeEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := CommentLikeStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
	return &graphql_models.CommentLikeConnection{
		Edges:	edges,
		PageInfo: &graphql_models.PageInfo{
			HasNextPage:		hasNextPage,
			HasPreviousPage:	hasPreviousPage,
			StartCursor:		startCursor,
			EndCursor:		endCursor,
		},
	}, nil
}

var CommentSortColumn = map[graphql_models.CommentSort]string{
	graphql_models.CommentSortID:		models.CommentColumns.ID,
	graphql_models.CommentSortContent:	models.CommentColumns.Content,
}

func CommentSortValueFromCursorValue(cursorValue string) (string, interface{}) {
	key, value := boilergql.FromCursorValue(cursorValue)
	column := CommentSortColumn[graphql_models.CommentSort(key)]

	if graphql_models.CommentSort(key) == graphql_models.CommentSortID {
		return column, boilergql.GetIDFromCursor(value)
	}

	return column, boilergql.StringToInterface(value)
}

func CommentSortCursorValue(sort graphql_models.CommentSort, m *graphql_models.Comment) interface{} {
	switch sort {
	case graphql_models.CommentSortID:
		return m.ID
	case graphql_models.CommentSortContent:
		return m.Content
	}
	return nil
}

func CommentSortDirection(ordering []*graphql_models.CommentOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func FromCommentCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	var columns []string
	var values []interface{}

	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := CommentSortValueFromCursorValue(cursorValue)
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

func ToCommentCursor(ordering []*graphql_models.CommentOrdering, m *graphql_models.Comment) string {
	var a []string
	var handledID bool

	for _, order := range ordering {
		if order.Sort == graphql_models.CommentSortID {
			handledID = true
		}
		value := CommentSortCursorValue(order.Sort, m)
		if value != nil {
			a = append(a, boilergql.ToCursorValue(string(order.Sort), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(graphql_models.CommentSortID), m.ID))
	}

	return boilergql.CursorValuesToString(a)
}

func CommentCursorType(ordering []*graphql_models.CommentOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func CommentCursorMods(ordering []*graphql_models.CommentOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if CommentCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromCommentCursor(*cursor, sign)
		}
		return boilergql.FromOffsetCursor(*cursor)
	}
	return nil
}

func CommentSortMods(ordering []*graphql_models.CommentOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, order := range ordering {
		if order.Sort == graphql_models.CommentSortID {
			handledID = true
		}
		column := CommentSortColumn[order.Sort]
		if column != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				column,
				boilergql.GetDirection(order.Direction, reverse),
			)))
		}
	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			models.CommentColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func CommentPaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*graphql_models.CommentOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := CommentSortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, CommentCursorMods(ordering, cursor, sign)...)
	mods = append(mods, CommentSortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func CommentPaginationMods(pagination boilergql.ConnectionPagination, ordering []*graphql_models.CommentOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	reverse := pagination.Backward != nil
	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	_, mods := CommentPaginationModsBase(pagination, ordering, reverse, limit)
	return mods, nil
}

func ToCommentCursorSwitch(ordering []*graphql_models.CommentOrdering, m *graphql_models.Comment, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToCommentCursor(ordering, m)
	}
	return ""
}

func CommentReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.CommentOrdering,
) (bool, error) {
	reverse := pagination.Forward != nil
	cursor, reverseMods := CommentPaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := CommentCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return models.Comments(reverseMods...).Count(ctx, db)
	})
}

func CommentEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*graphql_models.CommentOrdering) func(*models.Comment, int) *graphql_models.CommentEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), CommentCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *models.Comment, i int) *graphql_models.CommentEdge {
		n := CommentToGraphQL(m)
		return &graphql_models.CommentEdge{
			Cursor:	ToCommentCursorSwitch(ordering, n, cursorType, offset, i),
			Node:	n,
		}
	}
}

func CommentStartEndCursor(edges []*graphql_models.CommentEdge) (*string, *string) {
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

func CommentConnection(
	ctx context.Context,
	db *sql.DB,
	originalMods []qm.QueryMod,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.CommentOrdering,
) (*graphql_models.CommentConnection, error) {
	paginationMods, err := CommentPaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := CommentReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := models.Comments(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*graphql_models.CommentEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := CommentEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := CommentStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
	return &graphql_models.CommentConnection{
		Edges:	edges,
		PageInfo: &graphql_models.PageInfo{
			HasNextPage:		hasNextPage,
			HasPreviousPage:	hasPreviousPage,
			StartCursor:		startCursor,
			EndCursor:		endCursor,
		},
	}, nil
}

var FriendshipSortColumn = map[graphql_models.FriendshipSort]string{
	graphql_models.FriendshipSortID:	models.FriendshipColumns.ID,
	graphql_models.FriendshipSortCreatedAt:	models.FriendshipColumns.CreatedAt,
}

func FriendshipSortValueFromCursorValue(cursorValue string) (string, interface{}) {
	key, value := boilergql.FromCursorValue(cursorValue)
	column := FriendshipSortColumn[graphql_models.FriendshipSort(key)]

	if graphql_models.FriendshipSort(key) == graphql_models.FriendshipSortID {
		return column, boilergql.GetIDFromCursor(value)
	}

	return column, boilergql.StringToInterface(value)
}

func FriendshipSortCursorValue(sort graphql_models.FriendshipSort, m *graphql_models.Friendship) interface{} {
	switch sort {
	case graphql_models.FriendshipSortID:
		return m.ID
	case graphql_models.FriendshipSortCreatedAt:
		return m.CreatedAt
	}
	return nil
}

func FriendshipSortDirection(ordering []*graphql_models.FriendshipOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func FromFriendshipCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	var columns []string
	var values []interface{}

	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := FriendshipSortValueFromCursorValue(cursorValue)
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

func ToFriendshipCursor(ordering []*graphql_models.FriendshipOrdering, m *graphql_models.Friendship) string {
	var a []string
	var handledID bool

	for _, order := range ordering {
		if order.Sort == graphql_models.FriendshipSortID {
			handledID = true
		}
		value := FriendshipSortCursorValue(order.Sort, m)
		if value != nil {
			a = append(a, boilergql.ToCursorValue(string(order.Sort), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(graphql_models.FriendshipSortID), m.ID))
	}

	return boilergql.CursorValuesToString(a)
}

func FriendshipCursorType(ordering []*graphql_models.FriendshipOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func FriendshipCursorMods(ordering []*graphql_models.FriendshipOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if FriendshipCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromFriendshipCursor(*cursor, sign)
		}
		return boilergql.FromOffsetCursor(*cursor)
	}
	return nil
}

func FriendshipSortMods(ordering []*graphql_models.FriendshipOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, order := range ordering {
		if order.Sort == graphql_models.FriendshipSortID {
			handledID = true
		}
		column := FriendshipSortColumn[order.Sort]
		if column != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				column,
				boilergql.GetDirection(order.Direction, reverse),
			)))
		}
	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			models.FriendshipColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func FriendshipPaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*graphql_models.FriendshipOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := FriendshipSortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, FriendshipCursorMods(ordering, cursor, sign)...)
	mods = append(mods, FriendshipSortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func FriendshipPaginationMods(pagination boilergql.ConnectionPagination, ordering []*graphql_models.FriendshipOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	reverse := pagination.Backward != nil
	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	_, mods := FriendshipPaginationModsBase(pagination, ordering, reverse, limit)
	return mods, nil
}

func ToFriendshipCursorSwitch(ordering []*graphql_models.FriendshipOrdering, m *graphql_models.Friendship, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToFriendshipCursor(ordering, m)
	}
	return ""
}

func FriendshipReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.FriendshipOrdering,
) (bool, error) {
	reverse := pagination.Forward != nil
	cursor, reverseMods := FriendshipPaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := FriendshipCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return models.Friendships(reverseMods...).Count(ctx, db)
	})
}

func FriendshipEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*graphql_models.FriendshipOrdering) func(*models.Friendship, int) *graphql_models.FriendshipEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), FriendshipCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *models.Friendship, i int) *graphql_models.FriendshipEdge {
		n := FriendshipToGraphQL(m)
		return &graphql_models.FriendshipEdge{
			Cursor:	ToFriendshipCursorSwitch(ordering, n, cursorType, offset, i),
			Node:	n,
		}
	}
}

func FriendshipStartEndCursor(edges []*graphql_models.FriendshipEdge) (*string, *string) {
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

func FriendshipConnection(
	ctx context.Context,
	db *sql.DB,
	originalMods []qm.QueryMod,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.FriendshipOrdering,
) (*graphql_models.FriendshipConnection, error) {
	paginationMods, err := FriendshipPaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := FriendshipReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := models.Friendships(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*graphql_models.FriendshipEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := FriendshipEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := FriendshipStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
	return &graphql_models.FriendshipConnection{
		Edges:	edges,
		PageInfo: &graphql_models.PageInfo{
			HasNextPage:		hasNextPage,
			HasPreviousPage:	hasPreviousPage,
			StartCursor:		startCursor,
			EndCursor:		endCursor,
		},
	}, nil
}

var ImageSortColumn = map[graphql_models.ImageSort]string{
	graphql_models.ImageSortID:		models.ImageColumns.ID,
	graphql_models.ImageSortViews:		models.ImageColumns.Views,
	graphql_models.ImageSortOriginalURL:	models.ImageColumns.OriginalURL,
}

func ImageSortValueFromCursorValue(cursorValue string) (string, interface{}) {
	key, value := boilergql.FromCursorValue(cursorValue)
	column := ImageSortColumn[graphql_models.ImageSort(key)]

	if graphql_models.ImageSort(key) == graphql_models.ImageSortID {
		return column, boilergql.GetIDFromCursor(value)
	}

	return column, boilergql.StringToInterface(value)
}

func ImageSortCursorValue(sort graphql_models.ImageSort, m *graphql_models.Image) interface{} {
	switch sort {
	case graphql_models.ImageSortID:
		return m.ID
	case graphql_models.ImageSortViews:
		return m.Views
	case graphql_models.ImageSortOriginalURL:
		return m.OriginalURL
	}
	return nil
}

func ImageSortDirection(ordering []*graphql_models.ImageOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func FromImageCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	var columns []string
	var values []interface{}

	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := ImageSortValueFromCursorValue(cursorValue)
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

func ToImageCursor(ordering []*graphql_models.ImageOrdering, m *graphql_models.Image) string {
	var a []string
	var handledID bool

	for _, order := range ordering {
		if order.Sort == graphql_models.ImageSortID {
			handledID = true
		}
		value := ImageSortCursorValue(order.Sort, m)
		if value != nil {
			a = append(a, boilergql.ToCursorValue(string(order.Sort), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(graphql_models.ImageSortID), m.ID))
	}

	return boilergql.CursorValuesToString(a)
}

func ImageCursorType(ordering []*graphql_models.ImageOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func ImageCursorMods(ordering []*graphql_models.ImageOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if ImageCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromImageCursor(*cursor, sign)
		}
		return boilergql.FromOffsetCursor(*cursor)
	}
	return nil
}

func ImageSortMods(ordering []*graphql_models.ImageOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, order := range ordering {
		if order.Sort == graphql_models.ImageSortID {
			handledID = true
		}
		column := ImageSortColumn[order.Sort]
		if column != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				column,
				boilergql.GetDirection(order.Direction, reverse),
			)))
		}
	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			models.ImageColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func ImagePaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*graphql_models.ImageOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := ImageSortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, ImageCursorMods(ordering, cursor, sign)...)
	mods = append(mods, ImageSortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func ImagePaginationMods(pagination boilergql.ConnectionPagination, ordering []*graphql_models.ImageOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	reverse := pagination.Backward != nil
	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	_, mods := ImagePaginationModsBase(pagination, ordering, reverse, limit)
	return mods, nil
}

func ToImageCursorSwitch(ordering []*graphql_models.ImageOrdering, m *graphql_models.Image, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToImageCursor(ordering, m)
	}
	return ""
}

func ImageReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.ImageOrdering,
) (bool, error) {
	reverse := pagination.Forward != nil
	cursor, reverseMods := ImagePaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := ImageCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return models.Images(reverseMods...).Count(ctx, db)
	})
}

func ImageEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*graphql_models.ImageOrdering) func(*models.Image, int) *graphql_models.ImageEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), ImageCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *models.Image, i int) *graphql_models.ImageEdge {
		n := ImageToGraphQL(m)
		return &graphql_models.ImageEdge{
			Cursor:	ToImageCursorSwitch(ordering, n, cursorType, offset, i),
			Node:	n,
		}
	}
}

func ImageStartEndCursor(edges []*graphql_models.ImageEdge) (*string, *string) {
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

func ImageConnection(
	ctx context.Context,
	db *sql.DB,
	originalMods []qm.QueryMod,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.ImageOrdering,
) (*graphql_models.ImageConnection, error) {
	paginationMods, err := ImagePaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := ImageReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := models.Images(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*graphql_models.ImageEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := ImageEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := ImageStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
	return &graphql_models.ImageConnection{
		Edges:	edges,
		PageInfo: &graphql_models.PageInfo{
			HasNextPage:		hasNextPage,
			HasPreviousPage:	hasPreviousPage,
			StartCursor:		startCursor,
			EndCursor:		endCursor,
		},
	}, nil
}

var ImageVariationSortColumn = map[graphql_models.ImageVariationSort]string{
	graphql_models.ImageVariationSortID: models.ImageVariationColumns.ID,
}

func ImageVariationSortValueFromCursorValue(cursorValue string) (string, interface{}) {
	key, value := boilergql.FromCursorValue(cursorValue)
	column := ImageVariationSortColumn[graphql_models.ImageVariationSort(key)]

	if graphql_models.ImageVariationSort(key) == graphql_models.ImageVariationSortID {
		return column, boilergql.GetIDFromCursor(value)
	}

	return column, boilergql.StringToInterface(value)
}

func ImageVariationSortCursorValue(sort graphql_models.ImageVariationSort, m *graphql_models.ImageVariation) interface{} {
	switch sort {
	case graphql_models.ImageVariationSortID:
		return m.ID
	}
	return nil
}

func ImageVariationSortDirection(ordering []*graphql_models.ImageVariationOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func FromImageVariationCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	var columns []string
	var values []interface{}

	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := ImageVariationSortValueFromCursorValue(cursorValue)
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

func ToImageVariationCursor(ordering []*graphql_models.ImageVariationOrdering, m *graphql_models.ImageVariation) string {
	var a []string
	var handledID bool

	for _, order := range ordering {
		if order.Sort == graphql_models.ImageVariationSortID {
			handledID = true
		}
		value := ImageVariationSortCursorValue(order.Sort, m)
		if value != nil {
			a = append(a, boilergql.ToCursorValue(string(order.Sort), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(graphql_models.ImageVariationSortID), m.ID))
	}
	return boilergql.CursorValuesToString(a)
}

func ImageVariationCursorType(ordering []*graphql_models.ImageVariationOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func ImageVariationCursorMods(ordering []*graphql_models.ImageVariationOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if ImageVariationCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromImageVariationCursor(*cursor, sign)
		}
		return boilergql.FromOffsetCursor(*cursor)
	}
	return nil
}

func ImageVariationSortMods(ordering []*graphql_models.ImageVariationOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, order := range ordering {
		if order.Sort == graphql_models.ImageVariationSortID {
			handledID = true
		}
		column := ImageVariationSortColumn[order.Sort]
		if column != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				column,
				boilergql.GetDirection(order.Direction, reverse),
			)))
		}
	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			models.ImageVariationColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func ImageVariationPaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*graphql_models.ImageVariationOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := ImageVariationSortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, ImageVariationCursorMods(ordering, cursor, sign)...)
	mods = append(mods, ImageVariationSortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func ImageVariationPaginationMods(pagination boilergql.ConnectionPagination, ordering []*graphql_models.ImageVariationOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	reverse := pagination.Backward != nil
	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	_, mods := ImageVariationPaginationModsBase(pagination, ordering, reverse, limit)
	return mods, nil
}

func ToImageVariationCursorSwitch(ordering []*graphql_models.ImageVariationOrdering, m *graphql_models.ImageVariation, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToImageVariationCursor(ordering, m)
	}
	return ""
}

func ImageVariationReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.ImageVariationOrdering,
) (bool, error) {
	reverse := pagination.Forward != nil
	cursor, reverseMods := ImageVariationPaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := ImageVariationCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return models.ImageVariations(reverseMods...).Count(ctx, db)
	})
}

func ImageVariationEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*graphql_models.ImageVariationOrdering) func(*models.ImageVariation, int) *graphql_models.ImageVariationEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), ImageVariationCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *models.ImageVariation, i int) *graphql_models.ImageVariationEdge {
		n := ImageVariationToGraphQL(m)
		return &graphql_models.ImageVariationEdge{
			Cursor:	ToImageVariationCursorSwitch(ordering, n, cursorType, offset, i),
			Node:	n,
		}
	}
}

func ImageVariationStartEndCursor(edges []*graphql_models.ImageVariationEdge) (*string, *string) {
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

func ImageVariationConnection(
	ctx context.Context,
	db *sql.DB,
	originalMods []qm.QueryMod,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.ImageVariationOrdering,
) (*graphql_models.ImageVariationConnection, error) {
	paginationMods, err := ImageVariationPaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := ImageVariationReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := models.ImageVariations(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*graphql_models.ImageVariationEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := ImageVariationEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := ImageVariationStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
	return &graphql_models.ImageVariationConnection{
		Edges:	edges,
		PageInfo: &graphql_models.PageInfo{
			HasNextPage:		hasNextPage,
			HasPreviousPage:	hasPreviousPage,
			StartCursor:		startCursor,
			EndCursor:		endCursor,
		},
	}, nil
}

var LikeSortColumn = map[graphql_models.LikeSort]string{
	graphql_models.LikeSortID:		models.LikeColumns.ID,
	graphql_models.LikeSortLikeType:	models.LikeColumns.LikeType,
	graphql_models.LikeSortCreatedAt:	models.LikeColumns.CreatedAt,
}

func LikeSortValueFromCursorValue(cursorValue string) (string, interface{}) {
	key, value := boilergql.FromCursorValue(cursorValue)
	column := LikeSortColumn[graphql_models.LikeSort(key)]

	if graphql_models.LikeSort(key) == graphql_models.LikeSortID {
		return column, boilergql.GetIDFromCursor(value)
	}

	return column, boilergql.StringToInterface(value)
}

func LikeSortCursorValue(sort graphql_models.LikeSort, m *graphql_models.Like) interface{} {
	switch sort {
	case graphql_models.LikeSortID:
		return m.ID
	case graphql_models.LikeSortLikeType:
		return m.LikeType
	case graphql_models.LikeSortCreatedAt:
		return m.CreatedAt
	}
	return nil
}

func LikeSortDirection(ordering []*graphql_models.LikeOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func FromLikeCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	var columns []string
	var values []interface{}

	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := LikeSortValueFromCursorValue(cursorValue)
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

func ToLikeCursor(ordering []*graphql_models.LikeOrdering, m *graphql_models.Like) string {
	var a []string
	var handledID bool

	for _, order := range ordering {
		if order.Sort == graphql_models.LikeSortID {
			handledID = true
		}
		value := LikeSortCursorValue(order.Sort, m)
		if value != nil {
			a = append(a, boilergql.ToCursorValue(string(order.Sort), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(graphql_models.LikeSortID), m.ID))
	}

	return boilergql.CursorValuesToString(a)
}

func LikeCursorType(ordering []*graphql_models.LikeOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func LikeCursorMods(ordering []*graphql_models.LikeOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if LikeCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromLikeCursor(*cursor, sign)
		}
		return boilergql.FromOffsetCursor(*cursor)
	}
	return nil
}

func LikeSortMods(ordering []*graphql_models.LikeOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, order := range ordering {
		if order.Sort == graphql_models.LikeSortID {
			handledID = true
		}
		column := LikeSortColumn[order.Sort]
		if column != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				column,
				boilergql.GetDirection(order.Direction, reverse),
			)))
		}
	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			models.LikeColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func LikePaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*graphql_models.LikeOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := LikeSortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, LikeCursorMods(ordering, cursor, sign)...)
	mods = append(mods, LikeSortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func LikePaginationMods(pagination boilergql.ConnectionPagination, ordering []*graphql_models.LikeOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	reverse := pagination.Backward != nil
	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	_, mods := LikePaginationModsBase(pagination, ordering, reverse, limit)
	return mods, nil
}

func ToLikeCursorSwitch(ordering []*graphql_models.LikeOrdering, m *graphql_models.Like, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToLikeCursor(ordering, m)
	}
	return ""
}

func LikeReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.LikeOrdering,
) (bool, error) {
	reverse := pagination.Forward != nil
	cursor, reverseMods := LikePaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := LikeCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return models.Likes(reverseMods...).Count(ctx, db)
	})
}

func LikeEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*graphql_models.LikeOrdering) func(*models.Like, int) *graphql_models.LikeEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), LikeCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *models.Like, i int) *graphql_models.LikeEdge {
		n := LikeToGraphQL(m)
		return &graphql_models.LikeEdge{
			Cursor:	ToLikeCursorSwitch(ordering, n, cursorType, offset, i),
			Node:	n,
		}
	}
}

func LikeStartEndCursor(edges []*graphql_models.LikeEdge) (*string, *string) {
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

func LikeConnection(
	ctx context.Context,
	db *sql.DB,
	originalMods []qm.QueryMod,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.LikeOrdering,
) (*graphql_models.LikeConnection, error) {
	paginationMods, err := LikePaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := LikeReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := models.Likes(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*graphql_models.LikeEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := LikeEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := LikeStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
	return &graphql_models.LikeConnection{
		Edges:	edges,
		PageInfo: &graphql_models.PageInfo{
			HasNextPage:		hasNextPage,
			HasPreviousPage:	hasPreviousPage,
			StartCursor:		startCursor,
			EndCursor:		endCursor,
		},
	}, nil
}

var PostSortColumn = map[graphql_models.PostSort]string{
	graphql_models.PostSortID:	models.PostColumns.ID,
	graphql_models.PostSortContent:	models.PostColumns.Content,
}

func PostSortValueFromCursorValue(cursorValue string) (string, interface{}) {
	key, value := boilergql.FromCursorValue(cursorValue)
	column := PostSortColumn[graphql_models.PostSort(key)]

	if graphql_models.PostSort(key) == graphql_models.PostSortID {
		return column, boilergql.GetIDFromCursor(value)
	}

	return column, boilergql.StringToInterface(value)
}

func PostSortCursorValue(sort graphql_models.PostSort, m *graphql_models.Post) interface{} {
	switch sort {
	case graphql_models.PostSortID:
		return m.ID
	case graphql_models.PostSortContent:
		return m.Content
	}
	return nil
}

func PostSortDirection(ordering []*graphql_models.PostOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func FromPostCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	var columns []string
	var values []interface{}

	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := PostSortValueFromCursorValue(cursorValue)
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

func ToPostCursor(ordering []*graphql_models.PostOrdering, m *graphql_models.Post) string {
	var a []string
	var handledID bool

	for _, order := range ordering {
		if order.Sort == graphql_models.PostSortID {
			handledID = true
		}
		value := PostSortCursorValue(order.Sort, m)
		if value != nil {
			a = append(a, boilergql.ToCursorValue(string(order.Sort), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(graphql_models.PostSortID), m.ID))
	}

	return boilergql.CursorValuesToString(a)
}

func PostCursorType(ordering []*graphql_models.PostOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func PostCursorMods(ordering []*graphql_models.PostOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if PostCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromPostCursor(*cursor, sign)
		}
		return boilergql.FromOffsetCursor(*cursor)
	}
	return nil
}

func PostSortMods(ordering []*graphql_models.PostOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, order := range ordering {
		if order.Sort == graphql_models.PostSortID {
			handledID = true
		}
		column := PostSortColumn[order.Sort]
		if column != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				column,
				boilergql.GetDirection(order.Direction, reverse),
			)))
		}
	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			models.PostColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func PostPaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*graphql_models.PostOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := PostSortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, PostCursorMods(ordering, cursor, sign)...)
	mods = append(mods, PostSortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func PostPaginationMods(pagination boilergql.ConnectionPagination, ordering []*graphql_models.PostOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	reverse := pagination.Backward != nil
	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	_, mods := PostPaginationModsBase(pagination, ordering, reverse, limit)
	return mods, nil
}

func ToPostCursorSwitch(ordering []*graphql_models.PostOrdering, m *graphql_models.Post, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToPostCursor(ordering, m)
	}
	return ""
}

func PostReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.PostOrdering,
) (bool, error) {
	reverse := pagination.Forward != nil
	cursor, reverseMods := PostPaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := PostCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return models.Posts(reverseMods...).Count(ctx, db)
	})
}

func PostEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*graphql_models.PostOrdering) func(*models.Post, int) *graphql_models.PostEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), PostCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *models.Post, i int) *graphql_models.PostEdge {
		n := PostToGraphQL(m)
		return &graphql_models.PostEdge{
			Cursor:	ToPostCursorSwitch(ordering, n, cursorType, offset, i),
			Node:	n,
		}
	}
}

func PostStartEndCursor(edges []*graphql_models.PostEdge) (*string, *string) {
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

func PostConnection(
	ctx context.Context,
	db *sql.DB,
	originalMods []qm.QueryMod,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.PostOrdering,
) (*graphql_models.PostConnection, error) {
	paginationMods, err := PostPaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := PostReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := models.Posts(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*graphql_models.PostEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := PostEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := PostStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
	return &graphql_models.PostConnection{
		Edges:	edges,
		PageInfo: &graphql_models.PageInfo{
			HasNextPage:		hasNextPage,
			HasPreviousPage:	hasPreviousPage,
			StartCursor:		startCursor,
			EndCursor:		endCursor,
		},
	}, nil
}

var UserSortColumn = map[graphql_models.UserSort]string{
	graphql_models.UserSortID:		models.UserColumns.ID,
	graphql_models.UserSortFirstName:	models.UserColumns.FirstName,
	graphql_models.UserSortLastName:	models.UserColumns.LastName,
	graphql_models.UserSortEmail:		models.UserColumns.Email,
	graphql_models.UserSortPassword:	models.UserColumns.Password,
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
	case graphql_models.UserSortEmail:
		return m.Email
	case graphql_models.UserSortPassword:
		return m.Password
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
