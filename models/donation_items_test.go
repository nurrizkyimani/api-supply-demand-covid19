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

func testDonationItems(t *testing.T) {
	t.Parallel()

	query := DonationItems()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDonationItemsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DonationItem{}
	if err = randomize.Struct(seed, o, donationItemDBTypes, true, donationItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DonationItem struct: %s", err)
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

	count, err := DonationItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDonationItemsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DonationItem{}
	if err = randomize.Struct(seed, o, donationItemDBTypes, true, donationItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DonationItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := DonationItems().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DonationItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDonationItemsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DonationItem{}
	if err = randomize.Struct(seed, o, donationItemDBTypes, true, donationItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DonationItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DonationItemSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DonationItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDonationItemsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DonationItem{}
	if err = randomize.Struct(seed, o, donationItemDBTypes, true, donationItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DonationItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DonationItemExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if DonationItem exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DonationItemExists to return true, but got false.")
	}
}

func testDonationItemsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DonationItem{}
	if err = randomize.Struct(seed, o, donationItemDBTypes, true, donationItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DonationItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	donationItemFound, err := FindDonationItem(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if donationItemFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDonationItemsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DonationItem{}
	if err = randomize.Struct(seed, o, donationItemDBTypes, true, donationItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DonationItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = DonationItems().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDonationItemsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DonationItem{}
	if err = randomize.Struct(seed, o, donationItemDBTypes, true, donationItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DonationItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := DonationItems().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDonationItemsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	donationItemOne := &DonationItem{}
	donationItemTwo := &DonationItem{}
	if err = randomize.Struct(seed, donationItemOne, donationItemDBTypes, false, donationItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DonationItem struct: %s", err)
	}
	if err = randomize.Struct(seed, donationItemTwo, donationItemDBTypes, false, donationItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DonationItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = donationItemOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = donationItemTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DonationItems().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDonationItemsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	donationItemOne := &DonationItem{}
	donationItemTwo := &DonationItem{}
	if err = randomize.Struct(seed, donationItemOne, donationItemDBTypes, false, donationItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DonationItem struct: %s", err)
	}
	if err = randomize.Struct(seed, donationItemTwo, donationItemDBTypes, false, donationItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DonationItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = donationItemOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = donationItemTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DonationItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func donationItemBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *DonationItem) error {
	*o = DonationItem{}
	return nil
}

func donationItemAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *DonationItem) error {
	*o = DonationItem{}
	return nil
}

func donationItemAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *DonationItem) error {
	*o = DonationItem{}
	return nil
}

func donationItemBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DonationItem) error {
	*o = DonationItem{}
	return nil
}

func donationItemAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DonationItem) error {
	*o = DonationItem{}
	return nil
}

func donationItemBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DonationItem) error {
	*o = DonationItem{}
	return nil
}

func donationItemAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DonationItem) error {
	*o = DonationItem{}
	return nil
}

func donationItemBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DonationItem) error {
	*o = DonationItem{}
	return nil
}

func donationItemAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DonationItem) error {
	*o = DonationItem{}
	return nil
}

func testDonationItemsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &DonationItem{}
	o := &DonationItem{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, donationItemDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DonationItem object: %s", err)
	}

	AddDonationItemHook(boil.BeforeInsertHook, donationItemBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	donationItemBeforeInsertHooks = []DonationItemHook{}

	AddDonationItemHook(boil.AfterInsertHook, donationItemAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	donationItemAfterInsertHooks = []DonationItemHook{}

	AddDonationItemHook(boil.AfterSelectHook, donationItemAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	donationItemAfterSelectHooks = []DonationItemHook{}

	AddDonationItemHook(boil.BeforeUpdateHook, donationItemBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	donationItemBeforeUpdateHooks = []DonationItemHook{}

	AddDonationItemHook(boil.AfterUpdateHook, donationItemAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	donationItemAfterUpdateHooks = []DonationItemHook{}

	AddDonationItemHook(boil.BeforeDeleteHook, donationItemBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	donationItemBeforeDeleteHooks = []DonationItemHook{}

	AddDonationItemHook(boil.AfterDeleteHook, donationItemAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	donationItemAfterDeleteHooks = []DonationItemHook{}

	AddDonationItemHook(boil.BeforeUpsertHook, donationItemBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	donationItemBeforeUpsertHooks = []DonationItemHook{}

	AddDonationItemHook(boil.AfterUpsertHook, donationItemAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	donationItemAfterUpsertHooks = []DonationItemHook{}
}

func testDonationItemsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DonationItem{}
	if err = randomize.Struct(seed, o, donationItemDBTypes, true, donationItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DonationItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DonationItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDonationItemsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DonationItem{}
	if err = randomize.Struct(seed, o, donationItemDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DonationItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(donationItemColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := DonationItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDonationItemToOneDonationUsingDonation(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local DonationItem
	var foreign Donation

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, donationItemDBTypes, false, donationItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DonationItem struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, donationDBTypes, false, donationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Donation struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.DonationID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Donation().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := DonationItemSlice{&local}
	if err = local.L.LoadDonation(ctx, tx, false, (*[]*DonationItem)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Donation == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Donation = nil
	if err = local.L.LoadDonation(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Donation == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDonationItemToOneItemUsingItem(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local DonationItem
	var foreign Item

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, donationItemDBTypes, false, donationItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DonationItem struct: %s", err)
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

	slice := DonationItemSlice{&local}
	if err = local.L.LoadItem(ctx, tx, false, (*[]*DonationItem)(&slice), nil); err != nil {
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

func testDonationItemToOneUnitUsingUnit(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local DonationItem
	var foreign Unit

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, donationItemDBTypes, false, donationItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DonationItem struct: %s", err)
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

	slice := DonationItemSlice{&local}
	if err = local.L.LoadUnit(ctx, tx, false, (*[]*DonationItem)(&slice), nil); err != nil {
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

func testDonationItemToOneSetOpDonationUsingDonation(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DonationItem
	var b, c Donation

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, donationItemDBTypes, false, strmangle.SetComplement(donationItemPrimaryKeyColumns, donationItemColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, donationDBTypes, false, strmangle.SetComplement(donationPrimaryKeyColumns, donationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, donationDBTypes, false, strmangle.SetComplement(donationPrimaryKeyColumns, donationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Donation{&b, &c} {
		err = a.SetDonation(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Donation != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.DonationItems[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.DonationID != x.ID {
			t.Error("foreign key was wrong value", a.DonationID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.DonationID))
		reflect.Indirect(reflect.ValueOf(&a.DonationID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.DonationID != x.ID {
			t.Error("foreign key was wrong value", a.DonationID, x.ID)
		}
	}
}
func testDonationItemToOneSetOpItemUsingItem(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DonationItem
	var b, c Item

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, donationItemDBTypes, false, strmangle.SetComplement(donationItemPrimaryKeyColumns, donationItemColumnsWithoutDefault)...); err != nil {
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

		if x.R.DonationItems[0] != &a {
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
func testDonationItemToOneSetOpUnitUsingUnit(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DonationItem
	var b, c Unit

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, donationItemDBTypes, false, strmangle.SetComplement(donationItemPrimaryKeyColumns, donationItemColumnsWithoutDefault)...); err != nil {
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

		if x.R.DonationItems[0] != &a {
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

func testDonationItemsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DonationItem{}
	if err = randomize.Struct(seed, o, donationItemDBTypes, true, donationItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DonationItem struct: %s", err)
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

func testDonationItemsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DonationItem{}
	if err = randomize.Struct(seed, o, donationItemDBTypes, true, donationItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DonationItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DonationItemSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDonationItemsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DonationItem{}
	if err = randomize.Struct(seed, o, donationItemDBTypes, true, donationItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DonationItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DonationItems().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	donationItemDBTypes = map[string]string{`ID`: `text`, `DonationID`: `text`, `ItemID`: `text`, `UnitID`: `text`, `Quantity`: `numeric`}
	_                   = bytes.MinRead
)

func testDonationItemsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(donationItemPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(donationItemAllColumns) == len(donationItemPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DonationItem{}
	if err = randomize.Struct(seed, o, donationItemDBTypes, true, donationItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DonationItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DonationItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, donationItemDBTypes, true, donationItemPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DonationItem struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDonationItemsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(donationItemAllColumns) == len(donationItemPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DonationItem{}
	if err = randomize.Struct(seed, o, donationItemDBTypes, true, donationItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DonationItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DonationItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, donationItemDBTypes, true, donationItemPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DonationItem struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(donationItemAllColumns, donationItemPrimaryKeyColumns) {
		fields = donationItemAllColumns
	} else {
		fields = strmangle.SetComplement(
			donationItemAllColumns,
			donationItemPrimaryKeyColumns,
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

	slice := DonationItemSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDonationItemsUpsert(t *testing.T) {
	t.Parallel()

	if len(donationItemAllColumns) == len(donationItemPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := DonationItem{}
	if err = randomize.Struct(seed, &o, donationItemDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DonationItem struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DonationItem: %s", err)
	}

	count, err := DonationItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, donationItemDBTypes, false, donationItemPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DonationItem struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DonationItem: %s", err)
	}

	count, err = DonationItems().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
