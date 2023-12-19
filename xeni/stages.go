package xeni

import (
	parser "xen/xeni/parser"
	"xen/xeni/optimizer"
	"xen/xeni/vm"
)

func XenCompiler(sourceFilePath string) any {
	abstractTree, _ := parser.Parse(sourceFilePath)

	return abstractTree
}

func XenOptimizer(abstractTree any) any {

	optimizedTree, _ := optimizer.Optimize(abstractTree)

	return optimizedTree
}

func XenVM(abstractTree any) error {

	_ = vm.Execute(abstractTree)

	return nil
}