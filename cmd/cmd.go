package cmd

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"os"
)

// droneCmd represents the drone command
var droneCmd = &cobra.Command{
	Use:   "drone",
	Short: "drone is very good",
	Long:  `创建drone的指令`,
	RunE:  DroneGenerator, //步骤一
}

var (
	//步骤三
	DroneName   string
	GoPrivate   string
	ServiceName string
	ServiceType string
	GitBranch   string
	Registry    string
	Repo        string
	Tag         string
)

func Execute() {
	//os.Args = supportGoStdFlag(os.Args)
	if err := droneCmd.Execute(); err != nil {
		fmt.Println(aurora.Red(err.Error()))
		os.Exit(1)
	}
}

func init() {
	//droneCmd.AddCommand(droneCmd)

	// drone --droneName="base" --go_private="gitee.com" --service_name="baserpc.go" --service_type="rpc" --gitBranch="master" --registry="registry.cn-beijing.aliyuncs.com" --repo="registry.cn-beijing.aliyuncs.com/ctra_test/ctra-go-zhiye-rpc" --tag="latest"
	// drone -d="base" -g="gitee.com" -s="baserpc.go" -x="rpc" -b="master" -r="registry.cn-beijing.aliyuncs.com" -o="registry.cn-beijing.aliyuncs.com/ctra_test/ctra-go-zhiye-rpc" -t="latest"

	// 步骤二
	droneCmd.Flags().StringVarP(&DroneName, "droneName", "d", "", `.drone yaml name`)
	droneCmd.Flags().StringVarP(&GoPrivate, "go_private", "g", "", `switch git repo match your goprivate`)
	droneCmd.Flags().StringVarP(&ServiceName, "service_name", "s", "", `swagger save file name`)
	droneCmd.Flags().StringVarP(&ServiceType, "service_type", "x", "", `swagger save file name`)
	droneCmd.Flags().StringVarP(&GitBranch, "gitBranch", "b", "", `swagger save file name`)
	droneCmd.Flags().StringVarP(&Registry, "registry", "r", "", `swagger save file name`)
	droneCmd.Flags().StringVarP(&Repo, "repo", "o", "", `swagger save file name`)
	droneCmd.Flags().StringVarP(&Tag, "tag", "t", "", `swagger save file name`)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	droneCmd.PersistentFlags().String("droneName", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// droneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
