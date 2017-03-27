package commonregex

import (
	"reflect"
	"testing"
)

func TestCommonRegex_Date(t *testing.T) {
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
		if reflect.DeepEqual(parsed, []string{test}) == false {
			t.Errorf("%s is not matched with %s", parsed, []string{test})
		}
	}
}

func TestCommonRegex_Time(t *testing.T) {
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
		if reflect.DeepEqual(parsed, []string{test}) == false {
			t.Errorf("%s is not matched with %s", parsed, []string{test})
		}
	}
}

func TestCommonRegex_Phones(t *testing.T) {
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
		if reflect.DeepEqual(parsed, []string{test}) == false {
			t.Errorf("%s is not matched with %s", parsed, []string{test})
		}
	}
}

func TestCommonRegex_PhonesWithExts(t *testing.T) {
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
		if reflect.DeepEqual(parsed, []string{test}) == false {
			t.Errorf("%s is not matched with %s", parsed, []string{test})
		}
	}
}

func TestCommonRegex_Links(t *testing.T) {
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
		if reflect.DeepEqual(parsed, []string{test}) == false {
			t.Errorf("%s is not matched with %s", parsed, []string{test})
		}
	}
}

func TestCommonRegex_Emails(t *testing.T) {
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
		if reflect.DeepEqual(parsed, []string{test}) == false {
			t.Errorf("%s is not matched with %s", parsed, []string{test})
		}
	}
	for _, test := range failingTests {
		parsed := Emails(test)
		if reflect.DeepEqual(parsed, []string{test}) {
			t.Errorf("%s should not be equal with %s", parsed, []string{test})
		}
	}
}

func TestCommonRegex_IPs(t *testing.T) {
	tests := []string{
		"127.0.0.1",
		"192.168.1.1",
		"8.8.8.8",
	}
	for _, test := range tests {
		parsed := IPs(test)
		if reflect.DeepEqual(parsed, []string{test}) == false {
			t.Errorf("%s is not matched with %s", parsed, []string{test})
		}
	}
}

func TestCommonRegex_IPv6s(t *testing.T) {
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
		if reflect.DeepEqual(parsed, []string{test}) == false {
			t.Errorf("%s is not matched with %s", parsed, []string{test})
		}
	}
}

func TestCommonRegex_Prices(t *testing.T) {
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
		if reflect.DeepEqual(parsed, []string{test}) == false {
			t.Errorf("%s is not matched with %s", parsed, []string{test})
		}
	}
	for _, test := range failingTests {
		parsed := Prices(test)
		if reflect.DeepEqual(parsed, []string{test}) {
			t.Errorf("%s should not be equal with %s", parsed, []string{test})
		}
	}
}

func TestCommonRegex_HexColors(t *testing.T) {
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
		if reflect.DeepEqual(parsed, []string{test}) == false {
			t.Errorf("%s is not matched with %s", parsed, []string{test})
		}
	}
	for _, test := range failingTests {
		parsed := HexColors(test)
		if reflect.DeepEqual(parsed, []string{test}) {
			t.Errorf("%s should not be equal with %s", parsed, []string{test})
		}
	}
}

func TestCommonRegex_CreditCards(t *testing.T) {
	tests := []string{
		"0000-0000-0000-0000",
		"0123456789012345",
		"0000 0000 0000 0000",
		"012345678901234",
	}
	for _, test := range tests {
		parsed := CreditCards(test)
		if reflect.DeepEqual(parsed, []string{test}) == false {
			t.Errorf("%s is not matched with %s", parsed, []string{test})
		}
	}
}

func TestCommonRegex_BtcAddresses(t *testing.T) {
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
		if reflect.DeepEqual(parsed, []string{test}) == false {
			t.Errorf("%s is not matched with %s", parsed, []string{test})
		}
	}
	for _, test := range failingTests {
		parsed := BtcAddresses(test)
		if reflect.DeepEqual(parsed, []string{test}) {
			t.Errorf("%s should not be equal with %s", parsed, []string{test})
		}
	}
}

func TestCommonRegex_StreetAddresses(t *testing.T) {
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
		if reflect.DeepEqual(parsed, []string{test}) == false {
			t.Errorf("%s is not matched with %s", parsed, []string{test})
		}
	}
	for _, test := range failingTests {
		parsed := StreetAddresses(test)
		if reflect.DeepEqual(parsed, []string{test}) {
			t.Errorf("%s should not be equal with %s", parsed, []string{test})
		}
	}
}

func TestCommonRegex_ZipCodes(t *testing.T) {
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
		if reflect.DeepEqual(parsed, []string{test}) == false {
			t.Errorf("%s is not matched with %s", parsed, []string{test})
		}
	}
	for _, test := range failingTests {
		parsed := ZipCodes(test)
		if reflect.DeepEqual(parsed, []string{test}) {
			t.Errorf("%s should not be equal with %s", parsed, []string{test})
		}
	}
}

