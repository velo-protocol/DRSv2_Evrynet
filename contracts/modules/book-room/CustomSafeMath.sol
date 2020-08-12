pragma solidity ^0.5.0;

library CustomSafeMath {
	function mul(uint a, uint b) internal pure returns (uint) {
		uint c = a * b;
		require(a == 0 || c / a == b);
		return c;
	}

	function div(uint a, uint b) internal pure returns (uint) {
		require (b > 0 , "CustomSafeMath: division by zero") ;
		uint c = a / b;
		return c;
	}

	function sub(uint a, uint b) internal pure returns (uint) {
		require(b <= a);
		return a - b;
	}

	function add(uint a, uint b) internal pure returns (uint) {
		uint c = a + b;
		require(c >= a);
		return c;
	}

	function diff(uint a, uint b) internal pure returns (uint) {
		return a > b ? sub(a, b) : sub(b, a);
	}

	function gt(uint a, uint b) internal pure returns(bytes1) {
		bytes1 c;
		c = 0x00;
		if (a > b) {
			c = 0x01;
		}
		return c;
	}
}