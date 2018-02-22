package templatetype

import (
	"fmt"

	"github.com/kataras/iris"
)

const apiVersion = "1.0"

func main() {
	app := iris.Default()

	app.Party(fmt.Sprintf("/api/%s", apiVersion))
}
