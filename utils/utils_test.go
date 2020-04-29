package utils_test

import (
	"WhistleNewsBackend/utils"
	"testing"
	"time"
)

func TestTimeConversion(t *testing.T) {
	
	td := time.Now().AddDate(0, 0, -10)
	d:=utils.TimeSince(td)
	if d != "10 days ago" {
		t.Error("wrong duration output")
	}

	tm := time.Now().AddDate(0, -7, 0)
	m:=utils.TimeSince(tm)
	if m != "7 months ago" {
		t.Error("wrong duration output")
	}

	ty := time.Now().AddDate(-3, 0, 0)
	y:=utils.TimeSince(ty)
	if y != "3 years ago" {
		t.Error("wrong duration output")
	}

}