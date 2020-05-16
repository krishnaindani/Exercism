package letter

// FreqMap keep track to rune to frequency of it occuring
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns FreqMap
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

//ConcurrentFrequency takes a slice of string and run the Frequency function concurrently returning FreqMap
func ConcurrentFrequency(s []string) FreqMap {
	ch := make(chan FreqMap, 10)
	f := FreqMap{}
	for _, word := range s {
		go func(s string) {
			ch <- Frequency(s)
		}(word)
	}

	for range s {
		for key, value := range <-ch {
			f[key] += value
		}
	}

	return f
}
