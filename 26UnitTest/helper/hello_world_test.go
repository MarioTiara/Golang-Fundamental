package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	expected := "Hello mario"
	result := HelloWorld("mario")
	if result != expected {
		errmessage := fmt.Sprintf("\nresult: %s \nexpected: %s", result, expected)
		t.Error(errmessage)
		//t.Fail()
	}
}

func TestHelloWorldPratama(t *testing.T) {
	expected := "Hello pratama"
	result := HelloWorld("putra")
	if result != expected {
		errmessage := fmt.Sprintf("\nresult: %s \nexpected: %s", result, expected)
		t.Error(errmessage)
		//t.Fail()
	}
}
