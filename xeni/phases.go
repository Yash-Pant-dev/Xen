package xeni

import (
)

func XenCompiler(sourceFilePath string) any {
	
	abstractTree, _ := Parse(sourceFilePath)
	
	return abstractTree
}

func XenOptimizer(abstractTree any) any {

	optimizedTree, _ := Optimize(abstractTree)

	return optimizedTree
}

func XenVM(abstractTree any) error {

	_ = Execute(abstractTree)

	return nil
}