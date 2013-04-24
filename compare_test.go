package version

import (
	"strings"
	"testing"
)

var prepVersionValues = map[string]string{
	"1-stable":       "1.stable.",
	"1....0.0":       "1.0.0",
	"1-2_3-4":        "1.2.3.4",
	"1-2_3-..4-beta": "1.2.3.4.beta.",
}

func TestPrepVersion(t *testing.T) {
	for in, out := range prepVersionValues {
		if x := prepVersion(in); strings.Join(x, ".") != out {
			t.Errorf("FAIL: Normalize(%v) = %v: want %v", in, x, out)
		}
	}
}

var numVersionValues = map[string]int{
	"foo":   -7,
	"10":    10,
	"12212": 12212,
	"":      0,
	"dev":   -6,
	"alpha": -5,
	"a":     -5,
	"beta":  -4,
	"b":     -4,
	"RC":    -3,
	"rc":    -3,
	"#":     -2,
	"p":     1,
	"pl":    1,
}

func TestNumVersion(t *testing.T) {
	for in, out := range numVersionValues {
		if x := numVersion(in); x != out {
			t.Errorf("FAIL: Normalize(%v) = %v: want %v", in, x, out)
		}
	}
}

var compareVersionSimpleValues = map[string]int{
	"1|2":             -1,
	"10|2":            1,
	"1.0|1.1":         -1,
	"1.2|1.0.1":       1,
	"1.0-dev|1.0-dev": 0,
	"1.0-dev|1.0a1":   -1,
	"1.0-dev|1.0b1":   -1,
	"1.0-dev|1.0RC1":  -1,
	"1.0-dev|1.0rc1":  -1,
	"1.0-dev|1.0":     -1,
	"1.0-dev|1.0pl1":  -1,
	"1.0a1|1.0-dev":   1,
	"1.0a1|1.0a1":     0,
	"1.0a1|1.0b1":     -1,
	"1.0a1|1.0RC1":    -1,
	"1.0a1|1.0rc1":    -1,
	"1.0a1|1.0":       -1,
	"1.0a1|1.0pl1":    -1,
	"1.0b1|1.0-dev":   1,
	"1.0b1|1.0a1":     1,
	"1.0b1|1.0b1":     0,
	"1.0b1|1.0RC1":    -1,
	"1.0b1|1.0rc1":    -1,
	"1.0b1|1.0":       -1,
	"1.0b1|1.0pl1":    -1,
	"1.0RC1|1.0-dev":  1,
	"1.0RC1|1.0a1":    1,
	"1.0RC1|1.0b1":    1,
	"1.0RC1|1.0RC1":   0,
	"1.0RC1|1.0rc1":   0,
	"1.0RC1|1.0":      -1,
	"1.0RC1|1.0pl1":   -1,
	"1.0rc1|1.0-dev":  1,
	"1.0rc1|1.0a1":    1,
	"1.0rc1|1.0b1":    1,
	"1.0rc1|1.0RC1":   0,
	"1.0rc1|1.0rc1":   0,
	"1.0rc1|1.0":      -1,
	"1.0rc1|1.0pl1":   -1,
	"1.0|1.0-dev":     1,
	"1.0|1.0a1":       1,
	"1.0|1.0b1":       1,
	"1.0|1.0RC1":      1,
	"1.0|1.0rc1":      1,
	"1.0|1.0":         0,
	"1.0|1.0pl1":      -1,
	"1.0pl1|1.0-dev":  1,
	"1.0pl1|1.0a1":    1,
	"1.0pl1|1.0b1":    1,
	"1.0pl1|1.0RC1":   1,
	"1.0pl1|1.0rc1":   1,
	"1.0pl1|1.0":      1,
	"1.0pl1|1.0pl1":   0,
}

func TestCompareVersionSimple(t *testing.T) {
	for in, out := range compareVersionSimpleValues {
		v := strings.Split(in, "|")
		if x := CompareSimple(v[0], v[1]); x != out {
			t.Errorf("FAIL: CompareVersionSimple(%v) = %v: want %v", in, x, out)
		}
	}
}

