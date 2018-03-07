// Package flaggy is a input flag parsing tool that supports
// subcommands and any-position flags without complexeties.
//
// Parsing Order:
//   - Parse and asign any flags found in the format -key=var,
//     --key=var, or '-key var'.  Remove these variables from
//     further consideration.
//   - Detect any positional values
//   - Detect any subcommands and parse them
//   - Repeat parsing order on subcommands until out of subcommands
//
package flaggy

import "fmt"

// DebugMode indicates that debug output should be enabled
var DebugMode bool

// var mainArgumentParser *ArgumentParser

func init() {
	// TODO - allow usage like flaggy.StringVar or require usage of
	// NewArgumentParser for all usage?
	// mainArgumentParser = NewArgumentParser()?
}

// Pase parses flags as requested
// func Parse() {
// 	mainArgumentParser.Parse()
// }

// debugPrint prints if debugging is enabled
func debugPrint(i ...interface{}) {
	if DebugMode {
		fmt.Println(i...)
	}
}