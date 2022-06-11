package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	yyml "github.com/ghodss/yaml"
	sanitizederror "github.com/kyverno/kyverno/cmd/cli/kubectl-kyverno/utils/sanitizedError"
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

	// JSON file path
	paths := "./Play_with_JSON/JSON_into_YAML/value.json"

	// Reading JSON file --> Converted into JSONBytes
	valueInJsonBytes, err := ioutil.ReadFile(paths)
	if err != nil {
		errors = append(errors, sanitizederror.NewWithError("unable to read yaml", err))
	}

	// Convering directly JSON bytes into YAML bytes using liabrary yyml
	valuesInYamlBytes, err := yyml.JSONToYAML(valueInJsonBytes)
	if err != nil {
		errors = append(errors, sanitizederror.NewWithError("fail to convert JSON to YAML", err))
	}

	// Converting YAMLBytes into YAMLString
	valueInYAML := string(valuesInYamlBytes)
	fmt.Println("Converted into YAML fro JSON: ", valueInYAML)

	// Write YAML Value into the file, kyverno-test.yaml
	err = ioutil.WriteFile("./Play_with_JSON/JSON_into_YAML/kyverno-test.yaml", valuesInYamlBytes, 0644)
	if err != nil {
		errors = append(errors, sanitizederror.NewWithError("failed to Write file", err))
	}

	// Now store the JSON value in struct type..
	values := &Test{}
	if err = json.Unmarshal(valueInJsonBytes, values); err != nil {
		errors = append(errors, sanitizederror.NewWithError("failed to decode yaml", err))
	}

	fmt.Println("Value of struct type Test: ", values)

}
