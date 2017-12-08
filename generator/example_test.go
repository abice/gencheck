package generator

import (
	"fmt"
	"time"

	"github.com/abice/gencheck"
)

type NoImplicitvalid struct {
	StringField string

	NonStringKeyMap             map[int]Testvalid
	NestedNonStringKeyMap       map[string]map[int]map[string]*Testvalid
	NestedMapWithNonStructValue map[string]map[string]map[string]*int64

	SliceOfBuiltins       []string
	NestedSliceofBuiltins [][][]*int
}

type HasvalidImplicit struct {
	InvalidStruct *Testvalid
	ValidStruct   AlwaysValid

	MapOfStruct     map[string]Testvalid
	MapOfStructPtrs map[string]*Testvalid
	MapOfMaps       map[string]map[string]*Testvalid
	MapMapsOfSlices map[string]map[string][]*Testvalid
	MapOfInterfaces map[string]interface{}

	SimpleSlice           []*Testvalid
	SliceOfSlicesOfSlices [][][]*Testvalid

	MapOfSlicesOfMaps map[string][]map[string]*Testvalid
}

// Testvalid
type Testvalid struct {
	Valid bool
}

// valid
func (s Testvalid) valid() error {
	if !s.Valid {
		return gencheck.ValidationErrors{gencheck.NewFieldError("TestValid", "Valid", "", fmt.Errorf("invalid when false"))}
	}
	return nil
}

// AlwaysValid
type AlwaysValid struct{}

// valid
func (s AlwaysValid) valid() error {
	return nil
}

// Example
type MapContains struct {
	MapOfInterfaces map[string]interface{} `valid:"contains=myKey"`
}

// Example
type Example struct {
	MapOfInterfaces map[string]interface{} `valid:"notnil"`
}

type Inner struct {
	EqCSFieldString  string
	NeCSFieldString  string
	GtCSFieldString  string
	GteCSFieldString string
	LtCSFieldString  string
	LteCSFieldString string
}

type Embedded struct {
	FieldString string `valid:"required"`
}

type Test struct {
	Embedded          `valid:"dive"`
	Inner             Inner
	RequiredString    string    `validate:"required"`
	RequiredNumber    int       `validate:"required"`
	RequiredMultiple  []string  `validate:"required"`
	LenString         string    `validate:"len=1"`
	LenNumber         float64   `validate:"len=1113.00"`
	LenMultiple       []string  `validate:"len=7"`
	MinString         string    `validate:"min=1"`
	MinNumber         float64   `validate:"min=1113.00"`
	MinMultiple       []string  `validate:"min=7"`
	MaxString         string    `validate:"max=3"`
	MaxNumber         float64   `validate:"max=1113.00"`
	MaxMultiple       []string  `validate:"max=7"`
	EqString          string    `validate:"eq=3"`
	EqNumber          float64   `validate:"eq=2.33"`
	EqMultiple        []string  `validate:"eq=7"`
	NeString          string    `validate:"ne="`
	NeNumber          float64   `validate:"ne=0.00"`
	NeMultiple        []string  `validate:"ne=0"`
	LtString          string    `validate:"lt=3"`
	LtNumber          float64   `validate:"lt=5.56"`
	LtMultiple        []string  `validate:"lt=2"`
	LtTime            time.Time `validate:"lt"`
	LteString         string    `validate:"lte=3"`
	LteNumber         float64   `validate:"lte=5.56"`
	LteMultiple       []string  `validate:"lte=2"`
	LteTime           time.Time `validate:"lte"`
	GtString          string    `validate:"gt=3"`
	GtNumber          float64   `validate:"gt=5.56"`
	GtMultiple        []string  `validate:"gt=2"`
	GtTime            time.Time `validate:"gt"`
	GteString         string    `validate:"gte=3"`
	GteNumber         float64   `validate:"gte=5.56"`
	GteMultiple       []string  `validate:"gte=2"`
	GteTime           time.Time `validate:"gte"`
	EqFieldString     string    `validate:"eqfield=MaxString"`
	EqCSFieldString   string    `validate:"eqcsfield=Inner.EqCSFieldString"`
	NeCSFieldString   string    `validate:"necsfield=Inner.NeCSFieldString"`
	GtCSFieldString   string    `validate:"gtcsfield=Inner.GtCSFieldString"`
	GteCSFieldString  string    `validate:"gtecsfield=Inner.GteCSFieldString"`
	LtCSFieldString   string    `validate:"ltcsfield=Inner.LtCSFieldString"`
	LteCSFieldString  string    `validate:"ltecsfield=Inner.LteCSFieldString"`
	NeFieldString     string    `validate:"nefield=EqFieldString"`
	GtFieldString     string    `validate:"gtfield=MaxString"`
	GteFieldString    string    `validate:"gtefield=MaxString"`
	LtFieldString     string    `validate:"ltfield=MaxString"`
	LteFieldString    string    `validate:"ltefield=MaxString"`
	AlphaString       string    `validate:"alpha"`
	AlphanumString    string    `validate:"alphanum"`
	NumericString     string    `validate:"numeric"`
	NumberString      string    `validate:"number"`
	HexadecimalString string    `validate:"hexadecimal"`
	HexColorString    string    `validate:"hexcolor"`
	RGBColorString    string    `validate:"rgb"`
	RGBAColorString   string    `validate:"rgba"`
	HSLColorString    string    `validate:"hsl"`
	HSLAColorString   string    `validate:"hsla"`
	Email             string    `validate:"email"`
	URL               string    `validate:"url"`
	URI               string    `validate:"uri"`
	Base64            string    `validate:"base64"`
	Contains          string    `validate:"contains=purpose"`
	ContainsAny       string    `validate:"containsany=!@#$"`
	Excludes          string    `validate:"excludes=text"`
	ExcludesAll       string    `validate:"excludesall=!@#$"`
	ExcludesRune      string    `validate:"excludesrune=☻"`
	ISBN              string    `validate:"isbn"`
	ISBN10            string    `validate:"isbn10"`
	ISBN13            string    `validate:"isbn13"`
	UUID              string    `validate:"uuid"`
	UUID3             string    `validate:"uuid3"`
	UUID4             string    `validate:"uuid4"`
	UUID5             string    `validate:"uuid5"`
	ASCII             string    `validate:"ascii"`
	PrintableASCII    string    `validate:"printascii"`
	MultiByte         string    `validate:"multibyte"`
	DataURI           string    `validate:"datauri"`
	Latitude          string    `validate:"latitude"`
	Longitude         string    `validate:"longitude"`
	SSN               string    `validate:"ssn"`
	IP                string    `validate:"ip"`
	IPv4              string    `validate:"ipv4"`
	IPv6              string    `validate:"ipv6"`
	CIDR              string    `validate:"cidr"`
	CIDRv4            string    `validate:"cidrv4"`
	CIDRv6            string    `validate:"cidrv6"`
	TCPAddr           string    `validate:"tcp_addr"`
	TCPAddrv4         string    `validate:"tcp4_addr"`
	TCPAddrv6         string    `validate:"tcp6_addr"`
	UDPAddr           string    `validate:"udp_addr"`
	UDPAddrv4         string    `validate:"udp4_addr"`
	UDPAddrv6         string    `validate:"udp6_addr"`
	IPAddr            string    `validate:"ip_addr"`
	IPAddrv4          string    `validate:"ip4_addr"`
	IPAddrv6          string    `validate:"ip6_addr"`
	UnixAddr          string    `validate:"unix_addr"` // can't fail from within Go's net package currently, but maybe in the future
	MAC               string    `validate:"mac"`
	IsColor           string    `validate:"iscolor"`
}
