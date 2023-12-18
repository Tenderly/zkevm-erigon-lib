package compress

func convertOldSortOrderToNew[E any](cmp func(a, b E) bool) func(a, b E) int {
	return func(a, b E) int {
		if cmp(a, b) {
			return -1
		}
		if cmp(b, a) {
			return 1
		}
		return 0
	}
}