var compareVersionValues = map[string]bool{
	"1.0-dev lt 1.0-dev": false,
	"1.0-dev < 1.0-dev":  false,
	"1.0-dev le 1.0-dev": true,
	"1.0-dev <= 1.0-dev": true,
	"1.0-dev gt 1.0-dev": false,
	"1.0-dev > 1.0-dev":  false,
	"1.0-dev ge 1.0-dev": true,
	"1.0-dev >= 1.0-dev": true,
	"1.0-dev eq 1.0-dev": true,
	"1.0-dev = 1.0-dev":  true,
	"1.0-dev == 1.0-dev": true,
	"1.0-dev ne 1.0-dev": false,
	"1.0-dev <> 1.0-dev": false,
	"1.0-dev != 1.0-dev": false,
	"1.0-dev lt 1.0a1":   true,
	"1.0-dev < 1.0a1":    true,
	"1.0-dev le 1.0a1":   true,
	"1.0-dev <= 1.0a1":   true,
	"1.0-dev gt 1.0a1":   false,
	"1.0-dev > 1.0a1":    false,
	"1.0-dev ge 1.0a1":   false,
	"1.0-dev >= 1.0a1":   false,
	"1.0-dev eq 1.0a1":   false,
	"1.0-dev = 1.0a1":    false,
	"1.0-dev == 1.0a1":   false,
	"1.0-dev ne 1.0a1":   true,
	"1.0-dev <> 1.0a1":   true,
	"1.0-dev != 1.0a1":   true,
	"1.0-dev lt 1.0b1":   true,
	"1.0-dev < 1.0b1":    true,
	"1.0-dev le 1.0b1":   true,
	"1.0-dev <= 1.0b1":   true,
	"1.0-dev gt 1.0b1":   false,
	"1.0-dev > 1.0b1":    false,
	"1.0-dev ge 1.0b1":   false,
	"1.0-dev >= 1.0b1":   false,
	"1.0-dev eq 1.0b1":   false,
	"1.0-dev = 1.0b1":    false,
	"1.0-dev == 1.0b1":   false,
	"1.0-dev ne 1.0b1":   true,
	"1.0-dev <> 1.0b1":   true,
	"1.0-dev != 1.0b1":   true,
	"1.0-dev lt 1.0RC1":  true,
	"1.0-dev < 1.0RC1":   true,
	"1.0-dev le 1.0RC1":  true,
	"1.0-dev <= 1.0RC1":  true,
	"1.0-dev gt 1.0RC1":  false,
	"1.0-dev > 1.0RC1":   false,
	"1.0-dev ge 1.0RC1":  false,
	"1.0-dev >= 1.0RC1":  false,
	"1.0-dev eq 1.0RC1":  false,
	"1.0-dev = 1.0RC1":   false,
	"1.0-dev == 1.0RC1":  false,
	"1.0-dev ne 1.0RC1":  true,
	"1.0-dev <> 1.0RC1":  true,
	"1.0-dev != 1.0RC1":  true,
	"1.0-dev lt 1.0rc1":  true,
	"1.0-dev < 1.0rc1":   true,
	"1.0-dev le 1.0rc1":  true,
	"1.0-dev <= 1.0rc1":  true,
	"1.0-dev gt 1.0rc1":  false,
	"1.0-dev > 1.0rc1":   false,
	"1.0-dev ge 1.0rc1":  false,
	"1.0-dev >= 1.0rc1":  false,
	"1.0-dev eq 1.0rc1":  false,
	"1.0-dev = 1.0rc1":   false,
	"1.0-dev == 1.0rc1":  false,
	"1.0-dev ne 1.0rc1":  true,
	"1.0-dev <> 1.0rc1":  true,
	"1.0-dev != 1.0rc1":  true,
	"1.0-dev lt 1.0":     true,
	"1.0-dev < 1.0":      true,
	"1.0-dev le 1.0":     true,
	"1.0-dev <= 1.0":     true,
	"1.0-dev gt 1.0":     false,
	"1.0-dev > 1.0":      false,
	"1.0-dev ge 1.0":     false,
	"1.0-dev >= 1.0":     false,
	"1.0-dev eq 1.0":     false,
	"1.0-dev = 1.0":      false,
	"1.0-dev == 1.0":     false,
	"1.0-dev ne 1.0":     true,
	"1.0-dev <> 1.0":     true,
	"1.0-dev != 1.0":     true,
	"1.0-dev lt 1.0pl1":  true,
	"1.0-dev < 1.0pl1":   true,
	"1.0-dev le 1.0pl1":  true,
	"1.0-dev <= 1.0pl1":  true,
	"1.0-dev gt 1.0pl1":  false,
	"1.0-dev > 1.0pl1":   false,
	"1.0-dev ge 1.0pl1":  false,
	"1.0-dev >= 1.0pl1":  false,
	"1.0-dev eq 1.0pl1":  false,
	"1.0-dev = 1.0pl1":   false,
	"1.0-dev == 1.0pl1":  false,
	"1.0-dev ne 1.0pl1":  true,
	"1.0-dev <> 1.0pl1":  true,
	"1.0-dev != 1.0pl1":  true,
	"1.0a1 lt 1.0-dev":   false,
	"1.0a1 < 1.0-dev":    false,
	"1.0a1 le 1.0-dev":   false,
	"1.0a1 <= 1.0-dev":   false,
	"1.0a1 gt 1.0-dev":   true,
	"1.0a1 > 1.0-dev":    true,
	"1.0a1 ge 1.0-dev":   true,
	"1.0a1 >= 1.0-dev":   true,
	"1.0a1 eq 1.0-dev":   false,
	"1.0a1 = 1.0-dev":    false,
	"1.0a1 == 1.0-dev":   false,
	"1.0a1 ne 1.0-dev":   true,
	"1.0a1 <> 1.0-dev":   true,
	"1.0a1 != 1.0-dev":   true,
	"1.0a1 lt 1.0a1":     false,
	"1.0a1 < 1.0a1":      false,
	"1.0a1 le 1.0a1":     true,
	"1.0a1 <= 1.0a1":     true,
	"1.0a1 gt 1.0a1":     false,
	"1.0a1 > 1.0a1":      false,
	"1.0a1 ge 1.0a1":     true,
	"1.0a1 >= 1.0a1":     true,
	"1.0a1 eq 1.0a1":     true,
	"1.0a1 = 1.0a1":      true,
	"1.0a1 == 1.0a1":     true,
	"1.0a1 ne 1.0a1":     false,
	"1.0a1 <> 1.0a1":     false,
	"1.0a1 != 1.0a1":     false,
	"1.0a1 lt 1.0b1":     true,
	"1.0a1 < 1.0b1":      true,
	"1.0a1 le 1.0b1":     true,
	"1.0a1 <= 1.0b1":     true,
	"1.0a1 gt 1.0b1":     false,
	"1.0a1 > 1.0b1":      false,
	"1.0a1 ge 1.0b1":     false,
	"1.0a1 >= 1.0b1":     false,
	"1.0a1 eq 1.0b1":     false,
	"1.0a1 = 1.0b1":      false,
	"1.0a1 == 1.0b1":     false,
	"1.0a1 ne 1.0b1":     true,
	"1.0a1 <> 1.0b1":     true,
	"1.0a1 != 1.0b1":     true,
	"1.0a1 lt 1.0RC1":    true,
	"1.0a1 < 1.0RC1":     true,
	"1.0a1 le 1.0RC1":    true,
	"1.0a1 <= 1.0RC1":    true,
	"1.0a1 gt 1.0RC1":    false,
	"1.0a1 > 1.0RC1":     false,
	"1.0a1 ge 1.0RC1":    false,
	"1.0a1 >= 1.0RC1":    false,
	"1.0a1 eq 1.0RC1":    false,
	"1.0a1 = 1.0RC1":     false,
	"1.0a1 == 1.0RC1":    false,
	"1.0a1 ne 1.0RC1":    true,
	"1.0a1 <> 1.0RC1":    true,
	"1.0a1 != 1.0RC1":    true,
	"1.0a1 lt 1.0rc1":    true,
	"1.0a1 < 1.0rc1":     true,
	"1.0a1 le 1.0rc1":    true,
	"1.0a1 <= 1.0rc1":    true,
	"1.0a1 gt 1.0rc1":    false,
	"1.0a1 > 1.0rc1":     false,
	"1.0a1 ge 1.0rc1":    false,
	"1.0a1 >= 1.0rc1":    false,
	"1.0a1 eq 1.0rc1":    false,
	"1.0a1 = 1.0rc1":     false,
	"1.0a1 == 1.0rc1":    false,
	"1.0a1 ne 1.0rc1":    true,
	"1.0a1 <> 1.0rc1":    true,
	"1.0a1 != 1.0rc1":    true,
	"1.0a1 lt 1.0":       true,
	"1.0a1 < 1.0":        true,
	"1.0a1 le 1.0":       true,
	"1.0a1 <= 1.0":       true,
	"1.0a1 gt 1.0":       false,
	"1.0a1 > 1.0":        false,
	"1.0a1 ge 1.0":       false,
	"1.0a1 >= 1.0":       false,
	"1.0a1 eq 1.0":       false,
	"1.0a1 = 1.0":        false,
	"1.0a1 == 1.0":       false,
	"1.0a1 ne 1.0":       true,
	"1.0a1 <> 1.0":       true,
	"1.0a1 != 1.0":       true,
	"1.0a1 lt 1.0pl1":    true,
	"1.0a1 < 1.0pl1":     true,
	"1.0a1 le 1.0pl1":    true,
	"1.0a1 <= 1.0pl1":    true,
	"1.0a1 gt 1.0pl1":    false,
	"1.0a1 > 1.0pl1":     false,
	"1.0a1 ge 1.0pl1":    false,
	"1.0a1 >= 1.0pl1":    false,
	"1.0a1 eq 1.0pl1":    false,
	"1.0a1 = 1.0pl1":     false,
	"1.0a1 == 1.0pl1":    false,
	"1.0a1 ne 1.0pl1":    true,
	"1.0a1 <> 1.0pl1":    true,
	"1.0a1 != 1.0pl1":    true,
	"1.0b1 lt 1.0-dev":   false,
	"1.0b1 < 1.0-dev":    false,
	"1.0b1 le 1.0-dev":   false,
	"1.0b1 <= 1.0-dev":   false,
	"1.0b1 gt 1.0-dev":   true,
	"1.0b1 > 1.0-dev":    true,
	"1.0b1 ge 1.0-dev":   true,
	"1.0b1 >= 1.0-dev":   true,
	"1.0b1 eq 1.0-dev":   false,
	"1.0b1 = 1.0-dev":    false,
	"1.0b1 == 1.0-dev":   false,
	"1.0b1 ne 1.0-dev":   true,
	"1.0b1 <> 1.0-dev":   true,
	"1.0b1 != 1.0-dev":   true,
	"1.0b1 lt 1.0a1":     false,
	"1.0b1 < 1.0a1":      false,
	"1.0b1 le 1.0a1":     false,
	"1.0b1 <= 1.0a1":     false,
	"1.0b1 gt 1.0a1":     true,
	"1.0b1 > 1.0a1":      true,
	"1.0b1 ge 1.0a1":     true,
	"1.0b1 >= 1.0a1":     true,
	"1.0b1 eq 1.0a1":     false,
	"1.0b1 = 1.0a1":      false,
	"1.0b1 == 1.0a1":     false,
	"1.0b1 ne 1.0a1":     true,
	"1.0b1 <> 1.0a1":     true,
	"1.0b1 != 1.0a1":     true,
	"1.0b1 lt 1.0b1":     false,
	"1.0b1 < 1.0b1":      false,
	"1.0b1 le 1.0b1":     true,
	"1.0b1 <= 1.0b1":     true,
	"1.0b1 gt 1.0b1":     false,
	"1.0b1 > 1.0b1":      false,
	"1.0b1 ge 1.0b1":     true,
	"1.0b1 >= 1.0b1":     true,
	"1.0b1 eq 1.0b1":     true,
	"1.0b1 = 1.0b1":      true,
	"1.0b1 == 1.0b1":     true,
	"1.0b1 ne 1.0b1":     false,
	"1.0b1 <> 1.0b1":     false,
	"1.0b1 != 1.0b1":     false,
	"1.0b1 lt 1.0RC1":    true,
	"1.0b1 < 1.0RC1":     true,
	"1.0b1 le 1.0RC1":    true,
	"1.0b1 <= 1.0RC1":    true,
	"1.0b1 gt 1.0RC1":    false,
	"1.0b1 > 1.0RC1":     false,
	"1.0b1 ge 1.0RC1":    false,
	"1.0b1 >= 1.0RC1":    false,
	"1.0b1 eq 1.0RC1":    false,
	"1.0b1 = 1.0RC1":     false,
	"1.0b1 == 1.0RC1":    false,
	"1.0b1 ne 1.0RC1":    true,
	"1.0b1 <> 1.0RC1":    true,
	"1.0b1 != 1.0RC1":    true,
	"1.0b1 lt 1.0rc1":    true,
	"1.0b1 < 1.0rc1":     true,
	"1.0b1 le 1.0rc1":    true,
	"1.0b1 <= 1.0rc1":    true,
	"1.0b1 gt 1.0rc1":    false,
	"1.0b1 > 1.0rc1":     false,
	"1.0b1 ge 1.0rc1":    false,
	"1.0b1 >= 1.0rc1":    false,
	"1.0b1 eq 1.0rc1":    false,
	"1.0b1 = 1.0rc1":     false,
	"1.0b1 == 1.0rc1":    false,
	"1.0b1 ne 1.0rc1":    true,
	"1.0b1 <> 1.0rc1":    true,
	"1.0b1 != 1.0rc1":    true,
	"1.0b1 lt 1.0":       true,
	"1.0b1 < 1.0":        true,
	"1.0b1 le 1.0":       true,
	"1.0b1 <= 1.0":       true,
	"1.0b1 gt 1.0":       false,
	"1.0b1 > 1.0":        false,
	"1.0b1 ge 1.0":       false,
	"1.0b1 >= 1.0":       false,
	"1.0b1 eq 1.0":       false,
	"1.0b1 = 1.0":        false,
	"1.0b1 == 1.0":       false,
	"1.0b1 ne 1.0":       true,
	"1.0b1 <> 1.0":       true,
	"1.0b1 != 1.0":       true,
	"1.0b1 lt 1.0pl1":    true,
	"1.0b1 < 1.0pl1":     true,
	"1.0b1 le 1.0pl1":    true,
	"1.0b1 <= 1.0pl1":    true,
	"1.0b1 gt 1.0pl1":    false,
	"1.0b1 > 1.0pl1":     false,
	"1.0b1 ge 1.0pl1":    false,
	"1.0b1 >= 1.0pl1":    false,
	"1.0b1 eq 1.0pl1":    false,
	"1.0b1 = 1.0pl1":     false,
	"1.0b1 == 1.0pl1":    false,
	"1.0b1 ne 1.0pl1":    true,
	"1.0b1 <> 1.0pl1":    true,
	"1.0b1 != 1.0pl1":    true,
	"1.0RC1 lt 1.0-dev":  false,
	"1.0RC1 < 1.0-dev":   false,
	"1.0RC1 le 1.0-dev":  false,
	"1.0RC1 <= 1.0-dev":  false,
	"1.0RC1 gt 1.0-dev":  true,
	"1.0RC1 > 1.0-dev":   true,
	"1.0RC1 ge 1.0-dev":  true,
	"1.0RC1 >= 1.0-dev":  true,
	"1.0RC1 eq 1.0-dev":  false,
	"1.0RC1 = 1.0-dev":   false,
	"1.0RC1 == 1.0-dev":  false,
	"1.0RC1 ne 1.0-dev":  true,
	"1.0RC1 <> 1.0-dev":  true,
	"1.0RC1 != 1.0-dev":  true,
	"1.0RC1 lt 1.0a1":    false,
	"1.0RC1 < 1.0a1":     false,
	"1.0RC1 le 1.0a1":    false,
	"1.0RC1 <= 1.0a1":    false,
	"1.0RC1 gt 1.0a1":    true,
	"1.0RC1 > 1.0a1":     true,
	"1.0RC1 ge 1.0a1":    true,
	"1.0RC1 >= 1.0a1":    true,
	"1.0RC1 eq 1.0a1":    false,
	"1.0RC1 = 1.0a1":     false,
	"1.0RC1 == 1.0a1":    false,
	"1.0RC1 ne 1.0a1":    true,
	"1.0RC1 <> 1.0a1":    true,
	"1.0RC1 != 1.0a1":    true,
	"1.0RC1 lt 1.0b1":    false,
	"1.0RC1  < 1.0b1":    false,
	"1.0RC1 le 1.0b1":    false,
	"1.0RC1 <= 1.0b1":    false,
	"1.0RC1 gt 1.0b1":    true,
	"1.0RC1 > 1.0b1":     true,
	"1.0RC1 ge 1.0b1":    true,
	"1.0RC1 >= 1.0b1":    true,
	"1.0RC1 eq 1.0b1":    false,
	"1.0RC1 = 1.0b1":     false,
	"1.0RC1 == 1.0b1":    false,
	"1.0RC1 ne 1.0b1":    true,
	"1.0RC1 <> 1.0b1":    true,
	"1.0RC1 != 1.0b1":    true,
	"1.0RC1 lt 1.0RC1":   false,
	"1.0RC1 < 1.0RC1":    false,
	"1.0RC1 le 1.0RC1":   true,
	"1.0RC1 <= 1.0RC1":   true,
	"1.0RC1 gt 1.0RC1":   false,
	"1.0RC1 > 1.0RC1":    false,
	"1.0RC1 ge 1.0RC1":   true,
	"1.0RC1 >= 1.0RC1":   true,
	"1.0RC1 eq 1.0RC1":   true,
	"1.0RC1 = 1.0RC1":    true,
	"1.0RC1 == 1.0RC1":   true,
	"1.0RC1 ne 1.0RC1":   false,
	"1.0RC1 <> 1.0RC1":   false,
	"1.0RC1 != 1.0RC1":   false,
	"1.0RC1 lt 1.0rc1":   false,
	"1.0RC1 < 1.0rc1":    false,
	"1.0RC1 le 1.0rc1":   true,
	"1.0RC1 <= 1.0rc1":   true,
	"1.0RC1 gt 1.0rc1":   false,
	"1.0RC1 > 1.0rc1":    false,
	"1.0RC1 ge 1.0rc1":   true,
	"1.0RC1 >= 1.0rc1":   true,
	"1.0RC1 eq 1.0rc1":   true,
	"1.0RC1 = 1.0rc1":    true,
	"1.0RC1 == 1.0rc1":   true,
	"1.0RC1 ne 1.0rc1":   false,
	"1.0RC1 <> 1.0rc1":   false,
	"1.0RC1 != 1.0rc1":   false,
	"1.0RC1 lt 1.0":      true,
	"1.0RC1 < 1.0":       true,
	"1.0RC1 le 1.0":      true,
	"1.0RC1 <= 1.0":      true,
	"1.0RC1 gt 1.0":      false,
	"1.0RC1 > 1.0":       false,
	"1.0RC1 ge 1.0":      false,
	"1.0RC1 >= 1.0":      false,
	"1.0RC1 eq 1.0":      false,
	"1.0RC1 = 1.0":       false,
	"1.0RC1 == 1.0":      false,
	"1.0RC1 ne 1.0":      true,
	"1.0RC1 <> 1.0":      true,
	"1.0RC1 != 1.0":      true,
	"1.0RC1 lt 1.0pl1":   true,
	"1.0RC1 < 1.0pl1":    true,
	"1.0RC1 le 1.0pl1":   true,
	"1.0RC1 <= 1.0pl1":   true,
	"1.0RC1 gt 1.0pl1":   false,
	"1.0RC1 > 1.0pl1":    false,
	"1.0RC1 ge 1.0pl1":   false,
	"1.0RC1 >= 1.0pl1":   false,
	"1.0RC1 eq 1.0pl1":   false,
	"1.0RC1 = 1.0pl1":    false,
	"1.0RC1 == 1.0pl1":   false,
	"1.0RC1 ne 1.0pl1":   true,
	"1.0RC1 <> 1.0pl1":   true,
	"1.0RC1 != 1.0pl1":   true,
	"1.0rc1 lt 1.0-dev":  false,
	"1.0rc1 < 1.0-dev":   false,
	"1.0rc1 le 1.0-dev":  false,
	"1.0rc1 <= 1.0-dev":  false,
	"1.0rc1 gt 1.0-dev":  true,
	"1.0rc1 > 1.0-dev":   true,
	"1.0rc1 ge 1.0-dev":  true,
	"1.0rc1 >= 1.0-dev":  true,
	"1.0rc1 eq 1.0-dev":  false,
	"1.0rc1 = 1.0-dev":   false,
	"1.0rc1 == 1.0-dev":  false,
	"1.0rc1 ne 1.0-dev":  true,
	"1.0rc1 <> 1.0-dev":  true,
	"1.0rc1 != 1.0-dev":  true,
	"1.0rc1 lt 1.0a1":    false,
	"1.0rc1 < 1.0a1":     false,
	"1.0rc1 le 1.0a1":    false,
	"1.0rc1 <= 1.0a1":    false,
	"1.0rc1 gt 1.0a1":    true,
	"1.0rc1 > 1.0a1":     true,
	"1.0rc1 ge 1.0a1":    true,
	"1.0rc1 >= 1.0a1":    true,
	"1.0rc1 eq 1.0a1":    false,
	"1.0rc1 = 1.0a1":     false,
	"1.0rc1 == 1.0a1":    false,
	"1.0rc1 ne 1.0a1":    true,
	"1.0rc1 <> 1.0a1":    true,
	"1.0rc1 != 1.0a1":    true,
	"1.0rc1 lt 1.0b1":    false,
	"1.0rc1 < 1.0b1":     false,
	"1.0rc1 le 1.0b1":    false,
	"1.0rc1 <= 1.0b1":    false,
	"1.0rc1 gt 1.0b1":    true,
	"1.0rc1 > 1.0b1":     true,
	"1.0rc1 ge 1.0b1":    true,
	"1.0rc1 >= 1.0b1":    true,
	"1.0rc1 eq 1.0b1":    false,
	"1.0rc1 = 1.0b1":     false,
	"1.0rc1 == 1.0b1":    false,
	"1.0rc1 ne 1.0b1":    true,
	"1.0rc1 <> 1.0b1":    true,
	"1.0rc1 != 1.0b1":    true,
	"1.0rc1 lt 1.0RC1":   false,
	"1.0rc1 < 1.0RC1":    false,
	"1.0rc1 le 1.0RC1":   true,
	"1.0rc1 <= 1.0RC1":   true,
	"1.0rc1 gt 1.0RC1":   false,
	"1.0rc1 > 1.0RC1":    false,
	"1.0rc1 ge 1.0RC1":   true,
	"1.0rc1 >= 1.0RC1":   true,
	"1.0rc1 eq 1.0RC1":   true,
	"1.0rc1 = 1.0RC1":    true,
	"1.0rc1 == 1.0RC1":   true,
	"1.0rc1 ne 1.0RC1":   false,
	"1.0rc1 <> 1.0RC1":   false,
	"1.0rc1 != 1.0RC1":   false,
	"1.0rc1 lt 1.0rc1":   false,
	"1.0rc1 < 1.0rc1":    false,
	"1.0rc1 le 1.0rc1":   true,
	"1.0rc1 <= 1.0rc1":   true,
	"1.0rc1 gt 1.0rc1":   false,
	"1.0rc1 > 1.0rc1":    false,
	"1.0rc1 ge 1.0rc1":   true,
	"1.0rc1 >= 1.0rc1":   true,
	"1.0rc1 eq 1.0rc1":   true,
	"1.0rc1 = 1.0rc1":    true,
	"1.0rc1 == 1.0rc1":   true,
	"1.0rc1 ne 1.0rc1":   false,
	"1.0rc1 <> 1.0rc1":   false,
	"1.0rc1 != 1.0rc1":   false,
	"1.0rc1 lt 1.0":      true,
	"1.0rc1 < 1.0":       true,
	"1.0rc1 le 1.0":      true,
	"1.0rc1 <= 1.0":      true,
	"1.0rc1 gt 1.0":      false,
	"1.0rc1 > 1.0":       false,
	"1.0rc1 ge 1.0":      false,
	"1.0rc1 >= 1.0":      false,
	"1.0rc1 eq 1.0":      false,
	"1.0rc1 = 1.0":       false,
	"1.0rc1 == 1.0":      false,
	"1.0rc1 ne 1.0":      true,
	"1.0rc1 <> 1.0":      true,
	"1.0rc1 != 1.0":      true,
	"1.0rc1 lt 1.0pl1":   true,
	"1.0rc1 < 1.0pl1":    true,
	"1.0rc1 le 1.0pl1":   true,
	"1.0rc1 <= 1.0pl1":   true,
	"1.0rc1 gt 1.0pl1":   false,
	"1.0rc1 > 1.0pl1":    false,
	"1.0rc1 ge 1.0pl1":   false,
	"1.0rc1 >= 1.0pl1":   false,
	"1.0rc1 eq 1.0pl1":   false,
	"1.0rc1 = 1.0pl1":    false,
	"1.0rc1 == 1.0pl1":   false,
	"1.0rc1 ne 1.0pl1":   true,
	"1.0rc1 <> 1.0pl1":   true,
	"1.0rc1 != 1.0pl1":   true,
	"1.0 lt 1.0-dev":     false,
	"1.0 < 1.0-dev":      false,
	"1.0 le 1.0-dev":     false,
	"1.0 <= 1.0-dev":     false,
	"1.0 gt 1.0-dev":     true,
	"1.0 > 1.0-dev":      true,
	"1.0 ge 1.0-dev":     true,
	"1.0 >= 1.0-dev":     true,
	"1.0 eq 1.0-dev":     false,
	"1.0 = 1.0-dev":      false,
	"1.0 == 1.0-dev":     false,
	"1.0 ne 1.0-dev":     true,
	"1.0 <> 1.0-dev":     true,
	"1.0 != 1.0-dev":     true,
	"1.0 lt 1.0a1":       false,
	"1.0 < 1.0a1":        false,
	"1.0 le 1.0a1":       false,
	"1.0 <= 1.0a1":       false,
	"1.0 gt 1.0a1":       true,
	"1.0 > 1.0a1":        true,
	"1.0 ge 1.0a1":       true,
	"1.0 >= 1.0a1":       true,
	"1.0 eq 1.0a1":       false,
	"1.0 = 1.0a1":        false,
	"1.0 == 1.0a1":       false,
	"1.0 ne 1.0a1":       true,
	"1.0 <> 1.0a1":       true,
	"1.0 != 1.0a1":       true,
	"1.0 lt 1.0b1":       false,
	"1.0 < 1.0b1":        false,
	"1.0 le 1.0b1":       false,
	"1.0 <= 1.0b1":       false,
	"1.0 gt 1.0b1":       true,
	"1.0 > 1.0b1":        true,
	"1.0 ge 1.0b1":       true,
	"1.0 >= 1.0b1":       true,
	"1.0 eq 1.0b1":       false,
	"1.0 = 1.0b1":        false,
	"1.0 == 1.0b1":       false,
	"1.0 ne 1.0b1":       true,
	"1.0 <> 1.0b1":       true,
	"1.0 != 1.0b1":       true,
	"1.0 lt 1.0RC1":      false,
	"1.0 < 1.0RC1":       false,
	"1.0 le 1.0RC1":      false,
	"1.0 <= 1.0RC1":      false,
	"1.0 gt 1.0RC1":      true,
	"1.0 > 1.0RC1":       true,
	"1.0 ge 1.0RC1":      true,
	"1.0 >= 1.0RC1":      true,
	"1.0 eq 1.0RC1":      false,
	"1.0 = 1.0RC1":       false,
	"1.0 == 1.0RC1":      false,
	"1.0 ne 1.0RC1":      true,
	"1.0 <> 1.0RC1":      true,
	"1.0 != 1.0RC1":      true,
	"1.0 lt 1.0rc1":      false,
	"1.0 < 1.0rc1":       false,
	"1.0 le 1.0rc1":      false,
	"1.0 <= 1.0rc1":      false,
	"1.0 gt 1.0rc1":      true,
	"1.0 > 1.0rc1":       true,
	"1.0 ge 1.0rc1":      true,
	"1.0 >= 1.0rc1":      true,
	"1.0 eq 1.0rc1":      false,
	"1.0 = 1.0rc1":       false,
	"1.0 == 1.0rc1":      false,
	"1.0 ne 1.0rc1":      true,
	"1.0 <> 1.0rc1":      true,
	"1.0 != 1.0rc1":      true,
	"1.0 lt 1.0":         false,
	"1.0 < 1.0":          false,
	"1.0 le 1.0":         true,
	"1.0 <= 1.0":         true,
	"1.0 gt 1.0":         false,
	"1.0 > 1.0":          false,
	"1.0 ge 1.0":         true,
	"1.0 >= 1.0":         true,
	"1.0 eq 1.0":         true,
	"1.0 = 1.0":          true,
	"1.0 == 1.0":         true,
	"1.0 ne 1.0":         false,
	"1.0 <> 1.0":         false,
	"1.0 != 1.0":         false,
	"1.0 lt 1.0pl1":      true,
	"1.0 < 1.0pl1":       true,
	"1.0 le 1.0pl1":      true,
	"1.0 <= 1.0pl1":      true,
	"1.0 gt 1.0pl1":      false,
	"1.0 > 1.0pl1":       false,
	"1.0 ge 1.0pl1":      false,
	"1.0 >= 1.0pl1":      false,
	"1.0 eq 1.0pl1":      false,
	"1.0 = 1.0pl1":       false,
	"1.0 == 1.0pl1":      false,
	"1.0 ne 1.0pl1":      true,
	"1.0 <> 1.0pl1":      true,
	"1.0 != 1.0pl1":      true,
	"1.0pl1 lt 1.0-dev":  false,
	"1.0pl1 < 1.0-dev":   false,
	"1.0pl1 le 1.0-dev":  false,
	"1.0pl1 <= 1.0-dev":  false,
	"1.0pl1 gt 1.0-dev":  true,
	"1.0pl1 > 1.0-dev":   true,
	"1.0pl1 ge 1.0-dev":  true,
	"1.0pl1 >= 1.0-dev":  true,
	"1.0pl1 eq 1.0-dev":  false,
	"1.0pl1 = 1.0-dev":   false,
	"1.0pl1 == 1.0-dev":  false,
	"1.0pl1 ne 1.0-dev":  true,
	"1.0pl1 <> 1.0-dev":  true,
	"1.0pl1 != 1.0-dev":  true,
	"1.0pl1 lt 1.0a1":    false,
	"1.0pl1 < 1.0a1":     false,
	"1.0pl1 le 1.0a1":    false,
	"1.0pl1 <= 1.0a1":    false,
	"1.0pl1 gt 1.0a1":    true,
	"1.0pl1 > 1.0a1":     true,
	"1.0pl1 ge 1.0a1":    true,
	"1.0pl1 >= 1.0a1":    true,
	"1.0pl1 eq 1.0a1":    false,
	"1.0pl1 = 1.0a1":     false,
	"1.0pl1 == 1.0a1":    false,
	"1.0pl1 ne 1.0a1":    true,
	"1.0pl1 <> 1.0a1":    true,
	"1.0pl1 != 1.0a1":    true,
	"1.0pl1 lt 1.0b1":    false,
	"1.0pl1 < 1.0b1":     false,
	"1.0pl1 le 1.0b1":    false,
	"1.0pl1 <= 1.0b1":    false,
	"1.0pl1 gt 1.0b1":    true,
	"1.0pl1 > 1.0b1":     true,
	"1.0pl1 ge 1.0b1":    true,
	"1.0pl1 >= 1.0b1":    true,
	"1.0pl1 eq 1.0b1":    false,
	"1.0pl1 = 1.0b1":     false,
	"1.0pl1 == 1.0b1":    false,
	"1.0pl1 ne 1.0b1":    true,
	"1.0pl1 <> 1.0b1":    true,
	"1.0pl1 != 1.0b1":    true,
	"1.0pl1 lt 1.0RC1":   false,
	"1.0pl1 < 1.0RC1":    false,
	"1.0pl1 le 1.0RC1":   false,
	"1.0pl1 <= 1.0RC1":   false,
	"1.0pl1 gt 1.0RC1":   true,
	"1.0pl1 > 1.0RC1":    true,
	"1.0pl1 ge 1.0RC1":   true,
	"1.0pl1 >= 1.0RC1":   true,
	"1.0pl1 eq 1.0RC1":   false,
	"1.0pl1 = 1.0RC1":    false,
	"1.0pl1 == 1.0RC1":   false,
	"1.0pl1 ne 1.0RC1":   true,
	"1.0pl1 <> 1.0RC1":   true,
	"1.0pl1 != 1.0RC1":   true,
	"1.0pl1 lt 1.0rc1":   false,
	"1.0pl1 < 1.0rc1":    false,
	"1.0pl1 le 1.0rc1":   false,
	"1.0pl1 <= 1.0rc1":   false,
	"1.0pl1 gt 1.0rc1":   true,
	"1.0pl1 > 1.0rc1":    true,
	"1.0pl1 ge 1.0rc1":   true,
	"1.0pl1 >= 1.0rc1":   true,
	"1.0pl1 eq 1.0rc1":   false,
	"1.0pl1 = 1.0rc1":    false,
	"1.0pl1 == 1.0rc1":   false,
	"1.0pl1 ne 1.0rc1":   true,
	"1.0pl1 <> 1.0rc1":   true,
	"1.0pl1 != 1.0rc1":   true,
	"1.0pl1 lt 1.0":      false,
	"1.0pl1 < 1.0":       false,
	"1.0pl1 le 1.0":      false,
	"1.0pl1 <= 1.0":      false,
	"1.0pl1 gt 1.0":      true,
	"1.0pl1 > 1.0":       true,
	"1.0pl1 ge 1.0":      true,
	"1.0pl1 >= 1.0":      true,
	"1.0pl1 eq 1.0":      false,
	"1.0pl1 = 1.0":       false,
	"1.0pl1 == 1.0":      false,
	"1.0pl1 ne 1.0":      true,
	"1.0pl1 <> 1.0":      true,
	"1.0pl1 != 1.0":      true,
	"1.0pl1 lt 1.0pl1":   false,
	"1.0pl1 < 1.0pl1":    false,
	"1.0pl1 le 1.0pl1":   true,
	"1.0pl1 <= 1.0pl1":   true,
	"1.0pl1 gt 1.0pl1":   false,
	"1.0pl1 > 1.0pl1":    false,
	"1.0pl1 ge 1.0pl1":   true,
	"1.0pl1 >= 1.0pl1":   true,
	"1.0pl1 eq 1.0pl1":   true,
	"1.0pl1 = 1.0pl1":    true,
	"1.0pl1 == 1.0pl1":   true,
	"1.0pl1 ne 1.0pl1":   false,
	"1.0pl1 <> 1.0pl1":   false,
	"1.0pl1 != 1.0pl1":   false,
}

func TestCompareVersion(t *testing.T) {
	for in, out := range compareVersionValues {
		v := strings.Split(in, " ")
		if x := Compare(v[0], v[2], v[1]); x != out {
			t.Errorf("FAIL: CompareVersion(%v) = %v: want %v", in, x, out)
		}
	}
}
