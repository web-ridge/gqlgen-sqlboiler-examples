package helpers

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"

	"github.com/volatiletech/sqlboiler/v4/drivers"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/web-ridge/gqlgen-sqlboiler-examples/issue-38-plugin-panics/graphql_models"
	"github.com/web-ridge/gqlgen-sqlboiler-examples/issue-38-plugin-panics/models"
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

	member := reflect.ValueOf(q).Elem().FieldByName("dialect")
	dialectPtr := (**drivers.Dialect)(unsafe.Pointer(member.UnsafeAddr()))
	dialect := **dialectPtr
	dialect.UseIndexPlaceholders = false
	*dialectPtr = &dialect

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

func AdditiveFilterToMods(m *graphql_models.AdditiveFilter) []qm.QueryMod {
	if m == nil {
		return nil
	}
	if m.Search != nil || m.Where != nil {
		var queryMods []qm.QueryMod
		queryMods = append(queryMods, AdditiveSearchToMods(m.Search)...)
		queryMods = append(queryMods, AdditiveWhereToMods(m.Where, true, "")...)
		if len(queryMods) > 0 {
			return []qm.QueryMod{
				qm.Expr(queryMods...),
			}
		}
	}
	return nil
}
func AdditiveSearchToMods(search *string) []qm.QueryMod {

	return nil
}
func AdditiveInventoryFilterToMods(m *graphql_models.AdditiveInventoryFilter) []qm.QueryMod {
	if m == nil {
		return nil
	}
	if m.Search != nil || m.Where != nil {
		var queryMods []qm.QueryMod
		queryMods = append(queryMods, AdditiveInventorySearchToMods(m.Search)...)
		queryMods = append(queryMods, AdditiveInventoryWhereToMods(m.Where, true, "")...)
		if len(queryMods) > 0 {
			return []qm.QueryMod{
				qm.Expr(queryMods...),
			}
		}
	}
	return nil
}
func AdditiveInventorySearchToMods(search *string) []qm.QueryMod {

	return nil
}
func AdditiveInventoryWhereSubqueryToMods(m *graphql_models.AdditiveInventoryWhere, foreignColumn string, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	hasForeignKeyInRoot := foreignColumn != ""
	if hasForeignKeyInRoot {
		queryMods = append(queryMods, IDFilterToMods(m.ID, foreignColumn)...)
	}

	subQueryMods := AdditiveInventoryWhereToMods(m, !hasForeignKeyInRoot, parentTable)
	if len(subQueryMods) > 0 {
		subQuery := models.AdditiveInventories(append(subQueryMods, qm.Select("1"))...)
		queryMods = appendSubQuery(queryMods, subQuery.Query)
	}
	return queryMods
}

func AdditiveInventoryWhereToMods(m *graphql_models.AdditiveInventoryWhere, withPrimaryID bool, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	if withPrimaryID {
		queryMods = append(queryMods, IDFilterToMods(m.ID, models.AdditiveInventoryColumns.ID)...)
	}
	queryMods = append(queryMods, IntFilterToMods(m.PurchaseDate, models.AdditiveInventoryColumns.PurchaseDate)...)
	queryMods = append(queryMods, IntFilterToMods(m.ExpiryDate, models.AdditiveInventoryColumns.ExpiryDate)...)
	queryMods = append(queryMods, FloatFilterToMods(m.Cost, models.AdditiveInventoryColumns.Cost)...)
	queryMods = append(queryMods, FloatFilterToMods(m.Weight, models.AdditiveInventoryColumns.Weight)...)
	queryMods = append(queryMods, AdditiveWhereSubqueryToMods(m.Additive, models.AdditiveInventoryColumns.AdditiveID, models.TableNames.AdditiveInventory)...)
	queryMods = append(queryMods, SupplierWhereSubqueryToMods(m.Supplier, models.AdditiveInventoryColumns.SupplierID, models.TableNames.AdditiveInventory)...)
	queryMods = append(queryMods, IntFilterToMods(m.CreatedAt, models.AdditiveInventoryColumns.CreatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.UpdatedAt, models.AdditiveInventoryColumns.UpdatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.DeletedAt, models.AdditiveInventoryColumns.DeletedAt)...)
	if m.Or != nil {
		queryMods = append(queryMods, qm.Or2(qm.Expr(AdditiveInventoryWhereToMods(m.Or, true, "")...)))
	}
	if m.And != nil {
		queryMods = append(queryMods, qm.Expr(AdditiveInventoryWhereToMods(m.And, true, "")...))
	}

	if len(queryMods) > 0 && parentTable != "" {

		if parentTable == models.TableNames.Additive {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.AdditiveInventory, models.AdditiveInventoryColumns.AdditiveID, parentTable)))
		}
		if parentTable == models.TableNames.Supplier {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.AdditiveInventory, models.AdditiveInventoryColumns.SupplierID, parentTable)))
		}
	}

	return queryMods
}
func AdditiveWhereSubqueryToMods(m *graphql_models.AdditiveWhere, foreignColumn string, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	hasForeignKeyInRoot := foreignColumn != ""
	if hasForeignKeyInRoot {
		queryMods = append(queryMods, IDFilterToMods(m.ID, foreignColumn)...)
	}

	subQueryMods := AdditiveWhereToMods(m, !hasForeignKeyInRoot, parentTable)
	if len(subQueryMods) > 0 {
		subQuery := models.Additives(append(subQueryMods, qm.Select("1"))...)
		queryMods = appendSubQuery(queryMods, subQuery.Query)
	}
	return queryMods
}

func AdditiveWhereToMods(m *graphql_models.AdditiveWhere, withPrimaryID bool, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	if withPrimaryID {
		queryMods = append(queryMods, IDFilterToMods(m.ID, models.AdditiveColumns.ID)...)
	}
	queryMods = append(queryMods, StringFilterToMods(m.Name, models.AdditiveColumns.Name)...)
	queryMods = append(queryMods, StringFilterToMods(m.Note, models.AdditiveColumns.Note)...)
	queryMods = append(queryMods, IntFilterToMods(m.CreatedAt, models.AdditiveColumns.CreatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.UpdatedAt, models.AdditiveColumns.UpdatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.DeletedAt, models.AdditiveColumns.DeletedAt)...)
	queryMods = append(queryMods, RecipeAdditiveWhereSubqueryToMods(m.RecipeAdditive, "", models.TableNames.Additive)...)
	queryMods = append(queryMods, RecipeBatchAdditiveWhereSubqueryToMods(m.RecipeBatchAdditive, "", models.TableNames.Additive)...)
	queryMods = append(queryMods, AdditiveInventoryWhereSubqueryToMods(m.AdditiveInventories, "", models.TableNames.Additive)...)
	if m.Or != nil {
		queryMods = append(queryMods, qm.Or2(qm.Expr(AdditiveWhereToMods(m.Or, true, "")...)))
	}
	if m.And != nil {
		queryMods = append(queryMods, qm.Expr(AdditiveWhereToMods(m.And, true, "")...))
	}

	if len(queryMods) > 0 && parentTable != "" {

	}

	return queryMods
}
func AuthGroupFilterToMods(m *graphql_models.AuthGroupFilter) []qm.QueryMod {
	if m == nil {
		return nil
	}
	if m.Search != nil || m.Where != nil {
		var queryMods []qm.QueryMod
		queryMods = append(queryMods, AuthGroupSearchToMods(m.Search)...)
		queryMods = append(queryMods, AuthGroupWhereToMods(m.Where, true, "")...)
		if len(queryMods) > 0 {
			return []qm.QueryMod{
				qm.Expr(queryMods...),
			}
		}
	}
	return nil
}
func AuthGroupSearchToMods(search *string) []qm.QueryMod {

	return nil
}
func AuthGroupPermissionFilterToMods(m *graphql_models.AuthGroupPermissionFilter) []qm.QueryMod {
	if m == nil {
		return nil
	}
	if m.Search != nil || m.Where != nil {
		var queryMods []qm.QueryMod
		queryMods = append(queryMods, AuthGroupPermissionSearchToMods(m.Search)...)
		queryMods = append(queryMods, AuthGroupPermissionWhereToMods(m.Where, true, "")...)
		if len(queryMods) > 0 {
			return []qm.QueryMod{
				qm.Expr(queryMods...),
			}
		}
	}
	return nil
}
func AuthGroupPermissionSearchToMods(search *string) []qm.QueryMod {

	return nil
}
func AuthGroupPermissionWhereSubqueryToMods(m *graphql_models.AuthGroupPermissionWhere, foreignColumn string, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	hasForeignKeyInRoot := foreignColumn != ""
	if hasForeignKeyInRoot {
		queryMods = append(queryMods, IDFilterToMods(m.ID, foreignColumn)...)
	}

	subQueryMods := AuthGroupPermissionWhereToMods(m, !hasForeignKeyInRoot, parentTable)
	if len(subQueryMods) > 0 {
		subQuery := models.AuthGroupPermissions(append(subQueryMods, qm.Select("1"))...)
		queryMods = appendSubQuery(queryMods, subQuery.Query)
	}
	return queryMods
}

func AuthGroupPermissionWhereToMods(m *graphql_models.AuthGroupPermissionWhere, withPrimaryID bool, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	if withPrimaryID {
		queryMods = append(queryMods, IDFilterToMods(m.ID, models.AuthGroupPermissionColumns.ID)...)
	}
	queryMods = append(queryMods, AuthGroupWhereSubqueryToMods(m.Group, models.AuthGroupPermissionColumns.GroupID, models.TableNames.AuthGroupPermissions)...)
	queryMods = append(queryMods, AuthPermissionWhereSubqueryToMods(m.Permission, models.AuthGroupPermissionColumns.PermissionID, models.TableNames.AuthGroupPermissions)...)
	if m.Or != nil {
		queryMods = append(queryMods, qm.Or2(qm.Expr(AuthGroupPermissionWhereToMods(m.Or, true, "")...)))
	}
	if m.And != nil {
		queryMods = append(queryMods, qm.Expr(AuthGroupPermissionWhereToMods(m.And, true, "")...))
	}

	if len(queryMods) > 0 && parentTable != "" {

		if parentTable == models.TableNames.AuthGroup {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.AuthGroupPermissions, models.AuthGroupPermissionColumns.GroupID, parentTable)))
		}
		if parentTable == models.TableNames.AuthPermission {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.AuthGroupPermissions, models.AuthGroupPermissionColumns.PermissionID, parentTable)))
		}
	}

	return queryMods
}
func AuthGroupWhereSubqueryToMods(m *graphql_models.AuthGroupWhere, foreignColumn string, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	hasForeignKeyInRoot := foreignColumn != ""
	if hasForeignKeyInRoot {
		queryMods = append(queryMods, IDFilterToMods(m.ID, foreignColumn)...)
	}

	subQueryMods := AuthGroupWhereToMods(m, !hasForeignKeyInRoot, parentTable)
	if len(subQueryMods) > 0 {
		subQuery := models.AuthGroups(append(subQueryMods, qm.Select("1"))...)
		queryMods = appendSubQuery(queryMods, subQuery.Query)
	}
	return queryMods
}

