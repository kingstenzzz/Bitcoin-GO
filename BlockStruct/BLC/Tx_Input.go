package BLC

type TxInput struct {
	Vout      int    //引用上一笔脚底饿输出索引
	TxHash    []byte //交易Hash
	ScriptSig string
}

//验证引用得地址是否匹配

func (txInput *TxInput) CheckPubkeyWithAddress(address string) bool {
	return address == txInput.ScriptSig

}
