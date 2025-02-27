/*
Copyright 2021 The Knative Authors

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

package apiserversource_test

import (
	"embed"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1 "knative.dev/eventing/pkg/apis/sources/v1"

	"knative.dev/eventing/test/rekt/resources/apiserversource"

	duckv1 "knative.dev/pkg/apis/duck/v1"
	"knative.dev/reconciler-test/pkg/manifest"
)

//go:embed *.yaml
var yaml embed.FS

// The following examples validate the processing of the With* helper methods
// applied to config and go template parser.

func Example_min() {
	images := map[string]string{}
	cfg := map[string]interface{}{
		"name":      "foo",
		"namespace": "bar",
	}

	files, err := manifest.ExecuteYAML(yaml, images, cfg)
	if err != nil {
		panic(err)
	}

	manifest.OutputYAML(os.Stdout, files)
	// Output:
	// apiVersion: sources.knative.dev/v1
	// kind: ApiServerSource
	// metadata:
	//   name: foo
	//   namespace: bar
	// spec:
}

func Example_withServiceAccountName() {
	images := map[string]string{}
	cfg := map[string]interface{}{
		"name":      "foo",
		"namespace": "bar",
	}

	apiserversource.WithServiceAccountName("src-sa")(cfg)

	files, err := manifest.ExecuteYAML(yaml, images, cfg)
	if err != nil {
		panic(err)
	}

	manifest.OutputYAML(os.Stdout, files)
	// Output:
	// apiVersion: sources.knative.dev/v1
	// kind: ApiServerSource
	// metadata:
	//   name: foo
	//   namespace: bar
	// spec:
	//   serviceAccountName: src-sa
}

func Example_withEventMode() {
	images := map[string]string{}
	cfg := map[string]interface{}{
		"name":      "foo",
		"namespace": "bar",
	}

	apiserversource.WithEventMode(v1.ReferenceMode)(cfg)

	files, err := manifest.ExecuteYAML(yaml, images, cfg)
	if err != nil {
		panic(err)
	}

	manifest.OutputYAML(os.Stdout, files)
	// Output:
	// apiVersion: sources.knative.dev/v1
	// kind: ApiServerSource
	// metadata:
	//   name: foo
	//   namespace: bar
	// spec:
	//   mode: Reference
}

func Example_withSink() {
	images := map[string]string{}
	cfg := map[string]interface{}{
		"name":      "foo",
		"namespace": "bar",
	}

	sinkRef := &duckv1.KReference{
		Kind:       "sinkkind",
		Namespace:  "sinknamespace",
		Name:       "sinkname",
		APIVersion: "sinkversion",
	}
	apiserversource.WithSink(sinkRef, "uri/parts")(cfg)

	files, err := manifest.ExecuteYAML(yaml, images, cfg)
	if err != nil {
		panic(err)
	}

	manifest.OutputYAML(os.Stdout, files)
	// Output:
	// apiVersion: sources.knative.dev/v1
	// kind: ApiServerSource
	// metadata:
	//   name: foo
	//   namespace: bar
	// spec:
	//   sink:
	//     ref:
	//       kind: sinkkind
	//       namespace: sinknamespace
	//       name: sinkname
	//       apiVersion: sinkversion
	//     uri: uri/parts
}

func Example_withResources() {
	images := map[string]string{}
	cfg := map[string]interface{}{
		"name":      "foo",
		"namespace": "bar",
	}

	res1 := v1.APIVersionKindSelector{
		APIVersion:    "res1apiVersion",
		Kind:          "res1kind",
		LabelSelector: nil,
	}

	res2 := v1.APIVersionKindSelector{
		APIVersion: "res2apiVersion",
		Kind:       "res2kind",
		LabelSelector: &metav1.LabelSelector{
			MatchLabels:      map[string]string{"foo": "bar"},
			MatchExpressions: nil,
		},
	}

	res3 := v1.APIVersionKindSelector{
		APIVersion: "res3apiVersion",
		Kind:       "res3kind",
		LabelSelector: &metav1.LabelSelector{
			MatchLabels: map[string]string{"foo": "bar"},
			MatchExpressions: []metav1.LabelSelectorRequirement{{
				Key:      "daf",
				Operator: "uk",
				Values:   []string{"a", "b"},
			}},
		},
	}

	apiserversource.WithResources(res1, res2, res3)(cfg)

	files, err := manifest.ExecuteYAML(yaml, images, cfg)
	if err != nil {
		panic(err)
	}

	manifest.OutputYAML(os.Stdout, files)
	// Output:
	// apiVersion: sources.knative.dev/v1
	// kind: ApiServerSource
	// metadata:
	//   name: foo
	//   namespace: bar
	// spec:
	//   resources:
	//     - apiVersion: res1apiVersion
	//       kind: res1kind
	//     - apiVersion: res2apiVersion
	//       kind: res2kind
	//       selector:
	//         matchLabels:
	//           foo: bar
	//     - apiVersion: res3apiVersion
	//       kind: res3kind
	//       selector:
	//         matchLabels:
	//           foo: bar
	//         matchExpressions:
	//           - key: daf
	//             operator: uk
	//             values:
	//               - a
	//               - b
}

func Example_full() {
	images := map[string]string{}
	cfg := map[string]interface{}{
		"name":      "foo",
		"namespace": "bar",
	}

	sinkRef := &duckv1.KReference{
		Kind:       "sinkkind",
		Namespace:  "sinknamespace",
		Name:       "sinkname",
		APIVersion: "sinkversion",
	}

	res1 := v1.APIVersionKindSelector{
		APIVersion:    "res1apiVersion",
		Kind:          "res1kind",
		LabelSelector: nil,
	}

	res2 := v1.APIVersionKindSelector{
		APIVersion: "res2apiVersion",
		Kind:       "res2kind",
		LabelSelector: &metav1.LabelSelector{
			MatchLabels:      map[string]string{"foo": "bar"},
			MatchExpressions: nil,
		},
	}

	res3 := v1.APIVersionKindSelector{
		APIVersion: "res3apiVersion",
		Kind:       "res3kind",
		LabelSelector: &metav1.LabelSelector{
			MatchLabels: map[string]string{"foo": "bar"},
			MatchExpressions: []metav1.LabelSelectorRequirement{{
				Key:      "daf",
				Operator: "uk",
				Values:   []string{"a", "b"},
			}},
		},
	}

	apiserversource.WithServiceAccountName("src-sa")(cfg)
	apiserversource.WithEventMode(v1.ReferenceMode)(cfg)
	apiserversource.WithSink(sinkRef, "uri/parts")(cfg)
	apiserversource.WithResources(res1, res2, res3)(cfg)

	files, err := manifest.ExecuteYAML(yaml, images, cfg)
	if err != nil {
		panic(err)
	}

	manifest.OutputYAML(os.Stdout, files)
	// Output:
	// apiVersion: sources.knative.dev/v1
	// kind: ApiServerSource
	// metadata:
	//   name: foo
	//   namespace: bar
	// spec:
	//   serviceAccountName: src-sa
	//   mode: Reference
	//   resources:
	//     - apiVersion: res1apiVersion
	//       kind: res1kind
	//     - apiVersion: res2apiVersion
	//       kind: res2kind
	//       selector:
	//         matchLabels:
	//           foo: bar
	//     - apiVersion: res3apiVersion
	//       kind: res3kind
	//       selector:
	//         matchLabels:
	//           foo: bar
	//         matchExpressions:
	//           - key: daf
	//             operator: uk
	//             values:
	//               - a
	//               - b
	//   sink:
	//     ref:
	//       kind: sinkkind
	//       namespace: sinknamespace
	//       name: sinkname
	//       apiVersion: sinkversion
	//     uri: uri/parts
}
