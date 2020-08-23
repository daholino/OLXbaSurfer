package helpers

import "testing"

type TestCase struct {
	arg      string
	expected string
}

func TestStripSlash(t *testing.T) {
	cases := []TestCase{
		{
			arg:      "/tmp/OLXbaSurfer/",
			expected: "/tmp/OLXbaSurfer",
		},
		{
			arg:      "/tmp/OLXbaSurfer",
			expected: "/tmp/OLXbaSurfer",
		},
		{
			arg:      "/tmp/OLXbaSurfer///",
			expected: "/tmp/OLXbaSurfer",
		},
		{
			arg:      "/tmp/OLXbaSurfer//",
			expected: "/tmp/OLXbaSurfer",
		},
	}

	for _, tc := range cases {
		got := StripSlash(tc.arg)
		if got != tc.expected {
			t.Errorf("StripSlash(%s: Expected %s, got %s", tc.arg, tc.expected, got)
		}
	}
}

func TestByteToUint64AndViceVersaConversion(t *testing.T) {
	var ID uint64 = 1
	data := I64tob(ID)
	IDconverted := Btoi64(data)
	if ID != IDconverted {
		t.Error("Byte conversion failure for ID 1")
	}

	ID = 39478903
	data = I64tob(ID)
	IDconverted = Btoi64(data)
	if ID != IDconverted {
		t.Error("Byte conversion failure for ID 39478903")
	}
}
