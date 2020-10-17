// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testAuthUsers(t *testing.T) {
	t.Parallel()

	query := AuthUsers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testAuthUsersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthUser{}
	if err = randomize.Struct(seed, o, authUserDBTypes, true, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AuthUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthUsersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthUser{}
	if err = randomize.Struct(seed, o, authUserDBTypes, true, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := AuthUsers().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AuthUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthUsersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthUser{}
	if err = randomize.Struct(seed, o, authUserDBTypes, true, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AuthUserSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AuthUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthUsersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthUser{}
	if err = randomize.Struct(seed, o, authUserDBTypes, true, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := AuthUserExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if AuthUser exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AuthUserExists to return true, but got false.")
	}
}

func testAuthUsersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthUser{}
	if err = randomize.Struct(seed, o, authUserDBTypes, true, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	authUserFound, err := FindAuthUser(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if authUserFound == nil {
		t.Error("want a record, got nil")
	}
}

func testAuthUsersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthUser{}
	if err = randomize.Struct(seed, o, authUserDBTypes, true, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = AuthUsers().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testAuthUsersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthUser{}
	if err = randomize.Struct(seed, o, authUserDBTypes, true, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := AuthUsers().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAuthUsersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserOne := &AuthUser{}
	authUserTwo := &AuthUser{}
	if err = randomize.Struct(seed, authUserOne, authUserDBTypes, false, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}
	if err = randomize.Struct(seed, authUserTwo, authUserDBTypes, false, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = authUserOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = authUserTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AuthUsers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAuthUsersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	authUserOne := &AuthUser{}
	authUserTwo := &AuthUser{}
	if err = randomize.Struct(seed, authUserOne, authUserDBTypes, false, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}
	if err = randomize.Struct(seed, authUserTwo, authUserDBTypes, false, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = authUserOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = authUserTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AuthUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func authUserBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *AuthUser) error {
	*o = AuthUser{}
	return nil
}

func authUserAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *AuthUser) error {
	*o = AuthUser{}
	return nil
}

func authUserAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *AuthUser) error {
	*o = AuthUser{}
	return nil
}

func authUserBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *AuthUser) error {
	*o = AuthUser{}
	return nil
}

func authUserAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *AuthUser) error {
	*o = AuthUser{}
	return nil
}

func authUserBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *AuthUser) error {
	*o = AuthUser{}
	return nil
}

func authUserAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *AuthUser) error {
	*o = AuthUser{}
	return nil
}

func authUserBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *AuthUser) error {
	*o = AuthUser{}
	return nil
}

func authUserAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *AuthUser) error {
	*o = AuthUser{}
	return nil
}

func testAuthUsersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &AuthUser{}
	o := &AuthUser{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, authUserDBTypes, false); err != nil {
		t.Errorf("Unable to randomize AuthUser object: %s", err)
	}

	AddAuthUserHook(boil.BeforeInsertHook, authUserBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	authUserBeforeInsertHooks = []AuthUserHook{}

	AddAuthUserHook(boil.AfterInsertHook, authUserAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	authUserAfterInsertHooks = []AuthUserHook{}

	AddAuthUserHook(boil.AfterSelectHook, authUserAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	authUserAfterSelectHooks = []AuthUserHook{}

	AddAuthUserHook(boil.BeforeUpdateHook, authUserBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	authUserBeforeUpdateHooks = []AuthUserHook{}

	AddAuthUserHook(boil.AfterUpdateHook, authUserAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	authUserAfterUpdateHooks = []AuthUserHook{}

	AddAuthUserHook(boil.BeforeDeleteHook, authUserBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	authUserBeforeDeleteHooks = []AuthUserHook{}

	AddAuthUserHook(boil.AfterDeleteHook, authUserAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	authUserAfterDeleteHooks = []AuthUserHook{}

	AddAuthUserHook(boil.BeforeUpsertHook, authUserBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	authUserBeforeUpsertHooks = []AuthUserHook{}

	AddAuthUserHook(boil.AfterUpsertHook, authUserAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	authUserAfterUpsertHooks = []AuthUserHook{}
}

func testAuthUsersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthUser{}
	if err = randomize.Struct(seed, o, authUserDBTypes, true, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AuthUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAuthUsersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthUser{}
	if err = randomize.Struct(seed, o, authUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(authUserColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := AuthUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAuthUserToManyUserAuthUserGroups(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AuthUser
	var b, c AuthUserGroup

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authUserDBTypes, true, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, authUserGroupDBTypes, false, authUserGroupColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, authUserGroupDBTypes, false, authUserGroupColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.UserID = a.ID
	c.UserID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.UserAuthUserGroups().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.UserID == b.UserID {
			bFound = true
		}
		if v.UserID == c.UserID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := AuthUserSlice{&a}
	if err = a.L.LoadUserAuthUserGroups(ctx, tx, false, (*[]*AuthUser)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserAuthUserGroups); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.UserAuthUserGroups = nil
	if err = a.L.LoadUserAuthUserGroups(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserAuthUserGroups); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testAuthUserToManyUserAuthUserUserPermissions(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AuthUser
	var b, c AuthUserUserPermission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authUserDBTypes, true, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, authUserUserPermissionDBTypes, false, authUserUserPermissionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, authUserUserPermissionDBTypes, false, authUserUserPermissionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.UserID = a.ID
	c.UserID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.UserAuthUserUserPermissions().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.UserID == b.UserID {
			bFound = true
		}
		if v.UserID == c.UserID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := AuthUserSlice{&a}
	if err = a.L.LoadUserAuthUserUserPermissions(ctx, tx, false, (*[]*AuthUser)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserAuthUserUserPermissions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.UserAuthUserUserPermissions = nil
	if err = a.L.LoadUserAuthUserUserPermissions(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserAuthUserUserPermissions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testAuthUserToManyAddOpUserAuthUserGroups(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AuthUser
	var b, c, d, e AuthUserGroup

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authUserDBTypes, false, strmangle.SetComplement(authUserPrimaryKeyColumns, authUserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*AuthUserGroup{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, authUserGroupDBTypes, false, strmangle.SetComplement(authUserGroupPrimaryKeyColumns, authUserGroupColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*AuthUserGroup{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUserAuthUserGroups(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.UserID {
			t.Error("foreign key was wrong value", a.ID, first.UserID)
		}
		if a.ID != second.UserID {
			t.Error("foreign key was wrong value", a.ID, second.UserID)
		}

		if first.R.User != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.User != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.UserAuthUserGroups[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.UserAuthUserGroups[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.UserAuthUserGroups().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testAuthUserToManyAddOpUserAuthUserUserPermissions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AuthUser
	var b, c, d, e AuthUserUserPermission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authUserDBTypes, false, strmangle.SetComplement(authUserPrimaryKeyColumns, authUserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*AuthUserUserPermission{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, authUserUserPermissionDBTypes, false, strmangle.SetComplement(authUserUserPermissionPrimaryKeyColumns, authUserUserPermissionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*AuthUserUserPermission{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUserAuthUserUserPermissions(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.UserID {
			t.Error("foreign key was wrong value", a.ID, first.UserID)
		}
		if a.ID != second.UserID {
			t.Error("foreign key was wrong value", a.ID, second.UserID)
		}

		if first.R.User != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.User != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.UserAuthUserUserPermissions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.UserAuthUserUserPermissions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.UserAuthUserUserPermissions().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testAuthUsersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthUser{}
	if err = randomize.Struct(seed, o, authUserDBTypes, true, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testAuthUsersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthUser{}
	if err = randomize.Struct(seed, o, authUserDBTypes, true, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AuthUserSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testAuthUsersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AuthUser{}
	if err = randomize.Struct(seed, o, authUserDBTypes, true, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AuthUsers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	authUserDBTypes = map[string]string{`ID`: `integer`, `Password`: `character varying`, `LastLogin`: `timestamp with time zone`, `IsSuperuser`: `boolean`, `Username`: `character varying`, `FirstName`: `character varying`, `LastName`: `character varying`, `Email`: `character varying`, `IsStaff`: `boolean`, `IsActive`: `boolean`, `DateJoined`: `timestamp with time zone`}
	_               = bytes.MinRead
)

func testAuthUsersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(authUserPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(authUserAllColumns) == len(authUserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AuthUser{}
	if err = randomize.Struct(seed, o, authUserDBTypes, true, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AuthUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, authUserDBTypes, true, authUserPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testAuthUsersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(authUserAllColumns) == len(authUserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AuthUser{}
	if err = randomize.Struct(seed, o, authUserDBTypes, true, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AuthUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, authUserDBTypes, true, authUserPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(authUserAllColumns, authUserPrimaryKeyColumns) {
		fields = authUserAllColumns
	} else {
		fields = strmangle.SetComplement(
			authUserAllColumns,
			authUserPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := AuthUserSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testAuthUsersUpsert(t *testing.T) {
	t.Parallel()

	if len(authUserAllColumns) == len(authUserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := AuthUser{}
	if err = randomize.Struct(seed, &o, authUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert AuthUser: %s", err)
	}

	count, err := AuthUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, authUserDBTypes, false, authUserPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert AuthUser: %s", err)
	}

	count, err = AuthUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
