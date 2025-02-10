package parse

import (
	"os"
	"testing"
)

func TestParseAndConvertDef(t *testing.T) {
	input := `def TestFunc() {
    rest`
	expected := "def test_func() {\n    rest\n"

	tmpFile, err := os.CreateTemp("", "testFile")
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(input)
	tmpFile.Close()

	output, err := ProcessFileContent(tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	if output != expected {
		t.Fatalf("Expected %s, got %s", expected, output)
	}
}

func TestParseAndConvertFunc(t *testing.T) {
	input := `func TestFunc() {
    rest`
	expected := "func test_func() {\n    rest\n"

	tmpFile, err := os.CreateTemp("", "testFile")
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(input)
	tmpFile.Close()

	output, err := ProcessFileContent(tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	if output != expected {
		t.Fatalf("Expected %s, got %s", expected, output)
	}
}

func TestParseAndConvertFunction(t *testing.T) {
	input := `function TestFunc() {
    rest`
	expected := "function test_func() {\n    rest\n"

	tmpFile, err := os.CreateTemp("", "testFile")
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(input)
	tmpFile.Close()

	output, err := ProcessFileContent(tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	if output != expected {
		t.Fatalf("Expected %s, got %s", expected, output)
	}
}

func TestParseAndConvertFn(t *testing.T) {
	input := `fn TestFunc() {
    rest`
	expected := "fn test_func() {\n    rest\n"

	tmpFile, err := os.CreateTemp("", "testFile")
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(input)
	tmpFile.Close()

	output, err := ProcessFileContent(tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	if output != expected {
		t.Fatalf("Expected:\n%q\nGot:\n%q", expected, output)
	}
}

func TestParseAndConvertTSTest(t *testing.T) {
	expected := `function test_function() {
  return 1;
}
`

	output, err := ProcessFileContent("../../testFiles/testInput.ts")
	if err != nil {
		t.Fatal(err)
	}

	if output != expected {
		t.Fatalf("Expected:\n%q\nGot:\n%q", expected, output)
	}
}
