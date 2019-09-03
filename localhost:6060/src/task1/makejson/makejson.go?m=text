import (
	"fmt"
	"encoding/json"
)

func main() {
	mapp := make(map[string]string, 2)

	var name string
	var address string

	fmt.Print("Type the name: ")
	fmt.Scan(&name)
	fmt.Print("Type the address: ")
	fmt.Scan(&amp;address)

	mapp["name"] = name
	mapp["address"] = address

	barr, err := json.Marshal(mapp)

	if err == nil{
		fmt.Println("The JSON object is: ", string(barr))
	}
}
