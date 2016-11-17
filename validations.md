# Gencheck validations

Name | Params | Field Type | Description
---- | ------------------- | ------ | -----------
hex  | N/A | `(*)string` | Checks if a string field is a valid hexadecimal format number (0x prefix optional)
notnil | N/A | pointers, interfaces, chans, maps, slices | Checks and fails if a pointer is nil
len | 1 | `(*)string, slices, maps, chans(?)` | Uses golang's built in `len` method to determine the length of the field
len | 1 | `int` | Checks if the value equals the parameter
uuid | N/A | `(*)string` | Checks and fails if a string is not a valid UUID
required | N/A | Any Nullable Value (ptr, `map`, `slice`, `chan`) | Checks to see if the pointer value is nil
required | N/A | `string` | Checks to see if the `string != ""`
required | N/A | `struct` | Checks to see if the struct equals the empty value struct
min | 1 | `(*)string, slices, maps, chans` | Checks minimum length using `len(value)`
min | 1 | `numbers` | Checks a simple `value >= param`
max | 1 | `(*)string, slices, maps, chans` | Checks minimum length using `len(value)`
max | 1 | `numbers` | Checks a simple `value <= param`
lt(e) | 1 | See max | See max
lt(e) | 1 | `time` | Checks that the field is less than (or equal to) `time.Now().UTC()`
gt(e) | 1 | See min | See min
gt(e) | 1 | `time` | Checks that the field is greater than (or equal to) `time.Now().UTC()`
dive  | N/A | Any | Dive will call `gencheck.Validate()` on the struct/ptr/interface passed in.  That function will only call validate if it can be cast to the `gencheck.Validateable` interface.
contains  | 1 | `(*)string` | Checks to see if the field contains the parameter string value
contains | 1 | `[](*)string` | Checks to see if any of the array elements **equal** the parameter value
bcp47 | N/A | `(*)string` | Checks to see if the string value is one of the languages defined in https://tools.ietf.org/html/bcp47
ff  | N/A | Any | Fail Fast allows you to specify that if any validation within this field is invalid see [Fail Fast](#Fail Fast Flag) for more information
