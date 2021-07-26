package agp

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	hex := "0x11c642400a028ca14a428c20cc011080c61180a0820180604233082"
	want := &Genes{
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
	}
	bin, err := ParseHex(hex)
	if err != nil {
		t.Fatalf("Decode() error occured while parsing hex = %v", err)
		return
	}
	got, err := Decode(bin)
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
			got, err := GetBodySkin(tt.bin)
			if err == nil && tt.wantErr {
				t.Fatalf("GetBodySkin() expected an error")
				return
			}
			if err != nil {
				if !tt.wantErr {
					t.Fatalf("GetBodySkin() unexpected error = %v", err)
				}
				return
			}
			if !reflect.DeepEqual(*got, tt.want) {
				t.Fatalf("GetBodySkin() got = %v, want %v", *got, tt.want)
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
			got, err := GetClass(tt.bin)
			if err == nil && tt.wantErr {
				t.Fatalf("GetClass() expected an error")
				return
			}
			if err != nil {
				if !tt.wantErr {
					t.Fatalf("GetClass() unexpected error = %v", err)
				}
				return
			}
			if !reflect.DeepEqual(*got, tt.want) {
				t.Fatalf("GetClass() got = %v, want %v", *got, tt.want)
			}
		})
	}
}

func TestGetColorGenes(t *testing.T) {
	tests := []struct {
		name    string
		bin     *GeneBinGroup
		want    *ColorGene
		wantErr bool
	}{
		{"VALID_COLOR", &GeneBinGroup{Class: "0000", Color: "001000110010"}, &ColorGene{"ffec51", "ffa12a", "ffec51"}, false},
		{"INVALID_CLASS", &GeneBinGroup{Class: "1111", Color: "001000110010"}, nil, true},
		{"INVALID_COLORS", &GeneBinGroup{Class: "0000", Color: "1011010101010"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetColorGenes(tt.bin)
			if err == nil && tt.wantErr {
				t.Fatalf("GetColorGenes() expected an error")
				return
			}
			if err != nil {
				if !tt.wantErr {
					t.Fatalf("GetColorGenes() unexpected error = %v", err)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("GetColorGenes() got = %v, want %v", got, tt.want)
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
		want    *Part
		wantErr bool
	}{
		{
			"VALID_PART",
			args{Eyes, &GeneBinGroup{Region: "00000", Eyes: "00000000101000000010100011001010"}}, &Part{PartGene{"eyes-chubby", Beast, "", Eyes, "Chubby"}, PartGene{"eyes-chubby", Beast, "", Eyes, "Chubby"}, PartGene{"eyes-blossom", Plant, "", Eyes, "Blossom"}, false},
			false,
		},
		{
			"INVALID_PART",
			args{Eyes, &GeneBinGroup{Region: "00000", Eyes: "00101000101000000101100011001010"}},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPart(tt.args.partType, tt.args.gbg)
			if err == nil && tt.wantErr {
				t.Fatalf("GetPart() expected an error")
				return
			}
			if err != nil {
				if !tt.wantErr {
					t.Fatalf("GetPart() unexpected error = %v", err)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("GetPart() got = %v, want %v", got, tt.want)
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
		want    *PartGene
		wantErr bool
	}{
		{"VALID_PART_GENE", args{Ears, "Nut Cracker"}, &PartGene{"ears-nut-cracker", Beast, "", Ears, "Nut Cracker"}, false},
		{"INVALID_COMBINATION", args{Ears, "Chubby"}, nil, true},
		{"INVALID_PART_NAME", args{Ears, "Ballon Mouth"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPartGene(tt.args.partType, tt.args.partName)
			if err == nil && tt.wantErr {
				t.Fatalf("GetPartGene() expected an error")
				return
			}
			if err != nil {
				if !tt.wantErr {
					t.Fatalf("GetPartGene() unexpected error = %v", err)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("GetPartGene() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPartName(t *testing.T) {
	type args struct {
		class    Class
		partType PartType
		region   Region
		partBin  string
		skinBin  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"VALID_PART_NAME", args{Beast, Ears, Global, "001000", "00"}, "Zen", false},
		{"INVALID_PART_BIN", args{Beast, Ears, Global, "100100", "00"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPartName(tt.args.class, tt.args.partType, tt.args.region, tt.args.partBin, tt.args.skinBin)
			if err == nil && tt.wantErr {
				t.Fatalf("GetPartName() expected an error")
				return
			}
			if err != nil {
				if !tt.wantErr {
					t.Fatalf("GetPartName() unexpected error = %v", err)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("GetPartName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPartSkin(t *testing.T) {
	type args struct {
		region  Region
		skinBin string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"GLOBAL_SKIN", args{Global, "00"}, string(Global), false},
		{"XMAS_SKIN", args{Global, "10"}, "xmas", false},
		{"MYSTIC_SKIN", args{Global, "11"}, "mystic", false},
		{"INVALID_SKIN", args{Global, "01"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPartSkin(tt.args.region, tt.args.skinBin)
			if err == nil && tt.wantErr {
				t.Fatalf("GetPartSkin() expected an error")
				return
			}
			if err != nil {
				if !tt.wantErr {
					t.Fatalf("GetPartSkin() unexpected error = %v", err)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("GetPartSkin() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPatternGenes(t *testing.T) {
	tests := []struct {
		name    string
		bin     *GeneBinGroup
		want    *PatternGene
		wantErr bool
	}{
		{"VALID_PATTERN", &GeneBinGroup{Pattern: "000001000111000110"}, &PatternGene{"000001", "000111", "000110"}, false},
		{"INVALID_BINARY", &GeneBinGroup{Region: "01010101010"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPatternGenes(tt.bin)
			if err == nil && tt.wantErr {
				t.Fatalf("GetPatternGenes() expected an error")
				return
			}
			if err != nil {
				if !tt.wantErr {
					t.Fatalf("GetPatternGenes() unexpected error = %v", err)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("GetPatternGenes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRegion(t *testing.T) {
	tests := []struct {
		name    string
		bin     *GeneBinGroup
		want    Region
		wantErr bool
	}{
		{"VALID_REGION", &GeneBinGroup{Region: "00000"}, Global, false},
		{"INVALID_BINARY", &GeneBinGroup{Region: "000000"}, "", true},
		{"INVALID_REGION", &GeneBinGroup{Region: "11111"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRegion(tt.bin)
			if err == nil && tt.wantErr {
				t.Fatalf("GetRegion() expected an error")
				return
			}
			if err != nil {
				if !tt.wantErr {
					t.Fatalf("GetRegion() unexpected error = %v", err)
				}
				return
			}
			if !reflect.DeepEqual(*got, tt.want) {
				t.Fatalf("GetRegion() got = %v, want %v", *got, tt.want)
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
			got, err := GetTag(tt.bin)
			if err == nil && tt.wantErr {
				t.Fatalf("GetTag() expected an error")
				return
			}
			if err != nil {
				if !tt.wantErr {
					t.Fatalf("GetTag() unexpected error = %v", err)
				}
				return
			}
			if !reflect.DeepEqual(*got, tt.want) {
				t.Fatalf("GetTag() got = %v, want %v", *got, tt.want)
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
			got := fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s%s%s%s", gbg.Class, gbg.Reserved, gbg.Region, gbg.Tag, gbg.BodySkin, gbg.Xmas, gbg.Pattern, gbg.Color, gbg.Eyes, gbg.Mouth, gbg.Ears, gbg.Horn, gbg.Back, gbg.Tail)
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
