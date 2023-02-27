package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(3) + 3
	fmt.Printf("n = %d\n", n)

	table := make([][]int, n)
	for i := range table {
		table[i] = make([]int, n)
		for j := range table[i] {
			table[i][j] = rand.Intn(100)
		}
	}

	fmt.Println(table)

	fmt.Println("<table>")
	for i := range table {
		fmt.Println("  <tr>")
		for j := range table[i] {
			fmt.Printf("    <td>%d</td>\n", table[i][j])
		}
		fmt.Println("  </tr>")
	}
	fmt.Println("</table>")
}
