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

// RecipeLipid is an object representing the database table.
type RecipeLipid struct {
	ID         int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt  time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt  time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	DeletedAt  null.Time `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	Percentage float64   `boil:"percentage" json:"percentage" toml:"percentage" yaml:"percentage"`
	LipidID    int       `boil:"lipid_id" json:"lipid_id" toml:"lipid_id" yaml:"lipid_id"`
	RecipeID   int       `boil:"recipe_id" json:"recipe_id" toml:"recipe_id" yaml:"recipe_id"`

	R *recipeLipidR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L recipeLipidL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RecipeLipidColumns = struct {
	ID         string
	CreatedAt  string
	UpdatedAt  string
	DeletedAt  string
	Percentage string
	LipidID    string
	RecipeID   string
}{
	ID:         "id",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
	Percentage: "percentage",
	LipidID:    "lipid_id",
	RecipeID:   "recipe_id",
}

// Generated where

var RecipeLipidWhere = struct {
	ID         whereHelperint
	CreatedAt  whereHelpertime_Time
	UpdatedAt  whereHelpertime_Time
	DeletedAt  whereHelpernull_Time
	Percentage whereHelperfloat64
	LipidID    whereHelperint
	RecipeID   whereHelperint
}{
	ID:         whereHelperint{field: "\"recipe_lipid\".\"id\""},
	CreatedAt:  whereHelpertime_Time{field: "\"recipe_lipid\".\"created_at\""},
	UpdatedAt:  whereHelpertime_Time{field: "\"recipe_lipid\".\"updated_at\""},
	DeletedAt:  whereHelpernull_Time{field: "\"recipe_lipid\".\"deleted_at\""},
	Percentage: whereHelperfloat64{field: "\"recipe_lipid\".\"percentage\""},
	LipidID:    whereHelperint{field: "\"recipe_lipid\".\"lipid_id\""},
	RecipeID:   whereHelperint{field: "\"recipe_lipid\".\"recipe_id\""},
}

// RecipeLipidRels is where relationship names are stored.
var RecipeLipidRels = struct {
	Lipid  string
	Recipe string
}{
	Lipid:  "Lipid",
	Recipe: "Recipe",
}

// recipeLipidR is where relationships are stored.
type recipeLipidR struct {
	Lipid  *Lipid  `boil:"Lipid" json:"Lipid" toml:"Lipid" yaml:"Lipid"`
	Recipe *Recipe `boil:"Recipe" json:"Recipe" toml:"Recipe" yaml:"Recipe"`
}

// NewStruct creates a new relationship struct
func (*recipeLipidR) NewStruct() *recipeLipidR {
	return &recipeLipidR{}
}

// recipeLipidL is where Load methods for each relationship are stored.
type recipeLipidL struct{}

var (
	recipeLipidAllColumns            = []string{"id", "created_at", "updated_at", "deleted_at", "percentage", "lipid_id", "recipe_id"}
	recipeLipidColumnsWithoutDefault = []string{"created_at", "updated_at", "deleted_at", "percentage", "lipid_id", "recipe_id"}
	recipeLipidColumnsWithDefault    = []string{"id"}
	recipeLipidPrimaryKeyColumns     = []string{"id"}
)

type (
	// RecipeLipidSlice is an alias for a slice of pointers to RecipeLipid.
	// This should generally be used opposed to []RecipeLipid.
	RecipeLipidSlice []*RecipeLipid
	// RecipeLipidHook is the signature for custom RecipeLipid hook methods
	RecipeLipidHook func(context.Context, boil.ContextExecutor, *RecipeLipid) error

	recipeLipidQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	recipeLipidType                 = reflect.TypeOf(&RecipeLipid{})
	recipeLipidMapping              = queries.MakeStructMapping(recipeLipidType)
	recipeLipidPrimaryKeyMapping, _ = queries.BindMapping(recipeLipidType, recipeLipidMapping, recipeLipidPrimaryKeyColumns)
	recipeLipidInsertCacheMut       sync.RWMutex
	recipeLipidInsertCache          = make(map[string]insertCache)
	recipeLipidUpdateCacheMut       sync.RWMutex
	recipeLipidUpdateCache          = make(map[string]updateCache)
	recipeLipidUpsertCacheMut       sync.RWMutex
	recipeLipidUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var recipeLipidBeforeInsertHooks []RecipeLipidHook
var recipeLipidBeforeUpdateHooks []RecipeLipidHook
var recipeLipidBeforeDeleteHooks []RecipeLipidHook
var recipeLipidBeforeUpsertHooks []RecipeLipidHook

var recipeLipidAfterInsertHooks []RecipeLipidHook
var recipeLipidAfterSelectHooks []RecipeLipidHook
var recipeLipidAfterUpdateHooks []RecipeLipidHook
var recipeLipidAfterDeleteHooks []RecipeLipidHook
var recipeLipidAfterUpsertHooks []RecipeLipidHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *RecipeLipid) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeLipidBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *RecipeLipid) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeLipidBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *RecipeLipid) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeLipidBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *RecipeLipid) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeLipidBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *RecipeLipid) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeLipidAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *RecipeLipid) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeLipidAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *RecipeLipid) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeLipidAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *RecipeLipid) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeLipidAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *RecipeLipid) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeLipidAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRecipeLipidHook registers your hook function for all future operations.
func AddRecipeLipidHook(hookPoint boil.HookPoint, recipeLipidHook RecipeLipidHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		recipeLipidBeforeInsertHooks = append(recipeLipidBeforeInsertHooks, recipeLipidHook)
	case boil.BeforeUpdateHook:
		recipeLipidBeforeUpdateHooks = append(recipeLipidBeforeUpdateHooks, recipeLipidHook)
	case boil.BeforeDeleteHook:
		recipeLipidBeforeDeleteHooks = append(recipeLipidBeforeDeleteHooks, recipeLipidHook)
	case boil.BeforeUpsertHook:
		recipeLipidBeforeUpsertHooks = append(recipeLipidBeforeUpsertHooks, recipeLipidHook)
	case boil.AfterInsertHook:
		recipeLipidAfterInsertHooks = append(recipeLipidAfterInsertHooks, recipeLipidHook)
	case boil.AfterSelectHook:
		recipeLipidAfterSelectHooks = append(recipeLipidAfterSelectHooks, recipeLipidHook)
	case boil.AfterUpdateHook:
		recipeLipidAfterUpdateHooks = append(recipeLipidAfterUpdateHooks, recipeLipidHook)
	case boil.AfterDeleteHook:
		recipeLipidAfterDeleteHooks = append(recipeLipidAfterDeleteHooks, recipeLipidHook)
	case boil.AfterUpsertHook:
		recipeLipidAfterUpsertHooks = append(recipeLipidAfterUpsertHooks, recipeLipidHook)
	}
}

// One returns a single recipeLipid record from the query.
func (q recipeLipidQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RecipeLipid, error) {
	o := &RecipeLipid{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for recipe_lipid")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all RecipeLipid records from the query.
func (q recipeLipidQuery) All(ctx context.Context, exec boil.ContextExecutor) (RecipeLipidSlice, error) {
	var o []*RecipeLipid

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to RecipeLipid slice")
	}

	if len(recipeLipidAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all RecipeLipid records in the query.
func (q recipeLipidQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count recipe_lipid rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q recipeLipidQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if recipe_lipid exists")
	}

	return count > 0, nil
}

// Lipid pointed to by the foreign key.
func (o *RecipeLipid) Lipid(mods ...qm.QueryMod) lipidQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.LipidID),
	}

	queryMods = append(queryMods, mods...)

	query := Lipids(queryMods...)
	queries.SetFrom(query.Query, "\"lipid\"")

	return query
}

// Recipe pointed to by the foreign key.
func (o *RecipeLipid) Recipe(mods ...qm.QueryMod) recipeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.RecipeID),
	}

	queryMods = append(queryMods, mods...)

	query := Recipes(queryMods...)
	queries.SetFrom(query.Query, "\"recipe\"")

	return query
}

// LoadLipid allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (recipeLipidL) LoadLipid(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRecipeLipid interface{}, mods queries.Applicator) error {
	var slice []*RecipeLipid
	var object *RecipeLipid

	if singular {
		object = maybeRecipeLipid.(*RecipeLipid)
	} else {
		slice = *maybeRecipeLipid.(*[]*RecipeLipid)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &recipeLipidR{}
		}
		args = append(args, object.LipidID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &recipeLipidR{}
			}

			for _, a := range args {
				if a == obj.LipidID {
					continue Outer
				}
			}

			args = append(args, obj.LipidID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`lipid`),
		qm.WhereIn(`lipid.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Lipid")
	}

	var resultSlice []*Lipid
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Lipid")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for lipid")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for lipid")
	}

	if len(recipeLipidAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Lipid = foreign
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.LipidID == foreign.ID {
				local.R.Lipid = foreign
				break
			}
		}
	}

	return nil
}

