/*
Copyright © 2020 FairOS Authors

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

package user

import (
	"os"
	"path/filepath"
)

func (u *Users) ListAllUsers(dataDir string) ([]string, error) {
	var users []string
	destDir := filepath.Join(dataDir, userDirectoryName)
	err := filepath.Walk(destDir,
		func(path string, info os.FileInfo, err error) error {
			if info == nil {
				return nil
			}
			if info.Name() == userDirectoryName {
				return nil
			}
			userName := info.Name()
			userName = "<User> " + userName
			users = append(users, userName)
			return nil
		})
	if err != nil {
		return nil, err
	}
	return users, nil
}
