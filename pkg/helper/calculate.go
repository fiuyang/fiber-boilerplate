package helper

func CalculateQty(initial, deduct int64) int64 {
	if deduct != 0 {
		return initial - deduct
	}
	return initial
}
