// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2017-2018 Canonical Ltd
// Copyright (C) 2018-2019 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

// CHANGED BY CHRISTIAN ALEXANDER BAHRDT
// this file is derivative of
// https://github.com/edgexfoundry/device-sdk-go/blob/edinburgh/example/cmd/device-simple/main.go
package main

import (
    "github.com/datenente/device-bitflow-stream"
    "github.com/datenente/device-bitflow-stream/internal/driver"
	"github.com/edgexfoundry/device-sdk-go/pkg/startup"
)

const (
	serviceName string = "device-bitflow-stream"
)

func main() {
	sd := driver.Driver{}
	startup.Bootstrap(serviceName, device.Version, &sd)
}
