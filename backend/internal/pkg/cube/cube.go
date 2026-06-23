package cube

import (
	"math/rand"
)

// AddCubeRotation - records a rotation in the cube's history.
func (c *Cube) AddCubeRotation(rotation Rotation, fromScramble bool) {
	index := len(c.Rotations) // Index of the rotation to be added
	xRotation := NewCubeRotation(index, rotation, fromScramble)
	c.Rotations = append(c.Rotations, xRotation)
}

// Rotate - applies the given rotation to the cube and records it.
// Returns ErrInvalidRotation if the rotation is not a valid cube rotation.
func (c *Cube) Rotate(rotation Rotation, fromScramble bool) error {
	if !IsValidRotation(rotation) {
		return ErrInvalidRotation
	}

	switch rotation {
	case RotationF:
		c.F()
	case RotationF_:
		c.F_()
	case RotationB:
		c.B()
	case RotationB_:
		c.B_()
	case RotationU:
		c.U()
	case RotationU_:
		c.U_()
	case RotationD:
		c.D()
	case RotationD_:
		c.D_()
	case RotationL:
		c.L()
	case RotationL_:
		c.L_()
	case RotationR:
		c.R()
	case RotationR_:
		c.R_()
	}

	c.AddCubeRotation(rotation, fromScramble)

	return nil
}

// Scramble - applies random rotations to the cube and returns the applied rotations.
// If amount is <= 0, it defaults to 10
func (c *Cube) Scramble(amount int) {
	if amount <= 0 {
		amount = 10
	}

	for i := 0; i < amount; i++ {
		idx := rand.Intn(len(PossibleRotations)) // Get a random rotation index

		rotation := PossibleRotations[idx]

		_ = c.Rotate(rotation, true) // Apply rotation to the cube
	}
}

// Solved - returns true if the cube is solved.
func (c *Cube) Solved() bool {
	if c == nil {
		return false
	}

	return doCentersMatch(c.Centers, SolvedCube.Centers) &&
		doEdgesMatch(c.Edges, SolvedCube.Edges) &&
		doCornersMatch(c.Corners, SolvedCube.Corners)
}

// F - rotates the front face clockwise.
func (c *Cube) F() {
	clone, _ := c.clone()

	// EDGES
	fr := c.Edges["FR"]
	fr.Stickers[FaceR] = clone.Edges["UF"].Stickers[FaceU]
	fr.Stickers[FaceF] = clone.Edges["UF"].Stickers[FaceF]
	c.Edges["FR"] = fr

	df := c.Edges["DF"]
	df.Stickers[FaceD] = clone.Edges["FR"].Stickers[FaceR]
	df.Stickers[FaceF] = clone.Edges["FR"].Stickers[FaceF]
	c.Edges["DF"] = df

	fl := c.Edges["FL"]
	fl.Stickers[FaceL] = clone.Edges["DF"].Stickers[FaceD]
	fl.Stickers[FaceF] = clone.Edges["DF"].Stickers[FaceF]
	c.Edges["FL"] = fl

	uf := c.Edges["UF"]
	uf.Stickers[FaceU] = clone.Edges["FL"].Stickers[FaceL]
	uf.Stickers[FaceF] = clone.Edges["FL"].Stickers[FaceF]
	c.Edges["UF"] = uf

	// CORNERS
	dfr := c.Corners["DFR"]
	dfr.Stickers[FaceR] = clone.Corners["UFR"].Stickers[FaceU]
	dfr.Stickers[FaceD] = clone.Corners["UFR"].Stickers[FaceR]
	dfr.Stickers[FaceF] = clone.Corners["UFR"].Stickers[FaceF]
	c.Corners["DFR"] = dfr

	dlf := c.Corners["DLF"]
	dlf.Stickers[FaceL] = clone.Corners["DFR"].Stickers[FaceD]
	dlf.Stickers[FaceD] = clone.Corners["DFR"].Stickers[FaceR]
	dlf.Stickers[FaceF] = clone.Corners["DFR"].Stickers[FaceF]
	c.Corners["DLF"] = dlf

	ulf := c.Corners["ULF"]
	ulf.Stickers[FaceU] = clone.Corners["DLF"].Stickers[FaceL]
	ulf.Stickers[FaceL] = clone.Corners["DLF"].Stickers[FaceD]
	ulf.Stickers[FaceF] = clone.Corners["DLF"].Stickers[FaceF]
	c.Corners["ULF"] = ulf

	ufr := c.Corners["UFR"]
	ufr.Stickers[FaceU] = clone.Corners["ULF"].Stickers[FaceL]
	ufr.Stickers[FaceR] = clone.Corners["ULF"].Stickers[FaceU]
	ufr.Stickers[FaceF] = clone.Corners["ULF"].Stickers[FaceF]
	c.Corners["UFR"] = ufr
}

