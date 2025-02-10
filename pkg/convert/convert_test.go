package convert

import "testing"

var snake string = "hello_world"

func TestConvertCamel(t *testing.T) {
	// test convert

	camelCase := "helloWorld"

	if snake != Convert(camelCase) {
		t.Errorf("Expected %s, got %s", snake, Convert(camelCase))
	}
}

func TestConvertPascal(t *testing.T) {
	// test convert

	pascalCase := "HelloWorld"

	if snake != Convert(pascalCase) {
		t.Errorf("Expected %s, got %s", snake, Convert(pascalCase))
	}
}

func TestConvertKebab(t *testing.T) {
	// test convert

	kebabCase := "hello-world"

	if snake != Convert(kebabCase) {
		t.Errorf("Expected %s, got %s", snake, Convert(kebabCase))
	}
}
