package agp

type GeneBinGroup struct {
	Class    string
	Reserved string
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

type Genes struct {
	Class    Class       `json:"class,omitempty"`
	Region   Region      `json:"region,omitempty"`
	Tag      Tag         `json:"tag,omitempty"`
	BodySkin BodySkin    `json:"bodySkin,omitempty"`
	Pattern  PatternGene `json:"pattern,omitempty"`
	Color    ColorGene   `json:"color,omitempty"`
	Eyes     Part        `json:"eyes,omitempty"`
	Ears     Part        `json:"ears,omitempty"`
	Horn     Part        `json:"horn,omitempty"`
	Mouth    Part        `json:"mouth,omitempty"`
	Back     Part        `json:"back,omitempty"`
	Tail     Part        `json:"tail,omitempty"`
}

type Part struct {
	D      PartGene `json:"d1,omitempty"`
	R1     PartGene `json:"r1,omitempty"`
	R2     PartGene `json:"r2,omitempty"`
	Mystic bool     `json:"mystic,omitempty"`
}

type PartGene struct {
	PartId       string   `json:"partId,omitempty"`
	Class        Class    `json:"class,omitempty"`
	SpecialGenes string   `json:"specialGenes,omitempty"`
	Type         PartType `json:"type,omitempty"`
	Name         string   `json:"name,omitempty"`
}

type PatternGene struct {
	D  string `json:"d,omitempty"`
	R1 string `json:"r1,omitempty"`
	R2 string `json:"r2,omitempty"`
}

type ColorGene struct {
	D  string `json:"d,omitempty"`
	R1 string `json:"r1,omitempty"`
	R2 string `json:"r2,omitempty"`
}

type PartType string

const (
	Eyes  PartType = "eyes"
	Ears           = "ears"
	Mouth          = "mouth"
	Horn           = "horn"
	Back           = "back"
	Tail           = "tail"
)

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

type Region string

const (
	Global Region = "global"
	Japan         = "japan"
)

type Tag string

const (
	NoTag  Tag = ""
	Origin     = "origin"
	Meo1       = "meo1"
	Meo2       = "meo2"
)

type BodySkin string

const (
	DefBodySkin BodySkin = ""
	Frosty               = "frosty"
)
