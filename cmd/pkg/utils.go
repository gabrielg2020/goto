package pkg

// Returns the usage message
func Usage() string {
	return `Usage: goto [options] <pattern>
Options:
	-d <depth> Specify the maximum depth to search (default: 5)
	-h         Display this help message
`
}
