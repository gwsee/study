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

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/update-by-query.asciidoc#L593>
//
// --------------------------------------------------------------------------------
// POST twitter/_update_by_query?refresh&slices=5
// {
//   "script": {
//     "source": "ctx._source['extra'] = 'test'"
//   }
// }
// --------------------------------------------------------------------------------

func Test_docs_update_by_query_ea02de2dbe05091fcb0dac72c8ba5f83(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:ea02de2dbe05091fcb0dac72c8ba5f83[]
	res, err := es.UpdateByQuery(
		[]string{"twitter"},
		es.UpdateByQuery.WithBody(strings.NewReader(`{
		  "script": {
		    "source": "ctx._source['extra'] = 'test'"
		  }
		}`)),
		es.UpdateByQuery.WithRefresh(true),
		es.UpdateByQuery.WithSlices("5"),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:ea02de2dbe05091fcb0dac72c8ba5f83[]
}
