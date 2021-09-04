package agp

import (
	_ "embed"
	"encoding/json"
)

//go:embed assets/parts.json
var partsJson []byte

// partsJSON holds the content of the parts.json file.
type partsJSON map[string]PartGene

// getPartsJSON unmarshalls the contents of the parts.json file into a partsJSON object.
func getPartsJSON() (partsJSON, error) {
	var ret partsJSON
	err := json.Unmarshal(partsJson, &ret)
	return ret, err
}