// F_ - rotates the front face counter-clockwise.
func (c *Cube) F_() {
	c.F()
	c.F()
	c.F()
}

// R - rotates the right face clockwise.
func (c *Cube) R() {
	clone, _ := c.clone()

	// EDGES
	ur := c.Edges["UR"]
	ur.Stickers[FaceU] = clone.Edges["FR"].Stickers[FaceF]
	ur.Stickers[FaceR] = clone.Edges["FR"].Stickers[FaceR]
	c.Edges["UR"] = ur

	br := c.Edges["BR"]
	br.Stickers[FaceB] = clone.Edges["UR"].Stickers[FaceU]
	br.Stickers[FaceR] = clone.Edges["UR"].Stickers[FaceR]
	c.Edges["BR"] = br

	dr := c.Edges["DR"]
	dr.Stickers[FaceD] = clone.Edges["BR"].Stickers[FaceB]
	dr.Stickers[FaceR] = clone.Edges["BR"].Stickers[FaceR]
	c.Edges["DR"] = dr

	fr := c.Edges["FR"]
	fr.Stickers[FaceF] = clone.Edges["DR"].Stickers[FaceD]
	fr.Stickers[FaceR] = clone.Edges["DR"].Stickers[FaceR]
	c.Edges["FR"] = fr

	// CORNERS
	urb := c.Corners["URB"]
	urb.Stickers[FaceU] = clone.Corners["UFR"].Stickers[FaceF]
	urb.Stickers[FaceB] = clone.Corners["UFR"].Stickers[FaceU]
	urb.Stickers[FaceR] = clone.Corners["UFR"].Stickers[FaceR]
	c.Corners["URB"] = urb

	drb := c.Corners["DRB"]
	drb.Stickers[FaceD] = clone.Corners["URB"].Stickers[FaceB]
	drb.Stickers[FaceB] = clone.Corners["URB"].Stickers[FaceU]
	drb.Stickers[FaceR] = clone.Corners["URB"].Stickers[FaceR]
	c.Corners["DRB"] = drb

	dfr := c.Corners["DFR"]
	dfr.Stickers[FaceD] = clone.Corners["DRB"].Stickers[FaceB]
	dfr.Stickers[FaceF] = clone.Corners["DRB"].Stickers[FaceD]
	dfr.Stickers[FaceR] = clone.Corners["DRB"].Stickers[FaceR]
	c.Corners["DFR"] = dfr

	ufr := c.Corners["UFR"]
	ufr.Stickers[FaceU] = clone.Corners["DFR"].Stickers[FaceF]
	ufr.Stickers[FaceF] = clone.Corners["DFR"].Stickers[FaceD]
	ufr.Stickers[FaceR] = clone.Corners["DFR"].Stickers[FaceR]
	c.Corners["UFR"] = ufr
}

// R_ - rotates the right face counter-clockwise.
func (c *Cube) R_() {
	c.R()
	c.R()
	c.R()
}

// U - rotates the up face clockwise.
func (c *Cube) U() {
	clone, _ := c.clone()

	// EDGES
	ul := c.Edges["UL"]
	ul.Stickers[FaceL] = clone.Edges["UF"].Stickers[FaceF]
	ul.Stickers[FaceU] = clone.Edges["UF"].Stickers[FaceU]
	c.Edges["UL"] = ul

	ub := c.Edges["UB"]
	ub.Stickers[FaceB] = clone.Edges["UL"].Stickers[FaceL]
	ub.Stickers[FaceU] = clone.Edges["UL"].Stickers[FaceU]
	c.Edges["UB"] = ub

	ur := c.Edges["UR"]
	ur.Stickers[FaceR] = clone.Edges["UB"].Stickers[FaceB]
	ur.Stickers[FaceU] = clone.Edges["UB"].Stickers[FaceU]
	c.Edges["UR"] = ur

	uf := c.Edges["UF"]
	uf.Stickers[FaceF] = clone.Edges["UR"].Stickers[FaceR]
	uf.Stickers[FaceU] = clone.Edges["UR"].Stickers[FaceU]
	c.Edges["UF"] = uf

	// CORNERS
	ubl := c.Corners["UBL"]
	ubl.Stickers[FaceB] = clone.Corners["ULF"].Stickers[FaceL]
	ubl.Stickers[FaceL] = clone.Corners["ULF"].Stickers[FaceF]
	ubl.Stickers[FaceU] = clone.Corners["ULF"].Stickers[FaceU]
	c.Corners["UBL"] = ubl

	urb := c.Corners["URB"]
	urb.Stickers[FaceR] = clone.Corners["UBL"].Stickers[FaceB]
	urb.Stickers[FaceU] = clone.Corners["UBL"].Stickers[FaceU]
	urb.Stickers[FaceB] = clone.Corners["UBL"].Stickers[FaceL]
	c.Corners["URB"] = urb

	ufr := c.Corners["UFR"]
	ufr.Stickers[FaceF] = clone.Corners["URB"].Stickers[FaceR]
	ufr.Stickers[FaceU] = clone.Corners["URB"].Stickers[FaceU]
	ufr.Stickers[FaceR] = clone.Corners["URB"].Stickers[FaceB]
	c.Corners["UFR"] = ufr

	ulf := c.Corners["ULF"]
	ulf.Stickers[FaceL] = clone.Corners["UFR"].Stickers[FaceF]
	ulf.Stickers[FaceU] = clone.Corners["UFR"].Stickers[FaceU]
	ulf.Stickers[FaceF] = clone.Corners["UFR"].Stickers[FaceR]
	c.Corners["ULF"] = ulf
}