func AuthGroupWhereToMods(m *graphql_models.AuthGroupWhere, withPrimaryID bool, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	if withPrimaryID {
		queryMods = append(queryMods, IDFilterToMods(m.ID, models.AuthGroupColumns.ID)...)
	}
	queryMods = append(queryMods, StringFilterToMods(m.Name, models.AuthGroupColumns.Name)...)
	queryMods = append(queryMods, AuthGroupPermissionWhereSubqueryToMods(m.GroupAuthGroupPermissions, "", models.TableNames.AuthGroup)...)
	queryMods = append(queryMods, AuthUserGroupWhereSubqueryToMods(m.GroupAuthUserGroups, "", models.TableNames.AuthGroup)...)
	if m.Or != nil {
		queryMods = append(queryMods, qm.Or2(qm.Expr(AuthGroupWhereToMods(m.Or, true, "")...)))
	}
	if m.And != nil {
		queryMods = append(queryMods, qm.Expr(AuthGroupWhereToMods(m.And, true, "")...))
	}

	if len(queryMods) > 0 && parentTable != "" {

	}

	return queryMods
}
func AuthPermissionFilterToMods(m *graphql_models.AuthPermissionFilter) []qm.QueryMod {
	if m == nil {
		return nil
	}
	if m.Search != nil || m.Where != nil {
		var queryMods []qm.QueryMod
		queryMods = append(queryMods, AuthPermissionSearchToMods(m.Search)...)
		queryMods = append(queryMods, AuthPermissionWhereToMods(m.Where, true, "")...)
		if len(queryMods) > 0 {
			return []qm.QueryMod{
				qm.Expr(queryMods...),
			}
		}
	}
	return nil
}
func AuthPermissionSearchToMods(search *string) []qm.QueryMod {

	return nil
}
func AuthPermissionWhereSubqueryToMods(m *graphql_models.AuthPermissionWhere, foreignColumn string, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	hasForeignKeyInRoot := foreignColumn != ""
	if hasForeignKeyInRoot {
		queryMods = append(queryMods, IDFilterToMods(m.ID, foreignColumn)...)
	}

	subQueryMods := AuthPermissionWhereToMods(m, !hasForeignKeyInRoot, parentTable)
	if len(subQueryMods) > 0 {
		subQuery := models.AuthPermissions(append(subQueryMods, qm.Select("1"))...)
		queryMods = appendSubQuery(queryMods, subQuery.Query)
	}
	return queryMods
}

func AuthPermissionWhereToMods(m *graphql_models.AuthPermissionWhere, withPrimaryID bool, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	if withPrimaryID {
		queryMods = append(queryMods, IDFilterToMods(m.ID, models.AuthPermissionColumns.ID)...)
	}
	queryMods = append(queryMods, StringFilterToMods(m.Name, models.AuthPermissionColumns.Name)...)
	queryMods = append(queryMods, IDFilterToMods(m.ContentTypeID, models.AuthPermissionColumns.ContentTypeID)...)
	queryMods = append(queryMods, StringFilterToMods(m.Codename, models.AuthPermissionColumns.Codename)...)
	queryMods = append(queryMods, AuthGroupPermissionWhereSubqueryToMods(m.PermissionAuthGroupPermissions, "", models.TableNames.AuthPermission)...)
	queryMods = append(queryMods, AuthUserUserPermissionWhereSubqueryToMods(m.PermissionAuthUserUserPermissions, "", models.TableNames.AuthPermission)...)
	if m.Or != nil {
		queryMods = append(queryMods, qm.Or2(qm.Expr(AuthPermissionWhereToMods(m.Or, true, "")...)))
	}
	if m.And != nil {
		queryMods = append(queryMods, qm.Expr(AuthPermissionWhereToMods(m.And, true, "")...))
	}

	if len(queryMods) > 0 && parentTable != "" {

	}

	return queryMods
}
func AuthUserFilterToMods(m *graphql_models.AuthUserFilter) []qm.QueryMod {
	if m == nil {
		return nil
	}
	if m.Search != nil || m.Where != nil {
		var queryMods []qm.QueryMod
		queryMods = append(queryMods, AuthUserSearchToMods(m.Search)...)
		queryMods = append(queryMods, AuthUserWhereToMods(m.Where, true, "")...)
		if len(queryMods) > 0 {
			return []qm.QueryMod{
				qm.Expr(queryMods...),
			}
		}
	}
	return nil
}
func AuthUserSearchToMods(search *string) []qm.QueryMod {

	return nil
}
func AuthUserGroupFilterToMods(m *graphql_models.AuthUserGroupFilter) []qm.QueryMod {
	if m == nil {
		return nil
	}
	if m.Search != nil || m.Where != nil {
		var queryMods []qm.QueryMod
		queryMods = append(queryMods, AuthUserGroupSearchToMods(m.Search)...)
		queryMods = append(queryMods, AuthUserGroupWhereToMods(m.Where, true, "")...)
		if len(queryMods) > 0 {
			return []qm.QueryMod{
				qm.Expr(queryMods...),
			}
		}
	}
	return nil
}
func AuthUserGroupSearchToMods(search *string) []qm.QueryMod {

	return nil
}
func AuthUserGroupWhereSubqueryToMods(m *graphql_models.AuthUserGroupWhere, foreignColumn string, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	hasForeignKeyInRoot := foreignColumn != ""
	if hasForeignKeyInRoot {
		queryMods = append(queryMods, IDFilterToMods(m.ID, foreignColumn)...)
	}

	subQueryMods := AuthUserGroupWhereToMods(m, !hasForeignKeyInRoot, parentTable)
	if len(subQueryMods) > 0 {
		subQuery := models.AuthUserGroups(append(subQueryMods, qm.Select("1"))...)
		queryMods = appendSubQuery(queryMods, subQuery.Query)
	}
	return queryMods
}

func AuthUserGroupWhereToMods(m *graphql_models.AuthUserGroupWhere, withPrimaryID bool, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	if withPrimaryID {
		queryMods = append(queryMods, IDFilterToMods(m.ID, models.AuthUserGroupColumns.ID)...)
	}
	queryMods = append(queryMods, AuthUserWhereSubqueryToMods(m.User, models.AuthUserGroupColumns.UserID, models.TableNames.AuthUserGroups)...)
	queryMods = append(queryMods, AuthGroupWhereSubqueryToMods(m.Group, models.AuthUserGroupColumns.GroupID, models.TableNames.AuthUserGroups)...)
	if m.Or != nil {
		queryMods = append(queryMods, qm.Or2(qm.Expr(AuthUserGroupWhereToMods(m.Or, true, "")...)))
	}
	if m.And != nil {
		queryMods = append(queryMods, qm.Expr(AuthUserGroupWhereToMods(m.And, true, "")...))
	}

	if len(queryMods) > 0 && parentTable != "" {

		if parentTable == models.TableNames.AuthUser {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.AuthUserGroups, models.AuthUserGroupColumns.UserID, parentTable)))
		}
		if parentTable == models.TableNames.AuthGroup {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.AuthUserGroups, models.AuthUserGroupColumns.GroupID, parentTable)))
		}
	}

	return queryMods
}
func AuthUserUserPermissionFilterToMods(m *graphql_models.AuthUserUserPermissionFilter) []qm.QueryMod {
	if m == nil {
		return nil
	}
	if m.Search != nil || m.Where != nil {
		var queryMods []qm.QueryMod
		queryMods = append(queryMods, AuthUserUserPermissionSearchToMods(m.Search)...)
		queryMods = append(queryMods, AuthUserUserPermissionWhereToMods(m.Where, true, "")...)
		if len(queryMods) > 0 {
			return []qm.QueryMod{
				qm.Expr(queryMods...),
			}
		}
	}
	return nil
}
func AuthUserUserPermissionSearchToMods(search *string) []qm.QueryMod {

	return nil
}
func AuthUserUserPermissionWhereSubqueryToMods(m *graphql_models.AuthUserUserPermissionWhere, foreignColumn string, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	hasForeignKeyInRoot := foreignColumn != ""
	if hasForeignKeyInRoot {
		queryMods = append(queryMods, IDFilterToMods(m.ID, foreignColumn)...)
	}

	subQueryMods := AuthUserUserPermissionWhereToMods(m, !hasForeignKeyInRoot, parentTable)
	if len(subQueryMods) > 0 {
		subQuery := models.AuthUserUserPermissions(append(subQueryMods, qm.Select("1"))...)
		queryMods = appendSubQuery(queryMods, subQuery.Query)
	}
	return queryMods
}

func AuthUserUserPermissionWhereToMods(m *graphql_models.AuthUserUserPermissionWhere, withPrimaryID bool, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	if withPrimaryID {
		queryMods = append(queryMods, IDFilterToMods(m.ID, models.AuthUserUserPermissionColumns.ID)...)
	}
	queryMods = append(queryMods, AuthUserWhereSubqueryToMods(m.User, models.AuthUserUserPermissionColumns.UserID, models.TableNames.AuthUserUserPermissions)...)
	queryMods = append(queryMods, AuthPermissionWhereSubqueryToMods(m.Permission, models.AuthUserUserPermissionColumns.PermissionID, models.TableNames.AuthUserUserPermissions)...)
	if m.Or != nil {
		queryMods = append(queryMods, qm.Or2(qm.Expr(AuthUserUserPermissionWhereToMods(m.Or, true, "")...)))
	}
	if m.And != nil {
		queryMods = append(queryMods, qm.Expr(AuthUserUserPermissionWhereToMods(m.And, true, "")...))
	}

	if len(queryMods) > 0 && parentTable != "" {

		if parentTable == models.TableNames.AuthUser {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.AuthUserUserPermissions, models.AuthUserUserPermissionColumns.UserID, parentTable)))
		}
		if parentTable == models.TableNames.AuthPermission {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.AuthUserUserPermissions, models.AuthUserUserPermissionColumns.PermissionID, parentTable)))
		}
	}

	return queryMods
}
func AuthUserWhereSubqueryToMods(m *graphql_models.AuthUserWhere, foreignColumn string, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	hasForeignKeyInRoot := foreignColumn != ""
	if hasForeignKeyInRoot {
		queryMods = append(queryMods, IDFilterToMods(m.ID, foreignColumn)...)
	}

	subQueryMods := AuthUserWhereToMods(m, !hasForeignKeyInRoot, parentTable)
	if len(subQueryMods) > 0 {
		subQuery := models.AuthUsers(append(subQueryMods, qm.Select("1"))...)
		queryMods = appendSubQuery(queryMods, subQuery.Query)
	}
	return queryMods
}

