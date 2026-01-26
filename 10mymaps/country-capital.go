package main

func countryCapital(country string) string {
	ctry := make(map[string]string)

	ctry["Ghana"] = "Accra"
	ctry["France"] = "Paris"
	ctry["Togo"] = "Lome"
	ctry["USA"] = "Washington DC"

	capital := ctry[country]
	
	return capital
}
