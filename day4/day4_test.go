package main

import "testing"

func TestParseRoom(t *testing.T) {

	output := parseRoom("aaaaa-bbb-z-y-x-123[abxyz]")
	expected := Room{
		name:     "aaaaa-bbb-z-y-x",
		sectorId: 123,
		checksum: "abxyz",
	}

	if output != expected {
		t.Errorf("Expected %v, Got %v", expected, output)
	}
}

func TestNameDecrypted(t *testing.T) {

	input := "qzmt-zixmtkozy-ivhz-343[abc]"
	room := parseRoom(input)

	expected := "very encrypted name"

	if room.NameDecrypted() != expected {
		t.Errorf("Expected %v, Got %v", expected, room.NameDecrypted())
	}
}

func TestIsRealRoom(t *testing.T) {
	tests := []struct {
		name  string
		input Room
		want  bool
	}{
		{
			"Example 1",
			Room{
				name:     "aaaaabbbzyx",
				sectorId: 123,
				checksum: "abxyz",
			},
			true,
		},
		{
			"Example 2",
			Room{
				name:     "abcdefgh",
				sectorId: 987,
				checksum: "abcde",
			},
			true,
		},
		{
			"Example 3",
			Room{
				name:     "notarealroom",
				sectorId: 404,
				checksum: "oarel",
			},
			true,
		},
		{
			"Example 4",
			Room{
				name:     "totallyrealroom",
				sectorId: 200,
				checksum: "decoy",
			},
			false,
		},
	}

	for _, tt := range tests {
		if got := tt.input.isRealRoom(); got != tt.want {
			t.Errorf("%q. Room.IsReal() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
