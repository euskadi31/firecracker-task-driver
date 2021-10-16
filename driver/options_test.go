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
	"regexp"
	"testing"

	"github.com/firecracker-microvm/firecracker-go-sdk"
	"github.com/stretchr/testify/assert"
)

var (
	socketPathPattern = regexp.MustCompile(`/root/.firecracker.sock-\d+-\d+`)
	macAddressPattern = regexp.MustCompile(`^([0-9A-F]{2}:){5}[0-9A-F]{2}$`)
	vethPattern       = regexp.MustCompile(`veth([a-f0-9]+)`)
)

func TestGetSocketPath(t *testing.T) {
	path := getSocketPath()
	assert.True(t, socketPathPattern.MatchString(path), path)
}

func TestGenmacaddr(t *testing.T) {
	mac, err := genmacaddr()
	assert.NoError(t, err)

	assert.True(t, macAddressPattern.MatchString(mac), mac)
}

func TestRandomVethName(t *testing.T) {
	name, err := RandomVethName()
	assert.NoError(t, err)

	assert.True(t, vethPattern.MatchString(name), name)
}

func TestOptions(t *testing.T) {
	opts := newOptions()
	opts.FcMetadata = `{"foo":"bar"}`
	opts.FcKernelImage = "/vmlinux-v5.14"
	opts.FcNetworkName = "my-app"
	opts.FcRootDrivePath = "/ubuntu.rootfs.ext4"

	fcc, err := opts.getFirecrackerConfig("foo")
	assert.NoError(t, err)
	//assert.Nil(t, fcc)

	assert.True(t, socketPathPattern.MatchString(fcc.SocketPath))
	assert.Equal(t, "/vmlinux-v5.14", fcc.KernelImagePath)

	assert.Equal(t, 1, len(fcc.Drives))
	assert.Equal(t, &opts.FcRootDrivePath, fcc.Drives[0].PathOnHost)
	assert.Equal(t, firecracker.String("1"), fcc.Drives[0].DriveID)
	assert.Equal(t, firecracker.Bool(true), fcc.Drives[0].IsRootDevice)
	assert.Equal(t, firecracker.Bool(false), fcc.Drives[0].IsReadOnly)

	assert.Equal(t, 1, len(fcc.NetworkInterfaces))
	assert.True(t, vethPattern.MatchString(fcc.NetworkInterfaces[0].CNIConfiguration.IfName))
	assert.Equal(t, "my-app", fcc.NetworkInterfaces[0].CNIConfiguration.NetworkName)
}
