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
	"bytes"
	tabw "github.com/olekukonko/tablewriter"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
)

type onpData struct {
	ip     string
	dns    string
	uptime string
}

type iOnpData []onpData

func (this iOnpData) Len() int {
	return len(this)
}
func (this iOnpData) Less(i, j int) bool {
	return this[i].uptime < this[j].uptime
}
func (this iOnpData) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func callAPI(api string) string {
	resp, err := http.Get(api)
	if err != nil {
		log.Fatal(err)
		os.Exit(ERROR_EXIT)
	}
	defer resp.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return buf.String()
}

func List() {
	buf := callAPI(API + "?list")
	fieldsData := strings.Fields(buf)
	data := []onpData{}
	for i := 0; i < len(fieldsData); i += 5 {
		data = append(data, onpData{fieldsData[0+i], fieldsData[2+i], fieldsData[4+i]})
	}
	sort.Sort(sort.Reverse(iOnpData(data)))
	table := tabw.NewWriter(os.Stdout)
	table.SetHeader([]string{"IP", "DNS", "UPTIME (%)"})
	for _, data := range data {
		table.Append([]string{data.ip, data.dns, data.uptime})
	}
	table.Render()
}

func Bare() {
	buf := callAPI(API + "?bare")
	fieldsData := strings.Fields(buf)
	table := tabw.NewWriter(os.Stdout)
	table.SetHeader([]string{"IP"})
	for _, data := range fieldsData {
		table.Append([]string{data})
	}
	table.Render()
}
