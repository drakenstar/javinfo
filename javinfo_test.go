package javinfo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindByCode(t *testing.T) {
	info := New()
	titles, err := info.FindByCode(Code("IPX-666"))
	assert.NoError(t, err)
	assert.Len(t, titles, 1)
	assert.Contains(t, titles[0].Title, "Last Train")

	t.Fail()
}
