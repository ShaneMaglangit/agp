package agp

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math"
	"regexp"
	"strings"
)

// ParseHexDecode parses a given 256 hex into a Gene object. This combines ParseHex and Decode into a single function.
func ParseHexDecode(hex string) (Genes, error) {
	gbg, err := ParseHex(hex)
	if err != nil {
		return Genes{}, err
	}
	return Decode(&gbg)
}

// ParseHexDecode512 parses a given 512 hex into a Gene object. This combines ParseHex512 and Decode512 into a single function.
func ParseHexDecode512(hex string) (Genes, error) {
	gbg, err := ParseHex512(hex)
	if err != nil {
		return Genes{}, err
	}
	return Decode512(&gbg)
}

// ParseHex divide bits from the 256 hex representation of the string into their respective groups.
func ParseHex(hex string) (GeneBinGroup, error) {
	var gbg GeneBinGroup
	// Convert hex into binary
	bInt, err := hexutil.DecodeBig("0x" + strings.TrimLeft(hex[2:], "0"))
	if err != nil {
		return gbg, err
	}
	// Append leading zeroes to fill the 256 bit requirement.
	bStr := fmt.Sprintf("%0*s", 256, bInt.Text(2))
	gbg.Class = bStr[0:4]
	gbg.Region = bStr[8:13]
	gbg.Tag = bStr[13:18]
	gbg.BodySkin = bStr[18:22]
	gbg.Xmas = bStr[22:34]
	gbg.Pattern = bStr[34:52]
	gbg.Color = bStr[52:64]
	gbg.Eyes = bStr[64:96]
	gbg.Mouth = bStr[96:128]
	gbg.Ears = bStr[128:160]
	gbg.Horn = bStr[160:192]
	gbg.Back = bStr[192:224]
	gbg.Tail = bStr[224:256]
	return gbg, nil
}

// hexToBin converts a given 256 bit hex into binary.
func hexToBin(hex string) (string, error) {
	// Remove leading zeroes.
	p := regexp.MustCompile("^0+")
	str := p.ReplaceAllString(hex, "")
	bInt, err := hexutil.DecodeBig("0x" + str)
	if err != nil {
		return "", err
	}
	// Append back the leading zeroes to fill the 256 bits.
	return fmt.Sprintf("%0*s", 256, bInt.Text(2)), nil
}

// ParseHex512 divide bits from the 512 hex representation of the string into their respective groups.
func ParseHex512(hex string) (GeneBinGroup, error) {
	var gbg GeneBinGroup
	// Convert first 256 bit hex into binary.
	bStrL, err := hexToBin(hex[2:][:len(hex[2:])-64])
	if err != nil {
		return gbg, err
	}
	// Convert the next 256 bit hex into binary.
	bStrR, err := hexToBin(hex[2:][len(hex[2:])-64:])
	if err != nil {
		return gbg, err
	}
	// Merged the converted binaries.
	bStr := bStrL + bStrR
	gbg.Class = bStr[0:5]
	gbg.Region = bStr[22:40]
	gbg.Tag = bStr[40:55]
	gbg.BodySkin = bStr[61:65]
	gbg.Pattern = bStr[65:92]
	gbg.Color = bStr[92:110]
	gbg.Eyes = bStr[149:192]
	gbg.Mouth = bStr[213:256]
	gbg.Ears = bStr[277:320]
	gbg.Horn = bStr[341:384]
	gbg.Back = bStr[405:448]
	gbg.Tail = bStr[469:512]
	return gbg, nil
}

