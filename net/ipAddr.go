//ipv4 address is a 32 bit Integer 
//it addresses down to a network interface card on a single device

// IP address is composed of two parts
// (1) The address of the network in which device resides
// (2) The address of the device within the network 

// class A network : First Byte => identifies the network
// class B network : First 2 Bytes => identifies the network 
// Class C network : First 3 bytes => identifies the network

// to find the network of a device
// bitwise AND its IP address with the network mask 

// to find the device address within the subnet 
// create one's complement of the network mask 
// then bitwise AND it with the IP address 

// for 16 bit network address , netmask is 255.255.0.0
// for 24 bit network address , netmask is 255.255.255.0
// for 23 bit network address , netmask is 255.255.254.0
// for 22 bit network address , netmask is 255.255.252.0


