// See https://tumregels.github.io/Network-Programming-with-Go/
package net

import (
	"io/ioutil"
	"net"
	"testing"
)

func TestIP(t *testing.T) {
	ipAddrs := []string{
		"127.0.0.1",
		"10.13.128.230",
		"333.222.111.0",

		// IPv6 uses 128-bit addresses. Even bytes becomes cumbersome to express such addresses,
		// so hexadecimal digits are used, grouped into 4 digits and separated by a colon ":".
		// A typical address might be 2002:c0e8:82e7:0:0:0:c0e8:82e7
		// There are tricks to reducing some addresses, such as eliding zeroes and repeated digits.
		// For example, "localhost" is 0:0:0:0:0:0:0:1, which can be shortened to ::1
		"0:0:0:0:0:0:0:1",
		"fe80::1c78:189:e4b9",
		"0:0:0:0:0:0f",
	}

	for _, ip := range ipAddrs {
		addr := net.ParseIP(ip)
		if addr == nil {
			t.Logf("invalid address: %s", ip)
		} else {
			t.Logf("the address is: %s", addr.String())
		}
	}
}

func TestIPMask(t *testing.T) {
	// The IP address of any device is generally composed of two parts:
	// the address of the network in which the device resides,
	// and the address of the device within that network.
	// Given an IP address of a device, and knowing how many bits N are used
	// for the network address gives a relatively straightforward process
	// for extracting the network address and the device address within that network.
	// The netmask for 16 bit network addresses is 255.255.0.0, for 24 bit network addresses
	// it is 255.255.255.0, while for 23 bit addresses it would be 255.255.254.0
	// and for 22 bit addresses it would be 255.255.252.0
	ip := "192.168.0.110"
	addr := net.ParseIP(ip)
	if addr == nil {
		t.Errorf("invalid address: %s", ip)
		return
	}
	mask := addr.DefaultMask()
	ones, bits := mask.Size()
	network := addr.Mask(mask)
	t.Logf("address is %s, default mask length is %d, leading ones count is %d, mask is (hex) %s, network is %s",
		addr.String(), bits, ones, mask.String(), network.String())
}

func TestResolver(t *testing.T) {
	domains := []string{
		"www.google.com",
		"www.baidu.com",
		"archerwq.cn",
	}

	for _, d := range domains {
		addr, err := net.ResolveIPAddr("ip", d)
		if err != nil {
			t.Errorf("failed to resolve %s", d)
			return
		}
		t.Logf("resolve %s got %s", d, addr.String())
	}
}

func TestTCPClient(t *testing.T) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "www.baidu.com:80")
	if err != nil {
		t.Error(err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		t.Error(err)
	}

	if _, err := conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n")); err != nil {
		t.Error(err)
	}

	result, err := ioutil.ReadAll(conn)
	if err != nil {
		t.Error(err)
	}

	t.Logf("HEAD www.baidu.com got:\n %s \n", string(result))
}
