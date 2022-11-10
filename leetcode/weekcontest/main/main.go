package main

func main() {
	m := Constructor([]string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"}, []string{"korean", "japanese", "japanese", "greek", "japanese", "korean"}, []int{9, 12, 8, 15, 14, 7})
	m.ChangeRating("sushi", 16)
	m.HighestRated("japanese")
}
