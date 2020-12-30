package helpers

import (
	"fmt"
	"strings"

	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/web-ridge/gqlgen-sqlboiler-examples/issue-12-string-ids/graphql_models"
	"github.com/web-ridge/gqlgen-sqlboiler-examples/issue-12-string-ids/models"
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

func CommentFilterToMods(m *graphql_models.CommentFilter) []qm.QueryMod {
	if m == nil {
		return nil
	}
	if m.Search != nil || m.Where != nil {
		var queryMods []qm.QueryMod
		queryMods = append(queryMods, CommentSearchToMods(m.Search)...)
		queryMods = append(queryMods, CommentWhereToMods(m.Where, true, "")...)
		if len(queryMods) > 0 {
			return []qm.QueryMod{
				qm.Expr(queryMods...),
			}
		}
	}
	return nil
}
func CommentSearchToMods(search *string) []qm.QueryMod {

	return nil
}
func CommentLikeFilterToMods(m *graphql_models.CommentLikeFilter) []qm.QueryMod {
	if m == nil {
		return nil
	}
	if m.Search != nil || m.Where != nil {
		var queryMods []qm.QueryMod
		queryMods = append(queryMods, CommentLikeSearchToMods(m.Search)...)
		queryMods = append(queryMods, CommentLikeWhereToMods(m.Where, true, "")...)
		if len(queryMods) > 0 {
			return []qm.QueryMod{
				qm.Expr(queryMods...),
			}
		}
	}
	return nil
}
func CommentLikeSearchToMods(search *string) []qm.QueryMod {

	return nil
}
func CommentLikeWhereSubqueryToMods(m *graphql_models.CommentLikeWhere, foreignColumn string, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	hasForeignKeyInRoot := foreignColumn != ""
	if hasForeignKeyInRoot {
		queryMods = append(queryMods, IDFilterToMods(m.ID, foreignColumn)...)
	}

	subQueryMods := CommentLikeWhereToMods(m, !hasForeignKeyInRoot, parentTable)
	if len(subQueryMods) > 0 {
		subQuery := models.CommentLikes(append(subQueryMods, qm.Select("1"))...)
		queryMods = appendSubQuery(queryMods, subQuery.Query)
	}
	return queryMods
}

func CommentLikeWhereToMods(m *graphql_models.CommentLikeWhere, withPrimaryID bool, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	if withPrimaryID {
		queryMods = append(queryMods, IDFilterToMods(m.ID, models.CommentLikeColumns.ID)...)
	}
	queryMods = append(queryMods, CommentWhereSubqueryToMods(m.Comment, models.CommentLikeColumns.CommentID, models.TableNames.CommentLike)...)
	queryMods = append(queryMods, UserWhereSubqueryToMods(m.User, models.CommentLikeColumns.UserID, models.TableNames.CommentLike)...)
	queryMods = append(queryMods, StringFilterToMods(m.LikeType, models.CommentLikeColumns.LikeType)...)
	queryMods = append(queryMods, IntFilterToMods(m.CreatedAt, models.CommentLikeColumns.CreatedAt)...)
	if m.Or != nil {
		queryMods = append(queryMods, qm.Or2(qm.Expr(CommentLikeWhereToMods(m.Or, true, "")...)))
	}
	if m.And != nil {
		queryMods = append(queryMods, qm.Expr(CommentLikeWhereToMods(m.And, true, "")...))
	}

	if len(queryMods) > 0 && parentTable != "" {

		if parentTable == models.TableNames.Comment {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.CommentLike, models.CommentLikeColumns.CommentID, parentTable)))
		}
		if parentTable == models.TableNames.User {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.CommentLike, models.CommentLikeColumns.UserID, parentTable)))
		}
	}

	return queryMods
}
func CommentWhereSubqueryToMods(m *graphql_models.CommentWhere, foreignColumn string, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	hasForeignKeyInRoot := foreignColumn != ""
	if hasForeignKeyInRoot {
		queryMods = append(queryMods, IDFilterToMods(m.ID, foreignColumn)...)
	}

	subQueryMods := CommentWhereToMods(m, !hasForeignKeyInRoot, parentTable)
	if len(subQueryMods) > 0 {
		subQuery := models.Comments(append(subQueryMods, qm.Select("1"))...)
		queryMods = appendSubQuery(queryMods, subQuery.Query)
	}
	return queryMods
}

