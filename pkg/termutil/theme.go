package termutil

import (
	"fmt"
	"image/color"
	"strconv"
)

type Colour uint8

// See https://en.wikipedia.org/wiki/ANSI_escape_code#3-bit_and_4-bit
const (
	ColourBlack Colour = iota
	ColourRed
	ColourGreen
	ColourYellow
	ColourBlue
	ColourMagenta
	ColourCyan
	ColourWhite
	ColourBrightBlack
	ColourBrightRed
	ColourBrightGreen
	ColourBrightYellow
	ColourBrightBlue
	ColourBrightMagenta
	ColourBrightCyan
	ColourBrightWhite
)

type Theme struct {
	colourMap map[Colour]color.Color
}

var (
	DefaultTheme = Theme{
		colourMap: map[Colour]color.Color{
			ColourBlack: color.RGBA{
				A: 255,
			},
			ColourRed: color.RGBA{
				R: 222, G: 56, B: 43, A: 255,
			},
			ColourGreen: color.RGBA{
				R: 57, G: 181, B: 74, A: 255,
			},
			ColourYellow: color.RGBA{
				R: 255, G: 199, B: 6, A: 255,
			},
			ColourBlue: color.RGBA{
				G: 111, B: 184, A: 255,
			},
			ColourMagenta: color.RGBA{
				R: 118, G: 38, B: 113, A: 255,
			},
			ColourCyan: color.RGBA{
				R: 44, G: 181, B: 233, A: 255,
			},
			ColourWhite: color.RGBA{
				R: 204, G: 204, B: 204, A: 255,
			},
			ColourBrightBlack: color.RGBA{
				R: 128, G: 128, B: 128, A: 255,
			},
			ColourBrightRed: color.RGBA{
				R: 255, A: 255,
			},
			ColourBrightGreen: color.RGBA{
				G: 255, A: 255,
			},
			ColourBrightYellow: color.RGBA{
				R: 255, G: 255, A: 255,
			},
			ColourBrightBlue: color.RGBA{
				B: 255, A: 255,
			},
			ColourBrightMagenta: color.RGBA{
				R: 255, B: 255, A: 255,
			},
			ColourBrightCyan: color.RGBA{
				G: 255, B: 255, A: 255,
			},
			ColourBrightWhite: color.RGBA{
				R: 255, G: 255, B: 255, A: 255,
			},
		},
	}
	map4Bit = map[uint8]Colour{
		30:  ColourBlack,
		31:  ColourRed,
		32:  ColourGreen,
		33:  ColourYellow,
		34:  ColourBlue,
		35:  ColourMagenta,
		36:  ColourCyan,
		37:  ColourWhite,
		90:  ColourBrightBlack,
		91:  ColourBrightRed,
		92:  ColourBrightGreen,
		93:  ColourBrightYellow,
		94:  ColourBrightBlue,
		95:  ColourBrightMagenta,
		96:  ColourBrightCyan,
		97:  ColourBrightWhite,
		40:  ColourBlack,
		41:  ColourRed,
		42:  ColourGreen,
		43:  ColourYellow,
		44:  ColourBlue,
		45:  ColourMagenta,
		46:  ColourCyan,
		47:  ColourWhite,
		100: ColourBrightBlack,
		101: ColourBrightRed,
		102: ColourBrightGreen,
		103: ColourBrightYellow,
		104: ColourBrightBlue,
		105: ColourBrightMagenta,
		106: ColourBrightCyan,
		107: ColourBrightWhite,
	}
)

func (t *Theme) ColourFrom4Bit(code uint8) color.Color {
	colour, ok := map4Bit[code]
	if !ok {
		colour = ColourWhite
	}
	return t.colourMap[colour]
}

func (t *Theme) ColourFrom8Bit(n string) (color.Color, error) {

	index, err := strconv.Atoi(n)
	if err != nil {
		return nil, err
	}

	if index < 16 {
		return t.colourMap[Colour(index)], nil
	}

	if index >= 232 {
		c := ((index - 232) * 0xff) / 0x18
		return color.RGBA{
			R: byte(c),
			G: byte(c),
			B: byte(c),
			A: 0xff,
		}, nil
	}

	var r, g, b byte

	index = index - 16 // 0-216

	for i := 0; i < index; i++ {
		if b == 0 {
			b = 95
		} else if b < 255 {
			b += 40
		} else {
			b = 0
			if g == 0 {
				g = 95
			} else if g < 255 {
				g += 40
			} else {
				g = 0
				if r == 0 {
					r = 95
				} else if r < 255 {
					r += 40
				} else {
					break
				}
			}
		}
	}

	return color.RGBA{
		R: r,
		G: g,
		B: b,
		A: 0xff,
	}, nil
}

func (t *Theme) ColourFrom24Bit(r, g, b string) (color.Color, error) {
	ri, err := strconv.Atoi(r)
	if err != nil {
		return nil, err
	}
	gi, err := strconv.Atoi(g)
	if err != nil {
		return nil, err
	}
	bi, err := strconv.Atoi(b)
	if err != nil {
		return nil, err
	}
	return color.RGBA{
		R: byte(ri),
		G: byte(gi),
		B: byte(bi),
		A: 0xff,
	}, nil
}

func (t *Theme) ColourFromAnsi(ansi []string, bg bool) (color.Color, error) {

	if len(ansi) == 0 {
		return nil, fmt.Errorf("invalid ansi colour code")
	}

	switch ansi[0] {
	case "2":
		if len(ansi) != 4 {
			return nil, fmt.Errorf("invalid 24-bit ansi colour code")
		}
		return t.ColourFrom24Bit(ansi[1], ansi[2], ansi[3])
	case "5":
		if len(ansi) != 2 {
			return nil, fmt.Errorf("invalid 8-bit ansi colour code")
		}
		return t.ColourFrom8Bit(ansi[1])
	default:
		return nil, fmt.Errorf("invalid ansi colour code")
	}
}

func (t *Theme) ColourToANSI(color color.Color, bg bool) string {
	code := 38
	if bg {
		code = 48
	}

	// TODO support 3/4 bit and 24 bit output for non-true color terminals

	r, g, b, _ := color.RGBA()
	return fmt.Sprintf("\x1b[%d;%d;%d;%dm", code, r, g, b)
}
