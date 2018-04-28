package converter

import "github.com/flosell/hclencoder"

type HCLCondition struct {
	Test string `hcl:"test"`
	Variable string `hcl:"variable"`
	Values []string `hcl:"values"`
}

type HCLStatement struct {
	Sid    string `hcl:"sid"`
	Effect string `hcl:"effect"`
	Resources []string `hcl:"resources"`
	NotResources []string `hcl:"not_resources"`
	Actions []string `hcl:"actions"`
	NotActions []string `hcl:"not_actions"`
	Conditions []HCLCondition `hcl:"condition,squash"`
}

type HCLDataSource struct {
	Type       string         `hcl:",key"`
	Name       string         `hcl:",key"`
	Statements []HCLStatement `hcl:"statement,squash"`
}

type HclWholeFile struct {
	Entry HCLDataSource `hcl:"data"`
}

func Encode(dataSource HCLDataSource) (string, error) {
	b, err := hclencoder.Encode(HclWholeFile{Entry: dataSource})

	if err != nil {
		return "", err
	} else {
		return string(b), nil
	}
}
