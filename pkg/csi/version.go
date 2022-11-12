/*
Copyright 2022 zwwhdls.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package csi

import (
	"encoding/json"
	"fmt"
	"runtime"
)

// These are set during build time via -ldflags
var (
	driverVersion string
	gitCommit     string
	buildDate     string
)

// VersionInfo struct
type VersionInfo struct {
	DriverVersion string
	GitCommit     string
	BuildDate     string
	GoVersion     string
	Compiler      string
	Platform      string
}

// GetVersion returns VersionInfo
func GetVersion() VersionInfo {
	return VersionInfo{
		DriverVersion: driverVersion,
		GitCommit:     gitCommit,
		BuildDate:     buildDate,
		GoVersion:     runtime.Version(),
		Compiler:      runtime.Compiler,
		Platform:      fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

// GetVersionJSON returns version in JSON
func GetVersionJSON() (string, error) {
	info := GetVersion()
	marshalled, err := json.MarshalIndent(&info, "", "  ")
	if err != nil {
		return "", err
	}
	return string(marshalled), nil
}
