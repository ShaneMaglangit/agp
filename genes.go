package agp

// GeneBinGroup collectively stores each part of the parsed binary representation of the genes.
type GeneBinGroup struct {
	Class    string
	Region   string
	Tag      string
	BodySkin string
	Xmas     string
	Pattern  string
	Color    string
	Eyes     string
	Ears     string
	Horn     string
	Mouth    string
	Back     string
	Tail     string
}

// Genes contains the overall data about the Axie's gene.
type Genes struct {
	Class       Class       `json:"class,omitempty"`
	Region      Region      `json:"region,omitempty"`
	Tag         Tag         `json:"tag,omitempty"`
	BodySkin    BodySkin    `json:"bodySkin,omitempty"`
	Pattern     PatternGene `json:"pattern,omitempty"`
	Color       ColorGene   `json:"color,omitempty"`
	Eyes        Part        `json:"eyes,omitempty"`
	Mouth       Part        `json:"mouth,omitempty"`
	Ears        Part        `json:"ears,omitempty"`
	Horn        Part        `json:"horn,omitempty"`
	Back        Part        `json:"back,omitempty"`
	Tail        Part        `json:"tail,omitempty"`
	GeneQuality float64     `json:"geneQuality,omitempty"`
}

// Part stores the dominant and recessive genes of an Axie's part.
type Part struct {
	D      PartGene `json:"d1,omitempty"`
	R1     PartGene `json:"r1,omitempty"`
	R2     PartGene `json:"r2,omitempty"`
	Mystic bool     `json:"mystic,omitempty"`
}

// PartGene holds the data for a single gene of an Axie's part.
type PartGene struct {
	PartId       string   `json:"partId,omitempty"`
	Class        Class    `json:"class,omitempty"`
	SpecialGenes string   `json:"specialGenes,omitempty"`
	Type         PartType `json:"type,omitempty"`
	Name         string   `json:"name,omitempty"`
}

// PatternGene stores the dominant and recessive genes of an Axie's skin pattern.
type PatternGene struct {
	D  string `json:"d,omitempty"`
	R1 string `json:"r1,omitempty"`
	R2 string `json:"r2,omitempty"`
}

// ColorGene stores the dominant and recessive genes of an Axie's color.
type ColorGene struct {
	D  string `json:"d,omitempty"`
	R1 string `json:"r1,omitempty"`
	R2 string `json:"r2,omitempty"`
}

// PartType represents each of an Axies body parts including: Eeyes, Ears, Mouth, Horn, Back, Tail.
type PartType string

const (
	Eyes  PartType = "eyes"
	Ears           = "ears"
	Mouth          = "mouth"
	Horn           = "horn"
	Back           = "back"
	Tail           = "tail"
)

// Class represents the class of a given Axie.
// A class is among these values: Beast, Bug, Bird, Plant, Aquatic, Reptile, Mech, Dusk, Dawn.
type Class string

const (
	Beast   Class = "beast"
	Bug           = "bug"
	Bird          = "bird"
	Plant         = "plant"
	Aquatic       = "aquatic"
	Reptile       = "reptile"
	Mech          = "mech"
	Dusk          = "dusk"
	Dawn          = "dawn"
)

// Region represents the region origin of a given Axie. A region can either be Global or Japan.
type Region string

const (
	Global Region = "global"
	Japan         = "japan"
)

// Tag represents the tag attached to a given Axie. Special Axies may have the title Origin, Meo1, or Meo2 tag.
type Tag string

const (
	NoTag        Tag = ""
	Agamogenesis     = "agamogenesis"
	Origin           = "origin"
	Meo1             = "meo1"
	Meo2             = "meo2"
)

// BodySkin represents the special skin of an Axie's body. This can either be none (default) or Frosty.
type BodySkin string

const (
	DefBodySkin BodySkin = ""
	Frosty               = "frosty"
)

// PartSkin represents the special skin of an Axie's part. This can be Global, Japan, Xmas, Mystic, Bionic (Agamogenesis)
type PartSkin string

const (
	GlobalSkin PartSkin = "global"
	JapanSkin           = "japan"
	Xmas1               = "xmas1"
	Xmas2               = "xmas2"
	Mystic              = "mystic"
	Bionic              = "bionic"
)
