package main

func main() {
	numbers := readNumbers("input-1.txt")

	for _, n1 := range numbers {
		for _, n2 := range numbers {
			for _, n3 := range numbers {
				if n1+n2+n3 == 2020 {
					println(n1 * n2 * n3)
					return
				}
			}
		}
	}
}
