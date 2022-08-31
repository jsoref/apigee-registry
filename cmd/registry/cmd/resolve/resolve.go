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

package resolve

import (
	"context"
	"fmt"

	"github.com/apigee/registry/cmd/registry/controller"
	"github.com/apigee/registry/cmd/registry/core"
	"github.com/apigee/registry/log"
	"github.com/apigee/registry/pkg/connection"
	"github.com/apigee/registry/rpc"
	"github.com/apigee/registry/server/registry/names"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

func fetchManifest(
	ctx context.Context,
	client connection.RegistryClient,
	manifestName string) (*rpc.Manifest, error) {
	manifest := &rpc.Manifest{}
	body, err := client.GetArtifactContents(
		ctx,
		&rpc.GetArtifactContentsRequest{
			Name: manifestName,
		})
	if err != nil {
		return nil, err
	}

	contents := body.GetData()
	err = proto.Unmarshal(contents, manifest)
	if err != nil {
		return nil, err
	}

	return manifest, nil
}

func Command() *cobra.Command {
	var dryRun bool
	var jobs int
	var maxActions int
	cmd := &cobra.Command{
		Use:   "resolve MANIFEST_RESOURCE",
		Short: "resolve the dependencies and update the registry state (experimental)",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			c, err := connection.ActiveConfig()
			if err != nil {
				log.FromContext(ctx).WithError(err).Fatal("Failed to get config")
			}
			args[0] = c.FQName(args[0])

			name, err := names.ParseArtifact(args[0])
			if err != nil {
				log.FromContext(ctx).WithError(err).Fatal("Invalid manifest resource name")
			}

			registryClient, err := connection.NewRegistryClient(ctx)
			if err != nil {
				log.FromContext(ctx).WithError(err).Fatal("Failed to get client")
			}

			manifest, err := fetchManifest(ctx, registryClient, name.String())
			if err != nil {
				log.FromContext(ctx).WithError(err).Fatal("Failed to fetch manifest")
			}

			client := &controller.RegistryLister{RegistryClient: registryClient}

			log.Debug(ctx, "Generating the list of actions...")
			actions := controller.ProcessManifest(ctx, client, name.ProjectID(), manifest, maxActions)

			// The monitoring metrics/dashboards are built on top of the format of the log messages here.
			// Check the metric filters before making any changes to the format.
			// Location: registry/deployments/controller/dashboard/*
			if len(actions) == 0 {
				log.Debug(ctx, "Generated 0 actions. The registry is already in a resolved state.")
				return
			}

			log.Debugf(ctx, "Generated %d actions.", len(actions))

			// If dry_run is set to true, print the generated actions and exit
			if dryRun {
				for _, a := range actions {
					log.Debugf(ctx, "Action: %q", a.Command)
				}
				return
			}

			log.Debug(ctx, "Starting execution...")
			taskQueue, wait := core.WorkerPoolWithWarnings(ctx, jobs)
			defer wait()
			// Submit tasks to taskQueue
			for i := 0; i < len(actions) && i < maxActions; i++ {
				taskQueue <- &controller.ExecCommandTask{
					Action: actions[i],
					TaskID: fmt.Sprintf("%.8s", uuid.New()),
				}
			}
		},
	}

	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "if set, actions will only be printed and not executed")
	cmd.Flags().IntVarP(&jobs, "jobs", "j", 10, "Number of actions to execute simultaneously")
	cmd.Flags().IntVarP(&maxActions, "max-actions", "a", 100, "Maximum number of actions to execute")
	return cmd
}
