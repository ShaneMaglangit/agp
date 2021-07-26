package agp

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"strings"
)

//go:embed traits.json
var traitsJson []byte

type TraitsJSON map[Class]map[PartType]map[string]map[string]string

func getTraitsJSON() (TraitsJSON, error) {
	var ret TraitsJSON
	err := json.Unmarshal(traitsJson, &ret)
	return ret, err
}

//go:embed parts.json
var partsJson []byte

type PartsJSON map[string]PartGene

func getPartsJSON() (PartsJSON, error) {
	var ret PartsJSON
	err := json.Unmarshal(partsJson, &ret)
	return ret, err
}

func ParseHex(hex string) (*GeneBinGroup, error) {
	bInt, err := hexutil.DecodeBig(hex)
	if err != nil {
		return nil, err
	}
	bStr := fmt.Sprintf("%0*s", 256, bInt.Text(2))
	return &GeneBinGroup{
		Class:    bStr[0:4],
		Reserved: bStr[4:8],
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

func Decode(gbg *GeneBinGroup) (*Genes, error) {
	class, err := GetClass(gbg)
	if err != nil {
		return nil, err
	}
	region, err := GetRegion(gbg)
	if err != nil {
		return nil, err
	}
	tag, err := GetTag(gbg)
	if err != nil {
		return nil, err
	}
	bodySkin, err := GetBodySkin(gbg)
	if err != nil {
		return nil, err
	}
	pattern, err := GetPatternGenes(gbg)
	if err != nil {
		return nil, err
	}
	color, err := GetColorGenes(gbg)
	if err != nil {
		return nil, err
	}
	eyes, err := GetPart(Eyes, gbg)
	if err != nil {
		return nil, err
	}
	ears, err := GetPart(Ears, gbg)
	if err != nil {
		return nil, err
	}
	horn, err := GetPart(Horn, gbg)
	if err != nil {
		return nil, err
	}
	mouth, err := GetPart(Mouth, gbg)
	if err != nil {
		return nil, err
	}
	back, err := GetPart(Back, gbg)
	if err != nil {
		return nil, err
	}
	tail, err := GetPart(Tail, gbg)
	if err != nil {
		return nil, err
	}
	return &Genes{*class, *region, *tag, *bodySkin, *pattern, *color, *eyes, *ears, *horn, *mouth, *back, *tail}, nil
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
	"0111": Dawn,
}

func GetClass(gbg *GeneBinGroup) (*Class, error) {
	if len(gbg.Class) != 4 {
		return nil, errors.New("pattern binary must be of length 4")
	}
	ret, ok := binClassMap[gbg.Class]
	if !ok {
		return nil, errors.New("cannot recognize class")
	}
	return &ret, nil
}

var binRegionMap = map[string]Region{"00000": Global, "00001": Japan}

func GetRegion(gbg *GeneBinGroup) (*Region, error) {
	if len(gbg.Region) != 5 {
		return nil, errors.New("region binary must be of length 5")
	}
	ret, ok := binRegionMap[gbg.Region]
	if !ok {
		return nil, errors.New("cannot recognize region")
	}
	return &ret, nil
}

var binTagMap = map[string]Tag{"00000": NoTag, "00001": Origin, "00010": Meo1, "00011": Meo2}

func GetTag(gbg *GeneBinGroup) (*Tag, error) {
	if len(gbg.Tag) != 5 {
		return nil, errors.New("tag binary must be of length 5")
	}
	ret, ok := binTagMap[gbg.Tag]
	if !ok {
		return nil, errors.New("cannot recognize tag")
	}
	return &ret, nil
}

var binBodySkinMap = map[string]BodySkin{"0000": DefBodySkin, "0001": Frosty}

func GetBodySkin(gbg *GeneBinGroup) (*BodySkin, error) {
	if len(gbg.BodySkin) != 4 {
		return nil, errors.New("body skin binary must be of length 4")
	}
	ret, ok := binBodySkinMap[gbg.BodySkin]
	if !ok {
		return nil, errors.New("cannot recognize body skin")
	}
	return &ret, nil
}

func GetPatternGenes(gbg *GeneBinGroup) (*PatternGene, error) {
	if len(gbg.Pattern) != 18 {
		return nil, errors.New("pattern binary must be of length 18")
	}
	return &PatternGene{gbg.Pattern[0:6], gbg.Pattern[6:12], gbg.Pattern[12:18]}, nil
}

var classColorMap = map[Class]map[string]string{
	Beast:   {"0010": "ffec51", "0011": "ffa12a", "0100": "f0c66e", "0110": "60afce"},
	Bug:     {"0010": "ff7183", "0011": "ff6d61", "0100": "f74e4e"},
	Bird:    {"0010": "ff9ab8", "0011": "ffb4bb", "0100": "ff778e"},
	Plant:   {"0010": "ccef5e", "0011": "efd636", "0100": "c5ffd9"},
	Aquatic: {"0010": "4cffdf", "0011": "2de8f2", "0100": "759edb", "0110": "ff5a71"},
	Reptile: {"0010": "fdbcff", "0011": "ef93ff", "0100": "f5e1ff", "0110": "43e27d"},
	Mech:    {"0010": "D9D9D9", "0011": "D9D9D9", "0100": "D9D9D9", "0110": "D9D9D9"},
	Dusk:    {"0010": "D9D9D9", "0011": "D9D9D9", "0100": "D9D9D9", "0110": "D9D9D9"},
	Dawn:    {"0010": "D9D9D9", "0011": "D9D9D9", "0100": "D9D9D9", "0110": "D9D9D9"},
}

func GetColorGenes(gbg *GeneBinGroup) (*ColorGene, error) {
	if len(gbg.Color) != 12 {
		return nil, errors.New("color binary must be of length 12")
	}
	class, err := GetClass(gbg)
	if err != nil {
		return nil, err
	}
	return &ColorGene{classColorMap[*class][gbg.Color[0:4]], classColorMap[*class][gbg.Color[4:8]], classColorMap[*class][gbg.Color[8:12]]}, nil
}

func GetPart(partType PartType, gbg *GeneBinGroup) (*Part, error) {
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

	if len(partBin) != 32 {
		return nil, errors.New(fmt.Sprintf("%s binary must be of length 32", partType))
	}

	region, err := GetRegion(gbg)
	if err != nil {
		return nil, err
	}

	skinBin := partBin[0:2]

	dClass := binClassMap[partBin[2:6]]
	dBin := partBin[6:12]
	dName, err := GetPartName(dClass, partType, *region, dBin, skinBin)
	if err != nil {
		return nil, err
	}
	d, err := GetPartGene(partType, dName)
	if err != nil {
		return nil, err
	}

	r1Class := binClassMap[partBin[12:16]]
	r1Bin := partBin[16:22]
	r1Name, err := GetPartName(r1Class, partType, *region, r1Bin, "00")
	if err != nil {
		return nil, err
	}
	r1, err := GetPartGene(partType, r1Name)
	if err != nil {
		return nil, err
	}

	r2Class := binClassMap[partBin[22:26]]
	r2Bin := partBin[26:32]
	r2Name, err := GetPartName(r2Class, partType, *region, r2Bin, "00")
	if err != nil {
		return nil, err
	}
	r2, err := GetPartGene(partType, r2Name)
	if err != nil {
		return nil, err
	}

	return &Part{*d, *r1, *r2, skinBin == "11"}, nil
}

func GetPartName(class Class, partType PartType, region Region, partBin string, skinBin string) (string, error) {
	partSkin, err := GetPartSkin(region, skinBin)
	if err != nil {
		return "", err
	}
	traitsJson, err := getTraitsJSON()
	if err != nil {
		return "", err
	}
	if partName, ok := traitsJson[class][partType][partBin][partSkin]; ok {
		return partName, nil
	}
	return "", errors.New(fmt.Sprintf("error finding suitable part name: %s -> %s -> %s -> %s", class, partType, partBin, partSkin))
}

func GetPartGene(partType PartType, partName string) (*PartGene, error) {
	partName = strings.ReplaceAll(strings.ToLower(partName), " ", "-")
	partName = strings.ReplaceAll(partName, "'", "")
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

func GetPartSkin(region Region, skinBin string) (string, error) {
	if len(skinBin) != 2 {
		return "", errors.New("skin binary must be of length 2")
	}
	switch skinBin {
	case "00":
		return string(region), nil
	case "10":
		return "xmas", nil
	case "11":
		return "mystic", nil
	default:
		return "", errors.New("cannot recognize part skin")
	}
}
