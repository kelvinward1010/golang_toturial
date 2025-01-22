// 1. With the var keyword
// var variablename type = value

//2. With the := sign
// variablename := value

package variables

import (
	"fmt"
)

func Variables() {
	var student1 string = "John" //type is string
	var student2 = "Jane"        //type is inferred
	x := 2                       //type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
}
