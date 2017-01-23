package ddz

// https://compprog.wordpress.com/2007/10/17/generating-combinations-1/
// next_comb(int comb[], int k, int n)
// Generates the next combination of n elements as k after comb
//
// comb => the previous combination ( use (0, 1, 2, ..., k) for first)
// k => the size of the subsets to generate
// n => the size of the original set
//
// Returns: 1 if a valid combination was found
// 0, otherwise

func NextCombination(comb []int, k, n int) bool {
	i := k - 1
	comb[i]++

	for i > 0 && (comb[i] >= n-k+1+i) {
		i--
		comb[i]++
	}

	if comb[0] > n-k {
		return false
	}

	for i = i + 1; i < k; i++ {
		comb[i] = comb[i-1] + 1
	}

	return true
}
