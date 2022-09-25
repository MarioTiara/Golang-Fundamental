package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("Before Unit Test")
	m.Run()
	fmt.Println("After Unit Test")
}

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
	result := HelloWorld("pratama")
	if result != expected {
		errmessage := fmt.Sprintf("\nresult: %s \nexpected: %s", result, expected)
		t.Error(errmessage)
		//t.Fail()
	}
}

func TestHelloWorRequire(t *testing.T) {
	expected := "Hello pratama"
	result := HelloWorld("pratama")
	errmessage := fmt.Sprintf("result must be '%s'", expected)
	require.Equal(t, expected, result, errmessage)
}

func TestHelloWorldAssert(t *testing.T) {
	expected := "Hello pratama"
	result := HelloWorld("pratama")
	errmessage := fmt.Sprintf("result must be '%s'", expected)
	assert.Equal(t, expected, result, errmessage)
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Unit test can't running on windows")
	}
	expected := "Hello pratama"
	result := HelloWorld("pratama")
	errmessage := fmt.Sprintf("result must be '%s'", expected)
	assert.Equal(t, expected, result, errmessage)
}

func TestSubTest(t *testing.T) {
	t.Run("Mario", func(t *testing.T) {
		expected := "Hello Mario"
		result := HelloWorld("Mario")
		errmessage := fmt.Sprintf("result must be '%s'", expected)
		assert.Equal(t, expected, result, errmessage)
	})

	t.Run("Pratama", func(t *testing.T) {
		expected := "Hello Pratama"
		result := HelloWorld("Pratama")
		errmessage := fmt.Sprintf("result must be '%s'", expected)
		assert.Equal(t, expected, result, errmessage)
	})
}

func TestTableHelloworld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Mario",
			request:  "Mario",
			expected: "Hello Mario",
		},
		{
			name:     "Pratama",
			request:  "Pratama",
			expected: "Hello Pratama",
		},
		{
			name:     "Roby",
			request:  "Roby",
			expected: "Hello Roby",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}

}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Mario")
	}
}

func BenchmarkHelloWorldKurniawan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Kurniawan")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("mario", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("mario")
		}
	})
	b.Run("pratama", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("pratama")
		}
	})
}

func BenchmarkTable(b *testing.B) {
	bencmarks := []struct {
		name    string
		request string
	}{
		{name: "Mario", request: "Mario"},
		{name: "Pratama", request: "Pratama"},
		{name: "MarioTiaraPratama", request: "MarioTiaraPratama"},
	}

	for _, bench := range bencmarks {
		b.Run(bench.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(bench.request)
			}
		})
	}
}
