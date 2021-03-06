// Copyright 2019 Istio Authors
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

package helmreconciler

import (
	"istio.io/istio/operator/pkg/name"
)

func InsertChildrenRecursive(componentName name.ComponentName, tree ComponentTree, children ComponentNameToListMap) {
	tree[componentName] = make(ComponentTree)
	for _, child := range children[componentName] {
		InsertChildrenRecursive(child, tree[componentName].(ComponentTree), children)
	}
}
