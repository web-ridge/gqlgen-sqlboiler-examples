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

// RecipeBatchLipid is an object representing the database table.
type RecipeBatchLipid struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	DeletedAt null.Time `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	Weight    float64   `boil:"weight" json:"weight" toml:"weight" yaml:"weight"`
	Cost      float64   `boil:"cost" json:"cost" toml:"cost" yaml:"cost"`
	LipidID   int       `boil:"lipid_id" json:"lipid_id" toml:"lipid_id" yaml:"lipid_id"`
	BatchID   int       `boil:"batch_id" json:"batch_id" toml:"batch_id" yaml:"batch_id"`

	R *recipeBatchLipidR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L recipeBatchLipidL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RecipeBatchLipidColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	Weight    string
	Cost      string
	LipidID   string
	BatchID   string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Weight:    "weight",
	Cost:      "cost",
	LipidID:   "lipid_id",
	BatchID:   "batch_id",
}

// Generated where

var RecipeBatchLipidWhere = struct {
	ID        whereHelperint
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
	DeletedAt whereHelpernull_Time
	Weight    whereHelperfloat64
	Cost      whereHelperfloat64
	LipidID   whereHelperint
	BatchID   whereHelperint
}{
	ID:        whereHelperint{field: "\"recipe_batch_lipid\".\"id\""},
	CreatedAt: whereHelpertime_Time{field: "\"recipe_batch_lipid\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"recipe_batch_lipid\".\"updated_at\""},
	DeletedAt: whereHelpernull_Time{field: "\"recipe_batch_lipid\".\"deleted_at\""},
	Weight:    whereHelperfloat64{field: "\"recipe_batch_lipid\".\"weight\""},
	Cost:      whereHelperfloat64{field: "\"recipe_batch_lipid\".\"cost\""},
	LipidID:   whereHelperint{field: "\"recipe_batch_lipid\".\"lipid_id\""},
	BatchID:   whereHelperint{field: "\"recipe_batch_lipid\".\"batch_id\""},
}

// RecipeBatchLipidRels is where relationship names are stored.
var RecipeBatchLipidRels = struct {
	Batch string
	Lipid string
}{
	Batch: "Batch",
	Lipid: "Lipid",
}

// recipeBatchLipidR is where relationships are stored.
type recipeBatchLipidR struct {
	Batch *RecipeBatch `boil:"Batch" json:"Batch" toml:"Batch" yaml:"Batch"`
	Lipid *Lipid       `boil:"Lipid" json:"Lipid" toml:"Lipid" yaml:"Lipid"`
}

// NewStruct creates a new relationship struct
func (*recipeBatchLipidR) NewStruct() *recipeBatchLipidR {
	return &recipeBatchLipidR{}
}

// recipeBatchLipidL is where Load methods for each relationship are stored.
type recipeBatchLipidL struct{}

var (
	recipeBatchLipidAllColumns            = []string{"id", "created_at", "updated_at", "deleted_at", "weight", "cost", "lipid_id", "batch_id"}
	recipeBatchLipidColumnsWithoutDefault = []string{"created_at", "updated_at", "deleted_at", "weight", "cost", "lipid_id", "batch_id"}
	recipeBatchLipidColumnsWithDefault    = []string{"id"}
	recipeBatchLipidPrimaryKeyColumns     = []string{"id"}
)

type (
	// RecipeBatchLipidSlice is an alias for a slice of pointers to RecipeBatchLipid.
	// This should generally be used opposed to []RecipeBatchLipid.
	RecipeBatchLipidSlice []*RecipeBatchLipid
	// RecipeBatchLipidHook is the signature for custom RecipeBatchLipid hook methods
	RecipeBatchLipidHook func(context.Context, boil.ContextExecutor, *RecipeBatchLipid) error

	recipeBatchLipidQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	recipeBatchLipidType                 = reflect.TypeOf(&RecipeBatchLipid{})
	recipeBatchLipidMapping              = queries.MakeStructMapping(recipeBatchLipidType)
	recipeBatchLipidPrimaryKeyMapping, _ = queries.BindMapping(recipeBatchLipidType, recipeBatchLipidMapping, recipeBatchLipidPrimaryKeyColumns)
	recipeBatchLipidInsertCacheMut       sync.RWMutex
	recipeBatchLipidInsertCache          = make(map[string]insertCache)
	recipeBatchLipidUpdateCacheMut       sync.RWMutex
	recipeBatchLipidUpdateCache          = make(map[string]updateCache)
	recipeBatchLipidUpsertCacheMut       sync.RWMutex
	recipeBatchLipidUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var recipeBatchLipidBeforeInsertHooks []RecipeBatchLipidHook
var recipeBatchLipidBeforeUpdateHooks []RecipeBatchLipidHook
var recipeBatchLipidBeforeDeleteHooks []RecipeBatchLipidHook
var recipeBatchLipidBeforeUpsertHooks []RecipeBatchLipidHook

var recipeBatchLipidAfterInsertHooks []RecipeBatchLipidHook
var recipeBatchLipidAfterSelectHooks []RecipeBatchLipidHook
var recipeBatchLipidAfterUpdateHooks []RecipeBatchLipidHook
var recipeBatchLipidAfterDeleteHooks []RecipeBatchLipidHook
var recipeBatchLipidAfterUpsertHooks []RecipeBatchLipidHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *RecipeBatchLipid) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchLipidBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *RecipeBatchLipid) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchLipidBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *RecipeBatchLipid) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchLipidBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *RecipeBatchLipid) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchLipidBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *RecipeBatchLipid) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchLipidAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *RecipeBatchLipid) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchLipidAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *RecipeBatchLipid) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchLipidAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *RecipeBatchLipid) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchLipidAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *RecipeBatchLipid) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchLipidAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRecipeBatchLipidHook registers your hook function for all future operations.
func AddRecipeBatchLipidHook(hookPoint boil.HookPoint, recipeBatchLipidHook RecipeBatchLipidHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		recipeBatchLipidBeforeInsertHooks = append(recipeBatchLipidBeforeInsertHooks, recipeBatchLipidHook)
	case boil.BeforeUpdateHook:
		recipeBatchLipidBeforeUpdateHooks = append(recipeBatchLipidBeforeUpdateHooks, recipeBatchLipidHook)
	case boil.BeforeDeleteHook:
		recipeBatchLipidBeforeDeleteHooks = append(recipeBatchLipidBeforeDeleteHooks, recipeBatchLipidHook)
	case boil.BeforeUpsertHook:
		recipeBatchLipidBeforeUpsertHooks = append(recipeBatchLipidBeforeUpsertHooks, recipeBatchLipidHook)
	case boil.AfterInsertHook:
		recipeBatchLipidAfterInsertHooks = append(recipeBatchLipidAfterInsertHooks, recipeBatchLipidHook)
	case boil.AfterSelectHook:
		recipeBatchLipidAfterSelectHooks = append(recipeBatchLipidAfterSelectHooks, recipeBatchLipidHook)
	case boil.AfterUpdateHook:
		recipeBatchLipidAfterUpdateHooks = append(recipeBatchLipidAfterUpdateHooks, recipeBatchLipidHook)
	case boil.AfterDeleteHook:
		recipeBatchLipidAfterDeleteHooks = append(recipeBatchLipidAfterDeleteHooks, recipeBatchLipidHook)
	case boil.AfterUpsertHook:
		recipeBatchLipidAfterUpsertHooks = append(recipeBatchLipidAfterUpsertHooks, recipeBatchLipidHook)
	}
}

// One returns a single recipeBatchLipid record from the query.
func (q recipeBatchLipidQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RecipeBatchLipid, error) {
	o := &RecipeBatchLipid{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for recipe_batch_lipid")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all RecipeBatchLipid records from the query.
func (q recipeBatchLipidQuery) All(ctx context.Context, exec boil.ContextExecutor) (RecipeBatchLipidSlice, error) {
	var o []*RecipeBatchLipid

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to RecipeBatchLipid slice")
	}

	if len(recipeBatchLipidAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all RecipeBatchLipid records in the query.
func (q recipeBatchLipidQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count recipe_batch_lipid rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q recipeBatchLipidQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if recipe_batch_lipid exists")
	}

	return count > 0, nil
}

// Batch pointed to by the foreign key.
func (o *RecipeBatchLipid) Batch(mods ...qm.QueryMod) recipeBatchQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.BatchID),
	}

	queryMods = append(queryMods, mods...)

	query := RecipeBatches(queryMods...)
	queries.SetFrom(query.Query, "\"recipe_batch\"")

	return query
}

// Lipid pointed to by the foreign key.
func (o *RecipeBatchLipid) Lipid(mods ...qm.QueryMod) lipidQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.LipidID),
	}

	queryMods = append(queryMods, mods...)

	query := Lipids(queryMods...)
	queries.SetFrom(query.Query, "\"lipid\"")

	return query
}

// LoadBatch allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (recipeBatchLipidL) LoadBatch(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRecipeBatchLipid interface{}, mods queries.Applicator) error {
	var slice []*RecipeBatchLipid
	var object *RecipeBatchLipid

	if singular {
		object = maybeRecipeBatchLipid.(*RecipeBatchLipid)
	} else {
		slice = *maybeRecipeBatchLipid.(*[]*RecipeBatchLipid)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &recipeBatchLipidR{}
		}
		args = append(args, object.BatchID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &recipeBatchLipidR{}
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

	if len(recipeBatchLipidAfterSelectHooks) != 0 {
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

// LoadLipid allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (recipeBatchLipidL) LoadLipid(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRecipeBatchLipid interface{}, mods queries.Applicator) error {
	var slice []*RecipeBatchLipid
	var object *RecipeBatchLipid

	if singular {
		object = maybeRecipeBatchLipid.(*RecipeBatchLipid)
	} else {
		slice = *maybeRecipeBatchLipid.(*[]*RecipeBatchLipid)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &recipeBatchLipidR{}
		}
		args = append(args, object.LipidID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &recipeBatchLipidR{}
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

	if len(recipeBatchLipidAfterSelectHooks) != 0 {
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

// SetBatch of the recipeBatchLipid to the related item.
// Sets o.R.Batch to related.
// Adds o to related.R.BatchRecipeBatchLipids.
func (o *RecipeBatchLipid) SetBatch(ctx context.Context, exec boil.ContextExecutor, insert bool, related *RecipeBatch) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"recipe_batch_lipid\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"batch_id"}),
		strmangle.WhereClause("\"", "\"", 2, recipeBatchLipidPrimaryKeyColumns),
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
		o.R = &recipeBatchLipidR{
			Batch: related,
		}
	} else {
		o.R.Batch = related
	}

	if related.R == nil {
		related.R = &recipeBatchR{
			BatchRecipeBatchLipids: RecipeBatchLipidSlice{o},
		}
	} else {
		related.R.BatchRecipeBatchLipids = append(related.R.BatchRecipeBatchLipids, o)
	}

	return nil
}

// SetLipid of the recipeBatchLipid to the related item.
// Sets o.R.Lipid to related.
// Adds o to related.R.RecipeBatchLipid.
func (o *RecipeBatchLipid) SetLipid(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Lipid) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"recipe_batch_lipid\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"lipid_id"}),
		strmangle.WhereClause("\"", "\"", 2, recipeBatchLipidPrimaryKeyColumns),
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
		o.R = &recipeBatchLipidR{
			Lipid: related,
		}
	} else {
		o.R.Lipid = related
	}

	if related.R == nil {
		related.R = &lipidR{
			RecipeBatchLipid: o,
		}
	} else {
		related.R.RecipeBatchLipid = o
	}

	return nil
}

// RecipeBatchLipids retrieves all the records using an executor.
func RecipeBatchLipids(mods ...qm.QueryMod) recipeBatchLipidQuery {
	mods = append(mods, qm.From("\"recipe_batch_lipid\""))
	return recipeBatchLipidQuery{NewQuery(mods...)}
}

// FindRecipeBatchLipid retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRecipeBatchLipid(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*RecipeBatchLipid, error) {
	recipeBatchLipidObj := &RecipeBatchLipid{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"recipe_batch_lipid\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, recipeBatchLipidObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from recipe_batch_lipid")
	}

	return recipeBatchLipidObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RecipeBatchLipid) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no recipe_batch_lipid provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(recipeBatchLipidColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	recipeBatchLipidInsertCacheMut.RLock()
	cache, cached := recipeBatchLipidInsertCache[key]
	recipeBatchLipidInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			recipeBatchLipidAllColumns,
			recipeBatchLipidColumnsWithDefault,
			recipeBatchLipidColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(recipeBatchLipidType, recipeBatchLipidMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(recipeBatchLipidType, recipeBatchLipidMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"recipe_batch_lipid\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"recipe_batch_lipid\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into recipe_batch_lipid")
	}

	if !cached {
		recipeBatchLipidInsertCacheMut.Lock()
		recipeBatchLipidInsertCache[key] = cache
		recipeBatchLipidInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the RecipeBatchLipid.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *RecipeBatchLipid) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	recipeBatchLipidUpdateCacheMut.RLock()
	cache, cached := recipeBatchLipidUpdateCache[key]
	recipeBatchLipidUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			recipeBatchLipidAllColumns,
			recipeBatchLipidPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update recipe_batch_lipid, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"recipe_batch_lipid\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, recipeBatchLipidPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(recipeBatchLipidType, recipeBatchLipidMapping, append(wl, recipeBatchLipidPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update recipe_batch_lipid row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for recipe_batch_lipid")
	}

	if !cached {
		recipeBatchLipidUpdateCacheMut.Lock()
		recipeBatchLipidUpdateCache[key] = cache
		recipeBatchLipidUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q recipeBatchLipidQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for recipe_batch_lipid")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for recipe_batch_lipid")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RecipeBatchLipidSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), recipeBatchLipidPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"recipe_batch_lipid\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, recipeBatchLipidPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in recipeBatchLipid slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all recipeBatchLipid")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RecipeBatchLipid) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no recipe_batch_lipid provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(recipeBatchLipidColumnsWithDefault, o)

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

	recipeBatchLipidUpsertCacheMut.RLock()
	cache, cached := recipeBatchLipidUpsertCache[key]
	recipeBatchLipidUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			recipeBatchLipidAllColumns,
			recipeBatchLipidColumnsWithDefault,
			recipeBatchLipidColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			recipeBatchLipidAllColumns,
			recipeBatchLipidPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert recipe_batch_lipid, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(recipeBatchLipidPrimaryKeyColumns))
			copy(conflict, recipeBatchLipidPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"recipe_batch_lipid\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(recipeBatchLipidType, recipeBatchLipidMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(recipeBatchLipidType, recipeBatchLipidMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert recipe_batch_lipid")
	}

	if !cached {
		recipeBatchLipidUpsertCacheMut.Lock()
		recipeBatchLipidUpsertCache[key] = cache
		recipeBatchLipidUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single RecipeBatchLipid record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *RecipeBatchLipid) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RecipeBatchLipid provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), recipeBatchLipidPrimaryKeyMapping)
	sql := "DELETE FROM \"recipe_batch_lipid\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from recipe_batch_lipid")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for recipe_batch_lipid")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q recipeBatchLipidQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no recipeBatchLipidQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from recipe_batch_lipid")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for recipe_batch_lipid")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RecipeBatchLipidSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(recipeBatchLipidBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), recipeBatchLipidPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"recipe_batch_lipid\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, recipeBatchLipidPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from recipeBatchLipid slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for recipe_batch_lipid")
	}

	if len(recipeBatchLipidAfterDeleteHooks) != 0 {
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
func (o *RecipeBatchLipid) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRecipeBatchLipid(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RecipeBatchLipidSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RecipeBatchLipidSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), recipeBatchLipidPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"recipe_batch_lipid\".* FROM \"recipe_batch_lipid\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, recipeBatchLipidPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RecipeBatchLipidSlice")
	}

	*o = slice

	return nil
}

// RecipeBatchLipidExists checks if the RecipeBatchLipid row exists.
func RecipeBatchLipidExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"recipe_batch_lipid\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if recipe_batch_lipid exists")
	}

	return exists, nil
}
