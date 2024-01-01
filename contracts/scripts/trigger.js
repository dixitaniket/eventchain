const hre = require("hardhat");

async function main() {
    const testEvent = await hre.ethers.getContractAt("TestEvent","0x5FbDB2315678afecb367f032d93F642f64180aa3");
    let tx=await testEvent.Trigger("100","200");
    console.log(tx)
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});
