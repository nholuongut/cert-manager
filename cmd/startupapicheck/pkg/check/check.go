/*
Copyright 2023 The cert-manager Authors.

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

package check

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/nholuongut/cert-manager/startupapicheck-binary/pkg/check/api"
)

// NewCmdCheck returns a cobra command for checking cert-manager components.
func NewCmdCheck(ctx context.Context) *cobra.Command {
	cmds := NewCmdCreateBare()
	cmds.AddCommand(api.NewCmdCheckApi(ctx))

	return cmds
}

// NewCmdCreateBare returns bare cobra command for checking cert-manager components.
func NewCmdCreateBare() *cobra.Command {
	return &cobra.Command{
		Use:   "check",
		Short: "Check cert-manager components",
		Long:  `Check cert-manager components`,
	}
}
