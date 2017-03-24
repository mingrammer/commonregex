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
