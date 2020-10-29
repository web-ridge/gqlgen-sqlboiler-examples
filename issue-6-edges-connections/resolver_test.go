package main

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"

	//"github.com/stretchr/testify/assert"

	"github.com/web-ridge/gqlgen-sqlboiler-examples/issue-6-edges-connections/models"

	fm "github.com/web-ridge/gqlgen-sqlboiler-examples/issue-6-edges-connections/graphql_models"
)

func TestConnections(t *testing.T) {
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

	_, err := models.Users().DeleteAll(ctx, db)
	handleErr(t, err)

	// Eve
	for i := 0; i < 20; i++ {
		number := 20 - i
		user := models.User{
			FirstName: fmt.Sprintf("Eve%02d", number),
			LastName:  fmt.Sprintf("Eve%02d", number),
			Email:     fmt.Sprintf("eve%02d@gmail.com", number),
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
			Email:     fmt.Sprintf("adam%02d@gmail.com", number),
		}
		err = user.Insert(ctx, db, boil.Infer())
		handleErr(t, err)
	}

	////
	userConnection, err := resolver.Query().Users(ctx, fm.ConnectionPagination{
		Forward: &fm.ConnectionForwardPagination{
			First: 10,
			After: nil,
		},
	}, nil, nil)
	handleErr(t, err)
	assert.Equal(t, 10, len(userConnection.Edges), "edges not equal")
	assert.Equal(t, "Eve20", firstNameFromUser(userConnection.Edges[0]), "edges not equal")
	assert.Equal(t, "Eve11", firstNameFromUser(userConnection.Edges[9]), "edges not equal")
	assert.Equal(t, true, userConnection.PageInfo.HasNextPage, "nextPage not equal")
	assert.Equal(t, false, userConnection.PageInfo.HasPreviousPage, "previousPage not equal")

	endCursor := userConnection.PageInfo.EndCursor

	userConnection, err = resolver.Query().Users(ctx, fm.ConnectionPagination{
		Forward: &fm.ConnectionForwardPagination{
			First: 10,
			After: endCursor,
		},
	}, nil, nil)
	handleErr(t, err)
	assert.Equal(t, 10, len(userConnection.Edges), "edges not equal")
	assert.Equal(t, "Eve10", firstNameFromUser(userConnection.Edges[0]), "edges not equal")
	assert.Equal(t, "Eve01", firstNameFromUser(userConnection.Edges[9]), "edges not equal")
	assert.Equal(t, true, userConnection.PageInfo.HasNextPage, "nextPage not equal")
	assert.Equal(t, true, userConnection.PageInfo.HasPreviousPage, "previousPage not equal")

	endCursor = userConnection.PageInfo.EndCursor

	userConnection, err = resolver.Query().Users(ctx, fm.ConnectionPagination{
		Forward: &fm.ConnectionForwardPagination{
			First: 10,
			After: endCursor,
		},
	}, nil, nil)
	handleErr(t, err)
	assert.Equal(t, 10, len(userConnection.Edges), "edges not equal")
	assert.Equal(t, "Adam20", firstNameFromUser(userConnection.Edges[0]), "edges not equal")
	assert.Equal(t, "Adam11", firstNameFromUser(userConnection.Edges[9]), "edges not equal")
	assert.Equal(t, true, userConnection.PageInfo.HasNextPage, "nextPage not equal")
	assert.Equal(t, true, userConnection.PageInfo.HasPreviousPage, "previousPage not equal")

	userConnection, err = resolver.Query().Users(ctx, fm.ConnectionPagination{
		Forward: &fm.ConnectionForwardPagination{
			First: 100,
			After: nil,
		},
	}, nil, nil)
	handleErr(t, err)
	assert.Equal(t, 40, len(userConnection.Edges), "edges not equal")
	assert.Equal(t, "Eve20", firstNameFromUser(userConnection.Edges[0]), "edges not equal")
	assert.Equal(t, "Adam01", firstNameFromUser(userConnection.Edges[39]), "edges not equal")
	assert.Equal(t, false, userConnection.PageInfo.HasNextPage, "nextPage not equal")
	assert.Equal(t, false, userConnection.PageInfo.HasPreviousPage, "previousPage not equal")

	// BACKWARD PAGINATION
	userConnection, err = resolver.Query().Users(ctx, fm.ConnectionPagination{
		Backward: &fm.ConnectionBackwardPagination{
			Last:   10,
			Before: nil,
		},
	}, nil, nil)
	handleErr(t, err)
	assert.Equal(t, 10, len(userConnection.Edges), "edges not equal")
	assert.Equal(t, "Adam10", firstNameFromUser(userConnection.Edges[0]), "edges not equal")
	assert.Equal(t, "Adam01", firstNameFromUser(userConnection.Edges[9]), "edges not equal")
	assert.Equal(t, true, userConnection.PageInfo.HasPreviousPage, "nextPage not equal")
	assert.Equal(t, false, userConnection.PageInfo.HasNextPage, "previousPage not equal")

	startCursor := userConnection.PageInfo.StartCursor

	userConnection, err = resolver.Query().Users(ctx, fm.ConnectionPagination{
		Backward: &fm.ConnectionBackwardPagination{
			Last:   10,
			Before: startCursor,
		},
	}, nil, nil)
	handleErr(t, err)
	assert.Equal(t, 10, len(userConnection.Edges), "edges not equal")
	assert.Equal(t, "Adam20", firstNameFromUser(userConnection.Edges[0]), "edges not equal")
	assert.Equal(t, "Adam11", firstNameFromUser(userConnection.Edges[9]), "edges not equal")
	assert.Equal(t, true, userConnection.PageInfo.HasPreviousPage, "nextPage not equal")
	assert.Equal(t, true, userConnection.PageInfo.HasNextPage, "previousPage not equal")

	userConnection, err = resolver.Query().Users(ctx, fm.ConnectionPagination{
		Backward: &fm.ConnectionBackwardPagination{
			Last:   100,
			Before: nil,
		},
	}, nil, nil)
	handleErr(t, err)
	assert.Equal(t, "Eve20", firstNameFromUser(userConnection.Edges[0]), "edges not equal")
	assert.Equal(t, "Adam01", firstNameFromUser(userConnection.Edges[39]), "edges not equal")
	assert.Equal(t, 40, len(userConnection.Edges), "edges not equal")
	assert.Equal(t, false, userConnection.PageInfo.HasNextPage, "nextPage not equal")
	assert.Equal(t, false, userConnection.PageInfo.HasPreviousPage, "previousPage not equal")

	// Adam
	for i := 0; i < 20; i++ {
		m := models.User{
			FirstName: fmt.Sprintf("Dirk%02d", i),
			LastName:  fmt.Sprintf("Dirk%02d", i),
			Email:     fmt.Sprintf("Dirk%02d@gmail.com", i),
		}
		err = m.Insert(ctx, db, boil.Infer())
	}

	// With sorting
	sort := []*fm.UserOrdering{
		{
			Sort:      fm.UserSortFirstName,
			Direction: fm.SortDirectionAsc,
		},
		{
			Sort:      fm.UserSortLastName,
			Direction: fm.SortDirectionAsc,
		},
	}
	userConnection, err = resolver.Query().Users(ctx, fm.ConnectionPagination{
		Backward: &fm.ConnectionBackwardPagination{
			Last:   20,
			Before: nil,
		},
	}, sort, nil)
	handleErr(t, err)
	startCursor = userConnection.PageInfo.StartCursor
	for _, edge := range userConnection.Edges {
		expected := "Eve"
		if !strings.HasPrefix(firstNameFromUser(edge), expected) {
			t.Errorf("backward should start with %v but is %v", expected, firstNameFromUser(edge))
		}
	}
	userConnection, err = resolver.Query().Users(ctx, fm.ConnectionPagination{
		Backward: &fm.ConnectionBackwardPagination{
			Last:   20,
			Before: startCursor,
		},
	}, sort, nil)
	handleErr(t, err)
	startCursor = userConnection.PageInfo.StartCursor
	for _, edge := range userConnection.Edges {
		expected := "Dirk"
		if !strings.HasPrefix(firstNameFromUser(edge), expected) {
			t.Errorf("backward should start with %v but is %v", expected, firstNameFromUser(edge))
		}
	}

	userConnection, err = resolver.Query().Users(ctx, fm.ConnectionPagination{
		Backward: &fm.ConnectionBackwardPagination{
			Last:   20,
			Before: startCursor,
		},
	}, sort, nil)
	handleErr(t, err)
	startCursor = userConnection.PageInfo.StartCursor
	for _, edge := range userConnection.Edges {
		expected := "Adam"
		if !strings.HasPrefix(firstNameFromUser(edge), expected) {
			t.Errorf("backward should start with %v but is %v", expected, firstNameFromUser(edge))
		}
	}

	// other sort

	sort = []*fm.UserOrdering{
		{
			Sort:      fm.UserSortFirstName,
			Direction: fm.SortDirectionDesc,
		},
		{
			Sort:      fm.UserSortLastName,
			Direction: fm.SortDirectionDesc,
		},
	}
	userConnection, err = resolver.Query().Users(ctx, fm.ConnectionPagination{
		Backward: &fm.ConnectionBackwardPagination{
			Last:   20,
			Before: nil,
		},
	}, sort, nil)
	handleErr(t, err)
	startCursor = userConnection.PageInfo.StartCursor
	for _, edge := range userConnection.Edges {
		expected := "Adam"
		if !strings.HasPrefix(firstNameFromUser(edge), expected) {
			t.Errorf("backward should start with %v but is %v", expected, firstNameFromUser(edge))
		}
	}
	userConnection, err = resolver.Query().Users(ctx, fm.ConnectionPagination{
		Backward: &fm.ConnectionBackwardPagination{
			Last:   20,
			Before: startCursor,
		},
	}, sort, nil)
	handleErr(t, err)
	startCursor = userConnection.PageInfo.StartCursor
	for _, edge := range userConnection.Edges {
		expected := "Dirk"
		if !strings.HasPrefix(firstNameFromUser(edge), expected) {
			t.Errorf("backward should start with %v but is %v", expected, firstNameFromUser(edge))
		}
	}

	userConnection, err = resolver.Query().Users(ctx, fm.ConnectionPagination{
		Backward: &fm.ConnectionBackwardPagination{
			Last:   20,
			Before: startCursor,
		},
	}, sort, nil)
	handleErr(t, err)
	startCursor = userConnection.PageInfo.StartCursor
	for _, edge := range userConnection.Edges {
		expected := "Eve"
		if !strings.HasPrefix(firstNameFromUser(edge), expected) {
			t.Errorf("backward should start with %v but is %v", expected, firstNameFromUser(edge))
		}
	}
}

func firstNameFromUser(e *fm.UserEdge) string {
	return e.Node.FirstName
}

func handleErr(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
