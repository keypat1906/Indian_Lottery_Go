package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func checkerror(e error) {

	if e != nil {
		panic(e)
	}
}

// Intersection function to intersect two arrays

func intersect(slice1, slice2 []int) []int {
	var intersect []int
	for _, element1 := range slice1 {
		for _, element2 := range slice2 {
			if element1 == element2 {
				intersect = append(intersect, element1)
			}
		}
	}
	return intersect //return slice after intersection
}

// Main Function to pick the winners

func main() {

	t1 := time.Now()
	dat, err := os.Open("/Users/bracelet/Downloads/test.txt")
	checkerror(err)

	defer dat.Close()

	myDict := make(map[int]int)
	fmt.Println("Enter the numbers ")

	slice1 := make([]int, 5)
	for i := 0; i < 5; i++ {
		fmt.Scan(&slice1[i])
	}

	sc := bufio.NewScanner(dat)
	for sc.Scan() {
		var mystr = sc.Text() // GET the line string

		mystr2 := strings.Replace(mystr, "\\", "", -1)
		strs := strings.Split(mystr2, " ")
		slice2 := make([]int, len(strs))

		for i := range strs {
			slice2[i], _ = strconv.Atoi(strs[i])
		}

		intersect := intersect(slice1, slice2)

		if len(intersect) == 0 {
			continue
		}
		myDict[len(intersect)] = myDict[len(intersect)] + 1
	}

	fmt.Println("--------------------------------------")
	fmt.Println("| Numbers matching   |      Winners   ")
	fmt.Println("--------------------------------------")
	fmt.Println("| 5                  |      ", myDict[5])
	fmt.Println("| 4                  |      ", myDict[4])
	fmt.Println("| 3                  |      ", myDict[3])
	fmt.Println("| 2                  |      ", myDict[2])
	fmt.Println("--------------------------------------")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	t2 := time.Now()
	diff := t2.Sub(t1)
	fmt.Println("Total time taken", diff)

}
