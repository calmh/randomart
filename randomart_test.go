package randomart

import (
	"strings"
	"testing"
)

type TestCase struct {
	title    string
	subtitle string
	data     []byte
	clines   []string
}

// Test cases taken from OpenSSH_5.9p1 output, modified header row to fix "RSA
// 2048" text centering.

var testcases = []TestCase{
	{
		"RSA 2048",
		"",
		[]byte{
			0x9b, 0x4c, 0x7b, 0xce,
			0x7a, 0xbd, 0x0a, 0x13,
			0x61, 0xfb, 0x17, 0xc2,
			0x06, 0x12, 0x0c, 0xed,
		},
		[]string{
			"+---[RSA 2048]----+",
			"|    .+.          |",
			"|      o.         |",
			"|     .. +        |",
			"|      Eo =       |",
			"|        S + .    |",
			"|       o B . .   |",
			"|        B o..    |",
			"|         *...    |",
			"|        .o+...   |",
			"+-----------------+",
		},
	}, {
		"RSA 2048",
		"",
		[]byte{
			0x30, 0xaa, 0x88, 0x72,
			0x7d, 0xc8, 0x30, 0xd0,
			0x2b, 0x99, 0xc7, 0x8f,
			0xd1, 0x86, 0x59, 0xfc,
		},
		[]string{
			"+---[RSA 2048]----+",
			"|                 |",
			"| . .             |",
			"|. . o o          |",
			"| = * o o         |",
			"|+ X + E S        |",
			"|.+ @ .           |",
			"|+ + = .          |",
			"|..   .           |",
			"|                 |",
			"+-----------------+",
		},
	}, {
		"RSA 2048",
		"SHA256",
		[]byte{
			0x30, 0xaa, 0x88, 0x72,
			0x7d, 0xc8, 0x30, 0xd0,
			0x2b, 0x99, 0xc7, 0x8f,
			0xd1, 0x86, 0x59, 0xfc,
		},
		[]string{
			"+---[RSA 2048]----+",
			"|                 |",
			"| . .             |",
			"|. . o o          |",
			"| = * o o         |",
			"|+ X + E S        |",
			"|.+ @ .           |",
			"|+ + = .          |",
			"|..   .           |",
			"|                 |",
			"+----[SHA256]-----+",
		},
	}, {
		"ED25519 256",
		"SHA256",
		[]byte{
			0xa1, 0x37, 0xe4, 0xd4,
			0xdf, 0xd2, 0xa0, 0x96,
			0x1b, 0xc6, 0xf5, 0x9f,
			0xf5, 0x34, 0x05, 0x80,
			0xa7, 0xbd, 0x8f, 0x58,
			0x3d, 0x55, 0x92, 0xff,
			0x76, 0x1e, 0x4f, 0x6e,
			0x30, 0xbb, 0x9f, 0x75,
		},
		[]string{
			"+--[ED25519 256]--+",
			"|           ..... |",
			"|         .. . o..|",
			"|        + .+o  oo|",
			"|       = o.=.=  +|",
			"|      . S B oo+o+|",
			"|       . + oo.=+X|",
			"|          .o o @E|",
			"|          . . o O|",
			"|              .=.|",
			"+----[SHA256]-----+",
		},
	},
}

func TestRandomart(t *testing.T) {
	for _, tc := range testcases {
		verify(t, tc.title, tc.subtitle, tc.data, tc.clines)
	}
}

func verify(t *testing.T, title string, subtitle string, data []byte, clines []string) {
	generated := strings.TrimSpace(GenerateSubtitled(data, title, subtitle).String())
	glines := strings.Split(generated, "\n")

	if cl, gl := len(clines), len(glines); cl != gl {
		t.Errorf("Randomart length mismatch; %d != %d", gl, cl)
	}

	for i := range clines {
		if glines[i] != clines[i] {
			t.Errorf("Line %d mismatch %q != %q", i, glines[i], clines[i])
		}
	}
}
