package UltimateDivMod

func UltimateDivMod(a *int, b *int) {
	mytmp := *a
	*a = *a / *b
	*b = mytmp % *b
}
