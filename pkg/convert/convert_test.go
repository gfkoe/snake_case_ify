package convert

import "testing"

func TestConvertCamel(t *testing.T) {
	// test convert
	snake := "hello_world"

	camel_case := "helloWorld"

	if snake != Convert(camel_case) {
		t.Errorf("Expected %s, got %s", snake, Convert(camel_case))
	}
}
