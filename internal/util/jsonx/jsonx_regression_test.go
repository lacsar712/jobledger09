package jsonx

import "testing"

// Regression: DecodeMap previously discarded the json.Unmarshal error and
// always returned a nil error, so malformed input silently produced an empty
// map. These cases pin down that every failure path surfaces an error.
func TestDecodeMapErrorPaths(t *testing.T) {
	cases := [][]byte{
		[]byte(`{not-json`),
		[]byte(``),
		[]byte(`{`),
		[]byte(`{"n":}`),
		[]byte(`3`),       // valid JSON, wrong type
		[]byte(`"x"`),     // valid JSON, wrong type
		[]byte(`[1,2,3]`), // valid JSON, wrong type
	}
	for _, in := range cases {
		if m, err := DecodeMap(in); err == nil {
			t.Fatalf("DecodeMap(%q) returned nil error, map=%v; want error", in, m)
		}
	}
}

func TestDecodeMapValidPaths(t *testing.T) {
	t.Run("empty object", func(t *testing.T) {
		m, err := DecodeMap([]byte(`{}`))
		if err != nil {
			t.Fatal(err)
		}
		if len(m) != 0 {
			t.Fatalf("want empty map, got %v", m)
		}
	})

	t.Run("populated object", func(t *testing.T) {
		m, err := DecodeMap([]byte(`{"a":1,"b":2}`))
		if err != nil {
			t.Fatal(err)
		}
		if m["a"] != 1 || m["b"] != 2 {
			t.Fatalf("unexpected map: %v", m)
		}
	})
}
