package greetings

import (
    "testing"
    "regexp"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
    name := "Gladys"
    want := regexp.MustCompile(`\b`+name+`\b`)
    msg, err := Hello("Gladys")
    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}

func TestHelloNames(t *testing.T) {
    testCases := []struct {
		name string
		want string
	}{
		{`Gladys`, "Hi,Gladys.welcome!"},
		{`Jim`, "Hi,Jim.welcome!"},
		{`Bob`, "Hi,Bob.welcome!"},
	}
    for _, test := range testCases {
		msg, err := Hello(test.name)
		if err != nil || test.want != msg {
			t.Fatalf(`Hello("%s") = %q, %v, want match for %#q, nil`, test.name, msg, err, test.want)
		}
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
    msg, err := Hello("")
    if msg != "" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}

func BenchmarkHello(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Hello("Jim")
    }
}

