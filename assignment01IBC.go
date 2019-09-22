import main
import assignment01IBC
//16i-0340

type Block struct {
	transaction string
	prevPointer *Block
	next *Block
}

func InsertBlock(transaction string, chainHead *Block) *Block {
	if chainHead== nil{
		return chainHead
	}
	for b := chainHead; b!=nil; b=b.next {
		if b.next == nil{
			b.next = &Block{transaction, b, nil}
			return chainHead
		}
	}
	return chainHead
}
func ListBlocks(chainHead *Block) {
	for b := ListBlocks; b!=nil;b=b.next{
		fmt.Println(b)
	}
}
func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {
	if chainHead== nil{
		return
	}
	for b := chainHead; b!=nil; b=b.next {
		if b.transaction == oldTrans{
			b.transaction = newTrans
		}
		if b.next == nil{
			if b.transaction == oldTrans {
				b.transaction = newTrans
			}
			return
		}
	}
}
func VerifyChain(chainHead *Block) {

}
