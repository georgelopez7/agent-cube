package cube

import (
	"errors"
	"time"
)

type Cube struct {
	Centers   map[Face]Cubie   `json:"centers" bson:"centers"`
	Edges     map[string]Cubie `json:"edges" bson:"edges"`
	Corners   map[string]Cubie `json:"corners" bson:"corners"`
	Rotations []CubeRotation   `json:"rotations" bson:"rotations"`
}

// NewCube - creates a new cube. If cube is nil, the cube starts solved.
func NewCube() Cube {
	c, _ := SolvedCube.clone()
	return *c
}

type CubeState struct {
	Centers map[Face]Cubie   `json:"centers" bson:"centers"`
	Edges   map[string]Cubie `json:"edges" bson:"edges"`
	Corners map[string]Cubie `json:"corners" bson:"corners"`
}

// RawState - returns the raw state of the cube.
func (c *Cube) RawState() CubeState {
	return CubeState{
		Centers: c.Centers,
		Edges:   c.Edges,
		Corners: c.Corners,
	}
}

type Face string

const (
	FaceF Face = "F"
	FaceB Face = "B"
	FaceU Face = "U"
	FaceD Face = "D"
	FaceL Face = "L"
	FaceR Face = "R"
)

type FaceColor string

const (
	White  FaceColor = "white"
	Yellow FaceColor = "yellow"
	Orange FaceColor = "orange"
	Red    FaceColor = "red"
	Blue   FaceColor = "blue"
	Green  FaceColor = "green"
)

type Rotation string

const (
	RotationF  Rotation = "F"
	RotationF_ Rotation = "F'"
	RotationB  Rotation = "B"
	RotationB_ Rotation = "B'"
	RotationU  Rotation = "U"
	RotationU_ Rotation = "U'"
	RotationD  Rotation = "D"
	RotationD_ Rotation = "D'"
	RotationL  Rotation = "L"
	RotationL_ Rotation = "L'"
	RotationR  Rotation = "R"
	RotationR_ Rotation = "R'"
)

// ErrInvalidRotation - returned when an invalid rotation is provided.
var ErrInvalidRotation = errors.New("invalid rotation")

var PossibleRotations = []Rotation{
	RotationF,
	RotationF_,
	RotationB,
	RotationB_,
	RotationU,
	RotationU_,
	RotationD,
	RotationD_,
	RotationL,
	RotationL_,
	RotationR,
	RotationR_,
}

// IsValidRotation - returns true if the rotation is a valid cube rotation.
func IsValidRotation(rotation Rotation) bool {
	for _, r := range PossibleRotations {
		if r == rotation {
			return true
		}
	}

	return false
}

type CubeRotation struct {
	Rotation     Rotation  `json:"rotation" bson:"rotation"`
	Index        int       `json:"index" bson:"index"`
	FromScramble bool      `json:"from_scramble" bson:"from_scramble"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
}

func NewCubeRotation(index int, rotation Rotation, fromScramble bool) CubeRotation {
	return CubeRotation{
		Rotation:     rotation,
		Index:        index,
		FromScramble: fromScramble,
		CreatedAt:    time.Now().UTC(),
	}
}

// Stickers - maps a face to the color shown on that face of the cubie.
type Stickers map[Face]FaceColor

// Cubie - represents a single cube of the cube.
type Cubie struct {
	Stickers Stickers `json:"stickers"`
}

// SolvedCube - a solved cube
var SolvedCube = Cube{
	Centers: map[Face]Cubie{
		FaceF: {Stickers: Stickers{FaceF: Green}},
		FaceB: {Stickers: Stickers{FaceB: Blue}},
		FaceR: {Stickers: Stickers{FaceR: Red}},
		FaceL: {Stickers: Stickers{FaceL: Orange}},
		FaceU: {Stickers: Stickers{FaceU: White}},
		FaceD: {Stickers: Stickers{FaceD: Yellow}},
	},
	Edges: map[string]Cubie{
		"UF": {Stickers: Stickers{FaceU: White, FaceF: Green}},
		"UR": {Stickers: Stickers{FaceU: White, FaceR: Red}},
		"UB": {Stickers: Stickers{FaceU: White, FaceB: Blue}},
		"UL": {Stickers: Stickers{FaceU: White, FaceL: Orange}},

		"FR": {Stickers: Stickers{FaceF: Green, FaceR: Red}},
		"FL": {Stickers: Stickers{FaceF: Green, FaceL: Orange}},
		"BR": {Stickers: Stickers{FaceB: Blue, FaceR: Red}},
		"BL": {Stickers: Stickers{FaceB: Blue, FaceL: Orange}},

		"DF": {Stickers: Stickers{FaceD: Yellow, FaceF: Green}},
		"DR": {Stickers: Stickers{FaceD: Yellow, FaceR: Red}},
		"DB": {Stickers: Stickers{FaceD: Yellow, FaceB: Blue}},
		"DL": {Stickers: Stickers{FaceD: Yellow, FaceL: Orange}},
	},
	Corners: map[string]Cubie{
		"UFR": {Stickers: Stickers{FaceU: White, FaceF: Green, FaceR: Red}},
		"URB": {Stickers: Stickers{FaceU: White, FaceR: Red, FaceB: Blue}},
		"UBL": {Stickers: Stickers{FaceU: White, FaceB: Blue, FaceL: Orange}},
		"ULF": {Stickers: Stickers{FaceU: White, FaceL: Orange, FaceF: Green}},

		"DFR": {Stickers: Stickers{FaceD: Yellow, FaceF: Green, FaceR: Red}},
		"DRB": {Stickers: Stickers{FaceD: Yellow, FaceR: Red, FaceB: Blue}},
		"DBL": {Stickers: Stickers{FaceD: Yellow, FaceB: Blue, FaceL: Orange}},
		"DLF": {Stickers: Stickers{FaceD: Yellow, FaceL: Orange, FaceF: Green}},
	},
}
