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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// HashtagFollow is an object representing the database table.
type HashtagFollow struct {
	UserID    int       `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	HashtagID int       `boil:"hashtag_id" json:"hashtag_id" toml:"hashtag_id" yaml:"hashtag_id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *hashtagFollowR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L hashtagFollowL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var HashtagFollowColumns = struct {
	UserID    string
	HashtagID string
	CreatedAt string
	UpdatedAt string
}{
	UserID:    "user_id",
	HashtagID: "hashtag_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var HashtagFollowTableColumns = struct {
	UserID    string
	HashtagID string
	CreatedAt string
	UpdatedAt string
}{
	UserID:    "hashtag_follows.user_id",
	HashtagID: "hashtag_follows.hashtag_id",
	CreatedAt: "hashtag_follows.created_at",
	UpdatedAt: "hashtag_follows.updated_at",
}

// Generated where

var HashtagFollowWhere = struct {
	UserID    whereHelperint
	HashtagID whereHelperint
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	UserID:    whereHelperint{field: "`hashtag_follows`.`user_id`"},
	HashtagID: whereHelperint{field: "`hashtag_follows`.`hashtag_id`"},
	CreatedAt: whereHelpertime_Time{field: "`hashtag_follows`.`created_at`"},
	UpdatedAt: whereHelpertime_Time{field: "`hashtag_follows`.`updated_at`"},
}

// HashtagFollowRels is where relationship names are stored.
var HashtagFollowRels = struct {
}{}

// hashtagFollowR is where relationships are stored.
type hashtagFollowR struct {
}

// NewStruct creates a new relationship struct
func (*hashtagFollowR) NewStruct() *hashtagFollowR {
	return &hashtagFollowR{}
}

// hashtagFollowL is where Load methods for each relationship are stored.
type hashtagFollowL struct{}

var (
	hashtagFollowAllColumns            = []string{"user_id", "hashtag_id", "created_at", "updated_at"}
	hashtagFollowColumnsWithoutDefault = []string{"user_id", "hashtag_id"}
	hashtagFollowColumnsWithDefault    = []string{"created_at", "updated_at"}
	hashtagFollowPrimaryKeyColumns     = []string{"user_id", "hashtag_id"}
	hashtagFollowGeneratedColumns      = []string{}
)

type (
	// HashtagFollowSlice is an alias for a slice of pointers to HashtagFollow.
	// This should almost always be used instead of []HashtagFollow.
	HashtagFollowSlice []*HashtagFollow
	// HashtagFollowHook is the signature for custom HashtagFollow hook methods
	HashtagFollowHook func(context.Context, boil.ContextExecutor, *HashtagFollow) error

	hashtagFollowQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	hashtagFollowType                 = reflect.TypeOf(&HashtagFollow{})
	hashtagFollowMapping              = queries.MakeStructMapping(hashtagFollowType)
	hashtagFollowPrimaryKeyMapping, _ = queries.BindMapping(hashtagFollowType, hashtagFollowMapping, hashtagFollowPrimaryKeyColumns)
	hashtagFollowInsertCacheMut       sync.RWMutex
	hashtagFollowInsertCache          = make(map[string]insertCache)
	hashtagFollowUpdateCacheMut       sync.RWMutex
	hashtagFollowUpdateCache          = make(map[string]updateCache)
	hashtagFollowUpsertCacheMut       sync.RWMutex
	hashtagFollowUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var hashtagFollowAfterSelectHooks []HashtagFollowHook

var hashtagFollowBeforeInsertHooks []HashtagFollowHook
var hashtagFollowAfterInsertHooks []HashtagFollowHook

var hashtagFollowBeforeUpdateHooks []HashtagFollowHook
var hashtagFollowAfterUpdateHooks []HashtagFollowHook

var hashtagFollowBeforeDeleteHooks []HashtagFollowHook
var hashtagFollowAfterDeleteHooks []HashtagFollowHook

var hashtagFollowBeforeUpsertHooks []HashtagFollowHook
var hashtagFollowAfterUpsertHooks []HashtagFollowHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *HashtagFollow) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hashtagFollowAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *HashtagFollow) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hashtagFollowBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *HashtagFollow) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hashtagFollowAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *HashtagFollow) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hashtagFollowBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *HashtagFollow) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hashtagFollowAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *HashtagFollow) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hashtagFollowBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *HashtagFollow) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hashtagFollowAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *HashtagFollow) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hashtagFollowBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *HashtagFollow) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hashtagFollowAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddHashtagFollowHook registers your hook function for all future operations.
func AddHashtagFollowHook(hookPoint boil.HookPoint, hashtagFollowHook HashtagFollowHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		hashtagFollowAfterSelectHooks = append(hashtagFollowAfterSelectHooks, hashtagFollowHook)
	case boil.BeforeInsertHook:
		hashtagFollowBeforeInsertHooks = append(hashtagFollowBeforeInsertHooks, hashtagFollowHook)
	case boil.AfterInsertHook:
		hashtagFollowAfterInsertHooks = append(hashtagFollowAfterInsertHooks, hashtagFollowHook)
	case boil.BeforeUpdateHook:
		hashtagFollowBeforeUpdateHooks = append(hashtagFollowBeforeUpdateHooks, hashtagFollowHook)
	case boil.AfterUpdateHook:
		hashtagFollowAfterUpdateHooks = append(hashtagFollowAfterUpdateHooks, hashtagFollowHook)
	case boil.BeforeDeleteHook:
		hashtagFollowBeforeDeleteHooks = append(hashtagFollowBeforeDeleteHooks, hashtagFollowHook)
	case boil.AfterDeleteHook:
		hashtagFollowAfterDeleteHooks = append(hashtagFollowAfterDeleteHooks, hashtagFollowHook)
	case boil.BeforeUpsertHook:
		hashtagFollowBeforeUpsertHooks = append(hashtagFollowBeforeUpsertHooks, hashtagFollowHook)
	case boil.AfterUpsertHook:
		hashtagFollowAfterUpsertHooks = append(hashtagFollowAfterUpsertHooks, hashtagFollowHook)
	}
}

