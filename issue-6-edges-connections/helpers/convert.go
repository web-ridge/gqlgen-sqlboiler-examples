// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package helpers

import (
	null "github.com/volatiletech/null/v8"
	"github.com/web-ridge/gqlgen-sqlboiler-examples/issue-6-edges-connections/graphql_models"
	"github.com/web-ridge/gqlgen-sqlboiler-examples/issue-6-edges-connections/models"
	"github.com/web-ridge/utils-go/boilergql"
)

func NullDotStringToPointerSortDirection(v null.String) *graphql_models.SortDirection {
	s := StringToSortDirection(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToSortDirection(v null.String) graphql_models.SortDirection {
	if !v.Valid {
		return ""
	}
	return StringToSortDirection(v.String)
}

func StringToSortDirection(v string) graphql_models.SortDirection {
	if v == "asc" {
		return graphql_models.SortDirectionAsc
	}
	if v == "desc" {
		return graphql_models.SortDirectionDesc
	}
	return ""
}

func StringToPointerSortDirection(v string) *graphql_models.SortDirection {
	s := StringToSortDirection(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerSortDirectionToString(v *graphql_models.SortDirection) string {
	if v == nil {
		return ""
	}
	return SortDirectionToString(*v)
}

func PointerSortDirectionToNullDotString(v *graphql_models.SortDirection) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return SortDirectionToNullDotString(*v)
}

func SortDirectionToNullDotString(v graphql_models.SortDirection) null.String {
	s := SortDirectionToString(v)
	return null.NewString(s, s != "")
}

func SortDirectionToString(v graphql_models.SortDirection) string {
	if v == graphql_models.SortDirectionAsc {
		return "asc"
	}
	if v == graphql_models.SortDirectionDesc {
		return "desc"
	}
	return ""
}

func NullDotStringToPointerUserSort(v null.String) *graphql_models.UserSort {
	s := StringToUserSort(v.String)
	if s == "" {
		return nil
	}
	return &s
}

func NullDotStringToUserSort(v null.String) graphql_models.UserSort {
	if !v.Valid {
		return ""
	}
	return StringToUserSort(v.String)
}

func StringToUserSort(v string) graphql_models.UserSort {
	if v == "id" {
		return graphql_models.UserSortID
	}
	if v == "firstName" {
		return graphql_models.UserSortFirstName
	}
	if v == "lastName" {
		return graphql_models.UserSortLastName
	}
	if v == "age" {
		return graphql_models.UserSortAge
	}
	if v == "email" {
		return graphql_models.UserSortEmail
	}
	return ""
}

func StringToPointerUserSort(v string) *graphql_models.UserSort {
	s := StringToUserSort(v)
	if s == "" {
		return nil
	}
	return &s
}

func PointerUserSortToString(v *graphql_models.UserSort) string {
	if v == nil {
		return ""
	}
	return UserSortToString(*v)
}

func PointerUserSortToNullDotString(v *graphql_models.UserSort) null.String {
	if v == nil {
		return null.NewString("", false)
	}
	return UserSortToNullDotString(*v)
}

func UserSortToNullDotString(v graphql_models.UserSort) null.String {
	s := UserSortToString(v)
	return null.NewString(s, s != "")
}

func UserSortToString(v graphql_models.UserSort) string {
	if v == graphql_models.UserSortID {
		return "id"
	}
	if v == graphql_models.UserSortFirstName {
		return "firstName"
	}
	if v == graphql_models.UserSortLastName {
		return "lastName"
	}
	if v == graphql_models.UserSortAge {
		return "age"
	}
	if v == graphql_models.UserSortEmail {
		return "email"
	}
	return ""
}

func UserWithUintID(id uint) *graphql_models.User {
	return &graphql_models.User{
		ID: UserIDToGraphQL(id),
	}
}

func UserWithIntID(id int) *graphql_models.User {
	return UserWithUintID(uint(id))
}

func UserWithNullDotUintID(id null.Uint) *graphql_models.User {
	return UserWithUintID(id.Uint)
}

func UserWithNullDotIntID(id null.Int) *graphql_models.User {
	return UserWithUintID(uint(id.Int))
}

func UsersToGraphQL(am []*models.User) []*graphql_models.User {
	ar := make([]*graphql_models.User, len(am))
	for i, m := range am {
		ar[i] = UserToGraphQL(m)
	}
	return ar
}

func UserIDToGraphQL(v uint) string {
	return boilergql.IDToGraphQL(v, models.TableNames.User)
}

func UserToGraphQL(m *models.User) *graphql_models.User {
	if m == nil {
		return nil
	}

	r := &graphql_models.User{
		ID:        UserIDToGraphQL(m.ID),
		FirstName: m.FirstName,
		LastName:  m.LastName,
		Age:       boilergql.UintToInt(m.Age),
		Email:     boilergql.NullDotStringToPointerString(m.Email),
	}

	return r
}

func UserID(v string) uint {
	return boilergql.IDToBoilerUint(v)
}

func UserIDs(a []string) []uint {
	return boilergql.IDsToBoilerUint(a)
}
