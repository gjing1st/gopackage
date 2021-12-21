package gpip

import (
	"net"
	"strconv"
	"strings"
)


// IpToInt
// @description: ip转int
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/12/14 11:07
// @success:
func IpToInt(ip string) int64{
	ipArr := strings.Split(ip,".")
	_ = ipArr[3]
	ipInt0,_ := strconv.Atoi(ipArr[0])
	ipInt1,_ := strconv.Atoi(ipArr[1])
	ipInt2,_ := strconv.Atoi(ipArr[2])
	ipInt3,_ := strconv.Atoi(ipArr[3])
	return int64(ipInt0) << 24 + int64(ipInt1) << 16 + int64(ipInt2) << 8 +int64(ipInt3)
}

func IpToBytes(ip string) []byte{
	ipArr := strings.Split(ip,".")
	_ = ipArr[3]
	p := make([]byte,4)
	ipInt0,_ := strconv.Atoi(ipArr[0])
	ipInt1,_ := strconv.Atoi(ipArr[1])
	ipInt2,_ := strconv.Atoi(ipArr[2])
	ipInt3,_ := strconv.Atoi(ipArr[3])
	p[0] = byte(ipInt0)
	p[1] = byte(ipInt1)
	p[2] = byte(ipInt2)
	p[3] = byte(ipInt3)
	return p
}

func Int64ToIP(intIP int64) net.IP {
	var bytes [4]byte
	bytes[0] = byte(intIP & 0xFF)
	bytes[1] = byte((intIP >> 8) & 0xFF)
	bytes[2] = byte((intIP >> 16) & 0xFF)
	bytes[3] = byte((intIP >> 24) & 0xFF)

	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
}