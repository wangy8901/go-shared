package utils

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMapPropertySource(t *testing.T) {
	s := NewMapPropertySource(map[string]interface{}{
		"a": "foo",
		"b": 12,
		"c": true,
		"d": map[string]interface{}{
			"foo": "bar",
		},
	})
	assert.Equal(t, "foo", s.GetString("a"))
	assert.Equal(t, 12, s.GetInt("b"))
	assert.Equal(t, true, s.GetBool("c"))
	assert.True(t, reflect.DeepEqual(s.Get("d"), map[string]interface{}{
		"foo": "bar",
	}))
}
