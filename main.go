package main

import (
	"fmt"
	"math/rand"
	"os"
)

var AllTickets [][5][5]int // All tickets enter manually in count

var Bingo [30]int // the numbers that fell out
var Winners []int

func main() {

	var count int // amount of tickets
	fmt.Println("Enter the amount of tickets")
	fmt.Scanf("%d", &count)

	bingo()

	fmt.Println("Bingo = ", Bingo)
	for i := 0; i < count; i++ {
		createBilet()
		fmt.Println("Ticket ", i+1, "= ", AllTickets[i])
		checkTicket(i)
		f, err := os.Create(fmt.Sprintf("Ticket%d.txt", i+1))
		if err != nil {
			panic(err)
		}
		defer f.Close()
		f.WriteString((fmt.Sprintln(AllTickets[i])))
	}

	fmt.Println("Winners tickets = ", Winners)
	fmt.Println("Number of tickets = ", count)
	fmt.Println("Number of Winners = ", len(Winners))
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func createBilet() {
	var Bilet [5][5]int
	var slice []int
	for i := 1; i <= 90; i++ {
		slice = append(slice, i)
	}

	for i := range Bilet {
		for j := range Bilet[i] {

			a := rand.Intn(89 - j - i*5)
			Bilet[i][j] = slice[a]
			slice = remove(slice, a)

		}
	}
	AllTickets = append(AllTickets, Bilet)
	return
}

func bingo() {
	var slice2 []int // all numbers

	for i := 1; i <= 90; i++ {
		slice2 = append(slice2, i)
	}
	for i := 0; i < 30; i++ {
		a := rand.Intn(89 - i)
		Bingo[i] = slice2[a]
		slice2 = remove(slice2, a)
	}

}

func checkTicket(i int) []int {
	var coincidence []int
	var checkWinner [5][5]bool
	currentTicket := AllTickets[i]
	for a := range Bingo {
		for j := range currentTicket {
			for k := range currentTicket[j] {
				if Bingo[a] == currentTicket[j][k] {
					coincidence = append(coincidence, Bingo[a])
					checkWinner[j][k] = true
				}
			}
		}

	}
	for n := 0; n < 5; n++ {

		if checkWinner[n][0] && checkWinner[n][1] && checkWinner[n][2] && checkWinner[n][3] && checkWinner[n][4] == true {
			Winners = append(Winners, i+1)

		}
		if checkWinner[0][n] && checkWinner[1][n] && checkWinner[2][n] && checkWinner[3][n] && checkWinner[4][n] == true {
			Winners = append(Winners, i+1)

		}

	}
	if checkWinner[0][0] && checkWinner[1][1] && checkWinner[2][2] && checkWinner[3][3] && checkWinner[4][4] == true {
		Winners = append(Winners, i+1)

	}
	if checkWinner[0][4] && checkWinner[1][3] && checkWinner[2][2] && checkWinner[3][1] && checkWinner[4][0] == true {
		Winners = append(Winners, i+1)

	}
	fmt.Println("Coincidence numbers = ", coincidence)
	fmt.Println(checkWinner)
	fmt.Println("Number of coincidences = ", len(coincidence))
	return coincidence
}
