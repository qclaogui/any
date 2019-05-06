package any

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAnyUnmarshalJSON(t *testing.T) {
	type user struct {
		Name  string
		Age   int
		IsMan bool
		Point float64
		Kong  string
	}
	var u = user{"qc", 10, true, 20.5, ""}
	b, _ := json.Marshal(u)
	fmt.Printf("%s\n", b)

	type temp struct {
		Name  Value
		Age   Value
		IsMan Value
		Point Value
	}

	var tem temp
	err := json.Unmarshal(b, &tem)
	if df := cmp.Diff(err, nil); df != "" {
		t.Errorf("👉 \x1b[92m%s\x1b[39m", df)
	}

	fmt.Printf("%#v\n", tem)
	fmt.Printf("%#v\n", tem.Name.String())

	b1, _ := json.Marshal(tem)
	fmt.Printf("%s\n", b1)

	var u1 user
	err = json.Unmarshal(b, &u1)
	if df := cmp.Diff(err, nil); df != "" {
		t.Errorf("👉 \x1b[92m%s\x1b[39m", df)
	}

	fmt.Printf("%#v\n", u1)
}
