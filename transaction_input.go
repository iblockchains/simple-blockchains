package main

// TXInput 交易输入
type TXInput struct {
	Txid      []byte
	Vout      int
	ScriptSig string
}
