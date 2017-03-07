package boilingcore

import (
	"reflect"
	"sort"
	"testing"

	"github.com/pkg/errors"
	"github.com/vattle/sqlboiler/bdb"
)

func TestImportsSort(t *testing.T) {
	t.Parallel()

	a1 := importList{
		`"fmt"`,
		`"errors"`,
	}
	a2 := importList{
		`_ "github.com/lib/pq"`,
		`_ "github.com/gorilla/n"`,
		`"github.com/gorilla/mux"`,
		`"github.com/gorilla/websocket"`,
	}

	a1Expected := importList{`"errors"`, `"fmt"`}
	a2Expected := importList{
		`"github.com/gorilla/mux"`,
		`_ "github.com/gorilla/n"`,
		`"github.com/gorilla/websocket"`,
		`_ "github.com/lib/pq"`,
	}

	sort.Sort(a1)
	if !reflect.DeepEqual(a1, a1Expected) {
		t.Errorf("Expected a1 to match a1Expected, got: %v", a1)
	}

	for i, v := range a1 {
		if v != a1Expected[i] {
			t.Errorf("Expected a1[%d] to match a1Expected[%d]:\n%s\n%s\n", i, i, v, a1Expected[i])
		}
	}

	sort.Sort(a2)
	if !reflect.DeepEqual(a2, a2Expected) {
		t.Errorf("Expected a2 to match a2expected, got: %v", a2)
	}

	for i, v := range a2 {
		if v != a2Expected[i] {
			t.Errorf("Expected a2[%d] to match a2Expected[%d]:\n%s\n%s\n", i, i, v, a1Expected[i])
		}
	}
}

func TestCombineTypeImports(t *testing.T) {
	t.Parallel()

	imports1 := imports{
		standard: importList{
			`"errors"`,
			`"fmt"`,
		},
		thirdParty: importList{
			`"github.com/vattle/sqlboiler/boil"`,
		},
	}

	importsExpected := imports{
		standard: importList{
			`"errors"`,
			`"fmt"`,
			`"time"`,
		},
		thirdParty: importList{
			`"github.com/vattle/sqlboiler/boil"`,
			`"gopkg.in/nullbio/null.v6"`,
		},
	}

	cols := []bdb.Column{
		{
			Type: "null.Time",
		},
		{
			Type: "null.Time",
		},
		{
			Type: "time.Time",
		},
		{
			Type: "null.Float",
		},
	}

	res1 := combineTypeImports(imports1, importsBasedOnType, cols)

	if !reflect.DeepEqual(res1, importsExpected) {
		t.Errorf("Expected res1 to match importsExpected, got:\n\n%#v\n", res1)
	}

	imports2 := imports{
		standard: importList{
			`"errors"`,
			`"fmt"`,
			`"time"`,
		},
		thirdParty: importList{
			`"github.com/vattle/sqlboiler/boil"`,
			`"gopkg.in/nullbio/null.v6"`,
		},
	}

	res2 := combineTypeImports(imports2, importsBasedOnType, cols)

	if !reflect.DeepEqual(res2, importsExpected) {
		t.Errorf("Expected res2 to match importsExpected, got:\n\n%#v\n", res1)
	}
}

func TestCombineImports(t *testing.T) {
	t.Parallel()

	a := imports{
		standard:   importList{"fmt"},
		thirdParty: importList{"github.com/vattle/sqlboiler", "gopkg.in/nullbio/null.v6"},
	}
	b := imports{
		standard:   importList{"os"},
		thirdParty: importList{"github.com/vattle/sqlboiler"},
	}

	c := combineImports(a, b)

	if c.standard[0] != "fmt" && c.standard[1] != "os" {
		t.Errorf("Wanted: fmt, os got: %#v", c.standard)
	}
	if c.thirdParty[0] != "github.com/vattle/sqlboiler" && c.thirdParty[1] != "gopkg.in/nullbio/null.v6" {
		t.Errorf("Wanted: github.com/vattle/sqlboiler, gopkg.in/nullbio/null.v6 got: %#v", c.thirdParty)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	t.Parallel()

	hasDups := func(possible []string) error {
		for i := 0; i < len(possible)-1; i++ {
			for j := i + 1; j < len(possible); j++ {
				if possible[i] == possible[j] {
					return errors.Errorf("found duplicate: %s [%d] [%d]", possible[i], i, j)
				}
			}
		}

		return nil
	}

	if len(removeDuplicates([]string{})) != 0 {
		t.Error("It should have returned an empty slice")
	}

	oneItem := []string{"patrick"}
	slice := removeDuplicates(oneItem)
	if ln := len(slice); ln != 1 {
		t.Error("Length was wrong:", ln)
	} else if oneItem[0] != slice[0] {
		t.Errorf("Slices differ: %#v %#v", oneItem, slice)
	}

	slice = removeDuplicates([]string{"hello", "patrick", "hello"})
	if ln := len(slice); ln != 2 {
		t.Error("Length was wrong:", ln)
	}
	if err := hasDups(slice); err != nil {
		t.Error(err)
	}

	slice = removeDuplicates([]string{"five", "patrick", "hello", "hello", "patrick", "hello", "hello"})
	if ln := len(slice); ln != 3 {
		t.Error("Length was wrong:", ln)
	}
	if err := hasDups(slice); err != nil {
		t.Error(err)
	}
}

func TestCombineStringSlices(t *testing.T) {
	t.Parallel()

	var a, b []string
	slice := combineStringSlices(a, b)
	if ln := len(slice); ln != 0 {
		t.Error("Len was wrong:", ln)
	}

	a = []string{"1", "2"}
	slice = combineStringSlices(a, b)
	if ln := len(slice); ln != 2 {
		t.Error("Len was wrong:", ln)
	} else if slice[0] != a[0] || slice[1] != a[1] {
		t.Errorf("Slice mismatch: %#v %#v", a, slice)
	}

	b = a
	a = nil
	slice = combineStringSlices(a, b)
	if ln := len(slice); ln != 2 {
		t.Error("Len was wrong:", ln)
	} else if slice[0] != b[0] || slice[1] != b[1] {
		t.Errorf("Slice mismatch: %#v %#v", b, slice)
	}

	a = b
	b = []string{"3", "4"}
	slice = combineStringSlices(a, b)
	if ln := len(slice); ln != 4 {
		t.Error("Len was wrong:", ln)
	} else if slice[0] != a[0] || slice[1] != a[1] || slice[2] != b[0] || slice[3] != b[1] {
		t.Errorf("Slice mismatch: %#v + %#v != #%v", a, b, slice)
	}
}
