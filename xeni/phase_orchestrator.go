package xeni

import (
	"fmt"
	"xen/logger"
)

var debugMode = true
var compileOnlyMode = false
var executionOnlyMode = false

const logsFileName = "xeni_logs.log" 
const metricsFileName = "xeni_metrics.log"

var Log func(int, ...any)

/*
	Parse params & execute relevant phases
	The typical stages are : Compilation -> Optimization -> Execution
 */
func OrchestratePhases(params []string) {

	Log = logger.InitializeLogging(debugMode, logsFileName, metricsFileName)
	
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

/* 
	Flags:
	1. -d : Debug mode.
	2. -c : Only compiles, does not optimize or run.
	3. -e : Only executes the abstract tree in the given file.
*/
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
			logger.Log(4, "Flag does not exist", flag)
		}
	}
}
