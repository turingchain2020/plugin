pragma solidity ^0.5.0;

contract BridgeRegistry {

    address public turingchainBridge;
    address public bridgeBank;
    address public oracle;
    address public valset;
    uint256 public deployHeight;

    event LogContractsRegistered(
        address _turingchainBridge,
        address _bridgeBank,
        address _oracle,
        address _valset
    );
    
    constructor(
        address _turingchainBridge,
        address _bridgeBank,
        address _oracle,
        address _valset
    )
        public
    {
        turingchainBridge = _turingchainBridge;
        bridgeBank = _bridgeBank;
        oracle = _oracle;
        valset = _valset;
        deployHeight = block.number;

        emit LogContractsRegistered(
            turingchainBridge,
            bridgeBank,
            oracle,
            valset
        );
    }
}