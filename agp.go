// Package agp is a gene parsing package for Axie Infinity
package agp

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math"
	"strings"
)

//go:embed traits.json
var traitsJson []byte

// traitsJSON holds the content of the traits.json file.
type traitsJSON map[Class]map[PartType]map[string]map[string]string

// getTraitsJSON unmarshalls the content of the traits.json file into a traitsJSON object.
func getTraitsJSON() (traitsJSON, error) {
	var ret traitsJSON
	err := json.Unmarshal(traitsJson, &ret)
	return ret, err
}

//go:embed parts.json
var partsJson []byte

// partsJSON holds the content of the parts.json file.
type partsJSON map[string]PartGene

// getPartsJSON unmarshalls the contents of the parts.json file into a partsJSON object.
func getPartsJSON() (partsJSON, error) {
	var ret partsJSON
	err := json.Unmarshal(partsJson, &ret)
	return ret, err
}

// ParseHexDecode parses a given hex into a Gene object. This combines ParseHex and Decode into a single function.
func ParseHexDecode(hex string) (*Genes, error) {
	gbg, err := ParseHex(hex)
	if err != nil {
		return nil, err
	}
	return Decode(gbg)
}

// ParseHex converts a given hex into its binary representation and parse them into a GeneBinGroup object.
func ParseHex(hex string) (*GeneBinGroup, error) {
	// Convert the given hex into binary
	bInt, err := hexutil.DecodeBig(hex)
	if err != nil {
		return nil, err
	}
	// Fill the parsed binary to form a 256 bit string
	bStr := fmt.Sprintf("%0*s", 256, bInt.Text(2))
	// Create the GeneBinGroup object with the parsed binary string.
	return &GeneBinGroup{
		Class:    bStr[0:4],
		Region:   bStr[8:13],
		Tag:      bStr[13:18],
		BodySkin: bStr[18:22],
		Xmas:     bStr[22:34],
		Pattern:  bStr[34:52],
		Color:    bStr[52:64],
		Eyes:     bStr[64:96],
		Mouth:    bStr[96:128],
		Ears:     bStr[128:160],
		Horn:     bStr[160:192],
		Back:     bStr[192:224],
		Tail:     bStr[224:256],
	}, nil
}

// Decode parses each part of the gene binary group into its respective data to form a Gene object.
//
// If any error arises while the genes data are being decoded. The function will return a nil value along with
// the corresponding error message.
func Decode(gbg *GeneBinGroup) (*Genes, error) {
	class, err := getClass(gbg)
	if err != nil {
		return nil, err
	}
	region, err := getRegion(gbg)
	if err != nil {
		return nil, err
	}
	tag, err := getTag(gbg)
	if err != nil {
		return nil, err
	}
	bodySkin, err := getBodySkin(gbg)
	if err != nil {
		return nil, err
	}
	pattern, err := getPatternGenes(gbg)
	if err != nil {
		return nil, err
	}
	color, err := getColorGenes(gbg)
	if err != nil {
		return nil, err
	}
	eyes, err := getPart(Eyes, gbg)
	if err != nil {
		return nil, err
	}
	ears, err := getPart(Ears, gbg)
	if err != nil {
		return nil, err
	}
	horn, err := getPart(Horn, gbg)
	if err != nil {
		return nil, err
	}
	mouth, err := getPart(Mouth, gbg)
	if err != nil {
		return nil, err
	}
	back, err := getPart(Back, gbg)
	if err != nil {
		return nil, err
	}
	tail, err := getPart(Tail, gbg)
	if err != nil {
		return nil, err
	}

	genes := &Genes{*class, *region, *tag, *bodySkin, *pattern, *color, *eyes, *ears, *horn, *mouth, *back, *tail, 0}
	genes.GeneQuality = getGeneQuality(*genes)
	return genes, nil
}

var binClassMap = map[string]Class{
	"0000": Beast,
	"0001": Bug,
	"0010": Bird,
	"0011": Plant,
	"0100": Aquatic,
	"0101": Reptile,
	"1000": Mech,
	"1010": Dusk,
	"1001": Dawn,
}

// getClass extracts the gene class from the gene binary group.
// It will return a class among the 9 available Axie classes: Beast, Bug, Bird, Plant,
// Aquatic, Reptile, Mech, Dusk, and Dawn.
func getClass(gbg *GeneBinGroup) (*Class, error) {
	if len(gbg.Class) != 4 {
		return nil, errors.New("pattern binary must be of length 4")
	}
	if ret, ok := binClassMap[gbg.Class]; ok {
		return &ret, nil
	}
	return nil, errors.New("cannot recognize class")
}

var binRegionMap = map[string]Region{"00000": Global, "00001": Japan}

