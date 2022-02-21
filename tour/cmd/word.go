package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strings"
	"tour/internal/word"
)

const (
	MODE_UPPER = iota + 1 // 全部转为大写
	MODE_LOWER
	MODE_UNDERSCORE_TO_UPPER_CAMELCASE
	MODE_UNDERSCORE_TO_LOWER_CAMELCASE
	MODE_CAMELCASE_TO_UNDERSCORE
)

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换",
	"1：全部单词转为大写",
	"2：全部单词转为小写",
	"3：下划线单词转为大写驼峰单词",
	"4: 下划线单词转为小写驼峰单词",
	"5： 驼峰单词转为下划线单词",
}, "\n")

var mode int8
var str string

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case MODE_UPPER:
			content = word.ToUpper(str)
		case MODE_LOWER:
			content = word.ToLower(str)
		case MODE_UNDERSCORE_TO_UPPER_CAMELCASE:
			content = word.UnderscoreToUpperCamelCase(str)
		case MODE_UNDERSCORE_TO_LOWER_CAMELCASE:
			content = word.UnderscoreToLowerCameCase(str)
		case MODE_CAMELCASE_TO_UNDERSCORE:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该格式转换，请知悉help word 查看帮助")
		}

		log.Printf("输出结果 %s \n", content)

	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入待转换单词")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}
