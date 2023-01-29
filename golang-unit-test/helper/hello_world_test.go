package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("hai")
	}
}

func BenchmarkHelloWorldSub(b *testing.B) {
	b.Run("bench1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("hai")
		}
	})

	b.Run("bench2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("tes")
		}
	})
}

func BenchmarkHelloWorldTable(b *testing.B) {
	benchmarks := []struct {
		name string
		input string
	}{
		{
			name: "HelloWorld(hai)",
			input: "hai",
		},
		{
			name: "HelloWorld(tes)",
			input: "tes",
		},
	}

	for _, benchmark := range benchmarks{
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.input)
			}
		})
	}
}

func TestMain(m *testing.M){
	fmt.Println("Before Test")
	m.Run()
	fmt.Println("After Test")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("hai")
	if result != "Hello Gopher" {
		// failed
		t.Error("Expected Hello Gopher, got ", result)
	}
	fmt.Println("TestHelloWorld done")
}

func TestHelloWorld1(t *testing.T) {
	result := HelloWorld("hai")
	if result != "Hello Gopher" {
		// failed
		t.Fatal("Expected Hello Gopher, got ", result)
	}
	fmt.Println("TestHelloWorld1 done")
}

func TestHelloWordAsset(t *testing.T) {
	result := HelloWorld("hai")
	assert.Equal(t, "Hello Gopher", result, "they should be equal")
	fmt.Println("TestHelloWordAsset done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("hai")
	require.Equal(t, "Hello Gopher", result, "they should be equal")
	fmt.Println("TestHelloWorldRequire done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can't run on windows")
	}

	result := HelloWorld("hai")
	assert.Equal(t, "Hello hai", result, "they should be equal")
}

func TestSubTest(t *testing.T) {
	t.Run("A=1", func(t *testing.T) {
		// t.Parallel()
		result := HelloWorld("hai")
		assert.Equal(t, "Hello hai", result, "they should be equal")
	})

	t.Run("A=2", func(t *testing.T) {
		// t.Parallel()
		result := HelloWorld("tes")
		assert.Equal(t, "Hello tes", result, "they should be equal")
	})
}

func TestHelloWorldTable(t *testing.T){
	tests := []struct{
		name string
		input string
		expected string
	}{
		{
			name: "Hello hai", 
			input: "hai", 
			expected: "Hello hai",
		},
		{
			name: "Hello tes",
			input: "tes",
			expected: "Hello tes",
		},
		{
			name: "Hello gopher",
			input: "gopher",
			expected: "Hello gopher",
		},
	}

	for _, test := range tests{
		t.Run(test.name, func(t *testing.T){
			result := HelloWorld(test.input)
			require.Equal(t, test.expected, result, "they should be equal")
		})
	}
}