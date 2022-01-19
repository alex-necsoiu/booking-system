// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

//example hotel smart contract

contract BookRoom {


    struct Room {
        uint256 id;
        string name;
        address owner;
        // TimeSlot[12] timeslot;
    }
    struct TimeSlot {
        uint256 roomId;
        address renter;
        string checkIn;
        string checkOut;
        bool available;
    }

    // mapping acts as a associative array of key => value , here key is unsigned
    // int and value is structure
    // mapping (uint => Room) public rooms;

    // Store hotel count - we cannot iterate or find size in mapping
    uint256 public roomsCount;

    // Amount is send to owners address , owner will auto set to the deployed one
    address owner;

    // Constructor called when contract is deployed or migrated
    // constructor() public {
    //     // Set the contract creater as the owner
    //     owner = msg.sender;

    //     for (uint256 i = 1; i < 20; i++) {
    //         addRoom(i, owner);
    //     }
    // }
    Room[] public rooms;

    function createRoom(uint256 _id, string memory _name,address _owner) internal returns (bool) {
        rooms.push(Room({
            id:_id,
            name:_name,
            owner:_owner
            }));
        return true;
    }

    function getRoom(uint256 _id) public view returns (string memory name, uint id) {
        Room storage room =  rooms[_id];
        return (room.name, room.id);
    }

    // function bookRoom(uint256 _id, address renter, string memory checkIn, string memory checkOut) public returns (bool) {
    //     Room storage room =  rooms[_id];
    //         // if (room){

    //         // }
    //         room.renter=renter;
    //         room.checkIn=checkIn;
    //         room.checkOut= checkOut;

    //     return true;
    // }

}
