package main

func main() {
	groupCity := map[int][]string{
		10:   []string{}, // города с населением 10-99 тыс. человек
		100:  []string{}, // города с населением 100-999 тыс. человек
		1000: []string{}, // города с населением 1000 тыс. человек и более
	}
	cityPopulation := map[string]int{}

	badCities := append(groupCity[10], groupCity[1000]...) // []string
	for city, _ := range cityPopulation {
		for _, badCity := range badCities {
			if city == badCity {
				delete(cityPopulation, city)
			}
		}
	}
}
