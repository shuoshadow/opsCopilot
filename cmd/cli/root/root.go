package root

import (
	"fmt"
	"os"

	"github.com/shaowenchen/ops/cmd/cli/host"
	"github.com/shaowenchen/ops/cmd/cli/kubernetes"
	"github.com/shaowenchen/ops/cmd/cli/pipeline"
	"github.com/shaowenchen/ops/cmd/cli/storage"
	"github.com/shaowenchen/ops/cmd/cli/upgrade"
	"github.com/shaowenchen/ops/cmd/cli/version"
	"github.com/shaowenchen/ops/pkg/constants"
	"github.com/shaowenchen/ops/pkg/utils"
	"github.com/spf13/cobra"
)

func Execute() {
	RootCmd.AddCommand(host.HostCmd)
	RootCmd.AddCommand(kubernetes.KubernetesCmd)
	RootCmd.AddCommand(storage.StorageCmd)
	RootCmd.AddCommand(pipeline.PipelineCmd)
	RootCmd.AddCommand(version.VersionCmd)
	RootCmd.AddCommand(upgrade.UpgradeCmd)
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var RootCmd = &cobra.Command{
	Use:   "opscli",
	Short: "a cli tool",
	Long:  `This is a cli tool for ops.`,
}

func init() {
	utils.CreateDir(constants.GetOpsLogsDir())
}