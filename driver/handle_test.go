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
	"testing"

	"github.com/hashicorp/nomad/plugins/drivers"
	"github.com/stretchr/testify/assert"
)

func TestKeysToVal(t *testing.T) {
	key, val, err := keysToVal("foo 123")
	assert.NoError(t, err)
	assert.Equal(t, "foo", key)
	assert.Equal(t, uint64(123), val)
}

func TestKeysToValWithBadValue(t *testing.T) {
	key, val, err := keysToVal("foo")
	assert.EqualError(t, err, "line isn't a k/v pair")
	assert.Equal(t, "", key)
	assert.Equal(t, uint64(0), val)
}

func TestTaskHandle(t *testing.T) {
	task := &taskHandle{
		taskConfig: &drivers.TaskConfig{
			ID:   "foo",
			Name: "bar",
		},
		State: drivers.TaskStateRunning,
	}

	status := task.TaskStatus()

	assert.Equal(t, "foo", status.ID)
	assert.Equal(t, "bar", status.Name)

	assert.True(t, task.IsRunning())
}
