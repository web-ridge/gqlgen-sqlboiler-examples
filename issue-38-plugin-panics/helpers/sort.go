package helpers

import (
	"context"
	"errors"

	"github.com/web-ridge/gqlgen-sqlboiler-examples/issue-38-plugin-panics/graphql_models"
	"github.com/web-ridge/gqlgen-sqlboiler-examples/issue-38-plugin-panics/models"
	"github.com/web-ridge/utils-go/boilergql/v3"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"database/sql"
)

var AdditiveInventorySortColumn = map[graphql_models.AdditiveInventorySort]string{
	graphql_models.AdditiveInventorySortID:			models.AdditiveInventoryColumns.ID,
	graphql_models.AdditiveInventorySortPurchaseDate:	models.AdditiveInventoryColumns.PurchaseDate,
	graphql_models.AdditiveInventorySortExpiryDate:		models.AdditiveInventoryColumns.ExpiryDate,
	graphql_models.AdditiveInventorySortCost:		models.AdditiveInventoryColumns.Cost,
	graphql_models.AdditiveInventorySortWeight:		models.AdditiveInventoryColumns.Weight,
	graphql_models.AdditiveInventorySortCreatedAt:		models.AdditiveInventoryColumns.CreatedAt,
	graphql_models.AdditiveInventorySortUpdatedAt:		models.AdditiveInventoryColumns.UpdatedAt,
	graphql_models.AdditiveInventorySortDeletedAt:		models.AdditiveInventoryColumns.DeletedAt,
}

func AdditiveInventorySortValueFromCursorValue(cursorValue string) (string, interface{}) {
	key, value := boilergql.FromCursorValue(cursorValue)
	column := AdditiveInventorySortColumn[graphql_models.AdditiveInventorySort(key)]

	if graphql_models.AdditiveInventorySort(key) == graphql_models.AdditiveInventorySortID {
		return column, boilergql.GetIDFromCursor(value)
	}

	return column, boilergql.StringToInterface(value)
}

func AdditiveInventorySortCursorValue(sort graphql_models.AdditiveInventorySort, m *graphql_models.AdditiveInventory) interface{} {
	switch sort {
	case graphql_models.AdditiveInventorySortID:
		return m.ID
	case graphql_models.AdditiveInventorySortPurchaseDate:
		return m.PurchaseDate
	case graphql_models.AdditiveInventorySortExpiryDate:
		return m.ExpiryDate
	case graphql_models.AdditiveInventorySortCost:
		return m.Cost
	case graphql_models.AdditiveInventorySortWeight:
		return m.Weight
	case graphql_models.AdditiveInventorySortCreatedAt:
		return m.CreatedAt
	case graphql_models.AdditiveInventorySortUpdatedAt:
		return m.UpdatedAt
	case graphql_models.AdditiveInventorySortDeletedAt:
		return m.DeletedAt
	}
	return nil
}

func AdditiveInventorySortDirection(ordering []*graphql_models.AdditiveInventoryOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func FromAdditiveInventoryCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	var columns []string
	var values []interface{}

	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := AdditiveInventorySortValueFromCursorValue(cursorValue)
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

func ToAdditiveInventoryCursor(ordering []*graphql_models.AdditiveInventoryOrdering, m *graphql_models.AdditiveInventory) string {
	var a []string
	var handledID bool

	for _, order := range ordering {
		if order.Sort == graphql_models.AdditiveInventorySortID {
			handledID = true
		}
		value := AdditiveInventorySortCursorValue(order.Sort, m)
		if value != nil {
			a = append(a, boilergql.ToCursorValue(string(order.Sort), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(graphql_models.AdditiveInventorySortID), m.ID))
	}

	return boilergql.CursorValuesToString(a)
}

func AdditiveInventoryCursorType(ordering []*graphql_models.AdditiveInventoryOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func AdditiveInventoryCursorMods(ordering []*graphql_models.AdditiveInventoryOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if AdditiveInventoryCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromAdditiveInventoryCursor(*cursor, sign)
		}
		return boilergql.FromOffsetCursor(*cursor)
	}
	return nil
}

func AdditiveInventorySortMods(ordering []*graphql_models.AdditiveInventoryOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, order := range ordering {
		if order.Sort == graphql_models.AdditiveInventorySortID {
			handledID = true
		}
		column := AdditiveInventorySortColumn[order.Sort]
		if column != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				column,
				boilergql.GetDirection(order.Direction, reverse),
			)))
		}
	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			models.AdditiveInventoryColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func AdditiveInventoryPaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*graphql_models.AdditiveInventoryOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := AdditiveInventorySortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, AdditiveInventoryCursorMods(ordering, cursor, sign)...)
	mods = append(mods, AdditiveInventorySortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func AdditiveInventoryPaginationMods(pagination boilergql.ConnectionPagination, ordering []*graphql_models.AdditiveInventoryOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	reverse := pagination.Backward != nil
	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	_, mods := AdditiveInventoryPaginationModsBase(pagination, ordering, reverse, limit)
	return mods, nil
}

func ToAdditiveInventoryCursorSwitch(ordering []*graphql_models.AdditiveInventoryOrdering, m *graphql_models.AdditiveInventory, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToAdditiveInventoryCursor(ordering, m)
	}
	return ""
}

func AdditiveInventoryReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.AdditiveInventoryOrdering,
) (bool, error) {
	reverse := pagination.Forward != nil
	cursor, reverseMods := AdditiveInventoryPaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := AdditiveInventoryCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return models.AdditiveInventories(reverseMods...).Count(ctx, db)
	})
}

func AdditiveInventoryEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*graphql_models.AdditiveInventoryOrdering) func(*models.AdditiveInventory, int) *graphql_models.AdditiveInventoryEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), AdditiveInventoryCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *models.AdditiveInventory, i int) *graphql_models.AdditiveInventoryEdge {
		n := AdditiveInventoryToGraphQL(m)
		return &graphql_models.AdditiveInventoryEdge{
			Cursor:	ToAdditiveInventoryCursorSwitch(ordering, n, cursorType, offset, i),
			Node:	n,
		}
	}
}

func AdditiveInventoryStartEndCursor(edges []*graphql_models.AdditiveInventoryEdge) (*string, *string) {
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

func AdditiveInventoryConnection(
	ctx context.Context,
	db *sql.DB,
	originalMods []qm.QueryMod,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.AdditiveInventoryOrdering,
) (*graphql_models.AdditiveInventoryConnection, error) {
	paginationMods, err := AdditiveInventoryPaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := AdditiveInventoryReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := models.AdditiveInventories(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*graphql_models.AdditiveInventoryEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := AdditiveInventoryEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := AdditiveInventoryStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
	return &graphql_models.AdditiveInventoryConnection{
		Edges:	edges,
		PageInfo: &graphql_models.PageInfo{
			HasNextPage:		hasNextPage,
			HasPreviousPage:	hasPreviousPage,
			StartCursor:		startCursor,
			EndCursor:		endCursor,
		},
	}, nil
}

var AdditiveSortColumn = map[graphql_models.AdditiveSort]string{
	graphql_models.AdditiveSortID:		models.AdditiveColumns.ID,
	graphql_models.AdditiveSortName:	models.AdditiveColumns.Name,
	graphql_models.AdditiveSortNote:	models.AdditiveColumns.Note,
	graphql_models.AdditiveSortCreatedAt:	models.AdditiveColumns.CreatedAt,
	graphql_models.AdditiveSortUpdatedAt:	models.AdditiveColumns.UpdatedAt,
	graphql_models.AdditiveSortDeletedAt:	models.AdditiveColumns.DeletedAt,
}

func AdditiveSortValueFromCursorValue(cursorValue string) (string, interface{}) {
	key, value := boilergql.FromCursorValue(cursorValue)
	column := AdditiveSortColumn[graphql_models.AdditiveSort(key)]

	if graphql_models.AdditiveSort(key) == graphql_models.AdditiveSortID {
		return column, boilergql.GetIDFromCursor(value)
	}

	return column, boilergql.StringToInterface(value)
}

func AdditiveSortCursorValue(sort graphql_models.AdditiveSort, m *graphql_models.Additive) interface{} {
	switch sort {
	case graphql_models.AdditiveSortID:
		return m.ID
	case graphql_models.AdditiveSortName:
		return m.Name
	case graphql_models.AdditiveSortNote:
		return m.Note
	case graphql_models.AdditiveSortCreatedAt:
		return m.CreatedAt
	case graphql_models.AdditiveSortUpdatedAt:
		return m.UpdatedAt
	case graphql_models.AdditiveSortDeletedAt:
		return m.DeletedAt
	}
	return nil
}

func AdditiveSortDirection(ordering []*graphql_models.AdditiveOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func FromAdditiveCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	var columns []string
	var values []interface{}

	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := AdditiveSortValueFromCursorValue(cursorValue)
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

func ToAdditiveCursor(ordering []*graphql_models.AdditiveOrdering, m *graphql_models.Additive) string {
	var a []string
	var handledID bool

	for _, order := range ordering {
		if order.Sort == graphql_models.AdditiveSortID {
			handledID = true
		}
		value := AdditiveSortCursorValue(order.Sort, m)
		if value != nil {
			a = append(a, boilergql.ToCursorValue(string(order.Sort), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(graphql_models.AdditiveSortID), m.ID))
	}

	return boilergql.CursorValuesToString(a)
}

func AdditiveCursorType(ordering []*graphql_models.AdditiveOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func AdditiveCursorMods(ordering []*graphql_models.AdditiveOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if AdditiveCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromAdditiveCursor(*cursor, sign)
		}
		return boilergql.FromOffsetCursor(*cursor)
	}
	return nil
}

func AdditiveSortMods(ordering []*graphql_models.AdditiveOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, order := range ordering {
		if order.Sort == graphql_models.AdditiveSortID {
			handledID = true
		}
		column := AdditiveSortColumn[order.Sort]
		if column != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				column,
				boilergql.GetDirection(order.Direction, reverse),
			)))
		}
	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			models.AdditiveColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func AdditivePaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*graphql_models.AdditiveOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := AdditiveSortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, AdditiveCursorMods(ordering, cursor, sign)...)
	mods = append(mods, AdditiveSortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func AdditivePaginationMods(pagination boilergql.ConnectionPagination, ordering []*graphql_models.AdditiveOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	reverse := pagination.Backward != nil
	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	_, mods := AdditivePaginationModsBase(pagination, ordering, reverse, limit)
	return mods, nil
}

func ToAdditiveCursorSwitch(ordering []*graphql_models.AdditiveOrdering, m *graphql_models.Additive, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToAdditiveCursor(ordering, m)
	}
	return ""
}

func AdditiveReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.AdditiveOrdering,
) (bool, error) {
	reverse := pagination.Forward != nil
	cursor, reverseMods := AdditivePaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := AdditiveCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return models.Additives(reverseMods...).Count(ctx, db)
	})
}

func AdditiveEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*graphql_models.AdditiveOrdering) func(*models.Additive, int) *graphql_models.AdditiveEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), AdditiveCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *models.Additive, i int) *graphql_models.AdditiveEdge {
		n := AdditiveToGraphQL(m)
		return &graphql_models.AdditiveEdge{
			Cursor:	ToAdditiveCursorSwitch(ordering, n, cursorType, offset, i),
			Node:	n,
		}
	}
}

func AdditiveStartEndCursor(edges []*graphql_models.AdditiveEdge) (*string, *string) {
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

func AdditiveConnection(
	ctx context.Context,
	db *sql.DB,
	originalMods []qm.QueryMod,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.AdditiveOrdering,
) (*graphql_models.AdditiveConnection, error) {
	paginationMods, err := AdditivePaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := AdditiveReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := models.Additives(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*graphql_models.AdditiveEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := AdditiveEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := AdditiveStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
	return &graphql_models.AdditiveConnection{
		Edges:	edges,
		PageInfo: &graphql_models.PageInfo{
			HasNextPage:		hasNextPage,
			HasPreviousPage:	hasPreviousPage,
			StartCursor:		startCursor,
			EndCursor:		endCursor,
		},
	}, nil
}

var AuthGroupSortColumn = map[graphql_models.AuthGroupSort]string{
	graphql_models.AuthGroupSortID:		models.AuthGroupColumns.ID,
	graphql_models.AuthGroupSortName:	models.AuthGroupColumns.Name,
}

func AuthGroupSortValueFromCursorValue(cursorValue string) (string, interface{}) {
	key, value := boilergql.FromCursorValue(cursorValue)
	column := AuthGroupSortColumn[graphql_models.AuthGroupSort(key)]

	if graphql_models.AuthGroupSort(key) == graphql_models.AuthGroupSortID {
		return column, boilergql.GetIDFromCursor(value)
	}

	return column, boilergql.StringToInterface(value)
}

func AuthGroupSortCursorValue(sort graphql_models.AuthGroupSort, m *graphql_models.AuthGroup) interface{} {
	switch sort {
	case graphql_models.AuthGroupSortID:
		return m.ID
	case graphql_models.AuthGroupSortName:
		return m.Name
	}
	return nil
}

func AuthGroupSortDirection(ordering []*graphql_models.AuthGroupOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func FromAuthGroupCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	var columns []string
	var values []interface{}

	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := AuthGroupSortValueFromCursorValue(cursorValue)
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

func ToAuthGroupCursor(ordering []*graphql_models.AuthGroupOrdering, m *graphql_models.AuthGroup) string {
	var a []string
	var handledID bool

	for _, order := range ordering {
		if order.Sort == graphql_models.AuthGroupSortID {
			handledID = true
		}
		value := AuthGroupSortCursorValue(order.Sort, m)
		if value != nil {
			a = append(a, boilergql.ToCursorValue(string(order.Sort), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(graphql_models.AuthGroupSortID), m.ID))
	}

	return boilergql.CursorValuesToString(a)
}

func AuthGroupCursorType(ordering []*graphql_models.AuthGroupOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func AuthGroupCursorMods(ordering []*graphql_models.AuthGroupOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if AuthGroupCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromAuthGroupCursor(*cursor, sign)
		}
		return boilergql.FromOffsetCursor(*cursor)
	}
	return nil
}

func AuthGroupSortMods(ordering []*graphql_models.AuthGroupOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, order := range ordering {
		if order.Sort == graphql_models.AuthGroupSortID {
			handledID = true
		}
		column := AuthGroupSortColumn[order.Sort]
		if column != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				column,
				boilergql.GetDirection(order.Direction, reverse),
			)))
		}
	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			models.AuthGroupColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func AuthGroupPaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*graphql_models.AuthGroupOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := AuthGroupSortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, AuthGroupCursorMods(ordering, cursor, sign)...)
	mods = append(mods, AuthGroupSortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func AuthGroupPaginationMods(pagination boilergql.ConnectionPagination, ordering []*graphql_models.AuthGroupOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	reverse := pagination.Backward != nil
	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	_, mods := AuthGroupPaginationModsBase(pagination, ordering, reverse, limit)
	return mods, nil
}

func ToAuthGroupCursorSwitch(ordering []*graphql_models.AuthGroupOrdering, m *graphql_models.AuthGroup, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToAuthGroupCursor(ordering, m)
	}
	return ""
}

func AuthGroupReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.AuthGroupOrdering,
) (bool, error) {
	reverse := pagination.Forward != nil
	cursor, reverseMods := AuthGroupPaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := AuthGroupCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return models.AuthGroups(reverseMods...).Count(ctx, db)
	})
}

func AuthGroupEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*graphql_models.AuthGroupOrdering) func(*models.AuthGroup, int) *graphql_models.AuthGroupEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), AuthGroupCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *models.AuthGroup, i int) *graphql_models.AuthGroupEdge {
		n := AuthGroupToGraphQL(m)
		return &graphql_models.AuthGroupEdge{
			Cursor:	ToAuthGroupCursorSwitch(ordering, n, cursorType, offset, i),
			Node:	n,
		}
	}
}

func AuthGroupStartEndCursor(edges []*graphql_models.AuthGroupEdge) (*string, *string) {
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

func AuthGroupConnection(
	ctx context.Context,
	db *sql.DB,
	originalMods []qm.QueryMod,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.AuthGroupOrdering,
) (*graphql_models.AuthGroupConnection, error) {
	paginationMods, err := AuthGroupPaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := AuthGroupReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := models.AuthGroups(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*graphql_models.AuthGroupEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := AuthGroupEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := AuthGroupStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
	return &graphql_models.AuthGroupConnection{
		Edges:	edges,
		PageInfo: &graphql_models.PageInfo{
			HasNextPage:		hasNextPage,
			HasPreviousPage:	hasPreviousPage,
			StartCursor:		startCursor,
			EndCursor:		endCursor,
		},
	}, nil
}

var AuthGroupPermissionSortColumn = map[graphql_models.AuthGroupPermissionSort]string{
	graphql_models.AuthGroupPermissionSortID: models.AuthGroupPermissionColumns.ID,
}

func AuthGroupPermissionSortValueFromCursorValue(cursorValue string) (string, interface{}) {
	key, value := boilergql.FromCursorValue(cursorValue)
	column := AuthGroupPermissionSortColumn[graphql_models.AuthGroupPermissionSort(key)]

	if graphql_models.AuthGroupPermissionSort(key) == graphql_models.AuthGroupPermissionSortID {
		return column, boilergql.GetIDFromCursor(value)
	}

	return column, boilergql.StringToInterface(value)
}

func AuthGroupPermissionSortCursorValue(sort graphql_models.AuthGroupPermissionSort, m *graphql_models.AuthGroupPermission) interface{} {
	switch sort {
	case graphql_models.AuthGroupPermissionSortID:
		return m.ID
	}
	return nil
}

func AuthGroupPermissionSortDirection(ordering []*graphql_models.AuthGroupPermissionOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func FromAuthGroupPermissionCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	var columns []string
	var values []interface{}

	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := AuthGroupPermissionSortValueFromCursorValue(cursorValue)
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

func ToAuthGroupPermissionCursor(ordering []*graphql_models.AuthGroupPermissionOrdering, m *graphql_models.AuthGroupPermission) string {
	var a []string
	var handledID bool

	for _, order := range ordering {
		if order.Sort == graphql_models.AuthGroupPermissionSortID {
			handledID = true
		}
		value := AuthGroupPermissionSortCursorValue(order.Sort, m)
		if value != nil {
			a = append(a, boilergql.ToCursorValue(string(order.Sort), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(graphql_models.AuthGroupPermissionSortID), m.ID))
	}
	return boilergql.CursorValuesToString(a)
}

func AuthGroupPermissionCursorType(ordering []*graphql_models.AuthGroupPermissionOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func AuthGroupPermissionCursorMods(ordering []*graphql_models.AuthGroupPermissionOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if AuthGroupPermissionCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromAuthGroupPermissionCursor(*cursor, sign)
		}
		return boilergql.FromOffsetCursor(*cursor)
	}
	return nil
}

func AuthGroupPermissionSortMods(ordering []*graphql_models.AuthGroupPermissionOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, order := range ordering {
		if order.Sort == graphql_models.AuthGroupPermissionSortID {
			handledID = true
		}
		column := AuthGroupPermissionSortColumn[order.Sort]
		if column != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				column,
				boilergql.GetDirection(order.Direction, reverse),
			)))
		}
	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			models.AuthGroupPermissionColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func AuthGroupPermissionPaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*graphql_models.AuthGroupPermissionOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := AuthGroupPermissionSortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, AuthGroupPermissionCursorMods(ordering, cursor, sign)...)
	mods = append(mods, AuthGroupPermissionSortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func AuthGroupPermissionPaginationMods(pagination boilergql.ConnectionPagination, ordering []*graphql_models.AuthGroupPermissionOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	reverse := pagination.Backward != nil
	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	_, mods := AuthGroupPermissionPaginationModsBase(pagination, ordering, reverse, limit)
	return mods, nil
}

func ToAuthGroupPermissionCursorSwitch(ordering []*graphql_models.AuthGroupPermissionOrdering, m *graphql_models.AuthGroupPermission, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToAuthGroupPermissionCursor(ordering, m)
	}
	return ""
}

func AuthGroupPermissionReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.AuthGroupPermissionOrdering,
) (bool, error) {
	reverse := pagination.Forward != nil
	cursor, reverseMods := AuthGroupPermissionPaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := AuthGroupPermissionCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return models.AuthGroupPermissions(reverseMods...).Count(ctx, db)
	})
}

func AuthGroupPermissionEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*graphql_models.AuthGroupPermissionOrdering) func(*models.AuthGroupPermission, int) *graphql_models.AuthGroupPermissionEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), AuthGroupPermissionCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *models.AuthGroupPermission, i int) *graphql_models.AuthGroupPermissionEdge {
		n := AuthGroupPermissionToGraphQL(m)
		return &graphql_models.AuthGroupPermissionEdge{
			Cursor:	ToAuthGroupPermissionCursorSwitch(ordering, n, cursorType, offset, i),
			Node:	n,
		}
	}
}

func AuthGroupPermissionStartEndCursor(edges []*graphql_models.AuthGroupPermissionEdge) (*string, *string) {
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

func AuthGroupPermissionConnection(
	ctx context.Context,
	db *sql.DB,
	originalMods []qm.QueryMod,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.AuthGroupPermissionOrdering,
) (*graphql_models.AuthGroupPermissionConnection, error) {
	paginationMods, err := AuthGroupPermissionPaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := AuthGroupPermissionReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := models.AuthGroupPermissions(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*graphql_models.AuthGroupPermissionEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := AuthGroupPermissionEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := AuthGroupPermissionStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
	return &graphql_models.AuthGroupPermissionConnection{
		Edges:	edges,
		PageInfo: &graphql_models.PageInfo{
			HasNextPage:		hasNextPage,
			HasPreviousPage:	hasPreviousPage,
			StartCursor:		startCursor,
			EndCursor:		endCursor,
		},
	}, nil
}

var AuthPermissionSortColumn = map[graphql_models.AuthPermissionSort]string{
	graphql_models.AuthPermissionSortID:		models.AuthPermissionColumns.ID,
	graphql_models.AuthPermissionSortName:		models.AuthPermissionColumns.Name,
	graphql_models.AuthPermissionSortCodename:	models.AuthPermissionColumns.Codename,
}

func AuthPermissionSortValueFromCursorValue(cursorValue string) (string, interface{}) {
	key, value := boilergql.FromCursorValue(cursorValue)
	column := AuthPermissionSortColumn[graphql_models.AuthPermissionSort(key)]

	if graphql_models.AuthPermissionSort(key) == graphql_models.AuthPermissionSortID {
		return column, boilergql.GetIDFromCursor(value)
	}

	return column, boilergql.StringToInterface(value)
}

func AuthPermissionSortCursorValue(sort graphql_models.AuthPermissionSort, m *graphql_models.AuthPermission) interface{} {
	switch sort {
	case graphql_models.AuthPermissionSortID:
		return m.ID
	case graphql_models.AuthPermissionSortName:
		return m.Name
	case graphql_models.AuthPermissionSortCodename:
		return m.Codename
	}
	return nil
}

func AuthPermissionSortDirection(ordering []*graphql_models.AuthPermissionOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func FromAuthPermissionCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	var columns []string
	var values []interface{}

	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := AuthPermissionSortValueFromCursorValue(cursorValue)
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

func ToAuthPermissionCursor(ordering []*graphql_models.AuthPermissionOrdering, m *graphql_models.AuthPermission) string {
	var a []string
	var handledID bool

	for _, order := range ordering {
		if order.Sort == graphql_models.AuthPermissionSortID {
			handledID = true
		}
		value := AuthPermissionSortCursorValue(order.Sort, m)
		if value != nil {
			a = append(a, boilergql.ToCursorValue(string(order.Sort), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(graphql_models.AuthPermissionSortID), m.ID))
	}

	return boilergql.CursorValuesToString(a)
}

func AuthPermissionCursorType(ordering []*graphql_models.AuthPermissionOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func AuthPermissionCursorMods(ordering []*graphql_models.AuthPermissionOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if AuthPermissionCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromAuthPermissionCursor(*cursor, sign)
		}
		return boilergql.FromOffsetCursor(*cursor)
	}
	return nil
}

func AuthPermissionSortMods(ordering []*graphql_models.AuthPermissionOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, order := range ordering {
		if order.Sort == graphql_models.AuthPermissionSortID {
			handledID = true
		}
		column := AuthPermissionSortColumn[order.Sort]
		if column != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				column,
				boilergql.GetDirection(order.Direction, reverse),
			)))
		}
	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			models.AuthPermissionColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func AuthPermissionPaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*graphql_models.AuthPermissionOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := AuthPermissionSortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, AuthPermissionCursorMods(ordering, cursor, sign)...)
	mods = append(mods, AuthPermissionSortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func AuthPermissionPaginationMods(pagination boilergql.ConnectionPagination, ordering []*graphql_models.AuthPermissionOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	reverse := pagination.Backward != nil
	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	_, mods := AuthPermissionPaginationModsBase(pagination, ordering, reverse, limit)
	return mods, nil
}

func ToAuthPermissionCursorSwitch(ordering []*graphql_models.AuthPermissionOrdering, m *graphql_models.AuthPermission, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToAuthPermissionCursor(ordering, m)
	}
	return ""
}

func AuthPermissionReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.AuthPermissionOrdering,
) (bool, error) {
	reverse := pagination.Forward != nil
	cursor, reverseMods := AuthPermissionPaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := AuthPermissionCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return models.AuthPermissions(reverseMods...).Count(ctx, db)
	})
}

func AuthPermissionEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*graphql_models.AuthPermissionOrdering) func(*models.AuthPermission, int) *graphql_models.AuthPermissionEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), AuthPermissionCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *models.AuthPermission, i int) *graphql_models.AuthPermissionEdge {
		n := AuthPermissionToGraphQL(m)
		return &graphql_models.AuthPermissionEdge{
			Cursor:	ToAuthPermissionCursorSwitch(ordering, n, cursorType, offset, i),
			Node:	n,
		}
	}
}

func AuthPermissionStartEndCursor(edges []*graphql_models.AuthPermissionEdge) (*string, *string) {
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

func AuthPermissionConnection(
	ctx context.Context,
	db *sql.DB,
	originalMods []qm.QueryMod,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.AuthPermissionOrdering,
) (*graphql_models.AuthPermissionConnection, error) {
	paginationMods, err := AuthPermissionPaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := AuthPermissionReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := models.AuthPermissions(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*graphql_models.AuthPermissionEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := AuthPermissionEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := AuthPermissionStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
	return &graphql_models.AuthPermissionConnection{
		Edges:	edges,
		PageInfo: &graphql_models.PageInfo{
			HasNextPage:		hasNextPage,
			HasPreviousPage:	hasPreviousPage,
			StartCursor:		startCursor,
			EndCursor:		endCursor,
		},
	}, nil
}

var AuthUserGroupSortColumn = map[graphql_models.AuthUserGroupSort]string{
	graphql_models.AuthUserGroupSortID: models.AuthUserGroupColumns.ID,
}

func AuthUserGroupSortValueFromCursorValue(cursorValue string) (string, interface{}) {
	key, value := boilergql.FromCursorValue(cursorValue)
	column := AuthUserGroupSortColumn[graphql_models.AuthUserGroupSort(key)]

	if graphql_models.AuthUserGroupSort(key) == graphql_models.AuthUserGroupSortID {
		return column, boilergql.GetIDFromCursor(value)
	}

	return column, boilergql.StringToInterface(value)
}

func AuthUserGroupSortCursorValue(sort graphql_models.AuthUserGroupSort, m *graphql_models.AuthUserGroup) interface{} {
	switch sort {
	case graphql_models.AuthUserGroupSortID:
		return m.ID
	}
	return nil
}

func AuthUserGroupSortDirection(ordering []*graphql_models.AuthUserGroupOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func FromAuthUserGroupCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	var columns []string
	var values []interface{}

	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := AuthUserGroupSortValueFromCursorValue(cursorValue)
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

func ToAuthUserGroupCursor(ordering []*graphql_models.AuthUserGroupOrdering, m *graphql_models.AuthUserGroup) string {
	var a []string
	var handledID bool

	for _, order := range ordering {
		if order.Sort == graphql_models.AuthUserGroupSortID {
			handledID = true
		}
		value := AuthUserGroupSortCursorValue(order.Sort, m)
		if value != nil {
			a = append(a, boilergql.ToCursorValue(string(order.Sort), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(graphql_models.AuthUserGroupSortID), m.ID))
	}
	return boilergql.CursorValuesToString(a)
}

func AuthUserGroupCursorType(ordering []*graphql_models.AuthUserGroupOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func AuthUserGroupCursorMods(ordering []*graphql_models.AuthUserGroupOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if AuthUserGroupCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromAuthUserGroupCursor(*cursor, sign)
		}
		return boilergql.FromOffsetCursor(*cursor)
	}
	return nil
}

func AuthUserGroupSortMods(ordering []*graphql_models.AuthUserGroupOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, order := range ordering {
		if order.Sort == graphql_models.AuthUserGroupSortID {
			handledID = true
		}
		column := AuthUserGroupSortColumn[order.Sort]
		if column != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				column,
				boilergql.GetDirection(order.Direction, reverse),
			)))
		}
	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			models.AuthUserGroupColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func AuthUserGroupPaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*graphql_models.AuthUserGroupOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := AuthUserGroupSortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, AuthUserGroupCursorMods(ordering, cursor, sign)...)
	mods = append(mods, AuthUserGroupSortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func AuthUserGroupPaginationMods(pagination boilergql.ConnectionPagination, ordering []*graphql_models.AuthUserGroupOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	reverse := pagination.Backward != nil
	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	_, mods := AuthUserGroupPaginationModsBase(pagination, ordering, reverse, limit)
	return mods, nil
}

func ToAuthUserGroupCursorSwitch(ordering []*graphql_models.AuthUserGroupOrdering, m *graphql_models.AuthUserGroup, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToAuthUserGroupCursor(ordering, m)
	}
	return ""
}

func AuthUserGroupReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.AuthUserGroupOrdering,
) (bool, error) {
	reverse := pagination.Forward != nil
	cursor, reverseMods := AuthUserGroupPaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := AuthUserGroupCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return models.AuthUserGroups(reverseMods...).Count(ctx, db)
	})
}

func AuthUserGroupEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*graphql_models.AuthUserGroupOrdering) func(*models.AuthUserGroup, int) *graphql_models.AuthUserGroupEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), AuthUserGroupCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *models.AuthUserGroup, i int) *graphql_models.AuthUserGroupEdge {
		n := AuthUserGroupToGraphQL(m)
		return &graphql_models.AuthUserGroupEdge{
			Cursor:	ToAuthUserGroupCursorSwitch(ordering, n, cursorType, offset, i),
			Node:	n,
		}
	}
}

func AuthUserGroupStartEndCursor(edges []*graphql_models.AuthUserGroupEdge) (*string, *string) {
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

func AuthUserGroupConnection(
	ctx context.Context,
	db *sql.DB,
	originalMods []qm.QueryMod,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.AuthUserGroupOrdering,
) (*graphql_models.AuthUserGroupConnection, error) {
	paginationMods, err := AuthUserGroupPaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := AuthUserGroupReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := models.AuthUserGroups(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*graphql_models.AuthUserGroupEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := AuthUserGroupEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := AuthUserGroupStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
	return &graphql_models.AuthUserGroupConnection{
		Edges:	edges,
		PageInfo: &graphql_models.PageInfo{
			HasNextPage:		hasNextPage,
			HasPreviousPage:	hasPreviousPage,
			StartCursor:		startCursor,
			EndCursor:		endCursor,
		},
	}, nil
}

var AuthUserSortColumn = map[graphql_models.AuthUserSort]string{
	graphql_models.AuthUserSortID:		models.AuthUserColumns.ID,
	graphql_models.AuthUserSortPassword:	models.AuthUserColumns.Password,
	graphql_models.AuthUserSortLastLogin:	models.AuthUserColumns.LastLogin,
	graphql_models.AuthUserSortIsSuperuser:	models.AuthUserColumns.IsSuperuser,
	graphql_models.AuthUserSortUsername:	models.AuthUserColumns.Username,
	graphql_models.AuthUserSortFirstName:	models.AuthUserColumns.FirstName,
	graphql_models.AuthUserSortLastName:	models.AuthUserColumns.LastName,
	graphql_models.AuthUserSortEmail:	models.AuthUserColumns.Email,
	graphql_models.AuthUserSortIsStaff:	models.AuthUserColumns.IsStaff,
	graphql_models.AuthUserSortIsActive:	models.AuthUserColumns.IsActive,
	graphql_models.AuthUserSortDateJoined:	models.AuthUserColumns.DateJoined,
}

func AuthUserSortValueFromCursorValue(cursorValue string) (string, interface{}) {
	key, value := boilergql.FromCursorValue(cursorValue)
	column := AuthUserSortColumn[graphql_models.AuthUserSort(key)]

	if graphql_models.AuthUserSort(key) == graphql_models.AuthUserSortID {
		return column, boilergql.GetIDFromCursor(value)
	}

	return column, boilergql.StringToInterface(value)
}

func AuthUserSortCursorValue(sort graphql_models.AuthUserSort, m *graphql_models.AuthUser) interface{} {
	switch sort {
	case graphql_models.AuthUserSortID:
		return m.ID
	case graphql_models.AuthUserSortPassword:
		return m.Password
	case graphql_models.AuthUserSortLastLogin:
		return m.LastLogin
	case graphql_models.AuthUserSortIsSuperuser:
		return m.IsSuperuser
	case graphql_models.AuthUserSortUsername:
		return m.Username
	case graphql_models.AuthUserSortFirstName:
		return m.FirstName
	case graphql_models.AuthUserSortLastName:
		return m.LastName
	case graphql_models.AuthUserSortEmail:
		return m.Email
	case graphql_models.AuthUserSortIsStaff:
		return m.IsStaff
	case graphql_models.AuthUserSortIsActive:
		return m.IsActive
	case graphql_models.AuthUserSortDateJoined:
		return m.DateJoined
	}
	return nil
}

func AuthUserSortDirection(ordering []*graphql_models.AuthUserOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func FromAuthUserCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	var columns []string
	var values []interface{}

	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := AuthUserSortValueFromCursorValue(cursorValue)
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

func ToAuthUserCursor(ordering []*graphql_models.AuthUserOrdering, m *graphql_models.AuthUser) string {
	var a []string
	var handledID bool

	for _, order := range ordering {
		if order.Sort == graphql_models.AuthUserSortID {
			handledID = true
		}
		value := AuthUserSortCursorValue(order.Sort, m)
		if value != nil {
			a = append(a, boilergql.ToCursorValue(string(order.Sort), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(graphql_models.AuthUserSortID), m.ID))
	}

	return boilergql.CursorValuesToString(a)
}

func AuthUserCursorType(ordering []*graphql_models.AuthUserOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func AuthUserCursorMods(ordering []*graphql_models.AuthUserOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if AuthUserCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromAuthUserCursor(*cursor, sign)
		}
		return boilergql.FromOffsetCursor(*cursor)
	}
	return nil
}

func AuthUserSortMods(ordering []*graphql_models.AuthUserOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, order := range ordering {
		if order.Sort == graphql_models.AuthUserSortID {
			handledID = true
		}
		column := AuthUserSortColumn[order.Sort]
		if column != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				column,
				boilergql.GetDirection(order.Direction, reverse),
			)))
		}
	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			models.AuthUserColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func AuthUserPaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*graphql_models.AuthUserOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := AuthUserSortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, AuthUserCursorMods(ordering, cursor, sign)...)
	mods = append(mods, AuthUserSortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func AuthUserPaginationMods(pagination boilergql.ConnectionPagination, ordering []*graphql_models.AuthUserOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	reverse := pagination.Backward != nil
	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	_, mods := AuthUserPaginationModsBase(pagination, ordering, reverse, limit)
	return mods, nil
}

func ToAuthUserCursorSwitch(ordering []*graphql_models.AuthUserOrdering, m *graphql_models.AuthUser, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToAuthUserCursor(ordering, m)
	}
	return ""
}

func AuthUserReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.AuthUserOrdering,
) (bool, error) {
	reverse := pagination.Forward != nil
	cursor, reverseMods := AuthUserPaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := AuthUserCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return models.AuthUsers(reverseMods...).Count(ctx, db)
	})
}

func AuthUserEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*graphql_models.AuthUserOrdering) func(*models.AuthUser, int) *graphql_models.AuthUserEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), AuthUserCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *models.AuthUser, i int) *graphql_models.AuthUserEdge {
		n := AuthUserToGraphQL(m)
		return &graphql_models.AuthUserEdge{
			Cursor:	ToAuthUserCursorSwitch(ordering, n, cursorType, offset, i),
			Node:	n,
		}
	}
}

