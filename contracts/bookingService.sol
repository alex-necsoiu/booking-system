// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
 import "node_modules/@openzeppelin/contracts/utils/Strings.sol";

contract BookRoom {

    struct Room {
        uint256 id;
        string name;
        address owner;
        uint256 capacity;
        uint256 workingSlots;
    }
    
    // Store hotel count - we cannot iterate or find size in mapping
    uint8 public roomsCount;

    // Amount is send to owners address , owner will auto set to the deployed one
    address owner;
    // mapping (uint256 => Room[])public rooms;
   Room[] public rooms;
    // Constructor called when contract is deployed or migrated
    constructor(){
        // Set the contract creater as the owner
        owner = msg.sender;
        for (uint256 i = 1; i < 5; i++) {
            createRoom(string(abi.encodePacked("room",Strings.toString(i))), owner);
            roomsCount++;
        }
    }

    function createRoom( string memory _name,address _owner) public {
            rooms.push(
            Room({
                id:hashId(_owner),
                name:_name, 
                owner:_owner, 
                capacity:20, 
                workingSlots:10
                }));
        roomsCount++;
    }

    function allRoomsByDate()public view returns(Room  [] memory){
        return rooms;
    }
    function getRoom(uint idx) public view returns (string memory name, uint256 id) {
        return (rooms[idx].name, rooms[idx].id);
    }
    function getRoomCount()     public view returns(uint) {
        return rooms.length;
    }
    function hashId(address _owner) internal pure returns(uint256){ 
        return uint256(keccak256(abi.encodePacked(_owner)));
    }
}