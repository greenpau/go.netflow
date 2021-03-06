package net2

type IPProtocol struct {
	Keyword  string
	Protocol string
}

var IPProtocolMap = map[int]IPProtocol{
	0:   IPProtocol{"HOPOPT", "IPv6 Hop-by-Hop Option"},
	1:   IPProtocol{"ICMP", "Internet Control Message Protocol"},
	2:   IPProtocol{"IGMP", "Internet Group Management Protocol"},
	3:   IPProtocol{"GGP", "Gateway-to-Gateway Protocol"},
	4:   IPProtocol{"IP-in-IP", "IP-in-IP"},
	5:   IPProtocol{"ST", "Internet Stream Protocol"},
	6:   IPProtocol{"TCP", "Transmission Control Protocol"},
	7:   IPProtocol{"CBT", "Core-based trees"},
	8:   IPProtocol{"EGP", "Exterior Gateway Protocol"},
	9:   IPProtocol{"IGP", "Interior Gateway Protocol"},
	10:  IPProtocol{"BBN-RCC-MON", "BBN RCC Monitoring"},
	11:  IPProtocol{"NVP-II", "Network Voice Protocol"},
	12:  IPProtocol{"PUP", "Xerox PUP"},
	13:  IPProtocol{"ARGUS", "ARGUS"},
	14:  IPProtocol{"EMCON", "EMCON"},
	15:  IPProtocol{"XNET", "Cross Net Debugger"},
	16:  IPProtocol{"CHAOS", "Chaos"},
	17:  IPProtocol{"UDP", "User Datagram Protocol"},
	18:  IPProtocol{"MUX", "Multiplexing"},
	19:  IPProtocol{"DCN-MEAS", "DCN Measurement Subsystems"},
	20:  IPProtocol{"HMP", "Host Monitoring Protocol"},
	21:  IPProtocol{"PRM", "Packet Radio Measurement"},
	22:  IPProtocol{"XNS-IDP", "XEROX NS IDP"},
	23:  IPProtocol{"TRUNK-1", "Trunk-1"},
	24:  IPProtocol{"TRUNK-2", "Trunk-2"},
	25:  IPProtocol{"LEAF-1", "Leaf-1"},
	26:  IPProtocol{"LEAF-2", "Leaf-2"},
	27:  IPProtocol{"RDP", "Reliable Datagram Protocol"},
	28:  IPProtocol{"IRTP", "Internet Reliable Transaction Protocol"},
	29:  IPProtocol{"ISO-TP4", "ISO Transport Protocol Class 4"},
	30:  IPProtocol{"NETBLT", "Bulk Data Transfer Protocol"},
	31:  IPProtocol{"MFE-NSP", "MFE Network Services Protocol"},
	32:  IPProtocol{"MERIT-INP", "MERIT Internodal Protocol"},
	33:  IPProtocol{"DCCP", "Datagram Congestion Control Protocol"},
	34:  IPProtocol{"3PC", "Third Party Connect Protocol"},
	35:  IPProtocol{"IDPR", "Inter-Domain Policy Routing Protocol"},
	36:  IPProtocol{"XTP", "Xpress Transport Protocol"},
	37:  IPProtocol{"DDP", "Datagram Delivery Protocol"},
	38:  IPProtocol{"IDPR-CMTP", "IDPR Control Message Transport Protocol"},
	39:  IPProtocol{"TP++", "TP++ Transport Protocol"},
	40:  IPProtocol{"IL", "IL Transport Protocol"},
	41:  IPProtocol{"IPv6", "IPv6 Encapsulation"},
	42:  IPProtocol{"SDRP", "Source Demand Routing Protocol"},
	43:  IPProtocol{"IPv6-Route", "Routing Header for IPv6"},
	44:  IPProtocol{"IPv6-Frag", "Fragment Header for IPv6"},
	45:  IPProtocol{"IDRP", "Inter-Domain Routing Protocol"},
	46:  IPProtocol{"RSVP", "Resource Reservation Protocol"},
	47:  IPProtocol{"GRE", "Generic Routing Encapsulation"},
	48:  IPProtocol{"MHRP", "Mobile Host Routing Protocol"},
	49:  IPProtocol{"BNA", "BNA"},
	50:  IPProtocol{"ESP", "Encapsulating Security Payload"},
	51:  IPProtocol{"AH", "Authentication Header"},
	52:  IPProtocol{"I-NLSP", "Integrated Net Layer Security Protocol"},
	53:  IPProtocol{"SWIPE", "SwIPe"},
	54:  IPProtocol{"NARP", "NBMA Address Resolution Protocol"},
	55:  IPProtocol{"MOBILE", "IP Mobility"},
	56:  IPProtocol{"TLSP", "Transport Layer Security Protocol"},
	57:  IPProtocol{"SKIP", "Simple Key-Management for Internet Protocol"},
	58:  IPProtocol{"IPv6-ICMP", "ICMP for IPv6"},
	59:  IPProtocol{"IPv6-NoNxt", "No Next Header for IPv6"},
	60:  IPProtocol{"IPv6-Opts", "Destination Options for IPv6"},
	61:  IPProtocol{"", "Any host internal protocol"},
	62:  IPProtocol{"CFTP", "CFTP"},
	63:  IPProtocol{"", "Any local network"},
	64:  IPProtocol{"SAT-EXPAK", "SATNET and Backroom EXPAK"},
	65:  IPProtocol{"KRYPTOLAN", "Kryptolan"},
	66:  IPProtocol{"RVD", "MIT Remote Virtual Disk Protocol"},
	67:  IPProtocol{"IPPC", "Internet Pluribus Packet Core"},
	68:  IPProtocol{"", "Any distributed file system"},
	69:  IPProtocol{"SAT-MON", "SATNET Monitoring"},
	70:  IPProtocol{"VISA", "VISA Protocol"},
	71:  IPProtocol{"IPCU", "Internet Packet Core Utility"},
	72:  IPProtocol{"CPNX", "Computer Protocol Network Executive"},
	73:  IPProtocol{"CPHB", "Computer Protocol Heart Beat"},
	74:  IPProtocol{"WSN", "Wang Span Network"},
	75:  IPProtocol{"PVP", "Packet Video Protocol"},
	76:  IPProtocol{"BR-SAT-MON", "Backroom SATNET Monitoring"},
	77:  IPProtocol{"SUN-ND", "SUN ND PROTOCOL-Temporary"},
	78:  IPProtocol{"WB-MON", "WIDEBAND Monitoring"},
	79:  IPProtocol{"WB_EXPAK", "WIDEBAND EXPAK"},
	80:  IPProtocol{"ISO-IP", "International Organization for Standardization Internet Protocol"},
	81:  IPProtocol{"VMTP", "Versatile Message Transaction Protocol"},
	82:  IPProtocol{"SECURE-VMTP", "Secure Versatile Message Transaction Protocol"},
	83:  IPProtocol{"VINES", "VINES"},
	84:  IPProtocol{"TTP/IPTM", "TTP/Internet Protocol Traffic Manager"},
	85:  IPProtocol{"NSFNET-IGP", "NSFNET-IGP"},
	86:  IPProtocol{"DGP", "Dissimilar Gateway Protocol"},
	87:  IPProtocol{"TCF", "TCF"},
	88:  IPProtocol{"EIGRP", "EIGRP"},
	89:  IPProtocol{"OSPF", "OSPF"},
	90:  IPProtocol{"Sprite-RPC", "Sprite RPC Protocol"},
	91:  IPProtocol{"LARP", "Locus Address Resolution Protocol"},
	92:  IPProtocol{"MTP", "Multicast Transport Protocol"},
	93:  IPProtocol{"AX.25", "AX.25"},
	94:  IPProtocol{"IPIP", "IP-within-IP Encapsulation Protocol"},
	95:  IPProtocol{"MICP", "Mobile Internetworking Control Protocol"},
	96:  IPProtocol{"SCC-SP", "Semaphore Communications Sec. Pro"},
	97:  IPProtocol{"ETHERIP", "Ethernet-within-IP Encapsulation"},
	98:  IPProtocol{"ENCAP", "Encapsulation Header"},
	99:  IPProtocol{"", "Any private encryption scheme"},
	100: IPProtocol{"GMTP", "GMTP"},
	101: IPProtocol{"IFMP", "Ipsilon Flow Management Protocol"},
	102: IPProtocol{"PNNI", "PNNI over IP"},
	103: IPProtocol{"PIM", "Protocol Independent Multicast"},
	104: IPProtocol{"ARIS", "IBM's ARIS (Aggregate Route IP Switching) Protocol"},
	105: IPProtocol{"SCPS", "Space Communications Protocol Standards"},
	106: IPProtocol{"QNX", "QNX"},
	107: IPProtocol{"A/N", "Active Networks"},
	108: IPProtocol{"IPComp", "IP Payload Compression Protocol"},
	109: IPProtocol{"SNP", "Sitara Networks Protocol"},
	110: IPProtocol{"Compaq-Peer", "Compaq Peer Protocol"},
	111: IPProtocol{"IPX-in-IP", "IPX in IP"},
	112: IPProtocol{"VRRP", "Virtual Router Redundancy Protocol, Common Address Redundancy Protocol"},
	113: IPProtocol{"PGM", "PGM Reliable Transport Protocol"},
	114: IPProtocol{"", "Any 0-hop protocol"},
	115: IPProtocol{"L2TP", "Layer Two Tunning Protocol Version 3"},
	116: IPProtocol{"DDX", "D-II Data Exchange"},
	117: IPProtocol{"IATP", "Interactive Agent Transfer Protocol"},
	118: IPProtocol{"STP", "Schedule Transfer Protocol"},
	119: IPProtocol{"SRP", "SpectraLink Radio Protocol"},
	120: IPProtocol{"UTI", "Universal Transport Interface Protocol"},
	121: IPProtocol{"SMP", "Simple Message Protocol"},
	122: IPProtocol{"SM", "Simple Multicast Protocol"},
	123: IPProtocol{"PTP", "Performance Transparency Protocol"},
	124: IPProtocol{"IS-IS over IPv4", "Intermediate System to Intermediate System (IS-IS) Protocol over IPv4"},
	125: IPProtocol{"FIRE", "Flexible Intra-AS Routing Environment"},
	126: IPProtocol{"CRTP", "Combat Radio Transport Protocol"},
	127: IPProtocol{"CRUDP", "Combat Radio User Datagram"},
	128: IPProtocol{"SSCOPMCE", "Service-Specific Connection-Oriented Protocol in a Multilink and Connectionless Environment"},
	129: IPProtocol{"IPLT", ""},
	130: IPProtocol{"SPS", "Secure Packet Shield"},
	131: IPProtocol{"PIPE", "Private IP Encapsulation within IP"},
	132: IPProtocol{"SCTP", "Stream Control Transmission Protocol"},
	133: IPProtocol{"FC", "Fibre Channel"},
	134: IPProtocol{"RSVP-E2E-IGNORE", "Reservation Protocol (RSVP) End-to-End Ignore"},
	135: IPProtocol{"Mobility Header", "Mobility Extension Header for IPv6"},
	136: IPProtocol{"UDPLite", "Lightweight User Datagram Protocol"},
	137: IPProtocol{"MPLS-in-IP", "Multiprotocol Label Switching Ecapsulated in IP"},
	138: IPProtocol{"manet", "MANET Protocols"},
	139: IPProtocol{"HIP", "Host Identity Protocol"},
	140: IPProtocol{"Shim6", "Site Multihoming by IPv6 Intermediation"},
	141: IPProtocol{"WESP", "Wrapped Encapsulating Security Payload"},
	142: IPProtocol{"ROHC", "Robust Header Compression"},
}

