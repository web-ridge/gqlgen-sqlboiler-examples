// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql_models

import (
	"fmt"
	"io"
	"strconv"
)

type Node interface {
	IsNode()
}

type BooleanFilter struct {
	EqualTo    *bool `json:"equalTo"`
	NotEqualTo *bool `json:"notEqualTo"`
}

type ConnectionBackwardPagination struct {
	Last   int     `json:"last"`
	Before *string `json:"before"`
}

type ConnectionForwardPagination struct {
	First int     `json:"first"`
	After *string `json:"after"`
}

type ConnectionPagination struct {
	Forward  *ConnectionForwardPagination  `json:"forward"`
	Backward *ConnectionBackwardPagination `json:"backward"`
}

type FloatFilter struct {
	EqualTo           *float64  `json:"equalTo"`
	NotEqualTo        *float64  `json:"notEqualTo"`
	LessThan          *float64  `json:"lessThan"`
	LessThanOrEqualTo *float64  `json:"lessThanOrEqualTo"`
	MoreThan          *float64  `json:"moreThan"`
	MoreThanOrEqualTo *float64  `json:"moreThanOrEqualTo"`
	In                []float64 `json:"in"`
	NotIn             []float64 `json:"notIn"`
}

type IDFilter struct {
	EqualTo    *string  `json:"equalTo"`
	NotEqualTo *string  `json:"notEqualTo"`
	In         []string `json:"in"`
	NotIn      []string `json:"notIn"`
}

type IntFilter struct {
	EqualTo           *int  `json:"equalTo"`
	NotEqualTo        *int  `json:"notEqualTo"`
	LessThan          *int  `json:"lessThan"`
	LessThanOrEqualTo *int  `json:"lessThanOrEqualTo"`
	MoreThan          *int  `json:"moreThan"`
	MoreThanOrEqualTo *int  `json:"moreThanOrEqualTo"`
	In                []int `json:"in"`
	NotIn             []int `json:"notIn"`
}

type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor"`
	EndCursor       *string `json:"endCursor"`
}

type StringFilter struct {
	EqualTo            *string  `json:"equalTo"`
	NotEqualTo         *string  `json:"notEqualTo"`
	In                 []string `json:"in"`
	NotIn              []string `json:"notIn"`
	StartWith          *string  `json:"startWith"`
	NotStartWith       *string  `json:"notStartWith"`
	EndWith            *string  `json:"endWith"`
	NotEndWith         *string  `json:"notEndWith"`
	Contain            *string  `json:"contain"`
	NotContain         *string  `json:"notContain"`
	StartWithStrict    *string  `json:"startWithStrict"`
	NotStartWithStrict *string  `json:"notStartWithStrict"`
	EndWithStrict      *string  `json:"endWithStrict"`
	NotEndWithStrict   *string  `json:"notEndWithStrict"`
	ContainStrict      *string  `json:"containStrict"`
	NotContainStrict   *string  `json:"notContainStrict"`
}

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func (User) IsNode() {}

type UserConnection struct {
	Edges    []*UserEdge `json:"edges"`
	PageInfo *PageInfo   `json:"pageInfo"`
}

type UserEdge struct {
	Cursor string `json:"cursor"`
	Node   *User  `json:"node"`
}

type UserFilter struct {
	Search *string    `json:"search"`
	Where  *UserWhere `json:"where"`
}

type UserOrdering struct {
	Sort      UserSort      `json:"sort"`
	Direction SortDirection `json:"direction"`
}

type UserWhere struct {
	ID        *IDFilter     `json:"id"`
	FirstName *StringFilter `json:"firstName"`
	LastName  *StringFilter `json:"lastName"`
	Email     *StringFilter `json:"email"`
	Or        *UserWhere    `json:"or"`
	And       *UserWhere    `json:"and"`
}

type SortDirection string

const (
	SortDirectionAsc  SortDirection = "ASC"
	SortDirectionDesc SortDirection = "DESC"
)

var AllSortDirection = []SortDirection{
	SortDirectionAsc,
	SortDirectionDesc,
}

func (e SortDirection) IsValid() bool {
	switch e {
	case SortDirectionAsc, SortDirectionDesc:
		return true
	}
	return false
}

func (e SortDirection) String() string {
	return string(e)
}

func (e *SortDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SortDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SortDirection", str)
	}
	return nil
}

func (e SortDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserSort string

const (
	UserSortID        UserSort = "ID"
	UserSortFirstName UserSort = "FIRST_NAME"
	UserSortLastName  UserSort = "LAST_NAME"
	UserSortEmail     UserSort = "EMAIL"
)

var AllUserSort = []UserSort{
	UserSortID,
	UserSortFirstName,
	UserSortLastName,
	UserSortEmail,
}

func (e UserSort) IsValid() bool {
	switch e {
	case UserSortID, UserSortFirstName, UserSortLastName, UserSortEmail:
		return true
	}
	return false
}

func (e UserSort) String() string {
	return string(e)
}

func (e *UserSort) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserSort(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserSort", str)
	}
	return nil
}

func (e UserSort) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
