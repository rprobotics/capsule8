// Copyright 2017 Capsule8, Inc.
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

package main

import (
	"flag"

	"github.com/capsule8/capsule8/pkg/services"
	"github.com/capsule8/capsule8/pkg/version"
)

func main() {
	// Set "alsologtostderr" flag so that glog messages go stderr as well as /tmp.
	flag.Set("alsologtostderr", "true")
	flag.Parse()

	// Log version and build at "Starting ..." for debugging
	version.InitialBuildLog("sensor")

	services.Main()
}
