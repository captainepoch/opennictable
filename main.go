// onp -- An OpenNic Project cli
// Copyright (C) 2016  Adolfo "captainepoch" Santiago

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see
// https://www.gnu.org/licenses/gpl-3.0.txt

package main

import (
	"fmt"
	"os"
)

const (
	API          = "https://api.opennicproject.org/geoip/"
	SUCCESS_EXIT = 0
	ERROR_EXIT   = -1
	VERSION      = "0.2"
	BANNER       = "onp ~ Version %s\n"
	LICENSE      = `
onpcli  Copyright (C)  2016  Adolfo "captainepoch" Santiago

This program is licensed under the GNU GPL v3 license. There is a
copy of the license in the LICENSE file attached to this software.
This program comes with ABSOLUTELY NO WARRANTY. This is free software,
and you are welcome to redistribute it under certain conditions which
can be found on the LICENSE file.
`
)

func main() {
	if len(os.Args) > 2 ||
		len(os.Args) == 1 {
		usage()
		os.Exit(ERROR_EXIT)
	}
	switch os.Args[1] {
	case "list":
		List()
		break
	case "bare":
		Bare()
		break
	default:
		usage()
		break
	}
	os.Exit(SUCCESS_EXIT)
}

func usage() {
	fmt.Fprintf(os.Stdout, fmt.Sprintf(BANNER, VERSION))
	fmt.Println("\nUsage:")
	fmt.Println("  list -- Return an ordered list of DNS by uptime")
	fmt.Println("  bare -- Only list the server IP addresses")
	fmt.Println("  help -- Show this message")
	fmt.Print(LICENSE)
}
