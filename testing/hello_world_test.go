package testing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST")
	m.Run()
	fmt.Println("AFTER UNIT TEST")
}

func TestTable(t *testing.T){
	test := []struct {
		name	    string
		request   string
		expected  string
	} {
		{
			name: "pratito",
			request: "pratito",
			expected: "Hello pratito",
		},
	}
	for _, test := range test {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		}) 
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Alex", func(t *testing.T) {
		result := HelloWorld("Alex")
		require.Equal(t, "Hello Alex", result, "Resul must be hello alex")
	})
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Pratito")
	assert.Equal(t, "Hello Pratito", result, "Result must be Hello Pratito")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Pratito")
	if result != "Hello Pratito" {
		panic("result is not pratito")
	}
}	