func AuthUserWhereToMods(m *graphql_models.AuthUserWhere, withPrimaryID bool, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	if withPrimaryID {
		queryMods = append(queryMods, IDFilterToMods(m.ID, models.AuthUserColumns.ID)...)
	}
	queryMods = append(queryMods, StringFilterToMods(m.Password, models.AuthUserColumns.Password)...)
	queryMods = append(queryMods, IntFilterToMods(m.LastLogin, models.AuthUserColumns.LastLogin)...)
	queryMods = append(queryMods, BooleanFilterToMods(m.IsSuperuser, models.AuthUserColumns.IsSuperuser)...)
	queryMods = append(queryMods, StringFilterToMods(m.Username, models.AuthUserColumns.Username)...)
	queryMods = append(queryMods, StringFilterToMods(m.FirstName, models.AuthUserColumns.FirstName)...)
	queryMods = append(queryMods, StringFilterToMods(m.LastName, models.AuthUserColumns.LastName)...)
	queryMods = append(queryMods, StringFilterToMods(m.Email, models.AuthUserColumns.Email)...)
	queryMods = append(queryMods, BooleanFilterToMods(m.IsStaff, models.AuthUserColumns.IsStaff)...)
	queryMods = append(queryMods, BooleanFilterToMods(m.IsActive, models.AuthUserColumns.IsActive)...)
	queryMods = append(queryMods, IntFilterToMods(m.DateJoined, models.AuthUserColumns.DateJoined)...)
	queryMods = append(queryMods, AuthUserGroupWhereSubqueryToMods(m.UserAuthUserGroups, "", models.TableNames.AuthUser)...)
	queryMods = append(queryMods, AuthUserUserPermissionWhereSubqueryToMods(m.UserAuthUserUserPermissions, "", models.TableNames.AuthUser)...)
	if m.Or != nil {
		queryMods = append(queryMods, qm.Or2(qm.Expr(AuthUserWhereToMods(m.Or, true, "")...)))
	}
	if m.And != nil {
		queryMods = append(queryMods, qm.Expr(AuthUserWhereToMods(m.And, true, "")...))
	}

	if len(queryMods) > 0 && parentTable != "" {

	}

	return queryMods
}
func FragranceFilterToMods(m *graphql_models.FragranceFilter) []qm.QueryMod {
	if m == nil {
		return nil
	}
	if m.Search != nil || m.Where != nil {
		var queryMods []qm.QueryMod
		queryMods = append(queryMods, FragranceSearchToMods(m.Search)...)
		queryMods = append(queryMods, FragranceWhereToMods(m.Where, true, "")...)
		if len(queryMods) > 0 {
			return []qm.QueryMod{
				qm.Expr(queryMods...),
			}
		}
	}
	return nil
}
func FragranceSearchToMods(search *string) []qm.QueryMod {

	return nil
}
func FragranceInventoryFilterToMods(m *graphql_models.FragranceInventoryFilter) []qm.QueryMod {
	if m == nil {
		return nil
	}
	if m.Search != nil || m.Where != nil {
		var queryMods []qm.QueryMod
		queryMods = append(queryMods, FragranceInventorySearchToMods(m.Search)...)
		queryMods = append(queryMods, FragranceInventoryWhereToMods(m.Where, true, "")...)
		if len(queryMods) > 0 {
			return []qm.QueryMod{
				qm.Expr(queryMods...),
			}
		}
	}
	return nil
}
func FragranceInventorySearchToMods(search *string) []qm.QueryMod {

	return nil
}
func FragranceInventoryWhereSubqueryToMods(m *graphql_models.FragranceInventoryWhere, foreignColumn string, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	hasForeignKeyInRoot := foreignColumn != ""
	if hasForeignKeyInRoot {
		queryMods = append(queryMods, IDFilterToMods(m.ID, foreignColumn)...)
	}

	subQueryMods := FragranceInventoryWhereToMods(m, !hasForeignKeyInRoot, parentTable)
	if len(subQueryMods) > 0 {
		subQuery := models.FragranceInventories(append(subQueryMods, qm.Select("1"))...)
		queryMods = appendSubQuery(queryMods, subQuery.Query)
	}
	return queryMods
}

func FragranceInventoryWhereToMods(m *graphql_models.FragranceInventoryWhere, withPrimaryID bool, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	if withPrimaryID {
		queryMods = append(queryMods, IDFilterToMods(m.ID, models.FragranceInventoryColumns.ID)...)
	}
	queryMods = append(queryMods, IntFilterToMods(m.PurchaseDate, models.FragranceInventoryColumns.PurchaseDate)...)
	queryMods = append(queryMods, IntFilterToMods(m.ExpiryDate, models.FragranceInventoryColumns.ExpiryDate)...)
	queryMods = append(queryMods, FloatFilterToMods(m.Cost, models.FragranceInventoryColumns.Cost)...)
	queryMods = append(queryMods, FloatFilterToMods(m.Weight, models.FragranceInventoryColumns.Weight)...)
	queryMods = append(queryMods, FragranceWhereSubqueryToMods(m.Fragrance, models.FragranceInventoryColumns.FragranceID, models.TableNames.FragranceInventory)...)
	queryMods = append(queryMods, SupplierWhereSubqueryToMods(m.Supplier, models.FragranceInventoryColumns.SupplierID, models.TableNames.FragranceInventory)...)
	queryMods = append(queryMods, IntFilterToMods(m.CreatedAt, models.FragranceInventoryColumns.CreatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.UpdatedAt, models.FragranceInventoryColumns.UpdatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.DeletedAt, models.FragranceInventoryColumns.DeletedAt)...)
	if m.Or != nil {
		queryMods = append(queryMods, qm.Or2(qm.Expr(FragranceInventoryWhereToMods(m.Or, true, "")...)))
	}
	if m.And != nil {
		queryMods = append(queryMods, qm.Expr(FragranceInventoryWhereToMods(m.And, true, "")...))
	}

	if len(queryMods) > 0 && parentTable != "" {

		if parentTable == models.TableNames.Fragrance {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.FragranceInventory, models.FragranceInventoryColumns.FragranceID, parentTable)))
		}
		if parentTable == models.TableNames.Supplier {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.FragranceInventory, models.FragranceInventoryColumns.SupplierID, parentTable)))
		}
	}

	return queryMods
}
func FragranceWhereSubqueryToMods(m *graphql_models.FragranceWhere, foreignColumn string, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	hasForeignKeyInRoot := foreignColumn != ""
	if hasForeignKeyInRoot {
		queryMods = append(queryMods, IDFilterToMods(m.ID, foreignColumn)...)
	}

	subQueryMods := FragranceWhereToMods(m, !hasForeignKeyInRoot, parentTable)
	if len(subQueryMods) > 0 {
		subQuery := models.Fragrances(append(subQueryMods, qm.Select("1"))...)
		queryMods = appendSubQuery(queryMods, subQuery.Query)
	}
	return queryMods
}

func FragranceWhereToMods(m *graphql_models.FragranceWhere, withPrimaryID bool, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	if withPrimaryID {
		queryMods = append(queryMods, IDFilterToMods(m.ID, models.FragranceColumns.ID)...)
	}
	queryMods = append(queryMods, StringFilterToMods(m.Name, models.FragranceColumns.Name)...)
	queryMods = append(queryMods, StringFilterToMods(m.Note, models.FragranceColumns.Note)...)
	queryMods = append(queryMods, IntFilterToMods(m.CreatedAt, models.FragranceColumns.CreatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.UpdatedAt, models.FragranceColumns.UpdatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.DeletedAt, models.FragranceColumns.DeletedAt)...)
	queryMods = append(queryMods, RecipeBatchFragranceWhereSubqueryToMods(m.RecipeBatchFragrance, "", models.TableNames.Fragrance)...)
	queryMods = append(queryMods, RecipeFragranceWhereSubqueryToMods(m.RecipeFragrance, "", models.TableNames.Fragrance)...)
	queryMods = append(queryMods, FragranceInventoryWhereSubqueryToMods(m.FragranceInventories, "", models.TableNames.Fragrance)...)
	if m.Or != nil {
		queryMods = append(queryMods, qm.Or2(qm.Expr(FragranceWhereToMods(m.Or, true, "")...)))
	}
	if m.And != nil {
		queryMods = append(queryMods, qm.Expr(FragranceWhereToMods(m.And, true, "")...))
	}

	if len(queryMods) > 0 && parentTable != "" {

	}

	return queryMods
}
func LipidFilterToMods(m *graphql_models.LipidFilter) []qm.QueryMod {
	if m == nil {
		return nil
	}
	if m.Search != nil || m.Where != nil {
		var queryMods []qm.QueryMod
		queryMods = append(queryMods, LipidSearchToMods(m.Search)...)
		queryMods = append(queryMods, LipidWhereToMods(m.Where, true, "")...)
		if len(queryMods) > 0 {
			return []qm.QueryMod{
				qm.Expr(queryMods...),
			}
		}
	}
	return nil
}
func LipidSearchToMods(search *string) []qm.QueryMod {

	return nil
}
func LipidInventoryFilterToMods(m *graphql_models.LipidInventoryFilter) []qm.QueryMod {
	if m == nil {
		return nil
	}
	if m.Search != nil || m.Where != nil {
		var queryMods []qm.QueryMod
		queryMods = append(queryMods, LipidInventorySearchToMods(m.Search)...)
		queryMods = append(queryMods, LipidInventoryWhereToMods(m.Where, true, "")...)
		if len(queryMods) > 0 {
			return []qm.QueryMod{
				qm.Expr(queryMods...),
			}
		}
	}
	return nil
}
func LipidInventorySearchToMods(search *string) []qm.QueryMod {

	return nil
}
func LipidInventoryWhereSubqueryToMods(m *graphql_models.LipidInventoryWhere, foreignColumn string, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	hasForeignKeyInRoot := foreignColumn != ""
	if hasForeignKeyInRoot {
		queryMods = append(queryMods, IDFilterToMods(m.ID, foreignColumn)...)
	}

	subQueryMods := LipidInventoryWhereToMods(m, !hasForeignKeyInRoot, parentTable)
	if len(subQueryMods) > 0 {
		subQuery := models.LipidInventories(append(subQueryMods, qm.Select("1"))...)
		queryMods = appendSubQuery(queryMods, subQuery.Query)
	}
	return queryMods
}

