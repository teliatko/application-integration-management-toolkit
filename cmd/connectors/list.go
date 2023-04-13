// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package connectors

import (
	"errors"
	"internal/apiclient"

	"internal/client/connections"

	"github.com/spf13/cobra"
)

// ListCmd to list Connections
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all connections in the region",
	Long:  "List all connections in the region",
	Args: func(cmd *cobra.Command, args []string) (err error) {
		cmdProject := cmd.Flag("proj")
		cmdRegion := cmd.Flag("reg")

		if err = apiclient.SetRegion(cmdRegion.Value.String()); err != nil {
			return errors.Unwrap(err)
		}
		return apiclient.SetProjectID(cmdProject.Value.String())
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		_, err = connections.List(pageSize,
			cmd.Flag("pageToken").Value.String(),
			cmd.Flag("filter").Value.String(),
			cmd.Flag("orderBy").Value.String())
		return

	},
}

var pageSize int

func init() {
	var pageToken, filter, orderBy string

	ListCmd.Flags().IntVarP(&pageSize, "pageSize", "",
		-1, "The maximum number of versions to return")
	ListCmd.Flags().StringVarP(&pageToken, "pageToken", "",
		"", "A page token, received from a previous call")
	ListCmd.Flags().StringVarP(&filter, "filter", "",
		"", "Filter results")
	ListCmd.Flags().StringVarP(&orderBy, "orderBy", "",
		"", "The results would be returned in order")
}
