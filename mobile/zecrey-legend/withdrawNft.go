package zecrey_legend

import (
	"encoding/json"
	curve "github.com/zecrey-labs/zecrey-crypto/ecc/ztwistededwards/tebn254"
	"github.com/zecrey-labs/zecrey-crypto/wasm/zecrey-legend/legendTxTypes"
	"log"
)

func ProveWithdrawNft(seed string, segmentInfo string) (txInfo string, err error) {
	// parse segmentInfo
	sk, err := curve.GenerateEddsaPrivateKey(seed)
	if err != nil {
		return "", err
	}
	oTxInfo, err := legendTxTypes.ConstructWithdrawNftTxInfo(sk, segmentInfo)
	if err != nil {
		return "", err
	}
	txInfoBytes, err := json.Marshal(oTxInfo)
	if err != nil {
		log.Println("unable to marshal:", err)
		return "", err
	}
	return string(txInfoBytes), nil
}
