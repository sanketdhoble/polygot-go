package main

import (
	"fmt"
	_ "fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func isTwoSidedPrime(number int) bool {
	isPrime := true
	num := strconv.Itoa(number)
	// left trimming
	for len(num) > 0 {
		if checkPrime(num) {
			num = truncateLeftDigits(num)
		} else {
			isPrime = false
			break
		}
	}
	//right trimming
	for len(num) > 0 {
		if checkPrime(num) {
			num = truncateLeftDigits(num)
		} else {
			isPrime = false
			break
		}
	}
	return isPrime
}

func truncateLeftDigits(number string) string {
	num := number[1:]
	return num
}

func truncateRightDigits(number string) string {
	sz := len(number)
	num := number[:sz-1]
	return num
}

func checkPrime(number string) bool {
	num, err := strconv.Atoi(number)
	_ = err
	result := true
	if num <= 1 {
		result = false
	} else {
		if num <= 3 {
			result = true
		} else {
			if (num%2 == 0) || (num%3 == 0) {
				result = false
			} else {
				i := 5
				for (i * i) <= num {
					if ((num % i) == 0) || (num%(i+2) == 0) {
						result = false
						break
					}
					i = i + 6
				}
			}
		}
	}

	return result
}

func TwoSidedPrimeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strNum := vars["num"]
	num, err := strconv.ParseInt(strNum, 10, 32)
	if err != nil {
		fmt.Fprintln(w, "invalid integer")
	} else {
		fmt.Fprintln(w, isTwoSidedPrime(int(num)))
	}
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/isTwoSidedPrime/{num}", TwoSidedPrimeHandler)

	http.ListenAndServe(":8080", router)
}
