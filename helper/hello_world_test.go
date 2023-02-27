package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("Before unit test")

	m.Run()

	fmt.Println("After unit test")
}

func TestHelloWorldBudi(t *testing.T) {
	result := HelloWorld("Budi")

	if result != "Hello sd" {
		// t.Fail() // code selanjutnya akan tetap dijalankan
		t.Error("Result must be Hello Budi") // mirip dengan Fail tetapi memiliki parameter error message
	}
	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldUtomo(t *testing.T) {
	result := HelloWorld("Utomo")

	if result != "Hello ds" {
		// t.FailNow() // code selanjutnya tidak akan dijalankan
		t.Fatal("Result must be Hello Utomo") // mirip dengan FailNow tetapi memiliki parameter error message
	}

	fmt.Println("TestHelloWorldUtomo Done")
}

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Budi")
	assert.Equal(t, result, "Hello Budi", "Result must be Hello Budi")
	fmt.Println("TestHellowWorldWithAssertion done")
}

// assert -> kita gagal akan memanggil Fail
func TestHelloWorldAssertionNotPass(t *testing.T) {
	result := HelloWorld("Budi")
	assert.Equal(t, result, "Hello Utomo", "Result must be Hello Budi")
	fmt.Println("TestHellowWorldWithAssertion done")
}

// require -> kita gagal akan memanggil FailNow
func TestHelloWorldRequireNotPass(t *testing.T) {
	result := HelloWorld("Budi")
	require.Equal(t, result, "Hello Utomo", "Result must be Hello Budi")
	fmt.Println("TestHellowWorldWithAssertion done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Cannot run on window")
	}

	result := HelloWorld("Budi")
	require.Equal(t, result, "Hello Utomo", "Result must be Hello Budi")
	fmt.Println("TestSkip done", runtime.GOOS)
}

func TestSubTest(t *testing.T) {
	t.Run("Budi", func(t *testing.T) {
		result := HelloWorld("Budi")
		assert.Equal(t, "Hello Budi", result, "Result must be Hello Budi")
	})

	t.Run("Utomo", func(t *testing.T) {
		result := HelloWorld("Utomo")
		assert.Equal(t, "Hello Utomo", result, "Result must be Hello Utomo")

	})
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Budi",
			request:  "Budi",
			expected: "Hello Budi",
		},

		{
			name:     "Utomo",
			request:  "Utomo",
			expected: "Hello Utomo",
		},

		{
			name:     "Shinta",
			request:  "Shinta",
			expected: "Hello Shinta",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result, "Result must be Hello "+test.request)
		})
	}
}

// Benchmarking
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Budi")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Budi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Budi")
		}
	})

	b.Run("Utomo", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Budi")
		}
	})
}

func BenchmarkTable(b *testing.B) {
	benchamarks := []struct {
		name, request string
	}{
		{
			name:    "Budi",
			request: "Budi",
		},

		{
			name:    "Utomo",
			request: "Utomo",
		},

		{
			name:    "Shinta",
			request: "Shinta",
		},
	}

	for _, becbenchamark := range benchamarks {
		b.Run(becbenchamark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(becbenchamark.request)
			}
		})
	}
}
