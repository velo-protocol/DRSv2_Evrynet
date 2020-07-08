// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vabi

import (
	"math/big"
	"strings"

	evrynet "github.com/Evrynetlabs/evrynet-node"
	"github.com/Evrynetlabs/evrynet-node/accounts/abi"
	"github.com/Evrynetlabs/evrynet-node/accounts/abi/bind"
	"github.com/Evrynetlabs/evrynet-node/common"
	"github.com/Evrynetlabs/evrynet-node/core/types"
	"github.com/Evrynetlabs/evrynet-node/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = evrynet.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FeederABI is the input ABI used to generate the binding from.
const FeederABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pf1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pf2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pf3\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_fiatCode\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_collateralCode\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"timeInSecond\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AcceptPrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"timeInSecond\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"CommitPrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"SetValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"updater\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPriceFeed\",\"type\":\"address\"}],\"name\":\"UpdatePriceFeed\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"active\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"collateralCode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fiatCode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"firstPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeInSecond\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastOperationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeInSecond\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numOfPrices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operationCoolDown\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"priceFeed1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"priceFeed2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"priceFeed3\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"priceFeedTimeTol\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"priceTolInBP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"priceUpdateCoolDown\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"secondPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeInSecond\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"started\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"valueTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"}],\"name\":\"startOracle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"}],\"name\":\"commitPrice\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeder\",\"type\":\"address\"}],\"name\":\"updatePriceFeed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setValue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FeederBin is the compiled bytecode used for deploying new contracts.
const FeederBin = `0x60806040526101f4600955603c600a819055600b556000600c819055600e55600f805460ff1916905534801561003457600080fd5b50604051610e2a380380610e2a833981810160405260c081101561005757600080fd5b508051602082015160408301516060840151608085015160a090950151600f8054610100600160a81b0319166101006001600160a01b038089169190910291909117909155601080546001600160a01b031990811683881617909155601180548216838716179055601280549091169184169190911790556015869055601681905593949293919290916100e961010f565b600755505060006006555050600880546001600160a01b03191633179055506101139050565b4290565b610d08806101226000396000f3fe608060405234801561001057600080fd5b506004361061014d5760003560e01c806364bc5204116100c3578063acfc827e1161007c578063acfc827e14610288578063d1e61e24146102b4578063d8cf24fd146102bc578063dc9eaf38146102dd578063dcfe4c53146102e5578063e5693f41146102ed5761014d565b806364bc52041461022857806377680bb5146102305780637b8d56e3146102385780638da5cb5b1461025b578063a99ae65114610263578063ab0ca0e1146102805761014d565b806339a3da4a1161011557806339a3da4a146101dc578063427cb6fe146101e45780635587c436146102085780635623732e1461021057806357ef7f571461021857806361efd490146102205761014d565b8063028a82f41461015257806302fb0c5e14610183578063053f14da1461018b57806316d43a97146101ba5780631f2698ab146101d4575b600080fd5b61016f6004803603602081101561016857600080fd5b50356102f5565b604080519115158252519081900360200190f35b61016f610748565b610193610751565b6040805193845260208401929092526001600160a01b031682820152519081900360600190f35b6101c2610766565b60408051918252519081900360200190f35b61016f61076c565b610193610775565b6101ec61078a565b604080516001600160a01b039092168252519081900360200190f35b6101c2610799565b6101c261079f565b6101c26107a5565b6101c26107ab565b6101c26107b1565b6101c26107b7565b61016f6004803603604081101561024e57600080fd5b50803590602001356107bd565b6101ec610892565b61016f6004803603602081101561027957600080fd5b50356108a1565b6101ec6109b8565b61016f6004803603604081101561029e57600080fd5b50803590602001356001600160a01b03166109cc565b610193610ae6565b6102c4610afb565b6040805192835260208301919091528051918290030190f35b6101c2610b05565b6101c2610b0b565b6101ec610b11565b600f5460009061010090046001600160a01b031633148061032057506010546001600160a01b031633145b8061033557506011546001600160a01b031633145b61033e57600080fd5b6000610348610b20565b600f5490915060ff1680156103715750600b5460075461036d9163ffffffff610b2416565b8110155b61037a57600080fd5b6000600c546000141561045d5760065461039b90859063ffffffff610b3a16565b600954600654919250906103c7906103bb8461271063ffffffff610b5c16565b9063ffffffff610b7c16565b116103dc576103d7848333610b91565b610458565b6040805160608101825285815260208082018590523391830182905260008781556001869055600280546001600160a01b031916841790558351928352908201528151849287927fe95a7bb58ff5f75fa581ca00ff98b7a825caba174942abffb75a1d2b12d90dd9929081900390910190a3600c805460010190555b61073c565b600c546001141561060d57600b5460015461047d9163ffffffff610b2416565b8211156104c2576002546001600160a01b03163314156104a7576104a2848333610b91565b6103d7565b6000546002546103d7919084906001600160a01b0316610b91565b6002546001600160a01b03163314156104da57600080fd5b600a5460015483916104f2919063ffffffff610b2416565b10806105135750600a546001548391610511919063ffffffff610bfa16565b115b15610535576000546001546002546103d79291906001600160a01b0316610b91565b60005461054990859063ffffffff610b3a16565b60095460005491925090610569906103bb8461271063ffffffff610b5c16565b1161058b576000546001546002546103d79291906001600160a01b0316610b91565b6040805160608101825285815260208082018590523391830182905260038790556004859055600580546001600160a01b0319168317905582519182526001908201528151849287927fe95a7bb58ff5f75fa581ca00ff98b7a825caba174942abffb75a1d2b12d90dd9929081900390910190a3600c8054600101905561073c565b600c546002141561073157600b5460015401821115610674576002546001600160a01b031633148061064957506005546001600160a01b031633145b15610659576104a2848333610b91565b6003546005546103d7919084906001600160a01b0316610b91565b6002546001600160a01b0316331480159061069a57506005546001600160a01b03163314155b6106a357600080fd5b600a5460015460009184916106bd9163ffffffff610b2416565b10806106de5750600a5460015484916106dc919063ffffffff610bfa16565b115b156106ec5750600054610711565b6003548514156106fd575083610711565b60005460035461070e919087610c0c565b90505b60015460025461072b9183916001600160a01b0316610b91565b5061073c565b600092505050610743565b6001925050505b919050565b60145460ff1681565b6006546007546008546001600160a01b031683565b600d5481565b600f5460ff1681565b6000546001546002546001600160a01b031683565b6011546001600160a01b031681565b60165481565b600e5481565b600a5481565b60095481565b60155481565b600c5481565b6012546000906001600160a01b031633146108095760405162461bcd60e51b815260040180806020018281038252603a815260200180610c9a603a913960400191505060405180910390fd5b60008361081e57506009805490839055610848565b83600114156108355750600a805490839055610848565b836002141561014d5750600b8054908390555b604080518581526020810183905280820185905290517fabd08d77cf1eed600c8ba851f4210365f6695aa58b9500aa52a83db7d8b534ba9181900360600190a15060019392505050565b6012546001600160a01b031681565b600f5460009061010090046001600160a01b03163314806108cc57506010546001600160a01b031633145b806108e157506011546001600160a01b031633145b6108ea57600080fd5b600f5460ff16156108fa57600080fd5b6000821161093f576040805162461bcd60e51b815260206004820152600d60248201526c07072696365206e656564203e3609c1b604482015290519081900360640190fd5b6000610949610b20565b60078190556006849055600880546001600160a01b03191633908117909155600f805460ff191660011790556040805191825251919250829185917f60a26741549361639cff0096d0c4599ee53831ec6dec262b89d7fb1ddde1726e919081900360200190a350600192915050565b600f5461010090046001600160a01b031681565b6012546000906001600160a01b03163314610a185760405162461bcd60e51b815260040180806020018281038252603a815260200180610c9a603a913960400191505060405180910390fd5b60038310610a2557600080fd5b338284610a5157600f8054610100600160a81b0319166101006001600160a01b03841602179055610a96565b8460011415610a7a57601080546001600160a01b0319166001600160a01b038316179055610a96565b601180546001600160a01b0319166001600160a01b0383161790555b604080516001600160a01b0380851682528316602082015281517f0c19c511fd3cdf7b4524c00cd2176627ee9fdd2644d68683e6cc04b79280f316929181900390910190a1506001949350505050565b6003546004546005546001600160a01b031683565b6006546007549091565b60135481565b600b5481565b6010546001600160a01b031681565b4290565b600082820183811015610b3357fe5b9392505050565b6000818311610b5257610b4d8284610bfa565b610b33565b610b338383610bfa565b6000828202831580610b76575082848281610b7357fe5b04145b610b3357fe5b600080828481610b8857fe5b04949350505050565b60068390556007829055600880546001600160a01b0383166001600160a01b031990911681179091556000600c556040805191825251839185917f60a26741549361639cff0096d0c4599ee53831ec6dec262b89d7fb1ddde1726e9181900360200190a3505050565b600082821115610c0657fe5b50900390565b6000610c1e828563ffffffff610c8116565b610c2e858563ffffffff610c8116565b186001600160f81b031916610c44575082610b33565b610c54828463ffffffff610c8116565b610c64848663ffffffff610c8116565b186001600160f81b031916610c7a575081610b33565b5092915050565b60008082841115610b335750600160f81b939250505056fe4665656465722e6f6e6c794f776e65723a2063616c6c6572206d75737420626520616e206f776e6572206f66207468697320636f6e7472616374a265627a7a72315820abea8a2fd94ede33a5dc732e2cf3a4a092beec46401660e87eec010f42161c4964736f6c634300050f0032`

