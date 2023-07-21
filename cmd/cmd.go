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
	droneCmd.Flags().StringVarP(&DroneName, aurora.Yellow("droneName").String(), "d", "", aurora.Green(`The Drone name`).String())
	droneCmd.Flags().StringVarP(&GoPrivate, "go_private", "g", "", aurora.Green(`Go private (default "gitee.com")`).String())
	droneCmd.Flags().StringVarP(&ServiceName, "service_name", "s", "", aurora.Green(`The service name of the project`).String())
	droneCmd.Flags().StringVarP(&ServiceType, "service_type", "x", "", aurora.Green(`The service type, such as rpc | api`).String())
	droneCmd.Flags().StringVarP(&GitBranch, "gitBranch", "b", "", `The branch of the remote repo, it does work with --remote (default "master")`)
	droneCmd.Flags().StringVarP(&Registry, "registry", "r", "", `The remote git repo of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority  
    The git repo directory must be consistent with the https://github.com/zeromicro/go-zero-template directory structure`)
	droneCmd.Flags().StringVarP(&Repo, "repo", "o", "", aurora.Green(`The project git repository`).String())
	droneCmd.Flags().StringVarP(&Tag, "tag", "t", "", aurora.Green("Git tag (default \"latest\")").String())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	droneCmd.PersistentFlags().String("droneName", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// droneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
