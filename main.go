package main

import (
	"fmt"
	"sort"
)

type Typing struct {
	x    int
	y    int
	next string
}

func main() {
	var userInput = &Typing{next: "O"}
	var table = [3][3]string{{"-", "-", "-"}, {"-", "-", "-"}, {"-", "-", "-"}}
	var isOver = false
	var recordX = []int{}
	var recordO = []int{}
	var count = 1
	for isOver == false {
		// get user input
		fmt.Printf("------------- Round %v ---------------\n", count)
		userInput.catchInputAndStore(&table, &recordX, &recordO)
		renderTable(&table)
		checkOver(&isOver, recordX, recordO)
		count++
	}

}

func renderTable(table *[3][3]string) {
	fmt.Println("    =============")
	fmt.Printf("    | %s | %s | %s |\n", table[0][0], table[0][1], table[0][2])
	fmt.Println("    =============")
	fmt.Printf("    | %s | %s | %s |\n", table[1][0], table[1][1], table[1][2])
	fmt.Println("    =============")
	fmt.Printf("    | %s | %s | %s |\n", table[2][0], table[2][1], table[2][2])
	fmt.Println("    =============")
	fmt.Println()
}

func checkOver(isOver *bool, recordX, recordO []int) bool {
	if len(recordX)+len(recordO) == 9 {
		fmt.Println("No Winner..")
		*isOver = true
	}
	sort.Ints(recordX[:])
	sort.Ints(recordO[:])
	if checkRecord(recordX) {
		fmt.Println("X is Winner")
		*isOver = true
	}
	if checkRecord(recordO) {
		fmt.Println("O is Winner")
		*isOver = true
	}
	return *isOver
}

func checkRecord(record []int) bool {
	var checkOne []int
	var checkTen []int
	for i := range record {
		ones := record[i] % 10
		tens := record[i] / 10
		checkOne = append(checkOne, ones)
		checkTen = append(checkTen, tens)
	}
	// fmt.Println("checkOne: ", checkOne)
	// fmt.Println("checkTen: ", checkTen)
	if checkOrder(checkOne) && checkOrder(checkTen) {
		return true
	}
	return false
}

func checkOrder(arr []int) bool {
	// fmt.Println("arr: ", arr)
	var res = false
	var orderCount = 0
	var orderCountDesc = 0
	var sameCount = 0
	for i := 1; i < len(arr); i++ {
		if arr[i]-1 == arr[i-1] {
			res = true
			orderCount++
		} else if arr[i] == arr[i-1] {
			res = true
			sameCount++
		} else if arr[i] == arr[i-1]-1 {
			res = true
			orderCountDesc++
		}
	}
	if orderCount != 2 && sameCount != 2 && orderCountDesc != 2 {
		res = false
	}
	// fmt.Println("orderCount: ", orderCount)
	// fmt.Println("sameCount: ", sameCount)
	// fmt.Println("res: ", res)
	return res
}

func (t *Typing) catchInputAndStore(table *[3][3]string, recordX, recordO *[]int) {
	// use Scanln
	// ----------------------------------------------------------------
	var first, second int
	fmt.Printf("Current User is %s \n", t.next)
	fmt.Println("x and y must in the range of 1 to 3")
	fmt.Println("Please input follow this format => x space y ")
	fmt.Print("Your input: ")
	_, err := fmt.Scanln(&first, &second)
	if err != nil {
		fmt.Println("Input error: ", err)
	}

	if first > 3 || second > 3 {
		fmt.Println("x and y must in the range of 1 to 3")
		fmt.Println("-----------------------------------")
		t.catchInputAndStore(table, recordX, recordO)
	} else if first < 1 || second < 1 {
		fmt.Println("x and y must in the range of 1 to 3")
		fmt.Println("-----------------------------------")
		t.catchInputAndStore(table, recordX, recordO)
	} else {
		var sum = first*10 + second
		first = first - 1
		second = second - 1
		if table[first][second] == "-" {
			table[first][second] = t.next
			if t.next == "O" {
				t.next = "X"
				*recordO = append(*recordO, sum)
			} else {
				t.next = "O"
				*recordX = append(*recordX, sum)
			}
		} else {
			fmt.Println("This position is already used")
			fmt.Println("-----------------------------------")
			t.catchInputAndStore(table, recordX, recordO)
		}
	}
}
