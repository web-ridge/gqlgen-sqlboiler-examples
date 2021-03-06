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

// RecipeBatchFragrance is an object representing the database table.
type RecipeBatchFragrance struct {
	ID          int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt   time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt   time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	DeletedAt   null.Time `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	Weight      float64   `boil:"weight" json:"weight" toml:"weight" yaml:"weight"`
	Cost        float64   `boil:"cost" json:"cost" toml:"cost" yaml:"cost"`
	FragranceID int       `boil:"fragrance_id" json:"fragrance_id" toml:"fragrance_id" yaml:"fragrance_id"`
	BatchID     int       `boil:"batch_id" json:"batch_id" toml:"batch_id" yaml:"batch_id"`

	R *recipeBatchFragranceR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L recipeBatchFragranceL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RecipeBatchFragranceColumns = struct {
	ID          string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
	Weight      string
	Cost        string
	FragranceID string
	BatchID     string
}{
	ID:          "id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
	Weight:      "weight",
	Cost:        "cost",
	FragranceID: "fragrance_id",
	BatchID:     "batch_id",
}

// Generated where

var RecipeBatchFragranceWhere = struct {
	ID          whereHelperint
	CreatedAt   whereHelpertime_Time
	UpdatedAt   whereHelpertime_Time
	DeletedAt   whereHelpernull_Time
	Weight      whereHelperfloat64
	Cost        whereHelperfloat64
	FragranceID whereHelperint
	BatchID     whereHelperint
}{
	ID:          whereHelperint{field: "\"recipe_batch_fragrance\".\"id\""},
	CreatedAt:   whereHelpertime_Time{field: "\"recipe_batch_fragrance\".\"created_at\""},
	UpdatedAt:   whereHelpertime_Time{field: "\"recipe_batch_fragrance\".\"updated_at\""},
	DeletedAt:   whereHelpernull_Time{field: "\"recipe_batch_fragrance\".\"deleted_at\""},
	Weight:      whereHelperfloat64{field: "\"recipe_batch_fragrance\".\"weight\""},
	Cost:        whereHelperfloat64{field: "\"recipe_batch_fragrance\".\"cost\""},
	FragranceID: whereHelperint{field: "\"recipe_batch_fragrance\".\"fragrance_id\""},
	BatchID:     whereHelperint{field: "\"recipe_batch_fragrance\".\"batch_id\""},
}

// RecipeBatchFragranceRels is where relationship names are stored.
var RecipeBatchFragranceRels = struct {
	Batch     string
	Fragrance string
}{
	Batch:     "Batch",
	Fragrance: "Fragrance",
}

// recipeBatchFragranceR is where relationships are stored.
type recipeBatchFragranceR struct {
	Batch     *RecipeBatch `boil:"Batch" json:"Batch" toml:"Batch" yaml:"Batch"`
	Fragrance *Fragrance   `boil:"Fragrance" json:"Fragrance" toml:"Fragrance" yaml:"Fragrance"`
}

// NewStruct creates a new relationship struct
func (*recipeBatchFragranceR) NewStruct() *recipeBatchFragranceR {
	return &recipeBatchFragranceR{}
}

// recipeBatchFragranceL is where Load methods for each relationship are stored.
type recipeBatchFragranceL struct{}

var (
	recipeBatchFragranceAllColumns            = []string{"id", "created_at", "updated_at", "deleted_at", "weight", "cost", "fragrance_id", "batch_id"}
	recipeBatchFragranceColumnsWithoutDefault = []string{"created_at", "updated_at", "deleted_at", "weight", "cost", "fragrance_id", "batch_id"}
	recipeBatchFragranceColumnsWithDefault    = []string{"id"}
	recipeBatchFragrancePrimaryKeyColumns     = []string{"id"}
)

type (
	// RecipeBatchFragranceSlice is an alias for a slice of pointers to RecipeBatchFragrance.
	// This should generally be used opposed to []RecipeBatchFragrance.
	RecipeBatchFragranceSlice []*RecipeBatchFragrance
	// RecipeBatchFragranceHook is the signature for custom RecipeBatchFragrance hook methods
	RecipeBatchFragranceHook func(context.Context, boil.ContextExecutor, *RecipeBatchFragrance) error

	recipeBatchFragranceQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	recipeBatchFragranceType                 = reflect.TypeOf(&RecipeBatchFragrance{})
	recipeBatchFragranceMapping              = queries.MakeStructMapping(recipeBatchFragranceType)
	recipeBatchFragrancePrimaryKeyMapping, _ = queries.BindMapping(recipeBatchFragranceType, recipeBatchFragranceMapping, recipeBatchFragrancePrimaryKeyColumns)
	recipeBatchFragranceInsertCacheMut       sync.RWMutex
	recipeBatchFragranceInsertCache          = make(map[string]insertCache)
	recipeBatchFragranceUpdateCacheMut       sync.RWMutex
	recipeBatchFragranceUpdateCache          = make(map[string]updateCache)
	recipeBatchFragranceUpsertCacheMut       sync.RWMutex
	recipeBatchFragranceUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var recipeBatchFragranceBeforeInsertHooks []RecipeBatchFragranceHook
var recipeBatchFragranceBeforeUpdateHooks []RecipeBatchFragranceHook
var recipeBatchFragranceBeforeDeleteHooks []RecipeBatchFragranceHook
var recipeBatchFragranceBeforeUpsertHooks []RecipeBatchFragranceHook

var recipeBatchFragranceAfterInsertHooks []RecipeBatchFragranceHook
var recipeBatchFragranceAfterSelectHooks []RecipeBatchFragranceHook
var recipeBatchFragranceAfterUpdateHooks []RecipeBatchFragranceHook
var recipeBatchFragranceAfterDeleteHooks []RecipeBatchFragranceHook
var recipeBatchFragranceAfterUpsertHooks []RecipeBatchFragranceHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *RecipeBatchFragrance) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchFragranceBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *RecipeBatchFragrance) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchFragranceBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *RecipeBatchFragrance) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchFragranceBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *RecipeBatchFragrance) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchFragranceBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *RecipeBatchFragrance) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchFragranceAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *RecipeBatchFragrance) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchFragranceAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *RecipeBatchFragrance) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchFragranceAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *RecipeBatchFragrance) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchFragranceAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *RecipeBatchFragrance) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchFragranceAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRecipeBatchFragranceHook registers your hook function for all future operations.
func AddRecipeBatchFragranceHook(hookPoint boil.HookPoint, recipeBatchFragranceHook RecipeBatchFragranceHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		recipeBatchFragranceBeforeInsertHooks = append(recipeBatchFragranceBeforeInsertHooks, recipeBatchFragranceHook)
	case boil.BeforeUpdateHook:
		recipeBatchFragranceBeforeUpdateHooks = append(recipeBatchFragranceBeforeUpdateHooks, recipeBatchFragranceHook)
	case boil.BeforeDeleteHook:
		recipeBatchFragranceBeforeDeleteHooks = append(recipeBatchFragranceBeforeDeleteHooks, recipeBatchFragranceHook)
	case boil.BeforeUpsertHook:
		recipeBatchFragranceBeforeUpsertHooks = append(recipeBatchFragranceBeforeUpsertHooks, recipeBatchFragranceHook)
	case boil.AfterInsertHook:
		recipeBatchFragranceAfterInsertHooks = append(recipeBatchFragranceAfterInsertHooks, recipeBatchFragranceHook)
	case boil.AfterSelectHook:
		recipeBatchFragranceAfterSelectHooks = append(recipeBatchFragranceAfterSelectHooks, recipeBatchFragranceHook)
	case boil.AfterUpdateHook:
		recipeBatchFragranceAfterUpdateHooks = append(recipeBatchFragranceAfterUpdateHooks, recipeBatchFragranceHook)
	case boil.AfterDeleteHook:
		recipeBatchFragranceAfterDeleteHooks = append(recipeBatchFragranceAfterDeleteHooks, recipeBatchFragranceHook)
	case boil.AfterUpsertHook:
		recipeBatchFragranceAfterUpsertHooks = append(recipeBatchFragranceAfterUpsertHooks, recipeBatchFragranceHook)
	}
}

// One returns a single recipeBatchFragrance record from the query.
func (q recipeBatchFragranceQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RecipeBatchFragrance, error) {
	o := &RecipeBatchFragrance{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for recipe_batch_fragrance")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all RecipeBatchFragrance records from the query.
func (q recipeBatchFragranceQuery) All(ctx context.Context, exec boil.ContextExecutor) (RecipeBatchFragranceSlice, error) {
	var o []*RecipeBatchFragrance

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to RecipeBatchFragrance slice")
	}

	if len(recipeBatchFragranceAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all RecipeBatchFragrance records in the query.
func (q recipeBatchFragranceQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count recipe_batch_fragrance rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q recipeBatchFragranceQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if recipe_batch_fragrance exists")
	}

	return count > 0, nil
}

// Batch pointed to by the foreign key.
func (o *RecipeBatchFragrance) Batch(mods ...qm.QueryMod) recipeBatchQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.BatchID),
	}

	queryMods = append(queryMods, mods...)

	query := RecipeBatches(queryMods...)
	queries.SetFrom(query.Query, "\"recipe_batch\"")

	return query
}

// Fragrance pointed to by the foreign key.
func (o *RecipeBatchFragrance) Fragrance(mods ...qm.QueryMod) fragranceQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.FragranceID),
	}

	queryMods = append(queryMods, mods...)

	query := Fragrances(queryMods...)
	queries.SetFrom(query.Query, "\"fragrance\"")

	return query
}

// LoadBatch allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (recipeBatchFragranceL) LoadBatch(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRecipeBatchFragrance interface{}, mods queries.Applicator) error {
	var slice []*RecipeBatchFragrance
	var object *RecipeBatchFragrance

	if singular {
		object = maybeRecipeBatchFragrance.(*RecipeBatchFragrance)
	} else {
		slice = *maybeRecipeBatchFragrance.(*[]*RecipeBatchFragrance)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &recipeBatchFragranceR{}
		}
		args = append(args, object.BatchID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &recipeBatchFragranceR{}
			}

			for _, a := range args {
				if a == obj.BatchID {
					continue Outer
				}
			}

			args = append(args, obj.BatchID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`recipe_batch`),
		qm.WhereIn(`recipe_batch.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load RecipeBatch")
	}

	var resultSlice []*RecipeBatch
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice RecipeBatch")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for recipe_batch")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for recipe_batch")
	}

	if len(recipeBatchFragranceAfterSelectHooks) != 0 {
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
		object.R.Batch = foreign
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.BatchID == foreign.ID {
				local.R.Batch = foreign
				break
			}
		}
	}

	return nil
}

