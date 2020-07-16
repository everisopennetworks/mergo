package mergo

import (
	"testing"
)

func TestIssue123(t *testing.T) {
	src := map[string]interface{}{
		"col1": nil,
		"col2": 4,
		"col3": nil,
	}
	dst := map[string]interface{}{
		"col1": 2,
		"col2": 3,
		"col3": 3,
	}

	// Expected behavior
	if err := Merge(&dst, src, WithOverride); err != nil {
		t.Error(err)
	}
	testCases := []struct {
		key      string
		expected interface{}
	}{
		{
			"col1",
			nil,
		},
		{
			"col2",
			4,
		},
		{
			"col3",
			nil,
		},
	}
	for _, tC := range testCases {
		if dst[tC.key] != tC.expected {
			t.Errorf("expected %v in dst[%q], got %v", tC.expected, tC.key, dst[tC.key])
		}
	}
}