// GetRegion extracts the region from the gene binary group. It will either return Global or Japan.
func getRegion(gbg *GeneBinGroup) (*Region, error) {
	if len(gbg.Region) != 5 {
		return nil, errors.New("region binary must be of length 5")
	}
	if ret, ok := binRegionMap[gbg.Region]; ok {
		return &ret, nil
	}
	return nil, errors.New("cannot recognize region")
}

var binTagMap = map[string]Tag{"00000": NoTag, "00001": Origin, "00011": Meo1, "00100": Meo2}

// getTag extracts the tag from the gene binary group. It will return one among the values: Origin, Meo1, Meo2,
// an empty string (default).
func getTag(gbg *GeneBinGroup) (*Tag, error) {
	if len(gbg.Tag) != 5 {
		return nil, errors.New("tag binary must be of length 5")
	}
	if ret, ok := binTagMap[gbg.Tag]; ok {
		return &ret, nil
	}
	return nil, errors.New("cannot recognize tag")
}

var binBodySkinMap = map[string]BodySkin{"0000": DefBodySkin, "0001": Frosty}

// getBodySkin extracts the body skin from the gene binary group. It will either an Frost or an empty string (default).
func getBodySkin(gbg *GeneBinGroup) (*BodySkin, error) {
	if len(gbg.BodySkin) != 4 {
		return nil, errors.New("body skin binary must be of length 4")
	}
	if ret, ok := binBodySkinMap[gbg.BodySkin]; ok {
		return &ret, nil
	}
	return nil, errors.New("cannot recognize body skin")
}

// getPatternGenes extracts the pattern genes from the gene binary group. It will return one dominant and two recessive
// pattern genes of the Axie.
func getPatternGenes(gbg *GeneBinGroup) (*PatternGene, error) {
	if len(gbg.Pattern) != 18 {
		return nil, errors.New("pattern binary must be of length 18")
	}
	return &PatternGene{gbg.Pattern[0:6], gbg.Pattern[6:12], gbg.Pattern[12:18]}, nil
}

var classColorMap = map[Class]map[string]string{
	Beast:   {"0010": "ffec51", "0011": "ffa12a", "0100": "f0c66e", "0110": "60afce", "0000": "ffffff"},
	Bug:     {"0010": "ff7183", "0011": "ff6d61", "0100": "f74e4e", "0000": "ffffff"},
	Bird:    {"0010": "ff9ab8", "0011": "ffb4bb", "0100": "ff778e", "0000": "ffffff"},
	Plant:   {"0010": "ccef5e", "0011": "efd636", "0100": "c5ffd9", "0000": "ffffff"},
	Aquatic: {"0010": "4cffdf", "0011": "2de8f2", "0100": "759edb", "0110": "ff5a71", "0000": "ffffff"},
	Reptile: {"0010": "fdbcff", "0011": "ef93ff", "0100": "f5e1ff", "0110": "43e27d", "0000": "ffffff"},
	Mech:    {"0010": "D9D9D9", "0011": "D9D9D9", "0100": "D9D9D9", "0110": "D9D9D9", "0000": "ffffff"},
	Dusk:    {"0010": "D9D9D9", "0011": "D9D9D9", "0100": "D9D9D9", "0110": "D9D9D9", "0000": "ffffff"},
	Dawn:    {"0010": "D9D9D9", "0011": "D9D9D9", "0100": "D9D9D9", "0110": "D9D9D9", "0000": "ffffff"},
}

// getColorGenes extracts the color genes from the gene binary group. It will return one dominant and two recessive
// color genes of the Axie.
func getColorGenes(gbg *GeneBinGroup) (*ColorGene, error) {
	if len(gbg.Color) != 12 {
		return nil, errors.New("color binary must be of length 12")
	}
	class, err := getClass(gbg)
	if err != nil {
		return nil, err
	}
	return &ColorGene{classColorMap[*class][gbg.Color[0:4]], classColorMap[*class][gbg.Color[4:8]], classColorMap[*class][gbg.Color[8:12]]}, nil
}

