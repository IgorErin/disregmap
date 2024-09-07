package common


func Fmap[A any, B any](f func (A) B, sl []A) []B {
	res := make([]B, len(sl))
	for ind, item := range sl {
		res[ind] = f(item)
	}
	return res
}