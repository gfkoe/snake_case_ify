package convert

import "testing"

var (
	snakeNoSpace string = "hello_world()"
	snakeSpace   string = "hello_world ()"
)

func TestConvertCamel(t *testing.T) {
	// test convert

	camelNoSpace := "helloWorld()"
	camelSpace := "helloWorld ()"

	if snakeNoSpace != Convert(camelNoSpace) {
		t.Errorf("Expected %s, got %s", snakeNoSpace, Convert(camelNoSpace))
	}
	if snakeSpace != Convert(camelSpace) {
		t.Errorf("Expected %s, got %s", snakeSpace, Convert(camelSpace))
	}
}

func TestConvertPascal(t *testing.T) {
	// test convert

	pascalNoSpace := "HelloWorld()"
	pascalSpace := "HelloWorld ()"

	if snakeNoSpace != Convert(pascalNoSpace) {
		t.Errorf("Expected %s, got %s", snakeNoSpace, Convert(pascalNoSpace))
	}
	if snakeSpace != Convert(pascalSpace) {
		t.Errorf("Expected %s, got %s", snakeSpace, Convert(pascalSpace))
	}
}

func TestConvertKebab(t *testing.T) {
	// test convert

	kebabNoSpace := "hello-world()"
	kebabSpace := "hello-world ()"

	if snakeNoSpace != Convert(kebabNoSpace) {
		t.Errorf("Expected %s, got %s", snakeNoSpace, Convert(kebabNoSpace))
	}
	if snakeSpace != Convert(kebabSpace) {
		t.Errorf("Expected %s, got %s", snakeSpace, Convert(kebabSpace))
	}
}
