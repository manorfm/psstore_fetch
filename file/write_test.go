package file

import (
	"testing"
	"io/ioutil"
	"bytes"
	"encoding/json"
	"os"
)

func TestWriteFile(t *testing.T) {
	games := []Game { Game {
			AgeLimit: 16,
			Name: "God of War 4",
			ProviderName: "Sony",
		},
		Game {
			AgeLimit: 7,
			Name: "Crash",
			ProviderName: "Activision",
		},
	}
	
	Write(&games)

	path := "output.json"
	g, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatalf("failed reading file: %s", err)
	}

	buffers := new(bytes.Buffer)
    json.NewEncoder(buffers).Encode(games)
	
	if !bytes.Equal(buffers.Bytes(), g) {
		t.Errorf("written json does not match")
	}

	os.Remove(path)
}