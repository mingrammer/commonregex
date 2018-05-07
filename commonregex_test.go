package commonregex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommonRegex_Date(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	tests := []string{
		"3-23-17",
		"3.23.17",
		"03.23.17",
		"March 23th, 2017",
		"Mar 23th 2017",
		"Mar. 23th, 2017",
		"23 Mar 2017",
	}

	for _, test := range tests {
		parsed := Date(test)
		assert.Equal([]string{test}, parsed, "they should be matched")
	}
}

func TestCommonRegex_Time(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	tests := []string{
		"09:45",
		"9:45",
		"23:45",
		"9:00am",
		"9am",
		"9:00 A.M.",
		"9:00 pm",
	}

	for _, test := range tests {
		parsed := Time(test)
		assert.Equal([]string{test}, parsed, "they should be matched")
	}
}

func TestCommonRegex_Phones(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	tests := []string{
		"12345678900",
		"1234567890",
		"+1 234 567 8900",
		"234-567-8900",
		"1-234-567-8900",
		"1.234.567.8900",
		"5678900",
		"567-8900",
		"(003) 555-1212",
		"+41 22 730 5989",
		"+442345678900",
	}

	for _, test := range tests {
		parsed := Phones(test)
		assert.Equal([]string{test}, parsed, "they should be matched")
	}
}

func TestCommonRegex_PhonesWithExts(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	tests := []string{
		"(523)222-8888 ext 527",
		"(523)222-8888x623",
		"(523)222-8888 x623",
		"(523)222-8888 x 623",
		"(523)222-8888EXT623",
		"523-222-8888EXT623",
		"(523) 222-8888 x 623",
	}

	for _, test := range tests {
		parsed := PhonesWithExts(test)
		assert.Equal([]string{test}, parsed, "they should be matched")
	}
}

func TestCommonRegex_Links(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	tests := []string{
		"www.google.com",
		"http://www.google.com",
		"www.google.com/?query=dog",
		"sub.example.com",
		"http://www.google.com/%&#/?q=dog",
		"google.com",
	}

	for _, test := range tests {
		parsed := Links(test)
		assert.Equal([]string{test}, parsed, "they should be matched")
	}
}

func TestCommonRegex_Emails(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	tests := []string{
		"john.smith@gmail.com",
		"john_smith@gmail.com",
		"john@example.net",
		"John@example.net",
	}

	failingTests := []string{
		"john.smith@gmail..com",
	}

	for _, test := range tests {
		parsed := Emails(test)
		assert.Equal([]string{test}, parsed, "they should be matched")
	}

	for _, test := range failingTests {
		parsed := Emails(test)
		assert.NotEqual([]string{test}, parsed, "they should not be matched")
	}
}

func TestCommonRegex_IPv4s(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	tests := []string{
		"127.0.0.1",
		"192.168.1.1",
		"8.8.8.8",
		"192.30.253.113",
		"216.58.194.46",
	}

	for _, test := range tests {
		parsed := IPs(test)
		assert.Equal([]string{test}, parsed, "they should be matched")
	}
}

func TestCommonRegex_IPv6s(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	tests := []string{
		"fe80:0000:0000:0000:0204:61ff:fe9d:f156",
		"fe80:0:0:0:204:61ff:fe9d:f156",
		"fe80::204:61ff:fe9d:f156",
		"fe80:0000:0000:0000:0204:61ff:254.157.241.86",
		"fe80:0:0:0:0204:61ff:254.157.241.86",
		"::1",
	}

	for _, test := range tests {
		parsed := IPv6s(test)
		assert.Equal([]string{test}, parsed, "they should be matched")
	}
}

func TestCommonRegex_IPs(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	tests := []string{
		"127.0.0.1",
		"192.168.1.1",
		"8.8.8.8",
		"192.30.253.113",
		"216.58.194.46",
		"fe80:0000:0000:0000:0204:61ff:fe9d:f156",
		"fe80:0:0:0:204:61ff:fe9d:f156",
		"fe80::204:61ff:fe9d:f156",
		"fe80:0000:0000:0000:0204:61ff:254.157.241.86",
		"fe80:0:0:0:0204:61ff:254.157.241.86",
		"::1",
	}

	for _, test := range tests {
		parsed := IPs(test)
		assert.Equal([]string{test}, parsed, "they should be matched")
	}
}

