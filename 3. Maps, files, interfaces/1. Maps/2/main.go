package main

func main() {

	groupCity := map[int][]string{
		10:   []string{"10"},   // города с населением 10-99 тыс. человек
		100:  []string{"100"},  // города с населением 100-999 тыс. человек
		1000: []string{"1000"}, // города с населением 1000 тыс. человек и более
	}

	cityPopulation := map[string]int{
		"city": 100,
	}

	for city := range cityPopulation {
		check := false
		for _, groupCities := range groupCity[100] {
			if city == groupCities {
				check = true
				break
			}
		}
		if !check {
			delete(cityPopulation, city)
		}
	}

}
