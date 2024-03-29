// Code generated by SQLBoiler 4.14.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// Vtuber is an object representing the database table.
type Vtuber struct {
	ID              int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name            string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	BelongsTo       null.String `boil:"belongs_to" json:"belongs_to,omitempty" toml:"belongs_to" yaml:"belongs_to,omitempty"`
	ProfileImageURL null.String `boil:"profile_image_url" json:"profile_image_url,omitempty" toml:"profile_image_url" yaml:"profile_image_url,omitempty"`
	TwitterUserName null.String `boil:"twitter_user_name" json:"twitter_user_name,omitempty" toml:"twitter_user_name" yaml:"twitter_user_name,omitempty"`
	Channel         null.String `boil:"channel" json:"channel,omitempty" toml:"channel" yaml:"channel,omitempty"`
	CreatedAt       time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt       time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *vtuberR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L vtuberL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var VtuberColumns = struct {
	ID              string
	Name            string
	BelongsTo       string
	ProfileImageURL string
	TwitterUserName string
	Channel         string
	CreatedAt       string
	UpdatedAt       string
}{
	ID:              "id",
	Name:            "name",
	BelongsTo:       "belongs_to",
	ProfileImageURL: "profile_image_url",
	TwitterUserName: "twitter_user_name",
	Channel:         "channel",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}

var VtuberTableColumns = struct {
	ID              string
	Name            string
	BelongsTo       string
	ProfileImageURL string
	TwitterUserName string
	Channel         string
	CreatedAt       string
	UpdatedAt       string
}{
	ID:              "vtubers.id",
	Name:            "vtubers.name",
	BelongsTo:       "vtubers.belongs_to",
	ProfileImageURL: "vtubers.profile_image_url",
	TwitterUserName: "vtubers.twitter_user_name",
	Channel:         "vtubers.channel",
	CreatedAt:       "vtubers.created_at",
	UpdatedAt:       "vtubers.updated_at",
}

// Generated where

var VtuberWhere = struct {
	ID              whereHelperint
	Name            whereHelperstring
	BelongsTo       whereHelpernull_String
	ProfileImageURL whereHelpernull_String
	TwitterUserName whereHelpernull_String
	Channel         whereHelpernull_String
	CreatedAt       whereHelpertime_Time
	UpdatedAt       whereHelpertime_Time
}{
	ID:              whereHelperint{field: "`vtubers`.`id`"},
	Name:            whereHelperstring{field: "`vtubers`.`name`"},
	BelongsTo:       whereHelpernull_String{field: "`vtubers`.`belongs_to`"},
	ProfileImageURL: whereHelpernull_String{field: "`vtubers`.`profile_image_url`"},
	TwitterUserName: whereHelpernull_String{field: "`vtubers`.`twitter_user_name`"},
	Channel:         whereHelpernull_String{field: "`vtubers`.`channel`"},
	CreatedAt:       whereHelpertime_Time{field: "`vtubers`.`created_at`"},
	UpdatedAt:       whereHelpertime_Time{field: "`vtubers`.`updated_at`"},
}

// VtuberRels is where relationship names are stored.
var VtuberRels = struct {
}{}

// vtuberR is where relationships are stored.
type vtuberR struct {
}

// NewStruct creates a new relationship struct
func (*vtuberR) NewStruct() *vtuberR {
	return &vtuberR{}
}

// vtuberL is where Load methods for each relationship are stored.
type vtuberL struct{}

var (
	vtuberAllColumns            = []string{"id", "name", "belongs_to", "profile_image_url", "twitter_user_name", "channel", "created_at", "updated_at"}
	vtuberColumnsWithoutDefault = []string{"id", "name", "belongs_to", "profile_image_url", "twitter_user_name", "channel"}
	vtuberColumnsWithDefault    = []string{"created_at", "updated_at"}
	vtuberPrimaryKeyColumns     = []string{"id"}
	vtuberGeneratedColumns      = []string{}
)

type (
	// VtuberSlice is an alias for a slice of pointers to Vtuber.
	// This should almost always be used instead of []Vtuber.
	VtuberSlice []*Vtuber
	// VtuberHook is the signature for custom Vtuber hook methods
	VtuberHook func(context.Context, boil.ContextExecutor, *Vtuber) error

	vtuberQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	vtuberType                 = reflect.TypeOf(&Vtuber{})
	vtuberMapping              = queries.MakeStructMapping(vtuberType)
	vtuberPrimaryKeyMapping, _ = queries.BindMapping(vtuberType, vtuberMapping, vtuberPrimaryKeyColumns)
	vtuberInsertCacheMut       sync.RWMutex
	vtuberInsertCache          = make(map[string]insertCache)
	vtuberUpdateCacheMut       sync.RWMutex
	vtuberUpdateCache          = make(map[string]updateCache)
	vtuberUpsertCacheMut       sync.RWMutex
	vtuberUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var vtuberAfterSelectHooks []VtuberHook

var vtuberBeforeInsertHooks []VtuberHook
var vtuberAfterInsertHooks []VtuberHook

var vtuberBeforeUpdateHooks []VtuberHook
var vtuberAfterUpdateHooks []VtuberHook

var vtuberBeforeDeleteHooks []VtuberHook
var vtuberAfterDeleteHooks []VtuberHook

var vtuberBeforeUpsertHooks []VtuberHook
var vtuberAfterUpsertHooks []VtuberHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Vtuber) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vtuberAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Vtuber) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vtuberBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Vtuber) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vtuberAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Vtuber) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vtuberBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Vtuber) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vtuberAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Vtuber) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vtuberBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Vtuber) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vtuberAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Vtuber) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vtuberBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Vtuber) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range vtuberAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddVtuberHook registers your hook function for all future operations.
func AddVtuberHook(hookPoint boil.HookPoint, vtuberHook VtuberHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		vtuberAfterSelectHooks = append(vtuberAfterSelectHooks, vtuberHook)
	case boil.BeforeInsertHook:
		vtuberBeforeInsertHooks = append(vtuberBeforeInsertHooks, vtuberHook)
	case boil.AfterInsertHook:
		vtuberAfterInsertHooks = append(vtuberAfterInsertHooks, vtuberHook)
	case boil.BeforeUpdateHook:
		vtuberBeforeUpdateHooks = append(vtuberBeforeUpdateHooks, vtuberHook)
	case boil.AfterUpdateHook:
		vtuberAfterUpdateHooks = append(vtuberAfterUpdateHooks, vtuberHook)
	case boil.BeforeDeleteHook:
		vtuberBeforeDeleteHooks = append(vtuberBeforeDeleteHooks, vtuberHook)
	case boil.AfterDeleteHook:
		vtuberAfterDeleteHooks = append(vtuberAfterDeleteHooks, vtuberHook)
	case boil.BeforeUpsertHook:
		vtuberBeforeUpsertHooks = append(vtuberBeforeUpsertHooks, vtuberHook)
	case boil.AfterUpsertHook:
		vtuberAfterUpsertHooks = append(vtuberAfterUpsertHooks, vtuberHook)
	}
}

