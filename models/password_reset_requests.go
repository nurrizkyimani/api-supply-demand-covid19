// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// PasswordResetRequest is an object representing the database table.
type PasswordResetRequest struct {
	ID          string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID      string    `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	NewPassword string    `boil:"new_password" json:"new_password" toml:"new_password" yaml:"new_password"`
	Date        time.Time `boil:"date" json:"date" toml:"date" yaml:"date"`

	R *passwordResetRequestR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L passwordResetRequestL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PasswordResetRequestColumns = struct {
	ID          string
	UserID      string
	NewPassword string
	Date        string
}{
	ID:          "id",
	UserID:      "user_id",
	NewPassword: "new_password",
	Date:        "date",
}

// Generated where

var PasswordResetRequestWhere = struct {
	ID          whereHelperstring
	UserID      whereHelperstring
	NewPassword whereHelperstring
	Date        whereHelpertime_Time
}{
	ID:          whereHelperstring{field: "\"password_reset_requests\".\"id\""},
	UserID:      whereHelperstring{field: "\"password_reset_requests\".\"user_id\""},
	NewPassword: whereHelperstring{field: "\"password_reset_requests\".\"new_password\""},
	Date:        whereHelpertime_Time{field: "\"password_reset_requests\".\"date\""},
}

// PasswordResetRequestRels is where relationship names are stored.
var PasswordResetRequestRels = struct {
	User string
}{
	User: "User",
}

// passwordResetRequestR is where relationships are stored.
type passwordResetRequestR struct {
	User *User
}

// NewStruct creates a new relationship struct
func (*passwordResetRequestR) NewStruct() *passwordResetRequestR {
	return &passwordResetRequestR{}
}

// passwordResetRequestL is where Load methods for each relationship are stored.
type passwordResetRequestL struct{}

var (
	passwordResetRequestAllColumns            = []string{"id", "user_id", "new_password", "date"}
	passwordResetRequestColumnsWithoutDefault = []string{"id", "user_id", "new_password", "date"}
	passwordResetRequestColumnsWithDefault    = []string{}
	passwordResetRequestPrimaryKeyColumns     = []string{"id"}
)

type (
	// PasswordResetRequestSlice is an alias for a slice of pointers to PasswordResetRequest.
	// This should generally be used opposed to []PasswordResetRequest.
	PasswordResetRequestSlice []*PasswordResetRequest
	// PasswordResetRequestHook is the signature for custom PasswordResetRequest hook methods
	PasswordResetRequestHook func(context.Context, boil.ContextExecutor, *PasswordResetRequest) error

	passwordResetRequestQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	passwordResetRequestType                 = reflect.TypeOf(&PasswordResetRequest{})
	passwordResetRequestMapping              = queries.MakeStructMapping(passwordResetRequestType)
	passwordResetRequestPrimaryKeyMapping, _ = queries.BindMapping(passwordResetRequestType, passwordResetRequestMapping, passwordResetRequestPrimaryKeyColumns)
	passwordResetRequestInsertCacheMut       sync.RWMutex
	passwordResetRequestInsertCache          = make(map[string]insertCache)
	passwordResetRequestUpdateCacheMut       sync.RWMutex
	passwordResetRequestUpdateCache          = make(map[string]updateCache)
	passwordResetRequestUpsertCacheMut       sync.RWMutex
	passwordResetRequestUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var passwordResetRequestBeforeInsertHooks []PasswordResetRequestHook
var passwordResetRequestBeforeUpdateHooks []PasswordResetRequestHook
var passwordResetRequestBeforeDeleteHooks []PasswordResetRequestHook
var passwordResetRequestBeforeUpsertHooks []PasswordResetRequestHook

var passwordResetRequestAfterInsertHooks []PasswordResetRequestHook
var passwordResetRequestAfterSelectHooks []PasswordResetRequestHook
var passwordResetRequestAfterUpdateHooks []PasswordResetRequestHook
var passwordResetRequestAfterDeleteHooks []PasswordResetRequestHook
var passwordResetRequestAfterUpsertHooks []PasswordResetRequestHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *PasswordResetRequest) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range passwordResetRequestBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *PasswordResetRequest) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range passwordResetRequestBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *PasswordResetRequest) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range passwordResetRequestBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *PasswordResetRequest) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range passwordResetRequestBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *PasswordResetRequest) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range passwordResetRequestAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *PasswordResetRequest) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range passwordResetRequestAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *PasswordResetRequest) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range passwordResetRequestAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *PasswordResetRequest) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range passwordResetRequestAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *PasswordResetRequest) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range passwordResetRequestAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPasswordResetRequestHook registers your hook function for all future operations.
func AddPasswordResetRequestHook(hookPoint boil.HookPoint, passwordResetRequestHook PasswordResetRequestHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		passwordResetRequestBeforeInsertHooks = append(passwordResetRequestBeforeInsertHooks, passwordResetRequestHook)
	case boil.BeforeUpdateHook:
		passwordResetRequestBeforeUpdateHooks = append(passwordResetRequestBeforeUpdateHooks, passwordResetRequestHook)
	case boil.BeforeDeleteHook:
		passwordResetRequestBeforeDeleteHooks = append(passwordResetRequestBeforeDeleteHooks, passwordResetRequestHook)
	case boil.BeforeUpsertHook:
		passwordResetRequestBeforeUpsertHooks = append(passwordResetRequestBeforeUpsertHooks, passwordResetRequestHook)
	case boil.AfterInsertHook:
		passwordResetRequestAfterInsertHooks = append(passwordResetRequestAfterInsertHooks, passwordResetRequestHook)
	case boil.AfterSelectHook:
		passwordResetRequestAfterSelectHooks = append(passwordResetRequestAfterSelectHooks, passwordResetRequestHook)
	case boil.AfterUpdateHook:
		passwordResetRequestAfterUpdateHooks = append(passwordResetRequestAfterUpdateHooks, passwordResetRequestHook)
	case boil.AfterDeleteHook:
		passwordResetRequestAfterDeleteHooks = append(passwordResetRequestAfterDeleteHooks, passwordResetRequestHook)
	case boil.AfterUpsertHook:
		passwordResetRequestAfterUpsertHooks = append(passwordResetRequestAfterUpsertHooks, passwordResetRequestHook)
	}
}

// One returns a single passwordResetRequest record from the query.
func (q passwordResetRequestQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PasswordResetRequest, error) {
	o := &PasswordResetRequest{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for password_reset_requests")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all PasswordResetRequest records from the query.
func (q passwordResetRequestQuery) All(ctx context.Context, exec boil.ContextExecutor) (PasswordResetRequestSlice, error) {
	var o []*PasswordResetRequest

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to PasswordResetRequest slice")
	}

	if len(passwordResetRequestAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all PasswordResetRequest records in the query.
func (q passwordResetRequestQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count password_reset_requests rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q passwordResetRequestQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if password_reset_requests exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *PasswordResetRequest) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	return query
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (passwordResetRequestL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybePasswordResetRequest interface{}, mods queries.Applicator) error {
	var slice []*PasswordResetRequest
	var object *PasswordResetRequest

	if singular {
		object = maybePasswordResetRequest.(*PasswordResetRequest)
	} else {
		slice = *maybePasswordResetRequest.(*[]*PasswordResetRequest)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &passwordResetRequestR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &passwordResetRequestR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`users`), qm.WhereIn(`users.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(passwordResetRequestAfterSelectHooks) != 0 {
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
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.PasswordResetRequests = append(foreign.R.PasswordResetRequests, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.PasswordResetRequests = append(foreign.R.PasswordResetRequests, local)
				break
			}
		}
	}

	return nil
}

