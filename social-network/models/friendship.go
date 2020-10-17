// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Friendship is an object representing the database table.
type Friendship struct {
	ID        uint      `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`

	R *friendshipR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L friendshipL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var FriendshipColumns = struct {
	ID        string
	CreatedAt string
}{
	ID:        "id",
	CreatedAt: "created_at",
}

// Generated where

var FriendshipWhere = struct {
	ID        whereHelperuint
	CreatedAt whereHelpernull_Time
}{
	ID:        whereHelperuint{field: "`friendship`.`id`"},
	CreatedAt: whereHelpernull_Time{field: "`friendship`.`created_at`"},
}

// FriendshipRels is where relationship names are stored.
var FriendshipRels = struct {
	Users string
}{
	Users: "Users",
}

// friendshipR is where relationships are stored.
type friendshipR struct {
	Users UserSlice `boil:"Users" json:"Users" toml:"Users" yaml:"Users"`
}

// NewStruct creates a new relationship struct
func (*friendshipR) NewStruct() *friendshipR {
	return &friendshipR{}
}

// friendshipL is where Load methods for each relationship are stored.
type friendshipL struct{}

var (
	friendshipAllColumns            = []string{"id", "created_at"}
	friendshipColumnsWithoutDefault = []string{"created_at"}
	friendshipColumnsWithDefault    = []string{"id"}
	friendshipPrimaryKeyColumns     = []string{"id"}
)

type (
	// FriendshipSlice is an alias for a slice of pointers to Friendship.
	// This should generally be used opposed to []Friendship.
	FriendshipSlice []*Friendship
	// FriendshipHook is the signature for custom Friendship hook methods
	FriendshipHook func(context.Context, boil.ContextExecutor, *Friendship) error

	friendshipQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	friendshipType                 = reflect.TypeOf(&Friendship{})
	friendshipMapping              = queries.MakeStructMapping(friendshipType)
	friendshipPrimaryKeyMapping, _ = queries.BindMapping(friendshipType, friendshipMapping, friendshipPrimaryKeyColumns)
	friendshipInsertCacheMut       sync.RWMutex
	friendshipInsertCache          = make(map[string]insertCache)
	friendshipUpdateCacheMut       sync.RWMutex
	friendshipUpdateCache          = make(map[string]updateCache)
	friendshipUpsertCacheMut       sync.RWMutex
	friendshipUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var friendshipBeforeInsertHooks []FriendshipHook
var friendshipBeforeUpdateHooks []FriendshipHook
var friendshipBeforeDeleteHooks []FriendshipHook
var friendshipBeforeUpsertHooks []FriendshipHook

var friendshipAfterInsertHooks []FriendshipHook
var friendshipAfterSelectHooks []FriendshipHook
var friendshipAfterUpdateHooks []FriendshipHook
var friendshipAfterDeleteHooks []FriendshipHook
var friendshipAfterUpsertHooks []FriendshipHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Friendship) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range friendshipBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Friendship) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range friendshipBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Friendship) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range friendshipBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Friendship) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range friendshipBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Friendship) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range friendshipAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Friendship) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range friendshipAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Friendship) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range friendshipAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Friendship) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range friendshipAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Friendship) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range friendshipAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFriendshipHook registers your hook function for all future operations.
func AddFriendshipHook(hookPoint boil.HookPoint, friendshipHook FriendshipHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		friendshipBeforeInsertHooks = append(friendshipBeforeInsertHooks, friendshipHook)
	case boil.BeforeUpdateHook:
		friendshipBeforeUpdateHooks = append(friendshipBeforeUpdateHooks, friendshipHook)
	case boil.BeforeDeleteHook:
		friendshipBeforeDeleteHooks = append(friendshipBeforeDeleteHooks, friendshipHook)
	case boil.BeforeUpsertHook:
		friendshipBeforeUpsertHooks = append(friendshipBeforeUpsertHooks, friendshipHook)
	case boil.AfterInsertHook:
		friendshipAfterInsertHooks = append(friendshipAfterInsertHooks, friendshipHook)
	case boil.AfterSelectHook:
		friendshipAfterSelectHooks = append(friendshipAfterSelectHooks, friendshipHook)
	case boil.AfterUpdateHook:
		friendshipAfterUpdateHooks = append(friendshipAfterUpdateHooks, friendshipHook)
	case boil.AfterDeleteHook:
		friendshipAfterDeleteHooks = append(friendshipAfterDeleteHooks, friendshipHook)
	case boil.AfterUpsertHook:
		friendshipAfterUpsertHooks = append(friendshipAfterUpsertHooks, friendshipHook)
	}
}

