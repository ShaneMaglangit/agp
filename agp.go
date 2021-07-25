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

func GetTraitsJSON() (TraitsJSON, error) {
	var ret TraitsJSON
	err := json.Unmarshal(traitsJson, &ret)
	return ret, err
}

//go:embed parts.json
var partsJson []byte

type PartsJSON map[string]PartGene

func GetPartsJSON() (PartsJSON, error) {
	var ret PartsJSON
	err := json.Unmarshal(partsJson, &ret)
	return ret, err
}

func Decode(gene string) (Genes, error) {
	if len(gene) != 256 {
		return Genes{}, errors.New("gene must be 256 bit")
	}
	class, err := GetClass(gene)
	if err != nil {
		return Genes{}, err
	}
	region, err := GetRegion(gene)
	if err != nil {
		return Genes{}, err
	}
	pattern, err := GetPatternGenes(gene)
	if err != nil {
		return Genes{}, err
	}
	color, err := GetColorGenes(gene)
	if err != nil {
		return Genes{}, err
	}
	eyes, err := GetPart(Eyes, region, gene)
	if err != nil {
		return Genes{}, err
	}
	ears, err := GetPart(Ears, region, gene)
	if err != nil {
		return Genes{}, err
	}
	horn, err := GetPart(Horn, region, gene)
	if err != nil {
		return Genes{}, err
	}
	mouth, err := GetPart(Mouth, region, gene)
	if err != nil {
		return Genes{}, err
	}
	back, err := GetPart(Back, region, gene)
	if err != nil {
		return Genes{}, err
	}
	tail, err := GetPart(Tail, region, gene)
	if err != nil {
		return Genes{}, err
	}
	return Genes{class, region, pattern, color, eyes, ears, horn, mouth, back, tail}, nil
}

func ParseHex(hex string) (string, error) {
	bInt, err := hexutil.DecodeBig(hex)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%0*s", 256, bInt.Text(2)), nil
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

func GetClass(bytes string) (Class, error) {
	ret, ok := binClassMap[bytes[0:4]]
	if !ok {
		return ret, errors.New("cannot recognize class")
	}
	return ret, nil
}

var binRegionMap = map[string]Region{"00000": Global, "00001": Japan}

func GetRegion(bytes string) (Region, error) {
	ret, ok := binRegionMap[bytes[8:13]]
	if !ok {
		return ret, errors.New("cannot recognize region")
	}
	return ret, nil
}

var binTagMap = map[string]Tag{"00000": NoTag, "00001": Origin, "00010": Meo1, "00011": Meo2}

func GetTag(bytes string) (Tag, error) {
	ret, ok := binTagMap[bytes[13:18]]
	if !ok {
		return ret, errors.New("cannot recognize tag")
	}
	return ret, nil
}

var binBodySkinMap = map[string]BodySkin{"0000": DefBodySkin, "0001": Frosty}

func GetBodySkin(bytes string) (BodySkin, error) {
	ret, ok := binBodySkinMap[bytes[18:22]]
	if !ok {
		return ret, errors.New("cannot recognize body skin")
	}
	return ret, nil
}

func GetPatternGenes(bytes string) (PatternGene, error) {
	if len(bytes) != 256 {
		return PatternGene{}, errors.New("bytes must be 256 bit")
	}
	return PatternGene{bytes[22:28], bytes[28:34], bytes[34:40]}, nil
}

func GetColorGenes(bytes string) (ColorGene, error) {
	if len(bytes) != 256 {
		return ColorGene{}, errors.New("bytes must be 256 bit")
	}
	return ColorGene{bytes[40:44], bytes[44:48], bytes[48:52]}, nil
}

var offsetPartMap = map[PartType]int{Eyes: 52, Mouth: 84, Ears: 116, Horn: 148, Back: 180, Tail: 212}

func GetPart(partType PartType, region Region, bytes string) (Part, error) {
	offset := offsetPartMap[partType]

	skinBin := bytes[offset+0 : offset+2]

	dClass := binClassMap[bytes[offset+2:offset+6]]
	dBin := bytes[offset+6 : offset+12]
	dName, err := GetPartName(dClass, partType, region, dBin, skinBin)
	if err != nil {
		return Part{}, err
	}
	d, err := GetPartGene(partType, dName)
	if err != nil {
		return Part{}, err
	}

	r1Class := binClassMap[bytes[offset+12:offset+16]]
	r1Bin := bytes[offset+16 : offset+22]
	r1Name, err := GetPartName(r1Class, partType, region, r1Bin, "00")
	if err != nil {
		return Part{}, err
	}
	r1, err := GetPartGene(partType, r1Name)
	if err != nil {
		return Part{}, err
	}

	r2Class := binClassMap[bytes[offset+22:offset+26]]
	r2Bin := bytes[offset+26 : offset+32]
	r2Name, err := GetPartName(r2Class, partType, region, r2Bin, "00")
	if err != nil {
		return Part{}, err
	}
	r2, err := GetPartGene(partType, r2Name)
	if err != nil {
		return Part{}, err
	}

	return Part{d, r1, r2, skinBin == "11"}, nil
}

func GetPartName(class Class, partType PartType, region Region, partBin string, skinBin string) (string, error) {
	partSkin, err := GetPartSkin(region, skinBin)
	if err != nil {
		return "", err
	}
	traitsJson, err := GetTraitsJSON()
	if err != nil {
		return "", err
	}
	if partName, ok := traitsJson[class][partType][partBin][partSkin]; ok {
		return partName, nil
	}
	return "", errors.New("error finding suitable part name")
}

func GetPartGene(partType PartType, partName string) (PartGene, error) {
	partName = strings.ReplaceAll(strings.ToLower(partName), " ", "-")
	partName = strings.ReplaceAll(partName, "'", "")
	partId := fmt.Sprintf("%s-%s", partType, partName)
	partsJson, err := GetPartsJSON()
	if err != nil {
		return PartGene{}, err
	}
	if partGene, ok := partsJson[partId]; ok {
		return partGene, nil
	}
	return PartGene{}, errors.New("error finding suitable part name")
}

func GetPartSkin(region Region, skinBin string) (string, error) {
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
