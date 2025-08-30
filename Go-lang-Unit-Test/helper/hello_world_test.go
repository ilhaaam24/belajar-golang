package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
);




func TestTableHelloWorld(t *testing.T){

	tests := []struct{
		name string
		request string
		expected string
	}{
		{
			name: "Muhammad",
			request: "Muhammad",
			expected: "Hello Muhammad",
		},
		{
			name: "Ilham",
			request: "Ilham",
			expected: "Hello Ilham",
		},
		{
			name: "Joko",
			request: "Joko",
			expected: "Hello Joko",
		},
	}

	for _, test := range tests{
		t.Run(test.name, func(t *testing.T) {
			result:= HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}




func TestMain(m *testing.M){
	fmt.Println("Test HelloWorlIlham Start")


	m.Run()


	fmt.Println("Test HelloWorlIlham Done")
}


func TestSubTest(t *testing.T){
	t.Run("Muhammad", func(t *testing.T) {
		result := HelloWorld("Muhammad")
		assert.Equal(t, "Hello Muhammad", result, "result must be 'Hello Muhammad'")
	})
	t.Run("Ilham", func(t *testing.T) {
		result := HelloWorld("Ilham")
		assert.Equal(t, "Hello Ilham", result, "result must be 'Hello Ilham'")
	})
}


func  TestHelloWorlIlham(t *testing.T){

	 result := HelloWorld("Ilham")
		
	 if result != "Hello Ilham"{
		t.FailNow()
	 }

	 fmt.Println("Test HelloWorlIlham Done")
}

func  TestHelloWorlMuhammad(t *testing.T){

	 result := HelloWorld("Muhammad")

	 if result != "Hello Muhammad"{
		t.Fail()
	 }

	 fmt.Println("Test HelloWorlIlham Done")
}
func  TestHelloWorlMuhammadIlham(t *testing.T){

	 result := HelloWorld("Muhammad Ilham")

	 if result != "Hello Muhammad Ilham"{
		t.Error("name should be Muhammad Ilham")
	 }

	 fmt.Println("Test HelloWorlMuhammadIlham Done")
}
func  TestHelloWorlMuhammadIlham123(t *testing.T){

	 result := HelloWorld("Muhammad Ilham 123")

	 if result != "Hello Muhammad Ilham 123"{
		t.Fatal("name should be Muhammad Ilham 123")
	 }

	 fmt.Println("Test HelloWorlMuhammadIlham Done")
}


func TestHelloWorldAssert (t *testing.T){
	var result = HelloWorld("Ilham")
	assert.Equal(t, "Hello Ilham", result, "result must be 'Hello Ilham'")
	fmt.Println("Test HelloWorlAssert Done")
}
func TestHelloWorldRequire (t *testing.T){
	var result = HelloWorld("Ilham")
	require.Equal(t, "Hello Ilham", result, "result must be 'Hello Ilham'")
	fmt.Println("Test HelloWorlAssert Done")
}

func TestSkip (t *testing.T){
	if runtime.GOOS == "windows"{
		t.Skip("cannot run on windows")
	}


	result := HelloWorld("Ilham")
	assert.Equal(t, "Hello Ilham", result, "result must be 'Hello Ilham'")
	fmt.Println("Test HelloWorlAssert Done")
}