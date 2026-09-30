package cube

import "maps"

// clone - returns a deep copy of the cube.
func (c *Cube) clone() (*Cube, error) {
	out := &Cube{
		Centers:   make(map[Face]Cubie, len(c.Centers)),
		Edges:     make(map[string]Cubie, len(c.Edges)),
		Corners:   make(map[string]Cubie, len(c.Corners)),
		Rotations: make([]CubeRotation, len(c.Rotations)),
	}

	for k, v := range c.Centers {
		out.Centers[k] = cloneCubie(v)
	}

	for k, v := range c.Edges {
		out.Edges[k] = cloneCubie(v)
	}

	for k, v := range c.Corners {
		out.Corners[k] = cloneCubie(v)
	}

	copy(out.Rotations, c.Rotations)

	return out, nil
}

// cloneCubie - returns a deep copy of a single cubie.
func cloneCubie(c Cubie) Cubie {
	out := Cubie{
		Stickers: make(Stickers, len(c.Stickers)),
	}

	maps.Copy(out.Stickers, c.Stickers)

	return out
}

// doCentersMatch - returns true when the two center maps contain the same faces with the same stickers.
func doCentersMatch(a, b map[Face]Cubie) bool {
	if len(a) != len(b) {
		return false
	}

	for k, va := range a {
		vb, ok := b[k]
		if !ok {
			return false
		}

		if !stickersEqual(va.Stickers, vb.Stickers) {
			return false
		}
	}

	return true
}

// doEdgesMatch - returns true when the two edge maps contain the same keys with the same stickers.
func doEdgesMatch(a, b map[string]Cubie) bool {
	if len(a) != len(b) {
		return false
	}

	for k, va := range a {
		vb, ok := b[k]
		if !ok {
			return false
		}

		if !stickersEqual(va.Stickers, vb.Stickers) {
			return false
		}
	}

	return true
}

// doCornersMatch - returns true when the two corner maps contain the same keys with the same stickers.
func doCornersMatch(a, b map[string]Cubie) bool {
	if len(a) != len(b) {
		return false
	}

	for k, va := range a {
		vb, ok := b[k]
		if !ok {
			return false
		}

		if !stickersEqual(va.Stickers, vb.Stickers) {
			return false
		}
	}

	return true
}

// stickersEqual - returns true when the two sticker maps contain the same faces with the same colors.
func stickersEqual(a Stickers, b Stickers) bool {
	if len(a) != len(b) {
		return false
	}

	for k, va := range a {
		vb, ok := b[k]
		if !ok || va != vb {
			return false
		}
	}

	return true
}