// Decode parses the grouped binary and extracts the Axie information into a Gene object.
func Decode(gbg *GeneBinGroup) (Genes, error) {
	var genes Genes
	class, err := getClass(gbg)
	if err != nil {
		return genes, err
	}
	genes.Class = class
	region, err := getRegion(gbg)
	if err != nil {
		return genes, err
	}
	genes.Region = region
	tag, err := getTag(gbg)
	if err != nil {
		return genes, err
	}
	genes.Tag = tag
	bodySkin, err := getBodySkin(gbg)
	if err != nil {
		return genes, err
	}
	genes.BodySkin = bodySkin
	pattern, err := getPatternGenes(gbg)
	if err != nil {
		return genes, err
	}
	genes.Pattern = pattern
	color, err := getColorGenes(gbg)
	if err != nil {
		return genes, err
	}
	genes.Color = color
	eyes, err := getPart(gbg, gbg.Eyes, Eyes)
	if err != nil {
		return genes, err
	}
	genes.Eyes = eyes
	ears, err := getPart(gbg, gbg.Ears, Ears)
	if err != nil {
		return genes, err
	}
	genes.Ears = ears
	horn, err := getPart(gbg, gbg.Horn, Horn)
	if err != nil {
		return genes, err
	}
	genes.Horn = horn
	mouth, err := getPart(gbg, gbg.Mouth, Mouth)
	if err != nil {
		return genes, err
	}
	genes.Mouth = mouth
	back, err := getPart(gbg, gbg.Back, Back)
	if err != nil {
		return genes, err
	}
	genes.Back = back
	tail, err := getPart(gbg, gbg.Tail, Tail)
	if err != nil {
		return genes, err
	}
	genes.Tail = tail
	genes.GeneQuality = getGeneQuality(genes)
	return genes, nil
}

// Decode512 parses the grouped binary and extracts the Axie information into a Gene object.
func Decode512(gbg *GeneBinGroup) (Genes, error) {
	var genes Genes
	class, err := getClass(gbg)
	if err != nil {
		return genes, err
	}
	genes.Class = class
	region, err := getRegion(gbg)
	if err != nil {
		return genes, err
	}
	genes.Region = region
	tag, err := getTag(gbg)
	if err != nil {
		return genes, err
	}
	genes.Tag = tag
	bodySkin, err := getBodySkin(gbg)
	if err != nil {
		return genes, err
	}
	genes.BodySkin = bodySkin
	pattern, err := getPatternGenes(gbg)
	if err != nil {
		return genes, err
	}
	genes.Pattern = pattern
	color, err := getColorGenes(gbg)
	if err != nil {
		return genes, err
	}
	genes.Color = color
	eyes, err := getPart512(gbg, gbg.Eyes, Eyes)
	if err != nil {
		return genes, err
	}
	genes.Eyes = eyes
	ears, err := getPart512(gbg, gbg.Ears, Ears)
	if err != nil {
		return genes, err
	}
	genes.Ears = ears
	horn, err := getPart512(gbg, gbg.Horn, Horn)
	if err != nil {
		return genes, err
	}
	genes.Horn = horn
	mouth, err := getPart512(gbg, gbg.Mouth, Mouth)
	if err != nil {
		return genes, err
	}
	genes.Mouth = mouth
	back, err := getPart512(gbg, gbg.Back, Back)
	if err != nil {
		return genes, err
	}
	genes.Back = back
	tail, err := getPart512(gbg, gbg.Tail, Tail)
	if err != nil {
		return genes, err
	}
	genes.Tail = tail
	genes.GeneQuality = getGeneQuality(genes)
	return genes, nil
}

// binClassMap contains the details to map binary values into the class that it represents.
var binClassMap = map[string]Class{
	"0000": Beast, "0001": Bug, "0010": Bird, "0011": Plant, "0100": Aquatic, "0101": Reptile,
	"1000": Mech, "1010": Dusk, "1001": Dawn, "00000": Beast, "00001": Bug, "00010": Bird, "00011": Plant,
	"00100": Aquatic, "00101": Reptile, "10000": Mech, "10001": Dawn, "10010": Dusk,
}

// getClass parses binary values into the class that it represents.
func getClass(gbg *GeneBinGroup) (Class, error) {
	if ret, ok := binClassMap[gbg.Class]; ok {
		return ret, nil
	}
	return "", errors.New(fmt.Sprint("cannot recognize class:", gbg.Class))
}

// binRegionMap contains the details to map binary values into the region that it represents.
var binRegionMap = map[string]Region{"00000": Global, "00001": Japan}

// getRegion parses binary values into the region that it represents.
func getRegion(gbg *GeneBinGroup) (Region, error) {
	if ret, ok := binRegionMap[gbg.Region]; ok {
		return ret, nil
	}
	if len(gbg.Region) <= 4 {
		return Global, errors.New(fmt.Sprint("cannot recognize region:", gbg.Region))
	}
	if gbg.Eyes[0:4] == "0011" {
		return Japan, nil
	}
	if gbg.Ears[0:4] == "0011" {
		return Japan, nil
	}
	if gbg.Horn[0:4] == "0011" {
		return Japan, nil
	}
	if gbg.Mouth[0:4] == "0011" {
		return Japan, nil
	}
	if gbg.Back[0:4] == "0011" {
		return Japan, nil
	}
	if gbg.Tail[0:4] == "0011" {
		return Japan, nil
	}
	return Global, nil
}

