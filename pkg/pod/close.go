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

func (p *Pod) ClosePod(podName string) error {
	if !p.isPodOpened(podName) {
		return ErrPodNotOpened
	}

	podInfo, err := p.GetPodInfoFromPodMap(podName)
	if err != nil {
		return err
	}

	podIndex, err := p.getPodIndex(podName)
	if err != nil {
		return err
	}

	p.removePodFromPodMap(podName)
	podInfo.dir.RemoveFromDirectoryMap(podName)
	p.acc.DeletePodAccount(podIndex)
	return nil
}
