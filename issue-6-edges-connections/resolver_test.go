package main

import (
	"context"
	"database/sql"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/web-ridge/utils-go/boilergql"

	//"github.com/stretchr/testify/assert"

	"github.com/web-ridge/gqlgen-sqlboiler-examples/issue-6-edges-connections/models"

	fm "github.com/web-ridge/gqlgen-sqlboiler-examples/issue-6-edges-connections/graphql_models"
)

func getDependencies() (context.Context, *sql.DB, *Resolver) {
	operationCtx := &graphql.OperationContext{
		Variables: map[string]interface{}{},
	}
	fieldCtx := &graphql.FieldContext{
		Args: map[string]interface{}{},
	}
	ctx := context.Background()
	ctx = graphql.WithOperationContext(ctx, operationCtx)
	ctx = graphql.WithFieldContext(ctx, fieldCtx)
	db := getDatabase()
	resolver := &Resolver{
		db: db,
	}
	return ctx, db, resolver
}

func TestConnections(t *testing.T) {
	ctx, db, resolver := getDependencies()

	_, err := models.Users().DeleteAll(ctx, db)
	handleErr(t, err)

	// Eve
	for i := 0; i < 20; i++ {
		number := 20 - i
		user := models.User{
			FirstName: fmt.Sprintf("Eve%02d", number),
			LastName:  fmt.Sprintf("Eve%02d", number),
		}
		err = user.Insert(ctx, db, boil.Infer())
		handleErr(t, err)
	}

	// Adam
	for i := 0; i < 20; i++ {
		number := 20 - i
		user := models.User{
			FirstName: fmt.Sprintf("Adam%02d", number),
			LastName:  fmt.Sprintf("Adam%02d", number),
		}
		err = user.Insert(ctx, db, boil.Infer())
		handleErr(t, err)
	}

	////
	userConnection, err := resolver.Query().Users(ctx, boilergql.ConnectionPagination{
		Forward: &boilergql.ConnectionForwardPagination{
			First: 10,
			After: nil,
		},
	}, nil, nil)
	handleErr(t, err)
	if assert.Equal(t, 10, len(userConnection.Edges), "edges not equal") {
		assert.Equal(t, "Eve20", firstNameFromUser(userConnection.Edges[0]), "edges not equal")
		assert.Equal(t, "Eve11", firstNameFromUser(userConnection.Edges[9]), "edges not equal")
		assert.Equal(t, true, userConnection.PageInfo.HasNextPage, "nextPage not equal")
		assert.Equal(t, false, userConnection.PageInfo.HasPreviousPage, "previousPage not equal")
	}

	endCursor := userConnection.PageInfo.EndCursor

	userConnection, err = resolver.Query().Users(ctx, boilergql.ConnectionPagination{
		Forward: &boilergql.ConnectionForwardPagination{
			First: 10,
			After: endCursor,
		},
	}, nil, nil)
	handleErr(t, err)
	if assert.Equal(t, 10, len(userConnection.Edges), "edges not equal") {
		assert.Equal(t, "Eve10", firstNameFromUser(userConnection.Edges[0]), "edges not equal")
		assert.Equal(t, "Eve01", firstNameFromUser(userConnection.Edges[9]), "edges not equal")
		assert.Equal(t, true, userConnection.PageInfo.HasNextPage, "nextPage not equal")
		assert.Equal(t, true, userConnection.PageInfo.HasPreviousPage, "previousPage not equal")
	}

	endCursor = userConnection.PageInfo.EndCursor

	userConnection, err = resolver.Query().Users(ctx, boilergql.ConnectionPagination{
		Forward: &boilergql.ConnectionForwardPagination{
			First: 10,
			After: endCursor,
		},
	}, nil, nil)
	handleErr(t, err)
	if assert.Equal(t, 10, len(userConnection.Edges), "edges not equal") {
		assert.Equal(t, "Adam20", firstNameFromUser(userConnection.Edges[0]), "edges not equal")
		assert.Equal(t, "Adam11", firstNameFromUser(userConnection.Edges[9]), "edges not equal")
		assert.Equal(t, true, userConnection.PageInfo.HasNextPage, "nextPage not equal")
		assert.Equal(t, true, userConnection.PageInfo.HasPreviousPage, "previousPage not equal")
	}
	userConnection, err = resolver.Query().Users(ctx, boilergql.ConnectionPagination{
		Forward: &boilergql.ConnectionForwardPagination{
			First: 100,
			After: nil,
		},
	}, nil, nil)
	handleErr(t, err)
	if assert.Equal(t, 40, len(userConnection.Edges), "edges not equal") {
		assert.Equal(t, "Eve20", firstNameFromUser(userConnection.Edges[0]), "edges not equal")
		assert.Equal(t, "Adam01", firstNameFromUser(userConnection.Edges[39]), "edges not equal")
		assert.Equal(t, false, userConnection.PageInfo.HasNextPage, "nextPage not equal")
		assert.Equal(t, false, userConnection.PageInfo.HasPreviousPage, "previousPage not equal")
	}
	// BACKWARD PAGINATION
	userConnection, err = resolver.Query().Users(ctx, boilergql.ConnectionPagination{
		Backward: &boilergql.ConnectionBackwardPagination{
			Last:   10,
			Before: nil,
		},
	}, nil, nil)
	handleErr(t, err)
	if assert.Equal(t, 10, len(userConnection.Edges), "edges not equal") {
		assert.Equal(t, "Adam10", firstNameFromUser(userConnection.Edges[0]), "edges not equal")
		assert.Equal(t, "Adam01", firstNameFromUser(userConnection.Edges[9]), "edges not equal")
		assert.Equal(t, true, userConnection.PageInfo.HasPreviousPage, "nextPage not equal")
		assert.Equal(t, false, userConnection.PageInfo.HasNextPage, "previousPage not equal")
	}
	startCursor := userConnection.PageInfo.StartCursor

	userConnection, err = resolver.Query().Users(ctx, boilergql.ConnectionPagination{
		Backward: &boilergql.ConnectionBackwardPagination{
			Last:   10,
			Before: startCursor,
		},
	}, nil, nil)
	handleErr(t, err)

	if assert.Equal(t, 10, len(userConnection.Edges), "edges not equal") {
		assert.Equal(t, "Adam20", firstNameFromUser(userConnection.Edges[0]), "edges not equal")
		assert.Equal(t, "Adam11", firstNameFromUser(userConnection.Edges[9]), "edges not equal")
		assert.Equal(t, true, userConnection.PageInfo.HasPreviousPage, "nextPage not equal")
		assert.Equal(t, true, userConnection.PageInfo.HasNextPage, "previousPage not equal")
	}
	userConnection, err = resolver.Query().Users(ctx, boilergql.ConnectionPagination{
		Backward: &boilergql.ConnectionBackwardPagination{
			Last:   100,
			Before: nil,
		},
	}, nil, nil)
	handleErr(t, err)
	if assert.Equal(t, 40, len(userConnection.Edges), "edges not equal") {
		assert.Equal(t, "Eve20", firstNameFromUser(userConnection.Edges[0]), "edges not equal")
		assert.Equal(t, "Adam01", firstNameFromUser(userConnection.Edges[39]), "edges not equal")
		assert.Equal(t, false, userConnection.PageInfo.HasNextPage, "nextPage not equal")
		assert.Equal(t, false, userConnection.PageInfo.HasPreviousPage, "previousPage not equal")
	}

	// Adam
	for i := 0; i < 20; i++ {
		m := models.User{
			FirstName: fmt.Sprintf("Dirk%02d", i),
			LastName:  fmt.Sprintf("Dirk%02d", i),
		}
		err = m.Insert(ctx, db, boil.Infer())
		handleErr(t, err)
	}

	// With sorting
	sort := []*fm.UserOrdering{
		{
			Sort:      fm.UserSortFirstName,
			Direction: boilergql.SortDirectionAsc,
		},
		{
			Sort:      fm.UserSortLastName,
			Direction: boilergql.SortDirectionAsc,
		},
	}
	userConnection, err = resolver.Query().Users(ctx, boilergql.ConnectionPagination{
		Backward: &boilergql.ConnectionBackwardPagination{
			Last:   20,
			Before: nil,
		},
	}, sort, nil)
	handleErr(t, err)

	startCursor = userConnection.PageInfo.StartCursor
	if assert.Equal(t, 20, len(userConnection.Edges), "edges not equal backward sorting ASC EVE") {
		for _, edge := range userConnection.Edges {
			expected := "Eve"
			if !strings.HasPrefix(firstNameFromUser(edge), expected) {
				t.Errorf("backward should start with %v but is %v", expected, firstNameFromUser(edge))
			}
		}
	}
	userConnection, err = resolver.Query().Users(ctx, boilergql.ConnectionPagination{
		Backward: &boilergql.ConnectionBackwardPagination{
			Last:   20,
			Before: startCursor,
		},
	}, sort, nil)
	handleErr(t, err)
	startCursor = userConnection.PageInfo.StartCursor
	if assert.Equal(t, 20, len(userConnection.Edges), "edges not equal backward sorting ASC Dirk") {
		for _, edge := range userConnection.Edges {
			expected := "Dirk"
			if !strings.HasPrefix(firstNameFromUser(edge), expected) {
				t.Errorf("backward should start with %v but is %v", expected, firstNameFromUser(edge))
			}
		}
	}

	userConnection, err = resolver.Query().Users(ctx, boilergql.ConnectionPagination{
		Backward: &boilergql.ConnectionBackwardPagination{
			Last:   20,
			Before: startCursor,
		},
	}, sort, nil)
	handleErr(t, err)
	if assert.Equal(t, 20, len(userConnection.Edges), "edges not equal backward sorting ASC Adam") {
		for _, edge := range userConnection.Edges {
			expected := "Adam"
			if !strings.HasPrefix(firstNameFromUser(edge), expected) {
				t.Errorf("backward should start with %v but is %v", expected, firstNameFromUser(edge))
			}
		}
	}

	// other sort

	sort = []*fm.UserOrdering{
		{
			Sort:      fm.UserSortFirstName,
			Direction: boilergql.SortDirectionDesc,
		},
		{
			Sort:      fm.UserSortLastName,
			Direction: boilergql.SortDirectionDesc,
		},
	}
	userConnection, err = resolver.Query().Users(ctx, boilergql.ConnectionPagination{
		Backward: &boilergql.ConnectionBackwardPagination{
			Last:   20,
			Before: nil,
		},
	}, sort, nil)
	handleErr(t, err)
	startCursor = userConnection.PageInfo.StartCursor
	if assert.Equal(t, 20, len(userConnection.Edges), "edges not equal backward sorting DESC Adam") {
		for _, edge := range userConnection.Edges {
			expected := "Adam"
			if !strings.HasPrefix(firstNameFromUser(edge), expected) {
				t.Errorf("backward should start with %v but is %v", expected, firstNameFromUser(edge))
			}
		}
	}
	userConnection, err = resolver.Query().Users(ctx, boilergql.ConnectionPagination{
		Backward: &boilergql.ConnectionBackwardPagination{
			Last:   20,
			Before: startCursor,
		},
	}, sort, nil)
	handleErr(t, err)
	startCursor = userConnection.PageInfo.StartCursor
	if assert.Equal(t, 20, len(userConnection.Edges), "edges not equal backward sorting DESC Dirk") {
		for _, edge := range userConnection.Edges {
			expected := "Dirk"
			if !strings.HasPrefix(firstNameFromUser(edge), expected) {
				t.Errorf("backward should start with %v but is %v", expected, firstNameFromUser(edge))
			}
		}
	}

	userConnection, err = resolver.Query().Users(ctx, boilergql.ConnectionPagination{
		Backward: &boilergql.ConnectionBackwardPagination{
			Last:   20,
			Before: startCursor,
		},
	}, sort, nil)
	handleErr(t, err)

	startCursor = userConnection.PageInfo.StartCursor //nolint: ineffassign,staticcheck
	if assert.Equal(t, 20, len(userConnection.Edges), "edges not equal backward sorting DESC Eve") {
		for _, edge := range userConnection.Edges {
			expected := "Eve"
			if !strings.HasPrefix(firstNameFromUser(edge), expected) {
				t.Errorf("backward should start with %v but is %v", expected, firstNameFromUser(edge))
			}
		}
	}
}

