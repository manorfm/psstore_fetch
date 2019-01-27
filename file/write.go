package file

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
)

//Write that is used to write output.json file with all games fetched from api
func Write(games *[]Game) {

	log.Printf("Executing writing of %d games to output.json", len(*games))

	buffers := new(bytes.Buffer)
    json.NewEncoder(buffers).Encode(games)

    ioutil.WriteFile("output.json", buffers.Bytes(), 0644)
}
