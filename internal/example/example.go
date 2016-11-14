//go:generate gencheck -f=example.go

package example

import "time"

// Example struct for testing
type Example struct {
	MapOfInterfaces map[string]interface{} `valid:"notnil"`
}

// Inner is another example struct for testing
type Inner struct {
	EqCSFieldString  string
	NeCSFieldString  string
	GtCSFieldString  string
	GteCSFieldString string
	LtCSFieldString  string
	LteCSFieldString string
}

// Test is a test struct for testing.
type Test struct {
	Inner Inner
	// RequiredNumber    int       `valid:"required"`
	RequiredString    string     `valid:"required,ff"`
	RequiredMultiple  []string   `valid:"required"`
	LenString         string     `valid:"len=1"`
	LenNumber         float64    `valid:"len=1113.00"`
	LenMultiple       []string   `valid:"len=7"`
	MinString         string     `valid:"min=1"`
	MinNumber         float64    `valid:"min=1113.00"`
	MinMultiple       []string   `valid:"min=7"`
	MaxString         string     `valid:"max=3"`
	MaxNumber         float64    `valid:"max=1113.00"`
	MaxMultiple       []string   `valid:"max=7"`
	EqString          string     `valid:"eq=3"`
	EqNumber          float64    `valid:"eq=2.33"`
	EqMultiple        []string   `valid:"eq=7"`
	NeString          string     `valid:"ne="`
	NeNumber          float64    `valid:"ne=0.00"`
	NeMultiple        []string   `valid:"ne=0"`
	LtString          string     `valid:"lt=3"`
	LtNumber          float64    `valid:"lt=5.56"`
	LtMultiple        []string   `valid:"lt=2"`
	LtTime            time.Time  `valid:"lt"`
	LteString         string     `valid:"lte=3"`
	LteNumber         float64    `valid:"lte=5.56"`
	LteMultiple       []string   `valid:"lte=2"`
	LteTime           time.Time  `valid:"lte"`
	GtString          string     `valid:"gt=3"`
	GtNumber          float64    `valid:"gt=5.56"`
	GtMultiple        []string   `valid:"gt=2"`
	GtTime            time.Time  `valid:"gt"`
	GteString         string     `valid:"gte=3"`
	GteNumber         float64    `valid:"gte=5.56"`
	GteMultiple       []string   `valid:"gte=2"`
	GteTime           time.Time  `valid:"gte"`
	GteTimeVal        time.Time  `valid:"gte=1*time.Second"`
	GteTimePtr        *time.Time `valid:"gte"`
	EqFieldString     string     `valid:"eqfield=MaxString"`
	EqCSFieldString   string     `valid:"eqcsfield=Inner.EqCSFieldString"`
	NeCSFieldString   string     `valid:"necsfield=Inner.NeCSFieldString"`
	GtCSFieldString   string     `valid:"gtcsfield=Inner.GtCSFieldString"`
	GteCSFieldString  string     `valid:"gtecsfield=Inner.GteCSFieldString"`
	LtCSFieldString   string     `valid:"ltcsfield=Inner.LtCSFieldString"`
	LteCSFieldString  string     `valid:"ltecsfield=Inner.LteCSFieldString"`
	NeFieldString     string     `valid:"nefield=EqFieldString"`
	GtFieldString     string     `valid:"gtfield=MaxString"`
	GteFieldString    string     `valid:"gtefield=MaxString"`
	LtFieldString     string     `valid:"ltfield=MaxString"`
	LteFieldString    string     `valid:"ltefield=MaxString"`
	AlphaString       string     `valid:"alpha"`
	AlphanumString    string     `valid:"alphanum"`
	NumericString     string     `valid:"numeric"`
	NumberString      string     `valid:"number"`
	HexadecimalString string     `valid:"hexadecimal"`
	HexColorString    string     `valid:"hexcolor"`
	RGBColorString    string     `valid:"rgb"`
	RGBAColorString   string     `valid:"rgba"`
	HSLColorString    string     `valid:"hsl"`
	HSLAColorString   string     `valid:"hsla"`
	Email             string     `valid:"email"`
	URL               string     `valid:"url"`
	URI               string     `valid:"uri"`
	Base64            string     `valid:"base64"`
	Contains          string     `valid:"contains=purpose"`
	ContainsAny       string     `valid:"containsany=!@#$"`
	Excludes          string     `valid:"excludes=text"`
	ExcludesAll       string     `valid:"excludesall=!@#$"`
	ExcludesRune      string     `valid:"excludesrune=☻"`
	ISBN              string     `valid:"isbn"`
	ISBN10            string     `valid:"isbn10"`
	ISBN13            string     `valid:"isbn13"`
	UUID              string     `valid:"uuid"`
	UUID3             string     `valid:"uuid3"`
	UUID4             string     `valid:"uuid4"`
	UUID5             string     `valid:"uuid5"`
	ASCII             string     `valid:"ascii"`
	PrintableASCII    string     `valid:"printascii"`
	MultiByte         string     `valid:"multibyte"`
	DataURI           string     `valid:"datauri"`
	Latitude          string     `valid:"latitude"`
	Longitude         string     `valid:"longitude"`
	SSN               string     `valid:"ssn"`
	IP                string     `valid:"ip"`
	IPv4              string     `valid:"ipv4"`
	IPv6              string     `valid:"ipv6"`
	CIDR              string     `valid:"cidr"`
	CIDRv4            string     `valid:"cidrv4"`
	CIDRv6            string     `valid:"cidrv6"`
	TCPAddr           string     `valid:"tcp_addr"`
	TCPAddrv4         string     `valid:"tcp4_addr"`
	TCPAddrv6         string     `valid:"tcp6_addr"`
	UDPAddr           string     `valid:"udp_addr"`
	UDPAddrv4         string     `valid:"udp4_addr"`
	UDPAddrv6         string     `valid:"udp6_addr"`
	IPAddr            string     `valid:"ip_addr"`
	IPAddrv4          string     `valid:"ip4_addr"`
	IPAddrv6          string     `valid:"ip6_addr"`
	UnixAddr          string     `valid:"unix_addr"` // can't fail from within Go's net package currently, but maybe in the future
	MAC               string     `valid:"mac"`
	IsColor           string     `valid:"iscolor"`
	MinIntPtr         *int64     `valid:"required,min=1234"`
}
