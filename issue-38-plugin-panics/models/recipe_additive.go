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

// RecipeAdditive is an object representing the database table.
type RecipeAdditive struct {
	ID         int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt  time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt  time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	DeletedAt  null.Time `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	Percentage float64   `boil:"percentage" json:"percentage" toml:"percentage" yaml:"percentage"`
	AdditiveID int       `boil:"additive_id" json:"additive_id" toml:"additive_id" yaml:"additive_id"`
	RecipeID   int       `boil:"recipe_id" json:"recipe_id" toml:"recipe_id" yaml:"recipe_id"`

	R *recipeAdditiveR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L recipeAdditiveL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RecipeAdditiveColumns = struct {
	ID         string
	CreatedAt  string
	UpdatedAt  string
	DeletedAt  string
	Percentage string
	AdditiveID string
	RecipeID   string
}{
	ID:         "id",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
	Percentage: "percentage",
	AdditiveID: "additive_id",
	RecipeID:   "recipe_id",
}

// Generated where

var RecipeAdditiveWhere = struct {
	ID         whereHelperint
	CreatedAt  whereHelpertime_Time
	UpdatedAt  whereHelpertime_Time
	DeletedAt  whereHelpernull_Time
	Percentage whereHelperfloat64
	AdditiveID whereHelperint
	RecipeID   whereHelperint
}{
	ID:         whereHelperint{field: "\"recipe_additive\".\"id\""},
	CreatedAt:  whereHelpertime_Time{field: "\"recipe_additive\".\"created_at\""},
	UpdatedAt:  whereHelpertime_Time{field: "\"recipe_additive\".\"updated_at\""},
	DeletedAt:  whereHelpernull_Time{field: "\"recipe_additive\".\"deleted_at\""},
	Percentage: whereHelperfloat64{field: "\"recipe_additive\".\"percentage\""},
	AdditiveID: whereHelperint{field: "\"recipe_additive\".\"additive_id\""},
	RecipeID:   whereHelperint{field: "\"recipe_additive\".\"recipe_id\""},
}

// RecipeAdditiveRels is where relationship names are stored.
var RecipeAdditiveRels = struct {
	Additive string
	Recipe   string
}{
	Additive: "Additive",
	Recipe:   "Recipe",
}

// recipeAdditiveR is where relationships are stored.
type recipeAdditiveR struct {
	Additive *Additive `boil:"Additive" json:"Additive" toml:"Additive" yaml:"Additive"`
	Recipe   *Recipe   `boil:"Recipe" json:"Recipe" toml:"Recipe" yaml:"Recipe"`
}

// NewStruct creates a new relationship struct
func (*recipeAdditiveR) NewStruct() *recipeAdditiveR {
	return &recipeAdditiveR{}
}

// recipeAdditiveL is where Load methods for each relationship are stored.
type recipeAdditiveL struct{}

var (
	recipeAdditiveAllColumns            = []string{"id", "created_at", "updated_at", "deleted_at", "percentage", "additive_id", "recipe_id"}
	recipeAdditiveColumnsWithoutDefault = []string{"created_at", "updated_at", "deleted_at", "percentage", "additive_id", "recipe_id"}
	recipeAdditiveColumnsWithDefault    = []string{"id"}
	recipeAdditivePrimaryKeyColumns     = []string{"id"}
)

type (
	// RecipeAdditiveSlice is an alias for a slice of pointers to RecipeAdditive.
	// This should generally be used opposed to []RecipeAdditive.
	RecipeAdditiveSlice []*RecipeAdditive
	// RecipeAdditiveHook is the signature for custom RecipeAdditive hook methods
	RecipeAdditiveHook func(context.Context, boil.ContextExecutor, *RecipeAdditive) error

	recipeAdditiveQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	recipeAdditiveType                 = reflect.TypeOf(&RecipeAdditive{})
	recipeAdditiveMapping              = queries.MakeStructMapping(recipeAdditiveType)
	recipeAdditivePrimaryKeyMapping, _ = queries.BindMapping(recipeAdditiveType, recipeAdditiveMapping, recipeAdditivePrimaryKeyColumns)
	recipeAdditiveInsertCacheMut       sync.RWMutex
	recipeAdditiveInsertCache          = make(map[string]insertCache)
	recipeAdditiveUpdateCacheMut       sync.RWMutex
	recipeAdditiveUpdateCache          = make(map[string]updateCache)
	recipeAdditiveUpsertCacheMut       sync.RWMutex
	recipeAdditiveUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var recipeAdditiveBeforeInsertHooks []RecipeAdditiveHook
var recipeAdditiveBeforeUpdateHooks []RecipeAdditiveHook
var recipeAdditiveBeforeDeleteHooks []RecipeAdditiveHook
var recipeAdditiveBeforeUpsertHooks []RecipeAdditiveHook

var recipeAdditiveAfterInsertHooks []RecipeAdditiveHook
var recipeAdditiveAfterSelectHooks []RecipeAdditiveHook
var recipeAdditiveAfterUpdateHooks []RecipeAdditiveHook
var recipeAdditiveAfterDeleteHooks []RecipeAdditiveHook
var recipeAdditiveAfterUpsertHooks []RecipeAdditiveHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *RecipeAdditive) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeAdditiveBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *RecipeAdditive) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeAdditiveBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *RecipeAdditive) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeAdditiveBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *RecipeAdditive) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeAdditiveBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *RecipeAdditive) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeAdditiveAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *RecipeAdditive) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeAdditiveAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *RecipeAdditive) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeAdditiveAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *RecipeAdditive) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeAdditiveAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *RecipeAdditive) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeAdditiveAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRecipeAdditiveHook registers your hook function for all future operations.
func AddRecipeAdditiveHook(hookPoint boil.HookPoint, recipeAdditiveHook RecipeAdditiveHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		recipeAdditiveBeforeInsertHooks = append(recipeAdditiveBeforeInsertHooks, recipeAdditiveHook)
	case boil.BeforeUpdateHook:
		recipeAdditiveBeforeUpdateHooks = append(recipeAdditiveBeforeUpdateHooks, recipeAdditiveHook)
	case boil.BeforeDeleteHook:
		recipeAdditiveBeforeDeleteHooks = append(recipeAdditiveBeforeDeleteHooks, recipeAdditiveHook)
	case boil.BeforeUpsertHook:
		recipeAdditiveBeforeUpsertHooks = append(recipeAdditiveBeforeUpsertHooks, recipeAdditiveHook)
	case boil.AfterInsertHook:
		recipeAdditiveAfterInsertHooks = append(recipeAdditiveAfterInsertHooks, recipeAdditiveHook)
	case boil.AfterSelectHook:
		recipeAdditiveAfterSelectHooks = append(recipeAdditiveAfterSelectHooks, recipeAdditiveHook)
	case boil.AfterUpdateHook:
		recipeAdditiveAfterUpdateHooks = append(recipeAdditiveAfterUpdateHooks, recipeAdditiveHook)
	case boil.AfterDeleteHook:
		recipeAdditiveAfterDeleteHooks = append(recipeAdditiveAfterDeleteHooks, recipeAdditiveHook)
	case boil.AfterUpsertHook:
		recipeAdditiveAfterUpsertHooks = append(recipeAdditiveAfterUpsertHooks, recipeAdditiveHook)
	}
}

// One returns a single recipeAdditive record from the query.
func (q recipeAdditiveQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RecipeAdditive, error) {
	o := &RecipeAdditive{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for recipe_additive")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all RecipeAdditive records from the query.
func (q recipeAdditiveQuery) All(ctx context.Context, exec boil.ContextExecutor) (RecipeAdditiveSlice, error) {
	var o []*RecipeAdditive

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to RecipeAdditive slice")
	}

	if len(recipeAdditiveAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all RecipeAdditive records in the query.
func (q recipeAdditiveQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count recipe_additive rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q recipeAdditiveQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if recipe_additive exists")
	}

	return count > 0, nil
}

// Additive pointed to by the foreign key.
func (o *RecipeAdditive) Additive(mods ...qm.QueryMod) additiveQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.AdditiveID),
	}

	queryMods = append(queryMods, mods...)

	query := Additives(queryMods...)
	queries.SetFrom(query.Query, "\"additive\"")

	return query
}

// Recipe pointed to by the foreign key.
func (o *RecipeAdditive) Recipe(mods ...qm.QueryMod) recipeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.RecipeID),
	}

	queryMods = append(queryMods, mods...)

	query := Recipes(queryMods...)
	queries.SetFrom(query.Query, "\"recipe\"")

	return query
}

// LoadAdditive allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (recipeAdditiveL) LoadAdditive(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRecipeAdditive interface{}, mods queries.Applicator) error {
	var slice []*RecipeAdditive
	var object *RecipeAdditive

	if singular {
		object = maybeRecipeAdditive.(*RecipeAdditive)
	} else {
		slice = *maybeRecipeAdditive.(*[]*RecipeAdditive)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &recipeAdditiveR{}
		}
		args = append(args, object.AdditiveID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &recipeAdditiveR{}
			}

			for _, a := range args {
				if a == obj.AdditiveID {
					continue Outer
				}
			}

			args = append(args, obj.AdditiveID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`additive`),
		qm.WhereIn(`additive.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Additive")
	}

	var resultSlice []*Additive
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Additive")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for additive")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for additive")
	}

	if len(recipeAdditiveAfterSelectHooks) != 0 {
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
		object.R.Additive = foreign
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.AdditiveID == foreign.ID {
				local.R.Additive = foreign
				break
			}
		}
	}

	return nil
}

