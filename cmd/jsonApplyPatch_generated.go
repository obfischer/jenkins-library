// Code generated by piper's step-generator. DO NOT EDIT.

package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/SAP/jenkins-library/pkg/config"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/SAP/jenkins-library/pkg/telemetry"
	"github.com/spf13/cobra"
)

type jsonApplyPatchOptions struct {
	Input  string `json:"input,omitempty"`
	Patch  string `json:"patch,omitempty"`
	Output string `json:"output,omitempty"`
}

// JsonApplyPatchCommand Patches a json with a patch file
func JsonApplyPatchCommand() *cobra.Command {
	const STEP_NAME = "jsonApplyPatch"

	metadata := jsonApplyPatchMetadata()
	var stepConfig jsonApplyPatchOptions
	var startTime time.Time

	var createJsonApplyPatchCmd = &cobra.Command{
		Use:   STEP_NAME,
		Short: "Patches a json with a patch file",
		Long: `This steps patches a json file with patch file using the json patch standard.
This step can, e.g., be used if there is a json schema which needs to be patched.`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			startTime = time.Now()
			log.SetStepName(STEP_NAME)
			log.SetVerbose(GeneralConfig.Verbose)

			path, _ := os.Getwd()
			fatalHook := &log.FatalHook{CorrelationID: GeneralConfig.CorrelationID, Path: path}
			log.RegisterHook(fatalHook)

			err := PrepareConfig(cmd, &metadata, STEP_NAME, &stepConfig, config.OpenPiperFile)
			if err != nil {
				return err
			}

			if len(GeneralConfig.HookConfig.SentryConfig.Dsn) > 0 {
				sentryHook := log.NewSentryHook(GeneralConfig.HookConfig.SentryConfig.Dsn, GeneralConfig.CorrelationID)
				log.RegisterHook(&sentryHook)
			}

			return nil
		},
		Run: func(_ *cobra.Command, _ []string) {
			telemetryData := telemetry.CustomData{}
			telemetryData.ErrorCode = "1"
			handler := func() {
				telemetryData.Duration = fmt.Sprintf("%v", time.Since(startTime).Milliseconds())
				telemetry.Send(&telemetryData)
			}
			log.DeferExitHandler(handler)
			defer handler()
			telemetry.Initialize(GeneralConfig.NoTelemetry, STEP_NAME)
			jsonApplyPatch(stepConfig, &telemetryData)
			telemetryData.ErrorCode = "0"
			log.Entry().Info("SUCCESS")
		},
	}

	addJsonApplyPatchFlags(createJsonApplyPatchCmd, &stepConfig)
	return createJsonApplyPatchCmd
}

func addJsonApplyPatchFlags(cmd *cobra.Command, stepConfig *jsonApplyPatchOptions) {
	cmd.Flags().StringVar(&stepConfig.Input, "input", os.Getenv("PIPER_input"), "File path to the json file which schould be patched.")
	cmd.Flags().StringVar(&stepConfig.Patch, "patch", os.Getenv("PIPER_patch"), "File path to the patch which should be applied to the json file.")
	cmd.Flags().StringVar(&stepConfig.Output, "output", os.Getenv("PIPER_output"), "File path to destination of the patched json file.")

	cmd.MarkFlagRequired("input")
	cmd.MarkFlagRequired("patch")
	cmd.MarkFlagRequired("output")
}

// retrieve step metadata
func jsonApplyPatchMetadata() config.StepData {
	var theMetaData = config.StepData{
		Metadata: config.StepMetadata{
			Name:    "jsonApplyPatch",
			Aliases: []config.Alias{},
		},
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Parameters: []config.StepParameters{
					{
						Name:        "input",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "patch",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "output",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
				},
			},
		},
	}
	return theMetaData
}