func LipidInventoryWhereToMods(m *graphql_models.LipidInventoryWhere, withPrimaryID bool, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	if withPrimaryID {
		queryMods = append(queryMods, IDFilterToMods(m.ID, models.LipidInventoryColumns.ID)...)
	}
	queryMods = append(queryMods, IntFilterToMods(m.PurchaseDate, models.LipidInventoryColumns.PurchaseDate)...)
	queryMods = append(queryMods, IntFilterToMods(m.ExpiryDate, models.LipidInventoryColumns.ExpiryDate)...)
	queryMods = append(queryMods, FloatFilterToMods(m.Cost, models.LipidInventoryColumns.Cost)...)
	queryMods = append(queryMods, FloatFilterToMods(m.Weight, models.LipidInventoryColumns.Weight)...)
	queryMods = append(queryMods, FloatFilterToMods(m.Sap, models.LipidInventoryColumns.Sap)...)
	queryMods = append(queryMods, FloatFilterToMods(m.Naoh, models.LipidInventoryColumns.Naoh)...)
	queryMods = append(queryMods, FloatFilterToMods(m.Koh, models.LipidInventoryColumns.Koh)...)
	queryMods = append(queryMods, FloatFilterToMods(m.GramsPerLiter, models.LipidInventoryColumns.GramsPerLiter)...)
	queryMods = append(queryMods, LipidWhereSubqueryToMods(m.Lipid, models.LipidInventoryColumns.LipidID, models.TableNames.LipidInventory)...)
	queryMods = append(queryMods, SupplierWhereSubqueryToMods(m.Supplier, models.LipidInventoryColumns.SupplierID, models.TableNames.LipidInventory)...)
	queryMods = append(queryMods, IntFilterToMods(m.CreatedAt, models.LipidInventoryColumns.CreatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.UpdatedAt, models.LipidInventoryColumns.UpdatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.DeletedAt, models.LipidInventoryColumns.DeletedAt)...)
	if m.Or != nil {
		queryMods = append(queryMods, qm.Or2(qm.Expr(LipidInventoryWhereToMods(m.Or, true, "")...)))
	}
	if m.And != nil {
		queryMods = append(queryMods, qm.Expr(LipidInventoryWhereToMods(m.And, true, "")...))
	}

	if len(queryMods) > 0 && parentTable != "" {

		if parentTable == models.TableNames.Lipid {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.LipidInventory, models.LipidInventoryColumns.LipidID, parentTable)))
		}
		if parentTable == models.TableNames.Supplier {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.LipidInventory, models.LipidInventoryColumns.SupplierID, parentTable)))
		}
	}

	return queryMods
}
func LipidWhereSubqueryToMods(m *graphql_models.LipidWhere, foreignColumn string, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	hasForeignKeyInRoot := foreignColumn != ""
	if hasForeignKeyInRoot {
		queryMods = append(queryMods, IDFilterToMods(m.ID, foreignColumn)...)
	}

	subQueryMods := LipidWhereToMods(m, !hasForeignKeyInRoot, parentTable)
	if len(subQueryMods) > 0 {
		subQuery := models.Lipids(append(subQueryMods, qm.Select("1"))...)
		queryMods = appendSubQuery(queryMods, subQuery.Query)
	}
	return queryMods
}

func LipidWhereToMods(m *graphql_models.LipidWhere, withPrimaryID bool, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	if withPrimaryID {
		queryMods = append(queryMods, IDFilterToMods(m.ID, models.LipidColumns.ID)...)
	}
	queryMods = append(queryMods, StringFilterToMods(m.Name, models.LipidColumns.Name)...)
	queryMods = append(queryMods, IntFilterToMods(m.Lauric, models.LipidColumns.Lauric)...)
	queryMods = append(queryMods, IntFilterToMods(m.Myristic, models.LipidColumns.Myristic)...)
	queryMods = append(queryMods, IntFilterToMods(m.Palmitic, models.LipidColumns.Palmitic)...)
	queryMods = append(queryMods, IntFilterToMods(m.Stearic, models.LipidColumns.Stearic)...)
	queryMods = append(queryMods, IntFilterToMods(m.Ricinoleic, models.LipidColumns.Ricinoleic)...)
	queryMods = append(queryMods, IntFilterToMods(m.Oleic, models.LipidColumns.Oleic)...)
	queryMods = append(queryMods, IntFilterToMods(m.Linoleic, models.LipidColumns.Linoleic)...)
	queryMods = append(queryMods, IntFilterToMods(m.Linolenic, models.LipidColumns.Linolenic)...)
	queryMods = append(queryMods, IntFilterToMods(m.Hardness, models.LipidColumns.Hardness)...)
	queryMods = append(queryMods, IntFilterToMods(m.Cleansing, models.LipidColumns.Cleansing)...)
	queryMods = append(queryMods, IntFilterToMods(m.Conditioning, models.LipidColumns.Conditioning)...)
	queryMods = append(queryMods, IntFilterToMods(m.Bubbly, models.LipidColumns.Bubbly)...)
	queryMods = append(queryMods, IntFilterToMods(m.Creamy, models.LipidColumns.Creamy)...)
	queryMods = append(queryMods, IntFilterToMods(m.Iodine, models.LipidColumns.Iodine)...)
	queryMods = append(queryMods, IntFilterToMods(m.Ins, models.LipidColumns.Ins)...)
	queryMods = append(queryMods, StringFilterToMods(m.InciName, models.LipidColumns.InciName)...)
	queryMods = append(queryMods, StringFilterToMods(m.Family, models.LipidColumns.Family)...)
	queryMods = append(queryMods, FloatFilterToMods(m.Naoh, models.LipidColumns.Naoh)...)
	queryMods = append(queryMods, IntFilterToMods(m.CreatedAt, models.LipidColumns.CreatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.UpdatedAt, models.LipidColumns.UpdatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.DeletedAt, models.LipidColumns.DeletedAt)...)
	queryMods = append(queryMods, RecipeBatchLipidWhereSubqueryToMods(m.RecipeBatchLipid, "", models.TableNames.Lipid)...)
	queryMods = append(queryMods, RecipeLipidWhereSubqueryToMods(m.RecipeLipid, "", models.TableNames.Lipid)...)
	queryMods = append(queryMods, LipidInventoryWhereSubqueryToMods(m.LipidInventories, "", models.TableNames.Lipid)...)
	if m.Or != nil {
		queryMods = append(queryMods, qm.Or2(qm.Expr(LipidWhereToMods(m.Or, true, "")...)))
	}
	if m.And != nil {
		queryMods = append(queryMods, qm.Expr(LipidWhereToMods(m.And, true, "")...))
	}

	if len(queryMods) > 0 && parentTable != "" {

	}

	return queryMods
}
func LyeFilterToMods(m *graphql_models.LyeFilter) []qm.QueryMod {
	if m == nil {
		return nil
	}
	if m.Search != nil || m.Where != nil {
		var queryMods []qm.QueryMod
		queryMods = append(queryMods, LyeSearchToMods(m.Search)...)
		queryMods = append(queryMods, LyeWhereToMods(m.Where, true, "")...)
		if len(queryMods) > 0 {
			return []qm.QueryMod{
				qm.Expr(queryMods...),
			}
		}
	}
	return nil
}
func LyeSearchToMods(search *string) []qm.QueryMod {

	return nil
}
func LyeInventoryFilterToMods(m *graphql_models.LyeInventoryFilter) []qm.QueryMod {
	if m == nil {
		return nil
	}
	if m.Search != nil || m.Where != nil {
		var queryMods []qm.QueryMod
		queryMods = append(queryMods, LyeInventorySearchToMods(m.Search)...)
		queryMods = append(queryMods, LyeInventoryWhereToMods(m.Where, true, "")...)
		if len(queryMods) > 0 {
			return []qm.QueryMod{
				qm.Expr(queryMods...),
			}
		}
	}
	return nil
}
func LyeInventorySearchToMods(search *string) []qm.QueryMod {

	return nil
}
func LyeInventoryWhereSubqueryToMods(m *graphql_models.LyeInventoryWhere, foreignColumn string, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	hasForeignKeyInRoot := foreignColumn != ""
	if hasForeignKeyInRoot {
		queryMods = append(queryMods, IDFilterToMods(m.ID, foreignColumn)...)
	}

	subQueryMods := LyeInventoryWhereToMods(m, !hasForeignKeyInRoot, parentTable)
	if len(subQueryMods) > 0 {
		subQuery := models.LyeInventories(append(subQueryMods, qm.Select("1"))...)
		queryMods = appendSubQuery(queryMods, subQuery.Query)
	}
	return queryMods
}