// LoadRecipe allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (recipeAdditiveL) LoadRecipe(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRecipeAdditive interface{}, mods queries.Applicator) error {
	var slice []*RecipeAdditive
	var object *RecipeAdditive

	if singular {
		object = maybeRecipeAdditive.(*RecipeAdditive)
	} else {
		slice = *maybeRecipeAdditive.(*[]*RecipeAdditive)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &recipeAdditiveR{}
		}
		args = append(args, object.RecipeID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &recipeAdditiveR{}
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

	if len(recipeAdditiveAfterSelectHooks) != 0 {
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

// SetAdditive of the recipeAdditive to the related item.
// Sets o.R.Additive to related.
// Adds o to related.R.RecipeAdditive.
func (o *RecipeAdditive) SetAdditive(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Additive) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"recipe_additive\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"additive_id"}),
		strmangle.WhereClause("\"", "\"", 2, recipeAdditivePrimaryKeyColumns),
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

	o.AdditiveID = related.ID
	if o.R == nil {
		o.R = &recipeAdditiveR{
			Additive: related,
		}
	} else {
		o.R.Additive = related
	}

	if related.R == nil {
		related.R = &additiveR{
			RecipeAdditive: o,
		}
	} else {
		related.R.RecipeAdditive = o
	}

	return nil
}

// SetRecipe of the recipeAdditive to the related item.
// Sets o.R.Recipe to related.
// Adds o to related.R.RecipeAdditives.
func (o *RecipeAdditive) SetRecipe(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Recipe) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"recipe_additive\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"recipe_id"}),
		strmangle.WhereClause("\"", "\"", 2, recipeAdditivePrimaryKeyColumns),
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
		o.R = &recipeAdditiveR{
			Recipe: related,
		}
	} else {
		o.R.Recipe = related
	}

	if related.R == nil {
		related.R = &recipeR{
			RecipeAdditives: RecipeAdditiveSlice{o},
		}
	} else {
		related.R.RecipeAdditives = append(related.R.RecipeAdditives, o)
	}

	return nil
}

