package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
-> go test = menjalankan semua unit testing
-> go test -v = menjalankkan semua unit testing dengan melihat functionnya
-> go text -v -run {nama function} = menjalankan unit testing tertentu
*/

// test error testing
func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Qito")
	if result != "Hello Qito" {
		t.Error("Result must be Helo Qito")
	}
	fmt.Println("TestHelloWorld Done")
}

// test fatal testing
func TestHelloWorldQito(t *testing.T) {
	result := HelloWorld("Qito")
	if result != "Hello Qito" {
		t.Fatal("Result must be Hello Qito")
	}
	fmt.Println("TestHelloWorldQito Done")
}

// assert test testify
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Qito")
	assert.Equal(t, "Hello Qito", result, "Result harus Hello Qito")
	fmt.Println("Tested with Assert testify")
}

// require test testify
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Qito")
	require.Equal(t, "Hello Qito", result, "Result harus Hello Qito")
	fmt.Println("Tested with Require testify")
}

// skip testing
func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("tidak bisa dijalankan pada linux")
	}
	result := HelloWorld("Qito")
	require.Equal(t, "Hello Qito", result, "Result harus Hello Qito")
}

//before after testing
func TestMain(m *testing.M) {
	fmt.Println("Memulai unit test")

	m.Run() // eksekusi semua unit test

	fmt.Println("unit test selesai")
}

//sub test
func TestSubTest(t *testing.T) {
	t.Run("Qito", func(t *testing.T) {
		result := HelloWorld("Qito")
		require.Equal(t, "Hello Qito", result, "Result harus Hello Qito")
	})
}

// table test = test dinamis
func TestTableTest(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Qito",
			request:  "Qito",
			expected: "Hello Qito",
		},
		{
			name:     "Joko",
			request:  "Joko",
			expected: "Hello Joko",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

// benchmark test
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Qito")
	}
}

func BenchmarkHelloWorldQito(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Qito")
	}
}

// sub benchmark
func BenchmarkSub(b *testing.B) {
	b.Run("Qito", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Qito")
		}
	})
	b.Run("Rizqi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Rizqi")
		}
	})
}

// table benchmark
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Qito",
			request: "Qito",
		}, {
			name:    "Rizqi",
			request: "Rizqi",
		}, {
			name:    "Rizqi Qito",
			request: "Rizqi Qito",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.name)
			}
		})
	}
}
