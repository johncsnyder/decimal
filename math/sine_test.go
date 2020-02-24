package math

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/johncsnyder/decimal"
)

func TestSin(t *testing.T) {
	const N = 100
	diff := new(decimal.Big)
	eps := new(decimal.Big)
	for i, tt := range [...]struct {
		x, r string
	}{
		0:  {"0", "0"},
		1:  {pos(_pi_4, N), "0.7071067811865475244008443621048490392848359376884740365883398689953662392310535194251937671638207863"},
		2:  {neg(_pi_4, N), "-0.7071067811865475244008443621048490392848359376884740365883398689953662392310535194251937671638207863"},
		3:  {pos(_pi_3, N), "0.8660254037844386467637231707529361834714026269051903140279034897259665084544000185405730933786242877"},
		4:  {neg(_pi_3, N), "-0.8660254037844386467637231707529361834714026269051903140279034897259665084544000185405730933786242877"},
		5:  {pos(_pi_2, N), "1.000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"},
		6:  {neg(_pi_2, N), "-1.000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"},
		7:  {pos(_3pi_4, N), "0.7071067811865475244008443621048490392848359376884740365883398689953662392310535194251937671638207871"},
		8:  {neg(_3pi_4, N), "-0.7071067811865475244008443621048490392848359376884740365883398689953662392310535194251937671638207871"},
		9:  {pos(_pi, N), "9.821480865132823066470938446095505822317253594081284811174502841027019385211055596446229489549303820E-N"},
		10: {neg(_pi, N), "-9.821480865132823066470938446095505822317253594081284811174502841027019385211055596446229489549303820E-N"},
		11: {pos(_5pi_4, N), "-0.7071067811865475244008443621048490392848359376884740365883398689953662392310535194251937671638207857"},
		12: {neg(_5pi_4, N), "0.7071067811865475244008443621048490392848359376884740365883398689953662392310535194251937671638207857"},
		13: {pos(_3pi_2, N), "-1.000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"},
		14: {neg(_3pi_2, N), "1.000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"},
		15: {pos(_2pi, N), "-9.642961730265646132941876892191011644634507188162569622349005682054038770422111192892458979098607639E-N"},
		16: {neg(_2pi, N), "9.642961730265646132941876892191011644634507188162569622349005682054038770422111192892458979098607639E-N"},
		17: {"7.3303828583761842231", "0.86602540378443864677"},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			x, _ := new(decimal.Big).SetString(tt.x)
			r, _ := new(decimal.Big).SetString(tt.r)
			z := decimal.WithPrecision(r.Precision())

			Sin(z, x)
			diff.Context.Precision = z.Context.Precision
			eps.SetMantScale(1, z.Context.Precision)
			if z.Cmp(r) != 0 && diff.Sub(r, z).CmpAbs(eps) > 0 {
				t.Errorf(`#%d: Sin(%s)
wanted: %s
got   : %s
diff  : %s
`, i, x, r, z, diff)
			}
		})
	}
}

var sin_X, _ = new(decimal.Big).SetString("0.7853981633974483096156608458198757210492923498437764552437361480769541015715522496570087063355292670")

func BenchmarkSin(b *testing.B) {
	for _, prec := range benchPrecs {
		b.Run(fmt.Sprintf("%d", prec), func(b *testing.B) {
			z := decimal.WithPrecision(prec)
			for j := 0; j < b.N; j++ {
				Sin(z, sin_X)
			}
			gB = z
		})
	}
}