// binTagMap contains the details to map binary values into the tag that it represents.
var binTagMap = map[string]Tag{
	"00000": NoTag, "00001": Origin, "00010": Agamogenesis, "00011": Meo1, "00100": Meo2,
	"000000000000000": NoTag, "000000000000001": Origin, "000000000000010": Meo1, "000000000000011": Meo2,
}

// getTag parses binary values into the Tag it represents.
func getTag(gbg *GeneBinGroup) (Tag, error) {
	if gbg.Tag == "000000000000000" {
		eyesBionic, _ := getPartSkin(gbg, gbg.Eyes[0:4])
		earsBionic, _ := getPartSkin(gbg, gbg.Ears[0:4])
		hornBionic, _ := getPartSkin(gbg, gbg.Horn[0:4])
		mouthBionic, _ := getPartSkin(gbg, gbg.Mouth[0:4])
		backBionic, _ := getPartSkin(gbg, gbg.Back[0:4])
		tailBionic, _ := getPartSkin(gbg, gbg.Tail[0:4])
		if eyesBionic == Bionic || earsBionic == Bionic || hornBionic == Bionic ||
			mouthBionic == Bionic || backBionic == Bionic || tailBionic == Bionic {
			return Agamogenesis, nil
		}
	}
	if ret, ok := binTagMap[gbg.Tag]; ok {
		return ret, nil
	}
	return NoTag, errors.New(fmt.Sprint("cannot recognize tag:", gbg.Tag))
}

// binBodySkinMap contains the details to map binary values into the body skin that it represents.
var binBodySkinMap = map[string]BodySkin{"0000": DefBodySkin, "0001": Frosty}

// getBodySkin parses binary values into the BodySkin it represents.
func getBodySkin(gbg *GeneBinGroup) (BodySkin, error) {
	if ret, ok := binBodySkinMap[gbg.BodySkin]; ok {
		return ret, nil
	}
	return DefBodySkin, errors.New(fmt.Sprint("cannot recognize body skin", gbg.Tag))
}

// getPatternGenes parses binary values into the patterns that they represent.
func getPatternGenes(gbg *GeneBinGroup) (PatternGene, error) {
	bSize := len(gbg.Pattern) / 3
	return PatternGene{gbg.Pattern[0:bSize], gbg.Pattern[bSize : bSize*2], gbg.Pattern[bSize*2 : bSize*3]}, nil
}

// classColorMap contains the details to map binary values into the class colors it represents.
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

// getColorGenes parses binary values into the colors that they represent.
func getColorGenes(gbg *GeneBinGroup) (ColorGene, error) {
	bSize := len(gbg.Color) / 3
	class, err := getClass(gbg)
	if err != nil {
		return ColorGene{}, err
	}
	d := gbg.Color[0:bSize]
	r1 := gbg.Color[bSize : bSize*2]
	r2 := gbg.Color[bSize*2 : bSize*3]
	return ColorGene{
		classColorMap[class][d[len(d)-4:]],
		classColorMap[class][r1[len(r1)-4:]],
		classColorMap[class][r2[len(r2)-4:]],
	}, nil
}

// getPart parses binary values into the set of part genes that they represent.
func getPart(gbg *GeneBinGroup, partBin string, partType PartType) (Part, error) {
	var part Part
	dClass := binClassMap[partBin[2:6]]
	dBin := partBin[6:12]
	dSkin, err := getPartSkin(gbg, partBin[0:2])
	if err != nil {
		return part, err
	}
	dName, err := getPartName(dClass, partType, gbg.Region, dBin, dSkin)
	if err != nil {
		return part, err
	}
	part.D, err = getPartGene(partType, dName)
	if err != nil {
		return part, err
	}

	r1Class := binClassMap[partBin[12:16]]
	r1Bin := partBin[16:22]
	r1Name, err := getPartName(r1Class, partType, gbg.Region, r1Bin, "00")
	if err != nil {
		return part, err
	}
	part.R1, err = getPartGene(partType, r1Name)
	if err != nil {
		return part, err
	}

	r2Class := binClassMap[partBin[22:26]]
	r2Bin := partBin[26:32]
	r2Name, err := getPartName(r2Class, partType, gbg.Region, r2Bin, "00")
	if err != nil {
		return part, err
	}
	part.R2, err = getPartGene(partType, r2Name)
	if err != nil {
		return part, err
	}

	part.Mystic = dSkin == Mystic || partBin[0:2] == "0001"
	return part, nil
}

