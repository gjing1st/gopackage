package gpip

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strconv"
	"testing"
)

func TestIpToInt(t *testing.T) {

	privateKey := "a6a00c446a369a1f853b99fe7f6047c68924131f23887a0d5c9f021c9b9987e541f18b7c1e73e913661be238dead9ebddfc55992d912e232238c8e4ab577b504"
	publicKey := "73005218175930470277091168483563460203371166248613711116440644413998314621091"
	fmt.Println(len(privateKey))
	fmt.Println(len(publicKey))



	ip := IpToInt("192.168.0.127")
	ipByte := IpToBytes("192.168.0.127")
	fmt.Println("ip=parseIPv4",ipByte)
	ipBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(ipBytes, uint64(ip))
	fmt.Println("ipBytes=", ipBytes)
	ipHex := hex.EncodeToString(ipBytes)
	fmt.Println("ipHex=", ipHex)
	ip32 := make([]byte, 32)
	copy(ip32[32-len(ipBytes):], ipBytes)
	fmt.Println("ip32=", ip32)
	ip32Hex := hex.EncodeToString(ip32)
	fmt.Println("ip32Hex=",ip32Hex)

	s := strconv.FormatInt(ip, 2)
	fmt.Println(s)
	b := []byte(s)
	fmt.Println(b)
	fmt.Println(len(b))



	a := "3131303030303030313031303130303030303030303030303031313131313131"
	ab, _ := hex.DecodeString(a)
	fmt.Println("=====", string(ab))
	aI, _ := strconv.ParseInt(string(ab), 2, 64)
	fmt.Println("aint", aI)


}
