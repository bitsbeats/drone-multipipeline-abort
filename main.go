// Copyright 2019 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package main

import (
	"net/http"

	"github.com/bitsbeats/drone-multipipeline-abort/plugin"
	"github.com/drone/drone-go/plugin/webhook"

	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

// spec provides the plugin settings.
type spec struct {
	Bind   string `envconfig:"DRONE_BIND"`
	Debug  bool   `envconfig:"DRONE_DEBUG"`
	Secret string `envconfig:"DRONE_SECRET"`
	Server string `envconfig:"DRONE_SERVER"`
	Token  string `envconfig:"DRONE_TOKEN"`
}

func main() {
	spec := new(spec)
	err := envconfig.Process("", spec)
	if err != nil {
		logrus.Fatal(err)
	}

	if spec.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
	if spec.Secret == "" {
		logrus.Fatalln("missing secret key")
	}
	if spec.Server == "" {
		logrus.Fatalln("missing server url")
	}
	if spec.Token == "" {
		logrus.Fatalln("missing api token")
	}
	if spec.Bind == "" {
		spec.Bind = ":3000"
	}

	handler := webhook.Handler(
		plugin.New(spec.Server, spec.Token),
		spec.Secret,
		logrus.StandardLogger(),
	)

	logrus.Infof("server listening on address %s", spec.Bind)

	http.Handle("/", handler)
	logrus.Fatal(http.ListenAndServe(spec.Bind, nil))
}
