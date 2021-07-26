package agp

type GeneBinGroup struct {
	Class    string
	Region   string
	Tag      string
	BodySkin string
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
	Class   *Class       `json:"class"`
	Region  *Region      `json:"region"`
	Pattern *PatternGene `json:"pattern"`
	Color   *ColorGene
	Eyes    *Part
	Ears    *Part
	Horn    *Part
	Mouth   *Part
	Back    *Part
	Tail    *Part
}

type Part struct {
	D      *PartGene
	R1     *PartGene
	R2     *PartGene
	Mystic bool
}

type PartGene struct {
	PartId       string
	Class        *Class
	SpecialGenes *string
	Type         *PartType
	Name         string
}

type PatternGene struct {
	D  string
	R1 string
	R2 string
}

type ColorGene struct {
	D  string
	R1 string
	R2 string
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