// One returns a single friendship record from the query.
func (q friendshipQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Friendship, error) {
	o := &Friendship{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for friendship")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Friendship records from the query.
func (q friendshipQuery) All(ctx context.Context, exec boil.ContextExecutor) (FriendshipSlice, error) {
	var o []*Friendship

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Friendship slice")
	}

	if len(friendshipAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Friendship records in the query.
func (q friendshipQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count friendship rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q friendshipQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if friendship exists")
	}

	return count > 0, nil
}

// Users retrieves all the user's Users with an executor.
func (o *Friendship) Users(mods ...qm.QueryMod) userQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.InnerJoin("`user_has_friendship` on `user`.`id` = `user_has_friendship`.`user_id`"),
		qm.Where("`user_has_friendship`.`friendship_id`=?", o.ID),
	)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "`user`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`user`.*"})
	}

	return query
}

// LoadUsers allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (friendshipL) LoadUsers(ctx context.Context, e boil.ContextExecutor, singular bool, maybeFriendship interface{}, mods queries.Applicator) error {
	var slice []*Friendship
	var object *Friendship

	if singular {
		object = maybeFriendship.(*Friendship)
	} else {
		slice = *maybeFriendship.(*[]*Friendship)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &friendshipR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &friendshipR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.Select("`user`.*, `a`.`friendship_id`"),
		qm.From("`user`"),
		qm.InnerJoin("`user_has_friendship` as `a` on `user`.`id` = `a`.`user_id`"),
		qm.WhereIn("`a`.`friendship_id` in ?", args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load user")
	}

	var resultSlice []*User

	var localJoinCols []uint
	for results.Next() {
		one := new(User)
		var localJoinCol uint

		err = results.Scan(&one.ID, &one.FirstName, &one.LastName, &one.Email, &one.Password, &localJoinCol)
		if err != nil {
			return errors.Wrap(err, "failed to scan eager loaded results for user")
		}
		if err = results.Err(); err != nil {
			return errors.Wrap(err, "failed to plebian-bind eager loaded slice user")
		}

		resultSlice = append(resultSlice, one)
		localJoinCols = append(localJoinCols, localJoinCol)
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on user")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user")
	}

	if len(userAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Users = resultSlice
		return nil
	}

	for i, foreign := range resultSlice {
		localJoinCol := localJoinCols[i]
		for _, local := range slice {
			if local.ID == localJoinCol {
				local.R.Users = append(local.R.Users, foreign)
				break
			}
		}
	}

	return nil
}

// AddUsers adds the given related objects to the existing relationships
// of the friendship, optionally inserting them as new records.
// Appends related to o.R.Users.
// Sets related.R.Friendships appropriately.
func (o *Friendship) AddUsers(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*User) error {
	var err error
	for _, rel := range related {
		if insert {
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		}
	}

	for _, rel := range related {
		query := "insert into `user_has_friendship` (`friendship_id`, `user_id`) values (?, ?)"
		values := []interface{}{o.ID, rel.ID}

		if boil.IsDebug(ctx) {
			writer := boil.DebugWriterFrom(ctx)
			fmt.Fprintln(writer, query)
			fmt.Fprintln(writer, values)
		}
		_, err = exec.ExecContext(ctx, query, values...)
		if err != nil {
			return errors.Wrap(err, "failed to insert into join table")
		}
	}
	if o.R == nil {
		o.R = &friendshipR{
			Users: related,
		}
	} else {
		o.R.Users = append(o.R.Users, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &userR{
				Friendships: FriendshipSlice{o},
			}
		} else {
			rel.R.Friendships = append(rel.R.Friendships, o)
		}
	}
	return nil
}

// SetUsers removes all previously related items of the
// friendship replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Friendships's Users accordingly.
// Replaces o.R.Users with related.
// Sets related.R.Friendships's Users accordingly.
func (o *Friendship) SetUsers(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*User) error {
	query := "delete from `user_has_friendship` where `friendship_id` = ?"
	values := []interface{}{o.ID}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	removeUsersFromFriendshipsSlice(o, related)
	if o.R != nil {
		o.R.Users = nil
	}
	return o.AddUsers(ctx, exec, insert, related...)
}

