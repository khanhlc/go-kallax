// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

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
	"gopkg.in/nullbio/null.v6"
)

// Pet is an object representing the database table.
type Pet struct {
	ID       int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name     null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Kind     null.String `boil:"kind" json:"kind,omitempty" toml:"kind" yaml:"kind,omitempty"`
	PersonID null.Int    `boil:"person_id" json:"person_id,omitempty" toml:"person_id" yaml:"person_id,omitempty"`

	R *petR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L petL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// petR is where relationships are stored.
type petR struct {
	Person *Person
}

// petL is where Load methods for each relationship are stored.
type petL struct{}

var (
	petColumns               = []string{"id", "name", "kind", "person_id"}
	petColumnsWithoutDefault = []string{"name", "kind", "person_id"}
	petColumnsWithDefault    = []string{"id"}
	petPrimaryKeyColumns     = []string{"id"}
)

type (
	// PetSlice is an alias for a slice of pointers to Pet.
	// This should generally be used opposed to []Pet.
	PetSlice []*Pet

	petQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	petType                 = reflect.TypeOf(&Pet{})
	petMapping              = queries.MakeStructMapping(petType)
	petPrimaryKeyMapping, _ = queries.BindMapping(petType, petMapping, petPrimaryKeyColumns)
	petInsertCacheMut       sync.RWMutex
	petInsertCache          = make(map[string]insertCache)
	petUpdateCacheMut       sync.RWMutex
	petUpdateCache          = make(map[string]updateCache)
	petUpsertCacheMut       sync.RWMutex
	petUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single pet record from the query, and panics on error.
func (q petQuery) OneP() *Pet {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single pet record from the query.
func (q petQuery) One() (*Pet, error) {
	o := &Pet{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for pets")
	}

	return o, nil
}

// AllP returns all Pet records from the query, and panics on error.
func (q petQuery) AllP() PetSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Pet records from the query.
func (q petQuery) All() (PetSlice, error) {
	var o []*Pet

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Pet slice")
	}

	return o, nil
}

// CountP returns the count of all Pet records in the query, and panics on error.
func (q petQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Pet records in the query.
func (q petQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count pets rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q petQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q petQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if pets exists")
	}

	return count > 0, nil
}

// PersonG pointed to by the foreign key.
func (o *Pet) PersonG(mods ...qm.QueryMod) personQuery {
	return o.Person(boil.GetDB(), mods...)
}

// Person pointed to by the foreign key.
func (o *Pet) Person(exec boil.Executor, mods ...qm.QueryMod) personQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.PersonID),
	}

	queryMods = append(queryMods, mods...)

	query := People(exec, queryMods...)
	queries.SetFrom(query.Query, "\"people\"")

	return query
} // LoadPerson allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (petL) LoadPerson(e boil.Executor, singular bool, maybePet interface{}) error {
	var slice []*Pet
	var object *Pet

	count := 1
	if singular {
		object = maybePet.(*Pet)
	} else {
		slice = *maybePet.(*[]*Pet)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &petR{}
		}
		args[0] = object.PersonID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &petR{}
			}
			args[i] = obj.PersonID
		}
	}

	query := fmt.Sprintf(
		"select * from \"people\" where \"id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Person")
	}
	defer results.Close()

	var resultSlice []*Person
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Person")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		object.R.Person = resultSlice[0]
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.PersonID.Int == foreign.ID {
				local.R.Person = foreign
				break
			}
		}
	}

	return nil
}

// SetPersonG of the pet to the related item.
// Sets o.R.Person to related.
// Adds o to related.R.Pets.
// Uses the global database handle.
func (o *Pet) SetPersonG(insert bool, related *Person) error {
	return o.SetPerson(boil.GetDB(), insert, related)
}

