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

package pod

import (
	"fmt"
	"strings"
)

func (p *Pod) DeletePod(podName string) error {
	pods, sharedPods, err := p.loadUserPods()
	if err != nil {
		return err
	}
	found := false
	var podIndex int
	for index, pod := range pods {
		if strings.Trim(pod, "\n") == podName {
			delete(pods, index)
			podIndex = index
			found = true
		}
	}
	if !found {
		return fmt.Errorf("pod not found")
	}

	//delete the pod inode
	podInfo, err := p.GetPodInfoFromPodMap(podName)
	if err != nil {
		return err
	}
	err = podInfo.dir.DeletePodInode(podName)
	if err != nil {
		return err
	}

	// remove it from other data structures
	podInfo.dir.RemoveFromDirectoryMap(podName)
	p.removePodFromPodMap(podName)
	p.acc.DeletePodAccount(podIndex)

	// if last pod is deleted.. something should be there to update the feed
	if pods == nil {
		pods = make(map[int]string)
		pods[0] = ""
	}

	// remove the pod finally
	return p.storeUserPods(pods, sharedPods)
}
