// Package randomart generates OpenSSH style randomart images.
package randomart

import (
	"bytes"
)

// Dimensions of the generated image.
const (
	XDim = 17
	YDim = 9
)

const (
	start = -1
	end   = -2
)

// Board is a generated randomart board.
type Board struct {
	tiles [YDim][XDim]int8
	title string
}

// Generate creates a Board to represent the given data by applying the drunken
// bishop algorithm.
func Generate(data []byte, title string) Board {
	board := Board{title: title}
	var x, y int
	x = XDim / 2
	y = YDim / 2
	board.tiles[y][x] = start

	for _, b := range data {
		for s := uint(0); s < 8; s += 2 {
			d := (b >> s) & 3
			switch d {
			case 0, 1:
				// Up
				if y > 0 {
					y--
				}
			case 2, 3:
				// Down
				if y < YDim-1 {
					y++
				}
			}
			switch d {
			case 0, 2:
				// Left
				if x > 0 {
					x--
				}
			case 1, 3:
				// Right
				if x < XDim-1 {
					x++
				}
			}
			if board.tiles[y][x] >= 0 {
				board.tiles[y][x]++
			}
		}
	}
	if board.tiles[YDim/2][XDim/2] == 0 {
		board.tiles[YDim/2][XDim/2] = start
	}
	board.tiles[y][x] = end
	return board
}

// Returns the string representation of the Board, using the OpenSSH ASCII art
// character set.
func (board Board) String() string {
	var chars = []string{
		" ", ".", "o", "+", "=",
		"*", "B", "O", "X", "@",
		"%", "&", "#", "/", "^",
	}
	var buf bytes.Buffer

	if len(board.title) > 8 {
		board.title = board.title[:8]
	}

	buf.WriteString("+--[ " + board.title + " ]--")
	for i := 0; i < XDim-8-len(board.title); i++ {
		buf.WriteString("-")
	}
	buf.WriteString("+\n")

	for _, row := range board.tiles {
		buf.WriteString("|")
		for _, c := range row {
			var s string
			if c == start {
				s = "S"
			} else if c == end {
				s = "E"
			} else if int(c) < len(chars) {
				s = chars[c]
			} else {
				s = chars[len(chars)-1]
			}
			buf.WriteString(s)
		}
		buf.WriteString("|\n")
	}

	buf.WriteString("+")
	for i := 0; i < XDim; i++ {
		buf.WriteString("-")
	}
	buf.WriteString("+\n")

	return buf.String()
}
