package main

type Configuration struct {
	Input string `arg:"" help:"the file to read" optional:""`

	Language   string  `help:"code language" short:"l"`
	Output     string  `help:"output of the image" short:"o" default:"out.svg"`
	Window     bool    `help:"whether to show window controls" short:"w" default:"false"`
	Border     bool    `help:"whether to add an outline to the window" short:"b" default:"false"`
	Shadow     bool    `help:"whether to add a shadow to the window" short:"s" default:"false"`
	Radius     int     `help:"amount to round the corners" short:"r" default:"0"`
	Padding    []int   `help:"padding of the window" short:"p" default:"20,40,20,20"`
	Margin     []int   `help:"margin of the window" short:"m" default:"0"`
	FontFamily string  `help:"font family" short:"f" default:"JetBrains Mono"`
	FontSize   float64 `help:"font size" default:"14"`
	LineHeight float64 `help:"line height" default:"16.8"`
}

func expandPadding(p []int) []int {
	switch len(p) {
	case 1:
		return []int{p[top], p[top], p[top], p[top]}
	case 2:
		return []int{p[top], p[right], p[top], p[right]}
	case 4:
		return []int{p[top], p[right], p[bottom], p[left]}
	default:
		return []int{0, 0, 0, 0}

	}
}

var expandMargin = expandPadding

type side int

const (
	top    side = 0
	right  side = 1
	bottom side = 2
	left   side = 3
)