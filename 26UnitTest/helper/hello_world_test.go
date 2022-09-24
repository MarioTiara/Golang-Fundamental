package helper

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	expected := "Hello mario"
	result := HelloWorld("mario")
	if result != expected {
		//errmessage := fmt.Sprintf("result: %s expected: %s", result, expected)
		t.Fail()
	}
}

func TestHelloWorldPratama(t *testing.T) {
	expected := "Hello pratama"
	result := HelloWorld("putra")
	if result != expected {
		//errmessage := fmt.Sprintf("result: %s expected: %s", result, expected)
		t.Fail()
	}
}
