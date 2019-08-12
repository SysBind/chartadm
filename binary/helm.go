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

package binary

import (
	"fmt"
	"github.com/sysbind/chartadm/apis"
	"github.com/sysbind/chartadm/constants"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func Download(config apis.ChartAdmConfig) error {
	url := fmt.Sprintf("%s/%s", config.ReleaseURL, releaseFile(config.Version))
	get(url, "/tmp/archive.tar.gz")
	return nil
}

func releaseFile(version string) string {
	return fmt.Sprintf("helm-v%s-linux-amd64.tar.gz", version)
}

func get(url, archive string) error {
	log.Printf("[download] downloading helm from %s to %s\n", url, archive)
	argStr := fmt.Sprintf("--connect-timeout %v --progress-bar --location --output %v %v", int(constants.DefaultDownloadConnectTimeout/time.Second), archive, url)
	cmd := exec.Command("curl", strings.Fields(argStr)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		return err
	}
	err = cmd.Wait()
	if err != nil {
		return err
	}
	return nil
}
