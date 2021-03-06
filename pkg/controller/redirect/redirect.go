// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package redirect defines the controller for the deep link redirector.
package redirect

import (
	"context"
	"fmt"

	"github.com/google/exposure-notifications-verification-server/pkg/config"
	"github.com/google/exposure-notifications-verification-server/pkg/render"

	"github.com/google/exposure-notifications-server/pkg/logging"

	"go.uber.org/zap"
)

type Controller struct {
	config           *config.RedirectConfig
	h                *render.Renderer
	logger           *zap.SugaredLogger
	hostnameToRegion map[string]string
}

// New creates a new login controller.
func New(ctx context.Context, config *config.RedirectConfig, h *render.Renderer) (*Controller, error) {
	logger := logging.FromContext(ctx).Named("login")

	cfgMap, err := config.HostnameToRegion()
	if err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}
	logger.Infow("redirect configuration", "hostnameToRegion", cfgMap)

	return &Controller{
		config:           config,
		h:                h,
		logger:           logger,
		hostnameToRegion: cfgMap,
	}, nil
}
