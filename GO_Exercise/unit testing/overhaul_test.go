package overhaul

import (
	"regexp"
	"testing"
)

// TestCheck_msgName calls overhaul.Check_msg with a name, checking
// for a valid return value.
func TestCheck_msg(t *testing.T) {
	name := "kloudone"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Check_msg("kloudone")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Check_msg("kloudone") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestCheck_msgEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestCheck_msgEmpty(t *testing.T) {
	msg, err := Check_msg("")
	if msg != "" || err == nil {
		t.Fatalf(`Check_msg("") = %q, %v, want "", error`, msg, err)
	}
}
