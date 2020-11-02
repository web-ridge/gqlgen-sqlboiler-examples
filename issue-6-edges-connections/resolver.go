// Generated with https://github.com/web-ridge/gqlgen-sqlboiler.
package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

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

func getSortDirection(ordering []*fm.UserOrdering) fm.SortDirection {
	for _, o := range ordering {
		return o.Direction
	}
	return fm.SortDirectionAsc
}

func getCursorType(ordering []*fm.UserOrdering) boilergql.CursorType {
	var asc, desc int
	for _, o := range ordering {
		switch o.Direction {
		case fm.SortDirectionDesc:
			desc++
		case fm.SortDirectionAsc:
			asc++
		}
	}

	oneDirectionOnly := asc == 0 || desc == 0
	if oneDirectionOnly {
		return boilergql.CursorTypeCursor
	}
	return boilergql.CursorTypeOffset
}

func UserCursorMods(ordering []*fm.UserOrdering, cursor *string, sign boilergql.ComparisonSign) []qm.QueryMod {
	if cursor != nil {
		if getCursorType(ordering) == boilergql.CursorTypeCursor {
			return FromUserCursor(*cursor, sign)
		}
		return FromOffsetCursor(*cursor)
	}
	return nil
}

func UserPaginationModsBase(pagination fm.ConnectionPagination, ordering []*fm.UserOrdering, reverse bool, limit int) (*string, []qm.QueryMod) {
	direction := getSortDirection(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	sign := boilergql.GetComparison(pagination.Forward, pagination.Backward, reverse, direction)

	var mods []qm.QueryMod
	mods = append(mods, UserCursorMods(ordering, cursor, sign)...)
	mods = append(mods, UserSortMods(ordering, reverse, direction)...)
	mods = append(mods, qm.Limit(limit))
	return cursor, mods
}

func UserPaginationMods(pagination fm.ConnectionPagination, ordering []*fm.UserOrdering) ([]qm.QueryMod, error) {
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

func getDirection(direction fm.SortDirection, reverse bool) fm.SortDirection {
	if reverse {
		if direction == fm.SortDirectionAsc {
			return fm.SortDirectionDesc
		}
		return fm.SortDirectionAsc
	}
	return direction
}

func UserSortMods(ordering []*fm.UserOrdering, reverse bool, defaultDirection fm.SortDirection) []qm.QueryMod {
	var a []qm.QueryMod

	idSortDirection := defaultDirection
	for _, column := range ordering {
		switch column.Sort {
		case fm.UserSortID:
			idSortDirection = column.Direction
		case fm.UserSortAge:
			a = append(a, qm.OrderBy(dm.UserColumns.Age+" "+getDirection(column.Direction, reverse).String()))
		case fm.UserSortFirstName:
			a = append(a, qm.OrderBy(dm.UserColumns.FirstName+" "+getDirection(column.Direction, reverse).String()))
		case fm.UserSortLastName:
			a = append(a, qm.OrderBy(dm.UserColumns.LastName+" "+getDirection(column.Direction, reverse).String()))
		case fm.UserSortEmail:
			a = append(a, qm.OrderBy(dm.UserColumns.Email+" "+getDirection(column.Direction, reverse).String()))
		}
	}
	a = append(a, qm.OrderBy(dm.UserColumns.ID+" "+getDirection(idSortDirection, reverse).String()))
	return a
}

const (
	Separator1 = ",____,"
	Separator2 = ":"
)

func ToUserCursorSwitch(ordering []*fm.UserOrdering, user *fm.User, cursorType CursorType, offset int, index int) string {
	switch cursorType {
	case CursorTypeOffset:
		return ToOffsetCursor(offset + index)
	case CursorTypeCursor:
		return ToUserCursor(ordering, user)
	}
	return ""
}

func ToOffsetCursor(index int) string {
	return strconv.Itoa(index + 1)
}

func ToUserCursor(ordering []*fm.UserOrdering, user *fm.User) string {
	var a []string

	for _, column := range ordering {
		switch column.Sort {
		// case fm.UserSortID:
		//	a = append(a, fmt.Sprintf("%v%v%v", fm.UserSortID, Separator2, user.ID))
		case fm.UserSortAge:
			a = append(a, fmt.Sprintf("%v%v%v", fm.UserSortAge, Separator2, user.Age))
		case fm.UserSortFirstName:
			a = append(a, fmt.Sprintf("%v%v%v", fm.UserSortFirstName, Separator2, user.FirstName))
		case fm.UserSortLastName:
			a = append(a, fmt.Sprintf("%v%v%v", fm.UserSortLastName, Separator2, user.LastName))
		case fm.UserSortEmail:
			a = append(a, fmt.Sprintf("%v%v%v", fm.UserSortEmail, Separator2, user.Email))
		}
	}

	a = append(a, fmt.Sprintf("%v%v%v", fm.UserSortID, Separator2, user.ID))

	// return base64.StdEncoding.EncodeToString([]byte(strings.Join(a, Separator1)))
	return strings.Join(a, Separator1)
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

func FromUserCursor(cursor string, comparisonSign ComparisonSign) []qm.QueryMod {
	// b, _ := base64.StdEncoding.DecodeString(cursor)
	b := cursor
	var columns []string
	var values []interface{}

	// https://www.slideshare.net/MarkusWinand/p2d2-pagination-done-the-postgresql-way
	for _, columnAndValue := range strings.Split(string(b), Separator1) {
		s := strings.SplitN(columnAndValue, Separator2, 2)
		if len(s) != 2 {
			continue
		}
		column := fm.UserSort(s[0])
		value := s[1]
		switch column {
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
			qm.Where(
				fmt.Sprintf("(%v) %v (%v)",
					strings.Join(columns, ", "),
					comparisonSign,
					// strings.Join(columns, ", ")
					strings.TrimSuffix(strings.Repeat("?,", len(values)), ","),
				), values...),
		}
	}
	return nil
}

func HasReversePage(ctx context.Context, db *sql.DB, pagination fm.ConnectionPagination, ordering []*fm.UserOrdering) (bool, bool, error) {
	// reverse query to check hasPrevious / hasNext
	reverse := pagination.Forward != nil
	cursor, reverseMods := UserPaginationModsBase(pagination, ordering, reverse, 1)
	cursorType := getCursorType(ordering)

	var hasNextPage, hasPreviousPage bool

	if cursor != nil {
		var hasReverseObjects bool

		if cursorType == boilergql.CursorTypeCursor {
			reverseCount, err := dm.Users(reverseMods...).Count(ctx, db)
			if err != nil {
				return false, false, err
			}
			hasReverseObjects = reverseCount > 0
		} else {
			hasReverseObjects = true
		}

		if hasReverseObjects {
			if pagination.Backward != nil {
				hasNextPage = true
			}
			if pagination.Forward != nil {
				hasPreviousPage = true
			}
		}
	}
	return hasNextPage, hasPreviousPage, nil
}

func UserEdgeConverter(pagination fm.ConnectionPagination, ordering []*fm.UserOrdering) func(*dm.User, int) *fm.UserEdge {
	cursorType := getCursorType(ordering)
	cursor := boilergql.GetCursor(pagination.Forward, pagination.Backward)
	offset := boilergql.GetOffsetFromCursor(cursor)

	return func(m *dm.User, i int) *fm.UserEdge {
		n := UserToGraphQL(m)
		return &fm.UserEdge{
			Cursor: ToUserCursorSwitch(ordering, n, cursorType, offset, i),
			Node:   n,
		}
	}
}

func ToUserConnection(ctx context.Context, db *sql.DB, originalMods []qm.QueryMod, pagination fm.ConnectionPagination, ordering []*fm.UserOrdering) (*fm.UserConnection, error) {
	paginationMods, err := UserPaginationMods(pagination, ordering)
	if err != nil {
		return nil, err
	}

	hasNextPage, hasPreviousPage, err := HasReversePage(ctx, db, pagination, ordering)
	if err != nil {
		return nil, err
	}

	a, err := dm.Users(append(originalMods, paginationMods...)...).All(ctx, db)
	if err != nil {
		return nil, err
	}

	limit := boilergql.GetLimit(pagination.Forward, pagination.Backward)
	maxLength := limit - 1
	lowestLength := math.Min(float64(len(a)), float64(maxLength))
	edges := make([]*fm.UserEdge, 0, int(lowestLength))
	edgeConverter := UserEdgeConverter(pagination, ordering)

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
			edges = append(edges, edgeConverter(a[i], i))
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
			edges = append(edges, edgeConverter(row, i))
		}
	}

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