// LoadFragrance allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (recipeBatchFragranceL) LoadFragrance(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRecipeBatchFragrance interface{}, mods queries.Applicator) error {
	var slice []*RecipeBatchFragrance
	var object *RecipeBatchFragrance

	if singular {
		object = maybeRecipeBatchFragrance.(*RecipeBatchFragrance)
	} else {
		slice = *maybeRecipeBatchFragrance.(*[]*RecipeBatchFragrance)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &recipeBatchFragranceR{}
		}
		args = append(args, object.FragranceID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &recipeBatchFragranceR{}
			}

			for _, a := range args {
				if a == obj.FragranceID {
					continue Outer
				}
			}

			args = append(args, obj.FragranceID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`fragrance`),
		qm.WhereIn(`fragrance.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Fragrance")
	}

	var resultSlice []*Fragrance
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Fragrance")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for fragrance")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for fragrance")
	}

	if len(recipeBatchFragranceAfterSelectHooks) != 0 {
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
		object.R.Fragrance = foreign
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.FragranceID == foreign.ID {
				local.R.Fragrance = foreign
				break
			}
		}
	}

	return nil
}

// SetBatch of the recipeBatchFragrance to the related item.
// Sets o.R.Batch to related.
// Adds o to related.R.BatchRecipeBatchFragrances.
func (o *RecipeBatchFragrance) SetBatch(ctx context.Context, exec boil.ContextExecutor, insert bool, related *RecipeBatch) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"recipe_batch_fragrance\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"batch_id"}),
		strmangle.WhereClause("\"", "\"", 2, recipeBatchFragrancePrimaryKeyColumns),
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

	o.BatchID = related.ID
	if o.R == nil {
		o.R = &recipeBatchFragranceR{
			Batch: related,
		}
	} else {
		o.R.Batch = related
	}

	if related.R == nil {
		related.R = &recipeBatchR{
			BatchRecipeBatchFragrances: RecipeBatchFragranceSlice{o},
		}
	} else {
		related.R.BatchRecipeBatchFragrances = append(related.R.BatchRecipeBatchFragrances, o)
	}

	return nil
}

// SetFragrance of the recipeBatchFragrance to the related item.
// Sets o.R.Fragrance to related.
// Adds o to related.R.RecipeBatchFragrance.
func (o *RecipeBatchFragrance) SetFragrance(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Fragrance) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"recipe_batch_fragrance\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"fragrance_id"}),
		strmangle.WhereClause("\"", "\"", 2, recipeBatchFragrancePrimaryKeyColumns),
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

	o.FragranceID = related.ID
	if o.R == nil {
		o.R = &recipeBatchFragranceR{
			Fragrance: related,
		}
	} else {
		o.R.Fragrance = related
	}

	if related.R == nil {
		related.R = &fragranceR{
			RecipeBatchFragrance: o,
		}
	} else {
		related.R.RecipeBatchFragrance = o
	}

	return nil
}

// RecipeBatchFragrances retrieves all the records using an executor.
func RecipeBatchFragrances(mods ...qm.QueryMod) recipeBatchFragranceQuery {
	mods = append(mods, qm.From("\"recipe_batch_fragrance\""))
	return recipeBatchFragranceQuery{NewQuery(mods...)}
}

// FindRecipeBatchFragrance retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRecipeBatchFragrance(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*RecipeBatchFragrance, error) {
	recipeBatchFragranceObj := &RecipeBatchFragrance{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"recipe_batch_fragrance\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, recipeBatchFragranceObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from recipe_batch_fragrance")
	}

	return recipeBatchFragranceObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RecipeBatchFragrance) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no recipe_batch_fragrance provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(recipeBatchFragranceColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	recipeBatchFragranceInsertCacheMut.RLock()
	cache, cached := recipeBatchFragranceInsertCache[key]
	recipeBatchFragranceInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			recipeBatchFragranceAllColumns,
			recipeBatchFragranceColumnsWithDefault,
			recipeBatchFragranceColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(recipeBatchFragranceType, recipeBatchFragranceMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(recipeBatchFragranceType, recipeBatchFragranceMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"recipe_batch_fragrance\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"recipe_batch_fragrance\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into recipe_batch_fragrance")
	}

	if !cached {
		recipeBatchFragranceInsertCacheMut.Lock()
		recipeBatchFragranceInsertCache[key] = cache
		recipeBatchFragranceInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the RecipeBatchFragrance.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *RecipeBatchFragrance) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	recipeBatchFragranceUpdateCacheMut.RLock()
	cache, cached := recipeBatchFragranceUpdateCache[key]
	recipeBatchFragranceUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			recipeBatchFragranceAllColumns,
			recipeBatchFragrancePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update recipe_batch_fragrance, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"recipe_batch_fragrance\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, recipeBatchFragrancePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(recipeBatchFragranceType, recipeBatchFragranceMapping, append(wl, recipeBatchFragrancePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update recipe_batch_fragrance row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for recipe_batch_fragrance")
	}

	if !cached {
		recipeBatchFragranceUpdateCacheMut.Lock()
		recipeBatchFragranceUpdateCache[key] = cache
		recipeBatchFragranceUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q recipeBatchFragranceQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for recipe_batch_fragrance")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for recipe_batch_fragrance")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RecipeBatchFragranceSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), recipeBatchFragrancePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"recipe_batch_fragrance\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, recipeBatchFragrancePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in recipeBatchFragrance slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all recipeBatchFragrance")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RecipeBatchFragrance) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no recipe_batch_fragrance provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(recipeBatchFragranceColumnsWithDefault, o)

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

	recipeBatchFragranceUpsertCacheMut.RLock()
	cache, cached := recipeBatchFragranceUpsertCache[key]
	recipeBatchFragranceUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			recipeBatchFragranceAllColumns,
			recipeBatchFragranceColumnsWithDefault,
			recipeBatchFragranceColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			recipeBatchFragranceAllColumns,
			recipeBatchFragrancePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert recipe_batch_fragrance, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(recipeBatchFragrancePrimaryKeyColumns))
			copy(conflict, recipeBatchFragrancePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"recipe_batch_fragrance\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(recipeBatchFragranceType, recipeBatchFragranceMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(recipeBatchFragranceType, recipeBatchFragranceMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert recipe_batch_fragrance")
	}

	if !cached {
		recipeBatchFragranceUpsertCacheMut.Lock()
		recipeBatchFragranceUpsertCache[key] = cache
		recipeBatchFragranceUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single RecipeBatchFragrance record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *RecipeBatchFragrance) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RecipeBatchFragrance provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), recipeBatchFragrancePrimaryKeyMapping)
	sql := "DELETE FROM \"recipe_batch_fragrance\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from recipe_batch_fragrance")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for recipe_batch_fragrance")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q recipeBatchFragranceQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no recipeBatchFragranceQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from recipe_batch_fragrance")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for recipe_batch_fragrance")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RecipeBatchFragranceSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(recipeBatchFragranceBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), recipeBatchFragrancePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"recipe_batch_fragrance\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, recipeBatchFragrancePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from recipeBatchFragrance slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for recipe_batch_fragrance")
	}

	if len(recipeBatchFragranceAfterDeleteHooks) != 0 {
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
func (o *RecipeBatchFragrance) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRecipeBatchFragrance(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RecipeBatchFragranceSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RecipeBatchFragranceSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), recipeBatchFragrancePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"recipe_batch_fragrance\".* FROM \"recipe_batch_fragrance\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, recipeBatchFragrancePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RecipeBatchFragranceSlice")
	}

	*o = slice

	return nil
}

// RecipeBatchFragranceExists checks if the RecipeBatchFragrance row exists.
func RecipeBatchFragranceExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"recipe_batch_fragrance\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if recipe_batch_fragrance exists")
	}

	return exists, nil
}
