package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRotateT(t *testing.T) {
	require := require.New(t)
	shp := makeT(coords{x: 7, y: 3})
	expectedCenter := coords{x: 7, y: 3}
	require.Equal(expectedCenter, shp.squares()[0].position) // center
	expectedTop := coords{x: 7, y: 2}
	require.Equal(expectedTop, shp.squares()[1].position) // top
	expectedLeft := coords{x: 6, y: 3}
	require.Equal(expectedLeft, shp.squares()[2].position) // left
	expectedRight := coords{x: 8, y: 3}
	require.Equal(expectedRight, shp.squares()[3].position) // right

	shp.Rotate()
	expectedCenter = coords{x: 7, y: 3}
	require.Equal(expectedCenter, shp.squares()[0].position) // unchanged
	expectedTop = coords{x: 7, y: 2}
	require.Equal(expectedTop, shp.squares()[2].position) // left -> top
	expectedRight = coords{x: 8, y: 3}
	require.Equal(expectedRight, shp.squares()[1].position) // top -> right
	expectedBottom := coords{x: 7, y: 4}
	require.Equal(expectedBottom, shp.squares()[3].position) // right -> bottom
	// TODO: rotate 3 more times and test that we're back to normal
}

func TestRotateTBis(t *testing.T) {
	require := require.New(t)
	shp := makeT(coords{x: 5, y: 2})
	expectedCenter := coords{x: 5, y: 2}
	require.Equal(expectedCenter, shp.squares()[0].position) // center
	expectedTop := coords{x: 5, y: 1}
	require.Equal(expectedTop, shp.squares()[1].position) // top
	expectedLeft := coords{x: 4, y: 2}
	require.Equal(expectedLeft, shp.squares()[2].position) // left
	expectedRight := coords{x: 6, y: 2}
	require.Equal(expectedRight, shp.squares()[3].position) // right

	shp.Rotate()
	expectedCenter = coords{x: 5, y: 2}
	require.Equal(expectedCenter, shp.squares()[0].position) // unchanged
	expectedTop = coords{x: 5, y: 1}
	require.Equal(expectedTop, shp.squares()[2].position) // left -> top
	expectedRight = coords{x: 6, y: 2}
	require.Equal(expectedRight, shp.squares()[1].position) // top -> right
	expectedBottom := coords{x: 5, y: 3}
	require.Equal(expectedBottom, shp.squares()[3].position) // right -> bottom
	// TODO: rotate 3 more times and test that we're back to normal
}

func TestRotateJ(t *testing.T) {
	require := require.New(t)
	shp := makeJ(coords{x: 7, y: 4})
	expectedCenter := coords{x: 7, y: 3}
	require.Equal(expectedCenter, shp.squares()[0].position) // center
	expectedTop := coords{x: 7, y: 2}
	require.Equal(expectedTop, shp.squares()[1].position) // top
	expectedBranch := coords{x: 6, y: 4}
	require.Equal(expectedBranch, shp.squares()[2].position) // branch
	expectedBottom := coords{x: 7, y: 4}
	require.Equal(expectedBottom, shp.squares()[3].position) // bottom

	shp.Rotate()
	expectedCenter = coords{x: 7, y: 3}
	require.Equal(expectedCenter, shp.squares()[0].position)
	expectedTop = coords{x: 8, y: 3}
	require.Equal(expectedTop, shp.squares()[1].position)
	expectedBranch = coords{x: 6, y: 2}
	require.Equal(expectedBranch, shp.squares()[2].position)
	expectedBottom = coords{x: 6, y: 3}
	require.Equal(expectedBottom, shp.squares()[3].position)
}

func TestRotateI(t *testing.T) {
	require := require.New(t)
	shp := makeI(coords{x: 7, y: 10})
	expectedTop := coords{x: 7, y: 7}
	require.Equal(expectedTop, shp.squares()[0].position) // top
	expectedMiddle := coords{x: 7, y: 8}
	require.Equal(expectedMiddle, shp.squares()[1].position)
	expectedMiddle = coords{x: 7, y: 9}
	require.Equal(expectedMiddle, shp.squares()[2].position)
	expectedBottom := coords{x: 7, y: 10}
	require.Equal(expectedBottom, shp.squares()[3].position) // bottom

	shp.Rotate()
	expectedLeft := coords{x: 5, y: 8}
	require.Equal(expectedLeft, shp.squares()[0].position) // left
	expectedMiddle = coords{x: 6, y: 8}
	require.Equal(expectedMiddle, shp.squares()[1].position)
	expectedMiddle = coords{x: 7, y: 8}
	require.Equal(expectedMiddle, shp.squares()[2].position)
	expectedRight := coords{x: 8, y: 8}
	require.Equal(expectedRight, shp.squares()[3].position) // right

	shp.Rotate()
	expectedTop = coords{x: 7, y: 7}
	require.Equal(expectedTop, shp.squares()[0].position) // top
	expectedMiddle = coords{x: 7, y: 8}
	require.Equal(expectedMiddle, shp.squares()[1].position)
	expectedMiddle = coords{x: 7, y: 9}
	require.Equal(expectedMiddle, shp.squares()[2].position)
	expectedBottom = coords{x: 7, y: 10}
	require.Equal(expectedBottom, shp.squares()[3].position) // bottom
}
