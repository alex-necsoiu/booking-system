// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
contract X {
    struct PRODUCT {
        uint256 productid;
        uint price; 
        uint quantity;
    }

    mapping (uint256 => PRODUCT[]) products;

    function appendDetails(uint256 pid, uint256 productid, uint price, uint quantity) payable public {
        products[pid].push(PRODUCT(productid, price, quantity));
    }

    function getDetails(uint256 pid, uint idx) public view returns(uint256 productid, uint price, uint quantity) {
        PRODUCT storage p = products[pid][idx];
        productid = p.productid;
        price = p.price;
        quantity = p.quantity;
    }

    function getDetailsCount(uint256 pid)     public view returns(uint) {
        return products[pid].length;
    }

    function getDetailsByDate(uint256 pid)     public view returns(PRODUCT[] memory) {
        return products[pid];
    }
}