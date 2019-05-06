package any_test

import (
	"encoding/json"
	"testing"

	"github.com/qclaogui/any"

	"github.com/google/go-cmp/cmp"
)

func TestBoolCase(t *testing.T) {
	type user struct {
		Name  string    `json:"name"`
		Age   int       `json:"age"`
		IsMan any.Value `json:"is_man"`
		Point float64   `json:"point"`
	}
	tests := map[string]struct {
		data []byte
		want user
	}{
		"bool true":  {[]byte(`{"name":"qc","age":10,"is_man":true,"point":20.5}`), user{"qc", 10, "true", 20.5}},
		"bool false": {[]byte(`{"name":"qc","age":10,"is_man":false,"point":20.5}`), user{"qc", 10, "false", 20.5}},
		"bool 1":     {[]byte(`{"name":"qc","age":10,"is_man":1,"point":20.5}`), user{"qc", 10, "1", 20.5}},
		"bool 0":     {[]byte(`{"name":"qc","age":10,"is_man":0,"point":20.5}`), user{"qc", 10, "0", 20.5}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var got user
			if df := cmp.Diff(json.Unmarshal(test.data, &got), nil); df != "" {
				t.Errorf("👉 \x1b[92m%s\x1b[39m", df)
			}
			if df := cmp.Diff(got, test.want); df != "" {
				t.Errorf("👉 \x1b[92m%s\x1b[39m", df)
			}
		})
	}
}

func TestIntCase(t *testing.T) {
	type user struct {
		Name  string    `json:"name"`
		Age   any.Value `json:"age"`
		IsMan bool      `json:"is_man"`
		Point float64   `json:"point"`
	}
	tests := map[string]struct {
		data []byte
		want user
	}{
		"int To string":    {[]byte(`{"name":"qc","age":10,"is_man":true,"point":20.5}`), user{"qc", "10", true, 20.5}},
		"string To string": {[]byte(`{"name":"qc","age":"10","is_man":false,"point":20.5}`), user{"qc", "10", false, 20.5}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var got user
			if df := cmp.Diff(json.Unmarshal(test.data, &got), nil); df != "" {
				t.Errorf("👉 \x1b[92m%s\x1b[39m", df)
			}
			if df := cmp.Diff(got, test.want); df != "" {
				t.Errorf("👉 \x1b[92m%s\x1b[39m", df)
			}
		})
	}
}