func TestCommonRegex_NotKnownPorts(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	tests := []string{
		"1024",
		"2121",
		"8080",
		"12345",
		"55555",
		"65535",
	}

	failingTests := []string{
		"21",
		"80",
		"1023",
		"65536",
	}

	for _, test := range tests {
		parsed := NotKnownPorts(test)
		assert.Equal([]string{test}, parsed, "they should be matched")
	}

	for _, test := range failingTests {
		parsed := Emails(test)
		assert.NotEqual([]string{test}, parsed, "they should not be matched")
	}
}

func TestCommonRegex_Prices(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	tests := []string{
		"$1.23",
		"$1",
		"$1,000",
		"$10,000.00",
	}

	failingTests := []string{
		"$1,10,0",
		"$100.000",
	}

	for _, test := range tests {
		parsed := Prices(test)
		assert.Equal([]string{test}, parsed, "they should be matched")
	}

	for _, test := range failingTests {
		parsed := Prices(test)
		assert.NotEqual([]string{test}, parsed, "they should not be matched")
	}
}

func TestCommonRegex_HexColors(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	tests := []string{
		"#fff",
		"#123",
		"#4e32ff",
	}

	failingTests := []string{
		"#zzz",
	}

	for _, test := range tests {
		parsed := HexColors(test)
		assert.Equal([]string{test}, parsed, "they should be matched")
	}

	for _, test := range failingTests {
		parsed := HexColors(test)
		assert.NotEqual([]string{test}, parsed, "they should not be matched")
	}
}

func TestCommonRegex_CreditCards(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	tests := []string{
		"0000-0000-0000-0000",
		"0123456789012345",
		"0000 0000 0000 0000",
		"012345678901234",
	}

	for _, test := range tests {
		parsed := CreditCards(test)
		assert.Equal([]string{test}, parsed, "they should be matched")
	}
}

func TestCommonRegex_BtcAddresses(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	tests := []string{
		"1LgqButDNV2rVHe9DATt6WqD8tKZEKvaK2",
		"19P6EYhu6kZzRy9Au4wRRZVE8RemrxPbZP",
		"1bones8KbQge9euDn523z5wVhwkTP3uc1",
		"1Bow5EMqtDGV5n5xZVgdpRPJiiDK6XSjiC",
	}

	failingTests := []string{
		"2LgqButDNV2rVHe9DATt6WqD8tKZEKvaK2",
		"19Ry9Au4wRRZVE8RemrxPbZP",
		"1bones8KbQge9euDn523z5wVhwkTP3uc12939",
		"1Bow5EMqtDGV5n5xZVgdpR",
	}

	for _, test := range tests {
		parsed := BtcAddresses(test)
		assert.Equal([]string{test}, parsed, "they should be matched")
	}

	for _, test := range failingTests {
		parsed := BtcAddresses(test)
		assert.NotEqual([]string{test}, parsed, "they should not be matched")
	}
}

func TestCommonRegex_StreetAddresses(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	tests := []string{
		"101 main st.",
		"504 parkwood drive",
		"3 elm boulevard",
		"500 elm street ",
	}

	failingTests := []string{
		"101 main straight",
	}

	for _, test := range tests {
		parsed := StreetAddresses(test)
		assert.Equal([]string{test}, parsed, "they should be matched")
	}

	for _, test := range failingTests {
		parsed := StreetAddresses(test)
		assert.NotEqual([]string{test}, parsed, "they should not be matched")
	}
}

func TestCommonRegex_ZipCodes(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	tests := []string{
		"02540",
		"02540-4119",
	}

	failingTests := []string{
		"101 main straight",
		"123456",
	}

	for _, test := range tests {
		parsed := ZipCodes(test)
		assert.Equal([]string{test}, parsed, "they should be matched")
	}

	for _, test := range failingTests {
		parsed := ZipCodes(test)
		assert.NotEqual([]string{test}, parsed, "they should not be matched")
	}
}

func TestCommonRegex_PoBoxes(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	tests := []string{
		"PO Box 123456",
		"p.o. box 234234",
	}

	failingTests := []string{
		"101 main straight",
	}

	for _, test := range tests {
		parsed := PoBoxes(test)
		assert.Equal([]string{test}, parsed, "they should be matched")
	}

	for _, test := range failingTests {
		parsed := PoBoxes(test)
		assert.NotEqual([]string{test}, parsed, "they should not be matched")
	}
}

