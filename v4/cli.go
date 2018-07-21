package main

import (
	"os"
	"flag"
)

const Usage =`
	./exe addBlock --data DATA	"add a block to block chain"
	./exe printchain				"print all blocks"
`

type CLI struct {
	bc *BlockChain
}

func(cli *CLI)Run(){
	if len(os.Args)<2{
		println("too few parameters!", Usage)
		os.Exit(1)
	}
	//func NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet {
	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	//func (f *FlagSet) String(name string, value string, usage string) *string {
	//获取指定参数，
	//data ：命令行参数
	//value :默认值
	//usage ：描述
	addBlockCmdPara := addBlockCmd.String("data", "", "block info")
	switch os.Args[1] {
	case "addBlock":
		err := addBlockCmd.Parse(os.Args[2:])
		CheckErr(err)
		if addBlockCmd.Parsed(){
			if *addBlockCmdPara == ""{
				println("data is emty, please check!")
			}
			//cli.bc.AddBlock(*addBlockCmdPara)
			cli.AddBlock(*addBlockCmdPara)
		}
	case "printchain":
		err := printCmd.Parse(os.Args[2:])
		CheckErr(err)
		if printCmd.Parsed(){
			cli.PrintChain()
		}
	default:
		println("invalid cmd\n", Usage)
		os.Exit(1)
	}
}
