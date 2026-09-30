package cube_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"agent-cube/internal/pkg/cube"
)

func Test_Solved(t *testing.T) {
	t.Run("new cube is solved", func(t *testing.T) {
		c := cube.NewCube()
		require.True(t, c.Solved())
	})

	t.Run("rotated cube is not solved", func(t *testing.T) {
		c := cube.NewCube()
		c.F()
		require.False(t, c.Solved())
	})

	t.Run("full rotation returns to solved", func(t *testing.T) {
		c := cube.NewCube()
		c.F()
		c.F()
		c.F()
		c.F()
		require.True(t, c.Solved())
	})
}

func Test_Rotations(t *testing.T) {
	t.Run("F rotation", func(t *testing.T) {
		c := cube.NewCube()
		c.F()

		// EDGES
		require.Equal(t, cube.Orange, c.Edges["UF"].Stickers[cube.FaceU])
		require.Equal(t, cube.Green, c.Edges["UF"].Stickers[cube.FaceF])

		require.Equal(t, cube.White, c.Edges["FR"].Stickers[cube.FaceR])
		require.Equal(t, cube.Green, c.Edges["FR"].Stickers[cube.FaceF])

		require.Equal(t, cube.Red, c.Edges["DF"].Stickers[cube.FaceD])
		require.Equal(t, cube.Green, c.Edges["DF"].Stickers[cube.FaceF])

		require.Equal(t, cube.Yellow, c.Edges["FL"].Stickers[cube.FaceL])
		require.Equal(t, cube.Green, c.Edges["FL"].Stickers[cube.FaceF])

		// CORNERS
		require.Equal(t, cube.Orange, c.Corners["UFR"].Stickers[cube.FaceU])
		require.Equal(t, cube.White, c.Corners["UFR"].Stickers[cube.FaceR])
		require.Equal(t, cube.Green, c.Corners["UFR"].Stickers[cube.FaceF])

		require.Equal(t, cube.Red, c.Corners["DFR"].Stickers[cube.FaceD])
		require.Equal(t, cube.White, c.Corners["DFR"].Stickers[cube.FaceR])
		require.Equal(t, cube.Green, c.Corners["DFR"].Stickers[cube.FaceF])

		require.Equal(t, cube.Red, c.Corners["DLF"].Stickers[cube.FaceD])
		require.Equal(t, cube.Yellow, c.Corners["DLF"].Stickers[cube.FaceL])
		require.Equal(t, cube.Green, c.Corners["DLF"].Stickers[cube.FaceF])

		require.Equal(t, cube.Orange, c.Corners["ULF"].Stickers[cube.FaceU])
		require.Equal(t, cube.Yellow, c.Corners["ULF"].Stickers[cube.FaceL])
		require.Equal(t, cube.Green, c.Corners["ULF"].Stickers[cube.FaceF])
	})

	t.Run("R rotation", func(t *testing.T) {
		c := cube.NewCube()
		c.R()

		require.Equal(t, cube.Green, c.Edges["UR"].Stickers[cube.FaceU])
		require.Equal(t, cube.Red, c.Edges["UR"].Stickers[cube.FaceR])

		require.Equal(t, cube.White, c.Edges["BR"].Stickers[cube.FaceB])
		require.Equal(t, cube.Red, c.Edges["BR"].Stickers[cube.FaceR])

		require.Equal(t, cube.Blue, c.Edges["DR"].Stickers[cube.FaceD])
		require.Equal(t, cube.Red, c.Edges["DR"].Stickers[cube.FaceR])

		require.Equal(t, cube.Yellow, c.Edges["FR"].Stickers[cube.FaceF])
		require.Equal(t, cube.Red, c.Edges["FR"].Stickers[cube.FaceR])

		require.Equal(t, cube.Green, c.Corners["URB"].Stickers[cube.FaceU])
		require.Equal(t, cube.White, c.Corners["URB"].Stickers[cube.FaceB])
		require.Equal(t, cube.Red, c.Corners["URB"].Stickers[cube.FaceR])

		require.Equal(t, cube.Blue, c.Corners["DRB"].Stickers[cube.FaceD])
		require.Equal(t, cube.White, c.Corners["DRB"].Stickers[cube.FaceB])
		require.Equal(t, cube.Red, c.Corners["DRB"].Stickers[cube.FaceR])

		require.Equal(t, cube.Blue, c.Corners["DFR"].Stickers[cube.FaceD])
		require.Equal(t, cube.Yellow, c.Corners["DFR"].Stickers[cube.FaceF])
		require.Equal(t, cube.Red, c.Corners["DFR"].Stickers[cube.FaceR])

		require.Equal(t, cube.Green, c.Corners["UFR"].Stickers[cube.FaceU])
		require.Equal(t, cube.Yellow, c.Corners["UFR"].Stickers[cube.FaceF])
		require.Equal(t, cube.Red, c.Corners["UFR"].Stickers[cube.FaceR])
	})

	t.Run("U rotation", func(t *testing.T) {
		c := cube.NewCube()
		c.U()

		// EDGES
		require.Equal(t, cube.Green, c.Edges["UL"].Stickers[cube.FaceL])
		require.Equal(t, cube.White, c.Edges["UL"].Stickers[cube.FaceU])

		require.Equal(t, cube.Orange, c.Edges["UB"].Stickers[cube.FaceB])
		require.Equal(t, cube.White, c.Edges["UB"].Stickers[cube.FaceU])

		require.Equal(t, cube.Blue, c.Edges["UR"].Stickers[cube.FaceR])
		require.Equal(t, cube.White, c.Edges["UR"].Stickers[cube.FaceU])

		require.Equal(t, cube.Red, c.Edges["UF"].Stickers[cube.FaceF])
		require.Equal(t, cube.White, c.Edges["UF"].Stickers[cube.FaceU])

		// CORNERS
		require.Equal(t, cube.Orange, c.Corners["UBL"].Stickers[cube.FaceB])
		require.Equal(t, cube.Green, c.Corners["UBL"].Stickers[cube.FaceL])
		require.Equal(t, cube.White, c.Corners["UBL"].Stickers[cube.FaceU])

		require.Equal(t, cube.Blue, c.Corners["URB"].Stickers[cube.FaceR])
		require.Equal(t, cube.Orange, c.Corners["URB"].Stickers[cube.FaceB])
		require.Equal(t, cube.White, c.Corners["URB"].Stickers[cube.FaceU])

		require.Equal(t, cube.Red, c.Corners["UFR"].Stickers[cube.FaceF])
		require.Equal(t, cube.Blue, c.Corners["UFR"].Stickers[cube.FaceR])
		require.Equal(t, cube.White, c.Corners["UFR"].Stickers[cube.FaceU])

		require.Equal(t, cube.Green, c.Corners["ULF"].Stickers[cube.FaceL])
		require.Equal(t, cube.Red, c.Corners["ULF"].Stickers[cube.FaceF])
		require.Equal(t, cube.White, c.Corners["ULF"].Stickers[cube.FaceU])
	})

	t.Run("L rotation", func(t *testing.T) {
		c := cube.NewCube()
		c.L()

		// EDGES
		require.Equal(t, cube.White, c.Edges["FL"].Stickers[cube.FaceF])
		require.Equal(t, cube.Orange, c.Edges["FL"].Stickers[cube.FaceL])

		require.Equal(t, cube.Green, c.Edges["DL"].Stickers[cube.FaceD])
		require.Equal(t, cube.Orange, c.Edges["DL"].Stickers[cube.FaceL])

		require.Equal(t, cube.Yellow, c.Edges["BL"].Stickers[cube.FaceB])
		require.Equal(t, cube.Orange, c.Edges["BL"].Stickers[cube.FaceL])

		require.Equal(t, cube.Blue, c.Edges["UL"].Stickers[cube.FaceU])
		require.Equal(t, cube.Orange, c.Edges["UL"].Stickers[cube.FaceL])

		// CORNERS
		require.Equal(t, cube.Green, c.Corners["DLF"].Stickers[cube.FaceD])
		require.Equal(t, cube.White, c.Corners["DLF"].Stickers[cube.FaceF])
		require.Equal(t, cube.Orange, c.Corners["DLF"].Stickers[cube.FaceL])

		require.Equal(t, cube.Green, c.Corners["DBL"].Stickers[cube.FaceD])
		require.Equal(t, cube.Yellow, c.Corners["DBL"].Stickers[cube.FaceB])
		require.Equal(t, cube.Orange, c.Corners["DBL"].Stickers[cube.FaceL])

		require.Equal(t, cube.Blue, c.Corners["UBL"].Stickers[cube.FaceU])
		require.Equal(t, cube.Yellow, c.Corners["UBL"].Stickers[cube.FaceB])
		require.Equal(t, cube.Orange, c.Corners["UBL"].Stickers[cube.FaceL])

		require.Equal(t, cube.Blue, c.Corners["ULF"].Stickers[cube.FaceU])
		require.Equal(t, cube.White, c.Corners["ULF"].Stickers[cube.FaceF])
		require.Equal(t, cube.Orange, c.Corners["ULF"].Stickers[cube.FaceL])
	})

	t.Run("D rotation", func(t *testing.T) {
		c := cube.NewCube()
		c.D()

		// EDGES
		require.Equal(t, cube.Green, c.Edges["DR"].Stickers[cube.FaceR])
		require.Equal(t, cube.Yellow, c.Edges["DR"].Stickers[cube.FaceD])

		require.Equal(t, cube.Red, c.Edges["DB"].Stickers[cube.FaceB])
		require.Equal(t, cube.Yellow, c.Edges["DB"].Stickers[cube.FaceD])

		require.Equal(t, cube.Blue, c.Edges["DL"].Stickers[cube.FaceL])
		require.Equal(t, cube.Yellow, c.Edges["DL"].Stickers[cube.FaceD])

		require.Equal(t, cube.Orange, c.Edges["DF"].Stickers[cube.FaceF])
		require.Equal(t, cube.Yellow, c.Edges["DF"].Stickers[cube.FaceD])

		// CORNERS
		require.Equal(t, cube.Green, c.Corners["DRB"].Stickers[cube.FaceR])
		require.Equal(t, cube.Red, c.Corners["DRB"].Stickers[cube.FaceB])
		require.Equal(t, cube.Yellow, c.Corners["DRB"].Stickers[cube.FaceD])

		require.Equal(t, cube.Red, c.Corners["DBL"].Stickers[cube.FaceB])
		require.Equal(t, cube.Blue, c.Corners["DBL"].Stickers[cube.FaceL])
		require.Equal(t, cube.Yellow, c.Corners["DBL"].Stickers[cube.FaceD])

		require.Equal(t, cube.Blue, c.Corners["DLF"].Stickers[cube.FaceL])
		require.Equal(t, cube.Orange, c.Corners["DLF"].Stickers[cube.FaceF])
		require.Equal(t, cube.Yellow, c.Corners["DLF"].Stickers[cube.FaceD])

		require.Equal(t, cube.Orange, c.Corners["DFR"].Stickers[cube.FaceF])
		require.Equal(t, cube.Green, c.Corners["DFR"].Stickers[cube.FaceR])
		require.Equal(t, cube.Yellow, c.Corners["DFR"].Stickers[cube.FaceD])
	})

	t.Run("B rotation", func(t *testing.T) {
		c := cube.NewCube()
		c.B()

		// EDGES
		require.Equal(t, cube.White, c.Edges["BL"].Stickers[cube.FaceL])
		require.Equal(t, cube.Blue, c.Edges["BL"].Stickers[cube.FaceB])

		require.Equal(t, cube.Orange, c.Edges["DB"].Stickers[cube.FaceD])
		require.Equal(t, cube.Blue, c.Edges["DB"].Stickers[cube.FaceB])

		require.Equal(t, cube.Yellow, c.Edges["BR"].Stickers[cube.FaceR])
		require.Equal(t, cube.Blue, c.Edges["BR"].Stickers[cube.FaceB])

		require.Equal(t, cube.Red, c.Edges["UB"].Stickers[cube.FaceU])
		require.Equal(t, cube.Blue, c.Edges["UB"].Stickers[cube.FaceB])

		// CORNERS
		require.Equal(t, cube.Red, c.Corners["UBL"].Stickers[cube.FaceU])
		require.Equal(t, cube.White, c.Corners["UBL"].Stickers[cube.FaceL])
		require.Equal(t, cube.Blue, c.Corners["UBL"].Stickers[cube.FaceB])

		require.Equal(t, cube.Orange, c.Corners["DBL"].Stickers[cube.FaceD])
		require.Equal(t, cube.White, c.Corners["DBL"].Stickers[cube.FaceL])
		require.Equal(t, cube.Blue, c.Corners["DBL"].Stickers[cube.FaceB])

		require.Equal(t, cube.Orange, c.Corners["DRB"].Stickers[cube.FaceD])
		require.Equal(t, cube.Yellow, c.Corners["DRB"].Stickers[cube.FaceR])
		require.Equal(t, cube.Blue, c.Corners["DRB"].Stickers[cube.FaceB])

		require.Equal(t, cube.Red, c.Corners["URB"].Stickers[cube.FaceU])
		require.Equal(t, cube.Yellow, c.Corners["URB"].Stickers[cube.FaceR])
		require.Equal(t, cube.Blue, c.Corners["URB"].Stickers[cube.FaceB])
	})
}