func LyeInventoryWhereToMods(m *graphql_models.LyeInventoryWhere, withPrimaryID bool, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	if withPrimaryID {
		queryMods = append(queryMods, IDFilterToMods(m.ID, models.LyeInventoryColumns.ID)...)
	}
	queryMods = append(queryMods, IntFilterToMods(m.PurchaseDate, models.LyeInventoryColumns.PurchaseDate)...)
	queryMods = append(queryMods, IntFilterToMods(m.ExpiryDate, models.LyeInventoryColumns.ExpiryDate)...)
	queryMods = append(queryMods, FloatFilterToMods(m.Cost, models.LyeInventoryColumns.Cost)...)
	queryMods = append(queryMods, FloatFilterToMods(m.Weight, models.LyeInventoryColumns.Weight)...)
	queryMods = append(queryMods, FloatFilterToMods(m.Concentration, models.LyeInventoryColumns.Concentration)...)
	queryMods = append(queryMods, LyeWhereSubqueryToMods(m.Lye, models.LyeInventoryColumns.LyeID, models.TableNames.LyeInventory)...)
	queryMods = append(queryMods, SupplierWhereSubqueryToMods(m.Supplier, models.LyeInventoryColumns.SupplierID, models.TableNames.LyeInventory)...)
	queryMods = append(queryMods, IntFilterToMods(m.CreatedAt, models.LyeInventoryColumns.CreatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.UpdatedAt, models.LyeInventoryColumns.UpdatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.DeletedAt, models.LyeInventoryColumns.DeletedAt)...)
	if m.Or != nil {
		queryMods = append(queryMods, qm.Or2(qm.Expr(LyeInventoryWhereToMods(m.Or, true, "")...)))
	}
	if m.And != nil {
		queryMods = append(queryMods, qm.Expr(LyeInventoryWhereToMods(m.And, true, "")...))
	}

	if len(queryMods) > 0 && parentTable != "" {

		if parentTable == models.TableNames.Lye {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.LyeInventory, models.LyeInventoryColumns.LyeID, parentTable)))
		}
		if parentTable == models.TableNames.Supplier {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.LyeInventory, models.LyeInventoryColumns.SupplierID, parentTable)))
		}
	}

	return queryMods
}
func LyeWhereSubqueryToMods(m *graphql_models.LyeWhere, foreignColumn string, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	hasForeignKeyInRoot := foreignColumn != ""
	if hasForeignKeyInRoot {
		queryMods = append(queryMods, IDFilterToMods(m.ID, foreignColumn)...)
	}

	subQueryMods := LyeWhereToMods(m, !hasForeignKeyInRoot, parentTable)
	if len(subQueryMods) > 0 {
		subQuery := models.Lyes(append(subQueryMods, qm.Select("1"))...)
		queryMods = appendSubQuery(queryMods, subQuery.Query)
	}
	return queryMods
}

func LyeWhereToMods(m *graphql_models.LyeWhere, withPrimaryID bool, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	if withPrimaryID {
		queryMods = append(queryMods, IDFilterToMods(m.ID, models.LyeColumns.ID)...)
	}
	queryMods = append(queryMods, StringFilterToMods(m.Kind, models.LyeColumns.Kind)...)
	queryMods = append(queryMods, StringFilterToMods(m.Name, models.LyeColumns.Name)...)
	queryMods = append(queryMods, StringFilterToMods(m.Note, models.LyeColumns.Note)...)
	queryMods = append(queryMods, IntFilterToMods(m.CreatedAt, models.LyeColumns.CreatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.UpdatedAt, models.LyeColumns.UpdatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.DeletedAt, models.LyeColumns.DeletedAt)...)
	queryMods = append(queryMods, RecipeBatchLyeWhereSubqueryToMods(m.RecipeBatchLye, "", models.TableNames.Lye)...)
	queryMods = append(queryMods, LyeInventoryWhereSubqueryToMods(m.LyeInventories, "", models.TableNames.Lye)...)
	if m.Or != nil {
		queryMods = append(queryMods, qm.Or2(qm.Expr(LyeWhereToMods(m.Or, true, "")...)))
	}
	if m.And != nil {
		queryMods = append(queryMods, qm.Expr(LyeWhereToMods(m.And, true, "")...))
	}

	if len(queryMods) > 0 && parentTable != "" {

	}

	return queryMods
}
func RecipeAdditiveFilterToMods(m *graphql_models.RecipeAdditiveFilter) []qm.QueryMod {
	if m == nil {
		return nil
	}
	if m.Search != nil || m.Where != nil {
		var queryMods []qm.QueryMod
		queryMods = append(queryMods, RecipeAdditiveSearchToMods(m.Search)...)
		queryMods = append(queryMods, RecipeAdditiveWhereToMods(m.Where, true, "")...)
		if len(queryMods) > 0 {
			return []qm.QueryMod{
				qm.Expr(queryMods...),
			}
		}
	}
	return nil
}
func RecipeAdditiveSearchToMods(search *string) []qm.QueryMod {

	return nil
}
func RecipeAdditiveWhereSubqueryToMods(m *graphql_models.RecipeAdditiveWhere, foreignColumn string, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	hasForeignKeyInRoot := foreignColumn != ""
	if hasForeignKeyInRoot {
		queryMods = append(queryMods, IDFilterToMods(m.ID, foreignColumn)...)
	}

	subQueryMods := RecipeAdditiveWhereToMods(m, !hasForeignKeyInRoot, parentTable)
	if len(subQueryMods) > 0 {
		subQuery := models.RecipeAdditives(append(subQueryMods, qm.Select("1"))...)
		queryMods = appendSubQuery(queryMods, subQuery.Query)
	}
	return queryMods
}

func RecipeAdditiveWhereToMods(m *graphql_models.RecipeAdditiveWhere, withPrimaryID bool, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	if withPrimaryID {
		queryMods = append(queryMods, IDFilterToMods(m.ID, models.RecipeAdditiveColumns.ID)...)
	}
	queryMods = append(queryMods, FloatFilterToMods(m.Percentage, models.RecipeAdditiveColumns.Percentage)...)
	queryMods = append(queryMods, AdditiveWhereSubqueryToMods(m.Additive, models.RecipeAdditiveColumns.AdditiveID, models.TableNames.RecipeAdditive)...)
	queryMods = append(queryMods, RecipeWhereSubqueryToMods(m.Recipe, models.RecipeAdditiveColumns.RecipeID, models.TableNames.RecipeAdditive)...)
	queryMods = append(queryMods, IntFilterToMods(m.CreatedAt, models.RecipeAdditiveColumns.CreatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.UpdatedAt, models.RecipeAdditiveColumns.UpdatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.DeletedAt, models.RecipeAdditiveColumns.DeletedAt)...)
	if m.Or != nil {
		queryMods = append(queryMods, qm.Or2(qm.Expr(RecipeAdditiveWhereToMods(m.Or, true, "")...)))
	}
	if m.And != nil {
		queryMods = append(queryMods, qm.Expr(RecipeAdditiveWhereToMods(m.And, true, "")...))
	}

	if len(queryMods) > 0 && parentTable != "" {

		if parentTable == models.TableNames.Additive {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.RecipeAdditive, models.RecipeAdditiveColumns.AdditiveID, parentTable)))
		}
		if parentTable == models.TableNames.Recipe {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.RecipeAdditive, models.RecipeAdditiveColumns.RecipeID, parentTable)))
		}
	}

	return queryMods
}
func RecipeBatchAdditiveFilterToMods(m *graphql_models.RecipeBatchAdditiveFilter) []qm.QueryMod {
	if m == nil {
		return nil
	}
	if m.Search != nil || m.Where != nil {
		var queryMods []qm.QueryMod
		queryMods = append(queryMods, RecipeBatchAdditiveSearchToMods(m.Search)...)
		queryMods = append(queryMods, RecipeBatchAdditiveWhereToMods(m.Where, true, "")...)
		if len(queryMods) > 0 {
			return []qm.QueryMod{
				qm.Expr(queryMods...),
			}
		}
	}
	return nil
}
func RecipeBatchAdditiveSearchToMods(search *string) []qm.QueryMod {

	return nil
}
func RecipeBatchAdditiveWhereSubqueryToMods(m *graphql_models.RecipeBatchAdditiveWhere, foreignColumn string, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	hasForeignKeyInRoot := foreignColumn != ""
	if hasForeignKeyInRoot {
		queryMods = append(queryMods, IDFilterToMods(m.ID, foreignColumn)...)
	}

	subQueryMods := RecipeBatchAdditiveWhereToMods(m, !hasForeignKeyInRoot, parentTable)
	if len(subQueryMods) > 0 {
		subQuery := models.RecipeBatchAdditives(append(subQueryMods, qm.Select("1"))...)
		queryMods = appendSubQuery(queryMods, subQuery.Query)
	}
	return queryMods
}

func RecipeBatchAdditiveWhereToMods(m *graphql_models.RecipeBatchAdditiveWhere, withPrimaryID bool, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	if withPrimaryID {
		queryMods = append(queryMods, IDFilterToMods(m.ID, models.RecipeBatchAdditiveColumns.ID)...)
	}
	queryMods = append(queryMods, FloatFilterToMods(m.Weight, models.RecipeBatchAdditiveColumns.Weight)...)
	queryMods = append(queryMods, FloatFilterToMods(m.Cost, models.RecipeBatchAdditiveColumns.Cost)...)
	queryMods = append(queryMods, AdditiveWhereSubqueryToMods(m.Additive, models.RecipeBatchAdditiveColumns.AdditiveID, models.TableNames.RecipeBatchAdditive)...)
	queryMods = append(queryMods, RecipeBatchWhereSubqueryToMods(m.Batch, models.RecipeBatchAdditiveColumns.BatchID, models.TableNames.RecipeBatchAdditive)...)
	queryMods = append(queryMods, IntFilterToMods(m.CreatedAt, models.RecipeBatchAdditiveColumns.CreatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.UpdatedAt, models.RecipeBatchAdditiveColumns.UpdatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.DeletedAt, models.RecipeBatchAdditiveColumns.DeletedAt)...)
	if m.Or != nil {
		queryMods = append(queryMods, qm.Or2(qm.Expr(RecipeBatchAdditiveWhereToMods(m.Or, true, "")...)))
	}
	if m.And != nil {
		queryMods = append(queryMods, qm.Expr(RecipeBatchAdditiveWhereToMods(m.And, true, "")...))
	}

	if len(queryMods) > 0 && parentTable != "" {

		if parentTable == models.TableNames.Additive {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.RecipeBatchAdditive, models.RecipeBatchAdditiveColumns.AdditiveID, parentTable)))
		}
		if parentTable == models.TableNames.RecipeBatch {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.RecipeBatchAdditive, models.RecipeBatchAdditiveColumns.BatchID, parentTable)))
		}
	}

	return queryMods
}
func RecipeBatchFilterToMods(m *graphql_models.RecipeBatchFilter) []qm.QueryMod {
	if m == nil {
		return nil
	}
	if m.Search != nil || m.Where != nil {
		var queryMods []qm.QueryMod
		queryMods = append(queryMods, RecipeBatchSearchToMods(m.Search)...)
		queryMods = append(queryMods, RecipeBatchWhereToMods(m.Where, true, "")...)
		if len(queryMods) > 0 {
			return []qm.QueryMod{
				qm.Expr(queryMods...),
			}
		}
	}
	return nil
}
func RecipeBatchSearchToMods(search *string) []qm.QueryMod {

	return nil
}
func RecipeBatchFragranceFilterToMods(m *graphql_models.RecipeBatchFragranceFilter) []qm.QueryMod {
	if m == nil {
		return nil
	}
	if m.Search != nil || m.Where != nil {
		var queryMods []qm.QueryMod
		queryMods = append(queryMods, RecipeBatchFragranceSearchToMods(m.Search)...)
		queryMods = append(queryMods, RecipeBatchFragranceWhereToMods(m.Where, true, "")...)
		if len(queryMods) > 0 {
			return []qm.QueryMod{
				qm.Expr(queryMods...),
			}
		}
	}
	return nil
}
func RecipeBatchFragranceSearchToMods(search *string) []qm.QueryMod {

	return nil
}
func RecipeBatchFragranceWhereSubqueryToMods(m *graphql_models.RecipeBatchFragranceWhere, foreignColumn string, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	hasForeignKeyInRoot := foreignColumn != ""
	if hasForeignKeyInRoot {
		queryMods = append(queryMods, IDFilterToMods(m.ID, foreignColumn)...)
	}

	subQueryMods := RecipeBatchFragranceWhereToMods(m, !hasForeignKeyInRoot, parentTable)
	if len(subQueryMods) > 0 {
		subQuery := models.RecipeBatchFragrances(append(subQueryMods, qm.Select("1"))...)
		queryMods = appendSubQuery(queryMods, subQuery.Query)
	}
	return queryMods
}

