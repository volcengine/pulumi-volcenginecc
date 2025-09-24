// Copyright 2016-2023, Pulumi Corporation.
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
// SPDX-License-Identifier: MPL-2.0
// This file has been modified by Beijing Volcano Engine Technology Co., Ltd. on 2025

package volcenginecc

import (
	// embed is used to store bridge-metadata.json in the compiled binary
	_ "embed"
	"fmt"
	"path/filepath"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/info"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	"github.com/volcengine/pulumi-volcenginecc/provider/cmd/pulumi-resource-volcenginecc/token"
	"github.com/volcengine/pulumi-volcenginecc/provider/pkg/version"
	"github.com/volcengine/terraform-provider-volcenginecc/shim"

	pf "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
)

// all of the random token components used below.
const (
	volcengeccpkg = "volcenginecc"
)

//go:embed cmd/pulumi-resource-volcenginecc/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the random package.
func Provider() tfbridge.ProviderInfo {
	prov := tfbridge.ProviderInfo{
		P:           pf.ShimProvider(shim.NewProvider()),
		Name:        "volcenginecc",
		Description: "A Pulumi package to safely use randomness in Pulumi programs.",
		Keywords:    []string{"volcengine", "volcenginecc", "category/cloud"},
		License:     "MPL-2.0",
		Homepage:    "https://github.com/volcengine/pulumi-volcenginecc",
		Publisher:   "Volcengine",
		LogoURL:     "https://avatars.githubusercontent.com/u/67365215",
		Repository:  "https://github.com/volcengine/pulumi-volcenginecc",
		Version:     version.Version,
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		PluginDownloadURL: "github://api.github.com/volcengine",
		MetadataInfo:      tfbridge.NewProviderMetadata(metadata),
		Resources:         map[string]*tfbridge.ResourceInfo{},
		GitHubHost:        "github.com",
		GitHubOrg:         "volcengine",
		//UpstreamRepoPath:  "/Users/bytedance/workspace/terraform-provider-volcenginecc",
		Config: map[string]*tfbridge.SchemaInfo{
			"region": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VOLCENGINE_REGION"},
				},
			},
			"access_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VOLCENGINE_ACCESS_KEY"},
				},
			},
			"secret_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VOLCENGINE_SECRET_KEY"},
				},
			},
			"disable_ssl": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VOLCENGINE_DISABLE_SSL"},
				},
			},
			"customer_headers": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VOLCENGINE_CUSTOMER_HEADERS"},
				},
			},
			"proxy_url": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"VOLCENGINE_PROXY_URL"},
				},
			},
			"endpoints": {
				Fields: map[string]*info.Schema{
					"cloudcontrolapi": {
						Default: &tfbridge.DefaultInfo{
							EnvVars: []string{"VOLCENGINE_CC_ENDPOINT"},
						},
					},
					"sts": {
						Default: &tfbridge.DefaultInfo{
							EnvVars: []string{"VOLCENGINE_STS_ENDPOINT"},
						},
					},
				},
			},
			"assume_role": {
				Fields: map[string]*info.Schema{
					"assume_role_trn": {
						Default: &tfbridge.DefaultInfo{
							EnvVars: []string{"VOLCENGINE_ASSUME_ROLE_TRN"},
						},
					},
					"policy": {
						Default: &tfbridge.DefaultInfo{
							EnvVars: []string{"VOLCENGINE_ASSUME_ROLE_POLICY"},
						},
					},
					"duration_seconds": {
						Default: &tfbridge.DefaultInfo{
							EnvVars: []string{"VOLCENGINE_ASSUME_ROLE_DURATION_SECONDS"},
						},
					},
				},
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/volcengine/pulumi-%[1]s/sdk/", volcengeccpkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				volcengeccpkg,
			),

			GenerateResourceContainerTypes: true,
			RespectSchemaVersion:           true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			RootNamespace: "Volcengine.Pulumi",
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "pulumi_volcenginecc",
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions

			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			// Overlay: &tfbridge.OverlayInfo{},
			PackageName: "@volcengine/pulumi-volcenginecc",
		},
		Java: &tfbridge.JavaInfo{
			BasePackage: "com.volcengine",
		},
		EnableAccurateBridgePreview: true,
		EnableRawStateDelta:         true,
	}

	makeToken := func(module, name string) (string, error) {
		return tokens.MakeStandard(volcengeccpkg)(module, name)
	}
	prov.MustComputeTokens(token.VolcengineToken("volcenginecc_", makeToken))
	prov.MustApplyAutoAliases()

	return prov
}
