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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/update.asciidoc#L116>
//
// --------------------------------------------------------------------------------
// POST test/_update/1
// {
//   "script": {
//     "source": "ctx._source.tags.add(params.tag)",
//     "lang": "painless",
//     "params": {
//       "tag": "blue"
//     }
//   }
// }
// --------------------------------------------------------------------------------

func Test_docs_update_cdf400299acd1c7b1b7bb42e284e3d08(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:cdf400299acd1c7b1b7bb42e284e3d08[]
	res, err := es.Update(
		"test",
		"1",
		strings.NewReader(`{
		  "script": {
		    "source": "ctx._source.tags.add(params.tag)",
		    "lang": "painless",
		    "params": {
		      "tag": "blue"
		    }
		  }
		}`),
		es.Update.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:cdf400299acd1c7b1b7bb42e284e3d08[]
}
