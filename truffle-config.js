require('dotenv').config();
require("@babel/polyfill");
const PrivateKeyProvider = require("truffle-privatekey-provider");
const LedgerWalletProvider = require('truffle-ledger-provider');
//Contract data: Yes
// Browser Support: No

const ledgerOptions = {
  networkId: "*", // mainnet
  path: "44'/60'/0'/0", // ledger default derivation path
  askConfirm: false,
  accountsLength: 1,
  accountsOffset: 0
}; // use default options

module.exports = {
  networks: {
    local: {
      host: "127.0.0.1",
      port: 8545,
      network_id: "*"
    },
    dev: {
      provider: () => new PrivateKeyProvider(process.env.DEV_SCC_PK, process.env.DEV_SCC_HOST),
      network_id: "*",
      gasPrice: 1000000000,
      gas: 6357193,
      production: false
    },
    ledger: {
      provider: new LedgerWalletProvider(ledgerOptions, process.env.DEV_SCC_HOST),
      network_id: 3,
      gas: 4600000
    },
    coverage: {
      host: "localhost",
      network_id: "*",
      port: 8555,
      gas: 0xfffffffffff,
      gasPrice: 0x01
    }
  },
  // Set default mocha options here, use special reporters etc.
  mocha: {
    // timeout: 100000
  },
  compilers: {
    solc: {
      version: " 0.5.15",
      settings: {
        optimizer: {
          enabled: true,
          runs: 200
        },
        evmVersion: "constantinople"
      }
    }
  },
  plugins: ["solidity-coverage"]
};
