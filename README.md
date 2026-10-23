# go-graphalgo — brand assets

Official logos for the **go-graphalgo** organization.

<img src="avatar/go-graphalgo.png" width="160" alt="go-graphalgo">

The mark is a small graph, six vertices and their edges, with one path drawn thicker than the rest, the way a shortest path is highlighted on a figure. It is written by `mark/main.go`, which lays the points out from a short
table; nothing is placed by hand. Its gradient (`#8E9BF2` to `#3A3FA6`) sets
go-graphalgo apart from the other go-scikit organizations; its dark end reads on
white at 8.60:1.

## Formats

- **SVG** — `svg/color`, `svg/white` (white, transparent), `svg/black` (black, transparent)
- **PNG** 16→1024 px, 3 variants — `png/<variant>/<size>/`
- **JPG** (color, white background) — `jpg/<size>/`
- **ICO** (Windows) — `ico/`
- **ICNS** (macOS) — `icns/`
- **GitHub avatar** 512 px — `avatar/`
- **Social preview** 1280×640 (repo banner) — `social/`

`check/kit.py` counts the whole matrix on every pull request, and the `mark`
job runs `mark/main.go` again and fails if `svg/color/go-graphalgo.svg` differs.

---
*Generated from `svg/color/go-graphalgo.svg` by the fleet's Go brand pipeline (brandgen) — rounded square badge, line glyph, color / white / black variants.*