// RemoveUsers relationships from objects passed in.
// Removes related items from R.Users (uses pointer comparison, removal does not keep order)
// Sets related.R.Friendships.
func (o *Friendship) RemoveUsers(ctx context.Context, exec boil.ContextExecutor, related ...*User) error {
	var err error
	query := fmt.Sprintf(
		"delete from `user_has_friendship` where `friendship_id` = ? and `user_id` in (%s)",
		strmangle.Placeholders(dialect.UseIndexPlaceholders, len(related), 2, 1),
	)
	values := []interface{}{o.ID}
	for _, rel := range related {
		values = append(values, rel.ID)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err = exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}
	removeUsersFromFriendshipsSlice(o, related)
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Users {
			if rel != ri {
				continue
			}

			ln := len(o.R.Users)
			if ln > 1 && i < ln-1 {
				o.R.Users[i] = o.R.Users[ln-1]
			}
			o.R.Users = o.R.Users[:ln-1]
			break
		}
	}

	return nil
}

func removeUsersFromFriendshipsSlice(o *Friendship, related []*User) {
	for _, rel := range related {
		if rel.R == nil {
			continue
		}
		for i, ri := range rel.R.Friendships {
			if o.ID != ri.ID {
				continue
			}

			ln := len(rel.R.Friendships)
			if ln > 1 && i < ln-1 {
				rel.R.Friendships[i] = rel.R.Friendships[ln-1]
			}
			rel.R.Friendships = rel.R.Friendships[:ln-1]
			break
		}
	}
}

// Friendships retrieves all the records using an executor.
func Friendships(mods ...qm.QueryMod) friendshipQuery {
	mods = append(mods, qm.From("`friendship`"))
	return friendshipQuery{NewQuery(mods...)}
}

// FindFriendship retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFriendship(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*Friendship, error) {
	friendshipObj := &Friendship{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `friendship` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, friendshipObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from friendship")
	}

	return friendshipObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Friendship) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no friendship provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(friendshipColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	friendshipInsertCacheMut.RLock()
	cache, cached := friendshipInsertCache[key]
	friendshipInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			friendshipAllColumns,
			friendshipColumnsWithDefault,
			friendshipColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(friendshipType, friendshipMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(friendshipType, friendshipMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `friendship` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `friendship` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `friendship` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, friendshipPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into friendship")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == friendshipMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for friendship")
	}

CacheNoHooks:
	if !cached {
		friendshipInsertCacheMut.Lock()
		friendshipInsertCache[key] = cache
		friendshipInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Friendship.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Friendship) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	friendshipUpdateCacheMut.RLock()
	cache, cached := friendshipUpdateCache[key]
	friendshipUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			friendshipAllColumns,
			friendshipPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update friendship, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `friendship` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, friendshipPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(friendshipType, friendshipMapping, append(wl, friendshipPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update friendship row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for friendship")
	}

	if !cached {
		friendshipUpdateCacheMut.Lock()
		friendshipUpdateCache[key] = cache
		friendshipUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q friendshipQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for friendship")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for friendship")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FriendshipSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), friendshipPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `friendship` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, friendshipPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in friendship slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all friendship")
	}
	return rowsAff, nil
}

var mySQLFriendshipUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Friendship) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no friendship provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(friendshipColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLFriendshipUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	friendshipUpsertCacheMut.RLock()
	cache, cached := friendshipUpsertCache[key]
	friendshipUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			friendshipAllColumns,
			friendshipColumnsWithDefault,
			friendshipColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			friendshipAllColumns,
			friendshipPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert friendship, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "friendship", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `friendship` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(friendshipType, friendshipMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(friendshipType, friendshipMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for friendship")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == friendshipMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(friendshipType, friendshipMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for friendship")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for friendship")
	}

CacheNoHooks:
	if !cached {
		friendshipUpsertCacheMut.Lock()
		friendshipUpsertCache[key] = cache
		friendshipUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Friendship record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Friendship) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Friendship provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), friendshipPrimaryKeyMapping)
	sql := "DELETE FROM `friendship` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from friendship")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for friendship")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q friendshipQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no friendshipQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from friendship")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for friendship")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FriendshipSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(friendshipBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), friendshipPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `friendship` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, friendshipPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from friendship slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for friendship")
	}

	if len(friendshipAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Friendship) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindFriendship(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FriendshipSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := FriendshipSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), friendshipPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `friendship`.* FROM `friendship` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, friendshipPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in FriendshipSlice")
	}

	*o = slice

	return nil
}

// FriendshipExists checks if the Friendship row exists.
func FriendshipExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `friendship` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if friendship exists")
	}

	return exists, nil
}
