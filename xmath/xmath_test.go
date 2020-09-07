package xmath

import (
	"testing"

	"github.com/aca/go-misc/xstrconv"
	"github.com/stretchr/testify/assert"
)

func TestFloor(t *testing.T) {
	{
		expected := xstrconv.FormatFloat(100.4, 1)
		actual := xstrconv.FormatFloat(FloorN(100.44, 1), 1)
		assert.Equal(t, expected, actual)
	}
	{
		expected := xstrconv.FormatFloat(100.5, 1)
		actual := xstrconv.FormatFloat(RoundN(100.46, 1), 1)
		assert.Equal(t, expected, actual)
	}
	{
		expected := xstrconv.FormatFloat(100.5, 1)
		actual := xstrconv.FormatFloat(CeilN(100.46, 1), 1)
		assert.Equal(t, expected, actual)
	}
}
