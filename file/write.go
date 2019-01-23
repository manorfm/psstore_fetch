package file

import (
	"bytes"
    "encoding/json"
    "io/ioutil"
)

//Write that used to write file
func Write(games *[]Game) {

	buffers := new(bytes.Buffer)
    json.NewEncoder(buffers).Encode(games)

    ioutil.WriteFile("output.json", buffers.Bytes(), 0644)
}
