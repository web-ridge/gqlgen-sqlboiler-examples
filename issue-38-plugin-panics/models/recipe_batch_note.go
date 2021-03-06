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

// RecipeBatchNote is an object representing the database table.
type RecipeBatchNote struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	DeletedAt null.Time `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	Note      string    `boil:"note" json:"note" toml:"note" yaml:"note"`
	Link      string    `boil:"link" json:"link" toml:"link" yaml:"link"`
	BatchID   int       `boil:"batch_id" json:"batch_id" toml:"batch_id" yaml:"batch_id"`

	R *recipeBatchNoteR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L recipeBatchNoteL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RecipeBatchNoteColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	Note      string
	Link      string
	BatchID   string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Note:      "note",
	Link:      "link",
	BatchID:   "batch_id",
}

// Generated where

var RecipeBatchNoteWhere = struct {
	ID        whereHelperint
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
	DeletedAt whereHelpernull_Time
	Note      whereHelperstring
	Link      whereHelperstring
	BatchID   whereHelperint
}{
	ID:        whereHelperint{field: "\"recipe_batch_note\".\"id\""},
	CreatedAt: whereHelpertime_Time{field: "\"recipe_batch_note\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"recipe_batch_note\".\"updated_at\""},
	DeletedAt: whereHelpernull_Time{field: "\"recipe_batch_note\".\"deleted_at\""},
	Note:      whereHelperstring{field: "\"recipe_batch_note\".\"note\""},
	Link:      whereHelperstring{field: "\"recipe_batch_note\".\"link\""},
	BatchID:   whereHelperint{field: "\"recipe_batch_note\".\"batch_id\""},
}

// RecipeBatchNoteRels is where relationship names are stored.
var RecipeBatchNoteRels = struct {
	Batch string
}{
	Batch: "Batch",
}

// recipeBatchNoteR is where relationships are stored.
type recipeBatchNoteR struct {
	Batch *RecipeBatch `boil:"Batch" json:"Batch" toml:"Batch" yaml:"Batch"`
}

// NewStruct creates a new relationship struct
func (*recipeBatchNoteR) NewStruct() *recipeBatchNoteR {
	return &recipeBatchNoteR{}
}

// recipeBatchNoteL is where Load methods for each relationship are stored.
type recipeBatchNoteL struct{}

var (
	recipeBatchNoteAllColumns            = []string{"id", "created_at", "updated_at", "deleted_at", "note", "link", "batch_id"}
	recipeBatchNoteColumnsWithoutDefault = []string{"created_at", "updated_at", "deleted_at", "note", "link", "batch_id"}
	recipeBatchNoteColumnsWithDefault    = []string{"id"}
	recipeBatchNotePrimaryKeyColumns     = []string{"id"}
)

type (
	// RecipeBatchNoteSlice is an alias for a slice of pointers to RecipeBatchNote.
	// This should generally be used opposed to []RecipeBatchNote.
	RecipeBatchNoteSlice []*RecipeBatchNote
	// RecipeBatchNoteHook is the signature for custom RecipeBatchNote hook methods
	RecipeBatchNoteHook func(context.Context, boil.ContextExecutor, *RecipeBatchNote) error

	recipeBatchNoteQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	recipeBatchNoteType                 = reflect.TypeOf(&RecipeBatchNote{})
	recipeBatchNoteMapping              = queries.MakeStructMapping(recipeBatchNoteType)
	recipeBatchNotePrimaryKeyMapping, _ = queries.BindMapping(recipeBatchNoteType, recipeBatchNoteMapping, recipeBatchNotePrimaryKeyColumns)
	recipeBatchNoteInsertCacheMut       sync.RWMutex
	recipeBatchNoteInsertCache          = make(map[string]insertCache)
	recipeBatchNoteUpdateCacheMut       sync.RWMutex
	recipeBatchNoteUpdateCache          = make(map[string]updateCache)
	recipeBatchNoteUpsertCacheMut       sync.RWMutex
	recipeBatchNoteUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var recipeBatchNoteBeforeInsertHooks []RecipeBatchNoteHook
var recipeBatchNoteBeforeUpdateHooks []RecipeBatchNoteHook
var recipeBatchNoteBeforeDeleteHooks []RecipeBatchNoteHook
var recipeBatchNoteBeforeUpsertHooks []RecipeBatchNoteHook

var recipeBatchNoteAfterInsertHooks []RecipeBatchNoteHook
var recipeBatchNoteAfterSelectHooks []RecipeBatchNoteHook
var recipeBatchNoteAfterUpdateHooks []RecipeBatchNoteHook
var recipeBatchNoteAfterDeleteHooks []RecipeBatchNoteHook
var recipeBatchNoteAfterUpsertHooks []RecipeBatchNoteHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *RecipeBatchNote) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchNoteBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *RecipeBatchNote) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchNoteBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *RecipeBatchNote) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchNoteBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *RecipeBatchNote) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchNoteBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *RecipeBatchNote) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchNoteAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *RecipeBatchNote) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchNoteAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *RecipeBatchNote) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchNoteAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *RecipeBatchNote) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchNoteAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *RecipeBatchNote) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchNoteAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRecipeBatchNoteHook registers your hook function for all future operations.
func AddRecipeBatchNoteHook(hookPoint boil.HookPoint, recipeBatchNoteHook RecipeBatchNoteHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		recipeBatchNoteBeforeInsertHooks = append(recipeBatchNoteBeforeInsertHooks, recipeBatchNoteHook)
	case boil.BeforeUpdateHook:
		recipeBatchNoteBeforeUpdateHooks = append(recipeBatchNoteBeforeUpdateHooks, recipeBatchNoteHook)
	case boil.BeforeDeleteHook:
		recipeBatchNoteBeforeDeleteHooks = append(recipeBatchNoteBeforeDeleteHooks, recipeBatchNoteHook)
	case boil.BeforeUpsertHook:
		recipeBatchNoteBeforeUpsertHooks = append(recipeBatchNoteBeforeUpsertHooks, recipeBatchNoteHook)
	case boil.AfterInsertHook:
		recipeBatchNoteAfterInsertHooks = append(recipeBatchNoteAfterInsertHooks, recipeBatchNoteHook)
	case boil.AfterSelectHook:
		recipeBatchNoteAfterSelectHooks = append(recipeBatchNoteAfterSelectHooks, recipeBatchNoteHook)
	case boil.AfterUpdateHook:
		recipeBatchNoteAfterUpdateHooks = append(recipeBatchNoteAfterUpdateHooks, recipeBatchNoteHook)
	case boil.AfterDeleteHook:
		recipeBatchNoteAfterDeleteHooks = append(recipeBatchNoteAfterDeleteHooks, recipeBatchNoteHook)
	case boil.AfterUpsertHook:
		recipeBatchNoteAfterUpsertHooks = append(recipeBatchNoteAfterUpsertHooks, recipeBatchNoteHook)
	}
}

// One returns a single recipeBatchNote record from the query.
func (q recipeBatchNoteQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RecipeBatchNote, error) {
	o := &RecipeBatchNote{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for recipe_batch_note")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all RecipeBatchNote records from the query.
func (q recipeBatchNoteQuery) All(ctx context.Context, exec boil.ContextExecutor) (RecipeBatchNoteSlice, error) {
	var o []*RecipeBatchNote

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to RecipeBatchNote slice")
	}

	if len(recipeBatchNoteAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all RecipeBatchNote records in the query.
func (q recipeBatchNoteQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count recipe_batch_note rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q recipeBatchNoteQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if recipe_batch_note exists")
	}

	return count > 0, nil
}

// Batch pointed to by the foreign key.
func (o *RecipeBatchNote) Batch(mods ...qm.QueryMod) recipeBatchQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.BatchID),
	}

	queryMods = append(queryMods, mods...)

	query := RecipeBatches(queryMods...)
	queries.SetFrom(query.Query, "\"recipe_batch\"")

	return query
}

// LoadBatch allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (recipeBatchNoteL) LoadBatch(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRecipeBatchNote interface{}, mods queries.Applicator) error {
	var slice []*RecipeBatchNote
	var object *RecipeBatchNote

	if singular {
		object = maybeRecipeBatchNote.(*RecipeBatchNote)
	} else {
		slice = *maybeRecipeBatchNote.(*[]*RecipeBatchNote)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &recipeBatchNoteR{}
		}
		args = append(args, object.BatchID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &recipeBatchNoteR{}
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

	if len(recipeBatchNoteAfterSelectHooks) != 0 {
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

// SetBatch of the recipeBatchNote to the related item.
// Sets o.R.Batch to related.
// Adds o to related.R.BatchRecipeBatchNotes.
func (o *RecipeBatchNote) SetBatch(ctx context.Context, exec boil.ContextExecutor, insert bool, related *RecipeBatch) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"recipe_batch_note\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"batch_id"}),
		strmangle.WhereClause("\"", "\"", 2, recipeBatchNotePrimaryKeyColumns),
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
		o.R = &recipeBatchNoteR{
			Batch: related,
		}
	} else {
		o.R.Batch = related
	}

	if related.R == nil {
		related.R = &recipeBatchR{
			BatchRecipeBatchNotes: RecipeBatchNoteSlice{o},
		}
	} else {
		related.R.BatchRecipeBatchNotes = append(related.R.BatchRecipeBatchNotes, o)
	}

	return nil
}

// RecipeBatchNotes retrieves all the records using an executor.
func RecipeBatchNotes(mods ...qm.QueryMod) recipeBatchNoteQuery {
	mods = append(mods, qm.From("\"recipe_batch_note\""))
	return recipeBatchNoteQuery{NewQuery(mods...)}
}

// FindRecipeBatchNote retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRecipeBatchNote(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*RecipeBatchNote, error) {
	recipeBatchNoteObj := &RecipeBatchNote{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"recipe_batch_note\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, recipeBatchNoteObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from recipe_batch_note")
	}

	return recipeBatchNoteObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RecipeBatchNote) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no recipe_batch_note provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(recipeBatchNoteColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	recipeBatchNoteInsertCacheMut.RLock()
	cache, cached := recipeBatchNoteInsertCache[key]
	recipeBatchNoteInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			recipeBatchNoteAllColumns,
			recipeBatchNoteColumnsWithDefault,
			recipeBatchNoteColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(recipeBatchNoteType, recipeBatchNoteMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(recipeBatchNoteType, recipeBatchNoteMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"recipe_batch_note\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"recipe_batch_note\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into recipe_batch_note")
	}

	if !cached {
		recipeBatchNoteInsertCacheMut.Lock()
		recipeBatchNoteInsertCache[key] = cache
		recipeBatchNoteInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the RecipeBatchNote.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *RecipeBatchNote) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	recipeBatchNoteUpdateCacheMut.RLock()
	cache, cached := recipeBatchNoteUpdateCache[key]
	recipeBatchNoteUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			recipeBatchNoteAllColumns,
			recipeBatchNotePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update recipe_batch_note, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"recipe_batch_note\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, recipeBatchNotePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(recipeBatchNoteType, recipeBatchNoteMapping, append(wl, recipeBatchNotePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update recipe_batch_note row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for recipe_batch_note")
	}

	if !cached {
		recipeBatchNoteUpdateCacheMut.Lock()
		recipeBatchNoteUpdateCache[key] = cache
		recipeBatchNoteUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q recipeBatchNoteQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for recipe_batch_note")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for recipe_batch_note")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RecipeBatchNoteSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), recipeBatchNotePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"recipe_batch_note\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, recipeBatchNotePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in recipeBatchNote slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all recipeBatchNote")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RecipeBatchNote) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no recipe_batch_note provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(recipeBatchNoteColumnsWithDefault, o)

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

	recipeBatchNoteUpsertCacheMut.RLock()
	cache, cached := recipeBatchNoteUpsertCache[key]
	recipeBatchNoteUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			recipeBatchNoteAllColumns,
			recipeBatchNoteColumnsWithDefault,
			recipeBatchNoteColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			recipeBatchNoteAllColumns,
			recipeBatchNotePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert recipe_batch_note, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(recipeBatchNotePrimaryKeyColumns))
			copy(conflict, recipeBatchNotePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"recipe_batch_note\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(recipeBatchNoteType, recipeBatchNoteMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(recipeBatchNoteType, recipeBatchNoteMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert recipe_batch_note")
	}

	if !cached {
		recipeBatchNoteUpsertCacheMut.Lock()
		recipeBatchNoteUpsertCache[key] = cache
		recipeBatchNoteUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single RecipeBatchNote record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *RecipeBatchNote) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RecipeBatchNote provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), recipeBatchNotePrimaryKeyMapping)
	sql := "DELETE FROM \"recipe_batch_note\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from recipe_batch_note")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for recipe_batch_note")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q recipeBatchNoteQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no recipeBatchNoteQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from recipe_batch_note")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for recipe_batch_note")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RecipeBatchNoteSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(recipeBatchNoteBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), recipeBatchNotePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"recipe_batch_note\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, recipeBatchNotePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from recipeBatchNote slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for recipe_batch_note")
	}

	if len(recipeBatchNoteAfterDeleteHooks) != 0 {
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
func (o *RecipeBatchNote) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRecipeBatchNote(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RecipeBatchNoteSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RecipeBatchNoteSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), recipeBatchNotePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"recipe_batch_note\".* FROM \"recipe_batch_note\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, recipeBatchNotePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RecipeBatchNoteSlice")
	}

	*o = slice

	return nil
}

// RecipeBatchNoteExists checks if the RecipeBatchNote row exists.
func RecipeBatchNoteExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"recipe_batch_note\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if recipe_batch_note exists")
	}

	return exists, nil
}
