package versioned

import (
	"io/ioutil"
	"os"
	"testing"
)

const yamlContent = `---

version: 1.0.0`

func TestYAML(t *testing.T) {
	// prepare file
	file, err := ioutil.TempFile(os.TempDir(), "test")
	if err != nil {
		t.Fatalf("error creating temporary file: %s", err.Error())
		t.FailNow()
	}
	defer os.Remove(file.Name()) // clean up
	_, err = file.WriteString(yamlContent)
	if err != nil {
		t.Fatalf("error creating temporary file: %s", err.Error())
		t.FailNow()
	}

	// test version
	reader := NewVersionReader()
	version, err := reader.YAML.GetFromFile(file.Name())
	if err != nil {
		t.Log("unexpected exception")
		t.Fail()
	}
	if version == "" {
		t.Log("version empty")
		t.Fail()
	} else {
		if version != "1.0.0" {
			t.Logf("expected '1.0.0' got '%s'", version)
		}
	}
}
