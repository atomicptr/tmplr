package tmpl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchAndParseFilename(t *testing.T) {
	testCases := []struct {
		filename     string
		template     string
		shouldMatch  bool
		parsedValues map[string]string
	}{
		{"test.clj", "test.clj", true, nil},
		{"yolo.cpp", "[name].cpp", true, map[string]string{"name": "yolo"}},
		{"IndexController.php", "[controllerName]Controller.php", true, map[string]string{"controllerName": "Index"}},
		{"test123.js", "test[var].js", true, map[string]string{"var": "123"}},
		{"testXYZtest.clj", "test[var]test.clj", true, map[string]string{"var": "XYZ"}},
		{"fooBarbaz.ml", "[var1]Bar[var2].ml", true, map[string]string{"var1": "foo", "var2": "baz"}},
		{"example.txt", "[name].txt", true, map[string]string{"name": "example"}},
		{"example.txt", "[name].c", false, nil},
	}

	for _, tc := range testCases {
		if tc.shouldMatch {
			assert.True(t, MatchesFilename(tc.filename, tc.template), "%s should match %s", tc.filename, tc.template)
		} else {
			assert.False(t, MatchesFilename(tc.filename, tc.template), "%s should not match %s", tc.filename, tc.template)
		}

		parsedValues := ParseFilename(tc.filename, tc.template)

		assert.Equal(t, tc.parsedValues, parsedValues)
	}
}