// SetPersonP of the pet to the related item.
// Sets o.R.Person to related.
// Adds o to related.R.Pets.
// Panics on error.
func (o *Pet) SetPersonP(exec boil.Executor, insert bool, related *Person) {
	if err := o.SetPerson(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetPersonGP of the pet to the related item.
// Sets o.R.Person to related.
// Adds o to related.R.Pets.
// Uses the global database handle and panics on error.
func (o *Pet) SetPersonGP(insert bool, related *Person) {
	if err := o.SetPerson(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetPerson of the pet to the related item.
// Sets o.R.Person to related.
// Adds o to related.R.Pets.
func (o *Pet) SetPerson(exec boil.Executor, insert bool, related *Person) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"pets\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"person_id"}),
		strmangle.WhereClause("\"", "\"", 2, petPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PersonID.Int = related.ID
	o.PersonID.Valid = true

	if o.R == nil {
		o.R = &petR{
			Person: related,
		}
	} else {
		o.R.Person = related
	}

	if related.R == nil {
		related.R = &personR{
			Pets: PetSlice{o},
		}
	} else {
		related.R.Pets = append(related.R.Pets, o)
	}

	return nil
}

// RemovePersonG relationship.
// Sets o.R.Person to nil.
// Removes o from all passed in related items' relationships struct (Optional).
// Uses the global database handle.
func (o *Pet) RemovePersonG(related *Person) error {
	return o.RemovePerson(boil.GetDB(), related)
}

// RemovePersonP relationship.
// Sets o.R.Person to nil.
// Removes o from all passed in related items' relationships struct (Optional).
// Panics on error.
func (o *Pet) RemovePersonP(exec boil.Executor, related *Person) {
	if err := o.RemovePerson(exec, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// RemovePersonGP relationship.
// Sets o.R.Person to nil.
// Removes o from all passed in related items' relationships struct (Optional).
// Uses the global database handle and panics on error.
func (o *Pet) RemovePersonGP(related *Person) {
	if err := o.RemovePerson(boil.GetDB(), related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// RemovePerson relationship.
// Sets o.R.Person to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Pet) RemovePerson(exec boil.Executor, related *Person) error {
	var err error

	o.PersonID.Valid = false
	if err = o.Update(exec, "person_id"); err != nil {
		o.PersonID.Valid = true
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Person = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.Pets {
		if o.PersonID.Int != ri.PersonID.Int {
			continue
		}

		ln := len(related.R.Pets)
		if ln > 1 && i < ln-1 {
			related.R.Pets[i] = related.R.Pets[ln-1]
		}
		related.R.Pets = related.R.Pets[:ln-1]
		break
	}
	return nil
}

// PetsG retrieves all records.
func PetsG(mods ...qm.QueryMod) petQuery {
	return Pets(boil.GetDB(), mods...)
}

// Pets retrieves all the records using an executor.
func Pets(exec boil.Executor, mods ...qm.QueryMod) petQuery {
	mods = append(mods, qm.From("\"pets\""))
	return petQuery{NewQuery(exec, mods...)}
}

// FindPetG retrieves a single record by ID.
func FindPetG(id int, selectCols ...string) (*Pet, error) {
	return FindPet(boil.GetDB(), id, selectCols...)
}

// FindPetGP retrieves a single record by ID, and panics on error.
func FindPetGP(id int, selectCols ...string) *Pet {
	retobj, err := FindPet(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindPet retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPet(exec boil.Executor, id int, selectCols ...string) (*Pet, error) {
	petObj := &Pet{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"pets\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(petObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from pets")
	}

	return petObj, nil
}

// FindPetP retrieves a single record by ID with an executor, and panics on error.
func FindPetP(exec boil.Executor, id int, selectCols ...string) *Pet {
	retobj, err := FindPet(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Pet) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Pet) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Pet) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Pet) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no pets provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(petColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	petInsertCacheMut.RLock()
	cache, cached := petInsertCache[key]
	petInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			petColumns,
			petColumnsWithDefault,
			petColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(petType, petMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(petType, petMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"pets\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"pets\" DEFAULT VALUES"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		if len(wl) != 0 {
			cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
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
		return errors.Wrap(err, "models: unable to insert into pets")
	}

	if !cached {
		petInsertCacheMut.Lock()
		petInsertCache[key] = cache
		petInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single Pet record. See Update for
// whitelist behavior description.
func (o *Pet) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Pet record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Pet) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Pet, and panics on error.
// See Update for whitelist behavior description.
func (o *Pet) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Pet.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Pet) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	petUpdateCacheMut.RLock()
	cache, cached := petUpdateCache[key]
	petUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			petColumns,
			petPrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update pets, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"pets\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, petPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(petType, petMapping, append(wl, petPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update pets row")
	}

	if !cached {
		petUpdateCacheMut.Lock()
		petUpdateCache[key] = cache
		petUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q petQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q petQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for pets")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o PetSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o PetSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o PetSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PetSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), petPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"pets\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, petPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in pet slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Pet) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Pet) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Pet) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Pet) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no pets provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(petColumnsWithDefault, o)

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

	petUpsertCacheMut.RLock()
	cache, cached := petUpsertCache[key]
	petUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			petColumns,
			petColumnsWithDefault,
			petColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			petColumns,
			petPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert pets, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(petPrimaryKeyColumns))
			copy(conflict, petPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"pets\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(petType, petMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(petType, petMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert pets")
	}

	if !cached {
		petUpsertCacheMut.Lock()
		petUpsertCache[key] = cache
		petUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single Pet record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Pet) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Pet record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Pet) DeleteG() error {
	if o == nil {
		return errors.New("models: no Pet provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Pet record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Pet) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Pet record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Pet) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Pet provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), petPrimaryKeyMapping)
	sql := "DELETE FROM \"pets\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from pets")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q petQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q petQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no petQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from pets")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o PetSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o PetSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Pet slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o PetSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PetSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Pet slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), petPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"pets\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, petPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from pet slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Pet) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Pet) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Pet) ReloadG() error {
	if o == nil {
		return errors.New("models: no Pet provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Pet) Reload(exec boil.Executor) error {
	ret, err := FindPet(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *PetSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *PetSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PetSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty PetSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PetSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	pets := PetSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), petPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"pets\".* FROM \"pets\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, petPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&pets)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PetSlice")
	}

	*o = pets

	return nil
}

// PetExists checks if the Pet row exists.
func PetExists(exec boil.Executor, id int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"pets\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if pets exists")
	}

	return exists, nil
}

// PetExistsG checks if the Pet row exists.
func PetExistsG(id int) (bool, error) {
	return PetExists(boil.GetDB(), id)
}

// PetExistsGP checks if the Pet row exists. Panics on error.
func PetExistsGP(id int) bool {
	e, err := PetExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// PetExistsP checks if the Pet row exists. Panics on error.
func PetExistsP(exec boil.Executor, id int) bool {
	e, err := PetExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
