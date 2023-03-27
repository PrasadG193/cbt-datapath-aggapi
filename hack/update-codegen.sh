#!/usr/bin/env bash

# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname ${BASH_SOURCE})/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd ${SCRIPT_ROOT}; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}

# generate the code with:
# --output-base    because this script should also be able to run inside the vendor dir of
#                  k8s.io/kubernetes. The output-base is needed for the generators to output into the vendor dir
#                  instead of the $GOPATH directly. For normal projects this can be dropped.
# to use your own boilerplate text use:
#   --go-header-file ${SCRIPT_ROOT}/hack/custom-boilerplate.go.txt

for group in cbt;
do
  bash ${CODEGEN_PKG}/generate-internal-groups.sh \
    "deepcopy,client,lister,informer,openapi" \
    github.com/PrasadG193/cbt-datapath-aggapi/pkg/generated/${group} \
    github.com/PrasadG193/cbt-datapath-aggapi/pkg/apis \
    github.com/PrasadG193/cbt-datapath-aggapi/pkg/apis \
    ${group}:v1alpha1 \
    --output-base "$(dirname ${BASH_SOURCE})/../" \
    --go-header-file ${SCRIPT_ROOT}/hack/boilerplate.go.txt -v10
done

rm -rf pkg/generated
mkdir -p pkg/generated
cp -R github.com/PrasadG193/cbt-datapath-aggapi/pkg/generated/* pkg/generated/
cp -R github.com/PrasadG193/cbt-datapath-aggapi/pkg/apis/* pkg/apis/
rm -rf github.com
