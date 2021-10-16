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

	"github.com/stretchr/testify/assert"
)

func TestTaskStore(t *testing.T) {
	ts := newTaskStore()

	task, ok := ts.Get("foo")
	assert.False(t, ok)
	assert.Nil(t, task)

	expected := &taskHandle{}

	ts.Set("foo", expected)

	task, ok = ts.Get("foo")
	assert.True(t, ok)
	assert.Same(t, expected, task)

	ts.Delete("foo")

	task, ok = ts.Get("foo")
	assert.False(t, ok)
	assert.Nil(t, task)
}