func RecipeBatchFragranceWhereToMods(m *graphql_models.RecipeBatchFragranceWhere, withPrimaryID bool, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	if withPrimaryID {
		queryMods = append(queryMods, IDFilterToMods(m.ID, models.RecipeBatchFragranceColumns.ID)...)
	}
	queryMods = append(queryMods, FloatFilterToMods(m.Weight, models.RecipeBatchFragranceColumns.Weight)...)
	queryMods = append(queryMods, FloatFilterToMods(m.Cost, models.RecipeBatchFragranceColumns.Cost)...)
	queryMods = append(queryMods, FragranceWhereSubqueryToMods(m.Fragrance, models.RecipeBatchFragranceColumns.FragranceID, models.TableNames.RecipeBatchFragrance)...)
	queryMods = append(queryMods, RecipeBatchWhereSubqueryToMods(m.Batch, models.RecipeBatchFragranceColumns.BatchID, models.TableNames.RecipeBatchFragrance)...)
	queryMods = append(queryMods, IntFilterToMods(m.CreatedAt, models.RecipeBatchFragranceColumns.CreatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.UpdatedAt, models.RecipeBatchFragranceColumns.UpdatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.DeletedAt, models.RecipeBatchFragranceColumns.DeletedAt)...)
	if m.Or != nil {
		queryMods = append(queryMods, qm.Or2(qm.Expr(RecipeBatchFragranceWhereToMods(m.Or, true, "")...)))
	}
	if m.And != nil {
		queryMods = append(queryMods, qm.Expr(RecipeBatchFragranceWhereToMods(m.And, true, "")...))
	}

	if len(queryMods) > 0 && parentTable != "" {

		if parentTable == models.TableNames.Fragrance {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.RecipeBatchFragrance, models.RecipeBatchFragranceColumns.FragranceID, parentTable)))
		}
		if parentTable == models.TableNames.RecipeBatch {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.RecipeBatchFragrance, models.RecipeBatchFragranceColumns.BatchID, parentTable)))
		}
	}

	return queryMods
}
func RecipeBatchLipidFilterToMods(m *graphql_models.RecipeBatchLipidFilter) []qm.QueryMod {
	if m == nil {
		return nil
	}
	if m.Search != nil || m.Where != nil {
		var queryMods []qm.QueryMod
		queryMods = append(queryMods, RecipeBatchLipidSearchToMods(m.Search)...)
		queryMods = append(queryMods, RecipeBatchLipidWhereToMods(m.Where, true, "")...)
		if len(queryMods) > 0 {
			return []qm.QueryMod{
				qm.Expr(queryMods...),
			}
		}
	}
	return nil
}
func RecipeBatchLipidSearchToMods(search *string) []qm.QueryMod {

	return nil
}
func RecipeBatchLipidWhereSubqueryToMods(m *graphql_models.RecipeBatchLipidWhere, foreignColumn string, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	hasForeignKeyInRoot := foreignColumn != ""
	if hasForeignKeyInRoot {
		queryMods = append(queryMods, IDFilterToMods(m.ID, foreignColumn)...)
	}

	subQueryMods := RecipeBatchLipidWhereToMods(m, !hasForeignKeyInRoot, parentTable)
	if len(subQueryMods) > 0 {
		subQuery := models.RecipeBatchLipids(append(subQueryMods, qm.Select("1"))...)
		queryMods = appendSubQuery(queryMods, subQuery.Query)
	}
	return queryMods
}

func RecipeBatchLipidWhereToMods(m *graphql_models.RecipeBatchLipidWhere, withPrimaryID bool, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	if withPrimaryID {
		queryMods = append(queryMods, IDFilterToMods(m.ID, models.RecipeBatchLipidColumns.ID)...)
	}
	queryMods = append(queryMods, FloatFilterToMods(m.Weight, models.RecipeBatchLipidColumns.Weight)...)
	queryMods = append(queryMods, FloatFilterToMods(m.Cost, models.RecipeBatchLipidColumns.Cost)...)
	queryMods = append(queryMods, LipidWhereSubqueryToMods(m.Lipid, models.RecipeBatchLipidColumns.LipidID, models.TableNames.RecipeBatchLipid)...)
	queryMods = append(queryMods, RecipeBatchWhereSubqueryToMods(m.Batch, models.RecipeBatchLipidColumns.BatchID, models.TableNames.RecipeBatchLipid)...)
	queryMods = append(queryMods, IntFilterToMods(m.CreatedAt, models.RecipeBatchLipidColumns.CreatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.UpdatedAt, models.RecipeBatchLipidColumns.UpdatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.DeletedAt, models.RecipeBatchLipidColumns.DeletedAt)...)
	if m.Or != nil {
		queryMods = append(queryMods, qm.Or2(qm.Expr(RecipeBatchLipidWhereToMods(m.Or, true, "")...)))
	}
	if m.And != nil {
		queryMods = append(queryMods, qm.Expr(RecipeBatchLipidWhereToMods(m.And, true, "")...))
	}

	if len(queryMods) > 0 && parentTable != "" {

		if parentTable == models.TableNames.Lipid {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.RecipeBatchLipid, models.RecipeBatchLipidColumns.LipidID, parentTable)))
		}
		if parentTable == models.TableNames.RecipeBatch {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.RecipeBatchLipid, models.RecipeBatchLipidColumns.BatchID, parentTable)))
		}
	}

	return queryMods
}
func RecipeBatchLyeFilterToMods(m *graphql_models.RecipeBatchLyeFilter) []qm.QueryMod {
	if m == nil {
		return nil
	}
	if m.Search != nil || m.Where != nil {
		var queryMods []qm.QueryMod
		queryMods = append(queryMods, RecipeBatchLyeSearchToMods(m.Search)...)
		queryMods = append(queryMods, RecipeBatchLyeWhereToMods(m.Where, true, "")...)
		if len(queryMods) > 0 {
			return []qm.QueryMod{
				qm.Expr(queryMods...),
			}
		}
	}
	return nil
}
func RecipeBatchLyeSearchToMods(search *string) []qm.QueryMod {

	return nil
}
func RecipeBatchLyeWhereSubqueryToMods(m *graphql_models.RecipeBatchLyeWhere, foreignColumn string, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	hasForeignKeyInRoot := foreignColumn != ""
	if hasForeignKeyInRoot {
		queryMods = append(queryMods, IDFilterToMods(m.ID, foreignColumn)...)
	}

	subQueryMods := RecipeBatchLyeWhereToMods(m, !hasForeignKeyInRoot, parentTable)
	if len(subQueryMods) > 0 {
		subQuery := models.RecipeBatchLyes(append(subQueryMods, qm.Select("1"))...)
		queryMods = appendSubQuery(queryMods, subQuery.Query)
	}
	return queryMods
}

func RecipeBatchLyeWhereToMods(m *graphql_models.RecipeBatchLyeWhere, withPrimaryID bool, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	if withPrimaryID {
		queryMods = append(queryMods, IDFilterToMods(m.ID, models.RecipeBatchLyeColumns.ID)...)
	}
	queryMods = append(queryMods, FloatFilterToMods(m.Weight, models.RecipeBatchLyeColumns.Weight)...)
	queryMods = append(queryMods, FloatFilterToMods(m.Discount, models.RecipeBatchLyeColumns.Discount)...)
	queryMods = append(queryMods, FloatFilterToMods(m.Cost, models.RecipeBatchLyeColumns.Cost)...)
	queryMods = append(queryMods, LyeWhereSubqueryToMods(m.Lye, models.RecipeBatchLyeColumns.LyeID, models.TableNames.RecipeBatchLye)...)
	queryMods = append(queryMods, RecipeBatchWhereSubqueryToMods(m.Batch, models.RecipeBatchLyeColumns.BatchID, models.TableNames.RecipeBatchLye)...)
	queryMods = append(queryMods, IntFilterToMods(m.CreatedAt, models.RecipeBatchLyeColumns.CreatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.UpdatedAt, models.RecipeBatchLyeColumns.UpdatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.DeletedAt, models.RecipeBatchLyeColumns.DeletedAt)...)
	if m.Or != nil {
		queryMods = append(queryMods, qm.Or2(qm.Expr(RecipeBatchLyeWhereToMods(m.Or, true, "")...)))
	}
	if m.And != nil {
		queryMods = append(queryMods, qm.Expr(RecipeBatchLyeWhereToMods(m.And, true, "")...))
	}

	if len(queryMods) > 0 && parentTable != "" {

		if parentTable == models.TableNames.Lye {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.RecipeBatchLye, models.RecipeBatchLyeColumns.LyeID, parentTable)))
		}
		if parentTable == models.TableNames.RecipeBatch {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.RecipeBatchLye, models.RecipeBatchLyeColumns.BatchID, parentTable)))
		}
	}

	return queryMods
}
func RecipeBatchNoteFilterToMods(m *graphql_models.RecipeBatchNoteFilter) []qm.QueryMod {
	if m == nil {
		return nil
	}
	if m.Search != nil || m.Where != nil {
		var queryMods []qm.QueryMod
		queryMods = append(queryMods, RecipeBatchNoteSearchToMods(m.Search)...)
		queryMods = append(queryMods, RecipeBatchNoteWhereToMods(m.Where, true, "")...)
		if len(queryMods) > 0 {
			return []qm.QueryMod{
				qm.Expr(queryMods...),
			}
		}
	}
	return nil
}
func RecipeBatchNoteSearchToMods(search *string) []qm.QueryMod {

	return nil
}
func RecipeBatchNoteWhereSubqueryToMods(m *graphql_models.RecipeBatchNoteWhere, foreignColumn string, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	hasForeignKeyInRoot := foreignColumn != ""
	if hasForeignKeyInRoot {
		queryMods = append(queryMods, IDFilterToMods(m.ID, foreignColumn)...)
	}

	subQueryMods := RecipeBatchNoteWhereToMods(m, !hasForeignKeyInRoot, parentTable)
	if len(subQueryMods) > 0 {
		subQuery := models.RecipeBatchNotes(append(subQueryMods, qm.Select("1"))...)
		queryMods = appendSubQuery(queryMods, subQuery.Query)
	}
	return queryMods
}