func CommentWhereToMods(m *graphql_models.CommentWhere, withPrimaryID bool, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	if withPrimaryID {
		queryMods = append(queryMods, IDFilterToMods(m.ID, models.CommentColumns.ID)...)
	}
	queryMods = append(queryMods, StringFilterToMods(m.Content, models.CommentColumns.Content)...)
	queryMods = append(queryMods, PostWhereSubqueryToMods(m.Post, models.CommentColumns.PostID, models.TableNames.Comment)...)
	queryMods = append(queryMods, UserWhereSubqueryToMods(m.User, models.CommentColumns.UserID, models.TableNames.Comment)...)
	queryMods = append(queryMods, CommentLikeWhereSubqueryToMods(m.CommentLikes, "", models.TableNames.Comment)...)
	if m.Or != nil {
		queryMods = append(queryMods, qm.Or2(qm.Expr(CommentWhereToMods(m.Or, true, "")...)))
	}
	if m.And != nil {
		queryMods = append(queryMods, qm.Expr(CommentWhereToMods(m.And, true, "")...))
	}

	if len(queryMods) > 0 && parentTable != "" {

		if parentTable == models.TableNames.Post {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.Comment, models.CommentColumns.PostID, parentTable)))
		}
		if parentTable == models.TableNames.User {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.Comment, models.CommentColumns.UserID, parentTable)))
		}
	}

	return queryMods
}
func FriendshipFilterToMods(m *graphql_models.FriendshipFilter) []qm.QueryMod {
	if m == nil {
		return nil
	}
	if m.Search != nil || m.Where != nil {
		var queryMods []qm.QueryMod
		queryMods = append(queryMods, FriendshipSearchToMods(m.Search)...)
		queryMods = append(queryMods, FriendshipWhereToMods(m.Where, true, "")...)
		if len(queryMods) > 0 {
			return []qm.QueryMod{
				qm.Expr(queryMods...),
			}
		}
	}
	return nil
}
func FriendshipSearchToMods(search *string) []qm.QueryMod {

	return nil
}
func FriendshipWhereSubqueryToMods(m *graphql_models.FriendshipWhere, foreignColumn string, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	hasForeignKeyInRoot := foreignColumn != ""
	if hasForeignKeyInRoot {
		queryMods = append(queryMods, IDFilterToMods(m.ID, foreignColumn)...)
	}

	subQueryMods := FriendshipWhereToMods(m, !hasForeignKeyInRoot, parentTable)
	if len(subQueryMods) > 0 {
		subQuery := models.Friendships(append(subQueryMods, qm.Select("1"))...)
		queryMods = appendSubQuery(queryMods, subQuery.Query)
	}
	return queryMods
}

func FriendshipWhereToMods(m *graphql_models.FriendshipWhere, withPrimaryID bool, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	if withPrimaryID {
		queryMods = append(queryMods, IDFilterToMods(m.ID, models.FriendshipColumns.ID)...)
	}
	queryMods = append(queryMods, IntFilterToMods(m.CreatedAt, models.FriendshipColumns.CreatedAt)...)
	queryMods = append(queryMods, UserWhereSubqueryToMods(m.Users, "", models.TableNames.Friendship)...)
	if m.Or != nil {
		queryMods = append(queryMods, qm.Or2(qm.Expr(FriendshipWhereToMods(m.Or, true, "")...)))
	}
	if m.And != nil {
		queryMods = append(queryMods, qm.Expr(FriendshipWhereToMods(m.And, true, "")...))
	}

	if len(queryMods) > 0 && parentTable != "" {

	}

	return queryMods
}
func ImageFilterToMods(m *graphql_models.ImageFilter) []qm.QueryMod {
	if m == nil {
		return nil
	}
	if m.Search != nil || m.Where != nil {
		var queryMods []qm.QueryMod
		queryMods = append(queryMods, ImageSearchToMods(m.Search)...)
		queryMods = append(queryMods, ImageWhereToMods(m.Where, true, "")...)
		if len(queryMods) > 0 {
			return []qm.QueryMod{
				qm.Expr(queryMods...),
			}
		}
	}
	return nil
}
func ImageSearchToMods(search *string) []qm.QueryMod {

	return nil
}
func ImageVariationFilterToMods(m *graphql_models.ImageVariationFilter) []qm.QueryMod {
	if m == nil {
		return nil
	}
	if m.Search != nil || m.Where != nil {
		var queryMods []qm.QueryMod
		queryMods = append(queryMods, ImageVariationSearchToMods(m.Search)...)
		queryMods = append(queryMods, ImageVariationWhereToMods(m.Where, true, "")...)
		if len(queryMods) > 0 {
			return []qm.QueryMod{
				qm.Expr(queryMods...),
			}
		}
	}
	return nil
}
func ImageVariationSearchToMods(search *string) []qm.QueryMod {

	return nil
}
func ImageVariationWhereSubqueryToMods(m *graphql_models.ImageVariationWhere, foreignColumn string, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	hasForeignKeyInRoot := foreignColumn != ""
	if hasForeignKeyInRoot {
		queryMods = append(queryMods, IDFilterToMods(m.ID, foreignColumn)...)
	}

	subQueryMods := ImageVariationWhereToMods(m, !hasForeignKeyInRoot, parentTable)
	if len(subQueryMods) > 0 {
		subQuery := models.ImageVariations(append(subQueryMods, qm.Select("1"))...)
		queryMods = appendSubQuery(queryMods, subQuery.Query)
	}
	return queryMods
}

