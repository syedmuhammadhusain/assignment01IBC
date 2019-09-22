package main
package assignment01IBC

import (
a1 "github.com/syedmuhammadhusain/assignment01IBC"
)

func main() {
	var chainHead *a1.Block
	chainHead = a1.InsertBlock("GenesisBlock", nil)
	chainHead = a1.InsertBlock("AliceToBob", chainHead)
	chainHead = a1.InsertBlock("BobToCharlie", chainHead)
	a1.ListBlocks(chainHead)
	a1.ChangeBlock("AliceToBob", "AliceToTrudy", chainHead)
	a1.ListBlocks(chainHead)
	a1.VerifyChain(chainHead)

}
