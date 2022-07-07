package main
import (
	"github.com/goccy/go-yaml"
)


func main() {
	json := `{"test1":1,"test2":"abc"}`
	actual,err := yaml.JSONToYAML([]byte(json))
 	if err != nil {
		println(err)
	}
	println("yaml:",string(actual))


	yml := `
	test1:
	- ipaddr:"123"
	- data:1
	test2:
	- ipaddr:"456"
	- data:2
        `
	y_actual,yerr := yaml.YAMLToJSON([]byte(yml))
 	if yerr != nil {
		println(yerr)
	}
	println("json:",string(y_actual))


}