// LoadRecipe allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (recipeLipidL) LoadRecipe(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRecipeLipid interface{}, mods queries.Applicator) error {
	var slice []*RecipeLipid
	var object *RecipeLipid

	if singular {
		object = maybeRecipeLipid.(*RecipeLipid)
	} else {
		slice = *maybeRecipeLipid.(*[]*RecipeLipid)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &recipeLipidR{}
		}
		args = append(args, object.RecipeID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &recipeLipidR{}
			}

			for _, a := range args {
				if a == obj.RecipeID {
					continue Outer
				}
			}

			args = append(args, obj.RecipeID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`recipe`),
		qm.WhereIn(`recipe.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Recipe")
	}

	var resultSlice []*Recipe
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Recipe")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for recipe")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for recipe")
	}

	if len(recipeLipidAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Recipe = foreign
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.RecipeID == foreign.ID {
				local.R.Recipe = foreign
				break
			}
		}
	}

	return nil
}

// SetLipid of the recipeLipid to the related item.
// Sets o.R.Lipid to related.
// Adds o to related.R.RecipeLipid.
func (o *RecipeLipid) SetLipid(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Lipid) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"recipe_lipid\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"lipid_id"}),
		strmangle.WhereClause("\"", "\"", 2, recipeLipidPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.LipidID = related.ID
	if o.R == nil {
		o.R = &recipeLipidR{
			Lipid: related,
		}
	} else {
		o.R.Lipid = related
	}

	if related.R == nil {
		related.R = &lipidR{
			RecipeLipid: o,
		}
	} else {
		related.R.RecipeLipid = o
	}

	return nil
}

// SetRecipe of the recipeLipid to the related item.
// Sets o.R.Recipe to related.
// Adds o to related.R.RecipeLipids.
func (o *RecipeLipid) SetRecipe(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Recipe) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"recipe_lipid\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"recipe_id"}),
		strmangle.WhereClause("\"", "\"", 2, recipeLipidPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.RecipeID = related.ID
	if o.R == nil {
		o.R = &recipeLipidR{
			Recipe: related,
		}
	} else {
		o.R.Recipe = related
	}

	if related.R == nil {
		related.R = &recipeR{
			RecipeLipids: RecipeLipidSlice{o},
		}
	} else {
		related.R.RecipeLipids = append(related.R.RecipeLipids, o)
	}

	return nil
}

// RecipeLipids retrieves all the records using an executor.
func RecipeLipids(mods ...qm.QueryMod) recipeLipidQuery {
	mods = append(mods, qm.From("\"recipe_lipid\""))
	return recipeLipidQuery{NewQuery(mods...)}
}

// FindRecipeLipid retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRecipeLipid(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*RecipeLipid, error) {
	recipeLipidObj := &RecipeLipid{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"recipe_lipid\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, recipeLipidObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from recipe_lipid")
	}

	return recipeLipidObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RecipeLipid) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no recipe_lipid provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(recipeLipidColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	recipeLipidInsertCacheMut.RLock()
	cache, cached := recipeLipidInsertCache[key]
	recipeLipidInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			recipeLipidAllColumns,
			recipeLipidColumnsWithDefault,
			recipeLipidColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(recipeLipidType, recipeLipidMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(recipeLipidType, recipeLipidMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"recipe_lipid\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"recipe_lipid\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
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

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into recipe_lipid")
	}

	if !cached {
		recipeLipidInsertCacheMut.Lock()
		recipeLipidInsertCache[key] = cache
		recipeLipidInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the RecipeLipid.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *RecipeLipid) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	recipeLipidUpdateCacheMut.RLock()
	cache, cached := recipeLipidUpdateCache[key]
	recipeLipidUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			recipeLipidAllColumns,
			recipeLipidPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update recipe_lipid, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"recipe_lipid\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, recipeLipidPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(recipeLipidType, recipeLipidMapping, append(wl, recipeLipidPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update recipe_lipid row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for recipe_lipid")
	}

	if !cached {
		recipeLipidUpdateCacheMut.Lock()
		recipeLipidUpdateCache[key] = cache
		recipeLipidUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q recipeLipidQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for recipe_lipid")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for recipe_lipid")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RecipeLipidSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), recipeLipidPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"recipe_lipid\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, recipeLipidPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in recipeLipid slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all recipeLipid")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RecipeLipid) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no recipe_lipid provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(recipeLipidColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
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
	key := buf.String()
	strmangle.PutBuffer(buf)

	recipeLipidUpsertCacheMut.RLock()
	cache, cached := recipeLipidUpsertCache[key]
	recipeLipidUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			recipeLipidAllColumns,
			recipeLipidColumnsWithDefault,
			recipeLipidColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			recipeLipidAllColumns,
			recipeLipidPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert recipe_lipid, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(recipeLipidPrimaryKeyColumns))
			copy(conflict, recipeLipidPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"recipe_lipid\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(recipeLipidType, recipeLipidMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(recipeLipidType, recipeLipidMapping, ret)
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
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert recipe_lipid")
	}

	if !cached {
		recipeLipidUpsertCacheMut.Lock()
		recipeLipidUpsertCache[key] = cache
		recipeLipidUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single RecipeLipid record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *RecipeLipid) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RecipeLipid provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), recipeLipidPrimaryKeyMapping)
	sql := "DELETE FROM \"recipe_lipid\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from recipe_lipid")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for recipe_lipid")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q recipeLipidQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no recipeLipidQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from recipe_lipid")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for recipe_lipid")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RecipeLipidSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(recipeLipidBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), recipeLipidPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"recipe_lipid\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, recipeLipidPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from recipeLipid slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for recipe_lipid")
	}

	if len(recipeLipidAfterDeleteHooks) != 0 {
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
func (o *RecipeLipid) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRecipeLipid(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RecipeLipidSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RecipeLipidSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), recipeLipidPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"recipe_lipid\".* FROM \"recipe_lipid\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, recipeLipidPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RecipeLipidSlice")
	}

	*o = slice

	return nil
}

// RecipeLipidExists checks if the RecipeLipid row exists.
func RecipeLipidExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"recipe_lipid\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if recipe_lipid exists")
	}

	return exists, nil
}
