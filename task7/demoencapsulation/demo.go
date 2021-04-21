package demoencapsulation

import "fmt"

const ExportedConst = "ExportedConst is exported"

var unexprotedVar = "unexportedVar is not available outside of its package"

func ExprtedFunc() {
	fmt.Println("ExportedFunc is available to other packages")
	unexprotedFunc()
}

func unexprotedFunc() {
	fmt.Println("unexportedFunc is only available to its package")
	fmt.Println(unexprotedVar)
}
