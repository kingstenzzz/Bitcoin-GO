package BLC

type UTXO struct {
	//UTXO
	TxHash []byte
	//UTXO在其所属交易得输出列表中的索引
	Index int
	//Output
	Output *TxOutput
}
