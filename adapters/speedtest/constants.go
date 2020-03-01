package speedtest

const (
	Help                = "-h"
	Verbosity           = "-v"
	Version             = "-V"
	Servers             = "-L"
	SelectionDetails    = "--selection-details"
	SetServer           = "-s"
	SetHostname         = "-o"
	SetFormat           = "-f"
	HumanReadableFormat = "human-readable"
	CsvFormat           = "csv"
	TsvFormat           = "tsv"
	JsonFormat          = "json"
	JsonLinesFormat     = "jsonl"
	JsonPrettyFormat    = "json-pretty"
	OutputHeader        = "--output-header"
	UnitOfMeasure       = "-u"

	// Decimal prefix, bits per second:
	UnitBps  = "bps"
	UnitkBps = "kbps"
	UnitMBps = "Mbps"
	UnitGBps = "Gbps"
	// Decimal prefix, bytes per second:
	UnitBs  = "B/s"
	UnitkBs = "kB/s"
	UnitMBs = "MB/s"
	UnitGBs = "GB/s"
	// Binary prefix, bits per second:
	Unitkibps = "kibps"
	UnitMibps = "Mibps"
	UnitGibps = "Gibps"
	// Binary prefix, bytes per second:
	UnitkiBs = "kiB/s"
	UnitMiBs = "MiB/s"
	UnitGiBs = "GiB/s"
	// Auto-scaled prefix:
	UnitBinaryBits   = "-b"
	UnitBinaryBytes  = "-B"
	UnitDecimalBits  = "-a"
	UnitDecimalBytes = "-A"

	decimalPlaces      = "-P"
	progress           = "-p"
	specifiedInterface = "-I"
	ipAddress          = "-i"
	caCertificate      = "--ca-certificate="
)
