package main

import "fmt"

func main() {

	language := make(map[string]string, 4)

	language["C"] = "c"
	language["py"] = "python"
	language["js"] = "javascript"

	fmt.Println(language)
	//delete
	delete(language, "py")
	fmt.Println(language)

	//Looping map
	// We can use _ to ignore key/value
	for key, value := range language {

		if key == "js" {
			fmt.Println("Javascript: Node")
		}
		fmt.Println(key, value)
	}

}
