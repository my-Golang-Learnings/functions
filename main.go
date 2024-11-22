package main

func main() {
	// > Assigning variables with ':=' allows the compiler to infer their type
	firstString := "hey"
	// > or
	// firstString string = "hey"
	// > defines the type explicitly so i can use the more typical '=' to assign a value
	seconsString := "how are you"

	output := concatStrings(firstString, seconsString)
	println(concatStrings("heres your output :", output))
}

// > Funtion return types are stated after the function
func concatStrings(String1, String2 string) string {
	return (String1 + String2)
}
