package cmd

import (
	_ "embed"
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"os"
	"strings"
	"text/template"
)

var (
	//go:embed drone.tpl
	UsageTpl string
)

type Drone struct {
	//步骤三
	DroneName   string
	GoPrivate   string
	ServiceName string
	ServiceType string
	GitBranch   string
	Registry    string
	Repo        string
	Tag         string
}

func DroneGenerator(_ *cobra.Command, _ []string) error {

	// 步骤四
	// 对所有的传入的参数进行一一判断
	dronename := DroneName
	if len(dronename) == 0 {
		dronename = "dronegen-greet"
	}

	goprivate := GoPrivate
	fmt.Println(len(strings.Split(goprivate, ".")))
	if len(strings.Split(goprivate, ".")) <= 1 {
		return fmt.Errorf("error go private!")
	}

	serviceName := ServiceName
	serviceType := ServiceType
	gitBranch := GitBranch
	registry := Registry
	repo := Repo
	tag := Tag

	file, err := os.Create("drone.yml")
	if err != nil {
		fmt.Println("文件创建失败:", err)
		return err
	} else {
		fmt.Println("文件创建成功!")
	}

	defer file.Close()

	text, err := pathx.LoadTemplate("dronegen", "drone.tpl", UsageTpl)
	if err != nil {
		fmt.Println("打开模板失败:", err)
		return err
	} else {
		fmt.Println("打开模板成功!")
	}

	t := template.Must(template.New("dronegen").Parse(text))
	return t.Execute(file, Drone{
		DroneName:   dronename,
		GoPrivate:   goprivate,
		ServiceName: serviceName,
		ServiceType: serviceType,
		GitBranch:   gitBranch,
		Registry:    registry,
		Repo:        repo,
		Tag:         tag,
	})
	fmt.Println(aurora.Green("Done."))

	return nil
}
