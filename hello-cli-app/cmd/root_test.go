package cmd

import "testing"

func TestGreeting(t *testing.T) {
	emptyGreeting := greeting("")
	if emptyGreeting != "Hello, " {
		t.Errorf("greeting(\"\") failed! %v", emptyGreeting)
	}

	nonEmptyValue := greeting("Davor")
	if nonEmptyValue != "Hello, Davor" {
		t.Errorf("greeting(\"\") failed! %v", nonEmptyValue)
	}
}