// OneG returns a single vtuber record from the query using the global executor.
func (q vtuberQuery) OneG(ctx context.Context) (*Vtuber, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single vtuber record from the query.
func (q vtuberQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Vtuber, error) {
	o := &Vtuber{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for vtubers")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Vtuber records from the query using the global executor.
func (q vtuberQuery) AllG(ctx context.Context) (VtuberSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Vtuber records from the query.
func (q vtuberQuery) All(ctx context.Context, exec boil.ContextExecutor) (VtuberSlice, error) {
	var o []*Vtuber

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Vtuber slice")
	}

	if len(vtuberAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Vtuber records in the query using the global executor
func (q vtuberQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Vtuber records in the query.
func (q vtuberQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count vtubers rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q vtuberQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q vtuberQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if vtubers exists")
	}

	return count > 0, nil
}

// Vtubers retrieves all the records using an executor.
func Vtubers(mods ...qm.QueryMod) vtuberQuery {
	mods = append(mods, qm.From("`vtubers`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`vtubers`.*"})
	}

	return vtuberQuery{q}
}

// FindVtuberG retrieves a single record by ID.
func FindVtuberG(ctx context.Context, iD int, selectCols ...string) (*Vtuber, error) {
	return FindVtuber(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindVtuber retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindVtuber(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Vtuber, error) {
	vtuberObj := &Vtuber{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `vtubers` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, vtuberObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from vtubers")
	}

	if err = vtuberObj.doAfterSelectHooks(ctx, exec); err != nil {
		return vtuberObj, err
	}

	return vtuberObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Vtuber) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Vtuber) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no vtubers provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(vtuberColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	vtuberInsertCacheMut.RLock()
	cache, cached := vtuberInsertCache[key]
	vtuberInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			vtuberAllColumns,
			vtuberColumnsWithDefault,
			vtuberColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(vtuberType, vtuberMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(vtuberType, vtuberMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `vtubers` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `vtubers` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `vtubers` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, vtuberPrimaryKeyColumns))
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into vtubers")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
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
		return errors.Wrap(err, "models: unable to populate default values for vtubers")
	}

CacheNoHooks:
	if !cached {
		vtuberInsertCacheMut.Lock()
		vtuberInsertCache[key] = cache
		vtuberInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single Vtuber record using the global executor.
// See Update for more documentation.
func (o *Vtuber) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Vtuber.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Vtuber) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	vtuberUpdateCacheMut.RLock()
	cache, cached := vtuberUpdateCache[key]
	vtuberUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			vtuberAllColumns,
			vtuberPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update vtubers, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `vtubers` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, vtuberPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(vtuberType, vtuberMapping, append(wl, vtuberPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update vtubers row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for vtubers")
	}

	if !cached {
		vtuberUpdateCacheMut.Lock()
		vtuberUpdateCache[key] = cache
		vtuberUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q vtuberQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q vtuberQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for vtubers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for vtubers")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o VtuberSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o VtuberSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), vtuberPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `vtubers` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, vtuberPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in vtuber slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all vtuber")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Vtuber) UpsertG(ctx context.Context, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateColumns, insertColumns)
}

var mySQLVtuberUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Vtuber) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no vtubers provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(vtuberColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLVtuberUniqueColumns, o)

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

	vtuberUpsertCacheMut.RLock()
	cache, cached := vtuberUpsertCache[key]
	vtuberUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			vtuberAllColumns,
			vtuberColumnsWithDefault,
			vtuberColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			vtuberAllColumns,
			vtuberPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert vtubers, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`vtubers`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `vtubers` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(vtuberType, vtuberMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(vtuberType, vtuberMapping, ret)
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for vtubers")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(vtuberType, vtuberMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for vtubers")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for vtubers")
	}

CacheNoHooks:
	if !cached {
		vtuberUpsertCacheMut.Lock()
		vtuberUpsertCache[key] = cache
		vtuberUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single Vtuber record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Vtuber) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Vtuber record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Vtuber) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Vtuber provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), vtuberPrimaryKeyMapping)
	sql := "DELETE FROM `vtubers` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from vtubers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for vtubers")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q vtuberQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q vtuberQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no vtuberQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from vtubers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for vtubers")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o VtuberSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o VtuberSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(vtuberBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), vtuberPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `vtubers` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, vtuberPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from vtuber slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for vtubers")
	}

	if len(vtuberAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Vtuber) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no Vtuber provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Vtuber) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindVtuber(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *VtuberSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty VtuberSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *VtuberSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := VtuberSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), vtuberPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `vtubers`.* FROM `vtubers` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, vtuberPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in VtuberSlice")
	}

	*o = slice

	return nil
}

// VtuberExistsG checks if the Vtuber row exists.
func VtuberExistsG(ctx context.Context, iD int) (bool, error) {
	return VtuberExists(ctx, boil.GetContextDB(), iD)
}

// VtuberExists checks if the Vtuber row exists.
func VtuberExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `vtubers` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if vtubers exists")
	}

	return exists, nil
}

// Exists checks if the Vtuber row exists.
func (o *Vtuber) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return VtuberExists(ctx, exec, o.ID)
}
