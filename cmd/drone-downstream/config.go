// Copyright (c) 2020, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

// package main	is the main package for the drone-downstream plugin.
package main

import (
	"time"

	"github.com/urfave/cli/v2"

	"github.com/siderolabs/drone-downstream/internal/plugin"
)

// settingsFlags has the cli.Flags for the plugin.Settings.
func settingsFlags(settings *plugin.Settings) []cli.Flag {
	return []cli.Flag{
		&cli.StringSliceFlag{
			Name:        "repositories",
			Usage:       "List of repositories to trigger",
			EnvVars:     []string{"PLUGIN_REPOSITORIES"},
			Destination: &settings.Repos,
		},
		&cli.StringFlag{
			Name:        "server",
			Usage:       "Trigger a drone build on a custom server",
			EnvVars:     []string{"PLUGIN_SERVER", "DOWNSTREAM_SERVER"},
			Destination: &settings.Server,
		},
		&cli.StringFlag{
			Name:        "token",
			Usage:       "Drone API token from your user settings",
			EnvVars:     []string{"PLUGIN_TOKEN", "DRONE_TOKEN", "DOWNSTREAM_TOKEN"},
			Destination: &settings.Token,
		},
		&cli.BoolFlag{
			Name:        "wait",
			Usage:       "Wait for any currently running builds to finish",
			EnvVars:     []string{"PLUGIN_WAIT"},
			Destination: &settings.Wait,
		},
		&cli.DurationFlag{
			Name:        "timeout",
			Value:       time.Duration(60) * time.Second,
			Usage:       "How long to wait on any currently running builds",
			EnvVars:     []string{"PLUGIN_WAIT_TIMEOUT"},
			Destination: &settings.Timeout,
		},
		&cli.BoolFlag{
			Name:        "last-successful",
			Usage:       "Trigger last successful build",
			EnvVars:     []string{"PLUGIN_LAST_SUCCESSFUL"},
			Destination: &settings.LastSuccessful,
		},
		&cli.StringSliceFlag{
			Name:        "params",
			Usage:       "List of params (key=value or file paths of params) to pass to triggered builds",
			EnvVars:     []string{"PLUGIN_PARAMS"},
			Destination: &settings.Params,
		},
		&cli.StringSliceFlag{
			Name:        "params-from-env",
			Usage:       "List of environment variables to pass to triggered builds",
			EnvVars:     []string{"PLUGIN_PARAMS_FROM_ENV"},
			Destination: &settings.ParamsEnv,
		},
		&cli.StringFlag{
			Name:        "deploy",
			Usage:       "Environment to trigger deploy for the respective build",
			EnvVars:     []string{"PLUGIN_DEPLOY"},
			Destination: &settings.Deploy,
		},
		&cli.BoolFlag{
			Name:        "block",
			Usage:       "Block until the triggered build is finished, makes this build fail if triggered build fails",
			EnvVars:     []string{"PLUGIN_BLOCK"},
			Destination: &settings.Block,
		},
		&cli.DurationFlag{
			Name:        "blockTimeout",
			Value:       time.Duration(60) * time.Minute,
			Usage:       "How long to block until the triggered build is finished",
			EnvVars:     []string{"PLUGIN_BLOCK_TIMEOUT"},
			Destination: &settings.BlockTimeout,
		},
	}
}
