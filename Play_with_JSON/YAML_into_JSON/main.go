package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	sanitizederror "github.com/kyverno/kyverno/cmd/cli/kubectl-kyverno/utils/sanitizedError"
	"k8s.io/apimachinery/pkg/util/yaml"
)

type Test struct {
	Name      string        `json:"name"`
	Policies  []string      `json:"policies"`
	Resources []string      `json:"resources"`
	Results   []TestResults `json:"results"`
}

type TestResults struct {
	// Policy -->  name of the policy.
	Policy string `json:"policy"`

	// Rule --> name of the rule in the policy.
	Rule string `json:"rule"`

	// Resource mentions the name of the resource on which the policy is to be applied.
	Resource string `json:"resource"`

	// Kind --> type of Resource
	Kind string `json:"kind"`

	// Namespace of the Resource
	Namespace string `json:"namespace"`

	// Result
	Result string `json:"result"`
}

func main() {
	var errors []error

	// path of the file
	path := "./Play_with_JSON/YAML_into_JSON/kyverno-test.yaml"
	fmt.Println("path: ", path)

	// Get the file, currently its in a byte code + in YAML form.
	valueInYamlBytes, err := ioutil.ReadFile(path)
	if err != nil {
		errors = append(errors, sanitizederror.NewWithError("unable to read yaml", err))
	}

	// Now valueInYamlBytes coverted into string
	valueInYAML := string(valueInYamlBytes)
	fmt.Println("Value in YAML format: : ", valueInYAML)

	// Converts YAML bytes into JSON bytes
	valuesInJsonBytes, err := yaml.ToJSON(valueInYamlBytes)
	if err != nil {
		errors = append(errors, sanitizederror.NewWithError("unable to read yaml", err))
	}

	// Now valuesInJsonBytes coverted into string
	valueInJson := string(valuesInJsonBytes)
	fmt.Println("value in JSON format: ", valueInJson)

	// Now convert this JSON value into Struct of type Test
	values := &Test{}

	if err := json.Unmarshal(valuesInJsonBytes, values); err != nil {
		errors = append(errors, sanitizederror.NewWithError("failed to decode yaml", err))
	}

	fmt.Println("Value of struct type Test: ", values)

	// Now write this JSON into some file, say value.json
	err = ioutil.WriteFile("./Play_with_JSON/YAML_into_JSON/value.json", valuesInJsonBytes, 0644)
	if err != nil {
		errors = append(errors, sanitizederror.NewWithError("failed to Write file", err))
	}

}
