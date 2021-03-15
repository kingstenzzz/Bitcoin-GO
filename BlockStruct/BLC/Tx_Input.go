package BLC

type TxInput struct {

	//引用上一笔脚底饿输出索引
	vout      int
	TxHash    []byte //交易Hash
	ScriptSig string
}
