/*
Copyright 2023 The Alibaba Cloud Serverless Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import "time"

type Config struct {
	ClientAddr           string
	GcInterval           time.Duration
	IdleDurationBeforeGC time.Duration
	RctRate              float64
}

var DefaultConfig *Config

func init() {
	DefaultConfig = &Config{
		ClientAddr:           "127.0.0.1:50051",
		GcInterval:           1 * time.Second,
		IdleDurationBeforeGC: 5 * time.Minute,
		RctRate:              0.9,
	}
}