// TODO: fill..
var TCPUDPPortMap = map[int]map[string]string{
	0: {
		"TCP": "Programming technique for specifying system-allocated (dynamic) ports",
		"UDP": "Reserved",
	},
	1: {
		"TCP": "TCPMUX",
		"UDP": "TCPMUX",
	},
	7: {
		"TCP": "Echo Protocol",
		"UDP": "Echo Protocol",
	},
	20: {
		"TCP": "FTP data transfer",
		"UDP": "FTP data transfer",
	},
	21: {
		"TCP": "FTP control",
	},
	22: {
		"TCP": "SSH/SCP/SFTP",
		"UDP": "SSH/SCP/SFTP",
	},
	23: {
		"TCP": "Telnet",
		"UDP": "Telnet",
	},
	25: {
		"TCP": "SMTP",
	},
	42: {
		"TCP": "ARPA Host Name Server Protocol",
		"UDP": "ARPA Host Name Server Protocol",
	},
	43: {
		"TCP": "WHOIS",
	},
	53: {
		"TCP": "DNS",
		"UDP": "DNS",
	},
	69: {
		"UDP": "TFTP",
	},
	80: {
		"TCP": "HTTP",
		"UDP": "QUIC",
	},
	123: {
		"TCP": "NTP",
		"UDP": "NTP",
	},
	143: {
		"TCP": "IMAP",
	},
	161: {
		"UDP": "SNMP",
	},
	177: {
		"TCP": "XDMCP",
		"UDP": "XDMCP",
	},
	179: {
		"TCP": "BGP",
	},
	194: {
		"TCP": "IRC",
		"UDP": "IRC",
	},
	443: {
		"TCP": "HTTPS",
		"UDP": "QUIC",
	},
	554: {
		"TCP": "RTSP",
		"UDP": "RTSP",
	},
	944: {
		"UDP": "NFS",
	},
}
