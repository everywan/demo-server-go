package cmd

import (
	"github.com/spf13/cobra"
)

var demoTaskCmd = &cobra.Command{
	Use:   "demo_task",
	Short: "脚本任务",
	Long:  `二级命令, 用于启动一个脚本任务.`,
	Run: func(cmd *cobra.Command, args []string) {
		// var err error
		// //opts, err := loadOptions()
		// //handleInitError("load_options", err)
		// resource := &rules.Resource{}
		// ruleFactory := rules.NewFactory(resource)

		// for _, name := range ruleFactory.GetAllRuleNames() {
		// 	rule, _ := ruleFactory.Get(name)
		// 	err = rule.MakeEffective()
		// 	if err != nil {
		// 		log.Fatalf("error in exec rule make effective. error: %v", err)
		// 	}
		// }
		// log.Println("done")
	},
}

func init() {
	rootCmd.AddCommand(demoTaskCmd)
}
