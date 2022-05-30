package math

func Abs[T RealNumber](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
