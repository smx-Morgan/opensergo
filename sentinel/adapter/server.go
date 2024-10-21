// Copyright 2021 CloudWeGo authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package adapter

import (
	"github.com/cloudwego-contrib/cwgo-pkg/opensergo/sentinel/hzadapter"

	"github.com/cloudwego/hertz/pkg/app"
)

// SentinelServerMiddleware returns new app.HandlerFunc
// Default resource name is {method}:{path}, such as "GET:/api/users/:id"
// Default block fallback is returning 429 code
// Define your own behavior by setting serverOptions
func SentinelServerMiddleware(opts ...ServerOption) app.HandlerFunc {
	return hzadapter.SentinelServerMiddleware(opts...)
}
