// Copyright 2013 Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import "testing"

func testMetric(t testing.TB) {
	var scenarios = []struct {
		input  map[string]string
		hash   uint64
		rowkey string
	}{
		{
			input:  map[string]string{},
			rowkey: "02676020557754725067--0-",
			hash:   2676020557754725067,
		},
		{
			input: map[string]string{
				"first_name":   "electro",
				"occupation":   "robot",
				"manufacturer": "westinghouse",
			},
			rowkey: "04776841610193542734-f-6-t",
			hash:   4776841610193542734,
		},
		{
			input: map[string]string{
				"x": "y",
			},
			rowkey: "01306929544689993150-x-2-y",
			hash:   1306929544689993150,
		},
	}

	for i, scenario := range scenarios {
		metric := Metric{}
		for key, value := range scenario.input {
			metric[LabelName(key)] = LabelValue(value)
		}

		expectedRowKey := scenario.rowkey
		expectedHash := scenario.hash
		fingerprint := &Fingerprint{}
		fingerprint.LoadFromMetric(metric)
		actualRowKey := fingerprint.String()
		actualHash := fingerprint.Hash

		if expectedRowKey != actualRowKey {
			t.Errorf("%d. expected %s, got %s", i, expectedRowKey, actualRowKey)
		}
		if actualHash != expectedHash {
			t.Errorf("%d. expected %d, got %d", i, expectedHash, actualHash)
		}
	}
}

func TestMetric(t *testing.T) {
	testMetric(t)
}

func BenchmarkMetric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testMetric(b)
	}
}
