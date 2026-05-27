package config

// Driver interface.
// TODO refactor: rename GetDecoder() to Decode(), rename GetEncoder() to Encode()
type Driver interface {
	Name() string
	Aliases() []string // alias format names, use for resolve format name
	GetDecoder() Decoder
	GetEncoder() Encoder
}

// DriverV2 interface.
type DriverV2 interface {
	Name() string      // driver name, also is format name.
	Aliases() []string // alias format names, use for resolve format name
	Decode(blob []byte, v any) (err error)
	Encode(v any) (out []byte, err error)
}

// Decoder for decode yml,json,toml format content
type Decoder func(blob []byte, v any) (err error)

// Encoder for decode yml,json,toml format content
type Encoder func(v any) (out []byte, err error)

// StdDriver struct
type StdDriver struct {
	name    string
	aliases []string
	decoder Decoder
	encoder Encoder
}

// NewDriver new std driver instance.
func NewDriver(name string, dec Decoder, enc Encoder) *StdDriver {
	_ = "STUB: not implemented"
	return nil
}

// WithAliases set aliases for driver
func (d *StdDriver) WithAliases(aliases ...string) *StdDriver {
	_ = "STUB: not implemented"
	return nil
}

// WithAlias add alias for driver
func (d *StdDriver) WithAlias(alias string) *StdDriver { _ = "STUB: not implemented"; return nil }

// Name of driver
func (d *StdDriver) Name() string {
	_ = "STUB: not implemented"

	// Aliases format name of driver
	return ""
}

func (d *StdDriver) Aliases() []string {
	_ = "STUB: not implemented"

	// Decode of driver
	return nil
}

func (d *StdDriver) Decode(blob []byte, v any) (err error) { _ = "STUB: not implemented"; return nil }

// Encode of driver
func (d *StdDriver) Encode(v any) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// GetDecoder of driver
		nil
}

func (d *StdDriver) GetDecoder() Decoder {
	_ = "STUB: not implemented"

	// GetEncoder of driver
	return *new(Decoder)
}

func (d *StdDriver) GetEncoder() Encoder {
	_ = "STUB: not implemented"

	/*************************************************************
	 * JSON driver
	 *************************************************************/return *new(Encoder)
}

var (
	// JSONAllowComments support write comments on json file.
	JSONAllowComments = true

	// JSONMarshalIndent if not empty, will use json.MarshalIndent for encode data.
	//
	// Deprecated: please use JSONDriver.MarshalIndent
	JSONMarshalIndent string
)

// JSONDecoder for json decode
var JSONDecoder Decoder = func(data []byte, v any) (err error) {
	JSONDriver.ClearComments = JSONAllowComments
	return JSONDriver.Decode(data, v)
}

// JSONEncoder for json encode
var JSONEncoder Encoder = func(v any) (out []byte, err error) {
	JSONDriver.MarshalIndent = JSONMarshalIndent
	return JSONDriver.Encode(v)
}

// JSONDriver instance fot json
var JSONDriver = &jsonDriver{
	driverName:    JSON,
	ClearComments: JSONAllowComments,
	MarshalIndent: JSONMarshalIndent,
}

// jsonDriver for json format content
type jsonDriver struct {
	driverName string
	// ClearComments before parse JSON string.
	ClearComments bool
	// MarshalIndent if not empty, will use json.MarshalIndent for encode data.
	MarshalIndent string
}

// Name of the driver
func (d *jsonDriver) Name() string { _ = "STUB: not implemented"; return "" }

// Aliases of the driver
func (d *jsonDriver) Aliases() []string {
	_ = "STUB: not implemented"

	// Decode for the driver
	return nil
}

func (d *jsonDriver) Decode(data []byte, v any) error { _ = "STUB: not implemented"; return nil }

// GetDecoder for the driver
func (d *jsonDriver) GetDecoder() Decoder {
	_ = "STUB: not implemented"

	// Encode for the driver
	return *new(Decoder)
}

func (d *jsonDriver) Encode(v any) (out []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetEncoder for the driver
func (d *jsonDriver) GetEncoder() Encoder { _ = "STUB: not implemented"; return *new(Encoder) }