func TestCommonRegex_SSNs(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	tests := []string{
		"000-00-0000",
		"111-11-1111",
		"222-22-2222",
		"123-45-6789",
	}

	for _, test := range tests {
		parsed := SSNs(test)
		assert.Equal([]string{test}, parsed, "they should be matched")
	}
}

func TestCommonRegex_MD5Hexes(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	tests := []string{
		"b5ab01fad5a008d436f76aafc896f9c6",
		"00000000000000000000000000000000",
		"fffFFFfFFfFFFfFFFFfFfFfffffFfFFF",
	}

	failingTests := []string{
		"b5ab01fad5a008d436f76aafc896f9c600000000",
		"",
		"7TS5x1trQs652k4AZ3hJE83YCvJRy0U8",
		"b5ab01fad5a008-436f76aafc896f9c6",
	}

	for _, test := range tests {
		parsed := MD5Hexes(test)
		assert.Equal([]string{test}, parsed, "they should be matched")
	}

	for _, test := range failingTests {
		parsed := MD5Hexes(test)
		assert.NotEqual([]string{test}, parsed, "they should not be matched")
	}
}

func TestCommonRegex_SHA1Hexes(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	tests := []string{
		"b5ab01fad5a008d436f76aafc896f9c6abcd1234",
		"0000000000000000000000000000000000000000",
		"fffFFFfFFfFFFfFFFFfFfFfffffFfFFFffffFFFF",
	}

	failingTests := []string{
		"b5ab01fad5a008d436f76aafc896f9c600000000202020202020202020202020",
		"",
		"7TS5x1trQs652k4AZ3hJE83YCvJRy0U85x1trQs652k4AZ3hJE83YCvJRy0U8asd",
		"b5ab01fad5a008-436f76aafc896f9c6-436f76aafc896f9c6-436f76aafc896",
	}

	for _, test := range tests {
		parsed := SHA1Hexes(test)
		assert.Equal([]string{test}, parsed, "they should be matched")
	}

	for _, test := range failingTests {
		parsed := SHA1Hexes(test)
		assert.NotEqual([]string{test}, parsed, "they should not be matched")
	}
}

func TestCommonRegex_SHA256Hexes(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	tests := []string{
		"3f4146a1d0b5dac26562ff7dc6248573f4e996cf764a0f517318ff398dcfa792",
		"0000000000000000000000000000000000000000000000000000000000000000",
		"fffFFFfFFfFFFfFFFFfFfFfffffFfFFFffffFFFFfffffFFFFFffFFffFFffFFff",
	}

	failingTests := []string{
		"3f4146a1d0b5dac26562ff7dc6248573f4e996cf764a0f517318ff398dcfa7920",
		"",
		"e9iLS075z9HAJlUWg2ZpK5hRxjLeSpIqMKJO67c739VYf7Bj7eR1WjOO82IHcXVd",
		"b5ab01fad5a008-436f76aafc896f9c6-436f76aafc896f9c6-436f76aafc896",
	}

	for _, test := range tests {
		parsed := SHA256Hexes(test)
		assert.Equal([]string{test}, parsed, "they should be matched")
	}

	for _, test := range failingTests {
		parsed := SHA256Hexes(test)
		assert.NotEqual([]string{test}, parsed, "they should not be matched")
	}
}

func TestCommonRegex_GUIDs(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	tests := []string{
		"00000000-0000-0000-0000-000000000000",
		"00000000000000000000000000000000",
		"88a310ed-0ac0-4a3d-b3a2-958fa291d061",
		"27143ecab8a440cda6fb6effcf9b3c75",
	}

	failingTests := []string{
		"88a310ed-0ac0_4a3d_b3a2_958fa291d061",
		"88a310ed 0ac0 4a3d b3a2 958fa291d061",
		"",
		"Z8a310ed-0ac0-4a3d-b3a2-958fa291d061",
		"88a310ed-zac0-4a3d-b3a2-958fa291d061",
		"98a310ed-0ac0-za3d-b3a2-958fa291d061",
		"88a310ed-0ac0-4a3d-z3a2-958fa291d061",
		"88a310ed-0ac0-4a3d-b3a2-z58fa291d061",
	}

	for _, test := range tests {
		parsed := GUIDs(test)
		assert.Equal([]string{test}, parsed, "they should be matched")
	}

	for _, test := range failingTests {
		parsed := GUIDs(test)
		assert.NotEqual([]string{test}, parsed, "they should not be matched")
	}
}

