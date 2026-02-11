package main

func getName() string {
	var name string
	print("Enter your name: ")
	fmt.Scanln(&name)
	name = strings.TrimSpace(name)
	if !validateName(name) {
		return getName()
	}
	return name
}