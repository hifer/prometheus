// Copyright 2015 The Prometheus Authors
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

package api

import (
	"testing"
	"time"

	clientmodel "github.com/prometheus/client_golang/model"
)

func TestParseTimestampOrNow(t *testing.T) {
	ts, err := parseTimestampOrNow("", testNower)
	if err != nil {
		t.Fatalf("err = %s; want nil", err)
	}
	now := clientmodel.TimestampFromTime(testNower.Now())
	if !now.Equal(ts) {
		t.Fatalf("ts = %v; want ts = %v", ts, now)
	}

	ts, err = parseTimestampOrNow("1426956073.12345", testNower)
	if err != nil {
		t.Fatalf("err = %s; want nil", err)
	}
	expTS := clientmodel.TimestampFromUnixNano(1426956073123000000)
	if !ts.Equal(expTS) {
		t.Fatalf("ts = %v; want %v", ts, expTS)
	}

	_, err = parseTimestampOrNow("123.45foo", testNower)
	if err == nil {
		t.Fatalf("err = nil; want %s", err)
	}
}

func TestParseDuration(t *testing.T) {
	_, err := parseDuration("")
	if err == nil {
		t.Fatalf("err = nil; want %s", err)
	}

	_, err = parseDuration("1234.56foo")
	if err == nil {
		t.Fatalf("err = nil; want %s", err)
	}

	d, err := parseDuration("1234.56789")
	if err != nil {
		t.Fatalf("err = %s; want nil", err)
	}
	expD := time.Duration(1234.56789 * float64(time.Second))
	if d != expD {
		t.Fatalf("d = %v; want %v", d, expD)
	}
}
