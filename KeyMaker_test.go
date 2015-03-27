// KeyMaker_Test.go
package KeyMaker

import (
	"encoding/hex"
	"testing"
)

func Test1Key(t *testing.T) {
	var testString1 = "abc"
	var expectedValue = "900150983cd24fb0d6963f7d28e17f72"

	v, e := KeysToObjectId(testString1)

	var gotAsString = hex.EncodeToString(v[:])

	if e != nil {
		t.Error(
			"For", testString1,
			"expected", "no error",
			"got", e,
		)
	}

	if len(v) != 16 {
		t.Error(
			"For", testString1,
			"expected", 16,
			"got", len(v),
		)
	}

	if gotAsString != expectedValue {
		t.Error(
			"For", testString1,
			"expected", expectedValue,
			"got", gotAsString,
		)
	}
}

func Test2Key(t *testing.T) {
	var testString1 = "someemail@somedomain.com"
	var testString2 = "4000"
	var expectedValue = "a2cb40efd0d63e0283a615696eeb761d"

	v, e := KeysToObjectId(testString1, testString2)

	var gotAsString = hex.EncodeToString(v[:])

	if e != nil {
		t.Error(
			"For", testString1,
			"expected", "no error",
			"got", e,
		)
	}

	if len(v) != 16 {
		t.Error(
			"For", testString1,
			"expected", 16,
			"got", len(v),
		)
	}

	if gotAsString != expectedValue {
		t.Error(
			"For", testString1,
			"expected", expectedValue,
			"got", gotAsString,
		)
	}
}

func TestCaseInsensitivity(t *testing.T) {
	var testString1 = "someemail@somedomain.com"
	var testString2 = "soMeeMail@somEdOmain.com"

	v, e := KeysToObjectId(testString1)
	v2, e2 := KeysToObjectId(testString2)

	var gotAsString = hex.EncodeToString(v[:])
	var gotAsString2 = hex.EncodeToString(v2[:])

	if e != nil {
		t.Error(
			"For", testString1,
			"expected", "no error",
			"got", e,
		)
	}

	if e2 != nil {
		t.Error(
			"For", testString1,
			"expected", "no error",
			"got", e2,
		)
	}

	if len(v) != 16 {
		t.Error(
			"For", testString1,
			"expected", 16,
			"got", len(v),
		)
	}

	if len(v2) != 16 {
		t.Error(
			"For", testString1,
			"expected", 16,
			"got", len(v2),
		)
	}

	if gotAsString != gotAsString2 {
		t.Error(
			"For", testString1, "and", testString2,
			"expected", true,
			"got", false,
		)
	}
}
