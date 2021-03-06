package helpers

import (
	"fmt"
	"strings"

	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/web-ridge/gqlgen-sqlboiler-examples/issue-6-edges-connections/graphql_models"
	"github.com/web-ridge/gqlgen-sqlboiler-examples/issue-6-edges-connections/models"
	"github.com/web-ridge/utils-go/boilergql/v3"
)

const percentSign = `%`

func startsWithValue(v string) string	{ return v + percentSign }
func endsWithValue(v string) string	{ return percentSign + v }
func containsValue(v string) string	{ return percentSign + v + percentSign }

const isLike = " LIKE ?"
const in = " IN ?"
const notIn = " NOT IN ?"

func appendSubQuery(queryMods []qm.QueryMod, q *queries.Query) []qm.QueryMod {

	qs, args := queries.BuildQuery(q)
	qsClean := strings.TrimSuffix(qs, ";")
	return append(queryMods, qm.Where(fmt.Sprintf("EXISTS(%v)", qsClean), args...))
}

func BooleanFilterToMods(m *graphql_models.BooleanFilter, column string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod
	if m.EqualTo != nil {
		queryMods = append(queryMods, qmhelper.Where(column, qmhelper.EQ, *m.EqualTo))
	}
	if m.NotEqualTo != nil {
		queryMods = append(queryMods, qmhelper.Where(column, qmhelper.NEQ, *m.NotEqualTo))
	}
	return queryMods
}

func IDFilterToMods(m *graphql_models.IDFilter, column string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod
	if m.EqualTo != nil {
		queryMods = append(queryMods, qmhelper.Where(column, qmhelper.EQ, boilergql.IDToBoiler(*m.EqualTo)))
	}
	if m.NotEqualTo != nil {
		queryMods = append(queryMods, qmhelper.Where(column, qmhelper.NEQ, boilergql.IDToBoiler(*m.NotEqualTo)))
	}
	if len(m.In) > 0 {
		queryMods = append(queryMods, qm.WhereIn(column+in, boilergql.IDsToBoilerInterfaces(m.In)...))
	}
	if len(m.NotIn) > 0 {
		queryMods = append(queryMods, qm.WhereIn(column+notIn, boilergql.IDsToBoilerInterfaces(m.NotIn)...))
	}
	return queryMods
}

func StringFilterToMods(m *graphql_models.StringFilter, column string) []qm.QueryMod {
	if m == nil {
		return nil
	}

	var queryMods []qm.QueryMod
	if m.EqualTo != nil {
		queryMods = append(queryMods, qmhelper.Where(column, qmhelper.EQ, *m.EqualTo))
	}
	if m.NotEqualTo != nil {
		queryMods = append(queryMods, qmhelper.Where(column, qmhelper.NEQ, *m.NotEqualTo))
	}

	lowerColumn := "LOWER(" + column + ")"
	if m.StartWith != nil {
		queryMods = append(queryMods, qm.Where(lowerColumn+isLike, startsWithValue(strings.ToLower(*m.StartWith))))
	}
	if m.EndWith != nil {
		queryMods = append(queryMods, qm.Where(lowerColumn+isLike, endsWithValue(strings.ToLower(*m.EndWith))))
	}
	if m.Contain != nil {
		queryMods = append(queryMods, qm.Where(lowerColumn+isLike, containsValue(strings.ToLower(*m.Contain))))
	}

	if m.StartWithStrict != nil {
		queryMods = append(queryMods, qm.Where(column+isLike, startsWithValue(*m.StartWithStrict)))
	}
	if m.EndWithStrict != nil {
		queryMods = append(queryMods, qm.Where(column+isLike, endsWithValue(*m.EndWithStrict)))
	}
	if m.ContainStrict != nil {
		queryMods = append(queryMods, qm.Where(column+isLike, containsValue(*m.ContainStrict)))
	}

	if len(m.In) > 0 {
		queryMods = append(queryMods, qm.WhereIn(column+in, boilergql.IDsToBoilerInterfaces(m.In)...))
	}
	if len(m.NotIn) > 0 {
		queryMods = append(queryMods, qm.WhereIn(column+notIn, boilergql.IDsToBoilerInterfaces(m.NotIn)...))
	}

	return queryMods
}

