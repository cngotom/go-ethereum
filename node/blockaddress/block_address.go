package blockaddress

import "sync"
import "github.com/ethereum/go-ethereum/common"

var blockedAddress map[common.Address]bool = make(map[common.Address]bool)
var _bdlock sync.RWMutex

func init() {
	//	blockedAddress[common.HexToAddress("")] = true
}

func BlockAddress(address string) {
	_bdlock.Lock()
	defer _bdlock.Unlock()
	blockedAddress[common.HexToAddress(address)] = true
}

func UnBlockAddress(address string) {
	_bdlock.Lock()
	defer _bdlock.Unlock()
	delete(blockedAddress, common.HexToAddress(address))
}

func ListBlockAddress() []string {
	_bdlock.RLock()
	defer _bdlock.RUnlock()
	res := []string{}
	for addr, _ := range blockedAddress {
		res = append(res, addr.Hex())
	}
	return res
}

func InBlocked(address common.Address) bool {
	_bdlock.RLock()
	defer _bdlock.RUnlock()
	_, res := blockedAddress[address]
	return res
}
