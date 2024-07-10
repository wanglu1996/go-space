package main

import (
	"github.com/spf13/cobra"
)

func newApp() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "parseYaraRule",
		Short: "格式化yara规则",
		Long:  "使用该工具格式化yara规则。\n  使用方式： parseYaraRule -s [sourcfile] -d [dstfile]\n1.去除规制标识符的中文 \n2.对于规则标识符长于128字符的进行替换",
		Args:  cobra.NoArgs,
		RunE:  parseYaraRule,
	}
	rootCmd.Flags().StringP("src", "s", "", "源文件")
	rootCmd.Flags().StringP("dst", "d", "", "输出文件")

	return rootCmd
}

func parseYaraRule(cmd *cobra.Command, args []string) error {
	srcFile := cmd.Flag("src").Value.String()
	dstFile := cmd.Flag("dst").Value.String()
	parse(srcFile, dstFile)
	return nil
}
