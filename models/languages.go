package models

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/queries"
	"github.com/vattle/sqlboiler/queries/qm"
	"github.com/vattle/sqlboiler/strmangle"
)

// Language is an object representing the database table.
type Language struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Language  string    `boil:"language" json:"language" toml:"language" yaml:"language"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *languageR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L languageL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// languageR is where relationships are stored.
type languageR struct {
	Pilots PilotSlice
}

// languageL is where Load methods for each relationship are stored.
type languageL struct{}

var (
	languageColumns               = []string{"id", "language", "created_at", "updated_at"}
	languageColumnsWithoutDefault = []string{"language", "created_at", "updated_at"}
	languageColumnsWithDefault    = []string{"id"}
	languagePrimaryKeyColumns     = []string{"id"}
)

type (
	// LanguageSlice is an alias for a slice of pointers to Language.
	// This should generally be used opposed to []Language.
	LanguageSlice []*Language
	// LanguageHook is the signature for custom Language hook methods
	LanguageHook func(boil.Executor, *Language) error

	languageQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	languageType                 = reflect.TypeOf(&Language{})
	languageMapping              = queries.MakeStructMapping(languageType)
	languagePrimaryKeyMapping, _ = queries.BindMapping(languageType, languageMapping, languagePrimaryKeyColumns)
	languageInsertCacheMut       sync.RWMutex
	languageInsertCache          = make(map[string]insertCache)
	languageUpdateCacheMut       sync.RWMutex
	languageUpdateCache          = make(map[string]updateCache)
	languageUpsertCacheMut       sync.RWMutex
	languageUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var languageBeforeInsertHooks []LanguageHook
var languageBeforeUpdateHooks []LanguageHook
var languageBeforeDeleteHooks []LanguageHook
var languageBeforeUpsertHooks []LanguageHook

var languageAfterInsertHooks []LanguageHook
var languageAfterSelectHooks []LanguageHook
var languageAfterUpdateHooks []LanguageHook
var languageAfterDeleteHooks []LanguageHook
var languageAfterUpsertHooks []LanguageHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Language) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range languageBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Language) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range languageBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Language) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range languageBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Language) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range languageBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Language) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range languageAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Language) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range languageAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Language) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range languageAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Language) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range languageAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Language) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range languageAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddLanguageHook registers your hook function for all future operations.
func AddLanguageHook(hookPoint boil.HookPoint, languageHook LanguageHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		languageBeforeInsertHooks = append(languageBeforeInsertHooks, languageHook)
	case boil.BeforeUpdateHook:
		languageBeforeUpdateHooks = append(languageBeforeUpdateHooks, languageHook)
	case boil.BeforeDeleteHook:
		languageBeforeDeleteHooks = append(languageBeforeDeleteHooks, languageHook)
	case boil.BeforeUpsertHook:
		languageBeforeUpsertHooks = append(languageBeforeUpsertHooks, languageHook)
	case boil.AfterInsertHook:
		languageAfterInsertHooks = append(languageAfterInsertHooks, languageHook)
	case boil.AfterSelectHook:
		languageAfterSelectHooks = append(languageAfterSelectHooks, languageHook)
	case boil.AfterUpdateHook:
		languageAfterUpdateHooks = append(languageAfterUpdateHooks, languageHook)
	case boil.AfterDeleteHook:
		languageAfterDeleteHooks = append(languageAfterDeleteHooks, languageHook)
	case boil.AfterUpsertHook:
		languageAfterUpsertHooks = append(languageAfterUpsertHooks, languageHook)
	}
}

// OneP returns a single language record from the query, and panics on error.
func (q languageQuery) OneP() *Language {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single language record from the query.
func (q languageQuery) One() (*Language, error) {
	o := &Language{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for languages")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Language records from the query, and panics on error.
func (q languageQuery) AllP() LanguageSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Language records from the query.
func (q languageQuery) All() (LanguageSlice, error) {
	var o LanguageSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Language slice")
	}

	if len(languageAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Language records in the query, and panics on error.
func (q languageQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Language records in the query.
func (q languageQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count languages rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q languageQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q languageQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if languages exists")
	}

	return count > 0, nil
}

// PilotsG retrieves all the pilot's pilots.
func (o *Language) PilotsG(mods ...qm.QueryMod) pilotQuery {
	return o.Pilots(boil.GetDB(), mods...)
}

// Pilots retrieves all the pilot's pilots with an executor.
func (o *Language) Pilots(exec boil.Executor, mods ...qm.QueryMod) pilotQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.InnerJoin("\"pilot_languages\" as \"b\" on \"a\".\"id\" = \"b\".\"pilot_id\""),
		qm.Where("\"b\".\"language_id\"=?", o.ID),
	)

	query := Pilots(exec, queryMods...)
	queries.SetFrom(query.Query, "\"pilots\" as \"a\"")
	return query
}

// LoadPilots allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (languageL) LoadPilots(e boil.Executor, singular bool, maybeLanguage interface{}) error {
	var slice []*Language
	var object *Language

	count := 1
	if singular {
		object = maybeLanguage.(*Language)
	} else {
		slice = *maybeLanguage.(*LanguageSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &languageR{}
		}
		args[0] = object.ID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &languageR{}
			}
			args[i] = obj.ID
		}
	}

	query := fmt.Sprintf(
		"select \"a\".*, \"b\".\"language_id\" from \"pilots\" as \"a\" inner join \"pilot_languages\" as \"b\" on \"a\".\"id\" = \"b\".\"pilot_id\" where \"b\".\"language_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load pilots")
	}
	defer results.Close()

	var resultSlice []*Pilot

	var localJoinCols []int
	for results.Next() {
		one := new(Pilot)
		var localJoinCol int

		err = results.Scan(&one.ID, &one.Name, &one.Hobbies, &one.CreatedAt, &one.UpdatedAt, &localJoinCol)
		if err = results.Err(); err != nil {
			return errors.Wrap(err, "failed to plebian-bind eager loaded slice pilots")
		}

		resultSlice = append(resultSlice, one)
		localJoinCols = append(localJoinCols, localJoinCol)
	}

	if err = results.Err(); err != nil {
		return errors.Wrap(err, "failed to plebian-bind eager loaded slice pilots")
	}

	if len(pilotAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Pilots = resultSlice
		return nil
	}

	for i, foreign := range resultSlice {
		localJoinCol := localJoinCols[i]
		for _, local := range slice {
			if local.ID == localJoinCol {
				local.R.Pilots = append(local.R.Pilots, foreign)
				break
			}
		}
	}

	return nil
}

// AddPilotsG adds the given related objects to the existing relationships
// of the language, optionally inserting them as new records.
// Appends related to o.R.Pilots.
// Sets related.R.Languages appropriately.
// Uses the global database handle.
func (o *Language) AddPilotsG(insert bool, related ...*Pilot) error {
	return o.AddPilots(boil.GetDB(), insert, related...)
}

// AddPilotsP adds the given related objects to the existing relationships
// of the language, optionally inserting them as new records.
// Appends related to o.R.Pilots.
// Sets related.R.Languages appropriately.
// Panics on error.
func (o *Language) AddPilotsP(exec boil.Executor, insert bool, related ...*Pilot) {
	if err := o.AddPilots(exec, insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddPilotsGP adds the given related objects to the existing relationships
// of the language, optionally inserting them as new records.
// Appends related to o.R.Pilots.
// Sets related.R.Languages appropriately.
// Uses the global database handle and panics on error.
func (o *Language) AddPilotsGP(insert bool, related ...*Pilot) {
	if err := o.AddPilots(boil.GetDB(), insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddPilots adds the given related objects to the existing relationships
// of the language, optionally inserting them as new records.
// Appends related to o.R.Pilots.
// Sets related.R.Languages appropriately.
func (o *Language) AddPilots(exec boil.Executor, insert bool, related ...*Pilot) error {
	var err error
	for _, rel := range related {
		if insert {
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		}
	}

	for _, rel := range related {
		query := "insert into \"pilot_languages\" (\"language_id\", \"pilot_id\") values ($1, $2)"
		values := []interface{}{o.ID, rel.ID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, query)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		_, err = exec.Exec(query, values...)
		if err != nil {
			return errors.Wrap(err, "failed to insert into join table")
		}
	}
	if o.R == nil {
		o.R = &languageR{
			Pilots: related,
		}
	} else {
		o.R.Pilots = append(o.R.Pilots, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &pilotR{
				Languages: LanguageSlice{o},
			}
		} else {
			rel.R.Languages = append(rel.R.Languages, o)
		}
	}
	return nil
}

// SetPilotsG removes all previously related items of the
// language replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Languages's Pilots accordingly.
// Replaces o.R.Pilots with related.
// Sets related.R.Languages's Pilots accordingly.
// Uses the global database handle.
func (o *Language) SetPilotsG(insert bool, related ...*Pilot) error {
	return o.SetPilots(boil.GetDB(), insert, related...)
}

// SetPilotsP removes all previously related items of the
// language replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Languages's Pilots accordingly.
// Replaces o.R.Pilots with related.
// Sets related.R.Languages's Pilots accordingly.
// Panics on error.
func (o *Language) SetPilotsP(exec boil.Executor, insert bool, related ...*Pilot) {
	if err := o.SetPilots(exec, insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetPilotsGP removes all previously related items of the
// language replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Languages's Pilots accordingly.
// Replaces o.R.Pilots with related.
// Sets related.R.Languages's Pilots accordingly.
// Uses the global database handle and panics on error.
func (o *Language) SetPilotsGP(insert bool, related ...*Pilot) {
	if err := o.SetPilots(boil.GetDB(), insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetPilots removes all previously related items of the
// language replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Languages's Pilots accordingly.
// Replaces o.R.Pilots with related.
// Sets related.R.Languages's Pilots accordingly.
func (o *Language) SetPilots(exec boil.Executor, insert bool, related ...*Pilot) error {
	query := "delete from \"pilot_languages\" where \"language_id\" = $1"
	values := []interface{}{o.ID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	removePilotsFromLanguagesSlice(o, related)
	o.R.Pilots = nil
	return o.AddPilots(exec, insert, related...)
}

// RemovePilotsG relationships from objects passed in.
// Removes related items from R.Pilots (uses pointer comparison, removal does not keep order)
// Sets related.R.Languages.
// Uses the global database handle.
func (o *Language) RemovePilotsG(related ...*Pilot) error {
	return o.RemovePilots(boil.GetDB(), related...)
}

// RemovePilotsP relationships from objects passed in.
// Removes related items from R.Pilots (uses pointer comparison, removal does not keep order)
// Sets related.R.Languages.
// Panics on error.
func (o *Language) RemovePilotsP(exec boil.Executor, related ...*Pilot) {
	if err := o.RemovePilots(exec, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// RemovePilotsGP relationships from objects passed in.
// Removes related items from R.Pilots (uses pointer comparison, removal does not keep order)
// Sets related.R.Languages.
// Uses the global database handle and panics on error.
func (o *Language) RemovePilotsGP(related ...*Pilot) {
	if err := o.RemovePilots(boil.GetDB(), related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// RemovePilots relationships from objects passed in.
// Removes related items from R.Pilots (uses pointer comparison, removal does not keep order)
// Sets related.R.Languages.
func (o *Language) RemovePilots(exec boil.Executor, related ...*Pilot) error {
	var err error
	query := fmt.Sprintf(
		"delete from \"pilot_languages\" where \"language_id\" = $1 and \"pilot_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, len(related), 1, 1),
	)
	values := []interface{}{o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}
	removePilotsFromLanguagesSlice(o, related)
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Pilots {
			if rel != ri {
				continue
			}

			ln := len(o.R.Pilots)
			if ln > 1 && i < ln-1 {
				o.R.Pilots[i] = o.R.Pilots[ln-1]
			}
			o.R.Pilots = o.R.Pilots[:ln-1]
			break
		}
	}

	return nil
}

func removePilotsFromLanguagesSlice(o *Language, related []*Pilot) {
	for _, rel := range related {
		if rel.R == nil {
			continue
		}
		for i, ri := range rel.R.Languages {
			if o.ID != ri.ID {
				continue
			}

			ln := len(rel.R.Languages)
			if ln > 1 && i < ln-1 {
				rel.R.Languages[i] = rel.R.Languages[ln-1]
			}
			rel.R.Languages = rel.R.Languages[:ln-1]
			break
		}
	}
}

// LanguagesG retrieves all records.
func LanguagesG(mods ...qm.QueryMod) languageQuery {
	return Languages(boil.GetDB(), mods...)
}

// Languages retrieves all the records using an executor.
func Languages(exec boil.Executor, mods ...qm.QueryMod) languageQuery {
	mods = append(mods, qm.From("\"languages\""))
	return languageQuery{NewQuery(exec, mods...)}
}

// FindLanguageG retrieves a single record by ID.
func FindLanguageG(id int, selectCols ...string) (*Language, error) {
	return FindLanguage(boil.GetDB(), id, selectCols...)
}

// FindLanguageGP retrieves a single record by ID, and panics on error.
func FindLanguageGP(id int, selectCols ...string) *Language {
	retobj, err := FindLanguage(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindLanguage retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindLanguage(exec boil.Executor, id int, selectCols ...string) (*Language, error) {
	languageObj := &Language{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"languages\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(languageObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from languages")
	}

	return languageObj, nil
}

// FindLanguageP retrieves a single record by ID with an executor, and panics on error.
func FindLanguageP(exec boil.Executor, id int, selectCols ...string) *Language {
	retobj, err := FindLanguage(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Language) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Language) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Language) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Language) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no languages provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	if o.UpdatedAt.IsZero() {
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(languageColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	languageInsertCacheMut.RLock()
	cache, cached := languageInsertCache[key]
	languageInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			languageColumns,
			languageColumnsWithDefault,
			languageColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(languageType, languageMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(languageType, languageMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"languages\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

		if len(cache.retMapping) != 0 {
			cache.query += fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into languages")
	}

	if !cached {
		languageInsertCacheMut.Lock()
		languageInsertCache[key] = cache
		languageInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Language record. See Update for
// whitelist behavior description.
func (o *Language) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Language record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Language) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Language, and panics on error.
// See Update for whitelist behavior description.
func (o *Language) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Language.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Language) Update(exec boil.Executor, whitelist ...string) error {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	languageUpdateCacheMut.RLock()
	cache, cached := languageUpdateCache[key]
	languageUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(languageColumns, languagePrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update languages, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"languages\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, languagePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(languageType, languageMapping, append(wl, languagePrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update languages row")
	}

	if !cached {
		languageUpdateCacheMut.Lock()
		languageUpdateCache[key] = cache
		languageUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q languageQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q languageQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for languages")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o LanguageSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o LanguageSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o LanguageSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o LanguageSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), languagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"languages\" SET %s WHERE (\"id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(languagePrimaryKeyColumns), len(colNames)+1, len(languagePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in language slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Language) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Language) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Language) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Language) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no languages provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(languageColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
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
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	languageUpsertCacheMut.RLock()
	cache, cached := languageUpsertCache[key]
	languageUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			languageColumns,
			languageColumnsWithDefault,
			languageColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			languageColumns,
			languagePrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert languages, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(languagePrimaryKeyColumns))
			copy(conflict, languagePrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"languages\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(languageType, languageMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(languageType, languageMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert languages")
	}

	if !cached {
		languageUpsertCacheMut.Lock()
		languageUpsertCache[key] = cache
		languageUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Language record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Language) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Language record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Language) DeleteG() error {
	if o == nil {
		return errors.New("models: no Language provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Language record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Language) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Language record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Language) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Language provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), languagePrimaryKeyMapping)
	sql := "DELETE FROM \"languages\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from languages")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q languageQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q languageQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no languageQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from languages")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o LanguageSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o LanguageSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Language slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o LanguageSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o LanguageSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Language slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(languageBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), languagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"languages\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, languagePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(languagePrimaryKeyColumns), 1, len(languagePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from language slice")
	}

	if len(languageAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Language) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Language) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Language) ReloadG() error {
	if o == nil {
		return errors.New("models: no Language provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Language) Reload(exec boil.Executor) error {
	ret, err := FindLanguage(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *LanguageSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *LanguageSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LanguageSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty LanguageSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LanguageSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	languages := LanguageSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), languagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"languages\".* FROM \"languages\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, languagePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(languagePrimaryKeyColumns), 1, len(languagePrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&languages)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in LanguageSlice")
	}

	*o = languages

	return nil
}

// LanguageExists checks if the Language row exists.
func LanguageExists(exec boil.Executor, id int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"languages\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if languages exists")
	}

	return exists, nil
}

// LanguageExistsG checks if the Language row exists.
func LanguageExistsG(id int) (bool, error) {
	return LanguageExists(boil.GetDB(), id)
}

// LanguageExistsGP checks if the Language row exists. Panics on error.
func LanguageExistsGP(id int) bool {
	e, err := LanguageExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// LanguageExistsP checks if the Language row exists. Panics on error.
func LanguageExistsP(exec boil.Executor, id int) bool {
	e, err := LanguageExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