// DeployFeeder deploys a new Evrynet contract, binding an instance of Feeder to it.
func DeployFeeder(auth *bind.TransactOpts, backend bind.ContractBackend, pf1 common.Address, pf2 common.Address, pf3 common.Address, _owner common.Address, _fiatCode [32]byte, _collateralCode [32]byte) (common.Address, *types.Transaction, *Feeder, error) {
	parsed, err := abi.JSON(strings.NewReader(FeederABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FeederBin), backend, pf1, pf2, pf3, _owner, _fiatCode, _collateralCode)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Feeder{FeederCaller: FeederCaller{contract: contract}, FeederTransactor: FeederTransactor{contract: contract}, FeederFilterer: FeederFilterer{contract: contract}}, nil
}

// Feeder is an auto generated Go binding around an Evrynet contract.
type Feeder struct {
	FeederCaller     // Read-only binding to the contract
	FeederTransactor // Write-only binding to the contract
	FeederFilterer   // Log filterer for contract events
}

// FeederCaller is an auto generated read-only Go binding around an Evrynet contract.
type FeederCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeederTransactor is an auto generated write-only Go binding around an Evrynet contract.
type FeederTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeederFilterer is an auto generated log filtering Go binding around an Evrynet contract events.
type FeederFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeederSession is an auto generated Go binding around an Evrynet contract,
// with pre-set call and transact options.
type FeederSession struct {
	Contract     *Feeder           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeederCallerSession is an auto generated read-only Go binding around an Evrynet contract,
// with pre-set call options.
type FeederCallerSession struct {
	Contract *FeederCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FeederTransactorSession is an auto generated write-only Go binding around an Evrynet contract,
// with pre-set transact options.
type FeederTransactorSession struct {
	Contract     *FeederTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeederRaw is an auto generated low-level Go binding around an Evrynet contract.
type FeederRaw struct {
	Contract *Feeder // Generic contract binding to access the raw methods on
}

// FeederCallerRaw is an auto generated low-level read-only Go binding around an Evrynet contract.
type FeederCallerRaw struct {
	Contract *FeederCaller // Generic read-only contract binding to access the raw methods on
}

// FeederTransactorRaw is an auto generated low-level write-only Go binding around an Evrynet contract.
type FeederTransactorRaw struct {
	Contract *FeederTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeeder creates a new instance of Feeder, bound to a specific deployed contract.
func NewFeeder(address common.Address, backend bind.ContractBackend) (*Feeder, error) {
	contract, err := bindFeeder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Feeder{FeederCaller: FeederCaller{contract: contract}, FeederTransactor: FeederTransactor{contract: contract}, FeederFilterer: FeederFilterer{contract: contract}}, nil
}

// NewFeederCaller creates a new read-only instance of Feeder, bound to a specific deployed contract.
func NewFeederCaller(address common.Address, caller bind.ContractCaller) (*FeederCaller, error) {
	contract, err := bindFeeder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeederCaller{contract: contract}, nil
}

// NewFeederTransactor creates a new write-only instance of Feeder, bound to a specific deployed contract.
func NewFeederTransactor(address common.Address, transactor bind.ContractTransactor) (*FeederTransactor, error) {
	contract, err := bindFeeder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeederTransactor{contract: contract}, nil
}

// NewFeederFilterer creates a new log filterer instance of Feeder, bound to a specific deployed contract.
func NewFeederFilterer(address common.Address, filterer bind.ContractFilterer) (*FeederFilterer, error) {
	contract, err := bindFeeder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeederFilterer{contract: contract}, nil
}

// bindFeeder binds a generic wrapper to an already deployed contract.
func bindFeeder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FeederABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Feeder *FeederRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Feeder.Contract.FeederCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Feeder *FeederRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Feeder.Contract.FeederTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Feeder *FeederRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Feeder.Contract.FeederTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Feeder *FeederCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Feeder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Feeder *FeederTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Feeder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Feeder *FeederTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Feeder.Contract.contract.Transact(opts, method, params...)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() constant returns(bool)
func (_Feeder *FeederCaller) Active(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Feeder.contract.Call(opts, out, "active")
	return *ret0, err
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() constant returns(bool)
func (_Feeder *FeederSession) Active() (bool, error) {
	return _Feeder.Contract.Active(&_Feeder.CallOpts)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() constant returns(bool)
func (_Feeder *FeederCallerSession) Active() (bool, error) {
	return _Feeder.Contract.Active(&_Feeder.CallOpts)
}

// CollateralCode is a free data retrieval call binding the contract method 0x5587c436.
//
// Solidity: function collateralCode() constant returns(bytes32)
func (_Feeder *FeederCaller) CollateralCode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Feeder.contract.Call(opts, out, "collateralCode")
	return *ret0, err
}

// CollateralCode is a free data retrieval call binding the contract method 0x5587c436.
//
// Solidity: function collateralCode() constant returns(bytes32)
func (_Feeder *FeederSession) CollateralCode() ([32]byte, error) {
	return _Feeder.Contract.CollateralCode(&_Feeder.CallOpts)
}

// CollateralCode is a free data retrieval call binding the contract method 0x5587c436.
//
// Solidity: function collateralCode() constant returns(bytes32)
func (_Feeder *FeederCallerSession) CollateralCode() ([32]byte, error) {
	return _Feeder.Contract.CollateralCode(&_Feeder.CallOpts)
}

// FiatCode is a free data retrieval call binding the contract method 0x64bc5204.
//
// Solidity: function fiatCode() constant returns(bytes32)
func (_Feeder *FeederCaller) FiatCode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Feeder.contract.Call(opts, out, "fiatCode")
	return *ret0, err
}

// FiatCode is a free data retrieval call binding the contract method 0x64bc5204.
//
// Solidity: function fiatCode() constant returns(bytes32)
func (_Feeder *FeederSession) FiatCode() ([32]byte, error) {
	return _Feeder.Contract.FiatCode(&_Feeder.CallOpts)
}

// FiatCode is a free data retrieval call binding the contract method 0x64bc5204.
//
// Solidity: function fiatCode() constant returns(bytes32)
func (_Feeder *FeederCallerSession) FiatCode() ([32]byte, error) {
	return _Feeder.Contract.FiatCode(&_Feeder.CallOpts)
}

// FirstPrice is a free data retrieval call binding the contract method 0x39a3da4a.
//
// Solidity: function firstPrice() constant returns(uint256 priceInWei, uint256 timeInSecond, address source)
func (_Feeder *FeederCaller) FirstPrice(opts *bind.CallOpts) (struct {
	PriceInWei   *big.Int
	TimeInSecond *big.Int
	Source       common.Address
}, error) {
	ret := new(struct {
		PriceInWei   *big.Int
		TimeInSecond *big.Int
		Source       common.Address
	})
	out := ret
	err := _Feeder.contract.Call(opts, out, "firstPrice")
	return *ret, err
}

// FirstPrice is a free data retrieval call binding the contract method 0x39a3da4a.
//
// Solidity: function firstPrice() constant returns(uint256 priceInWei, uint256 timeInSecond, address source)
func (_Feeder *FeederSession) FirstPrice() (struct {
	PriceInWei   *big.Int
	TimeInSecond *big.Int
	Source       common.Address
}, error) {
	return _Feeder.Contract.FirstPrice(&_Feeder.CallOpts)
}

// FirstPrice is a free data retrieval call binding the contract method 0x39a3da4a.
//
// Solidity: function firstPrice() constant returns(uint256 priceInWei, uint256 timeInSecond, address source)
func (_Feeder *FeederCallerSession) FirstPrice() (struct {
	PriceInWei   *big.Int
	TimeInSecond *big.Int
	Source       common.Address
}, error) {
	return _Feeder.Contract.FirstPrice(&_Feeder.CallOpts)
}

// GetLastPrice is a free data retrieval call binding the contract method 0xd8cf24fd.
//
// Solidity: function getLastPrice() constant returns(uint256, uint256)
func (_Feeder *FeederCaller) GetLastPrice(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Feeder.contract.Call(opts, out, "getLastPrice")
	return *ret0, *ret1, err
}

// GetLastPrice is a free data retrieval call binding the contract method 0xd8cf24fd.
//
// Solidity: function getLastPrice() constant returns(uint256, uint256)
func (_Feeder *FeederSession) GetLastPrice() (*big.Int, *big.Int, error) {
	return _Feeder.Contract.GetLastPrice(&_Feeder.CallOpts)
}

// GetLastPrice is a free data retrieval call binding the contract method 0xd8cf24fd.
//
// Solidity: function getLastPrice() constant returns(uint256, uint256)
func (_Feeder *FeederCallerSession) GetLastPrice() (*big.Int, *big.Int, error) {
	return _Feeder.Contract.GetLastPrice(&_Feeder.CallOpts)
}

// LastOperationTime is a free data retrieval call binding the contract method 0x16d43a97.
//
// Solidity: function lastOperationTime() constant returns(uint256)
func (_Feeder *FeederCaller) LastOperationTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Feeder.contract.Call(opts, out, "lastOperationTime")
	return *ret0, err
}

// LastOperationTime is a free data retrieval call binding the contract method 0x16d43a97.
//
// Solidity: function lastOperationTime() constant returns(uint256)
func (_Feeder *FeederSession) LastOperationTime() (*big.Int, error) {
	return _Feeder.Contract.LastOperationTime(&_Feeder.CallOpts)
}

// LastOperationTime is a free data retrieval call binding the contract method 0x16d43a97.
//
// Solidity: function lastOperationTime() constant returns(uint256)
func (_Feeder *FeederCallerSession) LastOperationTime() (*big.Int, error) {
	return _Feeder.Contract.LastOperationTime(&_Feeder.CallOpts)
}

// LastPrice is a free data retrieval call binding the contract method 0x053f14da.
//
// Solidity: function lastPrice() constant returns(uint256 priceInWei, uint256 timeInSecond, address source)
func (_Feeder *FeederCaller) LastPrice(opts *bind.CallOpts) (struct {
	PriceInWei   *big.Int
	TimeInSecond *big.Int
	Source       common.Address
}, error) {
	ret := new(struct {
		PriceInWei   *big.Int
		TimeInSecond *big.Int
		Source       common.Address
	})
	out := ret
	err := _Feeder.contract.Call(opts, out, "lastPrice")
	return *ret, err
}

// LastPrice is a free data retrieval call binding the contract method 0x053f14da.
//
// Solidity: function lastPrice() constant returns(uint256 priceInWei, uint256 timeInSecond, address source)
func (_Feeder *FeederSession) LastPrice() (struct {
	PriceInWei   *big.Int
	TimeInSecond *big.Int
	Source       common.Address
}, error) {
	return _Feeder.Contract.LastPrice(&_Feeder.CallOpts)
}

// LastPrice is a free data retrieval call binding the contract method 0x053f14da.
//
// Solidity: function lastPrice() constant returns(uint256 priceInWei, uint256 timeInSecond, address source)
func (_Feeder *FeederCallerSession) LastPrice() (struct {
	PriceInWei   *big.Int
	TimeInSecond *big.Int
	Source       common.Address
}, error) {
	return _Feeder.Contract.LastPrice(&_Feeder.CallOpts)
}

// NumOfPrices is a free data retrieval call binding the contract method 0x77680bb5.
//
// Solidity: function numOfPrices() constant returns(uint256)
func (_Feeder *FeederCaller) NumOfPrices(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Feeder.contract.Call(opts, out, "numOfPrices")
	return *ret0, err
}

// NumOfPrices is a free data retrieval call binding the contract method 0x77680bb5.
//
// Solidity: function numOfPrices() constant returns(uint256)
func (_Feeder *FeederSession) NumOfPrices() (*big.Int, error) {
	return _Feeder.Contract.NumOfPrices(&_Feeder.CallOpts)
}

// NumOfPrices is a free data retrieval call binding the contract method 0x77680bb5.
//
// Solidity: function numOfPrices() constant returns(uint256)
func (_Feeder *FeederCallerSession) NumOfPrices() (*big.Int, error) {
	return _Feeder.Contract.NumOfPrices(&_Feeder.CallOpts)
}

// OperationCoolDown is a free data retrieval call binding the contract method 0x5623732e.
//
// Solidity: function operationCoolDown() constant returns(uint256)
func (_Feeder *FeederCaller) OperationCoolDown(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Feeder.contract.Call(opts, out, "operationCoolDown")
	return *ret0, err
}

// OperationCoolDown is a free data retrieval call binding the contract method 0x5623732e.
//
// Solidity: function operationCoolDown() constant returns(uint256)
func (_Feeder *FeederSession) OperationCoolDown() (*big.Int, error) {
	return _Feeder.Contract.OperationCoolDown(&_Feeder.CallOpts)
}

// OperationCoolDown is a free data retrieval call binding the contract method 0x5623732e.
//
// Solidity: function operationCoolDown() constant returns(uint256)
func (_Feeder *FeederCallerSession) OperationCoolDown() (*big.Int, error) {
	return _Feeder.Contract.OperationCoolDown(&_Feeder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Feeder *FeederCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Feeder.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Feeder *FeederSession) Owner() (common.Address, error) {
	return _Feeder.Contract.Owner(&_Feeder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Feeder *FeederCallerSession) Owner() (common.Address, error) {
	return _Feeder.Contract.Owner(&_Feeder.CallOpts)
}

// PriceFeed1 is a free data retrieval call binding the contract method 0xab0ca0e1.
//
// Solidity: function priceFeed1() constant returns(address)
func (_Feeder *FeederCaller) PriceFeed1(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Feeder.contract.Call(opts, out, "priceFeed1")
	return *ret0, err
}

// PriceFeed1 is a free data retrieval call binding the contract method 0xab0ca0e1.
//
// Solidity: function priceFeed1() constant returns(address)
func (_Feeder *FeederSession) PriceFeed1() (common.Address, error) {
	return _Feeder.Contract.PriceFeed1(&_Feeder.CallOpts)
}

// PriceFeed1 is a free data retrieval call binding the contract method 0xab0ca0e1.
//
// Solidity: function priceFeed1() constant returns(address)
func (_Feeder *FeederCallerSession) PriceFeed1() (common.Address, error) {
	return _Feeder.Contract.PriceFeed1(&_Feeder.CallOpts)
}

// PriceFeed2 is a free data retrieval call binding the contract method 0xe5693f41.
//
// Solidity: function priceFeed2() constant returns(address)
func (_Feeder *FeederCaller) PriceFeed2(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Feeder.contract.Call(opts, out, "priceFeed2")
	return *ret0, err
}

// PriceFeed2 is a free data retrieval call binding the contract method 0xe5693f41.
//
// Solidity: function priceFeed2() constant returns(address)
func (_Feeder *FeederSession) PriceFeed2() (common.Address, error) {
	return _Feeder.Contract.PriceFeed2(&_Feeder.CallOpts)
}

// PriceFeed2 is a free data retrieval call binding the contract method 0xe5693f41.
//
// Solidity: function priceFeed2() constant returns(address)
func (_Feeder *FeederCallerSession) PriceFeed2() (common.Address, error) {
	return _Feeder.Contract.PriceFeed2(&_Feeder.CallOpts)
}

// PriceFeed3 is a free data retrieval call binding the contract method 0x427cb6fe.
//
// Solidity: function priceFeed3() constant returns(address)
func (_Feeder *FeederCaller) PriceFeed3(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Feeder.contract.Call(opts, out, "priceFeed3")
	return *ret0, err
}

// PriceFeed3 is a free data retrieval call binding the contract method 0x427cb6fe.
//
// Solidity: function priceFeed3() constant returns(address)
func (_Feeder *FeederSession) PriceFeed3() (common.Address, error) {
	return _Feeder.Contract.PriceFeed3(&_Feeder.CallOpts)
}

// PriceFeed3 is a free data retrieval call binding the contract method 0x427cb6fe.
//
// Solidity: function priceFeed3() constant returns(address)
func (_Feeder *FeederCallerSession) PriceFeed3() (common.Address, error) {
	return _Feeder.Contract.PriceFeed3(&_Feeder.CallOpts)
}

// PriceFeedTimeTol is a free data retrieval call binding the contract method 0x57ef7f57.
//
// Solidity: function priceFeedTimeTol() constant returns(uint256)
func (_Feeder *FeederCaller) PriceFeedTimeTol(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Feeder.contract.Call(opts, out, "priceFeedTimeTol")
	return *ret0, err
}

// PriceFeedTimeTol is a free data retrieval call binding the contract method 0x57ef7f57.
//
// Solidity: function priceFeedTimeTol() constant returns(uint256)
func (_Feeder *FeederSession) PriceFeedTimeTol() (*big.Int, error) {
	return _Feeder.Contract.PriceFeedTimeTol(&_Feeder.CallOpts)
}

// PriceFeedTimeTol is a free data retrieval call binding the contract method 0x57ef7f57.
//
// Solidity: function priceFeedTimeTol() constant returns(uint256)
func (_Feeder *FeederCallerSession) PriceFeedTimeTol() (*big.Int, error) {
	return _Feeder.Contract.PriceFeedTimeTol(&_Feeder.CallOpts)
}

// PriceTolInBP is a free data retrieval call binding the contract method 0x61efd490.
//
// Solidity: function priceTolInBP() constant returns(uint256)
func (_Feeder *FeederCaller) PriceTolInBP(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Feeder.contract.Call(opts, out, "priceTolInBP")
	return *ret0, err
}

// PriceTolInBP is a free data retrieval call binding the contract method 0x61efd490.
//
// Solidity: function priceTolInBP() constant returns(uint256)
func (_Feeder *FeederSession) PriceTolInBP() (*big.Int, error) {
	return _Feeder.Contract.PriceTolInBP(&_Feeder.CallOpts)
}

// PriceTolInBP is a free data retrieval call binding the contract method 0x61efd490.
//
// Solidity: function priceTolInBP() constant returns(uint256)
func (_Feeder *FeederCallerSession) PriceTolInBP() (*big.Int, error) {
	return _Feeder.Contract.PriceTolInBP(&_Feeder.CallOpts)
}

// PriceUpdateCoolDown is a free data retrieval call binding the contract method 0xdcfe4c53.
//
// Solidity: function priceUpdateCoolDown() constant returns(uint256)
func (_Feeder *FeederCaller) PriceUpdateCoolDown(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Feeder.contract.Call(opts, out, "priceUpdateCoolDown")
	return *ret0, err
}

// PriceUpdateCoolDown is a free data retrieval call binding the contract method 0xdcfe4c53.
//
// Solidity: function priceUpdateCoolDown() constant returns(uint256)
func (_Feeder *FeederSession) PriceUpdateCoolDown() (*big.Int, error) {
	return _Feeder.Contract.PriceUpdateCoolDown(&_Feeder.CallOpts)
}

// PriceUpdateCoolDown is a free data retrieval call binding the contract method 0xdcfe4c53.
//
// Solidity: function priceUpdateCoolDown() constant returns(uint256)
func (_Feeder *FeederCallerSession) PriceUpdateCoolDown() (*big.Int, error) {
	return _Feeder.Contract.PriceUpdateCoolDown(&_Feeder.CallOpts)
}

// SecondPrice is a free data retrieval call binding the contract method 0xd1e61e24.
//
// Solidity: function secondPrice() constant returns(uint256 priceInWei, uint256 timeInSecond, address source)
func (_Feeder *FeederCaller) SecondPrice(opts *bind.CallOpts) (struct {
	PriceInWei   *big.Int
	TimeInSecond *big.Int
	Source       common.Address
}, error) {
	ret := new(struct {
		PriceInWei   *big.Int
		TimeInSecond *big.Int
		Source       common.Address
	})
	out := ret
	err := _Feeder.contract.Call(opts, out, "secondPrice")
	return *ret, err
}

// SecondPrice is a free data retrieval call binding the contract method 0xd1e61e24.
//
// Solidity: function secondPrice() constant returns(uint256 priceInWei, uint256 timeInSecond, address source)
func (_Feeder *FeederSession) SecondPrice() (struct {
	PriceInWei   *big.Int
	TimeInSecond *big.Int
	Source       common.Address
}, error) {
	return _Feeder.Contract.SecondPrice(&_Feeder.CallOpts)
}

// SecondPrice is a free data retrieval call binding the contract method 0xd1e61e24.
//
// Solidity: function secondPrice() constant returns(uint256 priceInWei, uint256 timeInSecond, address source)
func (_Feeder *FeederCallerSession) SecondPrice() (struct {
	PriceInWei   *big.Int
	TimeInSecond *big.Int
	Source       common.Address
}, error) {
	return _Feeder.Contract.SecondPrice(&_Feeder.CallOpts)
}

// Started is a free data retrieval call binding the contract method 0x1f2698ab.
//
// Solidity: function started() constant returns(bool)
func (_Feeder *FeederCaller) Started(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Feeder.contract.Call(opts, out, "started")
	return *ret0, err
}

// Started is a free data retrieval call binding the contract method 0x1f2698ab.
//
// Solidity: function started() constant returns(bool)
func (_Feeder *FeederSession) Started() (bool, error) {
	return _Feeder.Contract.Started(&_Feeder.CallOpts)
}

// Started is a free data retrieval call binding the contract method 0x1f2698ab.
//
// Solidity: function started() constant returns(bool)
func (_Feeder *FeederCallerSession) Started() (bool, error) {
	return _Feeder.Contract.Started(&_Feeder.CallOpts)
}

// ValueTimestamp is a free data retrieval call binding the contract method 0xdc9eaf38.
//
// Solidity: function valueTimestamp() constant returns(uint256)
func (_Feeder *FeederCaller) ValueTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Feeder.contract.Call(opts, out, "valueTimestamp")
	return *ret0, err
}

// ValueTimestamp is a free data retrieval call binding the contract method 0xdc9eaf38.
//
// Solidity: function valueTimestamp() constant returns(uint256)
func (_Feeder *FeederSession) ValueTimestamp() (*big.Int, error) {
	return _Feeder.Contract.ValueTimestamp(&_Feeder.CallOpts)
}

// ValueTimestamp is a free data retrieval call binding the contract method 0xdc9eaf38.
//
// Solidity: function valueTimestamp() constant returns(uint256)
func (_Feeder *FeederCallerSession) ValueTimestamp() (*big.Int, error) {
	return _Feeder.Contract.ValueTimestamp(&_Feeder.CallOpts)
}

// CommitPrice is a paid mutator transaction binding the contract method 0x028a82f4.
//
// Solidity: function commitPrice(uint256 priceInWei) returns(bool success)
func (_Feeder *FeederTransactor) CommitPrice(opts *bind.TransactOpts, priceInWei *big.Int) (*types.Transaction, error) {
	return _Feeder.contract.Transact(opts, "commitPrice", priceInWei)
}

// CommitPrice is a paid mutator transaction binding the contract method 0x028a82f4.
//
// Solidity: function commitPrice(uint256 priceInWei) returns(bool success)
func (_Feeder *FeederSession) CommitPrice(priceInWei *big.Int) (*types.Transaction, error) {
	return _Feeder.Contract.CommitPrice(&_Feeder.TransactOpts, priceInWei)
}

// CommitPrice is a paid mutator transaction binding the contract method 0x028a82f4.
//
// Solidity: function commitPrice(uint256 priceInWei) returns(bool success)
func (_Feeder *FeederTransactorSession) CommitPrice(priceInWei *big.Int) (*types.Transaction, error) {
	return _Feeder.Contract.CommitPrice(&_Feeder.TransactOpts, priceInWei)
}

// SetValue is a paid mutator transaction binding the contract method 0x7b8d56e3.
//
// Solidity: function setValue(uint256 idx, uint256 newValue) returns(bool success)
func (_Feeder *FeederTransactor) SetValue(opts *bind.TransactOpts, idx *big.Int, newValue *big.Int) (*types.Transaction, error) {
	return _Feeder.contract.Transact(opts, "setValue", idx, newValue)
}

// SetValue is a paid mutator transaction binding the contract method 0x7b8d56e3.
//
// Solidity: function setValue(uint256 idx, uint256 newValue) returns(bool success)
func (_Feeder *FeederSession) SetValue(idx *big.Int, newValue *big.Int) (*types.Transaction, error) {
	return _Feeder.Contract.SetValue(&_Feeder.TransactOpts, idx, newValue)
}

// SetValue is a paid mutator transaction binding the contract method 0x7b8d56e3.
//
// Solidity: function setValue(uint256 idx, uint256 newValue) returns(bool success)
func (_Feeder *FeederTransactorSession) SetValue(idx *big.Int, newValue *big.Int) (*types.Transaction, error) {
	return _Feeder.Contract.SetValue(&_Feeder.TransactOpts, idx, newValue)
}

// StartOracle is a paid mutator transaction binding the contract method 0xa99ae651.
//
// Solidity: function startOracle(uint256 priceInWei) returns(bool success)
func (_Feeder *FeederTransactor) StartOracle(opts *bind.TransactOpts, priceInWei *big.Int) (*types.Transaction, error) {
	return _Feeder.contract.Transact(opts, "startOracle", priceInWei)
}

// StartOracle is a paid mutator transaction binding the contract method 0xa99ae651.
//
// Solidity: function startOracle(uint256 priceInWei) returns(bool success)
func (_Feeder *FeederSession) StartOracle(priceInWei *big.Int) (*types.Transaction, error) {
	return _Feeder.Contract.StartOracle(&_Feeder.TransactOpts, priceInWei)
}

// StartOracle is a paid mutator transaction binding the contract method 0xa99ae651.
//
// Solidity: function startOracle(uint256 priceInWei) returns(bool success)
func (_Feeder *FeederTransactorSession) StartOracle(priceInWei *big.Int) (*types.Transaction, error) {
	return _Feeder.Contract.StartOracle(&_Feeder.TransactOpts, priceInWei)
}

// UpdatePriceFeed is a paid mutator transaction binding the contract method 0xacfc827e.
//
// Solidity: function updatePriceFeed(uint256 index, address feeder) returns(bool)
func (_Feeder *FeederTransactor) UpdatePriceFeed(opts *bind.TransactOpts, index *big.Int, feeder common.Address) (*types.Transaction, error) {
	return _Feeder.contract.Transact(opts, "updatePriceFeed", index, feeder)
}

// UpdatePriceFeed is a paid mutator transaction binding the contract method 0xacfc827e.
//
// Solidity: function updatePriceFeed(uint256 index, address feeder) returns(bool)
func (_Feeder *FeederSession) UpdatePriceFeed(index *big.Int, feeder common.Address) (*types.Transaction, error) {
	return _Feeder.Contract.UpdatePriceFeed(&_Feeder.TransactOpts, index, feeder)
}

// UpdatePriceFeed is a paid mutator transaction binding the contract method 0xacfc827e.
//
// Solidity: function updatePriceFeed(uint256 index, address feeder) returns(bool)
func (_Feeder *FeederTransactorSession) UpdatePriceFeed(index *big.Int, feeder common.Address) (*types.Transaction, error) {
	return _Feeder.Contract.UpdatePriceFeed(&_Feeder.TransactOpts, index, feeder)
}

// FeederAcceptPriceIterator is returned from FilterAcceptPrice and is used to iterate over the raw logs and unpacked data for AcceptPrice events raised by the Feeder contract.
type FeederAcceptPriceIterator struct {
	Event *FeederAcceptPrice // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log       // Log channel receiving the found contract events
	sub  evrynet.Subscription // Subscription for errors, completion and termination
	done bool                 // Whether the subscription completed delivering logs
	fail error                // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FeederAcceptPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeederAcceptPrice)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FeederAcceptPrice)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FeederAcceptPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeederAcceptPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeederAcceptPrice represents a AcceptPrice event raised by the Feeder contract.
type FeederAcceptPrice struct {
	PriceInWei   *big.Int
	TimeInSecond *big.Int
	Sender       common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAcceptPrice is a free log retrieval operation binding the contract event 0x60a26741549361639cff0096d0c4599ee53831ec6dec262b89d7fb1ddde1726e.
//
// Solidity: event AcceptPrice(uint256 indexed priceInWei, uint256 indexed timeInSecond, address sender)
func (_Feeder *FeederFilterer) FilterAcceptPrice(opts *bind.FilterOpts, priceInWei []*big.Int, timeInSecond []*big.Int) (*FeederAcceptPriceIterator, error) {

	var priceInWeiRule []interface{}
	for _, priceInWeiItem := range priceInWei {
		priceInWeiRule = append(priceInWeiRule, priceInWeiItem)
	}
	var timeInSecondRule []interface{}
	for _, timeInSecondItem := range timeInSecond {
		timeInSecondRule = append(timeInSecondRule, timeInSecondItem)
	}

	logs, sub, err := _Feeder.contract.FilterLogs(opts, "AcceptPrice", priceInWeiRule, timeInSecondRule)
	if err != nil {
		return nil, err
	}
	return &FeederAcceptPriceIterator{contract: _Feeder.contract, event: "AcceptPrice", logs: logs, sub: sub}, nil
}

// WatchAcceptPrice is a free log subscription operation binding the contract event 0x60a26741549361639cff0096d0c4599ee53831ec6dec262b89d7fb1ddde1726e.
//
// Solidity: event AcceptPrice(uint256 indexed priceInWei, uint256 indexed timeInSecond, address sender)
func (_Feeder *FeederFilterer) WatchAcceptPrice(opts *bind.WatchOpts, sink chan<- *FeederAcceptPrice, priceInWei []*big.Int, timeInSecond []*big.Int) (event.Subscription, error) {

	var priceInWeiRule []interface{}
	for _, priceInWeiItem := range priceInWei {
		priceInWeiRule = append(priceInWeiRule, priceInWeiItem)
	}
	var timeInSecondRule []interface{}
	for _, timeInSecondItem := range timeInSecond {
		timeInSecondRule = append(timeInSecondRule, timeInSecondItem)
	}

	logs, sub, err := _Feeder.contract.WatchLogs(opts, "AcceptPrice", priceInWeiRule, timeInSecondRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeederAcceptPrice)
				if err := _Feeder.contract.UnpackLog(event, "AcceptPrice", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// FeederCommitPriceIterator is returned from FilterCommitPrice and is used to iterate over the raw logs and unpacked data for CommitPrice events raised by the Feeder contract.
type FeederCommitPriceIterator struct {
	Event *FeederCommitPrice // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log       // Log channel receiving the found contract events
	sub  evrynet.Subscription // Subscription for errors, completion and termination
	done bool                 // Whether the subscription completed delivering logs
	fail error                // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FeederCommitPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeederCommitPrice)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FeederCommitPrice)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FeederCommitPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeederCommitPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeederCommitPrice represents a CommitPrice event raised by the Feeder contract.
type FeederCommitPrice struct {
	PriceInWei   *big.Int
	TimeInSecond *big.Int
	Sender       common.Address
	Index        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCommitPrice is a free log retrieval operation binding the contract event 0xe95a7bb58ff5f75fa581ca00ff98b7a825caba174942abffb75a1d2b12d90dd9.
//
// Solidity: event CommitPrice(uint256 indexed priceInWei, uint256 indexed timeInSecond, address sender, uint256 index)
func (_Feeder *FeederFilterer) FilterCommitPrice(opts *bind.FilterOpts, priceInWei []*big.Int, timeInSecond []*big.Int) (*FeederCommitPriceIterator, error) {

	var priceInWeiRule []interface{}
	for _, priceInWeiItem := range priceInWei {
		priceInWeiRule = append(priceInWeiRule, priceInWeiItem)
	}
	var timeInSecondRule []interface{}
	for _, timeInSecondItem := range timeInSecond {
		timeInSecondRule = append(timeInSecondRule, timeInSecondItem)
	}

	logs, sub, err := _Feeder.contract.FilterLogs(opts, "CommitPrice", priceInWeiRule, timeInSecondRule)
	if err != nil {
		return nil, err
	}
	return &FeederCommitPriceIterator{contract: _Feeder.contract, event: "CommitPrice", logs: logs, sub: sub}, nil
}

// WatchCommitPrice is a free log subscription operation binding the contract event 0xe95a7bb58ff5f75fa581ca00ff98b7a825caba174942abffb75a1d2b12d90dd9.
//
// Solidity: event CommitPrice(uint256 indexed priceInWei, uint256 indexed timeInSecond, address sender, uint256 index)
func (_Feeder *FeederFilterer) WatchCommitPrice(opts *bind.WatchOpts, sink chan<- *FeederCommitPrice, priceInWei []*big.Int, timeInSecond []*big.Int) (event.Subscription, error) {

	var priceInWeiRule []interface{}
	for _, priceInWeiItem := range priceInWei {
		priceInWeiRule = append(priceInWeiRule, priceInWeiItem)
	}
	var timeInSecondRule []interface{}
	for _, timeInSecondItem := range timeInSecond {
		timeInSecondRule = append(timeInSecondRule, timeInSecondItem)
	}

	logs, sub, err := _Feeder.contract.WatchLogs(opts, "CommitPrice", priceInWeiRule, timeInSecondRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeederCommitPrice)
				if err := _Feeder.contract.UnpackLog(event, "CommitPrice", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// FeederSetValueIterator is returned from FilterSetValue and is used to iterate over the raw logs and unpacked data for SetValue events raised by the Feeder contract.
type FeederSetValueIterator struct {
	Event *FeederSetValue // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log       // Log channel receiving the found contract events
	sub  evrynet.Subscription // Subscription for errors, completion and termination
	done bool                 // Whether the subscription completed delivering logs
	fail error                // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FeederSetValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeederSetValue)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FeederSetValue)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FeederSetValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeederSetValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeederSetValue represents a SetValue event raised by the Feeder contract.
type FeederSetValue struct {
	Index    *big.Int
	OldValue *big.Int
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetValue is a free log retrieval operation binding the contract event 0xabd08d77cf1eed600c8ba851f4210365f6695aa58b9500aa52a83db7d8b534ba.
//
// Solidity: event SetValue(uint256 index, uint256 oldValue, uint256 newValue)
func (_Feeder *FeederFilterer) FilterSetValue(opts *bind.FilterOpts) (*FeederSetValueIterator, error) {

	logs, sub, err := _Feeder.contract.FilterLogs(opts, "SetValue")
	if err != nil {
		return nil, err
	}
	return &FeederSetValueIterator{contract: _Feeder.contract, event: "SetValue", logs: logs, sub: sub}, nil
}

// WatchSetValue is a free log subscription operation binding the contract event 0xabd08d77cf1eed600c8ba851f4210365f6695aa58b9500aa52a83db7d8b534ba.
//
// Solidity: event SetValue(uint256 index, uint256 oldValue, uint256 newValue)
func (_Feeder *FeederFilterer) WatchSetValue(opts *bind.WatchOpts, sink chan<- *FeederSetValue) (event.Subscription, error) {

	logs, sub, err := _Feeder.contract.WatchLogs(opts, "SetValue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeederSetValue)
				if err := _Feeder.contract.UnpackLog(event, "SetValue", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// FeederUpdatePriceFeedIterator is returned from FilterUpdatePriceFeed and is used to iterate over the raw logs and unpacked data for UpdatePriceFeed events raised by the Feeder contract.
type FeederUpdatePriceFeedIterator struct {
	Event *FeederUpdatePriceFeed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log       // Log channel receiving the found contract events
	sub  evrynet.Subscription // Subscription for errors, completion and termination
	done bool                 // Whether the subscription completed delivering logs
	fail error                // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FeederUpdatePriceFeedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeederUpdatePriceFeed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FeederUpdatePriceFeed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FeederUpdatePriceFeedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeederUpdatePriceFeedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeederUpdatePriceFeed represents a UpdatePriceFeed event raised by the Feeder contract.
type FeederUpdatePriceFeed struct {
	Updater      common.Address
	NewPriceFeed common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatePriceFeed is a free log retrieval operation binding the contract event 0x0c19c511fd3cdf7b4524c00cd2176627ee9fdd2644d68683e6cc04b79280f316.
//
// Solidity: event UpdatePriceFeed(address updater, address newPriceFeed)
func (_Feeder *FeederFilterer) FilterUpdatePriceFeed(opts *bind.FilterOpts) (*FeederUpdatePriceFeedIterator, error) {

	logs, sub, err := _Feeder.contract.FilterLogs(opts, "UpdatePriceFeed")
	if err != nil {
		return nil, err
	}
	return &FeederUpdatePriceFeedIterator{contract: _Feeder.contract, event: "UpdatePriceFeed", logs: logs, sub: sub}, nil
}

// WatchUpdatePriceFeed is a free log subscription operation binding the contract event 0x0c19c511fd3cdf7b4524c00cd2176627ee9fdd2644d68683e6cc04b79280f316.
//
// Solidity: event UpdatePriceFeed(address updater, address newPriceFeed)
func (_Feeder *FeederFilterer) WatchUpdatePriceFeed(opts *bind.WatchOpts, sink chan<- *FeederUpdatePriceFeed) (event.Subscription, error) {

	logs, sub, err := _Feeder.contract.WatchLogs(opts, "UpdatePriceFeed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeederUpdatePriceFeed)
				if err := _Feeder.contract.UnpackLog(event, "UpdatePriceFeed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
