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
    // struct TimeSlot {
    //     uint256 roomId; // room ID + renter addres + checkIn + checkOut
    //     address renter;
    //     bool booked;
    // }
    
    // Store hotel count - we cannot iterate or find size in mapping
    uint8 public roomsCount;

    // Amount is send to owners address , owner will auto set to the deployed one
    address owner;
    // mapping(uint256 =>Room) RommsMap;

    // Constructor called when contract is deployed or migrated
    constructor(){
        // Set the contract creater as the owner
        owner = msg.sender;
        for (uint256 i = 1; i < 20; i++) {
            createRoom(string(abi.encodePacked("room",Strings.toString(i))), owner);
            roomsCount++;
        }
    }

    Room[] public rooms;
    // TimeSlot[] public timeSlots;

    function createRoom(string memory _name,address _owner) internal{
        //creating the object of the structure in solidity 
        //  Room storage room =  RommsMap[hashId(_owner)];

        uint256 id = hashId(_owner);
            rooms.push(Room({
            id:id,
            name:_name,
            owner:_owner,
            capacity:18,
            workingSlots:10
            }));
        // uint startDate = 1514764800;// 2018-01-01 00:00:00
        // createTimeSlot(id, _owner,startDate);
        // createTimeSlot(id, _owner,startDate);
        // createTimeSlot(id, _owner,startDate);
        // createTimeSlot(id, _owner,startDate);
    }

    // function createTimeSlot(uint256 _roomId, address _renter, uint256 _date) internal{
    //     timeSlots.push(TimeSlot({
    //         roomId:createDeterministicId( _roomId,_renter,_date),
    //         renter:_renter,
    //         booked:false
    //         }));
    // }
        
    function allrooms()public view returns(Room  [] memory){
        return rooms;
    }
    // function getRoom(address _address) public view returns(uint256,string memory ,address){
    //     return(RommsMap[hashId(_address)].id, RommsMap[hashId(_address)].name,  RommsMap[hashId(_address)].owner);
    // }

    function getRoom() public view returns (string memory name, uint256 id) {
        Room storage room =  rooms[1];
        return (room.name, room.id);
    }

    // function getRoomCapacity(address _roomOwner) public view returns (uint256 capacity) {
    //     Room storage room =  rooms[hashId(_roomOwner)];
    //     return (room.capacity);
    // }

    // function getRoomOwner(address _roomOwner) public view returns (address addr) {
    //     Room storage room =  rooms[hashId(_roomOwner)];
    //     return (room.owner);
    // }

    function createDeterministicId(uint256 _roomId, address _owner, uint256 date) internal pure returns(uint256){
        return uint256(keccak256(abi.encodePacked(_roomId,_owner, date)));
    }
    
    function hashId(address _owner) internal pure returns(uint256){ 
        return uint256(keccak256(abi.encodePacked(_owner)));
    }
}
    // // Magic word is "Solidity"
    // function guess(string memory _word) public view returns (bool) {
    //     return keccak256(abi.encodePacked(_word)) == answer;
    // }

    // function getTimeSlot(uint256 _id) public view returns (address renter, bool booked) {
    //     TimeSlot storage timeslot =  timeSlots[_id];
    //     return (timeslot.renter, timeslot.booked);
    // }
    

    // function getTimeSlot(uint256 _id) public view returns (TimeSlot timeslot) {
    //     TimeSlot storage timeslot;
    //     return timeslot;
    // }




    // function bookRoom(uint256 _id, address renter, string memory checkIn, string memory checkOut) public returns (bool) {
    //     Room storage room =  rooms[_id];
    //         // if (room){

    //         // }
    //         room.renter=renter;
    //         room.checkIn=checkIn;
    //         room.checkOut= checkOut;
    //     return true;
    // }

    // // Magic word is "Solidity"
    // function guess(string memory _word) public view returns (bool) {
    //     return keccak256(abi.encodePacked(_word)) == answer;
    // }

    // function getTimeSlot(uint256 _id) public view returns (address renter, bool booked) {
    //     TimeSlot storage timeslot =  timeSlots[_id];
    //     return (timeslot.renter, timeslot.booked);
    // }
    

    // function getTimeSlot(uint256 _id) public view returns (TimeSlot timeslot) {
    //     TimeSlot storage timeslot;
    //     return timeslot;
    // }




    // function bookRoom(uint256 _id, address renter, string memory checkIn, string memory checkOut) public returns (bool) {
    //     Room storage room =  rooms[_id];
    //         // if (room){

    //         // }
    //         room.renter=renter;
    //         room.checkIn=checkIn;
    //         room.checkOut= checkOut;
    //     return true;
    // }

