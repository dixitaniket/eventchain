const hre = require("hardhat");
const readline = require("readline").createInterface({
    input: process.stdin,
    output: process.stdout
});

async function main() {
    // get user input for num and toadd
    const arg1 = await new Promise(resolve => {
        readline.question("Enter num: ", (answer1) => {
            resolve(answer1);
        });
    });

    const arg2 = await new Promise(resolve => {
        readline.question("Enter to add: ", (answer2) => {
            resolve(answer2);
            readline.close();
        });
    });

    const testEvent = await hre.ethers.getContractAt("TestEvent", "0x5FbDB2315678afecb367f032d93F642f64180aa3");
    let tx = await testEvent.Trigger(arg1, arg2);
    console.log(tx);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});
