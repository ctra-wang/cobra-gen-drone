package cmd

import (
	_ "embed"
	"encoding/json"
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

	//步骤四
	//自己填充模版逻辑

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

	fmt.Println(serviceName)
	fmt.Println(serviceType)
	fmt.Println(gitBranch)
	fmt.Println(registry)
	fmt.Println(repo)
	fmt.Println(tag)

	out, err := os.Create(".drone.yml")
	if err != nil {
		return err
	}

	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			fmt.Println(aurora.Green(err))
		}
	}(out)

	text, err := pathx.LoadTemplate("dronegen", "drone.tpl", UsageTpl)
	if err != nil {
		return err
	}

	marshal, err := json.Marshal(text)
	if err != nil {
		return err
	}
	fmt.Println(string(marshal))

	t := template.Must(template.New("dronegen").Parse(text))
	t.Execute(out, Drone{
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
	//os.WriteFile("test2.txt", []byte("hello,Jessica"), os.ModePerm)

	return nil
}
