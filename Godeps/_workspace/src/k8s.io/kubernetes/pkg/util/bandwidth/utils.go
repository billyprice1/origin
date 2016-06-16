/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

package bandwidth

import (
	"fmt"

	"k8s.io/kubernetes/pkg/api/resource"
)

var minRsrc = resource.MustParse("1k")
var maxRsrc = resource.MustParse("1P")

func validateBandwidthIsReasonable(rsrc *resource.Quantity) error {
	if rsrc.Value() < minRsrc.Value() {
		return fmt.Errorf("resource is unreasonably small (< 1kbit)")
	}
	if rsrc.Value() > maxRsrc.Value() {
		return fmt.Errorf("resoruce is unreasonably large (> 1Pbit)")
	}
	return nil
}

func ExtractPodBandwidthResources(podAnnotations map[string]string) (ingress, egress *resource.Quantity, err error) {
	str, found := podAnnotations["kubernetes.io/ingress-bandwidth"]
	if found {
		if ingress, err = resource.ParseQuantity(str); err != nil {
			return nil, nil, err
		}
		if err := validateBandwidthIsReasonable(ingress); err != nil {
			return nil, nil, err
		}
	}
	str, found = podAnnotations["kubernetes.io/egress-bandwidth"]
	if found {
		if egress, err = resource.ParseQuantity(str); err != nil {
			return nil, nil, err
		}
		if err := validateBandwidthIsReasonable(egress); err != nil {
			return nil, nil, err
		}
	}
	return ingress, egress, nil
}