// RecipeAdditives retrieves all the records using an executor.
func RecipeAdditives(mods ...qm.QueryMod) recipeAdditiveQuery {
	mods = append(mods, qm.From("\"recipe_additive\""))
	return recipeAdditiveQuery{NewQuery(mods...)}
}

// FindRecipeAdditive retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRecipeAdditive(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*RecipeAdditive, error) {
	recipeAdditiveObj := &RecipeAdditive{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"recipe_additive\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, recipeAdditiveObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from recipe_additive")
	}

	return recipeAdditiveObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RecipeAdditive) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no recipe_additive provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(recipeAdditiveColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	recipeAdditiveInsertCacheMut.RLock()
	cache, cached := recipeAdditiveInsertCache[key]
	recipeAdditiveInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			recipeAdditiveAllColumns,
			recipeAdditiveColumnsWithDefault,
			recipeAdditiveColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(recipeAdditiveType, recipeAdditiveMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(recipeAdditiveType, recipeAdditiveMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"recipe_additive\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"recipe_additive\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into recipe_additive")
	}

	if !cached {
		recipeAdditiveInsertCacheMut.Lock()
		recipeAdditiveInsertCache[key] = cache
		recipeAdditiveInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the RecipeAdditive.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *RecipeAdditive) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	recipeAdditiveUpdateCacheMut.RLock()
	cache, cached := recipeAdditiveUpdateCache[key]
	recipeAdditiveUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			recipeAdditiveAllColumns,
			recipeAdditivePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update recipe_additive, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"recipe_additive\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, recipeAdditivePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(recipeAdditiveType, recipeAdditiveMapping, append(wl, recipeAdditivePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update recipe_additive row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for recipe_additive")
	}

	if !cached {
		recipeAdditiveUpdateCacheMut.Lock()
		recipeAdditiveUpdateCache[key] = cache
		recipeAdditiveUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q recipeAdditiveQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for recipe_additive")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for recipe_additive")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RecipeAdditiveSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), recipeAdditivePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"recipe_additive\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, recipeAdditivePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in recipeAdditive slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all recipeAdditive")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RecipeAdditive) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no recipe_additive provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(recipeAdditiveColumnsWithDefault, o)

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

	recipeAdditiveUpsertCacheMut.RLock()
	cache, cached := recipeAdditiveUpsertCache[key]
	recipeAdditiveUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			recipeAdditiveAllColumns,
			recipeAdditiveColumnsWithDefault,
			recipeAdditiveColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			recipeAdditiveAllColumns,
			recipeAdditivePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert recipe_additive, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(recipeAdditivePrimaryKeyColumns))
			copy(conflict, recipeAdditivePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"recipe_additive\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(recipeAdditiveType, recipeAdditiveMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(recipeAdditiveType, recipeAdditiveMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert recipe_additive")
	}

	if !cached {
		recipeAdditiveUpsertCacheMut.Lock()
		recipeAdditiveUpsertCache[key] = cache
		recipeAdditiveUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single RecipeAdditive record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *RecipeAdditive) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RecipeAdditive provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), recipeAdditivePrimaryKeyMapping)
	sql := "DELETE FROM \"recipe_additive\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from recipe_additive")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for recipe_additive")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q recipeAdditiveQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no recipeAdditiveQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from recipe_additive")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for recipe_additive")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RecipeAdditiveSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(recipeAdditiveBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), recipeAdditivePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"recipe_additive\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, recipeAdditivePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from recipeAdditive slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for recipe_additive")
	}

	if len(recipeAdditiveAfterDeleteHooks) != 0 {
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
func (o *RecipeAdditive) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRecipeAdditive(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RecipeAdditiveSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RecipeAdditiveSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), recipeAdditivePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"recipe_additive\".* FROM \"recipe_additive\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, recipeAdditivePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RecipeAdditiveSlice")
	}

	*o = slice

	return nil
}

// RecipeAdditiveExists checks if the RecipeAdditive row exists.
func RecipeAdditiveExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"recipe_additive\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if recipe_additive exists")
	}

	return exists, nil
}
