pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/AccessControl.sol";

contract Cicada is AccessControl {
  mapping(address => uint256) private _storageAllowance; // amount of gigabytes allowed per user)
  mapping(address => uint256) private _timeAllowance;

  uint256 public memoryRate; // cost of memory per gb in wei
  uint256 public monthlyFee;
  bytes32 public constant REGISTERED_USER = keccak256("REGISTERED_USER");

  constructor(uint256 initRate, uint256 monthly) {
    _setupRole(DEFAULT_ADMIN_ROLE, msg.sender);
    memoryRate = initRate;
    monthlyFee = monthly;
  }

  function setRate(uint256 rate) external onlyRole(DEFAULT_ADMIN_ROLE) {
    memoryRate = rate;
  }

  function setMonthlyFee(uint256 monthly) external onlyRole(DEFAULT_ADMIN_ROLE) {
    monthlyFee = monthly;
  }

  /*
    Pay by the gigabyte
  */
  function payForStorage(uint256 gigabytes) external payable {
    require(msg.value == gigabytes * memoryRate, "You did not pay the exact amount");
    require(_checkTimeAllowance(msg.sender), "You do not have an active subscription");
    _storageAllowance[msg.sender] += gigabytes;
  }

  /*
     Enter a user into the registry
  */
  function registerUser() external {
    _setupRole(REGISTERED_USER, msg.sender);
    _storageAllowance[msg.sender] = 1; // allot users 1gb of free storage
    _timeAllowance[msg.sender] = 31 days;
  }

  /*
     Renew a subscription to the service
  */
  function renewSubscription() external payable onlyRole(REGISTERED_USER) {
    require(block.timestamp - _timeAllowance[msg.sender] > 0, "You still have an active subscription");
    require(msg.value == monthlyFee, "You did not pay the correct amount");

    _timeAllowance[msg.sender] += 31 days;
  }

  /*
     Withdraw funds from contract
  */
  function withdraw() external onlyRole(DEFAULT_ADMIN_ROLE) {
    require(address(this).balance > 0, "You don't have any ETH to withdraw!");

    (bool success,) = payable(msg.sender).call{value: address(this).balance}("");
    require(success, "Something went wrong...");
  }

  function _checkTimeAllowance(address user) view internal returns (bool) {
    return _timeAllowance[user] > 0; // check if user has renewed subscription
  }
}
