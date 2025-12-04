package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// 自定义参数检查
var cusArgsCheckCmd = &cobra.Command{
	Use:"cusargs",
	
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args)<1{
			return errors.New("至少输入一个参数")
		}
		if len(args)>2{
			return errors.New("最多输入两个参数")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cusargs cmd run beging")
		fmt.Println("args:",args)
		fmt.Println("cusargs cmd run end")
	},
}

// 内置参数检查
var argsCheckCmd = &cobra.Command{
	Use:"args",
	// Args: cobra.MatchAll(cobra.MinimumNArgs(1),cobra.MaximumNArgs(2)),
	// 限制参数选择范围
	Args: cobra.OnlyValidArgs,
	ValidArgs: []string{"123","abc","345"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("args cmd run beging")
		fmt.Println("args:",args)
		fmt.Println("args cmd run end")
	},
}

func init(){
	rootCmd.AddCommand(cusArgsCheckCmd)
	rootCmd.AddCommand(argsCheckCmd)
}