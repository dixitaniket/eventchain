// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

contract TestEvent {
    event Launch(uint8 indexed number, uint8 indexed toadd);

    constructor() {}

    function Trigger(uint8 number, uint8 toadd) public {
        emit Launch(number,toadd);
    }
}
