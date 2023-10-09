// Copyright 2019 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin/webhook"
)

// New returns a new webhook extension.
func New(server, token string) webhook.Plugin {
	return &plugin{
		drone.NewClient(
			server,
			&http.Client{
				Transport: &transport{token},
			},
		),
	}
}

type plugin struct {
	client drone.Client
}

func (p *plugin) Deliver(ctx context.Context, req *webhook.Request) error {
	if req.Event != webhook.EventBuild {
		return nil
	}

	build := req.Build

	if slices.Contains([]string{drone.StatusError, drone.StatusFailing, drone.StatusKilled}, build.Status) {
		return nil
	}

	if slices.ContainsFunc(
		build.Stages,
		func(stage *drone.Stage) bool {
			return slices.Contains([]string{drone.StatusError, drone.StatusFailing, drone.StatusKilled}, stage.Status)
		},
	) {
		repos, err := p.client.RepoListAll(drone.ListOptions{})
		if err != nil {
			return err
		}

		repoIndex := slices.IndexFunc(
			repos,
			func(repo *drone.Repo) bool {
				return repo.ID == build.RepoID
			},
		)
		if repoIndex == -1 {
			return fmt.Errorf("could not find repo id %d", build.RepoID)
		}

		repo := repos[repoIndex]

		err = p.client.BuildCancel(repo.Namespace, repo.Name, int(build.Number))
		if err != nil {
			return err
		}
	}

	return nil
}
