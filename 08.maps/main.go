package main

import "fmt"

func main() {
	idiomas := make(map[string]string)
	idiomas["es"] = "Español"
	idiomas["en"] = "Ingles"
	idiomas["fr"] = "Frances"
	fmt.Println(idiomas)
	/*fmt.Println(idiomas["es"])
	fmt.Println(idiomas["en"])
	fmt.Println(idiomas["fr"])
	*/
	delete(idiomas, "en")
	fmt.Println(idiomas)
	fmt.Println(idiomas["br"])

}
