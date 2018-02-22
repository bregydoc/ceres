package templatetype

import (
	"fmt"

	"github.com/kataras/iris"
)

const apiVersion = "1.0"

func main() {
	app := iris.Default()

	apiParty := app.Party(fmt.Sprintf("/api/v%s", apiVersion))

	LinkWithUserType(apiParty)

	app.Run(iris.Addr(":8080"))

}
