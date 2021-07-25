package agp

import (
	"testing"
)

func TestParseHex(t *testing.T) {
	hex := "0x50000000051473330425108a14a32082046508860421088c0421084604c1304c"
	want := "0101000000000000000000000000000000000101000101000111001100110011000001000010010100010000100010100001010010100011001000001000001000000100011001010000100010000110000001000010000100001000100011000000010000100001000010000100011000000100110000010011000001001100"
	got, err := ParseHex(hex)
	if err != nil {
		t.Errorf("[ParseHex] unexpected error: %v", err)
	}
	if got != want {
		t.Errorf("[ParseHex] got = %v, want = %v", got, want)
	}
}

func TestGetClass(t *testing.T) {
	var want Class
	want = Reptile
	hex, _ := ParseHex("0x50000000051473330425108a14a32082046508860421088c0421084604c1304c")
	got, err := GetClass(hex)
	if err != nil {
		t.Errorf("[GetClass] unexpected error: %v", err)
	}
	if got != want {
		t.Errorf("[GetClass] got = %v, want = %v", got, want)
	}
}

func TestGetRegion(t *testing.T) {
	var want Region
	want = Global
	hex, _ := ParseHex("0x50000000051473330425108a14a32082046508860421088c0421084604c1304c")
	got, err := GetRegion(hex)
	if err != nil {
		t.Errorf("[GetRegion] unexpected error: %v", err)
	}
	if got != want {
		t.Errorf("[GetRegion] got = %v, want = %v", got, want)
	}
}

func TestGetTag(t *testing.T) {
	var want Tag
	want = NoTag
	hex, _ := ParseHex("0x50000000051473330425108a14a32082046508860421088c0421084604c1304c")
	got, err := GetTag(hex)
	if err != nil {
		t.Errorf("[GetTag] unexpected error: %v", err)
	}
	if got != want {
		t.Errorf("[GetTag] got = %v, want = %v", got, want)
	}
}

func TestGetBodySkin(t *testing.T) {
	var want BodySkin
	want = DefBodySkin
	hex, _ := ParseHex("0x50000000051473330425108a14a32082046508860421088c0421084604c1304c")
	got, err := GetBodySkin(hex)
	if err != nil {
		t.Errorf("[GetBodySkin] unexpected error: %v", err)
	}
	if got != want {
		t.Errorf("[GetBodySkin] got = %v, want = %v", got, want)
	}
}