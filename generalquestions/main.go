package generalquestions

// An Armstrong number is one whose sum of digits raised to the power three equals the number itself.
// 371, for example, is an Armstrong number because 3**3 + 7**3 + 1**3 = 371.

// Anagram: a word, phrase, or name formed by rearranging the letters of another, such as spar, formed from rasp.

func GetNthFibonacii(n int) int {

	if n == 2 {
		return 1
	}
	if n == 1 {
		return 0
	}

	// fmt.Println(n)
	return GetNthFibonacii(n-1) + GetNthFibonacii(n-2)

}

func palindromeCheck(str string) bool {

	strarray := []rune(str)

	for i, j := 0, len(strarray)-1; i <= j; {

		if strarray[i] == strarray[j] {
			i++
			j--
		} else {
			return false
		}

	}

	return true
}

func isPrime(n int) bool {

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
