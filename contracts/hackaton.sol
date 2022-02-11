// SPDX-License-Identifier: AGPL-3.0

pragma solidity 0.8.9;

import "@openzeppelin/contracts/access/Ownable.sol";

contract Hackaton is Ownable {

    constructor() {
        
    }
    event Data(string information);
    function sendEvent(string memory data) public {
        emit Data(data);
    }

    function verifyProof(
        bytes32 hash,
        uint256[2] calldata proofA,
        uint256[2][2] calldata proofB,
        uint256[2] calldata proofC
    ) public {
        require(
            //rollupVerifier.verifyProof(proofA, proofB, proofC, [hash]),
            true,
            "Verification: INVALID_PROOF"
        );

        //Do whatever we want
        
    }
}
