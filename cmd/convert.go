package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hjdr4/terrayaml/lib"
	"github.com/hjdr4/yaml"
	"github.com/spf13/cobra"
)

var fullTemplate string

var convert = &cobra.Command{
	Use:   "convert",
	Short: "Converts a set of YAML documents to Terraform formats",

	RunE: func(cmd *cobra.Command, args []string) error {
		switch format {
		case "json":
		case "hcl":
		default:
			return fmt.Errorf("%v", "Format needs to be either 'json' or 'hcl'")
		}

		codes := strings.Split(code, ",")
		sCode, err := lib.LoadData(codes)
		if err != nil {
			return err
		}
		sData := fullTemplate + "\n" + sCode

		//Render the code
		m := lib.MapSlice{}
		err = yaml.Unmarshal([]byte(sData), &m)
		if err != nil {
			return err
		}

		//Filter out the templates
		for i := len(m) - 1; i >= 0; i-- {
			v := m[i]
			if strings.Contains(v.Key.(string), "_template") {
				m = append(m[:i], m[i+1:]...)
			}
		}

		//Render into JSON
		b, err := json.MarshalIndent(m, "", "  ")
		if err != nil {
			return err
		}

		switch format {
		case "json":
			fmt.Println(string(b))
		case "hcl":
			hcl, err := lib.ToHCL(string(b))
			if err != nil {
				return err
			}
			fmt.Println(hcl)
		default:
			hcl, err := lib.ToHCL(string(b))
			if err != nil {
				return err
			}
			fmt.Println(hcl)
		}
		return nil
	},
}

var format string
var code string

func init() {
	convert.PersistentFlags().StringVarP(&format, "format", "f", "", "define the output format ('json' or 'hcl')")
	convert.MarkPersistentFlagRequired("format")
	convert.PersistentFlags().StringVarP(&code, "code", "c", "", "YAML code paths, separated by commas. They can be either directories containing '.yml' files or YAML files. Files are loaded in path order then lexicographical order")
	convert.MarkPersistentFlagRequired("code")
}
