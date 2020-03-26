// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("AllocationItems", testAllocationItems)
	t.Run("Allocations", testAllocations)
	t.Run("DonationItems", testDonationItems)
	t.Run("Donations", testDonations)
	t.Run("Items", testItems)
	t.Run("RequestItems", testRequestItems)
	t.Run("Requests", testRequests)
	t.Run("Stocks", testStocks)
	t.Run("Units", testUnits)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("AllocationItems", testAllocationItemsDelete)
	t.Run("Allocations", testAllocationsDelete)
	t.Run("DonationItems", testDonationItemsDelete)
	t.Run("Donations", testDonationsDelete)
	t.Run("Items", testItemsDelete)
	t.Run("RequestItems", testRequestItemsDelete)
	t.Run("Requests", testRequestsDelete)
	t.Run("Stocks", testStocksDelete)
	t.Run("Units", testUnitsDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("AllocationItems", testAllocationItemsQueryDeleteAll)
	t.Run("Allocations", testAllocationsQueryDeleteAll)
	t.Run("DonationItems", testDonationItemsQueryDeleteAll)
	t.Run("Donations", testDonationsQueryDeleteAll)
	t.Run("Items", testItemsQueryDeleteAll)
	t.Run("RequestItems", testRequestItemsQueryDeleteAll)
	t.Run("Requests", testRequestsQueryDeleteAll)
	t.Run("Stocks", testStocksQueryDeleteAll)
	t.Run("Units", testUnitsQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("AllocationItems", testAllocationItemsSliceDeleteAll)
	t.Run("Allocations", testAllocationsSliceDeleteAll)
	t.Run("DonationItems", testDonationItemsSliceDeleteAll)
	t.Run("Donations", testDonationsSliceDeleteAll)
	t.Run("Items", testItemsSliceDeleteAll)
	t.Run("RequestItems", testRequestItemsSliceDeleteAll)
	t.Run("Requests", testRequestsSliceDeleteAll)
	t.Run("Stocks", testStocksSliceDeleteAll)
	t.Run("Units", testUnitsSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("AllocationItems", testAllocationItemsExists)
	t.Run("Allocations", testAllocationsExists)
	t.Run("DonationItems", testDonationItemsExists)
	t.Run("Donations", testDonationsExists)
	t.Run("Items", testItemsExists)
	t.Run("RequestItems", testRequestItemsExists)
	t.Run("Requests", testRequestsExists)
	t.Run("Stocks", testStocksExists)
	t.Run("Units", testUnitsExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("AllocationItems", testAllocationItemsFind)
	t.Run("Allocations", testAllocationsFind)
	t.Run("DonationItems", testDonationItemsFind)
	t.Run("Donations", testDonationsFind)
	t.Run("Items", testItemsFind)
	t.Run("RequestItems", testRequestItemsFind)
	t.Run("Requests", testRequestsFind)
	t.Run("Stocks", testStocksFind)
	t.Run("Units", testUnitsFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("AllocationItems", testAllocationItemsBind)
	t.Run("Allocations", testAllocationsBind)
	t.Run("DonationItems", testDonationItemsBind)
	t.Run("Donations", testDonationsBind)
	t.Run("Items", testItemsBind)
	t.Run("RequestItems", testRequestItemsBind)
	t.Run("Requests", testRequestsBind)
	t.Run("Stocks", testStocksBind)
	t.Run("Units", testUnitsBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("AllocationItems", testAllocationItemsOne)
	t.Run("Allocations", testAllocationsOne)
	t.Run("DonationItems", testDonationItemsOne)
	t.Run("Donations", testDonationsOne)
	t.Run("Items", testItemsOne)
	t.Run("RequestItems", testRequestItemsOne)
	t.Run("Requests", testRequestsOne)
	t.Run("Stocks", testStocksOne)
	t.Run("Units", testUnitsOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("AllocationItems", testAllocationItemsAll)
	t.Run("Allocations", testAllocationsAll)
	t.Run("DonationItems", testDonationItemsAll)
	t.Run("Donations", testDonationsAll)
	t.Run("Items", testItemsAll)
	t.Run("RequestItems", testRequestItemsAll)
	t.Run("Requests", testRequestsAll)
	t.Run("Stocks", testStocksAll)
	t.Run("Units", testUnitsAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("AllocationItems", testAllocationItemsCount)
	t.Run("Allocations", testAllocationsCount)
	t.Run("DonationItems", testDonationItemsCount)
	t.Run("Donations", testDonationsCount)
	t.Run("Items", testItemsCount)
	t.Run("RequestItems", testRequestItemsCount)
	t.Run("Requests", testRequestsCount)
	t.Run("Stocks", testStocksCount)
	t.Run("Units", testUnitsCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("AllocationItems", testAllocationItemsHooks)
	t.Run("Allocations", testAllocationsHooks)
	t.Run("DonationItems", testDonationItemsHooks)
	t.Run("Donations", testDonationsHooks)
	t.Run("Items", testItemsHooks)
	t.Run("RequestItems", testRequestItemsHooks)
	t.Run("Requests", testRequestsHooks)
	t.Run("Stocks", testStocksHooks)
	t.Run("Units", testUnitsHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("AllocationItems", testAllocationItemsInsert)
	t.Run("AllocationItems", testAllocationItemsInsertWhitelist)
	t.Run("Allocations", testAllocationsInsert)
	t.Run("Allocations", testAllocationsInsertWhitelist)
	t.Run("DonationItems", testDonationItemsInsert)
	t.Run("DonationItems", testDonationItemsInsertWhitelist)
	t.Run("Donations", testDonationsInsert)
	t.Run("Donations", testDonationsInsertWhitelist)
	t.Run("Items", testItemsInsert)
	t.Run("Items", testItemsInsertWhitelist)
	t.Run("RequestItems", testRequestItemsInsert)
	t.Run("RequestItems", testRequestItemsInsertWhitelist)
	t.Run("Requests", testRequestsInsert)
	t.Run("Requests", testRequestsInsertWhitelist)
	t.Run("Stocks", testStocksInsert)
	t.Run("Stocks", testStocksInsertWhitelist)
	t.Run("Units", testUnitsInsert)
	t.Run("Units", testUnitsInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("AllocationItemToAllocationUsingAllocation", testAllocationItemToOneAllocationUsingAllocation)
	t.Run("AllocationItemToDonationItemUsingDonationItem", testAllocationItemToOneDonationItemUsingDonationItem)
	t.Run("AllocationItemToUnitUsingUnit", testAllocationItemToOneUnitUsingUnit)
	t.Run("AllocationToRequestUsingRequest", testAllocationToOneRequestUsingRequest)
	t.Run("AllocationToUserUsingAdmin", testAllocationToOneUserUsingAdmin)
	t.Run("DonationItemToDonationUsingDonation", testDonationItemToOneDonationUsingDonation)
	t.Run("DonationItemToItemUsingItem", testDonationItemToOneItemUsingItem)
	t.Run("DonationItemToUnitUsingUnit", testDonationItemToOneUnitUsingUnit)
	t.Run("DonationToUserUsingDonator", testDonationToOneUserUsingDonator)
	t.Run("RequestItemToItemUsingItem", testRequestItemToOneItemUsingItem)
	t.Run("RequestItemToRequestUsingRequest", testRequestItemToOneRequestUsingRequest)
	t.Run("RequestItemToUnitUsingUnit", testRequestItemToOneUnitUsingUnit)
	t.Run("RequestToUserUsingDonationApplicant", testRequestToOneUserUsingDonationApplicant)
	t.Run("StockToItemUsingItem", testStockToOneItemUsingItem)
	t.Run("StockToUnitUsingUnit", testStockToOneUnitUsingUnit)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("AllocationToAllocationItems", testAllocationToManyAllocationItems)
	t.Run("DonationItemToAllocationItems", testDonationItemToManyAllocationItems)
	t.Run("DonationToDonationItems", testDonationToManyDonationItems)
	t.Run("ItemToDonationItems", testItemToManyDonationItems)
	t.Run("ItemToRequestItems", testItemToManyRequestItems)
	t.Run("ItemToStocks", testItemToManyStocks)
	t.Run("RequestToAllocations", testRequestToManyAllocations)
	t.Run("RequestToRequestItems", testRequestToManyRequestItems)
	t.Run("UnitToAllocationItems", testUnitToManyAllocationItems)
	t.Run("UnitToDonationItems", testUnitToManyDonationItems)
	t.Run("UnitToRequestItems", testUnitToManyRequestItems)
	t.Run("UnitToStocks", testUnitToManyStocks)
	t.Run("UserToAdminAllocations", testUserToManyAdminAllocations)
	t.Run("UserToDonatorDonations", testUserToManyDonatorDonations)
	t.Run("UserToDonationApplicantRequests", testUserToManyDonationApplicantRequests)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("AllocationItemToAllocationUsingAllocationItems", testAllocationItemToOneSetOpAllocationUsingAllocation)
	t.Run("AllocationItemToDonationItemUsingAllocationItems", testAllocationItemToOneSetOpDonationItemUsingDonationItem)
	t.Run("AllocationItemToUnitUsingAllocationItems", testAllocationItemToOneSetOpUnitUsingUnit)
	t.Run("AllocationToRequestUsingAllocations", testAllocationToOneSetOpRequestUsingRequest)
	t.Run("AllocationToUserUsingAdminAllocations", testAllocationToOneSetOpUserUsingAdmin)
	t.Run("DonationItemToDonationUsingDonationItems", testDonationItemToOneSetOpDonationUsingDonation)
	t.Run("DonationItemToItemUsingDonationItems", testDonationItemToOneSetOpItemUsingItem)
	t.Run("DonationItemToUnitUsingDonationItems", testDonationItemToOneSetOpUnitUsingUnit)
	t.Run("DonationToUserUsingDonatorDonations", testDonationToOneSetOpUserUsingDonator)
	t.Run("RequestItemToItemUsingRequestItems", testRequestItemToOneSetOpItemUsingItem)
	t.Run("RequestItemToRequestUsingRequestItems", testRequestItemToOneSetOpRequestUsingRequest)
	t.Run("RequestItemToUnitUsingRequestItems", testRequestItemToOneSetOpUnitUsingUnit)
	t.Run("RequestToUserUsingDonationApplicantRequests", testRequestToOneSetOpUserUsingDonationApplicant)
	t.Run("StockToItemUsingStocks", testStockToOneSetOpItemUsingItem)
	t.Run("StockToUnitUsingStocks", testStockToOneSetOpUnitUsingUnit)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("AllocationToAllocationItems", testAllocationToManyAddOpAllocationItems)
	t.Run("DonationItemToAllocationItems", testDonationItemToManyAddOpAllocationItems)
	t.Run("DonationToDonationItems", testDonationToManyAddOpDonationItems)
	t.Run("ItemToDonationItems", testItemToManyAddOpDonationItems)
	t.Run("ItemToRequestItems", testItemToManyAddOpRequestItems)
	t.Run("ItemToStocks", testItemToManyAddOpStocks)
	t.Run("RequestToAllocations", testRequestToManyAddOpAllocations)
	t.Run("RequestToRequestItems", testRequestToManyAddOpRequestItems)
	t.Run("UnitToAllocationItems", testUnitToManyAddOpAllocationItems)
	t.Run("UnitToDonationItems", testUnitToManyAddOpDonationItems)
	t.Run("UnitToRequestItems", testUnitToManyAddOpRequestItems)
	t.Run("UnitToStocks", testUnitToManyAddOpStocks)
	t.Run("UserToAdminAllocations", testUserToManyAddOpAdminAllocations)
	t.Run("UserToDonatorDonations", testUserToManyAddOpDonatorDonations)
	t.Run("UserToDonationApplicantRequests", testUserToManyAddOpDonationApplicantRequests)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("AllocationItems", testAllocationItemsReload)
	t.Run("Allocations", testAllocationsReload)
	t.Run("DonationItems", testDonationItemsReload)
	t.Run("Donations", testDonationsReload)
	t.Run("Items", testItemsReload)
	t.Run("RequestItems", testRequestItemsReload)
	t.Run("Requests", testRequestsReload)
	t.Run("Stocks", testStocksReload)
	t.Run("Units", testUnitsReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("AllocationItems", testAllocationItemsReloadAll)
	t.Run("Allocations", testAllocationsReloadAll)
	t.Run("DonationItems", testDonationItemsReloadAll)
	t.Run("Donations", testDonationsReloadAll)
	t.Run("Items", testItemsReloadAll)
	t.Run("RequestItems", testRequestItemsReloadAll)
	t.Run("Requests", testRequestsReloadAll)
	t.Run("Stocks", testStocksReloadAll)
	t.Run("Units", testUnitsReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("AllocationItems", testAllocationItemsSelect)
	t.Run("Allocations", testAllocationsSelect)
	t.Run("DonationItems", testDonationItemsSelect)
	t.Run("Donations", testDonationsSelect)
	t.Run("Items", testItemsSelect)
	t.Run("RequestItems", testRequestItemsSelect)
	t.Run("Requests", testRequestsSelect)
	t.Run("Stocks", testStocksSelect)
	t.Run("Units", testUnitsSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("AllocationItems", testAllocationItemsUpdate)
	t.Run("Allocations", testAllocationsUpdate)
	t.Run("DonationItems", testDonationItemsUpdate)
	t.Run("Donations", testDonationsUpdate)
	t.Run("Items", testItemsUpdate)
	t.Run("RequestItems", testRequestItemsUpdate)
	t.Run("Requests", testRequestsUpdate)
	t.Run("Stocks", testStocksUpdate)
	t.Run("Units", testUnitsUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("AllocationItems", testAllocationItemsSliceUpdateAll)
	t.Run("Allocations", testAllocationsSliceUpdateAll)
	t.Run("DonationItems", testDonationItemsSliceUpdateAll)
	t.Run("Donations", testDonationsSliceUpdateAll)
	t.Run("Items", testItemsSliceUpdateAll)
	t.Run("RequestItems", testRequestItemsSliceUpdateAll)
	t.Run("Requests", testRequestsSliceUpdateAll)
	t.Run("Stocks", testStocksSliceUpdateAll)
	t.Run("Units", testUnitsSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
