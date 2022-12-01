package dao

import (
	"testing"
)

func TestConstructMonthLookUpList(t *testing.T) {
	monthLookUp := ConstructMonthLookUpList(0)
	if len(monthLookUp.MonthToLook) != 0 {
		t.Error("Month look up should have zero entries")
	}

	monthLookUp = ConstructMonthLookUpList(1)
	if len(monthLookUp.MonthToLook) != 1 {
		t.Error("Month look up should have just one entry")
	}

	monthLookUp = ConstructMonthLookUpList(2)
	if len(monthLookUp.MonthToLook) != 2 {
		t.Error("Month look up should have just two entries")
	}
}