func AuthUserStartEndCursor(edges []*graphql_models.AuthUserEdge) (*string, *string) {
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

func AuthUserConnection(
	ctx context.Context,
	db *sql.DB,
	originalMods []qm.QueryMod,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.AuthUserOrdering,
) (*graphql_models.AuthUserConnection, error) {
	paginationMods, err := AuthUserPaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := AuthUserReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := models.AuthUsers(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*graphql_models.AuthUserEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := AuthUserEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := AuthUserStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
	return &graphql_models.AuthUserConnection{
		Edges:	edges,
		PageInfo: &graphql_models.PageInfo{
			HasNextPage:		hasNextPage,
			HasPreviousPage:	hasPreviousPage,
			StartCursor:		startCursor,
			EndCursor:		endCursor,
		},
	}, nil
}

var AuthUserUserPermissionSortColumn = map[graphql_models.AuthUserUserPermissionSort]string{
	graphql_models.AuthUserUserPermissionSortID: models.AuthUserUserPermissionColumns.ID,
}

func AuthUserUserPermissionSortValueFromCursorValue(cursorValue string) (string, interface{}) {
	key, value := boilergql.FromCursorValue(cursorValue)
	column := AuthUserUserPermissionSortColumn[graphql_models.AuthUserUserPermissionSort(key)]

	if graphql_models.AuthUserUserPermissionSort(key) == graphql_models.AuthUserUserPermissionSortID {
		return column, boilergql.GetIDFromCursor(value)
	}

	return column, boilergql.StringToInterface(value)
}

func AuthUserUserPermissionSortCursorValue(sort graphql_models.AuthUserUserPermissionSort, m *graphql_models.AuthUserUserPermission) interface{} {
	switch sort {
	case graphql_models.AuthUserUserPermissionSortID:
		return m.ID
	}
	return nil
}

func AuthUserUserPermissionSortDirection(ordering []*graphql_models.AuthUserUserPermissionOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func FromAuthUserUserPermissionCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	var columns []string
	var values []interface{}

	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := AuthUserUserPermissionSortValueFromCursorValue(cursorValue)
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

func ToAuthUserUserPermissionCursor(ordering []*graphql_models.AuthUserUserPermissionOrdering, m *graphql_models.AuthUserUserPermission) string {
	var a []string
	var handledID bool

	for _, order := range ordering {
		if order.Sort == graphql_models.AuthUserUserPermissionSortID {
			handledID = true
		}
		value := AuthUserUserPermissionSortCursorValue(order.Sort, m)
		if value != nil {
			a = append(a, boilergql.ToCursorValue(string(order.Sort), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(graphql_models.AuthUserUserPermissionSortID), m.ID))
	}
	return boilergql.CursorValuesToString(a)
}

func AuthUserUserPermissionCursorType(ordering []*graphql_models.AuthUserUserPermissionOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func AuthUserUserPermissionCursorMods(ordering []*graphql_models.AuthUserUserPermissionOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if AuthUserUserPermissionCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromAuthUserUserPermissionCursor(*cursor, sign)
		}
		return boilergql.FromOffsetCursor(*cursor)
	}
	return nil
}

func AuthUserUserPermissionSortMods(ordering []*graphql_models.AuthUserUserPermissionOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, order := range ordering {
		if order.Sort == graphql_models.AuthUserUserPermissionSortID {
			handledID = true
		}
		column := AuthUserUserPermissionSortColumn[order.Sort]
		if column != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				column,
				boilergql.GetDirection(order.Direction, reverse),
			)))
		}
	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			models.AuthUserUserPermissionColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func AuthUserUserPermissionPaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*graphql_models.AuthUserUserPermissionOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := AuthUserUserPermissionSortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, AuthUserUserPermissionCursorMods(ordering, cursor, sign)...)
	mods = append(mods, AuthUserUserPermissionSortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func AuthUserUserPermissionPaginationMods(pagination boilergql.ConnectionPagination, ordering []*graphql_models.AuthUserUserPermissionOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	reverse := pagination.Backward != nil
	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	_, mods := AuthUserUserPermissionPaginationModsBase(pagination, ordering, reverse, limit)
	return mods, nil
}

func ToAuthUserUserPermissionCursorSwitch(ordering []*graphql_models.AuthUserUserPermissionOrdering, m *graphql_models.AuthUserUserPermission, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToAuthUserUserPermissionCursor(ordering, m)
	}
	return ""
}

func AuthUserUserPermissionReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.AuthUserUserPermissionOrdering,
) (bool, error) {
	reverse := pagination.Forward != nil
	cursor, reverseMods := AuthUserUserPermissionPaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := AuthUserUserPermissionCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return models.AuthUserUserPermissions(reverseMods...).Count(ctx, db)
	})
}

func AuthUserUserPermissionEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*graphql_models.AuthUserUserPermissionOrdering) func(*models.AuthUserUserPermission, int) *graphql_models.AuthUserUserPermissionEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), AuthUserUserPermissionCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *models.AuthUserUserPermission, i int) *graphql_models.AuthUserUserPermissionEdge {
		n := AuthUserUserPermissionToGraphQL(m)
		return &graphql_models.AuthUserUserPermissionEdge{
			Cursor:	ToAuthUserUserPermissionCursorSwitch(ordering, n, cursorType, offset, i),
			Node:	n,
		}
	}
}

func AuthUserUserPermissionStartEndCursor(edges []*graphql_models.AuthUserUserPermissionEdge) (*string, *string) {
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

func AuthUserUserPermissionConnection(
	ctx context.Context,
	db *sql.DB,
	originalMods []qm.QueryMod,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.AuthUserUserPermissionOrdering,
) (*graphql_models.AuthUserUserPermissionConnection, error) {
	paginationMods, err := AuthUserUserPermissionPaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := AuthUserUserPermissionReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := models.AuthUserUserPermissions(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*graphql_models.AuthUserUserPermissionEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := AuthUserUserPermissionEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := AuthUserUserPermissionStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
	return &graphql_models.AuthUserUserPermissionConnection{
		Edges:	edges,
		PageInfo: &graphql_models.PageInfo{
			HasNextPage:		hasNextPage,
			HasPreviousPage:	hasPreviousPage,
			StartCursor:		startCursor,
			EndCursor:		endCursor,
		},
	}, nil
}

var FragranceInventorySortColumn = map[graphql_models.FragranceInventorySort]string{
	graphql_models.FragranceInventorySortID:		models.FragranceInventoryColumns.ID,
	graphql_models.FragranceInventorySortPurchaseDate:	models.FragranceInventoryColumns.PurchaseDate,
	graphql_models.FragranceInventorySortExpiryDate:	models.FragranceInventoryColumns.ExpiryDate,
	graphql_models.FragranceInventorySortCost:		models.FragranceInventoryColumns.Cost,
	graphql_models.FragranceInventorySortWeight:		models.FragranceInventoryColumns.Weight,
	graphql_models.FragranceInventorySortCreatedAt:		models.FragranceInventoryColumns.CreatedAt,
	graphql_models.FragranceInventorySortUpdatedAt:		models.FragranceInventoryColumns.UpdatedAt,
	graphql_models.FragranceInventorySortDeletedAt:		models.FragranceInventoryColumns.DeletedAt,
}

func FragranceInventorySortValueFromCursorValue(cursorValue string) (string, interface{}) {
	key, value := boilergql.FromCursorValue(cursorValue)
	column := FragranceInventorySortColumn[graphql_models.FragranceInventorySort(key)]

	if graphql_models.FragranceInventorySort(key) == graphql_models.FragranceInventorySortID {
		return column, boilergql.GetIDFromCursor(value)
	}

	return column, boilergql.StringToInterface(value)
}

func FragranceInventorySortCursorValue(sort graphql_models.FragranceInventorySort, m *graphql_models.FragranceInventory) interface{} {
	switch sort {
	case graphql_models.FragranceInventorySortID:
		return m.ID
	case graphql_models.FragranceInventorySortPurchaseDate:
		return m.PurchaseDate
	case graphql_models.FragranceInventorySortExpiryDate:
		return m.ExpiryDate
	case graphql_models.FragranceInventorySortCost:
		return m.Cost
	case graphql_models.FragranceInventorySortWeight:
		return m.Weight
	case graphql_models.FragranceInventorySortCreatedAt:
		return m.CreatedAt
	case graphql_models.FragranceInventorySortUpdatedAt:
		return m.UpdatedAt
	case graphql_models.FragranceInventorySortDeletedAt:
		return m.DeletedAt
	}
	return nil
}

func FragranceInventorySortDirection(ordering []*graphql_models.FragranceInventoryOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func FromFragranceInventoryCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	var columns []string
	var values []interface{}

	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := FragranceInventorySortValueFromCursorValue(cursorValue)
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

func ToFragranceInventoryCursor(ordering []*graphql_models.FragranceInventoryOrdering, m *graphql_models.FragranceInventory) string {
	var a []string
	var handledID bool

	for _, order := range ordering {
		if order.Sort == graphql_models.FragranceInventorySortID {
			handledID = true
		}
		value := FragranceInventorySortCursorValue(order.Sort, m)
		if value != nil {
			a = append(a, boilergql.ToCursorValue(string(order.Sort), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(graphql_models.FragranceInventorySortID), m.ID))
	}

	return boilergql.CursorValuesToString(a)
}

func FragranceInventoryCursorType(ordering []*graphql_models.FragranceInventoryOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func FragranceInventoryCursorMods(ordering []*graphql_models.FragranceInventoryOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if FragranceInventoryCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromFragranceInventoryCursor(*cursor, sign)
		}
		return boilergql.FromOffsetCursor(*cursor)
	}
	return nil
}

func FragranceInventorySortMods(ordering []*graphql_models.FragranceInventoryOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, order := range ordering {
		if order.Sort == graphql_models.FragranceInventorySortID {
			handledID = true
		}
		column := FragranceInventorySortColumn[order.Sort]
		if column != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				column,
				boilergql.GetDirection(order.Direction, reverse),
			)))
		}
	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			models.FragranceInventoryColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func FragranceInventoryPaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*graphql_models.FragranceInventoryOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := FragranceInventorySortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, FragranceInventoryCursorMods(ordering, cursor, sign)...)
	mods = append(mods, FragranceInventorySortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func FragranceInventoryPaginationMods(pagination boilergql.ConnectionPagination, ordering []*graphql_models.FragranceInventoryOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	reverse := pagination.Backward != nil
	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	_, mods := FragranceInventoryPaginationModsBase(pagination, ordering, reverse, limit)
	return mods, nil
}

func ToFragranceInventoryCursorSwitch(ordering []*graphql_models.FragranceInventoryOrdering, m *graphql_models.FragranceInventory, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToFragranceInventoryCursor(ordering, m)
	}
	return ""
}

func FragranceInventoryReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.FragranceInventoryOrdering,
) (bool, error) {
	reverse := pagination.Forward != nil
	cursor, reverseMods := FragranceInventoryPaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := FragranceInventoryCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return models.FragranceInventories(reverseMods...).Count(ctx, db)
	})
}

func FragranceInventoryEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*graphql_models.FragranceInventoryOrdering) func(*models.FragranceInventory, int) *graphql_models.FragranceInventoryEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), FragranceInventoryCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *models.FragranceInventory, i int) *graphql_models.FragranceInventoryEdge {
		n := FragranceInventoryToGraphQL(m)
		return &graphql_models.FragranceInventoryEdge{
			Cursor:	ToFragranceInventoryCursorSwitch(ordering, n, cursorType, offset, i),
			Node:	n,
		}
	}
}

func FragranceInventoryStartEndCursor(edges []*graphql_models.FragranceInventoryEdge) (*string, *string) {
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

func FragranceInventoryConnection(
	ctx context.Context,
	db *sql.DB,
	originalMods []qm.QueryMod,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.FragranceInventoryOrdering,
) (*graphql_models.FragranceInventoryConnection, error) {
	paginationMods, err := FragranceInventoryPaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := FragranceInventoryReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := models.FragranceInventories(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*graphql_models.FragranceInventoryEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := FragranceInventoryEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := FragranceInventoryStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
	return &graphql_models.FragranceInventoryConnection{
		Edges:	edges,
		PageInfo: &graphql_models.PageInfo{
			HasNextPage:		hasNextPage,
			HasPreviousPage:	hasPreviousPage,
			StartCursor:		startCursor,
			EndCursor:		endCursor,
		},
	}, nil
}

var FragranceSortColumn = map[graphql_models.FragranceSort]string{
	graphql_models.FragranceSortID:		models.FragranceColumns.ID,
	graphql_models.FragranceSortName:	models.FragranceColumns.Name,
	graphql_models.FragranceSortNote:	models.FragranceColumns.Note,
	graphql_models.FragranceSortCreatedAt:	models.FragranceColumns.CreatedAt,
	graphql_models.FragranceSortUpdatedAt:	models.FragranceColumns.UpdatedAt,
	graphql_models.FragranceSortDeletedAt:	models.FragranceColumns.DeletedAt,
}

func FragranceSortValueFromCursorValue(cursorValue string) (string, interface{}) {
	key, value := boilergql.FromCursorValue(cursorValue)
	column := FragranceSortColumn[graphql_models.FragranceSort(key)]

	if graphql_models.FragranceSort(key) == graphql_models.FragranceSortID {
		return column, boilergql.GetIDFromCursor(value)
	}

	return column, boilergql.StringToInterface(value)
}

func FragranceSortCursorValue(sort graphql_models.FragranceSort, m *graphql_models.Fragrance) interface{} {
	switch sort {
	case graphql_models.FragranceSortID:
		return m.ID
	case graphql_models.FragranceSortName:
		return m.Name
	case graphql_models.FragranceSortNote:
		return m.Note
	case graphql_models.FragranceSortCreatedAt:
		return m.CreatedAt
	case graphql_models.FragranceSortUpdatedAt:
		return m.UpdatedAt
	case graphql_models.FragranceSortDeletedAt:
		return m.DeletedAt
	}
	return nil
}

func FragranceSortDirection(ordering []*graphql_models.FragranceOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func FromFragranceCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	var columns []string
	var values []interface{}

	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := FragranceSortValueFromCursorValue(cursorValue)
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

func ToFragranceCursor(ordering []*graphql_models.FragranceOrdering, m *graphql_models.Fragrance) string {
	var a []string
	var handledID bool

	for _, order := range ordering {
		if order.Sort == graphql_models.FragranceSortID {
			handledID = true
		}
		value := FragranceSortCursorValue(order.Sort, m)
		if value != nil {
			a = append(a, boilergql.ToCursorValue(string(order.Sort), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(graphql_models.FragranceSortID), m.ID))
	}

	return boilergql.CursorValuesToString(a)
}

func FragranceCursorType(ordering []*graphql_models.FragranceOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func FragranceCursorMods(ordering []*graphql_models.FragranceOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if FragranceCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromFragranceCursor(*cursor, sign)
		}
		return boilergql.FromOffsetCursor(*cursor)
	}
	return nil
}

func FragranceSortMods(ordering []*graphql_models.FragranceOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, order := range ordering {
		if order.Sort == graphql_models.FragranceSortID {
			handledID = true
		}
		column := FragranceSortColumn[order.Sort]
		if column != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				column,
				boilergql.GetDirection(order.Direction, reverse),
			)))
		}
	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			models.FragranceColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func FragrancePaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*graphql_models.FragranceOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := FragranceSortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, FragranceCursorMods(ordering, cursor, sign)...)
	mods = append(mods, FragranceSortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func FragrancePaginationMods(pagination boilergql.ConnectionPagination, ordering []*graphql_models.FragranceOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	reverse := pagination.Backward != nil
	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	_, mods := FragrancePaginationModsBase(pagination, ordering, reverse, limit)
	return mods, nil
}

func ToFragranceCursorSwitch(ordering []*graphql_models.FragranceOrdering, m *graphql_models.Fragrance, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToFragranceCursor(ordering, m)
	}
	return ""
}

func FragranceReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.FragranceOrdering,
) (bool, error) {
	reverse := pagination.Forward != nil
	cursor, reverseMods := FragrancePaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := FragranceCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return models.Fragrances(reverseMods...).Count(ctx, db)
	})
}

func FragranceEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*graphql_models.FragranceOrdering) func(*models.Fragrance, int) *graphql_models.FragranceEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), FragranceCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *models.Fragrance, i int) *graphql_models.FragranceEdge {
		n := FragranceToGraphQL(m)
		return &graphql_models.FragranceEdge{
			Cursor:	ToFragranceCursorSwitch(ordering, n, cursorType, offset, i),
			Node:	n,
		}
	}
}

func FragranceStartEndCursor(edges []*graphql_models.FragranceEdge) (*string, *string) {
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

func FragranceConnection(
	ctx context.Context,
	db *sql.DB,
	originalMods []qm.QueryMod,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.FragranceOrdering,
) (*graphql_models.FragranceConnection, error) {
	paginationMods, err := FragrancePaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := FragranceReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := models.Fragrances(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*graphql_models.FragranceEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := FragranceEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := FragranceStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
	return &graphql_models.FragranceConnection{
		Edges:	edges,
		PageInfo: &graphql_models.PageInfo{
			HasNextPage:		hasNextPage,
			HasPreviousPage:	hasPreviousPage,
			StartCursor:		startCursor,
			EndCursor:		endCursor,
		},
	}, nil
}

var LipidInventorySortColumn = map[graphql_models.LipidInventorySort]string{
	graphql_models.LipidInventorySortID:		models.LipidInventoryColumns.ID,
	graphql_models.LipidInventorySortPurchaseDate:	models.LipidInventoryColumns.PurchaseDate,
	graphql_models.LipidInventorySortExpiryDate:	models.LipidInventoryColumns.ExpiryDate,
	graphql_models.LipidInventorySortCost:		models.LipidInventoryColumns.Cost,
	graphql_models.LipidInventorySortWeight:	models.LipidInventoryColumns.Weight,
	graphql_models.LipidInventorySortSap:		models.LipidInventoryColumns.Sap,
	graphql_models.LipidInventorySortNaoh:		models.LipidInventoryColumns.Naoh,
	graphql_models.LipidInventorySortKoh:		models.LipidInventoryColumns.Koh,
	graphql_models.LipidInventorySortGramsPerLiter:	models.LipidInventoryColumns.GramsPerLiter,
	graphql_models.LipidInventorySortCreatedAt:	models.LipidInventoryColumns.CreatedAt,
	graphql_models.LipidInventorySortUpdatedAt:	models.LipidInventoryColumns.UpdatedAt,
	graphql_models.LipidInventorySortDeletedAt:	models.LipidInventoryColumns.DeletedAt,
}

func LipidInventorySortValueFromCursorValue(cursorValue string) (string, interface{}) {
	key, value := boilergql.FromCursorValue(cursorValue)
	column := LipidInventorySortColumn[graphql_models.LipidInventorySort(key)]

	if graphql_models.LipidInventorySort(key) == graphql_models.LipidInventorySortID {
		return column, boilergql.GetIDFromCursor(value)
	}

	return column, boilergql.StringToInterface(value)
}

func LipidInventorySortCursorValue(sort graphql_models.LipidInventorySort, m *graphql_models.LipidInventory) interface{} {
	switch sort {
	case graphql_models.LipidInventorySortID:
		return m.ID
	case graphql_models.LipidInventorySortPurchaseDate:
		return m.PurchaseDate
	case graphql_models.LipidInventorySortExpiryDate:
		return m.ExpiryDate
	case graphql_models.LipidInventorySortCost:
		return m.Cost
	case graphql_models.LipidInventorySortWeight:
		return m.Weight
	case graphql_models.LipidInventorySortSap:
		return m.Sap
	case graphql_models.LipidInventorySortNaoh:
		return m.Naoh
	case graphql_models.LipidInventorySortKoh:
		return m.Koh
	case graphql_models.LipidInventorySortGramsPerLiter:
		return m.GramsPerLiter
	case graphql_models.LipidInventorySortCreatedAt:
		return m.CreatedAt
	case graphql_models.LipidInventorySortUpdatedAt:
		return m.UpdatedAt
	case graphql_models.LipidInventorySortDeletedAt:
		return m.DeletedAt
	}
	return nil
}

func LipidInventorySortDirection(ordering []*graphql_models.LipidInventoryOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func FromLipidInventoryCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	var columns []string
	var values []interface{}

	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := LipidInventorySortValueFromCursorValue(cursorValue)
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

func ToLipidInventoryCursor(ordering []*graphql_models.LipidInventoryOrdering, m *graphql_models.LipidInventory) string {
	var a []string
	var handledID bool

	for _, order := range ordering {
		if order.Sort == graphql_models.LipidInventorySortID {
			handledID = true
		}
		value := LipidInventorySortCursorValue(order.Sort, m)
		if value != nil {
			a = append(a, boilergql.ToCursorValue(string(order.Sort), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(graphql_models.LipidInventorySortID), m.ID))
	}

	return boilergql.CursorValuesToString(a)
}

func LipidInventoryCursorType(ordering []*graphql_models.LipidInventoryOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func LipidInventoryCursorMods(ordering []*graphql_models.LipidInventoryOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if LipidInventoryCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromLipidInventoryCursor(*cursor, sign)
		}
		return boilergql.FromOffsetCursor(*cursor)
	}
	return nil
}

func LipidInventorySortMods(ordering []*graphql_models.LipidInventoryOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, order := range ordering {
		if order.Sort == graphql_models.LipidInventorySortID {
			handledID = true
		}
		column := LipidInventorySortColumn[order.Sort]
		if column != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				column,
				boilergql.GetDirection(order.Direction, reverse),
			)))
		}
	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			models.LipidInventoryColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func LipidInventoryPaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*graphql_models.LipidInventoryOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := LipidInventorySortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, LipidInventoryCursorMods(ordering, cursor, sign)...)
	mods = append(mods, LipidInventorySortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func LipidInventoryPaginationMods(pagination boilergql.ConnectionPagination, ordering []*graphql_models.LipidInventoryOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	reverse := pagination.Backward != nil
	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	_, mods := LipidInventoryPaginationModsBase(pagination, ordering, reverse, limit)
	return mods, nil
}

func ToLipidInventoryCursorSwitch(ordering []*graphql_models.LipidInventoryOrdering, m *graphql_models.LipidInventory, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToLipidInventoryCursor(ordering, m)
	}
	return ""
}

func LipidInventoryReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.LipidInventoryOrdering,
) (bool, error) {
	reverse := pagination.Forward != nil
	cursor, reverseMods := LipidInventoryPaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := LipidInventoryCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return models.LipidInventories(reverseMods...).Count(ctx, db)
	})
}

func LipidInventoryEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*graphql_models.LipidInventoryOrdering) func(*models.LipidInventory, int) *graphql_models.LipidInventoryEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), LipidInventoryCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *models.LipidInventory, i int) *graphql_models.LipidInventoryEdge {
		n := LipidInventoryToGraphQL(m)
		return &graphql_models.LipidInventoryEdge{
			Cursor:	ToLipidInventoryCursorSwitch(ordering, n, cursorType, offset, i),
			Node:	n,
		}
	}
}

func LipidInventoryStartEndCursor(edges []*graphql_models.LipidInventoryEdge) (*string, *string) {
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

func LipidInventoryConnection(
	ctx context.Context,
	db *sql.DB,
	originalMods []qm.QueryMod,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.LipidInventoryOrdering,
) (*graphql_models.LipidInventoryConnection, error) {
	paginationMods, err := LipidInventoryPaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := LipidInventoryReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := models.LipidInventories(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*graphql_models.LipidInventoryEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := LipidInventoryEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := LipidInventoryStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
	return &graphql_models.LipidInventoryConnection{
		Edges:	edges,
		PageInfo: &graphql_models.PageInfo{
			HasNextPage:		hasNextPage,
			HasPreviousPage:	hasPreviousPage,
			StartCursor:		startCursor,
			EndCursor:		endCursor,
		},
	}, nil
}

var LipidSortColumn = map[graphql_models.LipidSort]string{
	graphql_models.LipidSortID:		models.LipidColumns.ID,
	graphql_models.LipidSortName:		models.LipidColumns.Name,
	graphql_models.LipidSortLauric:		models.LipidColumns.Lauric,
	graphql_models.LipidSortMyristic:	models.LipidColumns.Myristic,
	graphql_models.LipidSortPalmitic:	models.LipidColumns.Palmitic,
	graphql_models.LipidSortStearic:	models.LipidColumns.Stearic,
	graphql_models.LipidSortRicinoleic:	models.LipidColumns.Ricinoleic,
	graphql_models.LipidSortOleic:		models.LipidColumns.Oleic,
	graphql_models.LipidSortLinoleic:	models.LipidColumns.Linoleic,
	graphql_models.LipidSortLinolenic:	models.LipidColumns.Linolenic,
	graphql_models.LipidSortHardness:	models.LipidColumns.Hardness,
	graphql_models.LipidSortCleansing:	models.LipidColumns.Cleansing,
	graphql_models.LipidSortConditioning:	models.LipidColumns.Conditioning,
	graphql_models.LipidSortBubbly:		models.LipidColumns.Bubbly,
	graphql_models.LipidSortCreamy:		models.LipidColumns.Creamy,
	graphql_models.LipidSortIodine:		models.LipidColumns.Iodine,
	graphql_models.LipidSortIns:		models.LipidColumns.Ins,
	graphql_models.LipidSortInciName:	models.LipidColumns.InciName,
	graphql_models.LipidSortFamily:		models.LipidColumns.Family,
	graphql_models.LipidSortNaoh:		models.LipidColumns.Naoh,
	graphql_models.LipidSortCreatedAt:	models.LipidColumns.CreatedAt,
	graphql_models.LipidSortUpdatedAt:	models.LipidColumns.UpdatedAt,
	graphql_models.LipidSortDeletedAt:	models.LipidColumns.DeletedAt,
}

func LipidSortValueFromCursorValue(cursorValue string) (string, interface{}) {
	key, value := boilergql.FromCursorValue(cursorValue)
	column := LipidSortColumn[graphql_models.LipidSort(key)]

	if graphql_models.LipidSort(key) == graphql_models.LipidSortID {
		return column, boilergql.GetIDFromCursor(value)
	}

	return column, boilergql.StringToInterface(value)
}

func LipidSortCursorValue(sort graphql_models.LipidSort, m *graphql_models.Lipid) interface{} {
	switch sort {
	case graphql_models.LipidSortID:
		return m.ID
	case graphql_models.LipidSortName:
		return m.Name
	case graphql_models.LipidSortLauric:
		return m.Lauric
	case graphql_models.LipidSortMyristic:
		return m.Myristic
	case graphql_models.LipidSortPalmitic:
		return m.Palmitic
	case graphql_models.LipidSortStearic:
		return m.Stearic
	case graphql_models.LipidSortRicinoleic:
		return m.Ricinoleic
	case graphql_models.LipidSortOleic:
		return m.Oleic
	case graphql_models.LipidSortLinoleic:
		return m.Linoleic
	case graphql_models.LipidSortLinolenic:
		return m.Linolenic
	case graphql_models.LipidSortHardness:
		return m.Hardness
	case graphql_models.LipidSortCleansing:
		return m.Cleansing
	case graphql_models.LipidSortConditioning:
		return m.Conditioning
	case graphql_models.LipidSortBubbly:
		return m.Bubbly
	case graphql_models.LipidSortCreamy:
		return m.Creamy
	case graphql_models.LipidSortIodine:
		return m.Iodine
	case graphql_models.LipidSortIns:
		return m.Ins
	case graphql_models.LipidSortInciName:
		return m.InciName
	case graphql_models.LipidSortFamily:
		return m.Family
	case graphql_models.LipidSortNaoh:
		return m.Naoh
	case graphql_models.LipidSortCreatedAt:
		return m.CreatedAt
	case graphql_models.LipidSortUpdatedAt:
		return m.UpdatedAt
	case graphql_models.LipidSortDeletedAt:
		return m.DeletedAt
	}
	return nil
}

func LipidSortDirection(ordering []*graphql_models.LipidOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func FromLipidCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	var columns []string
	var values []interface{}

	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := LipidSortValueFromCursorValue(cursorValue)
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

func ToLipidCursor(ordering []*graphql_models.LipidOrdering, m *graphql_models.Lipid) string {
	var a []string
	var handledID bool

	for _, order := range ordering {
		if order.Sort == graphql_models.LipidSortID {
			handledID = true
		}
		value := LipidSortCursorValue(order.Sort, m)
		if value != nil {
			a = append(a, boilergql.ToCursorValue(string(order.Sort), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(graphql_models.LipidSortID), m.ID))
	}

	return boilergql.CursorValuesToString(a)
}

func LipidCursorType(ordering []*graphql_models.LipidOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func LipidCursorMods(ordering []*graphql_models.LipidOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if LipidCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromLipidCursor(*cursor, sign)
		}
		return boilergql.FromOffsetCursor(*cursor)
	}
	return nil
}

func LipidSortMods(ordering []*graphql_models.LipidOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, order := range ordering {
		if order.Sort == graphql_models.LipidSortID {
			handledID = true
		}
		column := LipidSortColumn[order.Sort]
		if column != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				column,
				boilergql.GetDirection(order.Direction, reverse),
			)))
		}
	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			models.LipidColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func LipidPaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*graphql_models.LipidOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := LipidSortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, LipidCursorMods(ordering, cursor, sign)...)
	mods = append(mods, LipidSortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func LipidPaginationMods(pagination boilergql.ConnectionPagination, ordering []*graphql_models.LipidOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	reverse := pagination.Backward != nil
	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	_, mods := LipidPaginationModsBase(pagination, ordering, reverse, limit)
	return mods, nil
}

func ToLipidCursorSwitch(ordering []*graphql_models.LipidOrdering, m *graphql_models.Lipid, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToLipidCursor(ordering, m)
	}
	return ""
}

func LipidReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.LipidOrdering,
) (bool, error) {
	reverse := pagination.Forward != nil
	cursor, reverseMods := LipidPaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := LipidCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return models.Lipids(reverseMods...).Count(ctx, db)
	})
}

func LipidEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*graphql_models.LipidOrdering) func(*models.Lipid, int) *graphql_models.LipidEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), LipidCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *models.Lipid, i int) *graphql_models.LipidEdge {
		n := LipidToGraphQL(m)
		return &graphql_models.LipidEdge{
			Cursor:	ToLipidCursorSwitch(ordering, n, cursorType, offset, i),
			Node:	n,
		}
	}
}

func LipidStartEndCursor(edges []*graphql_models.LipidEdge) (*string, *string) {
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

func LipidConnection(
	ctx context.Context,
	db *sql.DB,
	originalMods []qm.QueryMod,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.LipidOrdering,
) (*graphql_models.LipidConnection, error) {
	paginationMods, err := LipidPaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := LipidReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := models.Lipids(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*graphql_models.LipidEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := LipidEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := LipidStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
	return &graphql_models.LipidConnection{
		Edges:	edges,
		PageInfo: &graphql_models.PageInfo{
			HasNextPage:		hasNextPage,
			HasPreviousPage:	hasPreviousPage,
			StartCursor:		startCursor,
			EndCursor:		endCursor,
		},
	}, nil
}

var LyeInventorySortColumn = map[graphql_models.LyeInventorySort]string{
	graphql_models.LyeInventorySortID:		models.LyeInventoryColumns.ID,
	graphql_models.LyeInventorySortPurchaseDate:	models.LyeInventoryColumns.PurchaseDate,
	graphql_models.LyeInventorySortExpiryDate:	models.LyeInventoryColumns.ExpiryDate,
	graphql_models.LyeInventorySortCost:		models.LyeInventoryColumns.Cost,
	graphql_models.LyeInventorySortWeight:		models.LyeInventoryColumns.Weight,
	graphql_models.LyeInventorySortConcentration:	models.LyeInventoryColumns.Concentration,
	graphql_models.LyeInventorySortCreatedAt:	models.LyeInventoryColumns.CreatedAt,
	graphql_models.LyeInventorySortUpdatedAt:	models.LyeInventoryColumns.UpdatedAt,
	graphql_models.LyeInventorySortDeletedAt:	models.LyeInventoryColumns.DeletedAt,
}

func LyeInventorySortValueFromCursorValue(cursorValue string) (string, interface{}) {
	key, value := boilergql.FromCursorValue(cursorValue)
	column := LyeInventorySortColumn[graphql_models.LyeInventorySort(key)]

	if graphql_models.LyeInventorySort(key) == graphql_models.LyeInventorySortID {
		return column, boilergql.GetIDFromCursor(value)
	}

	return column, boilergql.StringToInterface(value)
}

func LyeInventorySortCursorValue(sort graphql_models.LyeInventorySort, m *graphql_models.LyeInventory) interface{} {
	switch sort {
	case graphql_models.LyeInventorySortID:
		return m.ID
	case graphql_models.LyeInventorySortPurchaseDate:
		return m.PurchaseDate
	case graphql_models.LyeInventorySortExpiryDate:
		return m.ExpiryDate
	case graphql_models.LyeInventorySortCost:
		return m.Cost
	case graphql_models.LyeInventorySortWeight:
		return m.Weight
	case graphql_models.LyeInventorySortConcentration:
		return m.Concentration
	case graphql_models.LyeInventorySortCreatedAt:
		return m.CreatedAt
	case graphql_models.LyeInventorySortUpdatedAt:
		return m.UpdatedAt
	case graphql_models.LyeInventorySortDeletedAt:
		return m.DeletedAt
	}
	return nil
}

func LyeInventorySortDirection(ordering []*graphql_models.LyeInventoryOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func FromLyeInventoryCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	var columns []string
	var values []interface{}

	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := LyeInventorySortValueFromCursorValue(cursorValue)
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

func ToLyeInventoryCursor(ordering []*graphql_models.LyeInventoryOrdering, m *graphql_models.LyeInventory) string {
	var a []string
	var handledID bool

	for _, order := range ordering {
		if order.Sort == graphql_models.LyeInventorySortID {
			handledID = true
		}
		value := LyeInventorySortCursorValue(order.Sort, m)
		if value != nil {
			a = append(a, boilergql.ToCursorValue(string(order.Sort), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(graphql_models.LyeInventorySortID), m.ID))
	}

	return boilergql.CursorValuesToString(a)
}

func LyeInventoryCursorType(ordering []*graphql_models.LyeInventoryOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func LyeInventoryCursorMods(ordering []*graphql_models.LyeInventoryOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if LyeInventoryCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromLyeInventoryCursor(*cursor, sign)
		}
		return boilergql.FromOffsetCursor(*cursor)
	}
	return nil
}

func LyeInventorySortMods(ordering []*graphql_models.LyeInventoryOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, order := range ordering {
		if order.Sort == graphql_models.LyeInventorySortID {
			handledID = true
		}
		column := LyeInventorySortColumn[order.Sort]
		if column != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				column,
				boilergql.GetDirection(order.Direction, reverse),
			)))
		}
	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			models.LyeInventoryColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func LyeInventoryPaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*graphql_models.LyeInventoryOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := LyeInventorySortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, LyeInventoryCursorMods(ordering, cursor, sign)...)
	mods = append(mods, LyeInventorySortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func LyeInventoryPaginationMods(pagination boilergql.ConnectionPagination, ordering []*graphql_models.LyeInventoryOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	reverse := pagination.Backward != nil
	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	_, mods := LyeInventoryPaginationModsBase(pagination, ordering, reverse, limit)
	return mods, nil
}

func ToLyeInventoryCursorSwitch(ordering []*graphql_models.LyeInventoryOrdering, m *graphql_models.LyeInventory, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToLyeInventoryCursor(ordering, m)
	}
	return ""
}

func LyeInventoryReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.LyeInventoryOrdering,
) (bool, error) {
	reverse := pagination.Forward != nil
	cursor, reverseMods := LyeInventoryPaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := LyeInventoryCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return models.LyeInventories(reverseMods...).Count(ctx, db)
	})
}

func LyeInventoryEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*graphql_models.LyeInventoryOrdering) func(*models.LyeInventory, int) *graphql_models.LyeInventoryEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), LyeInventoryCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *models.LyeInventory, i int) *graphql_models.LyeInventoryEdge {
		n := LyeInventoryToGraphQL(m)
		return &graphql_models.LyeInventoryEdge{
			Cursor:	ToLyeInventoryCursorSwitch(ordering, n, cursorType, offset, i),
			Node:	n,
		}
	}
}

func LyeInventoryStartEndCursor(edges []*graphql_models.LyeInventoryEdge) (*string, *string) {
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

func LyeInventoryConnection(
	ctx context.Context,
	db *sql.DB,
	originalMods []qm.QueryMod,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.LyeInventoryOrdering,
) (*graphql_models.LyeInventoryConnection, error) {
	paginationMods, err := LyeInventoryPaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := LyeInventoryReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := models.LyeInventories(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*graphql_models.LyeInventoryEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := LyeInventoryEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := LyeInventoryStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
	return &graphql_models.LyeInventoryConnection{
		Edges:	edges,
		PageInfo: &graphql_models.PageInfo{
			HasNextPage:		hasNextPage,
			HasPreviousPage:	hasPreviousPage,
			StartCursor:		startCursor,
			EndCursor:		endCursor,
		},
	}, nil
}

var LyeSortColumn = map[graphql_models.LyeSort]string{
	graphql_models.LyeSortID:		models.LyeColumns.ID,
	graphql_models.LyeSortKind:		models.LyeColumns.Kind,
	graphql_models.LyeSortName:		models.LyeColumns.Name,
	graphql_models.LyeSortNote:		models.LyeColumns.Note,
	graphql_models.LyeSortCreatedAt:	models.LyeColumns.CreatedAt,
	graphql_models.LyeSortUpdatedAt:	models.LyeColumns.UpdatedAt,
	graphql_models.LyeSortDeletedAt:	models.LyeColumns.DeletedAt,
}

func LyeSortValueFromCursorValue(cursorValue string) (string, interface{}) {
	key, value := boilergql.FromCursorValue(cursorValue)
	column := LyeSortColumn[graphql_models.LyeSort(key)]

	if graphql_models.LyeSort(key) == graphql_models.LyeSortID {
		return column, boilergql.GetIDFromCursor(value)
	}

	return column, boilergql.StringToInterface(value)
}

func LyeSortCursorValue(sort graphql_models.LyeSort, m *graphql_models.Lye) interface{} {
	switch sort {
	case graphql_models.LyeSortID:
		return m.ID
	case graphql_models.LyeSortKind:
		return m.Kind
	case graphql_models.LyeSortName:
		return m.Name
	case graphql_models.LyeSortNote:
		return m.Note
	case graphql_models.LyeSortCreatedAt:
		return m.CreatedAt
	case graphql_models.LyeSortUpdatedAt:
		return m.UpdatedAt
	case graphql_models.LyeSortDeletedAt:
		return m.DeletedAt
	}
	return nil
}

func LyeSortDirection(ordering []*graphql_models.LyeOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func FromLyeCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	var columns []string
	var values []interface{}

	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := LyeSortValueFromCursorValue(cursorValue)
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

func ToLyeCursor(ordering []*graphql_models.LyeOrdering, m *graphql_models.Lye) string {
	var a []string
	var handledID bool

	for _, order := range ordering {
		if order.Sort == graphql_models.LyeSortID {
			handledID = true
		}
		value := LyeSortCursorValue(order.Sort, m)
		if value != nil {
			a = append(a, boilergql.ToCursorValue(string(order.Sort), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(graphql_models.LyeSortID), m.ID))
	}

	return boilergql.CursorValuesToString(a)
}

func LyeCursorType(ordering []*graphql_models.LyeOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func LyeCursorMods(ordering []*graphql_models.LyeOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if LyeCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromLyeCursor(*cursor, sign)
		}
		return boilergql.FromOffsetCursor(*cursor)
	}
	return nil
}

func LyeSortMods(ordering []*graphql_models.LyeOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, order := range ordering {
		if order.Sort == graphql_models.LyeSortID {
			handledID = true
		}
		column := LyeSortColumn[order.Sort]
		if column != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				column,
				boilergql.GetDirection(order.Direction, reverse),
			)))
		}
	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			models.LyeColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func LyePaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*graphql_models.LyeOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := LyeSortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, LyeCursorMods(ordering, cursor, sign)...)
	mods = append(mods, LyeSortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func LyePaginationMods(pagination boilergql.ConnectionPagination, ordering []*graphql_models.LyeOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	reverse := pagination.Backward != nil
	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	_, mods := LyePaginationModsBase(pagination, ordering, reverse, limit)
	return mods, nil
}

func ToLyeCursorSwitch(ordering []*graphql_models.LyeOrdering, m *graphql_models.Lye, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToLyeCursor(ordering, m)
	}
	return ""
}

func LyeReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.LyeOrdering,
) (bool, error) {
	reverse := pagination.Forward != nil
	cursor, reverseMods := LyePaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := LyeCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return models.Lyes(reverseMods...).Count(ctx, db)
	})
}

func LyeEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*graphql_models.LyeOrdering) func(*models.Lye, int) *graphql_models.LyeEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), LyeCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *models.Lye, i int) *graphql_models.LyeEdge {
		n := LyeToGraphQL(m)
		return &graphql_models.LyeEdge{
			Cursor:	ToLyeCursorSwitch(ordering, n, cursorType, offset, i),
			Node:	n,
		}
	}
}

func LyeStartEndCursor(edges []*graphql_models.LyeEdge) (*string, *string) {
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

func LyeConnection(
	ctx context.Context,
	db *sql.DB,
	originalMods []qm.QueryMod,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.LyeOrdering,
) (*graphql_models.LyeConnection, error) {
	paginationMods, err := LyePaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := LyeReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := models.Lyes(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*graphql_models.LyeEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := LyeEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := LyeStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
	return &graphql_models.LyeConnection{
		Edges:	edges,
		PageInfo: &graphql_models.PageInfo{
			HasNextPage:		hasNextPage,
			HasPreviousPage:	hasPreviousPage,
			StartCursor:		startCursor,
			EndCursor:		endCursor,
		},
	}, nil
}

var RecipeAdditiveSortColumn = map[graphql_models.RecipeAdditiveSort]string{
	graphql_models.RecipeAdditiveSortID:		models.RecipeAdditiveColumns.ID,
	graphql_models.RecipeAdditiveSortPercentage:	models.RecipeAdditiveColumns.Percentage,
	graphql_models.RecipeAdditiveSortCreatedAt:	models.RecipeAdditiveColumns.CreatedAt,
	graphql_models.RecipeAdditiveSortUpdatedAt:	models.RecipeAdditiveColumns.UpdatedAt,
	graphql_models.RecipeAdditiveSortDeletedAt:	models.RecipeAdditiveColumns.DeletedAt,
}

func RecipeAdditiveSortValueFromCursorValue(cursorValue string) (string, interface{}) {
	key, value := boilergql.FromCursorValue(cursorValue)
	column := RecipeAdditiveSortColumn[graphql_models.RecipeAdditiveSort(key)]

	if graphql_models.RecipeAdditiveSort(key) == graphql_models.RecipeAdditiveSortID {
		return column, boilergql.GetIDFromCursor(value)
	}

	return column, boilergql.StringToInterface(value)
}

func RecipeAdditiveSortCursorValue(sort graphql_models.RecipeAdditiveSort, m *graphql_models.RecipeAdditive) interface{} {
	switch sort {
	case graphql_models.RecipeAdditiveSortID:
		return m.ID
	case graphql_models.RecipeAdditiveSortPercentage:
		return m.Percentage
	case graphql_models.RecipeAdditiveSortCreatedAt:
		return m.CreatedAt
	case graphql_models.RecipeAdditiveSortUpdatedAt:
		return m.UpdatedAt
	case graphql_models.RecipeAdditiveSortDeletedAt:
		return m.DeletedAt
	}
	return nil
}

func RecipeAdditiveSortDirection(ordering []*graphql_models.RecipeAdditiveOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func FromRecipeAdditiveCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	var columns []string
	var values []interface{}

	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := RecipeAdditiveSortValueFromCursorValue(cursorValue)
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

func ToRecipeAdditiveCursor(ordering []*graphql_models.RecipeAdditiveOrdering, m *graphql_models.RecipeAdditive) string {
	var a []string
	var handledID bool

	for _, order := range ordering {
		if order.Sort == graphql_models.RecipeAdditiveSortID {
			handledID = true
		}
		value := RecipeAdditiveSortCursorValue(order.Sort, m)
		if value != nil {
			a = append(a, boilergql.ToCursorValue(string(order.Sort), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(graphql_models.RecipeAdditiveSortID), m.ID))
	}

	return boilergql.CursorValuesToString(a)
}

func RecipeAdditiveCursorType(ordering []*graphql_models.RecipeAdditiveOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func RecipeAdditiveCursorMods(ordering []*graphql_models.RecipeAdditiveOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if RecipeAdditiveCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromRecipeAdditiveCursor(*cursor, sign)
		}
		return boilergql.FromOffsetCursor(*cursor)
	}
	return nil
}

func RecipeAdditiveSortMods(ordering []*graphql_models.RecipeAdditiveOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, order := range ordering {
		if order.Sort == graphql_models.RecipeAdditiveSortID {
			handledID = true
		}
		column := RecipeAdditiveSortColumn[order.Sort]
		if column != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				column,
				boilergql.GetDirection(order.Direction, reverse),
			)))
		}
	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			models.RecipeAdditiveColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func RecipeAdditivePaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*graphql_models.RecipeAdditiveOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := RecipeAdditiveSortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, RecipeAdditiveCursorMods(ordering, cursor, sign)...)
	mods = append(mods, RecipeAdditiveSortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func RecipeAdditivePaginationMods(pagination boilergql.ConnectionPagination, ordering []*graphql_models.RecipeAdditiveOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	reverse := pagination.Backward != nil
	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	_, mods := RecipeAdditivePaginationModsBase(pagination, ordering, reverse, limit)
	return mods, nil
}

func ToRecipeAdditiveCursorSwitch(ordering []*graphql_models.RecipeAdditiveOrdering, m *graphql_models.RecipeAdditive, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToRecipeAdditiveCursor(ordering, m)
	}
	return ""
}

func RecipeAdditiveReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.RecipeAdditiveOrdering,
) (bool, error) {
	reverse := pagination.Forward != nil
	cursor, reverseMods := RecipeAdditivePaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := RecipeAdditiveCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return models.RecipeAdditives(reverseMods...).Count(ctx, db)
	})
}

func RecipeAdditiveEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*graphql_models.RecipeAdditiveOrdering) func(*models.RecipeAdditive, int) *graphql_models.RecipeAdditiveEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), RecipeAdditiveCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *models.RecipeAdditive, i int) *graphql_models.RecipeAdditiveEdge {
		n := RecipeAdditiveToGraphQL(m)
		return &graphql_models.RecipeAdditiveEdge{
			Cursor:	ToRecipeAdditiveCursorSwitch(ordering, n, cursorType, offset, i),
			Node:	n,
		}
	}
}

func RecipeAdditiveStartEndCursor(edges []*graphql_models.RecipeAdditiveEdge) (*string, *string) {
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

func RecipeAdditiveConnection(
	ctx context.Context,
	db *sql.DB,
	originalMods []qm.QueryMod,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.RecipeAdditiveOrdering,
) (*graphql_models.RecipeAdditiveConnection, error) {
	paginationMods, err := RecipeAdditivePaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := RecipeAdditiveReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := models.RecipeAdditives(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*graphql_models.RecipeAdditiveEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := RecipeAdditiveEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := RecipeAdditiveStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
	return &graphql_models.RecipeAdditiveConnection{
		Edges:	edges,
		PageInfo: &graphql_models.PageInfo{
			HasNextPage:		hasNextPage,
			HasPreviousPage:	hasPreviousPage,
			StartCursor:		startCursor,
			EndCursor:		endCursor,
		},
	}, nil
}

var RecipeBatchAdditiveSortColumn = map[graphql_models.RecipeBatchAdditiveSort]string{
	graphql_models.RecipeBatchAdditiveSortID:		models.RecipeBatchAdditiveColumns.ID,
	graphql_models.RecipeBatchAdditiveSortWeight:		models.RecipeBatchAdditiveColumns.Weight,
	graphql_models.RecipeBatchAdditiveSortCost:		models.RecipeBatchAdditiveColumns.Cost,
	graphql_models.RecipeBatchAdditiveSortCreatedAt:	models.RecipeBatchAdditiveColumns.CreatedAt,
	graphql_models.RecipeBatchAdditiveSortUpdatedAt:	models.RecipeBatchAdditiveColumns.UpdatedAt,
	graphql_models.RecipeBatchAdditiveSortDeletedAt:	models.RecipeBatchAdditiveColumns.DeletedAt,
}

func RecipeBatchAdditiveSortValueFromCursorValue(cursorValue string) (string, interface{}) {
	key, value := boilergql.FromCursorValue(cursorValue)
	column := RecipeBatchAdditiveSortColumn[graphql_models.RecipeBatchAdditiveSort(key)]

	if graphql_models.RecipeBatchAdditiveSort(key) == graphql_models.RecipeBatchAdditiveSortID {
		return column, boilergql.GetIDFromCursor(value)
	}

	return column, boilergql.StringToInterface(value)
}

func RecipeBatchAdditiveSortCursorValue(sort graphql_models.RecipeBatchAdditiveSort, m *graphql_models.RecipeBatchAdditive) interface{} {
	switch sort {
	case graphql_models.RecipeBatchAdditiveSortID:
		return m.ID
	case graphql_models.RecipeBatchAdditiveSortWeight:
		return m.Weight
	case graphql_models.RecipeBatchAdditiveSortCost:
		return m.Cost
	case graphql_models.RecipeBatchAdditiveSortCreatedAt:
		return m.CreatedAt
	case graphql_models.RecipeBatchAdditiveSortUpdatedAt:
		return m.UpdatedAt
	case graphql_models.RecipeBatchAdditiveSortDeletedAt:
		return m.DeletedAt
	}
	return nil
}

func RecipeBatchAdditiveSortDirection(ordering []*graphql_models.RecipeBatchAdditiveOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func FromRecipeBatchAdditiveCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	var columns []string
	var values []interface{}

	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := RecipeBatchAdditiveSortValueFromCursorValue(cursorValue)
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

func ToRecipeBatchAdditiveCursor(ordering []*graphql_models.RecipeBatchAdditiveOrdering, m *graphql_models.RecipeBatchAdditive) string {
	var a []string
	var handledID bool

	for _, order := range ordering {
		if order.Sort == graphql_models.RecipeBatchAdditiveSortID {
			handledID = true
		}
		value := RecipeBatchAdditiveSortCursorValue(order.Sort, m)
		if value != nil {
			a = append(a, boilergql.ToCursorValue(string(order.Sort), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(graphql_models.RecipeBatchAdditiveSortID), m.ID))
	}

	return boilergql.CursorValuesToString(a)
}

func RecipeBatchAdditiveCursorType(ordering []*graphql_models.RecipeBatchAdditiveOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func RecipeBatchAdditiveCursorMods(ordering []*graphql_models.RecipeBatchAdditiveOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if RecipeBatchAdditiveCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromRecipeBatchAdditiveCursor(*cursor, sign)
		}
		return boilergql.FromOffsetCursor(*cursor)
	}
	return nil
}

func RecipeBatchAdditiveSortMods(ordering []*graphql_models.RecipeBatchAdditiveOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, order := range ordering {
		if order.Sort == graphql_models.RecipeBatchAdditiveSortID {
			handledID = true
		}
		column := RecipeBatchAdditiveSortColumn[order.Sort]
		if column != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				column,
				boilergql.GetDirection(order.Direction, reverse),
			)))
		}
	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			models.RecipeBatchAdditiveColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func RecipeBatchAdditivePaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*graphql_models.RecipeBatchAdditiveOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := RecipeBatchAdditiveSortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, RecipeBatchAdditiveCursorMods(ordering, cursor, sign)...)
	mods = append(mods, RecipeBatchAdditiveSortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func RecipeBatchAdditivePaginationMods(pagination boilergql.ConnectionPagination, ordering []*graphql_models.RecipeBatchAdditiveOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	reverse := pagination.Backward != nil
	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	_, mods := RecipeBatchAdditivePaginationModsBase(pagination, ordering, reverse, limit)
	return mods, nil
}

func ToRecipeBatchAdditiveCursorSwitch(ordering []*graphql_models.RecipeBatchAdditiveOrdering, m *graphql_models.RecipeBatchAdditive, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToRecipeBatchAdditiveCursor(ordering, m)
	}
	return ""
}

func RecipeBatchAdditiveReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.RecipeBatchAdditiveOrdering,
) (bool, error) {
	reverse := pagination.Forward != nil
	cursor, reverseMods := RecipeBatchAdditivePaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := RecipeBatchAdditiveCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return models.RecipeBatchAdditives(reverseMods...).Count(ctx, db)
	})
}

func RecipeBatchAdditiveEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*graphql_models.RecipeBatchAdditiveOrdering) func(*models.RecipeBatchAdditive, int) *graphql_models.RecipeBatchAdditiveEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), RecipeBatchAdditiveCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *models.RecipeBatchAdditive, i int) *graphql_models.RecipeBatchAdditiveEdge {
		n := RecipeBatchAdditiveToGraphQL(m)
		return &graphql_models.RecipeBatchAdditiveEdge{
			Cursor:	ToRecipeBatchAdditiveCursorSwitch(ordering, n, cursorType, offset, i),
			Node:	n,
		}
	}
}

func RecipeBatchAdditiveStartEndCursor(edges []*graphql_models.RecipeBatchAdditiveEdge) (*string, *string) {
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

func RecipeBatchAdditiveConnection(
	ctx context.Context,
	db *sql.DB,
	originalMods []qm.QueryMod,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.RecipeBatchAdditiveOrdering,
) (*graphql_models.RecipeBatchAdditiveConnection, error) {
	paginationMods, err := RecipeBatchAdditivePaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := RecipeBatchAdditiveReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := models.RecipeBatchAdditives(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*graphql_models.RecipeBatchAdditiveEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := RecipeBatchAdditiveEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := RecipeBatchAdditiveStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
	return &graphql_models.RecipeBatchAdditiveConnection{
		Edges:	edges,
		PageInfo: &graphql_models.PageInfo{
			HasNextPage:		hasNextPage,
			HasPreviousPage:	hasPreviousPage,
			StartCursor:		startCursor,
			EndCursor:		endCursor,
		},
	}, nil
}

var RecipeBatchFragranceSortColumn = map[graphql_models.RecipeBatchFragranceSort]string{
	graphql_models.RecipeBatchFragranceSortID:		models.RecipeBatchFragranceColumns.ID,
	graphql_models.RecipeBatchFragranceSortWeight:		models.RecipeBatchFragranceColumns.Weight,
	graphql_models.RecipeBatchFragranceSortCost:		models.RecipeBatchFragranceColumns.Cost,
	graphql_models.RecipeBatchFragranceSortCreatedAt:	models.RecipeBatchFragranceColumns.CreatedAt,
	graphql_models.RecipeBatchFragranceSortUpdatedAt:	models.RecipeBatchFragranceColumns.UpdatedAt,
	graphql_models.RecipeBatchFragranceSortDeletedAt:	models.RecipeBatchFragranceColumns.DeletedAt,
}

func RecipeBatchFragranceSortValueFromCursorValue(cursorValue string) (string, interface{}) {
	key, value := boilergql.FromCursorValue(cursorValue)
	column := RecipeBatchFragranceSortColumn[graphql_models.RecipeBatchFragranceSort(key)]

	if graphql_models.RecipeBatchFragranceSort(key) == graphql_models.RecipeBatchFragranceSortID {
		return column, boilergql.GetIDFromCursor(value)
	}

	return column, boilergql.StringToInterface(value)
}

func RecipeBatchFragranceSortCursorValue(sort graphql_models.RecipeBatchFragranceSort, m *graphql_models.RecipeBatchFragrance) interface{} {
	switch sort {
	case graphql_models.RecipeBatchFragranceSortID:
		return m.ID
	case graphql_models.RecipeBatchFragranceSortWeight:
		return m.Weight
	case graphql_models.RecipeBatchFragranceSortCost:
		return m.Cost
	case graphql_models.RecipeBatchFragranceSortCreatedAt:
		return m.CreatedAt
	case graphql_models.RecipeBatchFragranceSortUpdatedAt:
		return m.UpdatedAt
	case graphql_models.RecipeBatchFragranceSortDeletedAt:
		return m.DeletedAt
	}
	return nil
}

func RecipeBatchFragranceSortDirection(ordering []*graphql_models.RecipeBatchFragranceOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func FromRecipeBatchFragranceCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	var columns []string
	var values []interface{}

	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := RecipeBatchFragranceSortValueFromCursorValue(cursorValue)
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

func ToRecipeBatchFragranceCursor(ordering []*graphql_models.RecipeBatchFragranceOrdering, m *graphql_models.RecipeBatchFragrance) string {
	var a []string
	var handledID bool

	for _, order := range ordering {
		if order.Sort == graphql_models.RecipeBatchFragranceSortID {
			handledID = true
		}
		value := RecipeBatchFragranceSortCursorValue(order.Sort, m)
		if value != nil {
			a = append(a, boilergql.ToCursorValue(string(order.Sort), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(graphql_models.RecipeBatchFragranceSortID), m.ID))
	}

	return boilergql.CursorValuesToString(a)
}

func RecipeBatchFragranceCursorType(ordering []*graphql_models.RecipeBatchFragranceOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func RecipeBatchFragranceCursorMods(ordering []*graphql_models.RecipeBatchFragranceOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if RecipeBatchFragranceCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromRecipeBatchFragranceCursor(*cursor, sign)
		}
		return boilergql.FromOffsetCursor(*cursor)
	}
	return nil
}

func RecipeBatchFragranceSortMods(ordering []*graphql_models.RecipeBatchFragranceOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, order := range ordering {
		if order.Sort == graphql_models.RecipeBatchFragranceSortID {
			handledID = true
		}
		column := RecipeBatchFragranceSortColumn[order.Sort]
		if column != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				column,
				boilergql.GetDirection(order.Direction, reverse),
			)))
		}
	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			models.RecipeBatchFragranceColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func RecipeBatchFragrancePaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*graphql_models.RecipeBatchFragranceOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := RecipeBatchFragranceSortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, RecipeBatchFragranceCursorMods(ordering, cursor, sign)...)
	mods = append(mods, RecipeBatchFragranceSortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func RecipeBatchFragrancePaginationMods(pagination boilergql.ConnectionPagination, ordering []*graphql_models.RecipeBatchFragranceOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	reverse := pagination.Backward != nil
	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	_, mods := RecipeBatchFragrancePaginationModsBase(pagination, ordering, reverse, limit)
	return mods, nil
}

func ToRecipeBatchFragranceCursorSwitch(ordering []*graphql_models.RecipeBatchFragranceOrdering, m *graphql_models.RecipeBatchFragrance, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToRecipeBatchFragranceCursor(ordering, m)
	}
	return ""
}

func RecipeBatchFragranceReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.RecipeBatchFragranceOrdering,
) (bool, error) {
	reverse := pagination.Forward != nil
	cursor, reverseMods := RecipeBatchFragrancePaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := RecipeBatchFragranceCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return models.RecipeBatchFragrances(reverseMods...).Count(ctx, db)
	})
}

func RecipeBatchFragranceEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*graphql_models.RecipeBatchFragranceOrdering) func(*models.RecipeBatchFragrance, int) *graphql_models.RecipeBatchFragranceEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), RecipeBatchFragranceCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *models.RecipeBatchFragrance, i int) *graphql_models.RecipeBatchFragranceEdge {
		n := RecipeBatchFragranceToGraphQL(m)
		return &graphql_models.RecipeBatchFragranceEdge{
			Cursor:	ToRecipeBatchFragranceCursorSwitch(ordering, n, cursorType, offset, i),
			Node:	n,
		}
	}
}

func RecipeBatchFragranceStartEndCursor(edges []*graphql_models.RecipeBatchFragranceEdge) (*string, *string) {
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

func RecipeBatchFragranceConnection(
	ctx context.Context,
	db *sql.DB,
	originalMods []qm.QueryMod,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.RecipeBatchFragranceOrdering,
) (*graphql_models.RecipeBatchFragranceConnection, error) {
	paginationMods, err := RecipeBatchFragrancePaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := RecipeBatchFragranceReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := models.RecipeBatchFragrances(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*graphql_models.RecipeBatchFragranceEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := RecipeBatchFragranceEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := RecipeBatchFragranceStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
	return &graphql_models.RecipeBatchFragranceConnection{
		Edges:	edges,
		PageInfo: &graphql_models.PageInfo{
			HasNextPage:		hasNextPage,
			HasPreviousPage:	hasPreviousPage,
			StartCursor:		startCursor,
			EndCursor:		endCursor,
		},
	}, nil
}

var RecipeBatchLipidSortColumn = map[graphql_models.RecipeBatchLipidSort]string{
	graphql_models.RecipeBatchLipidSortID:		models.RecipeBatchLipidColumns.ID,
	graphql_models.RecipeBatchLipidSortWeight:	models.RecipeBatchLipidColumns.Weight,
	graphql_models.RecipeBatchLipidSortCost:	models.RecipeBatchLipidColumns.Cost,
	graphql_models.RecipeBatchLipidSortCreatedAt:	models.RecipeBatchLipidColumns.CreatedAt,
	graphql_models.RecipeBatchLipidSortUpdatedAt:	models.RecipeBatchLipidColumns.UpdatedAt,
	graphql_models.RecipeBatchLipidSortDeletedAt:	models.RecipeBatchLipidColumns.DeletedAt,
}

func RecipeBatchLipidSortValueFromCursorValue(cursorValue string) (string, interface{}) {
	key, value := boilergql.FromCursorValue(cursorValue)
	column := RecipeBatchLipidSortColumn[graphql_models.RecipeBatchLipidSort(key)]

	if graphql_models.RecipeBatchLipidSort(key) == graphql_models.RecipeBatchLipidSortID {
		return column, boilergql.GetIDFromCursor(value)
	}

	return column, boilergql.StringToInterface(value)
}

func RecipeBatchLipidSortCursorValue(sort graphql_models.RecipeBatchLipidSort, m *graphql_models.RecipeBatchLipid) interface{} {
	switch sort {
	case graphql_models.RecipeBatchLipidSortID:
		return m.ID
	case graphql_models.RecipeBatchLipidSortWeight:
		return m.Weight
	case graphql_models.RecipeBatchLipidSortCost:
		return m.Cost
	case graphql_models.RecipeBatchLipidSortCreatedAt:
		return m.CreatedAt
	case graphql_models.RecipeBatchLipidSortUpdatedAt:
		return m.UpdatedAt
	case graphql_models.RecipeBatchLipidSortDeletedAt:
		return m.DeletedAt
	}
	return nil
}

func RecipeBatchLipidSortDirection(ordering []*graphql_models.RecipeBatchLipidOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func FromRecipeBatchLipidCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	var columns []string
	var values []interface{}

	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := RecipeBatchLipidSortValueFromCursorValue(cursorValue)
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

func ToRecipeBatchLipidCursor(ordering []*graphql_models.RecipeBatchLipidOrdering, m *graphql_models.RecipeBatchLipid) string {
	var a []string
	var handledID bool

	for _, order := range ordering {
		if order.Sort == graphql_models.RecipeBatchLipidSortID {
			handledID = true
		}
		value := RecipeBatchLipidSortCursorValue(order.Sort, m)
		if value != nil {
			a = append(a, boilergql.ToCursorValue(string(order.Sort), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(graphql_models.RecipeBatchLipidSortID), m.ID))
	}

	return boilergql.CursorValuesToString(a)
}

func RecipeBatchLipidCursorType(ordering []*graphql_models.RecipeBatchLipidOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func RecipeBatchLipidCursorMods(ordering []*graphql_models.RecipeBatchLipidOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if RecipeBatchLipidCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromRecipeBatchLipidCursor(*cursor, sign)
		}
		return boilergql.FromOffsetCursor(*cursor)
	}
	return nil
}

func RecipeBatchLipidSortMods(ordering []*graphql_models.RecipeBatchLipidOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, order := range ordering {
		if order.Sort == graphql_models.RecipeBatchLipidSortID {
			handledID = true
		}
		column := RecipeBatchLipidSortColumn[order.Sort]
		if column != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				column,
				boilergql.GetDirection(order.Direction, reverse),
			)))
		}
	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			models.RecipeBatchLipidColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func RecipeBatchLipidPaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*graphql_models.RecipeBatchLipidOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := RecipeBatchLipidSortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, RecipeBatchLipidCursorMods(ordering, cursor, sign)...)
	mods = append(mods, RecipeBatchLipidSortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func RecipeBatchLipidPaginationMods(pagination boilergql.ConnectionPagination, ordering []*graphql_models.RecipeBatchLipidOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	reverse := pagination.Backward != nil
	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	_, mods := RecipeBatchLipidPaginationModsBase(pagination, ordering, reverse, limit)
	return mods, nil
}

func ToRecipeBatchLipidCursorSwitch(ordering []*graphql_models.RecipeBatchLipidOrdering, m *graphql_models.RecipeBatchLipid, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToRecipeBatchLipidCursor(ordering, m)
	}
	return ""
}

func RecipeBatchLipidReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.RecipeBatchLipidOrdering,
) (bool, error) {
	reverse := pagination.Forward != nil
	cursor, reverseMods := RecipeBatchLipidPaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := RecipeBatchLipidCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return models.RecipeBatchLipids(reverseMods...).Count(ctx, db)
	})
}

func RecipeBatchLipidEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*graphql_models.RecipeBatchLipidOrdering) func(*models.RecipeBatchLipid, int) *graphql_models.RecipeBatchLipidEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), RecipeBatchLipidCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *models.RecipeBatchLipid, i int) *graphql_models.RecipeBatchLipidEdge {
		n := RecipeBatchLipidToGraphQL(m)
		return &graphql_models.RecipeBatchLipidEdge{
			Cursor:	ToRecipeBatchLipidCursorSwitch(ordering, n, cursorType, offset, i),
			Node:	n,
		}
	}
}

func RecipeBatchLipidStartEndCursor(edges []*graphql_models.RecipeBatchLipidEdge) (*string, *string) {
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

func RecipeBatchLipidConnection(
	ctx context.Context,
	db *sql.DB,
	originalMods []qm.QueryMod,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.RecipeBatchLipidOrdering,
) (*graphql_models.RecipeBatchLipidConnection, error) {
	paginationMods, err := RecipeBatchLipidPaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := RecipeBatchLipidReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := models.RecipeBatchLipids(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*graphql_models.RecipeBatchLipidEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := RecipeBatchLipidEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := RecipeBatchLipidStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
	return &graphql_models.RecipeBatchLipidConnection{
		Edges:	edges,
		PageInfo: &graphql_models.PageInfo{
			HasNextPage:		hasNextPage,
			HasPreviousPage:	hasPreviousPage,
			StartCursor:		startCursor,
			EndCursor:		endCursor,
		},
	}, nil
}

var RecipeBatchLyeSortColumn = map[graphql_models.RecipeBatchLyeSort]string{
	graphql_models.RecipeBatchLyeSortID:		models.RecipeBatchLyeColumns.ID,
	graphql_models.RecipeBatchLyeSortWeight:	models.RecipeBatchLyeColumns.Weight,
	graphql_models.RecipeBatchLyeSortDiscount:	models.RecipeBatchLyeColumns.Discount,
	graphql_models.RecipeBatchLyeSortCost:		models.RecipeBatchLyeColumns.Cost,
	graphql_models.RecipeBatchLyeSortCreatedAt:	models.RecipeBatchLyeColumns.CreatedAt,
	graphql_models.RecipeBatchLyeSortUpdatedAt:	models.RecipeBatchLyeColumns.UpdatedAt,
	graphql_models.RecipeBatchLyeSortDeletedAt:	models.RecipeBatchLyeColumns.DeletedAt,
}

func RecipeBatchLyeSortValueFromCursorValue(cursorValue string) (string, interface{}) {
	key, value := boilergql.FromCursorValue(cursorValue)
	column := RecipeBatchLyeSortColumn[graphql_models.RecipeBatchLyeSort(key)]

	if graphql_models.RecipeBatchLyeSort(key) == graphql_models.RecipeBatchLyeSortID {
		return column, boilergql.GetIDFromCursor(value)
	}

	return column, boilergql.StringToInterface(value)
}

func RecipeBatchLyeSortCursorValue(sort graphql_models.RecipeBatchLyeSort, m *graphql_models.RecipeBatchLye) interface{} {
	switch sort {
	case graphql_models.RecipeBatchLyeSortID:
		return m.ID
	case graphql_models.RecipeBatchLyeSortWeight:
		return m.Weight
	case graphql_models.RecipeBatchLyeSortDiscount:
		return m.Discount
	case graphql_models.RecipeBatchLyeSortCost:
		return m.Cost
	case graphql_models.RecipeBatchLyeSortCreatedAt:
		return m.CreatedAt
	case graphql_models.RecipeBatchLyeSortUpdatedAt:
		return m.UpdatedAt
	case graphql_models.RecipeBatchLyeSortDeletedAt:
		return m.DeletedAt
	}
	return nil
}

func RecipeBatchLyeSortDirection(ordering []*graphql_models.RecipeBatchLyeOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func FromRecipeBatchLyeCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	var columns []string
	var values []interface{}

	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := RecipeBatchLyeSortValueFromCursorValue(cursorValue)
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

func ToRecipeBatchLyeCursor(ordering []*graphql_models.RecipeBatchLyeOrdering, m *graphql_models.RecipeBatchLye) string {
	var a []string
	var handledID bool

	for _, order := range ordering {
		if order.Sort == graphql_models.RecipeBatchLyeSortID {
			handledID = true
		}
		value := RecipeBatchLyeSortCursorValue(order.Sort, m)
		if value != nil {
			a = append(a, boilergql.ToCursorValue(string(order.Sort), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(graphql_models.RecipeBatchLyeSortID), m.ID))
	}

	return boilergql.CursorValuesToString(a)
}

func RecipeBatchLyeCursorType(ordering []*graphql_models.RecipeBatchLyeOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func RecipeBatchLyeCursorMods(ordering []*graphql_models.RecipeBatchLyeOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if RecipeBatchLyeCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromRecipeBatchLyeCursor(*cursor, sign)
		}
		return boilergql.FromOffsetCursor(*cursor)
	}
	return nil
}

func RecipeBatchLyeSortMods(ordering []*graphql_models.RecipeBatchLyeOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, order := range ordering {
		if order.Sort == graphql_models.RecipeBatchLyeSortID {
			handledID = true
		}
		column := RecipeBatchLyeSortColumn[order.Sort]
		if column != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				column,
				boilergql.GetDirection(order.Direction, reverse),
			)))
		}
	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			models.RecipeBatchLyeColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func RecipeBatchLyePaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*graphql_models.RecipeBatchLyeOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := RecipeBatchLyeSortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, RecipeBatchLyeCursorMods(ordering, cursor, sign)...)
	mods = append(mods, RecipeBatchLyeSortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func RecipeBatchLyePaginationMods(pagination boilergql.ConnectionPagination, ordering []*graphql_models.RecipeBatchLyeOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	reverse := pagination.Backward != nil
	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	_, mods := RecipeBatchLyePaginationModsBase(pagination, ordering, reverse, limit)
	return mods, nil
}

func ToRecipeBatchLyeCursorSwitch(ordering []*graphql_models.RecipeBatchLyeOrdering, m *graphql_models.RecipeBatchLye, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToRecipeBatchLyeCursor(ordering, m)
	}
	return ""
}

func RecipeBatchLyeReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.RecipeBatchLyeOrdering,
) (bool, error) {
	reverse := pagination.Forward != nil
	cursor, reverseMods := RecipeBatchLyePaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := RecipeBatchLyeCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return models.RecipeBatchLyes(reverseMods...).Count(ctx, db)
	})
}

func RecipeBatchLyeEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*graphql_models.RecipeBatchLyeOrdering) func(*models.RecipeBatchLye, int) *graphql_models.RecipeBatchLyeEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), RecipeBatchLyeCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *models.RecipeBatchLye, i int) *graphql_models.RecipeBatchLyeEdge {
		n := RecipeBatchLyeToGraphQL(m)
		return &graphql_models.RecipeBatchLyeEdge{
			Cursor:	ToRecipeBatchLyeCursorSwitch(ordering, n, cursorType, offset, i),
			Node:	n,
		}
	}
}

func RecipeBatchLyeStartEndCursor(edges []*graphql_models.RecipeBatchLyeEdge) (*string, *string) {
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

func RecipeBatchLyeConnection(
	ctx context.Context,
	db *sql.DB,
	originalMods []qm.QueryMod,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.RecipeBatchLyeOrdering,
) (*graphql_models.RecipeBatchLyeConnection, error) {
	paginationMods, err := RecipeBatchLyePaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := RecipeBatchLyeReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := models.RecipeBatchLyes(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*graphql_models.RecipeBatchLyeEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := RecipeBatchLyeEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := RecipeBatchLyeStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
	return &graphql_models.RecipeBatchLyeConnection{
		Edges:	edges,
		PageInfo: &graphql_models.PageInfo{
			HasNextPage:		hasNextPage,
			HasPreviousPage:	hasPreviousPage,
			StartCursor:		startCursor,
			EndCursor:		endCursor,
		},
	}, nil
}

var RecipeBatchNoteSortColumn = map[graphql_models.RecipeBatchNoteSort]string{
	graphql_models.RecipeBatchNoteSortID:		models.RecipeBatchNoteColumns.ID,
	graphql_models.RecipeBatchNoteSortNote:		models.RecipeBatchNoteColumns.Note,
	graphql_models.RecipeBatchNoteSortLink:		models.RecipeBatchNoteColumns.Link,
	graphql_models.RecipeBatchNoteSortCreatedAt:	models.RecipeBatchNoteColumns.CreatedAt,
	graphql_models.RecipeBatchNoteSortUpdatedAt:	models.RecipeBatchNoteColumns.UpdatedAt,
	graphql_models.RecipeBatchNoteSortDeletedAt:	models.RecipeBatchNoteColumns.DeletedAt,
}

func RecipeBatchNoteSortValueFromCursorValue(cursorValue string) (string, interface{}) {
	key, value := boilergql.FromCursorValue(cursorValue)
	column := RecipeBatchNoteSortColumn[graphql_models.RecipeBatchNoteSort(key)]

	if graphql_models.RecipeBatchNoteSort(key) == graphql_models.RecipeBatchNoteSortID {
		return column, boilergql.GetIDFromCursor(value)
	}

	return column, boilergql.StringToInterface(value)
}

func RecipeBatchNoteSortCursorValue(sort graphql_models.RecipeBatchNoteSort, m *graphql_models.RecipeBatchNote) interface{} {
	switch sort {
	case graphql_models.RecipeBatchNoteSortID:
		return m.ID
	case graphql_models.RecipeBatchNoteSortNote:
		return m.Note
	case graphql_models.RecipeBatchNoteSortLink:
		return m.Link
	case graphql_models.RecipeBatchNoteSortCreatedAt:
		return m.CreatedAt
	case graphql_models.RecipeBatchNoteSortUpdatedAt:
		return m.UpdatedAt
	case graphql_models.RecipeBatchNoteSortDeletedAt:
		return m.DeletedAt
	}
	return nil
}

func RecipeBatchNoteSortDirection(ordering []*graphql_models.RecipeBatchNoteOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func FromRecipeBatchNoteCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	var columns []string
	var values []interface{}

	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := RecipeBatchNoteSortValueFromCursorValue(cursorValue)
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

func ToRecipeBatchNoteCursor(ordering []*graphql_models.RecipeBatchNoteOrdering, m *graphql_models.RecipeBatchNote) string {
	var a []string
	var handledID bool

	for _, order := range ordering {
		if order.Sort == graphql_models.RecipeBatchNoteSortID {
			handledID = true
		}
		value := RecipeBatchNoteSortCursorValue(order.Sort, m)
		if value != nil {
			a = append(a, boilergql.ToCursorValue(string(order.Sort), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(graphql_models.RecipeBatchNoteSortID), m.ID))
	}

	return boilergql.CursorValuesToString(a)
}

func RecipeBatchNoteCursorType(ordering []*graphql_models.RecipeBatchNoteOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func RecipeBatchNoteCursorMods(ordering []*graphql_models.RecipeBatchNoteOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if RecipeBatchNoteCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromRecipeBatchNoteCursor(*cursor, sign)
		}
		return boilergql.FromOffsetCursor(*cursor)
	}
	return nil
}

func RecipeBatchNoteSortMods(ordering []*graphql_models.RecipeBatchNoteOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, order := range ordering {
		if order.Sort == graphql_models.RecipeBatchNoteSortID {
			handledID = true
		}
		column := RecipeBatchNoteSortColumn[order.Sort]
		if column != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				column,
				boilergql.GetDirection(order.Direction, reverse),
			)))
		}
	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			models.RecipeBatchNoteColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func RecipeBatchNotePaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*graphql_models.RecipeBatchNoteOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := RecipeBatchNoteSortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, RecipeBatchNoteCursorMods(ordering, cursor, sign)...)
	mods = append(mods, RecipeBatchNoteSortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func RecipeBatchNotePaginationMods(pagination boilergql.ConnectionPagination, ordering []*graphql_models.RecipeBatchNoteOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	reverse := pagination.Backward != nil
	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	_, mods := RecipeBatchNotePaginationModsBase(pagination, ordering, reverse, limit)
	return mods, nil
}

func ToRecipeBatchNoteCursorSwitch(ordering []*graphql_models.RecipeBatchNoteOrdering, m *graphql_models.RecipeBatchNote, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToRecipeBatchNoteCursor(ordering, m)
	}
	return ""
}

func RecipeBatchNoteReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.RecipeBatchNoteOrdering,
) (bool, error) {
	reverse := pagination.Forward != nil
	cursor, reverseMods := RecipeBatchNotePaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := RecipeBatchNoteCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return models.RecipeBatchNotes(reverseMods...).Count(ctx, db)
	})
}

func RecipeBatchNoteEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*graphql_models.RecipeBatchNoteOrdering) func(*models.RecipeBatchNote, int) *graphql_models.RecipeBatchNoteEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), RecipeBatchNoteCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *models.RecipeBatchNote, i int) *graphql_models.RecipeBatchNoteEdge {
		n := RecipeBatchNoteToGraphQL(m)
		return &graphql_models.RecipeBatchNoteEdge{
			Cursor:	ToRecipeBatchNoteCursorSwitch(ordering, n, cursorType, offset, i),
			Node:	n,
		}
	}
}

func RecipeBatchNoteStartEndCursor(edges []*graphql_models.RecipeBatchNoteEdge) (*string, *string) {
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

func RecipeBatchNoteConnection(
	ctx context.Context,
	db *sql.DB,
	originalMods []qm.QueryMod,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.RecipeBatchNoteOrdering,
) (*graphql_models.RecipeBatchNoteConnection, error) {
	paginationMods, err := RecipeBatchNotePaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := RecipeBatchNoteReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := models.RecipeBatchNotes(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*graphql_models.RecipeBatchNoteEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := RecipeBatchNoteEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := RecipeBatchNoteStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
	return &graphql_models.RecipeBatchNoteConnection{
		Edges:	edges,
		PageInfo: &graphql_models.PageInfo{
			HasNextPage:		hasNextPage,
			HasPreviousPage:	hasPreviousPage,
			StartCursor:		startCursor,
			EndCursor:		endCursor,
		},
	}, nil
}

var RecipeBatchSortColumn = map[graphql_models.RecipeBatchSort]string{
	graphql_models.RecipeBatchSortID:		models.RecipeBatchColumns.ID,
	graphql_models.RecipeBatchSortTag:		models.RecipeBatchColumns.Tag,
	graphql_models.RecipeBatchSortProductionDate:	models.RecipeBatchColumns.ProductionDate,
	graphql_models.RecipeBatchSortSellableDate:	models.RecipeBatchColumns.SellableDate,
	graphql_models.RecipeBatchSortNote:		models.RecipeBatchColumns.Note,
	graphql_models.RecipeBatchSortLipidWeight:	models.RecipeBatchColumns.LipidWeight,
	graphql_models.RecipeBatchSortProductionWeight:	models.RecipeBatchColumns.ProductionWeight,
	graphql_models.RecipeBatchSortCuredWeight:	models.RecipeBatchColumns.CuredWeight,
	graphql_models.RecipeBatchSortCreatedAt:	models.RecipeBatchColumns.CreatedAt,
	graphql_models.RecipeBatchSortUpdatedAt:	models.RecipeBatchColumns.UpdatedAt,
	graphql_models.RecipeBatchSortDeletedAt:	models.RecipeBatchColumns.DeletedAt,
}

func RecipeBatchSortValueFromCursorValue(cursorValue string) (string, interface{}) {
	key, value := boilergql.FromCursorValue(cursorValue)
	column := RecipeBatchSortColumn[graphql_models.RecipeBatchSort(key)]

	if graphql_models.RecipeBatchSort(key) == graphql_models.RecipeBatchSortID {
		return column, boilergql.GetIDFromCursor(value)
	}

	return column, boilergql.StringToInterface(value)
}

func RecipeBatchSortCursorValue(sort graphql_models.RecipeBatchSort, m *graphql_models.RecipeBatch) interface{} {
	switch sort {
	case graphql_models.RecipeBatchSortID:
		return m.ID
	case graphql_models.RecipeBatchSortTag:
		return m.Tag
	case graphql_models.RecipeBatchSortProductionDate:
		return m.ProductionDate
	case graphql_models.RecipeBatchSortSellableDate:
		return m.SellableDate
	case graphql_models.RecipeBatchSortNote:
		return m.Note
	case graphql_models.RecipeBatchSortLipidWeight:
		return m.LipidWeight
	case graphql_models.RecipeBatchSortProductionWeight:
		return m.ProductionWeight
	case graphql_models.RecipeBatchSortCuredWeight:
		return m.CuredWeight
	case graphql_models.RecipeBatchSortCreatedAt:
		return m.CreatedAt
	case graphql_models.RecipeBatchSortUpdatedAt:
		return m.UpdatedAt
	case graphql_models.RecipeBatchSortDeletedAt:
		return m.DeletedAt
	}
	return nil
}

func RecipeBatchSortDirection(ordering []*graphql_models.RecipeBatchOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func FromRecipeBatchCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	var columns []string
	var values []interface{}

	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := RecipeBatchSortValueFromCursorValue(cursorValue)
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

func ToRecipeBatchCursor(ordering []*graphql_models.RecipeBatchOrdering, m *graphql_models.RecipeBatch) string {
	var a []string
	var handledID bool

	for _, order := range ordering {
		if order.Sort == graphql_models.RecipeBatchSortID {
			handledID = true
		}
		value := RecipeBatchSortCursorValue(order.Sort, m)
		if value != nil {
			a = append(a, boilergql.ToCursorValue(string(order.Sort), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(graphql_models.RecipeBatchSortID), m.ID))
	}

	return boilergql.CursorValuesToString(a)
}

func RecipeBatchCursorType(ordering []*graphql_models.RecipeBatchOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func RecipeBatchCursorMods(ordering []*graphql_models.RecipeBatchOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if RecipeBatchCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromRecipeBatchCursor(*cursor, sign)
		}
		return boilergql.FromOffsetCursor(*cursor)
	}
	return nil
}

func RecipeBatchSortMods(ordering []*graphql_models.RecipeBatchOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, order := range ordering {
		if order.Sort == graphql_models.RecipeBatchSortID {
			handledID = true
		}
		column := RecipeBatchSortColumn[order.Sort]
		if column != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				column,
				boilergql.GetDirection(order.Direction, reverse),
			)))
		}
	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			models.RecipeBatchColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func RecipeBatchPaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*graphql_models.RecipeBatchOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := RecipeBatchSortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, RecipeBatchCursorMods(ordering, cursor, sign)...)
	mods = append(mods, RecipeBatchSortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func RecipeBatchPaginationMods(pagination boilergql.ConnectionPagination, ordering []*graphql_models.RecipeBatchOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	reverse := pagination.Backward != nil
	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	_, mods := RecipeBatchPaginationModsBase(pagination, ordering, reverse, limit)
	return mods, nil
}

func ToRecipeBatchCursorSwitch(ordering []*graphql_models.RecipeBatchOrdering, m *graphql_models.RecipeBatch, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToRecipeBatchCursor(ordering, m)
	}
	return ""
}

func RecipeBatchReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.RecipeBatchOrdering,
) (bool, error) {
	reverse := pagination.Forward != nil
	cursor, reverseMods := RecipeBatchPaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := RecipeBatchCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return models.RecipeBatches(reverseMods...).Count(ctx, db)
	})
}

func RecipeBatchEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*graphql_models.RecipeBatchOrdering) func(*models.RecipeBatch, int) *graphql_models.RecipeBatchEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), RecipeBatchCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *models.RecipeBatch, i int) *graphql_models.RecipeBatchEdge {
		n := RecipeBatchToGraphQL(m)
		return &graphql_models.RecipeBatchEdge{
			Cursor:	ToRecipeBatchCursorSwitch(ordering, n, cursorType, offset, i),
			Node:	n,
		}
	}
}

func RecipeBatchStartEndCursor(edges []*graphql_models.RecipeBatchEdge) (*string, *string) {
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

func RecipeBatchConnection(
	ctx context.Context,
	db *sql.DB,
	originalMods []qm.QueryMod,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.RecipeBatchOrdering,
) (*graphql_models.RecipeBatchConnection, error) {
	paginationMods, err := RecipeBatchPaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := RecipeBatchReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := models.RecipeBatches(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*graphql_models.RecipeBatchEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := RecipeBatchEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := RecipeBatchStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
	return &graphql_models.RecipeBatchConnection{
		Edges:	edges,
		PageInfo: &graphql_models.PageInfo{
			HasNextPage:		hasNextPage,
			HasPreviousPage:	hasPreviousPage,
			StartCursor:		startCursor,
			EndCursor:		endCursor,
		},
	}, nil
}

var RecipeFragranceSortColumn = map[graphql_models.RecipeFragranceSort]string{
	graphql_models.RecipeFragranceSortID:		models.RecipeFragranceColumns.ID,
	graphql_models.RecipeFragranceSortPercentage:	models.RecipeFragranceColumns.Percentage,
	graphql_models.RecipeFragranceSortCreatedAt:	models.RecipeFragranceColumns.CreatedAt,
	graphql_models.RecipeFragranceSortUpdatedAt:	models.RecipeFragranceColumns.UpdatedAt,
	graphql_models.RecipeFragranceSortDeletedAt:	models.RecipeFragranceColumns.DeletedAt,
}

func RecipeFragranceSortValueFromCursorValue(cursorValue string) (string, interface{}) {
	key, value := boilergql.FromCursorValue(cursorValue)
	column := RecipeFragranceSortColumn[graphql_models.RecipeFragranceSort(key)]

	if graphql_models.RecipeFragranceSort(key) == graphql_models.RecipeFragranceSortID {
		return column, boilergql.GetIDFromCursor(value)
	}

	return column, boilergql.StringToInterface(value)
}

func RecipeFragranceSortCursorValue(sort graphql_models.RecipeFragranceSort, m *graphql_models.RecipeFragrance) interface{} {
	switch sort {
	case graphql_models.RecipeFragranceSortID:
		return m.ID
	case graphql_models.RecipeFragranceSortPercentage:
		return m.Percentage
	case graphql_models.RecipeFragranceSortCreatedAt:
		return m.CreatedAt
	case graphql_models.RecipeFragranceSortUpdatedAt:
		return m.UpdatedAt
	case graphql_models.RecipeFragranceSortDeletedAt:
		return m.DeletedAt
	}
	return nil
}

func RecipeFragranceSortDirection(ordering []*graphql_models.RecipeFragranceOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func FromRecipeFragranceCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	var columns []string
	var values []interface{}

	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := RecipeFragranceSortValueFromCursorValue(cursorValue)
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

func ToRecipeFragranceCursor(ordering []*graphql_models.RecipeFragranceOrdering, m *graphql_models.RecipeFragrance) string {
	var a []string
	var handledID bool

	for _, order := range ordering {
		if order.Sort == graphql_models.RecipeFragranceSortID {
			handledID = true
		}
		value := RecipeFragranceSortCursorValue(order.Sort, m)
		if value != nil {
			a = append(a, boilergql.ToCursorValue(string(order.Sort), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(graphql_models.RecipeFragranceSortID), m.ID))
	}

	return boilergql.CursorValuesToString(a)
}

func RecipeFragranceCursorType(ordering []*graphql_models.RecipeFragranceOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func RecipeFragranceCursorMods(ordering []*graphql_models.RecipeFragranceOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if RecipeFragranceCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromRecipeFragranceCursor(*cursor, sign)
		}
		return boilergql.FromOffsetCursor(*cursor)
	}
	return nil
}

func RecipeFragranceSortMods(ordering []*graphql_models.RecipeFragranceOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, order := range ordering {
		if order.Sort == graphql_models.RecipeFragranceSortID {
			handledID = true
		}
		column := RecipeFragranceSortColumn[order.Sort]
		if column != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				column,
				boilergql.GetDirection(order.Direction, reverse),
			)))
		}
	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			models.RecipeFragranceColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func RecipeFragrancePaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*graphql_models.RecipeFragranceOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := RecipeFragranceSortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, RecipeFragranceCursorMods(ordering, cursor, sign)...)
	mods = append(mods, RecipeFragranceSortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func RecipeFragrancePaginationMods(pagination boilergql.ConnectionPagination, ordering []*graphql_models.RecipeFragranceOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	reverse := pagination.Backward != nil
	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	_, mods := RecipeFragrancePaginationModsBase(pagination, ordering, reverse, limit)
	return mods, nil
}

func ToRecipeFragranceCursorSwitch(ordering []*graphql_models.RecipeFragranceOrdering, m *graphql_models.RecipeFragrance, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToRecipeFragranceCursor(ordering, m)
	}
	return ""
}

func RecipeFragranceReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.RecipeFragranceOrdering,
) (bool, error) {
	reverse := pagination.Forward != nil
	cursor, reverseMods := RecipeFragrancePaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := RecipeFragranceCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return models.RecipeFragrances(reverseMods...).Count(ctx, db)
	})
}

func RecipeFragranceEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*graphql_models.RecipeFragranceOrdering) func(*models.RecipeFragrance, int) *graphql_models.RecipeFragranceEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), RecipeFragranceCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *models.RecipeFragrance, i int) *graphql_models.RecipeFragranceEdge {
		n := RecipeFragranceToGraphQL(m)
		return &graphql_models.RecipeFragranceEdge{
			Cursor:	ToRecipeFragranceCursorSwitch(ordering, n, cursorType, offset, i),
			Node:	n,
		}
	}
}

func RecipeFragranceStartEndCursor(edges []*graphql_models.RecipeFragranceEdge) (*string, *string) {
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

func RecipeFragranceConnection(
	ctx context.Context,
	db *sql.DB,
	originalMods []qm.QueryMod,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.RecipeFragranceOrdering,
) (*graphql_models.RecipeFragranceConnection, error) {
	paginationMods, err := RecipeFragrancePaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := RecipeFragranceReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := models.RecipeFragrances(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*graphql_models.RecipeFragranceEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := RecipeFragranceEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := RecipeFragranceStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
	return &graphql_models.RecipeFragranceConnection{
		Edges:	edges,
		PageInfo: &graphql_models.PageInfo{
			HasNextPage:		hasNextPage,
			HasPreviousPage:	hasPreviousPage,
			StartCursor:		startCursor,
			EndCursor:		endCursor,
		},
	}, nil
}

var RecipeLipidSortColumn = map[graphql_models.RecipeLipidSort]string{
	graphql_models.RecipeLipidSortID:		models.RecipeLipidColumns.ID,
	graphql_models.RecipeLipidSortPercentage:	models.RecipeLipidColumns.Percentage,
	graphql_models.RecipeLipidSortCreatedAt:	models.RecipeLipidColumns.CreatedAt,
	graphql_models.RecipeLipidSortUpdatedAt:	models.RecipeLipidColumns.UpdatedAt,
	graphql_models.RecipeLipidSortDeletedAt:	models.RecipeLipidColumns.DeletedAt,
}

func RecipeLipidSortValueFromCursorValue(cursorValue string) (string, interface{}) {
	key, value := boilergql.FromCursorValue(cursorValue)
	column := RecipeLipidSortColumn[graphql_models.RecipeLipidSort(key)]

	if graphql_models.RecipeLipidSort(key) == graphql_models.RecipeLipidSortID {
		return column, boilergql.GetIDFromCursor(value)
	}

	return column, boilergql.StringToInterface(value)
}

func RecipeLipidSortCursorValue(sort graphql_models.RecipeLipidSort, m *graphql_models.RecipeLipid) interface{} {
	switch sort {
	case graphql_models.RecipeLipidSortID:
		return m.ID
	case graphql_models.RecipeLipidSortPercentage:
		return m.Percentage
	case graphql_models.RecipeLipidSortCreatedAt:
		return m.CreatedAt
	case graphql_models.RecipeLipidSortUpdatedAt:
		return m.UpdatedAt
	case graphql_models.RecipeLipidSortDeletedAt:
		return m.DeletedAt
	}
	return nil
}

func RecipeLipidSortDirection(ordering []*graphql_models.RecipeLipidOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func FromRecipeLipidCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	var columns []string
	var values []interface{}

	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := RecipeLipidSortValueFromCursorValue(cursorValue)
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

func ToRecipeLipidCursor(ordering []*graphql_models.RecipeLipidOrdering, m *graphql_models.RecipeLipid) string {
	var a []string
	var handledID bool

	for _, order := range ordering {
		if order.Sort == graphql_models.RecipeLipidSortID {
			handledID = true
		}
		value := RecipeLipidSortCursorValue(order.Sort, m)
		if value != nil {
			a = append(a, boilergql.ToCursorValue(string(order.Sort), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(graphql_models.RecipeLipidSortID), m.ID))
	}

	return boilergql.CursorValuesToString(a)
}

func RecipeLipidCursorType(ordering []*graphql_models.RecipeLipidOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func RecipeLipidCursorMods(ordering []*graphql_models.RecipeLipidOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if RecipeLipidCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromRecipeLipidCursor(*cursor, sign)
		}
		return boilergql.FromOffsetCursor(*cursor)
	}
	return nil
}

func RecipeLipidSortMods(ordering []*graphql_models.RecipeLipidOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, order := range ordering {
		if order.Sort == graphql_models.RecipeLipidSortID {
			handledID = true
		}
		column := RecipeLipidSortColumn[order.Sort]
		if column != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				column,
				boilergql.GetDirection(order.Direction, reverse),
			)))
		}
	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			models.RecipeLipidColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func RecipeLipidPaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*graphql_models.RecipeLipidOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := RecipeLipidSortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, RecipeLipidCursorMods(ordering, cursor, sign)...)
	mods = append(mods, RecipeLipidSortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func RecipeLipidPaginationMods(pagination boilergql.ConnectionPagination, ordering []*graphql_models.RecipeLipidOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	reverse := pagination.Backward != nil
	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	_, mods := RecipeLipidPaginationModsBase(pagination, ordering, reverse, limit)
	return mods, nil
}

func ToRecipeLipidCursorSwitch(ordering []*graphql_models.RecipeLipidOrdering, m *graphql_models.RecipeLipid, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToRecipeLipidCursor(ordering, m)
	}
	return ""
}

func RecipeLipidReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.RecipeLipidOrdering,
) (bool, error) {
	reverse := pagination.Forward != nil
	cursor, reverseMods := RecipeLipidPaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := RecipeLipidCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return models.RecipeLipids(reverseMods...).Count(ctx, db)
	})
}

func RecipeLipidEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*graphql_models.RecipeLipidOrdering) func(*models.RecipeLipid, int) *graphql_models.RecipeLipidEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), RecipeLipidCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *models.RecipeLipid, i int) *graphql_models.RecipeLipidEdge {
		n := RecipeLipidToGraphQL(m)
		return &graphql_models.RecipeLipidEdge{
			Cursor:	ToRecipeLipidCursorSwitch(ordering, n, cursorType, offset, i),
			Node:	n,
		}
	}
}

func RecipeLipidStartEndCursor(edges []*graphql_models.RecipeLipidEdge) (*string, *string) {
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

func RecipeLipidConnection(
	ctx context.Context,
	db *sql.DB,
	originalMods []qm.QueryMod,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.RecipeLipidOrdering,
) (*graphql_models.RecipeLipidConnection, error) {
	paginationMods, err := RecipeLipidPaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := RecipeLipidReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := models.RecipeLipids(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*graphql_models.RecipeLipidEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := RecipeLipidEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := RecipeLipidStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
	return &graphql_models.RecipeLipidConnection{
		Edges:	edges,
		PageInfo: &graphql_models.PageInfo{
			HasNextPage:		hasNextPage,
			HasPreviousPage:	hasPreviousPage,
			StartCursor:		startCursor,
			EndCursor:		endCursor,
		},
	}, nil
}

var RecipeSortColumn = map[graphql_models.RecipeSort]string{
	graphql_models.RecipeSortID:		models.RecipeColumns.ID,
	graphql_models.RecipeSortName:		models.RecipeColumns.Name,
	graphql_models.RecipeSortNote:		models.RecipeColumns.Note,
	graphql_models.RecipeSortCreatedAt:	models.RecipeColumns.CreatedAt,
	graphql_models.RecipeSortUpdatedAt:	models.RecipeColumns.UpdatedAt,
	graphql_models.RecipeSortDeletedAt:	models.RecipeColumns.DeletedAt,
}

func RecipeSortValueFromCursorValue(cursorValue string) (string, interface{}) {
	key, value := boilergql.FromCursorValue(cursorValue)
	column := RecipeSortColumn[graphql_models.RecipeSort(key)]

	if graphql_models.RecipeSort(key) == graphql_models.RecipeSortID {
		return column, boilergql.GetIDFromCursor(value)
	}

	return column, boilergql.StringToInterface(value)
}

func RecipeSortCursorValue(sort graphql_models.RecipeSort, m *graphql_models.Recipe) interface{} {
	switch sort {
	case graphql_models.RecipeSortID:
		return m.ID
	case graphql_models.RecipeSortName:
		return m.Name
	case graphql_models.RecipeSortNote:
		return m.Note
	case graphql_models.RecipeSortCreatedAt:
		return m.CreatedAt
	case graphql_models.RecipeSortUpdatedAt:
		return m.UpdatedAt
	case graphql_models.RecipeSortDeletedAt:
		return m.DeletedAt
	}
	return nil
}

func RecipeSortDirection(ordering []*graphql_models.RecipeOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func FromRecipeCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	var columns []string
	var values []interface{}

	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := RecipeSortValueFromCursorValue(cursorValue)
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

func ToRecipeCursor(ordering []*graphql_models.RecipeOrdering, m *graphql_models.Recipe) string {
	var a []string
	var handledID bool

	for _, order := range ordering {
		if order.Sort == graphql_models.RecipeSortID {
			handledID = true
		}
		value := RecipeSortCursorValue(order.Sort, m)
		if value != nil {
			a = append(a, boilergql.ToCursorValue(string(order.Sort), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(graphql_models.RecipeSortID), m.ID))
	}

	return boilergql.CursorValuesToString(a)
}

func RecipeCursorType(ordering []*graphql_models.RecipeOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func RecipeCursorMods(ordering []*graphql_models.RecipeOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if RecipeCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromRecipeCursor(*cursor, sign)
		}
		return boilergql.FromOffsetCursor(*cursor)
	}
	return nil
}

func RecipeSortMods(ordering []*graphql_models.RecipeOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, order := range ordering {
		if order.Sort == graphql_models.RecipeSortID {
			handledID = true
		}
		column := RecipeSortColumn[order.Sort]
		if column != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				column,
				boilergql.GetDirection(order.Direction, reverse),
			)))
		}
	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			models.RecipeColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func RecipePaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*graphql_models.RecipeOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := RecipeSortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, RecipeCursorMods(ordering, cursor, sign)...)
	mods = append(mods, RecipeSortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func RecipePaginationMods(pagination boilergql.ConnectionPagination, ordering []*graphql_models.RecipeOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	reverse := pagination.Backward != nil
	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	_, mods := RecipePaginationModsBase(pagination, ordering, reverse, limit)
	return mods, nil
}

func ToRecipeCursorSwitch(ordering []*graphql_models.RecipeOrdering, m *graphql_models.Recipe, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToRecipeCursor(ordering, m)
	}
	return ""
}

func RecipeReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.RecipeOrdering,
) (bool, error) {
	reverse := pagination.Forward != nil
	cursor, reverseMods := RecipePaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := RecipeCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return models.Recipes(reverseMods...).Count(ctx, db)
	})
}

func RecipeEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*graphql_models.RecipeOrdering) func(*models.Recipe, int) *graphql_models.RecipeEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), RecipeCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *models.Recipe, i int) *graphql_models.RecipeEdge {
		n := RecipeToGraphQL(m)
		return &graphql_models.RecipeEdge{
			Cursor:	ToRecipeCursorSwitch(ordering, n, cursorType, offset, i),
			Node:	n,
		}
	}
}

func RecipeStartEndCursor(edges []*graphql_models.RecipeEdge) (*string, *string) {
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

func RecipeConnection(
	ctx context.Context,
	db *sql.DB,
	originalMods []qm.QueryMod,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.RecipeOrdering,
) (*graphql_models.RecipeConnection, error) {
	paginationMods, err := RecipePaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := RecipeReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := models.Recipes(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*graphql_models.RecipeEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := RecipeEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := RecipeStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
	return &graphql_models.RecipeConnection{
		Edges:	edges,
		PageInfo: &graphql_models.PageInfo{
			HasNextPage:		hasNextPage,
			HasPreviousPage:	hasPreviousPage,
			StartCursor:		startCursor,
			EndCursor:		endCursor,
		},
	}, nil
}

var RecipeStepSortColumn = map[graphql_models.RecipeStepSort]string{
	graphql_models.RecipeStepSortID:	models.RecipeStepColumns.ID,
	graphql_models.RecipeStepSortNum:	models.RecipeStepColumns.Num,
	graphql_models.RecipeStepSortNote:	models.RecipeStepColumns.Note,
	graphql_models.RecipeStepSortCreatedAt:	models.RecipeStepColumns.CreatedAt,
	graphql_models.RecipeStepSortUpdatedAt:	models.RecipeStepColumns.UpdatedAt,
	graphql_models.RecipeStepSortDeletedAt:	models.RecipeStepColumns.DeletedAt,
}

func RecipeStepSortValueFromCursorValue(cursorValue string) (string, interface{}) {
	key, value := boilergql.FromCursorValue(cursorValue)
	column := RecipeStepSortColumn[graphql_models.RecipeStepSort(key)]

	if graphql_models.RecipeStepSort(key) == graphql_models.RecipeStepSortID {
		return column, boilergql.GetIDFromCursor(value)
	}

	return column, boilergql.StringToInterface(value)
}

func RecipeStepSortCursorValue(sort graphql_models.RecipeStepSort, m *graphql_models.RecipeStep) interface{} {
	switch sort {
	case graphql_models.RecipeStepSortID:
		return m.ID
	case graphql_models.RecipeStepSortNum:
		return m.Num
	case graphql_models.RecipeStepSortNote:
		return m.Note
	case graphql_models.RecipeStepSortCreatedAt:
		return m.CreatedAt
	case graphql_models.RecipeStepSortUpdatedAt:
		return m.UpdatedAt
	case graphql_models.RecipeStepSortDeletedAt:
		return m.DeletedAt
	}
	return nil
}

func RecipeStepSortDirection(ordering []*graphql_models.RecipeStepOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func FromRecipeStepCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	var columns []string
	var values []interface{}

	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := RecipeStepSortValueFromCursorValue(cursorValue)
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

func ToRecipeStepCursor(ordering []*graphql_models.RecipeStepOrdering, m *graphql_models.RecipeStep) string {
	var a []string
	var handledID bool

	for _, order := range ordering {
		if order.Sort == graphql_models.RecipeStepSortID {
			handledID = true
		}
		value := RecipeStepSortCursorValue(order.Sort, m)
		if value != nil {
			a = append(a, boilergql.ToCursorValue(string(order.Sort), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(graphql_models.RecipeStepSortID), m.ID))
	}

	return boilergql.CursorValuesToString(a)
}

func RecipeStepCursorType(ordering []*graphql_models.RecipeStepOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func RecipeStepCursorMods(ordering []*graphql_models.RecipeStepOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if RecipeStepCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromRecipeStepCursor(*cursor, sign)
		}
		return boilergql.FromOffsetCursor(*cursor)
	}
	return nil
}

func RecipeStepSortMods(ordering []*graphql_models.RecipeStepOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, order := range ordering {
		if order.Sort == graphql_models.RecipeStepSortID {
			handledID = true
		}
		column := RecipeStepSortColumn[order.Sort]
		if column != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				column,
				boilergql.GetDirection(order.Direction, reverse),
			)))
		}
	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			models.RecipeStepColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func RecipeStepPaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*graphql_models.RecipeStepOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := RecipeStepSortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, RecipeStepCursorMods(ordering, cursor, sign)...)
	mods = append(mods, RecipeStepSortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func RecipeStepPaginationMods(pagination boilergql.ConnectionPagination, ordering []*graphql_models.RecipeStepOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	reverse := pagination.Backward != nil
	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	_, mods := RecipeStepPaginationModsBase(pagination, ordering, reverse, limit)
	return mods, nil
}

func ToRecipeStepCursorSwitch(ordering []*graphql_models.RecipeStepOrdering, m *graphql_models.RecipeStep, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToRecipeStepCursor(ordering, m)
	}
	return ""
}

func RecipeStepReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.RecipeStepOrdering,
) (bool, error) {
	reverse := pagination.Forward != nil
	cursor, reverseMods := RecipeStepPaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := RecipeStepCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return models.RecipeSteps(reverseMods...).Count(ctx, db)
	})
}

func RecipeStepEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*graphql_models.RecipeStepOrdering) func(*models.RecipeStep, int) *graphql_models.RecipeStepEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), RecipeStepCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *models.RecipeStep, i int) *graphql_models.RecipeStepEdge {
		n := RecipeStepToGraphQL(m)
		return &graphql_models.RecipeStepEdge{
			Cursor:	ToRecipeStepCursorSwitch(ordering, n, cursorType, offset, i),
			Node:	n,
		}
	}
}

func RecipeStepStartEndCursor(edges []*graphql_models.RecipeStepEdge) (*string, *string) {
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

func RecipeStepConnection(
	ctx context.Context,
	db *sql.DB,
	originalMods []qm.QueryMod,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.RecipeStepOrdering,
) (*graphql_models.RecipeStepConnection, error) {
	paginationMods, err := RecipeStepPaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := RecipeStepReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := models.RecipeSteps(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*graphql_models.RecipeStepEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := RecipeStepEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := RecipeStepStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
	return &graphql_models.RecipeStepConnection{
		Edges:	edges,
		PageInfo: &graphql_models.PageInfo{
			HasNextPage:		hasNextPage,
			HasPreviousPage:	hasPreviousPage,
			StartCursor:		startCursor,
			EndCursor:		endCursor,
		},
	}, nil
}

var SupplierSortColumn = map[graphql_models.SupplierSort]string{
	graphql_models.SupplierSortID:		models.SupplierColumns.ID,
	graphql_models.SupplierSortName:	models.SupplierColumns.Name,
	graphql_models.SupplierSortWebsite:	models.SupplierColumns.Website,
	graphql_models.SupplierSortNote:	models.SupplierColumns.Note,
	graphql_models.SupplierSortCreatedAt:	models.SupplierColumns.CreatedAt,
	graphql_models.SupplierSortUpdatedAt:	models.SupplierColumns.UpdatedAt,
	graphql_models.SupplierSortDeletedAt:	models.SupplierColumns.DeletedAt,
}

func SupplierSortValueFromCursorValue(cursorValue string) (string, interface{}) {
	key, value := boilergql.FromCursorValue(cursorValue)
	column := SupplierSortColumn[graphql_models.SupplierSort(key)]

	if graphql_models.SupplierSort(key) == graphql_models.SupplierSortID {
		return column, boilergql.GetIDFromCursor(value)
	}

	return column, boilergql.StringToInterface(value)
}

func SupplierSortCursorValue(sort graphql_models.SupplierSort, m *graphql_models.Supplier) interface{} {
	switch sort {
	case graphql_models.SupplierSortID:
		return m.ID
	case graphql_models.SupplierSortName:
		return m.Name
	case graphql_models.SupplierSortWebsite:
		return m.Website
	case graphql_models.SupplierSortNote:
		return m.Note
	case graphql_models.SupplierSortCreatedAt:
		return m.CreatedAt
	case graphql_models.SupplierSortUpdatedAt:
		return m.UpdatedAt
	case graphql_models.SupplierSortDeletedAt:
		return m.DeletedAt
	}
	return nil
}

func SupplierSortDirection(ordering []*graphql_models.SupplierOrdering) boilergql.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return boilergql.SortDirectionAsc
}

func FromSupplierCursor(cursor string, comparisonSign boilergql.ComparisonSign) []qm.QueryMod {
	var columns []string
	var values []interface{}

	for _, cursorValue := range boilergql.CursorStringToValues(cursor) {
		column, value := SupplierSortValueFromCursorValue(cursorValue)
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

func ToSupplierCursor(ordering []*graphql_models.SupplierOrdering, m *graphql_models.Supplier) string {
	var a []string
	var handledID bool

	for _, order := range ordering {
		if order.Sort == graphql_models.SupplierSortID {
			handledID = true
		}
		value := SupplierSortCursorValue(order.Sort, m)
		if value != nil {
			a = append(a, boilergql.ToCursorValue(string(order.Sort), value))
		}
	}

	if !handledID {
		a = append(a, boilergql.ToCursorValue(string(graphql_models.SupplierSortID), m.ID))
	}

	return boilergql.CursorValuesToString(a)
}

func SupplierCursorType(ordering []*graphql_models.SupplierOrdering) boilergql.CursorType {
	countDirection, result := boilergql.CursorTypeCounter()
	for _, o := range ordering {
		countDirection(o.Direction)
	}
	return result()
}

func SupplierCursorMods(ordering []*graphql_models.SupplierOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if SupplierCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromSupplierCursor(*cursor, sign)
		}
		return boilergql.FromOffsetCursor(*cursor)
	}
	return nil
}

func SupplierSortMods(ordering []*graphql_models.SupplierOrdering, reverse bool, defaultDirection boilergql.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	var handledID bool
	for _, order := range ordering {
		if order.Sort == graphql_models.SupplierSortID {
			handledID = true
		}
		column := SupplierSortColumn[order.Sort]
		if column != "" {
			a = append(a, qm.OrderBy(boilergql.GetOrderBy(
				column,
				boilergql.GetDirection(order.Direction, reverse),
			)))
		}
	}
	if !handledID {
		a = append(a, qm.OrderBy(boilergql.GetOrderBy(
			models.SupplierColumns.ID,
			boilergql.GetDirection(defaultDirection, reverse),
		)))
	}
	return a
}

func SupplierPaginationModsBase(pagination boilergql.ConnectionPagination, ordering []*graphql_models.SupplierOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := SupplierSortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, SupplierCursorMods(ordering, cursor, sign)...)
	mods = append(mods, SupplierSortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func SupplierPaginationMods(pagination boilergql.ConnectionPagination, ordering []*graphql_models.SupplierOrdering) ([]qm.QueryMod, error) {
	if pagination.Forward != nil && pagination.Backward != nil {
		return nil, errors.New("can not use forward and backward pagination at once")
	}
	if pagination.Forward == nil && pagination.Backward == nil {
		return nil, errors.New("no forward or backward pagination provided")
	}

	reverse := pagination.Backward != nil
	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	_, mods := SupplierPaginationModsBase(pagination, ordering, reverse, limit)
	return mods, nil
}

func ToSupplierCursorSwitch(ordering []*graphql_models.SupplierOrdering, m *graphql_models.Supplier, cursorType boilergql.CursorType, offset int, index int) string {
	switch cursorType {
	case boilergql.CursorTypeOffset:
		return boilergql.ToOffsetCursor(offset + index)
	case boilergql.CursorTypeCursor:
		return ToSupplierCursor(ordering, m)
	}
	return ""
}

func SupplierReversePageInformation(
	ctx context.Context,
	db *sql.DB,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.SupplierOrdering,
) (bool, error) {
	reverse := pagination.Forward != nil
	cursor, reverseMods := SupplierPaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := SupplierCursorType(ordering)
	return boilergql.HasReversePage(cursor, pagination, cursorType, func() (int64, error) {
		return models.Suppliers(reverseMods...).Count(ctx, db)
	})
}

func SupplierEdgeConverter(pagination boilergql.ConnectionPagination, ordering []*graphql_models.SupplierOrdering) func(*models.Supplier, int) *graphql_models.SupplierEdge {
	cursor, cursorType := boilergql.GetCursor(pagination.Forward, pagination.Backward), SupplierCursorType(ordering)
	offset := boilergql.GetOffsetFromCursor(cursor)
	return func(m *models.Supplier, i int) *graphql_models.SupplierEdge {
		n := SupplierToGraphQL(m)
		return &graphql_models.SupplierEdge{
			Cursor:	ToSupplierCursorSwitch(ordering, n, cursorType, offset, i),
			Node:	n,
		}
	}
}

func SupplierStartEndCursor(edges []*graphql_models.SupplierEdge) (*string, *string) {
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

func SupplierConnection(
	ctx context.Context,
	db *sql.DB,
	originalMods []qm.QueryMod,
	pagination boilergql.ConnectionPagination,
	ordering []*graphql_models.SupplierOrdering,
) (*graphql_models.SupplierConnection, error) {
	paginationMods, err := SupplierPaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasMoreReversed, err := SupplierReversePageInformation(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := models.Suppliers(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}
	edges := make([]*graphql_models.SupplierEdge, 0, boilergql.EdgeLength(pagination, len(a)))
	edgeConverter := SupplierEdgeConverter(pagination, ordering)
	hasMore := boilergql.BaseConnection(pagination, len(a), func(i int) {
		edges = append(edges, edgeConverter(a[i], i))
	})
	startCursor, endCursor := SupplierStartEndCursor(edges)
	hasNextPage, hasPreviousPage := boilergql.HasNextAndPreviousPage(pagination, hasMore, hasMoreReversed)
	return &graphql_models.SupplierConnection{
		Edges:	edges,
		PageInfo: &graphql_models.PageInfo{
			HasNextPage:		hasNextPage,
			HasPreviousPage:	hasPreviousPage,
			StartCursor:		startCursor,
			EndCursor:		endCursor,
		},
	}, nil
}