func ImageVariationWhereToMods(m *graphql_models.ImageVariationWhere, withPrimaryID bool, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	if withPrimaryID {
		queryMods = append(queryMods, IDFilterToMods(m.ID, models.ImageVariationColumns.ID)...)
	}
	queryMods = append(queryMods, ImageWhereSubqueryToMods(m.Image, models.ImageVariationColumns.ImageID, models.TableNames.ImageVariation)...)
	if m.Or != nil {
		queryMods = append(queryMods, qm.Or2(qm.Expr(ImageVariationWhereToMods(m.Or, true, "")...)))
	}
	if m.And != nil {
		queryMods = append(queryMods, qm.Expr(ImageVariationWhereToMods(m.And, true, "")...))
	}

	if len(queryMods) > 0 && parentTable != "" {

		if parentTable == models.TableNames.Image {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.ImageVariation, models.ImageVariationColumns.ImageID, parentTable)))
		}
	}

	return queryMods
}
func ImageWhereSubqueryToMods(m *graphql_models.ImageWhere, foreignColumn string, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	hasForeignKeyInRoot := foreignColumn != ""
	if hasForeignKeyInRoot {
		queryMods = append(queryMods, IDFilterToMods(m.ID, foreignColumn)...)
	}

	subQueryMods := ImageWhereToMods(m, !hasForeignKeyInRoot, parentTable)
	if len(subQueryMods) > 0 {
		subQuery := models.Images(append(subQueryMods, qm.Select("1"))...)
		queryMods = appendSubQuery(queryMods, subQuery.Query)
	}
	return queryMods
}

func ImageWhereToMods(m *graphql_models.ImageWhere, withPrimaryID bool, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	if withPrimaryID {
		queryMods = append(queryMods, IDFilterToMods(m.ID, models.ImageColumns.ID)...)
	}
	queryMods = append(queryMods, PostWhereSubqueryToMods(m.Post, models.ImageColumns.PostID, models.TableNames.Image)...)
	queryMods = append(queryMods, IntFilterToMods(m.Views, models.ImageColumns.Views)...)
	queryMods = append(queryMods, StringFilterToMods(m.OriginalURL, models.ImageColumns.OriginalURL)...)
	queryMods = append(queryMods, ImageVariationWhereSubqueryToMods(m.ImageVariations, "", models.TableNames.Image)...)
	if m.Or != nil {
		queryMods = append(queryMods, qm.Or2(qm.Expr(ImageWhereToMods(m.Or, true, "")...)))
	}
	if m.And != nil {
		queryMods = append(queryMods, qm.Expr(ImageWhereToMods(m.And, true, "")...))
	}

	if len(queryMods) > 0 && parentTable != "" {

		if parentTable == models.TableNames.Post {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.Image, models.ImageColumns.PostID, parentTable)))
		}
	}

	return queryMods
}
func LikeFilterToMods(m *graphql_models.LikeFilter) []qm.QueryMod {
	if m == nil {
		return nil
	}
	if m.Search != nil || m.Where != nil {
		var queryMods []qm.QueryMod
		queryMods = append(queryMods, LikeSearchToMods(m.Search)...)
		queryMods = append(queryMods, LikeWhereToMods(m.Where, true, "")...)
		if len(queryMods) > 0 {
			return []qm.QueryMod{
				qm.Expr(queryMods...),
			}
		}
	}
	return nil
}
func LikeSearchToMods(search *string) []qm.QueryMod {

	return nil
}
func LikeWhereSubqueryToMods(m *graphql_models.LikeWhere, foreignColumn string, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	hasForeignKeyInRoot := foreignColumn != ""
	if hasForeignKeyInRoot {
		queryMods = append(queryMods, IDFilterToMods(m.ID, foreignColumn)...)
	}

	subQueryMods := LikeWhereToMods(m, !hasForeignKeyInRoot, parentTable)
	if len(subQueryMods) > 0 {
		subQuery := models.Likes(append(subQueryMods, qm.Select("1"))...)
		queryMods = appendSubQuery(queryMods, subQuery.Query)
	}
	return queryMods
}

