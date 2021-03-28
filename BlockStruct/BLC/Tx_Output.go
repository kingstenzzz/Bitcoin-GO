package BLC

type TxOutput struct {

	//引用上一笔脚底饿输出索引

	Value        int
	ScriptPubkey string
}

func (txOutout *TxOutput) CheckPubkeyWithAdress(adress string) bool {
	return adress == txOutout.ScriptPubkey

}
