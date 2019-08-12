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

package apis

import (
	"github.com/sysbind/chartadm/constants"
)

// ChartAdmConfig holds chartadm global configuration
type ChartAdmConfig struct {
	Version    string
	ReleaseURL string
	CacheDir   string
}

// SetDefaults sets configuration values defined at build time
func SetDefaults(cfg *ChartAdmConfig) {
	cfg.Version = constants.DefaultVersion
	cfg.ReleaseURL = constants.ReleaseURL
	cfg.CacheDir = constants.DefaultCacheBaseDir
}
