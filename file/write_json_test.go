package file

import (
	"bytes"
	"encoding/json"
	"github.com/goinggo/tracelog"
	"io/ioutil"
	"os"

	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWriteFile(t *testing.T) {
	games := []Game {{
			AgeLimit: 16,
			Name: "God of War 4",
			ProviderName: "Sony",
		},
		{
			AgeLimit: 7,
			Name: "Crash",
			ProviderName: "Activision",
		},
	}
	
	Write(&games)

	path := "output.json"
	g, err := ioutil.ReadFile(path)
	assert.Nil(t, err, "failed reading file %v", err)


	buffers := new(bytes.Buffer)
    json.NewEncoder(buffers).Encode(games)

	assert.True(t, bytes.Equal(buffers.Bytes(), g), "written json does not match")

	os.Remove(path)
}

func TestMain(m *testing.M) {
	tracelog.Start(tracelog.LevelTrace)
	code := m.Run()
	tracelog.Stop()
	os.Exit(code)
}