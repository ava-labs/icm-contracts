pragma solidity 0.8.25;

import {IValidatorManager} from "./interfaces/IValidatorManager.sol";
import {PoSValidatorInfo} from "./interfaces/IStakingManager.sol";
import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/IERC20.sol";
import {PChainOwner} from "./interfaces";
//very dumb mock coin im making, mainly for the fun of it but the plan is for it to be used with the other smart contract

//I wrote all this without actually looking at ERC-20 standards, so now I am and its not too far off from what I have so cool

struct AllowanceInfo {
    address spender;
    address owner;
    uint256 originalAllowance;
    uint256 allowanceRemaining;
}

contract EmCoin is IERC20 {
    mapping (address => uint256) balances;
    //kinda scuffed but im not sure how else to do this in an efficient way
    mapping (bytes32 _hashedAddr => AllowanceInfo) allowances;
    uint256 public wallet = 1e9;
    // uint256 receiptNum = 1;
    function name() public pure returns (string memory) {
        return "EmCoin";
    }
    function symbol() public pure returns (string memory) {
        return "EMC";
    }
    function decimals() public pure returns (uint8) {
        return 0;
    }
    function totalSupply() pure public returns (uint256) {
        return 1e9;
    }
    function balanceOf(address _owner) public view returns (uint256 balance) {
        return balances[_owner];
    }
    function transfer(address _to, uint256 _value) public returns (bool success) {
        assert(balances[msg.sender] >= _value);
        return transferFunds(msg.sender, _to, _value);
    }

    function transferFrom(address _from, address _to, uint256 _value) public returns (bool success) {
        success = false;
        bytes32 _hashedAddr = hashAddresses(_from, _to);

        require(allowances[_hashedAddr].allowanceRemaining != 0);
        require(allowances[_hashedAddr].spender == _to);
        require(allowances[_hashedAddr].owner == _from);
        require(allowances[_hashedAddr].allowanceRemaining >= _value);
        allowances[_hashedAddr].allowanceRemaining -= _value;
        return transferFunds(_from, _to, _value);
    }

    function approve(address _spender, uint256 _value) public returns (bool success) {
        bytes32 _hashedAddr = hashAddresses(msg.sender, _spender);
        if (allowances[_hashedAddr].allowanceRemaining == allowances[_hashedAddr].originalAllowance) {
            allowances[_hashedAddr] = AllowanceInfo(_spender, msg.sender, _value, _value);
            emit Approval(msg.sender, _spender, _value);
            return true;
        }
        else {
            //sets the balance back to zero if money was taken out, incase of the event where money was taken
            //prematurely on the chain, saves them from being taken again
            allowances[_hashedAddr] = AllowanceInfo(_spender, msg.sender, 0, 0);
            return false;
        }
    }
    function hashAddresses(address _from, address _to)pure internal returns (bytes32) {
        return keccak256(abi.encodePacked(_from, _to));
    }
    function allowance(address _owner, address _spender) public view returns (uint256 remaining) {
        bytes32 _hashedAddr = hashAddresses(_owner, _spender);
        return allowances[_hashedAddr].allowanceRemaining;
    }

    //obviously would want something in return but for now this works just for the sake of getting tokens
    function buyCoin(uint val) public {
        balances[msg.sender] += val;
        wallet -= val;
    }

    function transferFunds(address _from, address _to, uint256 _value) internal returns (bool) {
        if (balances[_from] < _value) return false;
        balances[_from] -= _value;
        balances[_to] += _value;
        emit Transfer(_from, _to, _value);
        return true;
    }


}

contract EmreValidator {
    address constant public EmTokenAddress; //once I deploy I should be able to put the address here
    uint8 public validatorCount = 0;
    mapping (address nodeID => bytes32 validationID) validators;
    address constant public validatorManagerAddress;
    function becomeValidator (
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner
        ) external returns (bytes32){

        require (validatorCount < 10); //check to see if there are slots open
        require (validators[nodeID] == 0); // check to see if node if addr is already validating
        validatorCount++;

        IERC20(EmTokenAddress).transferFrom(nodeID, block.number, 10);
        //potential problem, dont want to make them a validator before marking them as a validator
        //might be smart to have it point to a struct that says if its a validator and holds its validationID, but for now well try this out
        bytes32 _validationID = IValidatorManager(validatorManagerAddress).initiateValidatorRegistration(nodeID, blsPublicKey, remainingBalanceOwner, disableOwner, 100);
        validators[nodeID] = _validationID;
        return _validationID;
    } 

    function stopValidating(bytes32 validationID) external {
        require (validatorCount > 0);
        require (validators[msg.sender] == validationID);
        validators[msg.sender] = false;
        validatorCount--;
        IValidatorManager(validatorManagerAddress).initiateValidatorRemoval(validationID);
        IERC20(EmTokenAddress).transfer(msg.sender, 10);
        //call the functions for emtoken to send the money back and remove them as validator
    }
}