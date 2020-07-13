const Heart = artifacts.require('Heart');
const ReserveManager = artifacts.require('ReserveManager');
const Price = artifacts.require('Price');
const DRS = artifacts.require('DigitalReserveSystem');
const Token = artifacts.require('Token');
const Hasher = artifacts.require('Hasher');
const StableCredit = artifacts.require('StableCredit');
const h = require("../testhelper");
const MockContract = artifacts.require("MockContract");

let drs, heart, price, reserveManager, veloCollateralAsset, otherCollateralAsset, stableCreditVUSD, mocks;

const velo = "VELO";
const veloBytes32 = web3.utils.padRight(web3.utils.fromAscii(velo), 64);
const usdBytes32 = web3.utils.padRight(web3.utils.fromAscii("USD"), 64);

contract("ReserveManager Scenario Test", async accounts => {

  before(async () => {
    mocks = {
      price: await MockContract.new(),
    }
    price = await Price.at(mocks.price.address);
  });

  it("should work!", async () => {
    const calInputs = {
      issuanceFeeRate: h.decimal7(0.05), // 0.05 (5%)
      price: h.decimal7(10.00), // 10.00 USD/VELO
      collateralRatio: h.decimal7(1.3), // 1.30
      peggedValue: h.decimal7(1.00) // 1.00
    };

    const [ bob] = accounts;

    // Setup the whole ecosystem
    heart = await Heart.new();
    drs = await DRS.new(heart.address);
    reserveManager = await ReserveManager.new(heart.address);
    veloCollateralAsset = await Token.new(velo, velo, 7);

    await mocks.price.givenMethodReturnUint(
      price.contract.methods.get().encodeABI(),
      calInputs.price
    );

    // Setup heart
    heart.addPrice(web3.utils.soliditySha3(veloBytes32, usdBytes32), price.address);
    heart.setReserveManager(reserveManager.address);
    heart.setDrsAddress(drs.address);
    // Mint VELO for Bob, so he can spend VELO afterward
    await veloCollateralAsset.mint(bob, 10000);
    // Setup Heart
    await heart.setCollateralAsset(veloBytes32, veloCollateralAsset.address, calInputs.collateralRatio);
    await heart.setTrustedPartner(bob);
    await heart.setReserveFreeze(veloBytes32,1);
   const tim=  await heart.getReserveFreeze(veloBytes32);
   console.log("Freeze  time==>"+tim);
    await heart.setCreditIssuanceFee(calInputs.issuanceFeeRate);
    await heart.setAllowedLink(web3.utils.soliditySha3(veloBytes32, usdBytes32), true);
    // Approve RM to spend VELO
    await veloCollateralAsset.approve(reserveManager.address, 1000, {from: bob});
    const lockReserveResult = await reserveManager.lockReserve(veloBytes32,bob,1000)
    const  lockReserveEvent=lockReserveResult.logs[0].args
    h.assert.equalNumber(await veloCollateralAsset.balanceOf(bob), 9000);
    try {
      await reserveManager.releaseReserve(lockReserveEvent.lockedReserveId,veloBytes32,1000,{from: bob})
    } catch (err) {
      assert.equal(err.reason, "release time not reach")
    }

  });
});
