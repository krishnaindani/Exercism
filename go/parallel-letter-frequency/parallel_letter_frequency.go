package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

//ConcurrentFrequency ...
func ConcurrentFrequency(s []string) FreqMap {
	ch := make(chan FreqMap, len(s))
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