func RecipeBatchNoteWhereToMods(m *graphql_models.RecipeBatchNoteWhere, withPrimaryID bool, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	if withPrimaryID {
		queryMods = append(queryMods, IDFilterToMods(m.ID, models.RecipeBatchNoteColumns.ID)...)
	}
	queryMods = append(queryMods, StringFilterToMods(m.Note, models.RecipeBatchNoteColumns.Note)...)
	queryMods = append(queryMods, StringFilterToMods(m.Link, models.RecipeBatchNoteColumns.Link)...)
	queryMods = append(queryMods, RecipeBatchWhereSubqueryToMods(m.Batch, models.RecipeBatchNoteColumns.BatchID, models.TableNames.RecipeBatchNote)...)
	queryMods = append(queryMods, IntFilterToMods(m.CreatedAt, models.RecipeBatchNoteColumns.CreatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.UpdatedAt, models.RecipeBatchNoteColumns.UpdatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.DeletedAt, models.RecipeBatchNoteColumns.DeletedAt)...)
	if m.Or != nil {
		queryMods = append(queryMods, qm.Or2(qm.Expr(RecipeBatchNoteWhereToMods(m.Or, true, "")...)))
	}
	if m.And != nil {
		queryMods = append(queryMods, qm.Expr(RecipeBatchNoteWhereToMods(m.And, true, "")...))
	}

	if len(queryMods) > 0 && parentTable != "" {

		if parentTable == models.TableNames.RecipeBatch {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.RecipeBatchNote, models.RecipeBatchNoteColumns.BatchID, parentTable)))
		}
	}

	return queryMods
}
func RecipeBatchWhereSubqueryToMods(m *graphql_models.RecipeBatchWhere, foreignColumn string, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	hasForeignKeyInRoot := foreignColumn != ""
	if hasForeignKeyInRoot {
		queryMods = append(queryMods, IDFilterToMods(m.ID, foreignColumn)...)
	}

	subQueryMods := RecipeBatchWhereToMods(m, !hasForeignKeyInRoot, parentTable)
	if len(subQueryMods) > 0 {
		subQuery := models.RecipeBatches(append(subQueryMods, qm.Select("1"))...)
		queryMods = appendSubQuery(queryMods, subQuery.Query)
	}
	return queryMods
}

func RecipeBatchWhereToMods(m *graphql_models.RecipeBatchWhere, withPrimaryID bool, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	if withPrimaryID {
		queryMods = append(queryMods, IDFilterToMods(m.ID, models.RecipeBatchColumns.ID)...)
	}
	queryMods = append(queryMods, StringFilterToMods(m.Tag, models.RecipeBatchColumns.Tag)...)
	queryMods = append(queryMods, IntFilterToMods(m.ProductionDate, models.RecipeBatchColumns.ProductionDate)...)
	queryMods = append(queryMods, IntFilterToMods(m.SellableDate, models.RecipeBatchColumns.SellableDate)...)
	queryMods = append(queryMods, StringFilterToMods(m.Note, models.RecipeBatchColumns.Note)...)
	queryMods = append(queryMods, FloatFilterToMods(m.LipidWeight, models.RecipeBatchColumns.LipidWeight)...)
	queryMods = append(queryMods, FloatFilterToMods(m.ProductionWeight, models.RecipeBatchColumns.ProductionWeight)...)
	queryMods = append(queryMods, FloatFilterToMods(m.CuredWeight, models.RecipeBatchColumns.CuredWeight)...)
	queryMods = append(queryMods, RecipeWhereSubqueryToMods(m.Recipe, models.RecipeBatchColumns.RecipeID, models.TableNames.RecipeBatch)...)
	queryMods = append(queryMods, IntFilterToMods(m.CreatedAt, models.RecipeBatchColumns.CreatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.UpdatedAt, models.RecipeBatchColumns.UpdatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.DeletedAt, models.RecipeBatchColumns.DeletedAt)...)
	queryMods = append(queryMods, RecipeBatchAdditiveWhereSubqueryToMods(m.BatchRecipeBatchAdditives, "", models.TableNames.RecipeBatch)...)
	queryMods = append(queryMods, RecipeBatchFragranceWhereSubqueryToMods(m.BatchRecipeBatchFragrances, "", models.TableNames.RecipeBatch)...)
	queryMods = append(queryMods, RecipeBatchLipidWhereSubqueryToMods(m.BatchRecipeBatchLipids, "", models.TableNames.RecipeBatch)...)
	queryMods = append(queryMods, RecipeBatchLyeWhereSubqueryToMods(m.BatchRecipeBatchLyes, "", models.TableNames.RecipeBatch)...)
	queryMods = append(queryMods, RecipeBatchNoteWhereSubqueryToMods(m.BatchRecipeBatchNotes, "", models.TableNames.RecipeBatch)...)
	if m.Or != nil {
		queryMods = append(queryMods, qm.Or2(qm.Expr(RecipeBatchWhereToMods(m.Or, true, "")...)))
	}
	if m.And != nil {
		queryMods = append(queryMods, qm.Expr(RecipeBatchWhereToMods(m.And, true, "")...))
	}

	if len(queryMods) > 0 && parentTable != "" {

		if parentTable == models.TableNames.Recipe {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.RecipeBatch, models.RecipeBatchColumns.RecipeID, parentTable)))
		}
	}

	return queryMods
}
func RecipeFilterToMods(m *graphql_models.RecipeFilter) []qm.QueryMod {
	if m == nil {
		return nil
	}
	if m.Search != nil || m.Where != nil {
		var queryMods []qm.QueryMod
		queryMods = append(queryMods, RecipeSearchToMods(m.Search)...)
		queryMods = append(queryMods, RecipeWhereToMods(m.Where, true, "")...)
		if len(queryMods) > 0 {
			return []qm.QueryMod{
				qm.Expr(queryMods...),
			}
		}
	}
	return nil
}
func RecipeSearchToMods(search *string) []qm.QueryMod {

	return nil
}
func RecipeFragranceFilterToMods(m *graphql_models.RecipeFragranceFilter) []qm.QueryMod {
	if m == nil {
		return nil
	}
	if m.Search != nil || m.Where != nil {
		var queryMods []qm.QueryMod
		queryMods = append(queryMods, RecipeFragranceSearchToMods(m.Search)...)
		queryMods = append(queryMods, RecipeFragranceWhereToMods(m.Where, true, "")...)
		if len(queryMods) > 0 {
			return []qm.QueryMod{
				qm.Expr(queryMods...),
			}
		}
	}
	return nil
}
func RecipeFragranceSearchToMods(search *string) []qm.QueryMod {

	return nil
}
func RecipeFragranceWhereSubqueryToMods(m *graphql_models.RecipeFragranceWhere, foreignColumn string, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	hasForeignKeyInRoot := foreignColumn != ""
	if hasForeignKeyInRoot {
		queryMods = append(queryMods, IDFilterToMods(m.ID, foreignColumn)...)
	}

	subQueryMods := RecipeFragranceWhereToMods(m, !hasForeignKeyInRoot, parentTable)
	if len(subQueryMods) > 0 {
		subQuery := models.RecipeFragrances(append(subQueryMods, qm.Select("1"))...)
		queryMods = appendSubQuery(queryMods, subQuery.Query)
	}
	return queryMods
}

func RecipeFragranceWhereToMods(m *graphql_models.RecipeFragranceWhere, withPrimaryID bool, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	if withPrimaryID {
		queryMods = append(queryMods, IDFilterToMods(m.ID, models.RecipeFragranceColumns.ID)...)
	}
	queryMods = append(queryMods, FloatFilterToMods(m.Percentage, models.RecipeFragranceColumns.Percentage)...)
	queryMods = append(queryMods, FragranceWhereSubqueryToMods(m.Fragrance, models.RecipeFragranceColumns.FragranceID, models.TableNames.RecipeFragrance)...)
	queryMods = append(queryMods, RecipeWhereSubqueryToMods(m.Recipe, models.RecipeFragranceColumns.RecipeID, models.TableNames.RecipeFragrance)...)
	queryMods = append(queryMods, IntFilterToMods(m.CreatedAt, models.RecipeFragranceColumns.CreatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.UpdatedAt, models.RecipeFragranceColumns.UpdatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.DeletedAt, models.RecipeFragranceColumns.DeletedAt)...)
	if m.Or != nil {
		queryMods = append(queryMods, qm.Or2(qm.Expr(RecipeFragranceWhereToMods(m.Or, true, "")...)))
	}
	if m.And != nil {
		queryMods = append(queryMods, qm.Expr(RecipeFragranceWhereToMods(m.And, true, "")...))
	}

	if len(queryMods) > 0 && parentTable != "" {

		if parentTable == models.TableNames.Fragrance {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.RecipeFragrance, models.RecipeFragranceColumns.FragranceID, parentTable)))
		}
		if parentTable == models.TableNames.Recipe {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.RecipeFragrance, models.RecipeFragranceColumns.RecipeID, parentTable)))
		}
	}

	return queryMods
}
func RecipeLipidFilterToMods(m *graphql_models.RecipeLipidFilter) []qm.QueryMod {
	if m == nil {
		return nil
	}
	if m.Search != nil || m.Where != nil {
		var queryMods []qm.QueryMod
		queryMods = append(queryMods, RecipeLipidSearchToMods(m.Search)...)
		queryMods = append(queryMods, RecipeLipidWhereToMods(m.Where, true, "")...)
		if len(queryMods) > 0 {
			return []qm.QueryMod{
				qm.Expr(queryMods...),
			}
		}
	}
	return nil
}
func RecipeLipidSearchToMods(search *string) []qm.QueryMod {

	return nil
}
func RecipeLipidWhereSubqueryToMods(m *graphql_models.RecipeLipidWhere, foreignColumn string, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	hasForeignKeyInRoot := foreignColumn != ""
	if hasForeignKeyInRoot {
		queryMods = append(queryMods, IDFilterToMods(m.ID, foreignColumn)...)
	}

	subQueryMods := RecipeLipidWhereToMods(m, !hasForeignKeyInRoot, parentTable)
	if len(subQueryMods) > 0 {
		subQuery := models.RecipeLipids(append(subQueryMods, qm.Select("1"))...)
		queryMods = appendSubQuery(queryMods, subQuery.Query)
	}
	return queryMods
}

