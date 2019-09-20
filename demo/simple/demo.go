package main

import (
	"github.com/Pingze-github/go-package-plantuml/codeanalysis"
	log "github.com/Sirupsen/logrus"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	log.SetLevel(log.InfoLevel)

	//config := codeanalysis.Config{
	//	CodeDir: "/appdev/go-demo/src/github.com/Pingze-github/go-package-plantuml/testdata/a",
	//	GopathDir : "/appdev/go-demo",
	//}
	//
	//result := codeanalysis.AnalysisCode(config)
	//
	//result.OutputToFile("/tmp/uml.txt")
	//"/appdev/go-demo"
	//--gopath $GOPATH$ --codedir $FileDir$ --outputfile $FileDir$/$FileDirName$.puml
	dir, _ := os.Getwd()
	config := codeanalysis.Config{
		//C:\Users\lizhe\go\src\github.com\nsqio\nsq\apps\nsq_to_http
		CodeDir:   `C:\Users\lizhe\go\src\github.com\Pingze-github\go-package-plantuml\testdata\haha.go`,
		GopathDir: `C:\Users\lizhe\go`,
		OutputFile:dir+"\\uml.puml",//`C:\Users\lizhe\go\src\github.com\Pingze-github\go-package-plantuml\demo\simple\uml.puml`,
	}

	result := codeanalysis.AnalysisCode(config)

	result.OutputToFile(config.OutputFile)

	bytes, _ := ioutil.ReadFile(config.OutputFile)

	fmt.Println(string(bytes))

	// java -jar /app/plantuml.jar  /tmp/uml.txt -tsvg && open2 /tmp/uml.svg

}
