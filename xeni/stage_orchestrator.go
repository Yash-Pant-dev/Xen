package xeni

import (
	"fmt"
)

var debugMode = true
var compileOnlyMode = false
var executionOnlyMode = false

// Parse params & Execute relevant stages
func OrchestrateStages(params []string) {

	initializeLogging()

	switch len(params) {
	case 0:
		panic(fmt.Sprintf("[%v]", E_Orch_InsufficientArgs))
	case 1:
		XenVM(XenOptimizer(XenCompiler(params[0])))
	default:
		parseFlags(params[1:])
		if compileOnlyMode {
			XenCompiler(params[0])
		}
		if executionOnlyMode {
			XenVM(params[0])
		}
		if !compileOnlyMode && !executionOnlyMode {
			XenVM(XenOptimizer(XenCompiler(params[0])))
		}
	}

}

func parseFlags(flags []string) {
	for _, flag := range flags {

		switch flag {
		case "-d": // debug mode
			debugMode = true
		case "-c": // compile only
			compileOnlyMode = true
		case "-e": // execution only
			executionOnlyMode = true
		default:
			Log(4, "Flag does not exist", flag)
		}

	}
}
