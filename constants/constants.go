/*
Copyright © 5769 (2019) Asaf Ohayon <asaf@sysbind.co.il>

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

package constants

import "time"

// Command-line flag defaults
const (
	ReleaseURL                    = "https://get.helm.sh"
	DefaultVersion                = "2.14.3"
	DefaultInstallDir             = "/opt/bin/"
	DefaultDownloadConnectTimeout = 10 * time.Second
)
