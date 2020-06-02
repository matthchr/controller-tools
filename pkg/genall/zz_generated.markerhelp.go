// +build !ignore_autogenerated

/*
Copyright2019 The Kubernetes Authors.

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

// Code generated by helpgen. DO NOT EDIT.

package genall

import (
	"github.com/matthchr/controller-tools/pkg/markers"
)

func (InputPaths) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "",
		DetailedHelp: markers.DetailedHelp{
			Summary: "represents paths and go-style path patterns to use as package roots.",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (OutputArtifacts) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "",
		DetailedHelp: markers.DetailedHelp{
			Summary: "outputs artifacts to different locations, depending on whether they're package-associated or not. ",
			Details: "Non-package associated artifacts are output to the Config directory, while package-associated ones are output to their package's source files' directory, unless an alternate path is specified in Code.",
		},
		FieldHelp: map[string]markers.DetailedHelp{
			"Config": markers.DetailedHelp{
				Summary: "points to the directory to which to write configuration.",
				Details: "",
			},
			"Code": markers.DetailedHelp{
				Summary: "overrides the directory in which to write new code (defaults to where the existing code lives).",
				Details: "",
			},
		},
	}
}

func (OutputToDirectory) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "",
		DetailedHelp: markers.DetailedHelp{
			Summary: "outputs each artifact to the given directory, regardless of if it's package-associated or not.",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (outputToNothing) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "",
		DetailedHelp: markers.DetailedHelp{
			Summary: "skips outputting anything.",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}

func (outputToStdout) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "",
		DetailedHelp: markers.DetailedHelp{
			Summary: "outputs everything to standard-out, with no separation. ",
			Details: "Generally useful for single-artifact outputs.",
		},
		FieldHelp: map[string]markers.DetailedHelp{},
	}
}
