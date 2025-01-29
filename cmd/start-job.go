// Copyright 2025 "Google LLC"
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

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startJob)
}

// type PersistentVolumeClaim struct {
// 	ClaimName string `yaml:"claimName"`
// 	MountPath string `yaml:"mountPath"`
// }

// type ClusterTrainingConfig struct {
// 	name          string
// 	nodes         int
// 	nTasksPerNode int
// }

// type Cluster struct {
// 	clusterConfig string
// 	instanceType  string
// }

// type Config struct {
// 	Container             string
// 	Cluster               Cluster
// 	clusterTrainingConfig ClusterTrainingConfig
// 	envVrs                interface{}
// 	baseResultsDir        string
// }

// cli should validate that the file exists and output the command to be run
func addStartJobFlags(cmd *cobra.Command) *cobra.Command {
	cmd.Flags().StringP("recipe", "r", "", "The recipe to use for the deployment (required)")
	cmd.MarkFlagRequired("recipe") // Make --recipe required

	return cmd
}

var (
	startJob = addStartJobFlags(&cobra.Command{
		Use:   "start-job --recipe <RECIPE-PATH>",
		Short: "Creates a new training job.",
		Long:  "Create a new training job based on a provided blueprint.",
		Run:   runStartJobCmd,
	})
)

func validateRecipePath(path string) error { // No change here
	if _, err := os.Lstat(path); err != nil {
		return fmt.Errorf("%q does not exist", path)
	}
	return nil
}

func runStartJobCmd(cmd *cobra.Command, args []string) {
	recipe, err := cmd.Flags().GetString("recipe") // Get recipe from flag
	if err != nil {
		fmt.Println("Error getting recipe:", err)
		return
	}
	if err := validateRecipePath(recipe); err == nil {
		recipe, err := cmd.Flags().GetString("recipe")
		if err != nil {
			fmt.Println("error getting recipes:", err)
		}
		// pvClaims, _ := cmd.Flags().GetStringSlice("persistent-volume-claims")
		// overrideParams, _ := cmd.Flags().GetString("override-parameters")

		// ... use recipe, pvClaims, and overrideParams in your doCreate function ...
		doStartJob(recipe) // Modified doCreate signature
		// ... rest of the function (logging and instructions) ...
	} else {
		fmt.Println("Recipe does not exist", err)
	}

}

// Update doCreate to accept the new parameters
func doStartJob(recipe string) {
	fmt.Println("Recipe:", recipe)
}

// func doRunJobCmd(recipe string,
// 	overrideParameters string,
// 	stringjobName string,
// 	configFile string,
// 	volumes string,
// 	persistentVolumeClaims string,
// 	autoResume bool,
// 	labelSelector string,
// 	maxRetry int) {
// 	// list all the aprams for the job

// }
