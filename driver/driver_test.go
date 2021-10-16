/* Firecracker-task-driver is a task driver for Hashicorp's nomad that allows
 * to create microvms using AWS Firecracker vmm
 * Copyright (C) 2019 Carlos Neira cneirabustos@gmail.com
 * Copyright (C) 2021 Axel Etcheverry axel@etcheverry.biz
 *
 * Licensed under the Apache License, Version 2.0 (the "License")
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 */

package firevm

import (
	"context"
	"testing"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/nomad/plugins/base"
	"github.com/stretchr/testify/assert"
)

func TestNewFirecrackerDriver(t *testing.T) {
	logger := hclog.NewNullLogger()
	driver := NewFirecrackerDriver(logger)

	info, err := driver.PluginInfo()
	assert.NoError(t, err)

	assert.Equal(t, "firecracker-task-driver", info.Name)
	assert.Equal(t, base.PluginTypeDriver, info.Type)
	assert.Equal(t, []string{"v0.1.0"}, info.PluginApiVersions)

	spec, err := driver.ConfigSchema()
	assert.NoError(t, err)
	assert.Nil(t, spec)

	spec, err = driver.TaskConfigSchema()
	assert.NoError(t, err)
	assert.Same(t, taskConfigSpec, spec)

	caps, err := driver.Capabilities()
	assert.NoError(t, err)
	assert.Same(t, capabilities, caps)

	ctx := context.Background()

	fingerprintCh, err := driver.Fingerprint(ctx)

	fingerprint := <-fingerprintCh

	assert.Equal(t, "ready", fingerprint.HealthDescription)

	result, err := driver.ExecTask("foo", []string{}, 2*time.Second)
	assert.EqualError(t, err, "Firecracker-task-driver does not support exec")
	assert.Nil(t, result)
}
