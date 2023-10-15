package components

import "fmt"

func Test(name, sex string, age int) string {

	return fmt.Sprintf(`
		<div>My name is: %s</div>
		<div>sex : %s</div>
		<div>age: %d</div>
	`, name, sex, age)
}
