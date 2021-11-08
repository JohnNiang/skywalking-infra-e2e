// Licensed to Apache Software Foundation (ASF) under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Apache Software Foundation (ASF) licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package util

import (
	_ "embed"
	"testing"
	"time"
)

func TestRetryAfter(t *testing.T) {
	count := 0
	type args struct {
		retry      RetryFunc
		retryAfter time.Duration
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
		wantErr   bool
	}{{
		name: "Zero retryAfter",
		args: args{
			retry: func() error {
				count++
				return nil
			},
			retryAfter: 0,
		},
		wantCount: 1,
	}, {
		name: "Positive retryAfter",
		args: args{
			retry: func() error {
				count++
				return nil
			},
			retryAfter: 1,
		},
		wantCount: 2,
	},
		{
			name: "Negative retryAfter",
			args: args{
				retry: func() error {
					count++
					return nil
				},
				retryAfter: -1,
			},
			wantCount: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			count = 0
			if err := RetryAfter(tt.args.retry, tt.args.retryAfter); (err != nil) != tt.wantErr {
				t.Errorf("RetryAfter() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantCount != count {
				t.Errorf("RetryAfter() got count=%v, want count=%v", count, tt.wantCount)
			}
		})
	}
}