func TestCommonRegex_PoBoxes(t *testing.T) {
	tests := []string{
		"PO Box 123456",
		"p.o. box 234234",
	}
	failingTests := []string{
		"101 main straight",
	}
	for _, test := range tests {
		parsed := PoBoxes(test)
		if reflect.DeepEqual(parsed, []string{test}) == false {
			t.Errorf("%s is not matched with %s", parsed, []string{test})
		}
	}
	for _, test := range failingTests {
		parsed := PoBoxes(test)
		if reflect.DeepEqual(parsed, []string{test}) {
			t.Errorf("%s should not be equal with %s", parsed, []string{test})
		}
	}
}

func TestCommonRegex_SSNs(t *testing.T) {
	tests := []string{
		"000-00-0000",
		"111-11-1111",
		"222-22-2222",
		"123-45-6789",
	}
	for _, test := range tests {
		parsed := SSNs(test)
		if reflect.DeepEqual(parsed, []string{test}) == false {
			t.Errorf("%s is not matched with %s", parsed, []string{test})
		}
	}
}

func TestCommonRegex_MD5Hexes(t *testing.T) {
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
		if reflect.DeepEqual(parsed, []string{test}) == false {
			t.Errorf("%s is not matched with %s", parsed, []string{test})
		}
	}
	for _, test := range failingTests {
		parsed := MD5Hexes(test)
		if reflect.DeepEqual(parsed, []string{test}) {
			t.Errorf("%s should not be equal with %s", parsed, []string{test})
		}
	}
}

func TestCommonRegex_SHA1Hexes(t *testing.T) {
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
		if reflect.DeepEqual(parsed, []string{test}) == false {
			t.Errorf("%s is not matched with %s", parsed, []string{test})
		}
	}
	for _, test := range failingTests {
		parsed := SHA1Hexes(test)
		if reflect.DeepEqual(parsed, []string{test}) {
			t.Errorf("%s should not be equal with %s", parsed, []string{test})
		}
	}
}

func TestCommonRegex_SHA256Hexes(t *testing.T) {
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
		if reflect.DeepEqual(parsed, []string{test}) == false {
			t.Errorf("%s is not matched with %s", parsed, []string{test})
		}
	}
	for _, test := range failingTests {
		parsed := SHA256Hexes(test)
		if reflect.DeepEqual(parsed, []string{test}) {
			t.Errorf("%s should not be equal with %s", parsed, []string{test})
		}
	}
}

func TestCommonRegex_GUIDs(t *testing.T) {
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
		if reflect.DeepEqual(parsed, []string{test}) == false {
			t.Errorf("%s is not matched with %s", parsed, []string{test})
		}
	}
	for _, test := range failingTests {
		parsed := GUIDs(test)
		if reflect.DeepEqual(parsed, []string{test}) {
			t.Errorf("%s should not be equal with %s", parsed, []string{test})
		}
	}
}

func TestCommonRegex_ISBN13s(t *testing.T) {
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
		if reflect.DeepEqual(parsed, []string{test}) == false {
			t.Errorf("%s is not matched with %s", parsed, []string{test})
		}
	}
	for _, test := range failingTests {
		parsed := ISBN13s(test)
		if reflect.DeepEqual(parsed, []string{test}) {
			t.Errorf("%s should not be equal with %s", parsed, []string{test})
		}
	}
}

func TestCommonRegex_ISBN10s(t *testing.T) {
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
		if reflect.DeepEqual(parsed, []string{test}) == false {
			t.Errorf("%s is not matched with %s", parsed, []string{test})
		}
	}
	for _, test := range failingTests {
		parsed := ISBN10s(test)
		if reflect.DeepEqual(parsed, []string{test}) {
			t.Errorf("%s should not be equal with %s", parsed, []string{test})
		}
	}
}

func TestCommonRegex_VISACreditCards(t *testing.T) {
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
		if reflect.DeepEqual(parsed, []string{test}) == false {
			t.Errorf("%s is not matched with %s", parsed, []string{test})
		}
	}
	for _, test := range failingTests {
		parsed := VISACreditCards(test)
		if reflect.DeepEqual(parsed, []string{test}) {
			t.Errorf("%s should not be equal with %s", parsed, []string{test})
		}
	}
}

func TestCommonRegex_MCCreditCards(t *testing.T) {
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
		if reflect.DeepEqual(parsed, []string{test}) == false {
			t.Errorf("%s is not matched with %s", parsed, []string{test})
		}
	}
	for _, test := range failingTests {
		parsed := MCCreditCards(test)
		if reflect.DeepEqual(parsed, []string{test}) {
			t.Errorf("%s should not be equal with %s", parsed, []string{test})
		}
	}
}

func TestCommonRegex_MACAddresses(t *testing.T) {
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
		if reflect.DeepEqual(parsed, []string{test}) == false {
			t.Errorf("%s is not matched with %s", parsed, []string{test})
		}
	}
	for _, test := range failingTests {
		parsed := MACAddresses(test)
		if reflect.DeepEqual(parsed, []string{test}) {
			t.Errorf("%s should not be equal with %s", parsed, []string{test})
		}
	}
}