func TestAscDescSortingAtOnce(t *testing.T) {
	ctx, db, resolver := getDependencies()
	_, err := models.Users().DeleteAll(ctx, db)
	handleErr(t, err)

	for i := 0; i < 20; i++ {
		user := models.User{
			FirstName: "Adam",
			LastName:  "",
			Age:       uint(50 - i),
		}
		err = user.Insert(ctx, db, boil.Infer())
		handleErr(t, err)
	}
	for i := 0; i < 20; i++ {
		user := models.User{
			FirstName: "Zara",
			LastName:  "",
			Age:       uint(50 - i),
		}
		err = user.Insert(ctx, db, boil.Infer())
		handleErr(t, err)
	}
	for i := 0; i < 20; i++ {
		user := models.User{
			FirstName: "Eve",
			LastName:  "",
			Age:       uint(i + 20),
		}
		err = user.Insert(ctx, db, boil.Infer())
		handleErr(t, err)
	}
	fmt.Println(resolver)

	dbUsers, err := models.Users().All(ctx, db)
	handleErr(t, err)
	for _, firstNameDirection := range boilergql.AllSortDirection {
		for _, ageDirection := range boilergql.AllSortDirection {
			testIdentifier := fmt.Sprintf("sortFirstName%v and sortAge%v", firstNameDirection, ageDirection)
			t.Logf("Starting test: %v", testIdentifier)
			ordering := []*fm.UserOrdering{
				{
					Sort:      fm.UserSortFirstName,
					Direction: firstNameDirection,
				},
				{
					Sort:      fm.UserSortAge,
					Direction: ageDirection,
				},
			}
			allExpectedFirstNames, allExpectedAges := sortFirstNameAndAge(dbUsers, firstNameDirection, ageDirection)

			// forward
			var endCursor *string
			for i := 0; i < 3; i++ {
				userConnection, err := resolver.Query().Users(ctx, boilergql.ConnectionPagination{
					Forward: &boilergql.ConnectionForwardPagination{
						First: 20,
						After: endCursor,
					},
				}, ordering, nil)
				handleErr(t, err)
				endCursor = userConnection.PageInfo.EndCursor

				offset := i * 20
				firstNames, ages := pickFirstNamesAndEdges(userConnection)
				expectedFirstNames, expectedAges := allExpectedFirstNames[offset:offset+20],
					allExpectedAges[offset:offset+20]

				assert.Equal(t, i < 2, userConnection.PageInfo.HasNextPage, testIdentifier+
					" > forward pagination > has next page > "+fmt.Sprintf("offset: %v", offset))
				assert.Equal(t, i > 0, userConnection.PageInfo.HasPreviousPage, testIdentifier+
					" > forward pagination > has previous page > "+fmt.Sprintf("offset: %v", offset))
				assert.Equal(t, commaS(expectedFirstNames), commaS(firstNames), testIdentifier+
					" > forward pagination > firstnames not equal > "+fmt.Sprintf("offset: %v", offset))
				assert.Equal(t, commaU(expectedAges), commaU(ages), testIdentifier+
					" > forward pagination > ages not equal > "+fmt.Sprintf("offset: %v", offset))
			}

			// backward
			var startCursorBackward *string
			for i := 0; i < 3; i++ {
				userConnection, err := resolver.Query().Users(ctx, boilergql.ConnectionPagination{
					Backward: &boilergql.ConnectionBackwardPagination{
						Last:   20,
						Before: startCursorBackward,
					},
				}, ordering, nil)
				handleErr(t, err)
				startCursorBackward = userConnection.PageInfo.StartCursor

				offset := len(dbUsers) - i*20
				firstNames, ages := pickFirstNamesAndEdges(userConnection)
				expectedFirstNames, expectedAges := allExpectedFirstNames[offset-20:offset], allExpectedAges[offset-20:offset]
				assert.Equal(t, i > 0, userConnection.PageInfo.HasNextPage, testIdentifier+
					" > backward pagination > has next page > "+fmt.Sprintf("offset: %v", offset))
				assert.Equal(t, i < 2, userConnection.PageInfo.HasPreviousPage, testIdentifier+
					" > backward pagination > has previous page > "+fmt.Sprintf("offset: %v", offset))
				assert.Equal(t, commaS(expectedFirstNames), commaS(firstNames), testIdentifier+
					" > backward pagination > firstnames not equal > "+fmt.Sprintf("offset: %v", offset))
				assert.Equal(t, commaU(expectedAges), commaU(ages), testIdentifier+
					" > backward pagination > ages not equal > "+fmt.Sprintf("offset: %v", offset))
			}
		}
	}
}

