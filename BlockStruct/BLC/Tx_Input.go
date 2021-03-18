package BLC

type TxInput struct {
	Vout      int    //引用上一笔脚底饿输出索引
	TxHash    []byte //交易Hash
	ScriptSig string
}