// getPart extracts the part genes from the gene binary group. It accepts an additional parameter "partType"
// which describes which part would the function look for. It will return one dominant and two recessive genes of
// the given part.
func getPart(partType PartType, gbg *GeneBinGroup) (*Part, error) {
	// Get the binary representation of the selected part
	var partBin string
	switch partType {
	case Eyes:
		partBin = gbg.Eyes
	case Ears:
		partBin = gbg.Ears
	case Horn:
		partBin = gbg.Horn
	case Mouth:
		partBin = gbg.Mouth
	case Back:
		partBin = gbg.Back
	default:
		partBin = gbg.Tail
	}

	// Check if the binary for this part is 32 bit. Otherwise, throw an error.
	if len(partBin) != 32 {
		return nil, errors.New(fmt.Sprintf("%s binary must be of length 32", partType))
	}

	skinBin := partBin[0:2]

	// Get the dominant genes
	dClass := binClassMap[partBin[2:6]]
	dBin := partBin[6:12]
	dName, err := getPartName(dClass, partType, gbg.Region, gbg.Xmas, dBin, skinBin)
	if err != nil {
		return nil, err
	}
	d, err := getPartGene(partType, dName)
	if err != nil {
		return nil, err
	}

	// Get the recessive 1 genes
	r1Class := binClassMap[partBin[12:16]]
	r1Bin := partBin[16:22]
	r1Name, err := getPartName(r1Class, partType, gbg.Region, gbg.Xmas, r1Bin, "00")
	if err != nil {
		return nil, err
	}
	r1, err := getPartGene(partType, r1Name)
	if err != nil {
		return nil, err
	}

	// Get the recessive 2 genes
	r2Class := binClassMap[partBin[22:26]]
	r2Bin := partBin[26:32]
	r2Name, err := getPartName(r2Class, partType, gbg.Region, gbg.Xmas, r2Bin, "00")
	if err != nil {
		return nil, err
	}
	r2, err := getPartGene(partType, r2Name)
	if err != nil {
		return nil, err
	}

	return &Part{*d, *r1, *r2, skinBin == "11"}, nil
}

// getPartName finds the part name based on the parameters provided.
func getPartName(class Class, partType PartType, regionBin string, xmasBin string, partBin string, skinBin string) (string, error) {
	partSkin, err := getPartSkin(xmasBin, regionBin, skinBin)
	if err != nil {
		return "", err
	}
	traitsJson, err := getTraitsJSON()
	if err != nil {
		return "", err
	}
	if partName, ok := traitsJson[class][partType][partBin][string(partSkin)]; ok {
		return partName, nil
	}
	if partName, ok := traitsJson[class][partType][partBin][string(GlobalSkin)]; ok {
		return partName, nil
	}
	return "", errors.New(fmt.Sprintf("error finding suitable part name: %s -> %s -> %s -> %s", class, partType, partBin, partSkin))
}

// getPartGene extracts a single part gene of a given part from the gene binary group. A Part is composed of three
// part genes (d, r1, r2).
func getPartGene(partType PartType, partName string) (*PartGene, error) {
	// Set '-' as the string separator
	partName = strings.ReplaceAll(strings.ToLower(partName), " ", "-")
	partName = strings.ReplaceAll(partName, ".", "")
	// Remove hyphens from the part names
	partName = strings.ReplaceAll(partName, "'", "")
	// Form the partId (<partType>-<partName>)
	partId := fmt.Sprintf("%s-%s", partType, partName)
	partsJson, err := getPartsJSON()
	if err != nil {
		return nil, err
	}
	if partGene, ok := partsJson[partId]; ok {
		return &partGene, nil
	}
	return nil, errors.New(fmt.Sprintf("error finding suitable part gene for %s", partId))
}

var binPartSkinMap = map[string]PartSkin{
	"00000":        GlobalSkin,
	"00001":        JapanSkin,
	"010101010101": Xmas1,
	"10":           Xmas2,
	"11":           Mystic,
}

// getPartSkin extracts the skin of the given part from the gene binary group. It may only return one among the four
// values Global, Japan, Xmas, and Mystic.
func getPartSkin(xmasBin string, regionBin string, skinBin string) (PartSkin, error) {
	if len(skinBin) != 2 {
		return "", errors.New("skin binary must be of length 2")
	}
	if skin, ok := binPartSkinMap[skinBin]; ok {
		return skin, nil
	}
	if skinBin == "00" {
		if xmasBin == "010101010101" {
			return Xmas1, nil
		}
		if skin, ok := binPartSkinMap[regionBin]; ok {
			return skin, nil
		}
	}
	return "", errors.New("cannot recognize part skin")
}

// getGeneQuality computes the gene quality of the Axie.
func getGeneQuality(genes Genes) float64 {
	geneQuality := 0.0
	geneQuality += getPartQuality(genes.Class, genes.Eyes)
	geneQuality += getPartQuality(genes.Class, genes.Ears)
	geneQuality += getPartQuality(genes.Class, genes.Horn)
	geneQuality += getPartQuality(genes.Class, genes.Mouth)
	geneQuality += getPartQuality(genes.Class, genes.Back)
	geneQuality += getPartQuality(genes.Class, genes.Tail)
	return math.Round(geneQuality*100) / 100
}

// getPartQuality computes the quality of Axie's body part.
func getPartQuality(class Class, part Part) float64 {
	partQuality := 0.0
	if part.D.Class == class {
		partQuality += 76.0 / 6
	}
	if part.R1.Class == class {
		partQuality += 3
	}
	if part.R2.Class == class {
		partQuality += 1
	}
	return partQuality
}
