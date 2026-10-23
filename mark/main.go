// Command mark writes the go-graphalgo mark, svg/color/go-graphalgo.svg:
//
//	cd mark && go run . > ../svg/color/go-graphalgo.svg
//
// go-graphalgo is graph algorithms: shortest paths, flows, matchings,
// isomorphism, centrality. The mark is a small graph: six vertices, the edges
// between them, and one path through the graph drawn thicker than the rest,
// the way a shortest path is highlighted on a figure.
//
// Nothing is placed by hand. The program lists the vertices as positions on
// the square and the edges as pairs of vertex numbers; the highlighted path is
// a list of vertex numbers too. Change a pair and the mark redraws itself.
//
// The rest is the family grammar every go-* mark shares: a 256-unit rounded
// square (rx 56) filled with a diagonal two-stop gradient, a white line glyph
// at stroke-width 13 with round caps and joins. logos/brandgen renders the
// format matrix from the file this writes; the kit workflow runs this again
// and fails if a byte of svg/color differs.
package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	light = "#8E9BF2" // the gradient's light stop
	dark  = "#3A3FA6" // its dark stop

	stroke    = 13.0 // the family's line width
	nodeR     = 17.0 // radius of a vertex
	pathExtra = 8.0  // how much thicker the highlighted path is drawn
)

type point struct{ x, y float64 }

// vertices, by number, on the 256-unit square.
var vertices = []point{
	{58, 70},   // 0
	{128, 44},  // 1
	{198, 84},  // 2
	{74, 156},  // 3
	{146, 134}, // 4
	{188, 202}, // 5
}

// edges, as pairs of vertex numbers.
var edges = [][2]int{
	{0, 1}, {1, 2}, {0, 3}, {1, 4}, {2, 4}, {3, 4}, {4, 5}, {2, 5},
}

// route is the highlighted path, as vertex numbers in order. Every
// consecutive pair must be an edge: main checks it.
var route = []int{0, 1, 4, 5}

func num(v float64) string {
	s := fmt.Sprintf("%.1f", math.Round(v*10)/10)
	s = strings.TrimSuffix(s, ".0")
	if s == "-0" {
		s = "0"
	}
	return s
}

func isEdge(a, b int) bool {
	for _, e := range edges {
		if (e[0] == a && e[1] == b) || (e[0] == b && e[1] == a) {
			return true
		}
	}
	return false
}

func main() {
	for i := 1; i < len(route); i++ {
		if !isEdge(route[i-1], route[i]) {
			fmt.Fprintf(os.Stderr, "mark: route step %d-%d is not an edge\n", route[i-1], route[i])
			os.Exit(1)
		}
	}

	var lines []string
	for _, e := range edges {
		a, b := vertices[e[0]], vertices[e[1]]
		lines = append(lines, fmt.Sprintf(`<path d="M%s %sL%s %s"/>`, num(a.x), num(a.y), num(b.x), num(b.y)))
	}

	var d []string
	for i, v := range route {
		cmd := "L"
		if i == 0 {
			cmd = "M"
		}
		d = append(d, cmd+num(vertices[v].x)+" "+num(vertices[v].y))
	}
	highlight := fmt.Sprintf(`<path d="%s" stroke-width="%s"/>`, strings.Join(d, ""), num(stroke+pathExtra))

	var dots []string
	for _, v := range vertices {
		dots = append(dots, fmt.Sprintf(`<circle cx="%s" cy="%s" r="%s"/>`, num(v.x), num(v.y), num(nodeR)))
	}

	fmt.Fprintf(os.Stdout, `<svg viewBox="0 0 256 256" xmlns="http://www.w3.org/2000/svg" role="img" aria-label="go-graphalgo">
  <defs><linearGradient id="gagrad" x1="0" y1="0" x2="1" y2="1"><stop offset="0" stop-color="%s"/><stop offset="1" stop-color="%s"/></linearGradient></defs>
  <rect width="256" height="256" rx="56" fill="url(#gagrad)"/>
  <g fill="none" stroke="#FFFFFF" stroke-width="%s" stroke-linecap="round" stroke-linejoin="round">
    %s
    %s
  </g>
  <g fill="#FFFFFF">
    %s
  </g>
</svg>
`, light, dark, num(stroke), strings.Join(lines, "\n    "), highlight, strings.Join(dots, "\n    "))
}
