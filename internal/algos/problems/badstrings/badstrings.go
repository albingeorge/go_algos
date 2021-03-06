package badstrings

var vowels = []string{"a", "e", "i", "o", "u"}

const (
	vowel     = "vowel"
	consonant = "consonant"
)

/*
Badstrings ...Find bad strings
Returns 1 if good, 0 if bad
*/
func Badstrings(input string) int {
	cCount, vCount := 0, 0
	prev := ""
	for _, v := range input {
		s := string(v)

		if s == "?" {
			cCount++
			vCount++
		} else if contains(vowels, s) == true {
			if prev == consonant {
				cCount, vCount = 0, 0
			}
			prev = vowel
			vCount++
		} else {
			if prev == vowel {
				cCount, vCount = 0, 0
			}
			prev = consonant
			cCount++
		}

		if cCount > 3 || vCount > 5 {
			return 0
		}
	}

	return 1
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
