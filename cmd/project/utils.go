package project

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/yourusername/airfoil/api"
)

func contains(input string, choices []string) bool {
	for _, choice := range choices {
		if input == choice {
			return true
		}
	}
	return false
}

func promptChoice(message string, choices []string, defaultChoice string) string {
	var selection string = ""
	for !contains(selection, choices) {
		selection = ""
		fmt.Println(message)
		fmt.Print("   Available options: ")
		for _, choice := range choices {
			fmt.Printf("%s", choice)
			if choice == defaultChoice {
				fmt.Print(" (default)")
			}
			if choice != choices[len(choices)-1] {
				fmt.Print(", ")
			}
		}

		fmt.Print("\n   > ")

		fmt.Scanln(&selection)

		if selection == "" {
			return defaultChoice
		}
	}
	return selection
}

func selectNetworkVolume() (networkVolumeId string, err error) {
	networkVolumes, err := api.GetNetworkVolumes()
	if err != nil {
		fmt.Println("Error fetching network volumes:", err)
		return "", err
	}
	if len(networkVolumes) == 0 {
		fmt.Println("No network volumes found. Please create one and try again. (https://runpod.io/console/user/storage)")
		return "", fmt.Errorf("no network volumes found")
	}

	promptTemplates := &promptui.SelectTemplates{
		Label:    inputPromptPrefix + "{{ . }}",
		Active:   ` {{ "●" | cyan }} {{ .Name | cyan }}`,
		Inactive: `   {{ .Name | white }}`,
		Selected: `   {{ .Name | white }}`,
	}

	options := []NetVolOption{}
	for _, networkVolume := range networkVolumes {
		options = append(options, NetVolOption{Name: fmt.Sprintf("%s: %s (%d GB, %s)", networkVolume.Id, networkVolume.Name, networkVolume.Size, networkVolume.DataCenterId), Value: networkVolume.Id})
	}
	getNetworkVolume := promptui.Select{
		Label:     "Select a Network Volume:",
		Items:     options,
		Templates: promptTemplates,
	}
	i, _, err := getNetworkVolume.Run()
	if err != nil {
		return "", err
	}
	networkVolumeId = options[i].Value
	return networkVolumeId, nil
}

func selectStarterTemplate() (template string, err error) {
	type StarterTemplateOption struct {
		Name  string
		Value string
	}
	templates, err := starterTemplates.ReadDir("starter_examples")
	if err != nil {
		fmt.Println("Something went wrong trying to fetch the starter project.")
		fmt.Println(err)
		return "", err
	}
	promptTemplates := &promptui.SelectTemplates{
		Label:    inputPromptPrefix + "{{ . }}",
		Active:   ` {{ "●" | cyan }} {{ .Name | cyan }}`,
		Inactive: `   {{ .Name | white }}`,
		Selected: `   {{ .Name | white }}`,
	}
	options := []StarterTemplateOption{}
	for _, template := range templates {
		name := template.Name()
		name = strings.Replace(name, "_", " ", -1)
		options = append(options, StarterTemplateOption{Name: name, Value: template.Name()})
	}
	getStarterTemplate := promptui.Select{
		Label:     "Select a Starter Project:",
		Items:     options,
		Templates: promptTemplates,
	}
	i, _, err := getStarterTemplate.Run()
	if err != nil {
		return "", err
	}
	template = options[i].Value
	return template, nil
}

type NetVolOption struct {
	Name  string
	Value string
}