func FloatFilterToMods(m *graphql_models.FloatFilter, column string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod
	if m.EqualTo != nil {
		queryMods = append(queryMods, qmhelper.Where(column, qmhelper.EQ, *m.EqualTo))
	}
	if m.NotEqualTo != nil {
		queryMods = append(queryMods, qmhelper.Where(column, qmhelper.NEQ, *m.NotEqualTo))
	}
	if m.LessThan != nil {
		queryMods = append(queryMods, qmhelper.Where(column, qmhelper.LT, *m.LessThan))
	}
	if m.MoreThan != nil {
		queryMods = append(queryMods, qmhelper.Where(column, qmhelper.GT, *m.MoreThan))
	}
	if m.LessThanOrEqualTo != nil {
		queryMods = append(queryMods, qmhelper.Where(column, qmhelper.LTE, *m.LessThanOrEqualTo))
	}
	if m.MoreThanOrEqualTo != nil {
		queryMods = append(queryMods, qmhelper.Where(column, qmhelper.GTE, *m.MoreThanOrEqualTo))
	}
	if len(m.In) > 0 {
		queryMods = append(queryMods, qm.WhereIn(column+in, boilergql.FloatsToInterfaces(m.In)...))
	}
	if len(m.NotIn) > 0 {
		queryMods = append(queryMods, qm.WhereIn(column+notIn, boilergql.FloatsToInterfaces(m.NotIn)...))
	}
	return queryMods
}

func IntFilterToMods(m *graphql_models.IntFilter, column string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod
	if m.EqualTo != nil {
		queryMods = append(queryMods, qmhelper.Where(column, qmhelper.EQ, *m.EqualTo))
	}
	if m.NotEqualTo != nil {
		queryMods = append(queryMods, qmhelper.Where(column, qmhelper.NEQ, *m.NotEqualTo))
	}
	if m.LessThan != nil {
		queryMods = append(queryMods, qmhelper.Where(column, qmhelper.LT, *m.LessThan))
	}
	if m.MoreThan != nil {
		queryMods = append(queryMods, qmhelper.Where(column, qmhelper.GT, *m.MoreThan))
	}
	if m.LessThanOrEqualTo != nil {
		queryMods = append(queryMods, qmhelper.Where(column, qmhelper.LTE, *m.LessThanOrEqualTo))
	}
	if m.MoreThanOrEqualTo != nil {
		queryMods = append(queryMods, qmhelper.Where(column, qmhelper.GTE, *m.MoreThanOrEqualTo))
	}
	if len(m.In) > 0 {
		queryMods = append(queryMods, qm.WhereIn(column+in, boilergql.IntsToInterfaces(m.In)...))
	}
	if len(m.NotIn) > 0 {
		queryMods = append(queryMods, qm.WhereIn(column+notIn, boilergql.IntsToInterfaces(m.NotIn)...))
	}
	return queryMods
}

func UserFilterToMods(m *graphql_models.UserFilter) []qm.QueryMod {
	if m == nil {
		return nil
	}
	if m.Search != nil || m.Where != nil {
		var queryMods []qm.QueryMod
		queryMods = append(queryMods, UserSearchToMods(m.Search)...)
		queryMods = append(queryMods, UserWhereToMods(m.Where, true, "")...)
		if len(queryMods) > 0 {
			return []qm.QueryMod{
				qm.Expr(queryMods...),
			}
		}
	}
	return nil
}
func UserSearchToMods(search *string) []qm.QueryMod {

	return nil
}
func UserWhereSubqueryToMods(m *graphql_models.UserWhere, foreignColumn string, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	hasForeignKeyInRoot := foreignColumn != ""
	if hasForeignKeyInRoot {
		queryMods = append(queryMods, IDFilterToMods(m.ID, foreignColumn)...)
	}

	subQueryMods := UserWhereToMods(m, !hasForeignKeyInRoot, parentTable)
	if len(subQueryMods) > 0 {
		subQuery := models.Users(append(subQueryMods, qm.Select("1"))...)
		queryMods = appendSubQuery(queryMods, subQuery.Query)
	}
	return queryMods
}

func UserWhereToMods(m *graphql_models.UserWhere, withPrimaryID bool, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	if withPrimaryID {
		queryMods = append(queryMods, IDFilterToMods(m.ID, models.UserColumns.ID)...)
	}
	queryMods = append(queryMods, StringFilterToMods(m.FirstName, models.UserColumns.FirstName)...)
	queryMods = append(queryMods, StringFilterToMods(m.LastName, models.UserColumns.LastName)...)
	queryMods = append(queryMods, IntFilterToMods(m.Age, models.UserColumns.Age)...)
	queryMods = append(queryMods, StringFilterToMods(m.Email, models.UserColumns.Email)...)
	if m.Or != nil {
		queryMods = append(queryMods, qm.Or2(qm.Expr(UserWhereToMods(m.Or, true, "")...)))
	}
	if m.And != nil {
		queryMods = append(queryMods, qm.Expr(UserWhereToMods(m.And, true, "")...))
	}

	if len(queryMods) > 0 && parentTable != "" {

	}

	return queryMods
}
