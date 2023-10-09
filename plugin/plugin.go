// Copyright 2019 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"

	"github.com/drone/drone-go/plugin/webhook"
)

// New returns a new webhook extension.
func New() webhook.Plugin {
	return &plugin{}
}

type plugin struct {}

func (p *plugin) Deliver(ctx context.Context, req *webhook.Request) error {
	return nil
}
