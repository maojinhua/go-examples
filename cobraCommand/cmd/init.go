package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	// 将 init 命令重命名成 add ，在使用的时候是用 go run main.go add 来执行
	Use: "add",
	Short: "简短介绍 init 命令",
	Long: "详细介绍 init 命令",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init cmd beging")
		fmt.Println(
			// 需要父命令里 TraverseChildren 设置为 true
			// go run main.go -s local  add  --config=myconf.yaml --viper=true -a nick -l apache
			cmd.Parent().Flags().Lookup("source").Value.String(),
			cmd.Flags().Lookup("viper").Value.String(),
			cmd.Flags().Lookup("author").Value.String(),
			cmd.Flags().Lookup("config").Value.String(),
			cmd.Flags().Lookup("license").Value.String(),
		)
		fmt.Println("init cmd end")
	},
}

func init(){
	initCmd.Flags().StringP("init_arg","i","","")
	rootCmd.AddCommand(initCmd)
}