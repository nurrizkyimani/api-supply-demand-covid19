// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testRequestItems(t *testing.T) {
	t.Parallel()

	query := RequestItems()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRequestItemsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestItem{}
	if err = randomize.Struct(seed, o, requestItemDBTypes, true, requestItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RequestItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRequestItemsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestItem{}
	if err = randomize.Struct(seed, o, requestItemDBTypes, true, requestItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := RequestItems().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RequestItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRequestItemsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestItem{}
	if err = randomize.Struct(seed, o, requestItemDBTypes, true, requestItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RequestItemSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RequestItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRequestItemsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestItem{}
	if err = randomize.Struct(seed, o, requestItemDBTypes, true, requestItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RequestItemExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if RequestItem exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RequestItemExists to return true, but got false.")
	}
}

func testRequestItemsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestItem{}
	if err = randomize.Struct(seed, o, requestItemDBTypes, true, requestItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	requestItemFound, err := FindRequestItem(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if requestItemFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRequestItemsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestItem{}
	if err = randomize.Struct(seed, o, requestItemDBTypes, true, requestItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = RequestItems().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRequestItemsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestItem{}
	if err = randomize.Struct(seed, o, requestItemDBTypes, true, requestItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := RequestItems().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRequestItemsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	requestItemOne := &RequestItem{}
	requestItemTwo := &RequestItem{}
	if err = randomize.Struct(seed, requestItemOne, requestItemDBTypes, false, requestItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItem struct: %s", err)
	}
	if err = randomize.Struct(seed, requestItemTwo, requestItemDBTypes, false, requestItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = requestItemOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = requestItemTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RequestItems().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRequestItemsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	requestItemOne := &RequestItem{}
	requestItemTwo := &RequestItem{}
	if err = randomize.Struct(seed, requestItemOne, requestItemDBTypes, false, requestItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItem struct: %s", err)
	}
	if err = randomize.Struct(seed, requestItemTwo, requestItemDBTypes, false, requestItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = requestItemOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = requestItemTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RequestItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func requestItemBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *RequestItem) error {
	*o = RequestItem{}
	return nil
}

func requestItemAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *RequestItem) error {
	*o = RequestItem{}
	return nil
}

func requestItemAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *RequestItem) error {
	*o = RequestItem{}
	return nil
}

func requestItemBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RequestItem) error {
	*o = RequestItem{}
	return nil
}

func requestItemAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RequestItem) error {
	*o = RequestItem{}
	return nil
}

func requestItemBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RequestItem) error {
	*o = RequestItem{}
	return nil
}

func requestItemAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RequestItem) error {
	*o = RequestItem{}
	return nil
}

func requestItemBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RequestItem) error {
	*o = RequestItem{}
	return nil
}

func requestItemAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RequestItem) error {
	*o = RequestItem{}
	return nil
}

func testRequestItemsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &RequestItem{}
	o := &RequestItem{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, requestItemDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RequestItem object: %s", err)
	}

	AddRequestItemHook(boil.BeforeInsertHook, requestItemBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	requestItemBeforeInsertHooks = []RequestItemHook{}

	AddRequestItemHook(boil.AfterInsertHook, requestItemAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	requestItemAfterInsertHooks = []RequestItemHook{}

	AddRequestItemHook(boil.AfterSelectHook, requestItemAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	requestItemAfterSelectHooks = []RequestItemHook{}

	AddRequestItemHook(boil.BeforeUpdateHook, requestItemBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	requestItemBeforeUpdateHooks = []RequestItemHook{}

	AddRequestItemHook(boil.AfterUpdateHook, requestItemAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	requestItemAfterUpdateHooks = []RequestItemHook{}

	AddRequestItemHook(boil.BeforeDeleteHook, requestItemBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	requestItemBeforeDeleteHooks = []RequestItemHook{}

	AddRequestItemHook(boil.AfterDeleteHook, requestItemAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	requestItemAfterDeleteHooks = []RequestItemHook{}

	AddRequestItemHook(boil.BeforeUpsertHook, requestItemBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	requestItemBeforeUpsertHooks = []RequestItemHook{}

	AddRequestItemHook(boil.AfterUpsertHook, requestItemAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	requestItemAfterUpsertHooks = []RequestItemHook{}
}

func testRequestItemsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestItem{}
	if err = randomize.Struct(seed, o, requestItemDBTypes, true, requestItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RequestItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRequestItemsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestItem{}
	if err = randomize.Struct(seed, o, requestItemDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RequestItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(requestItemColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := RequestItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRequestItemToOneItemUsingItem(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RequestItem
	var foreign Item

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, requestItemDBTypes, false, requestItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItem struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, itemDBTypes, false, itemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Item struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ItemID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Item().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := RequestItemSlice{&local}
	if err = local.L.LoadItem(ctx, tx, false, (*[]*RequestItem)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Item == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Item = nil
	if err = local.L.LoadItem(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Item == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testRequestItemToOneRequestUsingRequest(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RequestItem
	var foreign Request

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, requestItemDBTypes, false, requestItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItem struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, requestDBTypes, false, requestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Request struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.RequestID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Request().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := RequestItemSlice{&local}
	if err = local.L.LoadRequest(ctx, tx, false, (*[]*RequestItem)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Request == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Request = nil
	if err = local.L.LoadRequest(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Request == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testRequestItemToOneUnitUsingUnit(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RequestItem
	var foreign Unit

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, requestItemDBTypes, false, requestItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItem struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, unitDBTypes, false, unitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Unit struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UnitID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Unit().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := RequestItemSlice{&local}
	if err = local.L.LoadUnit(ctx, tx, false, (*[]*RequestItem)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Unit == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Unit = nil
	if err = local.L.LoadUnit(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Unit == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testRequestItemToOneSetOpItemUsingItem(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RequestItem
	var b, c Item

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, requestItemDBTypes, false, strmangle.SetComplement(requestItemPrimaryKeyColumns, requestItemColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, itemDBTypes, false, strmangle.SetComplement(itemPrimaryKeyColumns, itemColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, itemDBTypes, false, strmangle.SetComplement(itemPrimaryKeyColumns, itemColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Item{&b, &c} {
		err = a.SetItem(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Item != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RequestItems[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ItemID != x.ID {
			t.Error("foreign key was wrong value", a.ItemID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ItemID))
		reflect.Indirect(reflect.ValueOf(&a.ItemID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ItemID != x.ID {
			t.Error("foreign key was wrong value", a.ItemID, x.ID)
		}
	}
}
func testRequestItemToOneSetOpRequestUsingRequest(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RequestItem
	var b, c Request

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, requestItemDBTypes, false, strmangle.SetComplement(requestItemPrimaryKeyColumns, requestItemColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, requestDBTypes, false, strmangle.SetComplement(requestPrimaryKeyColumns, requestColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, requestDBTypes, false, strmangle.SetComplement(requestPrimaryKeyColumns, requestColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Request{&b, &c} {
		err = a.SetRequest(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Request != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RequestItems[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.RequestID != x.ID {
			t.Error("foreign key was wrong value", a.RequestID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.RequestID))
		reflect.Indirect(reflect.ValueOf(&a.RequestID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.RequestID != x.ID {
			t.Error("foreign key was wrong value", a.RequestID, x.ID)
		}
	}
}
func testRequestItemToOneSetOpUnitUsingUnit(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RequestItem
	var b, c Unit

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, requestItemDBTypes, false, strmangle.SetComplement(requestItemPrimaryKeyColumns, requestItemColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, unitDBTypes, false, strmangle.SetComplement(unitPrimaryKeyColumns, unitColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, unitDBTypes, false, strmangle.SetComplement(unitPrimaryKeyColumns, unitColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Unit{&b, &c} {
		err = a.SetUnit(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Unit != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RequestItems[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UnitID != x.ID {
			t.Error("foreign key was wrong value", a.UnitID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UnitID))
		reflect.Indirect(reflect.ValueOf(&a.UnitID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UnitID != x.ID {
			t.Error("foreign key was wrong value", a.UnitID, x.ID)
		}
	}
}

func testRequestItemsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestItem{}
	if err = randomize.Struct(seed, o, requestItemDBTypes, true, requestItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRequestItemsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestItem{}
	if err = randomize.Struct(seed, o, requestItemDBTypes, true, requestItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RequestItemSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRequestItemsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestItem{}
	if err = randomize.Struct(seed, o, requestItemDBTypes, true, requestItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RequestItems().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	requestItemDBTypes = map[string]string{`ID`: `text`, `ItemID`: `text`, `UnitID`: `text`, `Quantity`: `numeric`, `RequestID`: `text`}
	_                  = bytes.MinRead
)

func testRequestItemsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(requestItemPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(requestItemAllColumns) == len(requestItemPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RequestItem{}
	if err = randomize.Struct(seed, o, requestItemDBTypes, true, requestItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RequestItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, requestItemDBTypes, true, requestItemPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RequestItem struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRequestItemsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(requestItemAllColumns) == len(requestItemPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RequestItem{}
	if err = randomize.Struct(seed, o, requestItemDBTypes, true, requestItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RequestItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, requestItemDBTypes, true, requestItemPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RequestItem struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(requestItemAllColumns, requestItemPrimaryKeyColumns) {
		fields = requestItemAllColumns
	} else {
		fields = strmangle.SetComplement(
			requestItemAllColumns,
			requestItemPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := RequestItemSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRequestItemsUpsert(t *testing.T) {
	t.Parallel()

	if len(requestItemAllColumns) == len(requestItemPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := RequestItem{}
	if err = randomize.Struct(seed, &o, requestItemDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RequestItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RequestItem: %s", err)
	}

	count, err := RequestItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, requestItemDBTypes, false, requestItemPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RequestItem struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RequestItem: %s", err)
	}

	count, err = RequestItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