// SetUser of the passwordResetRequest to the related item.
// Sets o.R.User to related.
// Adds o to related.R.PasswordResetRequests.
func (o *PasswordResetRequest) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"password_reset_requests\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, passwordResetRequestPrimaryKeyColumns),
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

	o.UserID = related.ID
	if o.R == nil {
		o.R = &passwordResetRequestR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			PasswordResetRequests: PasswordResetRequestSlice{o},
		}
	} else {
		related.R.PasswordResetRequests = append(related.R.PasswordResetRequests, o)
	}

	return nil
}

// PasswordResetRequests retrieves all the records using an executor.
func PasswordResetRequests(mods ...qm.QueryMod) passwordResetRequestQuery {
	mods = append(mods, qm.From("\"password_reset_requests\""))
	return passwordResetRequestQuery{NewQuery(mods...)}
}

// FindPasswordResetRequest retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPasswordResetRequest(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*PasswordResetRequest, error) {
	passwordResetRequestObj := &PasswordResetRequest{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"password_reset_requests\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, passwordResetRequestObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from password_reset_requests")
	}

	return passwordResetRequestObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PasswordResetRequest) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no password_reset_requests provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(passwordResetRequestColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	passwordResetRequestInsertCacheMut.RLock()
	cache, cached := passwordResetRequestInsertCache[key]
	passwordResetRequestInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			passwordResetRequestAllColumns,
			passwordResetRequestColumnsWithDefault,
			passwordResetRequestColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(passwordResetRequestType, passwordResetRequestMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(passwordResetRequestType, passwordResetRequestMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"password_reset_requests\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"password_reset_requests\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into password_reset_requests")
	}

	if !cached {
		passwordResetRequestInsertCacheMut.Lock()
		passwordResetRequestInsertCache[key] = cache
		passwordResetRequestInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the PasswordResetRequest.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PasswordResetRequest) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	passwordResetRequestUpdateCacheMut.RLock()
	cache, cached := passwordResetRequestUpdateCache[key]
	passwordResetRequestUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			passwordResetRequestAllColumns,
			passwordResetRequestPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update password_reset_requests, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"password_reset_requests\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, passwordResetRequestPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(passwordResetRequestType, passwordResetRequestMapping, append(wl, passwordResetRequestPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update password_reset_requests row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for password_reset_requests")
	}

	if !cached {
		passwordResetRequestUpdateCacheMut.Lock()
		passwordResetRequestUpdateCache[key] = cache
		passwordResetRequestUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q passwordResetRequestQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for password_reset_requests")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for password_reset_requests")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PasswordResetRequestSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), passwordResetRequestPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"password_reset_requests\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, passwordResetRequestPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in passwordResetRequest slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all passwordResetRequest")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *PasswordResetRequest) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no password_reset_requests provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(passwordResetRequestColumnsWithDefault, o)

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

	passwordResetRequestUpsertCacheMut.RLock()
	cache, cached := passwordResetRequestUpsertCache[key]
	passwordResetRequestUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			passwordResetRequestAllColumns,
			passwordResetRequestColumnsWithDefault,
			passwordResetRequestColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			passwordResetRequestAllColumns,
			passwordResetRequestPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert password_reset_requests, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(passwordResetRequestPrimaryKeyColumns))
			copy(conflict, passwordResetRequestPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"password_reset_requests\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(passwordResetRequestType, passwordResetRequestMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(passwordResetRequestType, passwordResetRequestMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert password_reset_requests")
	}

	if !cached {
		passwordResetRequestUpsertCacheMut.Lock()
		passwordResetRequestUpsertCache[key] = cache
		passwordResetRequestUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single PasswordResetRequest record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PasswordResetRequest) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no PasswordResetRequest provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), passwordResetRequestPrimaryKeyMapping)
	sql := "DELETE FROM \"password_reset_requests\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from password_reset_requests")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for password_reset_requests")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q passwordResetRequestQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no passwordResetRequestQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from password_reset_requests")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for password_reset_requests")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PasswordResetRequestSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(passwordResetRequestBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), passwordResetRequestPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"password_reset_requests\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, passwordResetRequestPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from passwordResetRequest slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for password_reset_requests")
	}

	if len(passwordResetRequestAfterDeleteHooks) != 0 {
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
func (o *PasswordResetRequest) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPasswordResetRequest(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PasswordResetRequestSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PasswordResetRequestSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), passwordResetRequestPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"password_reset_requests\".* FROM \"password_reset_requests\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, passwordResetRequestPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PasswordResetRequestSlice")
	}

	*o = slice

	return nil
}

// PasswordResetRequestExists checks if the PasswordResetRequest row exists.
func PasswordResetRequestExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"password_reset_requests\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if password_reset_requests exists")
	}

	return exists, nil
}
