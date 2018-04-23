// Copyright © 2017 Microsoft <wastore@microsoft.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"github.com/Azure/azure-storage-azcopy/common"
	"github.com/spf13/cobra"
)

func init() {
	raw := rawCopyCmdArgs{}
	// deleteCmd represents the delete command
	var deleteCmd = &cobra.Command{
		Use:        "remove",
		Aliases:    []string{"rm", "r"},
		SuggestFor: []string{"delete", "del"},
		Short:      "Coming soon: remove(rm) deletes blobs or containers in Azure Storage.",
		Long:       `Coming soon: remove(rm) deletes blobs or containers in Azure Storage.`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("remove command only takes 1 arguments. Passed %d arguments", len(args))
			}
			argsLocation := inferArgumentLocation(args[0])
			if argsLocation != argsLocation.Blob() {
				return fmt.Errorf("remove command supports delete of blob only. Passed %v as an argument", argsLocation)
			}
			raw.src = args[0]
			raw.fromTo = common.EFromTo.BlobTrash().String()
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cooked, err := raw.cook()
			if err != nil {
				return fmt.Errorf("failed to parse user input due to error %s", err)
			}

			err = cooked.process()
			if err != nil {
				return fmt.Errorf("failed to perform copy command due to error %s", err)
			}
			return nil
		},
	}
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.PersistentFlags().BoolVar(&raw.recursive, "recursive", false, "Filter: Look into sub-directories recursively when deleting from container.")
	deleteCmd.PersistentFlags().Uint8Var(&raw.logVerbosity, "Logging", uint8(pipeline.LogWarning), "defines the log verbosity to be saved to log file")
}
