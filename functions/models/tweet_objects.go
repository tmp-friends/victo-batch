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

// TweetObject is an object representing the database table.
type TweetObject struct {
	ID           string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Text         null.String `boil:"text" json:"text,omitempty" toml:"text" yaml:"text,omitempty"`
	RetweetCount int         `boil:"retweet_count" json:"retweet_count" toml:"retweet_count" yaml:"retweet_count"`
	LikeCount    int         `boil:"like_count" json:"like_count" toml:"like_count" yaml:"like_count"`
	AuthorID     string      `boil:"author_id" json:"author_id" toml:"author_id" yaml:"author_id"`
	URL          string      `boil:"url" json:"url" toml:"url" yaml:"url"`
	TweetedAt    time.Time   `boil:"tweeted_at" json:"tweeted_at" toml:"tweeted_at" yaml:"tweeted_at"`
	CreatedAt    time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt    time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	HashtagID    int         `boil:"hashtag_id" json:"hashtag_id" toml:"hashtag_id" yaml:"hashtag_id"`

	R *tweetObjectR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L tweetObjectL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TweetObjectColumns = struct {
	ID           string
	Text         string
	RetweetCount string
	LikeCount    string
	AuthorID     string
	URL          string
	TweetedAt    string
	CreatedAt    string
	UpdatedAt    string
	HashtagID    string
}{
	ID:           "id",
	Text:         "text",
	RetweetCount: "retweet_count",
	LikeCount:    "like_count",
	AuthorID:     "author_id",
	URL:          "url",
	TweetedAt:    "tweeted_at",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	HashtagID:    "hashtag_id",
}

var TweetObjectTableColumns = struct {
	ID           string
	Text         string
	RetweetCount string
	LikeCount    string
	AuthorID     string
	URL          string
	TweetedAt    string
	CreatedAt    string
	UpdatedAt    string
	HashtagID    string
}{
	ID:           "tweet_objects.id",
	Text:         "tweet_objects.text",
	RetweetCount: "tweet_objects.retweet_count",
	LikeCount:    "tweet_objects.like_count",
	AuthorID:     "tweet_objects.author_id",
	URL:          "tweet_objects.url",
	TweetedAt:    "tweet_objects.tweeted_at",
	CreatedAt:    "tweet_objects.created_at",
	UpdatedAt:    "tweet_objects.updated_at",
	HashtagID:    "tweet_objects.hashtag_id",
}

// Generated where

var TweetObjectWhere = struct {
	ID           whereHelperstring
	Text         whereHelpernull_String
	RetweetCount whereHelperint
	LikeCount    whereHelperint
	AuthorID     whereHelperstring
	URL          whereHelperstring
	TweetedAt    whereHelpertime_Time
	CreatedAt    whereHelpertime_Time
	UpdatedAt    whereHelpertime_Time
	HashtagID    whereHelperint
}{
	ID:           whereHelperstring{field: "`tweet_objects`.`id`"},
	Text:         whereHelpernull_String{field: "`tweet_objects`.`text`"},
	RetweetCount: whereHelperint{field: "`tweet_objects`.`retweet_count`"},
	LikeCount:    whereHelperint{field: "`tweet_objects`.`like_count`"},
	AuthorID:     whereHelperstring{field: "`tweet_objects`.`author_id`"},
	URL:          whereHelperstring{field: "`tweet_objects`.`url`"},
	TweetedAt:    whereHelpertime_Time{field: "`tweet_objects`.`tweeted_at`"},
	CreatedAt:    whereHelpertime_Time{field: "`tweet_objects`.`created_at`"},
	UpdatedAt:    whereHelpertime_Time{field: "`tweet_objects`.`updated_at`"},
	HashtagID:    whereHelperint{field: "`tweet_objects`.`hashtag_id`"},
}

// TweetObjectRels is where relationship names are stored.
var TweetObjectRels = struct {
}{}

// tweetObjectR is where relationships are stored.
type tweetObjectR struct {
}

// NewStruct creates a new relationship struct
func (*tweetObjectR) NewStruct() *tweetObjectR {
	return &tweetObjectR{}
}

// tweetObjectL is where Load methods for each relationship are stored.
type tweetObjectL struct{}

var (
	tweetObjectAllColumns            = []string{"id", "text", "retweet_count", "like_count", "author_id", "url", "tweeted_at", "created_at", "updated_at", "hashtag_id"}
	tweetObjectColumnsWithoutDefault = []string{"id", "text", "retweet_count", "like_count", "author_id", "url", "tweeted_at", "hashtag_id"}
	tweetObjectColumnsWithDefault    = []string{"created_at", "updated_at"}
	tweetObjectPrimaryKeyColumns     = []string{"id"}
	tweetObjectGeneratedColumns      = []string{}
)

type (
	// TweetObjectSlice is an alias for a slice of pointers to TweetObject.
	// This should almost always be used instead of []TweetObject.
	TweetObjectSlice []*TweetObject
	// TweetObjectHook is the signature for custom TweetObject hook methods
	TweetObjectHook func(context.Context, boil.ContextExecutor, *TweetObject) error

	tweetObjectQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	tweetObjectType                 = reflect.TypeOf(&TweetObject{})
	tweetObjectMapping              = queries.MakeStructMapping(tweetObjectType)
	tweetObjectPrimaryKeyMapping, _ = queries.BindMapping(tweetObjectType, tweetObjectMapping, tweetObjectPrimaryKeyColumns)
	tweetObjectInsertCacheMut       sync.RWMutex
	tweetObjectInsertCache          = make(map[string]insertCache)
	tweetObjectUpdateCacheMut       sync.RWMutex
	tweetObjectUpdateCache          = make(map[string]updateCache)
	tweetObjectUpsertCacheMut       sync.RWMutex
	tweetObjectUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var tweetObjectAfterSelectHooks []TweetObjectHook

var tweetObjectBeforeInsertHooks []TweetObjectHook
var tweetObjectAfterInsertHooks []TweetObjectHook

var tweetObjectBeforeUpdateHooks []TweetObjectHook
var tweetObjectAfterUpdateHooks []TweetObjectHook

var tweetObjectBeforeDeleteHooks []TweetObjectHook
var tweetObjectAfterDeleteHooks []TweetObjectHook

var tweetObjectBeforeUpsertHooks []TweetObjectHook
var tweetObjectAfterUpsertHooks []TweetObjectHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TweetObject) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tweetObjectAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TweetObject) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tweetObjectBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TweetObject) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tweetObjectAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TweetObject) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tweetObjectBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TweetObject) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tweetObjectAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TweetObject) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tweetObjectBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TweetObject) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tweetObjectAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TweetObject) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tweetObjectBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TweetObject) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tweetObjectAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTweetObjectHook registers your hook function for all future operations.
func AddTweetObjectHook(hookPoint boil.HookPoint, tweetObjectHook TweetObjectHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		tweetObjectAfterSelectHooks = append(tweetObjectAfterSelectHooks, tweetObjectHook)
	case boil.BeforeInsertHook:
		tweetObjectBeforeInsertHooks = append(tweetObjectBeforeInsertHooks, tweetObjectHook)
	case boil.AfterInsertHook:
		tweetObjectAfterInsertHooks = append(tweetObjectAfterInsertHooks, tweetObjectHook)
	case boil.BeforeUpdateHook:
		tweetObjectBeforeUpdateHooks = append(tweetObjectBeforeUpdateHooks, tweetObjectHook)
	case boil.AfterUpdateHook:
		tweetObjectAfterUpdateHooks = append(tweetObjectAfterUpdateHooks, tweetObjectHook)
	case boil.BeforeDeleteHook:
		tweetObjectBeforeDeleteHooks = append(tweetObjectBeforeDeleteHooks, tweetObjectHook)
	case boil.AfterDeleteHook:
		tweetObjectAfterDeleteHooks = append(tweetObjectAfterDeleteHooks, tweetObjectHook)
	case boil.BeforeUpsertHook:
		tweetObjectBeforeUpsertHooks = append(tweetObjectBeforeUpsertHooks, tweetObjectHook)
	case boil.AfterUpsertHook:
		tweetObjectAfterUpsertHooks = append(tweetObjectAfterUpsertHooks, tweetObjectHook)
	}
}