func RecipeLipidWhereToMods(m *graphql_models.RecipeLipidWhere, withPrimaryID bool, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	if withPrimaryID {
		queryMods = append(queryMods, IDFilterToMods(m.ID, models.RecipeLipidColumns.ID)...)
	}
	queryMods = append(queryMods, FloatFilterToMods(m.Percentage, models.RecipeLipidColumns.Percentage)...)
	queryMods = append(queryMods, LipidWhereSubqueryToMods(m.Lipid, models.RecipeLipidColumns.LipidID, models.TableNames.RecipeLipid)...)
	queryMods = append(queryMods, RecipeWhereSubqueryToMods(m.Recipe, models.RecipeLipidColumns.RecipeID, models.TableNames.RecipeLipid)...)
	queryMods = append(queryMods, IntFilterToMods(m.CreatedAt, models.RecipeLipidColumns.CreatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.UpdatedAt, models.RecipeLipidColumns.UpdatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.DeletedAt, models.RecipeLipidColumns.DeletedAt)...)
	if m.Or != nil {
		queryMods = append(queryMods, qm.Or2(qm.Expr(RecipeLipidWhereToMods(m.Or, true, "")...)))
	}
	if m.And != nil {
		queryMods = append(queryMods, qm.Expr(RecipeLipidWhereToMods(m.And, true, "")...))
	}

	if len(queryMods) > 0 && parentTable != "" {

		if parentTable == models.TableNames.Lipid {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.RecipeLipid, models.RecipeLipidColumns.LipidID, parentTable)))
		}
		if parentTable == models.TableNames.Recipe {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.RecipeLipid, models.RecipeLipidColumns.RecipeID, parentTable)))
		}
	}

	return queryMods
}
func RecipeStepFilterToMods(m *graphql_models.RecipeStepFilter) []qm.QueryMod {
	if m == nil {
		return nil
	}
	if m.Search != nil || m.Where != nil {
		var queryMods []qm.QueryMod
		queryMods = append(queryMods, RecipeStepSearchToMods(m.Search)...)
		queryMods = append(queryMods, RecipeStepWhereToMods(m.Where, true, "")...)
		if len(queryMods) > 0 {
			return []qm.QueryMod{
				qm.Expr(queryMods...),
			}
		}
	}
	return nil
}
func RecipeStepSearchToMods(search *string) []qm.QueryMod {

	return nil
}
func RecipeStepWhereSubqueryToMods(m *graphql_models.RecipeStepWhere, foreignColumn string, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	hasForeignKeyInRoot := foreignColumn != ""
	if hasForeignKeyInRoot {
		queryMods = append(queryMods, IDFilterToMods(m.ID, foreignColumn)...)
	}

	subQueryMods := RecipeStepWhereToMods(m, !hasForeignKeyInRoot, parentTable)
	if len(subQueryMods) > 0 {
		subQuery := models.RecipeSteps(append(subQueryMods, qm.Select("1"))...)
		queryMods = appendSubQuery(queryMods, subQuery.Query)
	}
	return queryMods
}

func RecipeStepWhereToMods(m *graphql_models.RecipeStepWhere, withPrimaryID bool, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	if withPrimaryID {
		queryMods = append(queryMods, IDFilterToMods(m.ID, models.RecipeStepColumns.ID)...)
	}
	queryMods = append(queryMods, IntFilterToMods(m.Num, models.RecipeStepColumns.Num)...)
	queryMods = append(queryMods, StringFilterToMods(m.Note, models.RecipeStepColumns.Note)...)
	queryMods = append(queryMods, RecipeWhereSubqueryToMods(m.Recipe, models.RecipeStepColumns.RecipeID, models.TableNames.RecipeStep)...)
	queryMods = append(queryMods, IntFilterToMods(m.CreatedAt, models.RecipeStepColumns.CreatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.UpdatedAt, models.RecipeStepColumns.UpdatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.DeletedAt, models.RecipeStepColumns.DeletedAt)...)
	if m.Or != nil {
		queryMods = append(queryMods, qm.Or2(qm.Expr(RecipeStepWhereToMods(m.Or, true, "")...)))
	}
	if m.And != nil {
		queryMods = append(queryMods, qm.Expr(RecipeStepWhereToMods(m.And, true, "")...))
	}

	if len(queryMods) > 0 && parentTable != "" {

		if parentTable == models.TableNames.Recipe {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.RecipeStep, models.RecipeStepColumns.RecipeID, parentTable)))
		}
	}

	return queryMods
}
func RecipeWhereSubqueryToMods(m *graphql_models.RecipeWhere, foreignColumn string, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	hasForeignKeyInRoot := foreignColumn != ""
	if hasForeignKeyInRoot {
		queryMods = append(queryMods, IDFilterToMods(m.ID, foreignColumn)...)
	}

	subQueryMods := RecipeWhereToMods(m, !hasForeignKeyInRoot, parentTable)
	if len(subQueryMods) > 0 {
		subQuery := models.Recipes(append(subQueryMods, qm.Select("1"))...)
		queryMods = appendSubQuery(queryMods, subQuery.Query)
	}
	return queryMods
}

func RecipeWhereToMods(m *graphql_models.RecipeWhere, withPrimaryID bool, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	if withPrimaryID {
		queryMods = append(queryMods, IDFilterToMods(m.ID, models.RecipeColumns.ID)...)
	}
	queryMods = append(queryMods, StringFilterToMods(m.Name, models.RecipeColumns.Name)...)
	queryMods = append(queryMods, StringFilterToMods(m.Note, models.RecipeColumns.Note)...)
	queryMods = append(queryMods, IntFilterToMods(m.CreatedAt, models.RecipeColumns.CreatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.UpdatedAt, models.RecipeColumns.UpdatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.DeletedAt, models.RecipeColumns.DeletedAt)...)
	queryMods = append(queryMods, RecipeAdditiveWhereSubqueryToMods(m.RecipeAdditives, "", models.TableNames.Recipe)...)
	queryMods = append(queryMods, RecipeBatchWhereSubqueryToMods(m.RecipeBatches, "", models.TableNames.Recipe)...)
	queryMods = append(queryMods, RecipeFragranceWhereSubqueryToMods(m.RecipeFragrances, "", models.TableNames.Recipe)...)
	queryMods = append(queryMods, RecipeLipidWhereSubqueryToMods(m.RecipeLipids, "", models.TableNames.Recipe)...)
	queryMods = append(queryMods, RecipeStepWhereSubqueryToMods(m.RecipeSteps, "", models.TableNames.Recipe)...)
	if m.Or != nil {
		queryMods = append(queryMods, qm.Or2(qm.Expr(RecipeWhereToMods(m.Or, true, "")...)))
	}
	if m.And != nil {
		queryMods = append(queryMods, qm.Expr(RecipeWhereToMods(m.And, true, "")...))
	}

	if len(queryMods) > 0 && parentTable != "" {

	}

	return queryMods
}
func SupplierFilterToMods(m *graphql_models.SupplierFilter) []qm.QueryMod {
	if m == nil {
		return nil
	}
	if m.Search != nil || m.Where != nil {
		var queryMods []qm.QueryMod
		queryMods = append(queryMods, SupplierSearchToMods(m.Search)...)
		queryMods = append(queryMods, SupplierWhereToMods(m.Where, true, "")...)
		if len(queryMods) > 0 {
			return []qm.QueryMod{
				qm.Expr(queryMods...),
			}
		}
	}
	return nil
}
func SupplierSearchToMods(search *string) []qm.QueryMod {

	return nil
}
func SupplierWhereSubqueryToMods(m *graphql_models.SupplierWhere, foreignColumn string, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	hasForeignKeyInRoot := foreignColumn != ""
	if hasForeignKeyInRoot {
		queryMods = append(queryMods, IDFilterToMods(m.ID, foreignColumn)...)
	}

	subQueryMods := SupplierWhereToMods(m, !hasForeignKeyInRoot, parentTable)
	if len(subQueryMods) > 0 {
		subQuery := models.Suppliers(append(subQueryMods, qm.Select("1"))...)
		queryMods = appendSubQuery(queryMods, subQuery.Query)
	}
	return queryMods
}

func SupplierWhereToMods(m *graphql_models.SupplierWhere, withPrimaryID bool, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	if withPrimaryID {
		queryMods = append(queryMods, IDFilterToMods(m.ID, models.SupplierColumns.ID)...)
	}
	queryMods = append(queryMods, StringFilterToMods(m.Name, models.SupplierColumns.Name)...)
	queryMods = append(queryMods, StringFilterToMods(m.Website, models.SupplierColumns.Website)...)
	queryMods = append(queryMods, StringFilterToMods(m.Note, models.SupplierColumns.Note)...)
	queryMods = append(queryMods, IntFilterToMods(m.CreatedAt, models.SupplierColumns.CreatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.UpdatedAt, models.SupplierColumns.UpdatedAt)...)
	queryMods = append(queryMods, IntFilterToMods(m.DeletedAt, models.SupplierColumns.DeletedAt)...)
	queryMods = append(queryMods, AdditiveInventoryWhereSubqueryToMods(m.AdditiveInventories, "", models.TableNames.Supplier)...)
	queryMods = append(queryMods, FragranceInventoryWhereSubqueryToMods(m.FragranceInventories, "", models.TableNames.Supplier)...)
	queryMods = append(queryMods, LipidInventoryWhereSubqueryToMods(m.LipidInventories, "", models.TableNames.Supplier)...)
	queryMods = append(queryMods, LyeInventoryWhereSubqueryToMods(m.LyeInventories, "", models.TableNames.Supplier)...)
	if m.Or != nil {
		queryMods = append(queryMods, qm.Or2(qm.Expr(SupplierWhereToMods(m.Or, true, "")...)))
	}
	if m.And != nil {
		queryMods = append(queryMods, qm.Expr(SupplierWhereToMods(m.And, true, "")...))
	}

	if len(queryMods) > 0 && parentTable != "" {

	}

	return queryMods
}
