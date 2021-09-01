package agp

import (
	_ "embed"
	"encoding/json"
)

//go:embed assets/traits.json
var traitsJson []byte

// traitsJSON holds the content of the traits.json file.
type traitsJSON map[Class]map[PartType]map[string]map[string]string

// getTraitsJSON unmarshalls the content of the traits.json file into a traitsJSON object.
func getTraitsJSON() (traitsJSON, error) {
	var ret traitsJSON
	err := json.Unmarshal(traitsJson, &ret)
	return ret, err
}