func sortFirstNameAsc(a, b string) bool {
	return a < b
}

func sortAgeAsc(a, b uint) bool {
	return a < b
}

func sortFirstNameDesc(a, b string) bool {
	return a > b
}

func sortAgeDesc(a, b uint) bool {
	return a > b
}

func commaS(a []string) string {
	return strings.Join(a, ", ")
}

func commaU(a []uint) string {
	return uintJoin(a, ", ")
}

func uintJoin(a []uint, sep string) string {
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(int(v))
	}
	return strings.Join(b, sep)
}

func pickFirstNamesAndEdges(uc *fm.UserConnection) ([]string, []uint) {
	firstNames := make([]string, len(uc.Edges))
	ages := make([]uint, len(uc.Edges))
	for i, e := range uc.Edges {
		firstNames[i] = e.Node.FirstName
		ages[i] = uint(e.Node.Age)
	}
	return firstNames, ages
}

func pickFirstNames(users []*models.User) []string {
	pa := make([]string, len(users))
	for i, u := range users {
		pa[i] = u.FirstName
	}
	return pa
}

func pickAges(users []*models.User) []uint {
	pa := make([]uint, len(users))
	for i, u := range users {
		pa[i] = u.Age
	}
	return pa
}

func sortFirstNameAndAge(
	users []*models.User,
	firstNameDirection boilergql.SortDirection,
	ageDirection boilergql.SortDirection,
) ([]string, []uint) {
	switch firstNameDirection {
	case boilergql.SortDirectionAsc:
		switch ageDirection {
		case boilergql.SortDirectionAsc:
			return sortFirstNameAscAndAgeAsc(users)
		case boilergql.SortDirectionDesc:
			return sortFirstNameAscAndAgeDesc(users)
		}
	case boilergql.SortDirectionDesc:
		switch ageDirection {
		case boilergql.SortDirectionAsc:
			return sortFirstNameDescAndAgeAsc(users)
		case boilergql.SortDirectionDesc:
			return sortFirstNameDescAndAgeDesc(users)
		}
	}
	return nil, nil
}

