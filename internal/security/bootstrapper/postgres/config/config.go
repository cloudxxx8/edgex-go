//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package config

import (
	bootstrapConfig "github.com/edgexfoundry/go-mod-bootstrap/v4/config"
)

// ConfigurationStruct has a 1:1 relationship to the configuration.yaml for the service. Writable is
// the runtime extension of the static configuration.
type ConfigurationStruct struct {
	LogLevel       string
	Database       bootstrapConfig.Database
	DatabaseConfig DatabaseBootstrapInitInfo
}

// DatabaseBootstrapInitInfo contains the initialization db scripts properties for bootstrapping the database
type DatabaseBootstrapInitInfo struct {
	Path       string
	Name       string
	MaxClients int
}

// Implement interface.Configuration

// UpdateFromRaw converts configuration received from the registry to a service-specific
// configuration struct which is then used to overwrite the service's existing configuration struct.
func (c *ConfigurationStruct) UpdateFromRaw(rawConfig interface{}) bool {
	return false
}

// EmptyWritablePtr returns a pointer to a service-specific empty WritableInfo struct.  It is used
// by the bootstrap to provide the appropriate structure to registry.Client's WatchForChanges().
func (c *ConfigurationStruct) EmptyWritablePtr() interface{} {
	return nil
}

// GetWritablePtr returns pointer to the writable section
// Not needed for this service, so return nil
func (c *ConfigurationStruct) GetWritablePtr() any {
	return nil
}

// UpdateWritableFromRaw converts configuration received from the registry to a service-specific
// WritableInfo struct which is then used to overwrite the service's existing configuration's
// WritableInfo struct.
func (c *ConfigurationStruct) UpdateWritableFromRaw(rawWritable interface{}) bool {
	return false
}

// GetBootstrap returns the configuration elements required by the bootstrap.  Currently, a copy of
// the configuration data is returned.  This is intended to be temporary -- since
// ConfigurationStruct drives the configuration.yaml's structure -- until we can make
// backwards-breaking configuration.yaml changes (which would consolidate these fields into an
// bootstrapConfig.BootstrapConfiguration struct contained within ConfigurationStruct).
func (c *ConfigurationStruct) GetBootstrap() bootstrapConfig.BootstrapConfiguration {
	// temporary until we can make backwards-breaking configuration.yaml change
	return bootstrapConfig.BootstrapConfiguration{}
}

// GetLogLevel returns the current ConfigurationStruct's log level.
func (c *ConfigurationStruct) GetLogLevel() string {
	return c.LogLevel
}

// GetRegistryInfo returns the RegistryInfo from the ConfigurationStruct.
func (c *ConfigurationStruct) GetRegistryInfo() bootstrapConfig.RegistryInfo {
	return bootstrapConfig.RegistryInfo{}
}

// GetDatabaseInfo returns a database information.
func (c *ConfigurationStruct) GetDatabaseInfo() bootstrapConfig.Database {
	return c.Database
}

// GetInsecureSecrets returns the service's InsecureSecrets which this service doesn't support
func (c *ConfigurationStruct) GetInsecureSecrets() bootstrapConfig.InsecureSecrets {
	return nil
}

// GetTelemetryInfo returns the service's Telemetry settings of which this service doesn't have. I.e. service has no metrics
func (c *ConfigurationStruct) GetTelemetryInfo() *bootstrapConfig.TelemetryInfo {
	return &bootstrapConfig.TelemetryInfo{}
}