// U_ - rotates the up face counter-clockwise.
func (c *Cube) U_() {
	c.U()
	c.U()
	c.U()
}

// L - rotates the left face clockwise.
func (c *Cube) L() {
	clone, _ := c.clone()

	// EDGES
	fl := c.Edges["FL"]
	fl.Stickers[FaceF] = clone.Edges["UL"].Stickers[FaceU]
	fl.Stickers[FaceL] = clone.Edges["UL"].Stickers[FaceL]
	c.Edges["FL"] = fl

	dl := c.Edges["DL"]
	dl.Stickers[FaceD] = clone.Edges["FL"].Stickers[FaceF]
	dl.Stickers[FaceL] = clone.Edges["FL"].Stickers[FaceL]
	c.Edges["DL"] = dl

	bl := c.Edges["BL"]
	bl.Stickers[FaceB] = clone.Edges["DL"].Stickers[FaceD]
	bl.Stickers[FaceL] = clone.Edges["DL"].Stickers[FaceL]
	c.Edges["BL"] = bl

	ul := c.Edges["UL"]
	ul.Stickers[FaceU] = clone.Edges["BL"].Stickers[FaceB]
	ul.Stickers[FaceL] = clone.Edges["BL"].Stickers[FaceL]
	c.Edges["UL"] = ul

	// CORNERS
	dlf := c.Corners["DLF"]
	dlf.Stickers[FaceD] = clone.Corners["ULF"].Stickers[FaceF]
	dlf.Stickers[FaceF] = clone.Corners["ULF"].Stickers[FaceU]
	dlf.Stickers[FaceL] = clone.Corners["ULF"].Stickers[FaceL]
	c.Corners["DLF"] = dlf

	dbl := c.Corners["DBL"]
	dbl.Stickers[FaceD] = clone.Corners["DLF"].Stickers[FaceF]
	dbl.Stickers[FaceB] = clone.Corners["DLF"].Stickers[FaceD]
	dbl.Stickers[FaceL] = clone.Corners["DLF"].Stickers[FaceL]
	c.Corners["DBL"] = dbl

	ubl := c.Corners["UBL"]
	ubl.Stickers[FaceU] = clone.Corners["DBL"].Stickers[FaceB]
	ubl.Stickers[FaceB] = clone.Corners["DBL"].Stickers[FaceD]
	ubl.Stickers[FaceL] = clone.Corners["DBL"].Stickers[FaceL]
	c.Corners["UBL"] = ubl

	ulf := c.Corners["ULF"]
	ulf.Stickers[FaceF] = clone.Corners["UBL"].Stickers[FaceU]
	ulf.Stickers[FaceU] = clone.Corners["UBL"].Stickers[FaceB]
	ulf.Stickers[FaceL] = clone.Corners["UBL"].Stickers[FaceL]
	c.Corners["ULF"] = ulf
}

// L_ - rotates the left face counter-clockwise.
func (c *Cube) L_() {
	c.L()
	c.L()
	c.L()
}

