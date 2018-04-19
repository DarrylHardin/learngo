package main

import "fmt"

func main() {
	name := "Darryl Hardin"
	tpl := `<!DOCTYPE html>
<html>
<head>
	<title></title>
</head>
<body>
 <h1>` + name + `</h1>
</body>
</html>`

	fmt.Println(tpl)
}
