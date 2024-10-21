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
	"context"
	"github.com/cloudwego-contrib/cwgo-pkg/opensergo/sentinel/hzadapter"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol"
)

type (
	ServerOption = hzadapter.ServerOption
)

type (
	ClientOption = hzadapter.ClientOption
)

// WithServerResourceExtractor sets the resource extractor of the web requests for server side.
func WithServerResourceExtractor(fn func(context.Context, *app.RequestContext) string) ServerOption {
	return hzadapter.WithServerResourceExtractor(fn)
}

// WithServerBlockFallback sets the fallback handler when requests are blocked for server side.
func WithServerBlockFallback(fn func(context.Context, *app.RequestContext)) ServerOption {
	return hzadapter.WithServerBlockFallback(fn)
}

// WithClientResourceExtractor sets the resource extractor of the web requests for client side.
func WithClientResourceExtractor(fn func(context.Context, *protocol.Request,
	*protocol.Response) string,
) ClientOption {
	return hzadapter.WithClientResourceExtractor(fn)
}

// WithClientBlockFallback sets the fallback handler when requests are blocked for client side.
func WithClientBlockFallback(fn func(ctx context.Context, req *protocol.Request,
	resp *protocol.Response, blockError error) error,
) ClientOption {
	return hzadapter.WithClientBlockFallback(fn)
}