func sortFirstNameAscAndAgeDesc(users []*models.User) ([]string, []uint) {
	sort.Slice(users, func(i, j int) bool {
		a, b := users[i], users[j]
		if a.FirstName == b.FirstName {
			return sortAgeDesc(a.Age, b.Age)
		}
		return sortFirstNameAsc(a.FirstName, b.FirstName)
	})
	return pickFirstNames(users), pickAges(users)
}

func sortFirstNameDescAndAgeAsc(users []*models.User) ([]string, []uint) {
	sort.Slice(users, func(i, j int) bool {
		a, b := users[i], users[j]
		if a.FirstName == b.FirstName {
			return sortAgeAsc(a.Age, b.Age)
		}
		return sortFirstNameDesc(a.FirstName, b.FirstName)
	})
	return pickFirstNames(users), pickAges(users)
}

func sortFirstNameAscAndAgeAsc(users []*models.User) ([]string, []uint) {
	sort.Slice(users, func(i, j int) bool {
		a, b := users[i], users[j]
		if a.FirstName == b.FirstName {
			return sortAgeAsc(a.Age, b.Age)
		}
		return sortFirstNameAsc(a.FirstName, b.FirstName)
	})
	return pickFirstNames(users), pickAges(users)
}

func sortFirstNameDescAndAgeDesc(users []*models.User) ([]string, []uint) {
	sort.Slice(users, func(i, j int) bool {
		a, b := users[i], users[j]
		if a.FirstName == b.FirstName {
			return sortAgeDesc(a.Age, b.Age)
		}
		return sortFirstNameDesc(a.FirstName, b.FirstName)
	})
	return pickFirstNames(users), pickAges(users)
}

func firstNameFromUser(e *fm.UserEdge) string {
	return e.Node.FirstName
}

func handleErr(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
