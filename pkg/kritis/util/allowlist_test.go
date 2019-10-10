/*
Copyright 2018 Google LLC

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
package util

import (
	"testing"

	"github.com/grafeas/kritis/pkg/kritis/testutil"
)

func Test_RemoveGloballyAllowedImages(t *testing.T) {
	tests := []struct {
		name     string
		images   []string
		expected []string
	}{
		{
			name: "images in allowlist",
			images: []string{
				"gcr.io/kritis-project/kritis-server:tag",
				"gcr.io/kritis-project/kritis-server@sha256:0000000000000000000000000000000000000000000000000000000000000000",
			},
			expected: []string{},
		},
		{
			name: "some images not allowlisted",
			images: []string{
				"gcr.io/kritis-project/kritis-server:tag",
				"gcr.io/some/image@sha256:0000000000000000000000000000000000000000000000000000000000000000",
			},
			expected: []string{"gcr.io/some/image@sha256:0000000000000000000000000000000000000000000000000000000000000000"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := RemoveGloballyAllowedImages(test.images)
			testutil.DeepEqual(t, test.expected, actual)
		})
	}
}

func Test_imageInAllowlist(t *testing.T) {
	tests := []struct {
		name     string
		image    string
		expected bool
	}{
		{
			name:     "test image in allowlist",
			image:    "gcr.io/kritis-project/kritis-server:tag",
			expected: true,
		},
		{
			name:     "test image with digest in allowlist",
			image:    "gcr.io/kritis-project/kritis-server@sha:123",
			expected: true,
		},
		{
			name:     "test image not in allowlist",
			image:    "some/image",
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := imageInAllowlist(test.image)
			testutil.CheckErrorAndDeepEqual(t, false, err, test.expected, actual)
		})
	}
}
