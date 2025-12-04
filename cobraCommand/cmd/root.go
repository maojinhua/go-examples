package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:"root" ,
	// 简短描述，
	Short: "short desc",
	Long: "详细描述",
	// Run 指定命令的处理函数
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root cmd run beging")
		// 打印 flag
		fmt.Println(
			cmd.Flags().Lookup("source").Value.String(),
			cmd.Flags().Lookup("viper").Value.String(),
			cmd.PersistentFlags().Lookup("author").Value.String(),
			cmd.PersistentFlags().Lookup("config").Value.String(),
			cmd.PersistentFlags().Lookup("license").Value.String(),
		)
		fmt.Println("--------------------------------")
		// 打印 viper 的值
		fmt.Println(
			viper.Get("author"),
			viper.Get("license"),
		)
		fmt.Println("root cmd run end")
	},
	// 子命令也可以访问 Flags 本地标志的值
	TraverseChildren: true,
}

func Execute(){
	rootCmd.Execute()
}

var cfgFile string
var userLicense string
// 定义根命令上的命令
func init(){
	cobra.OnInitialize(initConfig)
	//  定义持久化标志，持久化标志可以传递给他的子命令
	rootCmd.PersistentFlags().Bool("viper",true,"")
	// 指定 flag 缩写.带 P 的是定义缩写， -a 和 --author
	rootCmd.PersistentFlags().StringP("author","a","YOUR NAME","")
	// 通过指针，将值定义到字段. 带 Var 的是指定一个变量接收命令参数
	rootCmd.PersistentFlags().StringVar(&cfgFile,"config","","")
	rootCmd.PersistentFlags().StringVarP(&userLicense,"license","l","","")
	// 本地标志，只能在当前命令使用
	rootCmd.Flags().StringP("source","s","","")

	// 配置绑定，配置绑定可以动态改变 viper 的值，
	// 如果构建参数里有设置值则优先级最高，如果指定配置文件则会使用配置文件里的值，
	// 如果构建参数里有值，并且也有指定配置文件，则以构建参数里的值为准
	// 如果都没有则使用默认值
	// viper 内容的优先级
	// 1. 命令行参数
	// 2. 环境变量
	// 3. 配置文件
	// 4. 默认值
	// 例子：构建参数 -a "build author" -l "build license"，配置文件 myconf.yaml 里的 author 和 license 的值为 "config author" 和 "config license"，则最终的值为 "build author" 和 "build license"
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("license", rootCmd.PersistentFlags().Lookup("license"))
	viper.SetDefault("author", "default author")
	viper.SetDefault("license", "default license")
}

// 初始化配置文件
func initConfig(){
	if cfgFile!=""{
		viper.SetConfigFile(cfgFile)
	}else{
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}
	// 检查环境变量，将配置的键值加载到 viper 中
	viper.AutomaticEnv()
	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)
	}
	fmt.Println("using config file:", viper.ConfigFileUsed())
}