// getPart512 parses binary values into the set of part genes that they represent.
func getPart512(gbg *GeneBinGroup, partBin string, partType PartType) (Part, error) {
	var part Part
	dClass := binClassMap[partBin[4:9]]
	dBin := partBin[11:17]
	dSkin, err := getPartSkin(gbg, partBin[0:4])
	if err != nil {
		return part, err
	}
	dName, err := getPartName(dClass, partType, gbg.Region, dBin, dSkin)
	if err != nil {
		return part, err
	}
	part.D, err = getPartGene(partType, dName)
	if err != nil {
		return part, err
	}

	r1Class := binClassMap[partBin[17:22]]
	r1Bin := partBin[24:30]
	r1Name, err := getPartName(r1Class, partType, gbg.Region, r1Bin, dSkin)
	if err != nil {
		return part, err
	}
	part.R1, err = getPartGene(partType, r1Name)
	if err != nil {
		return part, err
	}

	r2Class := binClassMap[partBin[30:35]]
	r2Bin := partBin[37:43]
	r2Name, err := getPartName(r2Class, partType, gbg.Region, r2Bin, dSkin)
	if err != nil {
		return part, err
	}
	part.R2, err = getPartGene(partType, r2Name)
	if err != nil {
		return part, err
	}

	part.Mystic = dSkin == Mystic || partBin[0:4] == "0001"
	return part, nil
}

// getPartName parses binary values into the part name that they represent.
func getPartName(class Class, partType PartType, regionBin string, partBin string, skin PartSkin) (string, error) {
	traitsJson, err := getTraitsJSON()
	if err != nil {
		return "", err
	}
	part, ok := traitsJson[class][partType][partBin]
	if !ok {
		return "", errors.New(fmt.Sprint("cannot recognize part name:", partType, regionBin, partBin))
	}
	if partName := part[string(skin)]; partName != "" {
		return partName, nil
	}
	if partName := part[string(Global)]; partName != "" {
		return partName, nil
	}
	return "", errors.New(fmt.Sprint("cannot recognize part name:", partType, regionBin, partBin))
}

// getPartGene parses binary values and extract the part information that it represents.
func getPartGene(partType PartType, partName string) (PartGene, error) {
	partName = strings.ReplaceAll(strings.ToLower(partName), " ", "-")
	partName = strings.ReplaceAll(partName, ".", "")
	partName = strings.ReplaceAll(partName, "'", "")
	partId := fmt.Sprintf("%s-%s", partType, partName)
	partsJson, err := getPartsJSON()
	if err != nil {
		return PartGene{}, err
	}
	if partGene, ok := partsJson[partId]; ok {
		return partGene, nil
	}
	return PartGene{}, errors.New(fmt.Sprint("cannot recognize part:", partId))
}

// binPartSkinMap contains the details to map binary values into the part skin that it represents.
var binPartSkinMap = map[string]PartSkin{
	"00000":        GlobalSkin,
	"00001":        JapanSkin,
	"010101010101": Xmas1,
	"10":           Xmas2,
	"11":           Mystic,
	"01":           Bionic,
	"0000":         GlobalSkin,
	"0001":         Mystic,
	"0011":         Japan,
	"0100":         Xmas1,
	"0101":         Xmas2,
	"0010":         Bionic,
}

// getPartSkin parses binary values and extract the part skin that it represents.
func getPartSkin(gbg *GeneBinGroup, skinBin string) (PartSkin, error) {
	partSkin := binPartSkinMap[skinBin]
	if skinBin == "00" {
		if gbg.Xmas == "010101010101" {
			partSkin = Xmas1
		} else {
			partSkin = binPartSkinMap[gbg.Region]
		}
	}
	if partSkin == "" {
		return partSkin, errors.New(fmt.Sprint("cannot recognize part skin:", skinBin, gbg.Xmas))
	}
	return partSkin, nil
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
