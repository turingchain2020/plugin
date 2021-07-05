package turingchain

import (
	turingchainCommon "github.com/turingchain2020/turingchain/common"
	"github.com/ethereum/go-ethereum/crypto"

	//dbm "github.com/turingchain2020/turingchain/common/db"
	turingchainTypes "github.com/turingchain2020/turingchain/types"
	wcom "github.com/turingchain2020/turingchain/wallet/common"
	x2ethTypes "github.com/turingchain2020/plugin/plugin/dapp/x2ethereum/ebrelayer/types"
)

var (
	turingchainAccountKey = []byte("TuringchainAccount4Relayer")
	start             = int(1)
)

//GetAccount ...
func (turingchainRelayer *Relayer4Turingchain) GetAccount(passphrase string) (privateKey, addr string, err error) {
	accountInfo, err := turingchainRelayer.db.Get(turingchainAccountKey)
	if nil != err {
		return "", "", err
	}
	ethAccount := &x2ethTypes.Account4Relayer{}
	if err := turingchainTypes.Decode(accountInfo, ethAccount); nil != err {
		return "", "", err
	}
	decryptered := wcom.CBCDecrypterPrivkey([]byte(passphrase), ethAccount.Privkey)
	privateKey = turingchainCommon.ToHex(decryptered)
	addr = ethAccount.Addr
	return
}

//GetAccountAddr ...
func (turingchainRelayer *Relayer4Turingchain) GetAccountAddr() (addr string, err error) {
	accountInfo, err := turingchainRelayer.db.Get(turingchainAccountKey)
	if nil != err {
		relayerLog.Info("GetValidatorAddr", "Failed to get account from db due to:", err.Error())
		return "", err
	}
	ethAccount := &x2ethTypes.Account4Relayer{}
	if err := turingchainTypes.Decode(accountInfo, ethAccount); nil != err {
		relayerLog.Info("GetValidatorAddr", "Failed to decode due to:", err.Error())
		return "", err
	}
	addr = ethAccount.Addr
	return
}

//ImportPrivateKey ...
func (turingchainRelayer *Relayer4Turingchain) ImportPrivateKey(passphrase, privateKeyStr string) (addr string, err error) {
	privateKeySlice, err := turingchainCommon.FromHex(privateKeyStr)
	if nil != err {
		return "", err
	}
	privateKey, err := crypto.ToECDSA(privateKeySlice)
	if nil != err {
		return "", err
	}

	ethSender := crypto.PubkeyToAddress(privateKey.PublicKey)
	turingchainRelayer.privateKey4Ethereum = privateKey
	turingchainRelayer.ethSender = ethSender
	turingchainRelayer.unlock <- start

	addr = turingchainCommon.ToHex(ethSender.Bytes())
	encryptered := wcom.CBCEncrypterPrivkey([]byte(passphrase), privateKeySlice)
	ethAccount := &x2ethTypes.Account4Relayer{
		Privkey: encryptered,
		Addr:    addr,
	}
	encodedInfo := turingchainTypes.Encode(ethAccount)
	err = turingchainRelayer.db.SetSync(turingchainAccountKey, encodedInfo)

	return
}

//StoreAccountWithNewPassphase ...
func (turingchainRelayer *Relayer4Turingchain) StoreAccountWithNewPassphase(newPassphrase, oldPassphrase string) error {
	accountInfo, err := turingchainRelayer.db.Get(turingchainAccountKey)
	if nil != err {
		relayerLog.Info("StoreAccountWithNewPassphase", "pls check account is created already, err", err)
		return err
	}
	ethAccount := &x2ethTypes.Account4Relayer{}
	if err := turingchainTypes.Decode(accountInfo, ethAccount); nil != err {
		return err
	}
	decryptered := wcom.CBCDecrypterPrivkey([]byte(oldPassphrase), ethAccount.Privkey)
	encryptered := wcom.CBCEncrypterPrivkey([]byte(newPassphrase), decryptered)
	ethAccount.Privkey = encryptered
	encodedInfo := turingchainTypes.Encode(ethAccount)
	return turingchainRelayer.db.SetSync(turingchainAccountKey, encodedInfo)
}

//RestorePrivateKeys ...
func (turingchainRelayer *Relayer4Turingchain) RestorePrivateKeys(passphrase string) error {
	accountInfo, err := turingchainRelayer.db.Get(turingchainAccountKey)
	if nil != err {
		relayerLog.Info("No private key saved for Relayer4Turingchain")
		return nil
	}
	ethAccount := &x2ethTypes.Account4Relayer{}
	if err := turingchainTypes.Decode(accountInfo, ethAccount); nil != err {
		relayerLog.Info("RestorePrivateKeys", "Failed to decode due to:", err.Error())
		return err
	}
	decryptered := wcom.CBCDecrypterPrivkey([]byte(passphrase), ethAccount.Privkey)
	privateKey, err := crypto.ToECDSA(decryptered)
	if nil != err {
		relayerLog.Info("RestorePrivateKeys", "Failed to ToECDSA:", err.Error())
		return err
	}

	turingchainRelayer.rwLock.Lock()
	turingchainRelayer.privateKey4Ethereum = privateKey
	turingchainRelayer.ethSender = crypto.PubkeyToAddress(privateKey.PublicKey)
	turingchainRelayer.rwLock.Unlock()
	turingchainRelayer.unlock <- start
	return nil
}

//func (turingchainRelayer *Relayer4Turingchain) UpdatePrivateKey(Passphrase, privateKey string) error {
//	return nil
//}
