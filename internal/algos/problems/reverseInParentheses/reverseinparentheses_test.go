package reverseinparentheses

import "testing"

func TestReverseinparentheses(t *testing.T) {
	input := "foo(bar(baz))blim"
	expected := "foobazrabblim"
	got := ReverseInParentheses(input)

	if expected != got {
		t.Errorf("expected: %v; received: %v", expected, got)
	}
}
