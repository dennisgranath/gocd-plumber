package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleMerge(t *testing.T) {
	var simpleTests = []struct {
		base     map[string]interface{}
		over     map[string]interface{}
		expected map[string]interface{}
	}{
		{
			map[string]interface{}{"base": "test"},
			map[string]interface{}{"over": "test"},
			map[string]interface{}{"base": "test", "over": "test"},
		},
		{
			map[string]interface{}{"base": "test"},
			map[string]interface{}{"base": "override"},
			map[string]interface{}{"base": "override"},
		},
		{
			map[string]interface{}{"base": "test"},
			map[string]interface{}{"over": map[string]interface{}{"net": "eth1"}},
			map[string]interface{}{"base": "test", "over": map[string]interface{}{"net": "eth1"}},
		},
		{
			map[string]interface{}{"base": map[string]interface{}{"base": "1234"}},
			map[string]interface{}{"base": map[string]interface{}{"over": "4321"}},
			map[string]interface{}{"base": map[string]interface{}{"base": "1234", "over": "4321"}},
		},
		{
			map[string]interface{}{"base": map[string]interface{}{"base": "1234"}},
			map[string]interface{}{"base": map[string]interface{}{"over": "4321", "another": 1}, "over": 2},
			map[string]interface{}{"base": map[string]interface{}{"base": "1234", "over": "4321", "another": 1}, "over": 2},
		},
	}

	for _, tt := range simpleTests {
		actual := merge(tt.base, tt.over)
		assert.Equal(t, tt.expected, actual)
	}
}

func TestSimpleMergeNegative(t *testing.T) {
	var simpleTests = []struct {
		base     map[string]interface{}
		over     map[string]interface{}
		expected map[string]interface{}
	}{
		{
			map[string]interface{}{"base": "test"},
			map[string]interface{}{"over": "test"},
			map[string]interface{}{"base": "test"},
		},
		{
			map[string]interface{}{"base": "test"},
			map[string]interface{}{"base": "override"},
			map[string]interface{}{"base": "test"},
		},
		{
			map[string]interface{}{"base": "test"},
			map[string]interface{}{"over": map[string]interface{}{"net": "eth1"}},
			map[string]interface{}{"over": map[string]interface{}{"net": "eth1"}},
		},
		{
			map[string]interface{}{"base": map[string]interface{}{"base": "1234"}},
			map[string]interface{}{"base": map[string]interface{}{"over": "4321"}},
			map[string]interface{}{"base": map[string]interface{}{"over": "4321"}},
		},
		{
			map[string]interface{}{"base": map[string]interface{}{"base": "1234"}},
			map[string]interface{}{"base": map[string]interface{}{"over": "4321", "another": 1}, "over": 2},
			map[string]interface{}{"base": map[string]interface{}{"base": "1234", "over": "4321", "another": 2}, "over": 3},
		},
	}

	for _, tt := range simpleTests {
		actual := merge(tt.base, tt.over)
		assert.NotEqual(t, tt.expected, actual)
	}
}