func LikeWhereToMods(m *graphql_models.LikeWhere, withPrimaryID bool, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	if withPrimaryID {
		queryMods = append(queryMods, IDFilterToMods(m.ID, models.LikeColumns.ID)...)
	}
	queryMods = append(queryMods, PostWhereSubqueryToMods(m.Post, models.LikeColumns.PostID, models.TableNames.Like)...)
	queryMods = append(queryMods, UserWhereSubqueryToMods(m.User, models.LikeColumns.UserID, models.TableNames.Like)...)
	queryMods = append(queryMods, StringFilterToMods(m.LikeType, models.LikeColumns.LikeType)...)
	queryMods = append(queryMods, IntFilterToMods(m.CreatedAt, models.LikeColumns.CreatedAt)...)
	if m.Or != nil {
		queryMods = append(queryMods, qm.Or2(qm.Expr(LikeWhereToMods(m.Or, true, "")...)))
	}
	if m.And != nil {
		queryMods = append(queryMods, qm.Expr(LikeWhereToMods(m.And, true, "")...))
	}

	if len(queryMods) > 0 && parentTable != "" {

		if parentTable == models.TableNames.Post {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.Like, models.LikeColumns.PostID, parentTable)))
		}
		if parentTable == models.TableNames.User {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.Like, models.LikeColumns.UserID, parentTable)))
		}
	}

	return queryMods
}
func PostFilterToMods(m *graphql_models.PostFilter) []qm.QueryMod {
	if m == nil {
		return nil
	}
	if m.Search != nil || m.Where != nil {
		var queryMods []qm.QueryMod
		queryMods = append(queryMods, PostSearchToMods(m.Search)...)
		queryMods = append(queryMods, PostWhereToMods(m.Where, true, "")...)
		if len(queryMods) > 0 {
			return []qm.QueryMod{
				qm.Expr(queryMods...),
			}
		}
	}
	return nil
}
func PostSearchToMods(search *string) []qm.QueryMod {

	return nil
}
func PostWhereSubqueryToMods(m *graphql_models.PostWhere, foreignColumn string, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	hasForeignKeyInRoot := foreignColumn != ""
	if hasForeignKeyInRoot {
		queryMods = append(queryMods, IDFilterToMods(m.ID, foreignColumn)...)
	}

	subQueryMods := PostWhereToMods(m, !hasForeignKeyInRoot, parentTable)
	if len(subQueryMods) > 0 {
		subQuery := models.Posts(append(subQueryMods, qm.Select("1"))...)
		queryMods = appendSubQuery(queryMods, subQuery.Query)
	}
	return queryMods
}

func PostWhereToMods(m *graphql_models.PostWhere, withPrimaryID bool, parentTable string) []qm.QueryMod {
	if m == nil {
		return nil
	}
	var queryMods []qm.QueryMod

	if withPrimaryID {
		queryMods = append(queryMods, IDFilterToMods(m.ID, models.PostColumns.ID)...)
	}
	queryMods = append(queryMods, StringFilterToMods(m.Content, models.PostColumns.Content)...)
	queryMods = append(queryMods, UserWhereSubqueryToMods(m.User, models.PostColumns.UserID, models.TableNames.Post)...)
	queryMods = append(queryMods, CommentWhereSubqueryToMods(m.Comments, "", models.TableNames.Post)...)
	queryMods = append(queryMods, ImageWhereSubqueryToMods(m.Images, "", models.TableNames.Post)...)
	queryMods = append(queryMods, LikeWhereSubqueryToMods(m.Likes, "", models.TableNames.Post)...)
	if m.Or != nil {
		queryMods = append(queryMods, qm.Or2(qm.Expr(PostWhereToMods(m.Or, true, "")...)))
	}
	if m.And != nil {
		queryMods = append(queryMods, qm.Expr(PostWhereToMods(m.And, true, "")...))
	}

	if len(queryMods) > 0 && parentTable != "" {

		if parentTable == models.TableNames.User {
			queryMods = append(queryMods, qm.Where(fmt.Sprintf("%v.%v = %v.id", models.TableNames.Post, models.PostColumns.UserID, parentTable)))
		}
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
	queryMods = append(queryMods, StringFilterToMods(m.Email, models.UserColumns.Email)...)
	queryMods = append(queryMods, StringFilterToMods(m.Password, models.UserColumns.Password)...)
	queryMods = append(queryMods, CommentWhereSubqueryToMods(m.Comments, "", models.TableNames.User)...)
	queryMods = append(queryMods, CommentLikeWhereSubqueryToMods(m.CommentLikes, "", models.TableNames.User)...)
	queryMods = append(queryMods, LikeWhereSubqueryToMods(m.Likes, "", models.TableNames.User)...)
	queryMods = append(queryMods, PostWhereSubqueryToMods(m.Posts, "", models.TableNames.User)...)
	queryMods = append(queryMods, FriendshipWhereSubqueryToMods(m.Friendships, "", models.TableNames.User)...)
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
