package protocols

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"log"
)

type Protocol struct {
	Name        string // Protocol name
	FromIP      string // From IP address
	ToIP        string // To IP addess
	FromPortNum string // From port number
	ToPortNum   string // To port number
	Description string // Other protocol specific information
}

// ADD the implemented protocol here
// this list cannot be exported
var protocolsList = [...]func(gopacket.Packet, chan Protocol){
	DetectHTTP,
	DetectSSH,
}

func Detect(packet gopacket.Packet, detectedProt chan Protocol) {
	for i := range protocolsList {
		// call the corresponding detect function
		protocolsList[i](packet, detectedProt)
	}
}

func GetIPAddresses(packet gopacket.Packet) (string, string) {
	/*
		Returns the IP address of source and destination.
	*/
	if ipLayer := packet.Layer(layers.LayerTypeIPv4); ipLayer != nil {
		ip, _ := ipLayer.(*layers.IPv4)
		return ip.SrcIP.String(), ip.DstIP.String()
	} else if ipLayer := packet.Layer(layers.LayerTypeIPv6); ipLayer != nil {
		ip, _ := ipLayer.(*layers.IPv4)
		return ip.SrcIP.String(), ip.DstIP.String()
	}
	// ?
	log.Fatal("No IPv4/IPv6 layer found in the packet.")
	return "", ""
}

func GetPortAddresses(packet gopacket.Packet) (string, string) {
	/*
		Returns the port numbers of source and destination
	*/
	if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
		tcp, _ := tcpLayer.(*layers.TCP)
		return tcp.SrcPort.String(), tcp.DstPort.String()
	}
	// ?
	log.Fatal("No TCP layer found in the packet.")
	return "", ""
}
