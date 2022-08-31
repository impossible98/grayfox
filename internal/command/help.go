package command

import (
	// import built-in packages
	"fmt"
	"strings"
	// import local packages
	"github.com/impossible98/grayfox/global"
)

var serverCommandHelp = "\tserve\t\tStart the server\n"
var helpHelp = "\t--help\t\tShow help\n"
var versionHelp = "\t--version\tShow version\n"

func Help() {
	fmt.Println(strings.ToLower(global.APP_NAME) + " help:\n\n" + serverCommandHelp + "\n" + helpHelp + versionHelp)
}