// D - rotates the down face clockwise.
func (c *Cube) D() {
	clone, _ := c.clone()

	// EDGES
	dr := c.Edges["DR"]
	dr.Stickers[FaceR] = clone.Edges["DF"].Stickers[FaceF]
	dr.Stickers[FaceD] = clone.Edges["DF"].Stickers[FaceD]
	c.Edges["DR"] = dr

	db := c.Edges["DB"]
	db.Stickers[FaceB] = clone.Edges["DR"].Stickers[FaceR]
	db.Stickers[FaceD] = clone.Edges["DR"].Stickers[FaceD]
	c.Edges["DB"] = db

	dl := c.Edges["DL"]
	dl.Stickers[FaceL] = clone.Edges["DB"].Stickers[FaceB]
	dl.Stickers[FaceD] = clone.Edges["DB"].Stickers[FaceD]
	c.Edges["DL"] = dl

	df := c.Edges["DF"]
	df.Stickers[FaceF] = clone.Edges["DL"].Stickers[FaceL]
	df.Stickers[FaceD] = clone.Edges["DL"].Stickers[FaceD]
	c.Edges["DF"] = df

	// CORNERS
	drb := c.Corners["DRB"]
	drb.Stickers[FaceR] = clone.Corners["DFR"].Stickers[FaceF]
	drb.Stickers[FaceB] = clone.Corners["DFR"].Stickers[FaceR]
	drb.Stickers[FaceD] = clone.Corners["DFR"].Stickers[FaceD]
	c.Corners["DRB"] = drb

	dbl := c.Corners["DBL"]
	dbl.Stickers[FaceB] = clone.Corners["DRB"].Stickers[FaceR]
	dbl.Stickers[FaceL] = clone.Corners["DRB"].Stickers[FaceB]
	dbl.Stickers[FaceD] = clone.Corners["DRB"].Stickers[FaceD]
	c.Corners["DBL"] = dbl

	dlf := c.Corners["DLF"]
	dlf.Stickers[FaceL] = clone.Corners["DBL"].Stickers[FaceB]
	dlf.Stickers[FaceF] = clone.Corners["DBL"].Stickers[FaceL]
	dlf.Stickers[FaceD] = clone.Corners["DBL"].Stickers[FaceD]
	c.Corners["DLF"] = dlf

	dfr := c.Corners["DFR"]
	dfr.Stickers[FaceF] = clone.Corners["DLF"].Stickers[FaceL]
	dfr.Stickers[FaceR] = clone.Corners["DLF"].Stickers[FaceF]
	dfr.Stickers[FaceD] = clone.Corners["DLF"].Stickers[FaceD]
	c.Corners["DFR"] = dfr
}

// D_ - rotates the down face counter-clockwise.
func (c *Cube) D_() {
	c.D()
	c.D()
	c.D()
}

// B - rotates the back face clockwise.
func (c *Cube) B() {
	clone, _ := c.clone()

	// EDGES
	bl := c.Edges["BL"]
	bl.Stickers[FaceL] = clone.Edges["UB"].Stickers[FaceU]
	bl.Stickers[FaceB] = clone.Edges["UB"].Stickers[FaceB]
	c.Edges["BL"] = bl

	db := c.Edges["DB"]
	db.Stickers[FaceD] = clone.Edges["BL"].Stickers[FaceL]
	db.Stickers[FaceB] = clone.Edges["BL"].Stickers[FaceB]
	c.Edges["DB"] = db

	br := c.Edges["BR"]
	br.Stickers[FaceR] = clone.Edges["DB"].Stickers[FaceD]
	br.Stickers[FaceB] = clone.Edges["DB"].Stickers[FaceB]
	c.Edges["BR"] = br

	ub := c.Edges["UB"]
	ub.Stickers[FaceU] = clone.Edges["BR"].Stickers[FaceR]
	ub.Stickers[FaceB] = clone.Edges["BR"].Stickers[FaceB]
	c.Edges["UB"] = ub

	// CORNERS
	ubl := c.Corners["UBL"]
	ubl.Stickers[FaceL] = clone.Corners["URB"].Stickers[FaceU]
	ubl.Stickers[FaceU] = clone.Corners["URB"].Stickers[FaceR]
	ubl.Stickers[FaceB] = clone.Corners["URB"].Stickers[FaceB]
	c.Corners["UBL"] = ubl

	dbl := c.Corners["DBL"]
	dbl.Stickers[FaceL] = clone.Corners["UBL"].Stickers[FaceU]
	dbl.Stickers[FaceD] = clone.Corners["UBL"].Stickers[FaceL]
	dbl.Stickers[FaceB] = clone.Corners["UBL"].Stickers[FaceB]
	c.Corners["DBL"] = dbl

	drb := c.Corners["DRB"]
	drb.Stickers[FaceR] = clone.Corners["DBL"].Stickers[FaceD]
	drb.Stickers[FaceD] = clone.Corners["DBL"].Stickers[FaceL]
	drb.Stickers[FaceB] = clone.Corners["DBL"].Stickers[FaceB]
	c.Corners["DRB"] = drb

	urb := c.Corners["URB"]
	urb.Stickers[FaceU] = clone.Corners["DRB"].Stickers[FaceR]
	urb.Stickers[FaceR] = clone.Corners["DRB"].Stickers[FaceD]
	urb.Stickers[FaceB] = clone.Corners["DRB"].Stickers[FaceB]
	c.Corners["URB"] = urb
}

// B_ - rotates the back face counter-clockwise.
func (c *Cube) B_() {
	c.B()
	c.B()
	c.B()
}
