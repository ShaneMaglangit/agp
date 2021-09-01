package agp

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	hex := "0x11c642400a028ca14a428c20cc011080c61180a0820180604233082"
	want := Genes{
		Class:    Beast,
		Region:   Global,
		Tag:      NoTag,
		BodySkin: DefBodySkin,
		Pattern:  PatternGene{"000001", "000111", "000110"},
		Color:    ColorGene{"f0c66e", "ffec51", "f0c66e"},
		Eyes: Part{
			D:  PartGene{"eyes-chubby", Beast, "", Eyes, "Chubby"},
			R1: PartGene{"eyes-chubby", Beast, "", Eyes, "Chubby"},
			R2: PartGene{"eyes-blossom", Plant, "", Eyes, "Blossom"},
		},
		Ears: Part{
			D:  PartGene{"ears-lotus", Plant, "", Ears, "Lotus"},
			R1: PartGene{"ears-nut-cracker", Beast, "", Ears, "Nut Cracker"},
			R2: PartGene{"ears-inkling", Aquatic, "", Ears, "Inkling"},
		},
		Horn: Part{
			D:  PartGene{"horn-rose-bud", Plant, "", Horn, "Rose Bud"},
			R1: PartGene{"horn-caterpillars", Bug, "", Horn, "Caterpillars"},
			R2: PartGene{"horn-dual-blade", Beast, "", Horn, "Dual Blade"},
		},
		Mouth: Part{
			D:  PartGene{"mouth-tiny-turtle", Reptile, "", Mouth, "Tiny Turtle"},
			R1: PartGene{"mouth-piranha", Aquatic, "", Mouth, "Piranha"},
			R2: PartGene{"mouth-serious", Plant, "", Mouth, "Serious"},
		},
		Back: Part{
			D:  PartGene{"back-balloon", Bird, "", Back, "Balloon"},
			R1: PartGene{"back-jaguar", Beast, "", Back, "Jaguar"},
			R2: PartGene{"back-jaguar", Beast, "", Back, "Jaguar"},
		},
		Tail: Part{
			D:  PartGene{"tail-ant", Bug, "", Tail, "Ant"},
			R1: PartGene{"tail-hot-butt", Plant, "", Tail, "Hot Butt"},
			R2: PartGene{"tail-swallow", Bird, "", Tail, "Swallow"},
		},
		GeneQuality: 23.67,
	}
	bin, err := ParseHex(hex)
	if err != nil {
		t.Fatalf("Decode() error occured while parsing hex = %v", err)
		return
	}
	got, err := Decode(&bin)
	if err != nil {
		t.Fatalf("Decode() unexpected error = %v", err)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Decode() got = %v,\nwant %v", got, want)
	}
}

