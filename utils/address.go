package utils

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"regexp"
	"strconv"
	"strings"
)

func CheckETHAddress(addr string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(addr)
}

func GetMessage(msg string) common.Hash {
	data := []byte("\x19Ethereum Signed Message:\n" + strconv.Itoa(len(msg)) + msg)
	return crypto.Keccak256Hash(data)
}
func VerifyMessage(message string, signedMessage string) (string, error) {
	hash := GetMessage(message)
	decodedMessage := hexutil.MustDecode(signedMessage)
	// Handles cases where EIP-115 is not implemented (most wallets don't implement it)
	if decodedMessage[64] == 27 || decodedMessage[64] == 28 {
		decodedMessage[64] -= 27
	}
	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), decodedMessage)
	if sigPublicKeyECDSA == nil {
		err = errors.New("could not get a public get from the message signature")
	}
	if err != nil {
		return "", err
	}

	return crypto.PubkeyToAddress(*sigPublicKeyECDSA).String(), nil
}

func CheckSign(message string, signedMessage string, addr string) (bool, error) {
	address, err := VerifyMessage(message, signedMessage)
	if err != nil {
		return false, err
	}
	return strings.EqualFold(address, addr), nil
}
