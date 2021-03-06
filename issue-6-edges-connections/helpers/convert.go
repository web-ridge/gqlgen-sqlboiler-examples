package helpers

import (
	"github.com/web-ridge/gqlgen-sqlboiler-examples/issue-6-edges-connections/graphql_models"
	"github.com/web-ridge/gqlgen-sqlboiler-examples/issue-6-edges-connections/models"
	"github.com/web-ridge/utils-go/boilergql/v3"

	"github.com/volatiletech/null/v8"
)

type UserSort string

const (
	UserSortID		UserSort	= "id"
	UserSortFirstName	UserSort	= "firstName"
	UserSortLastName	UserSort	= "lastName"
	UserSortAge		UserSort	= "age"
	UserSortEmail		UserSort	= "email"
)

var UserSortDBValue = map[graphql_models.UserSort]UserSort{
	graphql_models.UserSortID:		UserSortID,
	graphql_models.UserSortFirstName:	UserSortFirstName,
	graphql_models.UserSortLastName:	UserSortLastName,
	graphql_models.UserSortAge:		UserSortAge,
	graphql_models.UserSortEmail:		UserSortEmail,
}
var UserSortAPIValue = map[UserSort]graphql_models.UserSort{
	UserSortID:		graphql_models.UserSortID,
	UserSortFirstName:	graphql_models.UserSortFirstName,
	UserSortLastName:	graphql_models.UserSortLastName,
	UserSortAge:		graphql_models.UserSortAge,
	UserSortEmail:		graphql_models.UserSortEmail,
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
	s := UserSortAPIValue[UserSort(v)]
	return s
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
	s := UserSortDBValue[v]
	return string(s)
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
		ID:		UserIDToGraphQL(m.ID),
		FirstName:	m.FirstName,
		LastName:	m.LastName,
		Age:		boilergql.UintToInt(m.Age),
		Email:		boilergql.NullDotStringToPointerString(m.Email),
	}

	return r
}

func UserID(v string) uint {
	return boilergql.IDToBoilerUint(v)
}

func UserIDs(a []string) []uint {
	return boilergql.IDsToBoilerUint(a)
}
