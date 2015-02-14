package decomposition

import (
	"testing"

	"github.com/ready-steady/support/assert"
)

func TestSymEig(t *testing.T) {
	m := uint(5)

	A := []float64{
		0.814723686393179,
		0.097540404999410,
		0.157613081677548,
		0.141886338627215,
		0.655740699156587,
		0.097540404999410,
		0.278498218867048,
		0.970592781760616,
		0.421761282626275,
		0.035711678574190,
		0.157613081677548,
		0.970592781760616,
		0.957166948242946,
		0.915735525189067,
		0.849129305868777,
		0.141886338627215,
		0.421761282626275,
		0.915735525189067,
		0.792207329559554,
		0.933993247757551,
		0.655740699156587,
		0.035711678574190,
		0.849129305868777,
		0.933993247757551,
		0.678735154857773,
	}

	eigenVectors := []float64{
		+0.200767588469279,
		-0.613521879994358,
		+0.529492579537623,
		+0.161735212201923,
		-0.526082320114459,
		-0.241005628008408,
		-0.272281143378657,
		+0.443280672960843,
		-0.675165120368165,
		+0.464148221418878,
		+0.509762909240926,
		+0.555609456752178,
		+0.244072927029371,
		-0.492754485897426,
		-0.359251069377747,
		-0.766321363493223,
		+0.386556170387878,
		+0.341170928524320,
		+0.084643789583352,
		-0.373849864790357,
		+0.233456648876442,
		+0.302202482503382,
		+0.589211894835079,
		+0.517708631263932,
		+0.488854547655902,
	}

	eigenValues := []float64{
		-0.671640666831794,
		-0.230366398529950,
		+0.397221322493687,
		+0.999582068576074,
		+3.026535012212483,
	}

	U := make([]float64, m*m)
	Λ := make([]float64, m*1)

	SymEig(A, U, Λ, 5)

	assert.AlmostEqual(U, eigenVectors, t)
	assert.AlmostEqual(Λ, eigenValues, t)
}