// OneG returns a single hashtagFollow record from the query using the global executor.
func (q hashtagFollowQuery) OneG(ctx context.Context) (*HashtagFollow, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single hashtagFollow record from the query.
func (q hashtagFollowQuery) One(ctx context.Context, exec boil.ContextExecutor) (*HashtagFollow, error) {
	o := &HashtagFollow{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for hashtag_follows")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all HashtagFollow records from the query using the global executor.
func (q hashtagFollowQuery) AllG(ctx context.Context) (HashtagFollowSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all HashtagFollow records from the query.
func (q hashtagFollowQuery) All(ctx context.Context, exec boil.ContextExecutor) (HashtagFollowSlice, error) {
	var o []*HashtagFollow

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to HashtagFollow slice")
	}

	if len(hashtagFollowAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all HashtagFollow records in the query using the global executor
func (q hashtagFollowQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all HashtagFollow records in the query.
func (q hashtagFollowQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count hashtag_follows rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q hashtagFollowQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q hashtagFollowQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if hashtag_follows exists")
	}

	return count > 0, nil
}

// HashtagFollows retrieves all the records using an executor.
func HashtagFollows(mods ...qm.QueryMod) hashtagFollowQuery {
	mods = append(mods, qm.From("`hashtag_follows`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`hashtag_follows`.*"})
	}

	return hashtagFollowQuery{q}
}

// FindHashtagFollowG retrieves a single record by ID.
func FindHashtagFollowG(ctx context.Context, userID int, hashtagID int, selectCols ...string) (*HashtagFollow, error) {
	return FindHashtagFollow(ctx, boil.GetContextDB(), userID, hashtagID, selectCols...)
}

// FindHashtagFollow retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindHashtagFollow(ctx context.Context, exec boil.ContextExecutor, userID int, hashtagID int, selectCols ...string) (*HashtagFollow, error) {
	hashtagFollowObj := &HashtagFollow{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `hashtag_follows` where `user_id`=? AND `hashtag_id`=?", sel,
	)

	q := queries.Raw(query, userID, hashtagID)

	err := q.Bind(ctx, exec, hashtagFollowObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from hashtag_follows")
	}

	if err = hashtagFollowObj.doAfterSelectHooks(ctx, exec); err != nil {
		return hashtagFollowObj, err
	}

	return hashtagFollowObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *HashtagFollow) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *HashtagFollow) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no hashtag_follows provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(hashtagFollowColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	hashtagFollowInsertCacheMut.RLock()
	cache, cached := hashtagFollowInsertCache[key]
	hashtagFollowInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			hashtagFollowAllColumns,
			hashtagFollowColumnsWithDefault,
			hashtagFollowColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(hashtagFollowType, hashtagFollowMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(hashtagFollowType, hashtagFollowMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `hashtag_follows` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `hashtag_follows` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `hashtag_follows` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, hashtagFollowPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into hashtag_follows")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.UserID,
		o.HashtagID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for hashtag_follows")
	}

CacheNoHooks:
	if !cached {
		hashtagFollowInsertCacheMut.Lock()
		hashtagFollowInsertCache[key] = cache
		hashtagFollowInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single HashtagFollow record using the global executor.
// See Update for more documentation.
func (o *HashtagFollow) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the HashtagFollow.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *HashtagFollow) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	hashtagFollowUpdateCacheMut.RLock()
	cache, cached := hashtagFollowUpdateCache[key]
	hashtagFollowUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			hashtagFollowAllColumns,
			hashtagFollowPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update hashtag_follows, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `hashtag_follows` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, hashtagFollowPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(hashtagFollowType, hashtagFollowMapping, append(wl, hashtagFollowPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update hashtag_follows row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for hashtag_follows")
	}

	if !cached {
		hashtagFollowUpdateCacheMut.Lock()
		hashtagFollowUpdateCache[key] = cache
		hashtagFollowUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q hashtagFollowQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q hashtagFollowQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for hashtag_follows")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for hashtag_follows")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o HashtagFollowSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o HashtagFollowSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), hashtagFollowPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `hashtag_follows` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, hashtagFollowPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in hashtagFollow slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all hashtagFollow")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *HashtagFollow) UpsertG(ctx context.Context, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateColumns, insertColumns)
}

var mySQLHashtagFollowUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *HashtagFollow) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no hashtag_follows provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(hashtagFollowColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLHashtagFollowUniqueColumns, o)

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

	hashtagFollowUpsertCacheMut.RLock()
	cache, cached := hashtagFollowUpsertCache[key]
	hashtagFollowUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			hashtagFollowAllColumns,
			hashtagFollowColumnsWithDefault,
			hashtagFollowColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			hashtagFollowAllColumns,
			hashtagFollowPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert hashtag_follows, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`hashtag_follows`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `hashtag_follows` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(hashtagFollowType, hashtagFollowMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(hashtagFollowType, hashtagFollowMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for hashtag_follows")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(hashtagFollowType, hashtagFollowMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for hashtag_follows")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for hashtag_follows")
	}

CacheNoHooks:
	if !cached {
		hashtagFollowUpsertCacheMut.Lock()
		hashtagFollowUpsertCache[key] = cache
		hashtagFollowUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single HashtagFollow record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *HashtagFollow) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single HashtagFollow record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *HashtagFollow) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no HashtagFollow provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), hashtagFollowPrimaryKeyMapping)
	sql := "DELETE FROM `hashtag_follows` WHERE `user_id`=? AND `hashtag_id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from hashtag_follows")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for hashtag_follows")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q hashtagFollowQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q hashtagFollowQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no hashtagFollowQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from hashtag_follows")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for hashtag_follows")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o HashtagFollowSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o HashtagFollowSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(hashtagFollowBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), hashtagFollowPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `hashtag_follows` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, hashtagFollowPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from hashtagFollow slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for hashtag_follows")
	}

	if len(hashtagFollowAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *HashtagFollow) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no HashtagFollow provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *HashtagFollow) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindHashtagFollow(ctx, exec, o.UserID, o.HashtagID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *HashtagFollowSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty HashtagFollowSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *HashtagFollowSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := HashtagFollowSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), hashtagFollowPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `hashtag_follows`.* FROM `hashtag_follows` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, hashtagFollowPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in HashtagFollowSlice")
	}

	*o = slice

	return nil
}

// HashtagFollowExistsG checks if the HashtagFollow row exists.
func HashtagFollowExistsG(ctx context.Context, userID int, hashtagID int) (bool, error) {
	return HashtagFollowExists(ctx, boil.GetContextDB(), userID, hashtagID)
}

// HashtagFollowExists checks if the HashtagFollow row exists.
func HashtagFollowExists(ctx context.Context, exec boil.ContextExecutor, userID int, hashtagID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `hashtag_follows` where `user_id`=? AND `hashtag_id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, userID, hashtagID)
	}
	row := exec.QueryRowContext(ctx, sql, userID, hashtagID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if hashtag_follows exists")
	}

	return exists, nil
}

// Exists checks if the HashtagFollow row exists.
func (o *HashtagFollow) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return HashtagFollowExists(ctx, exec, o.UserID, o.HashtagID)
}