func TestCommonRegex_ISBN13s(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	tests := []string{
		"978-3-16-148410-0",
		"978-1-56619-909-4",
		"133-1-12144-909-9",
	}

	failingTests := []string{
		"1-56619-909-3",
		"1-33342-100-1",
		"2-33342-362-9",
	}

	for _, test := range tests {
		parsed := ISBN13s(test)
		assert.Equal([]string{test}, parsed, "they should be matched")
	}

	for _, test := range failingTests {
		parsed := ISBN13s(test)
		assert.NotEqual([]string{test}, parsed, "they should not be matched")
	}
}

func TestCommonRegex_ISBN10s(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	tests := []string{
		"1-56619-909-3",
		"1-33342-100-1",
		"2-33342-362-9",
	}

	failingTests := []string{
		"978-3-16-148410-0",
		"978-1-56619-909-4",
		"133-1-12144-909-9",
	}

	for _, test := range tests {
		parsed := ISBN10s(test)
		assert.Equal([]string{test}, parsed, "they should be matched")
	}

	for _, test := range failingTests {
		parsed := ISBN10s(test)
		assert.NotEqual([]string{test}, parsed, "they should not be matched")
	}
}

func TestCommonRegex_VISACreditCards(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	tests := []string{
		"4111 1111 1111 1111",
		"4222 2222 2222 2222",
	}

	failingTests := []string{
		"5500 0000 0000 0004",
		"3400 0000 0000 009",
		"3000 0000 0000 04",
	}

	for _, test := range tests {
		parsed := VISACreditCards(test)
		assert.Equal([]string{test}, parsed, "they should be matched")
	}

	for _, test := range failingTests {
		parsed := VISACreditCards(test)
		assert.NotEqual([]string{test}, parsed, "they should not be matched")
	}
}

func TestCommonRegex_MCCreditCards(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	tests := []string{
		"5500 0000 0000 0004",
		"5500 3334 0000 1234",
	}

	failingTests := []string{
		"4111 1111 1111 1111",
		"4222 2222 2222 2222",
		"3400 0000 0000 009",
		"3000 0000 0000 04",
	}

	for _, test := range tests {
		parsed := MCCreditCards(test)
		assert.Equal([]string{test}, parsed, "they should be matched")
	}

	for _, test := range failingTests {
		parsed := MCCreditCards(test)
		assert.NotEqual([]string{test}, parsed, "they should not be matched")
	}
}

func TestCommonRegex_MACAddresses(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	tests := []string{
		"f8:2f:a4:fe:76:d2",
		"F8:2F:A4:FE:76:D2",
		"3D-F2-C9-A6-B3-4F",
	}

	failingTests := []string{
		"3D:F2:C9:A6:B3:4G",
		"f0:2f:P4:Be:96:J5",
	}

	for _, test := range tests {
		parsed := MACAddresses(test)
		assert.Equal([]string{test}, parsed, "they should be matched")
	}

	for _, test := range failingTests {
		parsed := MACAddresses(test)
		assert.NotEqual([]string{test}, parsed, "they should not be matched")
	}
}

func TestCommonRegex_IBANs(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	tests := []string{
		"FR1420041010050500013M02606",
		"MU17BOMM0101101030300200000MUR",
		"NO9386011117947",
	}

	failingTests := []string{
		"424220041010050500013M02606",
		"GB29RBOS601613",
	}

	for _, test := range tests {
		parsed := IBANs(test)
		assert.Equal([]string{test}, parsed, "they should be matched")
	}

	for _, test := range failingTests {
		parsed := IBANs(test)
		assert.NotEqual([]string{test}, parsed, "they should not be matched")
	}
}

func TestCommonRegex_GitRepos(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	tests := []string{
		"https://github.com/mingrammer/commonregex.git",
		"git@github.com:mingrammer/commonregex.git",
	}

	failingTests := []string{
		"https://github.com/mingrammer/commonregex",
		"test@github.com:mingrammer/commonregex.git",
	}

	for _, test := range tests {
		parsed := GitRepos(test)
		assert.Equal([]string{test}, parsed, "they should be matched")
	}

	for _, test := range failingTests {
		parsed := GitRepos(test)
		assert.NotEqual([]string{test}, parsed, "they should not be matched")
	}
}