// OneG returns a single tweetObject record from the query using the global executor.
func (q tweetObjectQuery) OneG(ctx context.Context) (*TweetObject, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single tweetObject record from the query.
func (q tweetObjectQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TweetObject, error) {
	o := &TweetObject{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for tweet_objects")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all TweetObject records from the query using the global executor.
func (q tweetObjectQuery) AllG(ctx context.Context) (TweetObjectSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all TweetObject records from the query.
func (q tweetObjectQuery) All(ctx context.Context, exec boil.ContextExecutor) (TweetObjectSlice, error) {
	var o []*TweetObject

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TweetObject slice")
	}

	if len(tweetObjectAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all TweetObject records in the query using the global executor
func (q tweetObjectQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all TweetObject records in the query.
func (q tweetObjectQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count tweet_objects rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q tweetObjectQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q tweetObjectQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if tweet_objects exists")
	}

	return count > 0, nil
}

// TweetObjects retrieves all the records using an executor.
func TweetObjects(mods ...qm.QueryMod) tweetObjectQuery {
	mods = append(mods, qm.From("`tweet_objects`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`tweet_objects`.*"})
	}

	return tweetObjectQuery{q}
}

// FindTweetObjectG retrieves a single record by ID.
func FindTweetObjectG(ctx context.Context, iD string, selectCols ...string) (*TweetObject, error) {
	return FindTweetObject(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindTweetObject retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTweetObject(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*TweetObject, error) {
	tweetObjectObj := &TweetObject{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `tweet_objects` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, tweetObjectObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from tweet_objects")
	}

	if err = tweetObjectObj.doAfterSelectHooks(ctx, exec); err != nil {
		return tweetObjectObj, err
	}

	return tweetObjectObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *TweetObject) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TweetObject) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no tweet_objects provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(tweetObjectColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	tweetObjectInsertCacheMut.RLock()
	cache, cached := tweetObjectInsertCache[key]
	tweetObjectInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			tweetObjectAllColumns,
			tweetObjectColumnsWithDefault,
			tweetObjectColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(tweetObjectType, tweetObjectMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(tweetObjectType, tweetObjectMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `tweet_objects` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `tweet_objects` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `tweet_objects` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, tweetObjectPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into tweet_objects")
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
		return errors.Wrap(err, "models: unable to populate default values for tweet_objects")
	}

CacheNoHooks:
	if !cached {
		tweetObjectInsertCacheMut.Lock()
		tweetObjectInsertCache[key] = cache
		tweetObjectInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single TweetObject record using the global executor.
// See Update for more documentation.
func (o *TweetObject) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the TweetObject.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TweetObject) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	tweetObjectUpdateCacheMut.RLock()
	cache, cached := tweetObjectUpdateCache[key]
	tweetObjectUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			tweetObjectAllColumns,
			tweetObjectPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update tweet_objects, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `tweet_objects` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, tweetObjectPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(tweetObjectType, tweetObjectMapping, append(wl, tweetObjectPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update tweet_objects row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for tweet_objects")
	}

	if !cached {
		tweetObjectUpdateCacheMut.Lock()
		tweetObjectUpdateCache[key] = cache
		tweetObjectUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q tweetObjectQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q tweetObjectQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for tweet_objects")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for tweet_objects")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o TweetObjectSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TweetObjectSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tweetObjectPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `tweet_objects` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tweetObjectPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in tweetObject slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all tweetObject")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *TweetObject) UpsertG(ctx context.Context, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateColumns, insertColumns)
}

var mySQLTweetObjectUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TweetObject) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no tweet_objects provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(tweetObjectColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLTweetObjectUniqueColumns, o)

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

	tweetObjectUpsertCacheMut.RLock()
	cache, cached := tweetObjectUpsertCache[key]
	tweetObjectUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			tweetObjectAllColumns,
			tweetObjectColumnsWithDefault,
			tweetObjectColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			tweetObjectAllColumns,
			tweetObjectPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert tweet_objects, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`tweet_objects`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `tweet_objects` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(tweetObjectType, tweetObjectMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(tweetObjectType, tweetObjectMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for tweet_objects")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(tweetObjectType, tweetObjectMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for tweet_objects")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for tweet_objects")
	}

CacheNoHooks:
	if !cached {
		tweetObjectUpsertCacheMut.Lock()
		tweetObjectUpsertCache[key] = cache
		tweetObjectUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single TweetObject record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *TweetObject) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single TweetObject record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TweetObject) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TweetObject provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), tweetObjectPrimaryKeyMapping)
	sql := "DELETE FROM `tweet_objects` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from tweet_objects")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for tweet_objects")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q tweetObjectQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q tweetObjectQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no tweetObjectQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from tweet_objects")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for tweet_objects")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o TweetObjectSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TweetObjectSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(tweetObjectBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tweetObjectPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `tweet_objects` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tweetObjectPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from tweetObject slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for tweet_objects")
	}

	if len(tweetObjectAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *TweetObject) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no TweetObject provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *TweetObject) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTweetObject(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TweetObjectSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty TweetObjectSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TweetObjectSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TweetObjectSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tweetObjectPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `tweet_objects`.* FROM `tweet_objects` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tweetObjectPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TweetObjectSlice")
	}

	*o = slice

	return nil
}

// TweetObjectExistsG checks if the TweetObject row exists.
func TweetObjectExistsG(ctx context.Context, iD string) (bool, error) {
	return TweetObjectExists(ctx, boil.GetContextDB(), iD)
}

// TweetObjectExists checks if the TweetObject row exists.
func TweetObjectExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `tweet_objects` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if tweet_objects exists")
	}

	return exists, nil
}

// Exists checks if the TweetObject row exists.
func (o *TweetObject) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return TweetObjectExists(ctx, exec, o.ID)
}
