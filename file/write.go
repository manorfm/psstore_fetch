package file

import (
	"bytes"
    "encoding/json"
	"github.com/goinggo/tracelog"
	"io/ioutil"
)

//Write that used to write file
func Write(games *[]Game) {

	tracelog.Info("PsStore", "Write",
		"executing writing of %d games", len(*games))

	buffers := new(bytes.Buffer)
    json.NewEncoder(buffers).Encode(games)

    ioutil.WriteFile("output.json", buffers.Bytes(), 0644)
}
