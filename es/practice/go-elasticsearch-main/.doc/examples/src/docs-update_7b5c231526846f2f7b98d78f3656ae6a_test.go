// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
// Code generated, DO NOT EDIT

package elasticsearch_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	_ = fmt.Printf
	_ = os.Stdout
	_ = elasticsearch.NewDefaultClient
)

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/update.asciidoc#L327>
//
// --------------------------------------------------------------------------------
// POST test/_update/1
// {
//   "doc": {
//     "name": "new_name"
//   },
//   "doc_as_upsert": true
// }
// --------------------------------------------------------------------------------

func Test_docs_update_7b5c231526846f2f7b98d78f3656ae6a(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:7b5c231526846f2f7b98d78f3656ae6a[]
	res, err := es.Update(
		"test",
		"1",
		strings.NewReader(`{
		  "doc": {
		    "name": "new_name"
		  },
		  "doc_as_upsert": true
		}`),
		es.Update.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:7b5c231526846f2f7b98d78f3656ae6a[]
}
