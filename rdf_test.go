package rdf

import (
	"testing"
	"time"
)

func TestTermTypeBlank(t *testing.T) {
	_, err := NewBlank("")
	if err != ErrBlankNodeMissingID {
		t.Errorf("NewBlank(\" \") => %v; want ErrBlankNodeMissingID", err)
	}

	_, err = NewBlank(" \n\r \t ")
	if err != ErrBlankNodeMissingID {
		t.Errorf("NewBlank(\" \n\r \t \") => %v; want ErrBlankNodeMissingID", nil)
	}

	b, err := NewBlank("a")
	if err != nil {
		t.Fatalf("NewBlank(\"a\") failed with %v", err)
	}

	want := "_:a"
	if b.String() != want {
		t.Errorf("NewBlank(\"a\").String() => %v; want %v", b.String(), want)
	}

	b2 := Blank{ID: "a"}
	b3 := Blank{ID: "b"}

	if !b.Eq(b2) {
		t.Errorf("two blank nodes with same ID should be equal")
	}

	if b.Eq(b3) {
		t.Errorf("two blank nodes with different IDs should not be equal")
	}
}

func TestTermTypeIRI(t *testing.T) {
	uri := IRI{IRI: "x://y/z"}
	want := "<x://y/z>"
	if uri.String() != want {
		t.Errorf("NewIRI(\"x://y/z\").String() => %s; want %s", uri.String(), want)
	}

	_, err := NewIRI("")
	if err != ErrIRIEmptyInput {
		t.Errorf("NewIRI(\" \") => %v want ErrIRIEmptyInput", err)
	}

	_, err = NewIRI("<&httoop.dott")
	if err != ErrIRIInvalidCharacters {
		t.Errorf("NewIRI(\"<&http.dott\") => %v; want errIRIInvalidCharacters", err)
	}

	a := IRI{IRI: "abba"}
	b := IRI{IRI: "ABBA"}
	c := IRI{IRI: "abba"}

	if a.Eq(b) {
		t.Errorf("two different IRIs should not be equal")
	}

	if !a.Eq(c) {
		t.Errorf("two identical IRIs should be equal")
	}

}

func TestTermTypeLiteral(t *testing.T) {
	_, err := NewLiteral([]int{1, 2, 3})
	if err == nil {
		t.Errorf("Expected an error creating Literal, got nil")
	}

	l1, _ := NewLiteral(42)
	l2, _ := NewLiteral(42.00001)
	l3, _ := NewLiteral(true)
	l4, _ := NewLiteral(false)
	l5 := NewLangLiteral("fisk", "nno")
	l6 := NewLangLiteral("fisk", "no")
	l7, _ := NewLiteral("fisk")
	l8, _ := NewLiteral(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))

	var eqTests = []struct {
		a, b Literal
		want bool
	}{
		{l1, l2, false},
		{l1, l3, false},
		{l3, l4, false},
		{l5, l6, false},
		{l6, l7, false},
	}

	for _, tt := range eqTests {
		if tt.a.Eq(tt.b) != tt.want {
			t.Errorf("%v.Eq(%v) => %v, want %v", tt.a, tt.b, tt.want, tt.a.Eq(tt.b))
		}
	}

	var formatTests = []struct {
		l    Literal
		want string
	}{
		{l1, "42"},
		{l2, "42.00001"},
		{l3, "true"},
		{l4, "false"},
		{l5, `"fisk"@nno`},
		{l7, `"fisk"`},
		{l8, `"2009-11-10T23:00:00Z"^^<http://www.w3.org/2001/XMLSchema#dateTime>`},
	}

	for _, tt := range formatTests {
		if tt.l.String() != tt.want {
			t.Errorf("Literal string formatting \"%v\", want \"%v\"", tt.want, tt.l.String())
		}
	}
}
