package main

import (
	"net"
	"strconv"

	"github.com/jedib0t/go-pretty/v6/table"
)

type CidrCmds struct {
	Show  CidrShowCmd  `cmd:""`
	Range CidrRangeCmd `cmd:""`
}

type CidrShowCmd struct {
	Cidr     string `arg:""`
	Division bool   `short:"d" help:"interactive division"`
}

func (cmd *CidrShowCmd) Run(g *GlobalSettings) error {
	if cmd.Division {
		return interativeCidrDivision(cmd.Cidr)
	}

	cidr, err := getCidrDivision(cmd.Cidr)
	if err != nil {
		return err
	}

	t := table.NewWriter()
	t.SetTitle("IP Range")
	t.SetStyle(table.StyleLight)
	t.Style().Options.SeparateRows = true
	t.AppendRow(table.Row{"range", cidr.start.String() + " -> " + cidr.end.String()})
	t.AppendRow(table.Row{"capacity", cidr.capacity})
	t.AppendRow(table.Row{"mask", net.IP(cidr.mask).String() + " (" + cidr.mask.String() + ")"})
	println(t.Render())

	return nil
}

type cidrBlock struct {
	start, end net.IP
	capacity   uint32
	mask       net.IPMask
}

func getCidrDivision(cidr string) (*cidrBlock, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	mask := ipnet.Mask
	// e.g.
	// cidr = 192.168.0.0/24
	// ones, bits = 24, 32
	ones, _ := mask.Size()
	start := ip.Mask(mask)
	end := getIPEnd(start, ones)
	capacity := ip2Int(end) - ip2Int(start) + 1

	return &cidrBlock{start, end, capacity, mask}, nil
}

func getIPEnd(start net.IP, prefix int) net.IP {
	end := net.IP(make([]byte, 4))
	copy(end, start)

	// fragment index, byte index
	fi, bi := prefix/8, prefix%8

	// mask the first byte
	end[fi] |= 1<<(8-bi) - 1
	// set following bytes to 255
	for i := fi + 1; i < 4; i++ {
		end[i] = 255
	}

	return end
}

func getIPNext(start net.IP, prefix int) net.IP {
	end := net.IP(make([]byte, 4))
	copy(end, start)
	fi, bi := prefix/8, prefix%8
	end[fi] += 1 << (8 - bi)
	if end[fi] > 255 {
		end[fi] = 0
		end[fi+1] += 1
	}

	return end
}

func interativeCidrDivision(cidr string) error {
	// TODO interactive with bubbletea
	// <-, merge
	// ->, divide
	// c, confirm
	println("Interactive CIDR division is not implemented yet, use: https://www.davidc.net/sites/default/subnets/subnets.html")
	return nil
}

type CidrRangeCmd struct {
	Ips []string `arg:""`
}

func (cmd *CidrRangeCmd) Run(g *GlobalSettings) error {
	ips := make([]net.IP, len(cmd.Ips))
	for i, ipStr := range cmd.Ips {
		ips[i] = net.ParseIP(ipStr)
	}

	cidr := getSmallestCidr(ips)
	println("Smallest CIDR: " + cidr)
	return nil
}

func getSmallestCidr(ips []net.IP) string {
	// 找到共同的前缀位数
	ones := getCommonPrefixLength(ips)

	// 创建掩码
	mask := net.CIDRMask(ones, 32)

	// 计算CIDR
	ip := ips[0].Mask(mask)
	ipStr := ip.String()

	// 将掩码位数添加到CIDR
	cidr := ipStr + "/" + strconv.Itoa(ones)

	return cidr
}

func getCommonPrefixLength(ips []net.IP) int {
	minIPInt := ip2Int(ips[0])
	maxIPInt := ip2Int(ips[0])
	for _, ip := range ips {
		ipInt := ip2Int(ip)
		if ipInt < minIPInt {
			minIPInt = ipInt
		}
		if ipInt > maxIPInt {
			maxIPInt = ipInt
		}
	}

	// 异或后，最大一个 1 所在位置之前的位数即为共同前缀位数
	commonPrefix := minIPInt ^ maxIPInt
	ones := 0
	for commonPrefix > 0 {
		commonPrefix >>= 1
		ones++
	}

	// e.g.
	// 0.0.0.7 ^ 0.0.0.0
	// prefix = 00000000 00000000 00000000 00000111
	// ones   = 32 - 3 = 29
	// cidr   = xxx.xxx.xxx.xxx/29

	return 32 - ones
}

func ip2Int(ip net.IP) uint32 {
	ip = ip.To4()
	return (uint32(ip[0]) << 24) |
		(uint32(ip[1]) << 16) |
		(uint32(ip[2]) << 8) |
		uint32(ip[3])
}