func TestGetBodySkin(t *testing.T) {
	tests := []struct {
		name    string
		bin     *GeneBinGroup
		want    BodySkin
		wantErr bool
	}{
		{"VALID_BODY_SKIN", &GeneBinGroup{BodySkin: "0001"}, Frosty, false},
		{"INVALID_BODY_SKIN", &GeneBinGroup{BodySkin: "0011"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getBodySkin(tt.bin)
			if err == nil && tt.wantErr {
				t.Fatalf("getBodySkin() expected an error")
				return
			}
			if err != nil {
				if !tt.wantErr {
					t.Fatalf("getBodySkin() unexpected error = %v", err)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("getBodySkin() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetClass(t *testing.T) {
	tests := []struct {
		name    string
		bin     *GeneBinGroup
		want    Class
		wantErr bool
	}{
		{"VALID_CLASS", &GeneBinGroup{Class: "0000"}, Beast, false},
		{"INVALID_CLASS", &GeneBinGroup{Class: "1111"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getClass(tt.bin)
			if err == nil && tt.wantErr {
				t.Fatalf("getClass() expected an error")
				return
			}
			if err != nil {
				if !tt.wantErr {
					t.Fatalf("getClass() unexpected error = %v", err)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("getClass() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetColorGenes(t *testing.T) {
	tests := []struct {
		name string
		bin  *GeneBinGroup
		want ColorGene
	}{
		{"VALID_COLOR", &GeneBinGroup{Class: "0000", Color: "001000110010"}, ColorGene{"ffec51", "ffa12a", "ffec51"}},
		{"INVALID_CLASS", &GeneBinGroup{Class: "0010", Color: "001100110011"}, ColorGene{"ffb4bb", "ffb4bb", "ffb4bb"}},
		{"INVALID_COLORS", &GeneBinGroup{Class: "0100", Color: "001000110010"}, ColorGene{"4cffdf", "2de8f2", "4cffdf"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := getColorGenes(tt.bin)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("getColorGenes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPart(t *testing.T) {
	type args struct {
		partType PartType
		gbg      *GeneBinGroup
	}
	tests := []struct {
		name    string
		args    args
		want    Part
		wantErr bool
	}{
		{
			"VALID_PART",
			args{Eyes, &GeneBinGroup{Region: "00000", Eyes: "00000000101000000010100011001010"}}, Part{PartGene{"eyes-chubby", Beast, "", Eyes, "Chubby"}, PartGene{"eyes-chubby", Beast, "", Eyes, "Chubby"}, PartGene{"eyes-blossom", Plant, "", Eyes, "Blossom"}, false},
			false,
		},
		{
			"INVALID_PART",
			args{Eyes, &GeneBinGroup{Region: "00000", Eyes: "00101000101000000101100011001010"}},
			Part{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getPart(tt.args.gbg, tt.args.gbg.Eyes, Eyes)
			if err == nil && tt.wantErr {
				t.Fatalf("getPart() expected an error")
				return
			}
			if err != nil {
				if !tt.wantErr {
					t.Fatalf("getPart() unexpected error = %v", err)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("getPart() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPartGene(t *testing.T) {
	type args struct {
		partType PartType
		partName string
	}
	tests := []struct {
		name    string
		args    args
		want    PartGene
		wantErr bool
	}{
		{"VALID_PART_GENE", args{Ears, "Nut Cracker"}, PartGene{"ears-nut-cracker", Beast, "", Ears, "Nut Cracker"}, false},
		{"INVALID_COMBINATION", args{Ears, "Chubby"}, PartGene{}, true},
		{"INVALID_PART_NAME", args{Ears, "Ballon Mouth"}, PartGene{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getPartGene(tt.args.partType, tt.args.partName)
			if err == nil && tt.wantErr {
				t.Fatalf("getPartGene() expected an error")
				return
			}
			if err != nil {
				if !tt.wantErr {
					t.Fatalf("getPartGene() unexpected error = %v", err)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("getPartGene() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPartName(t *testing.T) {
	type args struct {
		class     Class
		partType  PartType
		regionBin string
		partBin   string
		partSkin  PartSkin
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"VALID_PART_NAME", args{Beast, Ears, "00000", "001000", GlobalSkin}, "Zen", false},
		{"INVALID_PART_BIN", args{Beast, Ears, "00000", "100100", GlobalSkin}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getPartName(tt.args.class, tt.args.partType, tt.args.regionBin, tt.args.partBin, tt.args.partSkin)
			if err == nil && tt.wantErr {
				t.Fatalf("getPartName() expected an error")
				return
			}
			if err != nil {
				if !tt.wantErr {
					t.Fatalf("getPartName() unexpected error = %v", err)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("getPartName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPartSkin(t *testing.T) {
	type args struct {
		regionBin string
		skinBin   string
	}
	tests := []struct {
		name    string
		args    args
		want    PartSkin
		wantErr bool
	}{
		{"GLOBAL_SKIN", args{"00000", "00"}, GlobalSkin, false},
		{"XMAS_SKIN", args{"00000", "10"}, Xmas2, false},
		{"MYSTIC_SKIN", args{"00000", "11"}, Mystic, false},
		{"INVALID_SKIN", args{"00000", "01"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getPartSkin(&GeneBinGroup{Region: tt.args.regionBin}, tt.args.skinBin)
			if err == nil && tt.wantErr {
				t.Fatalf("getPartSkin() expected an error")
				return
			}
			if err != nil {
				if !tt.wantErr {
					t.Fatalf("getPartSkin() unexpected error = %v", err)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("getPartSkin() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPatternGenes(t *testing.T) {
	tests := []struct {
		name string
		bin  *GeneBinGroup
		want PatternGene
	}{
		{"VALID_PATTERN", &GeneBinGroup{Pattern: "000001000111000110"}, PatternGene{"000001", "000111", "000110"}},
		{"INVALID_BINARY", &GeneBinGroup{Region: "01010101010"}, PatternGene{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := getPatternGenes(tt.bin)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("getPatternGenes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRegion(t *testing.T) {
	tests := []struct {
		name string
		bin  *GeneBinGroup
		want Region
	}{
		{"VALID_REGION", &GeneBinGroup{Region: "00000", Eyes: "0000", Ears: "0000", Horn: "0000", Mouth: "0000", Back: "0000", Tail: "0000"}, Global},
		{"INVALID_REGION", &GeneBinGroup{Region: "000000000000000000", Eyes: "0000", Ears: "0011", Horn: "0000", Mouth: "0000", Back: "0000", Tail: "0000"}, Japan},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := getRegion(tt.bin)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("getRegion() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTag(t *testing.T) {
	tests := []struct {
		name    string
		bin     *GeneBinGroup
		want    Tag
		wantErr bool
	}{
		{"VALID_TAG", &GeneBinGroup{Tag: "00001"}, Origin, false},
		{"INVALID_BINARY", &GeneBinGroup{Tag: "000010"}, "", true},
		{"INVALID_TAG", &GeneBinGroup{Tag: "11111"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getTag(tt.bin)
			if err == nil && tt.wantErr {
				t.Fatalf("getTag() expected an error")
				return
			}
			if err != nil {
				if !tt.wantErr {
					t.Fatalf("getTag() unexpected error = %v", err)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("getTag() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseHex(t *testing.T) {
	tests := []struct {
		name    string
		hex     string
		want    string
		wantErr bool
	}{
		{
			"VALID_HEX",
			"0x11c642400a028ca14a428c20cc011080c61180a0820180604233082",
			"0000000000000000000000000000000000000001000111000110010000100100000000001010000000101000110010100001010010100100001010001100001000001100110000000001000100001000000011000110000100011000000010100000100000100000000110000000011000000100001000110011000010000010",
			false,
		},
		{
			"INVALID_HEX",
			"0x20000000081882220c82hg288a08a22882018880c6010840c62194a",
			"",
			true,
		},
		{
			"EMPTY_HEX",
			"",
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gbg, err := ParseHex(tt.hex)
			if err == nil && tt.wantErr {
				t.Fatalf("ParseHex() expected an error")
				return
			}
			if err != nil {
				if !tt.wantErr {
					t.Fatalf("ParseHex() unexpected error = %v", err)
				}
				return
			}
			tt.want = fmt.Sprintf("%0*s", 256, tt.want)
			got := fmt.Sprintf("%s0000%s%s%s%s%s%s%s%s%s%s%s%s", gbg.Class, gbg.Region, gbg.Tag, gbg.BodySkin, gbg.Xmas, gbg.Pattern, gbg.Color, gbg.Eyes, gbg.Mouth, gbg.Ears, gbg.Horn, gbg.Back, gbg.Tail)
			if !reflect.DeepEqual(got, tt.want) {
				for i := 0; i < 256; i++ {
					if got[i] != tt.want[i] {
						t.Fatalf("ParseHex() got = %v, want %v at index %d", string(got[i]), string(tt.want[i]), i)
						return
					}
				}
			}
		})
	}
}

func TestGetGeneQuality(t *testing.T) {
	tests := []struct {
		name string
		hex  string
		want float64
	}{
		{"ZERO_QUALITY", "0x10000000080c144410a0294208a220881040080a0c24180410c3194200200904", 0},
		{"MID_QUALITY", "0xd34c44414a028c40023114400802082004130040025280200a0280a", 75.33},
		{"HIGH_QUALITY", "0x30000000041040230c4310c40c2308c20ca330ca0c6318ca0cc330cc0c2308c2", 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			genes, _ := ParseHexDecode(tt.hex)
			if got := getGeneQuality(genes); got != tt.want {
				t.Fatalf("getGeneQuality() = %v, want %v", got, tt.want)
			}
		})
	}
}
