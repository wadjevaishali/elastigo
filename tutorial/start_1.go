// Copyright 2013 Matthew Baird
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	elastigo "github.com/wadjevaishali/elastigo/lib"
)

var (
	host *string = flag.String("host", "localhost", "Elasticsearch Host")
)

func main() {
	c := elastigo.NewConn()
	log.SetFlags(log.LstdFlags)
	flag.Parse()

	fmt.Println("host = ", *host)
	// Set the Elasticsearch Host to Connect to
	c.Domain = *host

	// Index a document
	_, err := c.Index("testindex", "user", "docid_1", nil, `{"name":"bob martin"}`)
	exitIfErr(err)

	// Index a doc using a map of values
	_, err = c.Index("testindex", "user", "docid_2", nil, map[string]string{"name": "bob mosbey"})
	exitIfErr(err)

	// Index a doc using Structs
	_, err = c.Index("testindex", "user", "docid_3", nil, MyUser{"vaishali", 22})
	exitIfErr(err)

	// Search Using Raw json String
	searchJson := `{
	    "query" : {
				"match": {
					 "name" : "bob"
				}
	    }
	}`
	out, err := c.Search("testindex", "user", nil, searchJson)

	result, _ := json.MarshalIndent(out, "", "  ")
	fmt.Printf("%s\n", result)
	exitIfErr(err)

}
func exitIfErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
}

type MyUser struct {
	Name string
	Age  int
}
