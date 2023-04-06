package cmd

import (
	"errors"
	"log"
	db "oss-cli/minio"
	"oss-cli/utils"
	"sync"

	"github.com/spf13/cobra"
)

var (
	maxGoroutine int

	inputPath  string
	outputPath string

	rootCMD = &cobra.Command{
		Use:   "upload",
		Short: "上传文件",
	}
)

func uploadDirCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dir",
		Short: "上传文件夹",
		Run:   UploadDir,
	}

	cmd.Flags().IntVarP(&maxGoroutine, "maxGoroutine", "m", 10, "最大并发数")
	cmd.Flags().StringVarP(&inputPath, "inputPath", "i", "", "输入文件路径")
	cmd.Flags().StringVarP(&outputPath, "outputPath", "o", "", "输出文件路径")

	cmd.MarkFlagRequired("inputPath")
	cmd.MarkFlagRequired("outputPath")

	return cmd
}

func init() {
	rootCMD.AddCommand(uploadDirCmd())
}

func uploadCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upload",
		Short: "上传文件",
		Run:   UploadFile,
	}

	cmd.Flags().IntVarP(&maxGoroutine, "maxGoroutine", "m", 10, "最大并发数")
	cmd.Flags().StringVarP(&inputPath, "inputPath", "i", "", "输入文件路径")
	cmd.Flags().StringVarP(&outputPath, "outputPath", "o", "", "输出文件路径")

	cmd.MarkFlagRequired("inputPath")
	cmd.MarkFlagRequired("outputPath")

	return cmd
}

// create md5 dir
// split file
// upload to minio
func UploadFile(cmd *cobra.Command, args []string) {
	/*if exist := utils.FileIsExist(inputPath); !exist {
		log.Fatal(errors.New("input file is not exist"))
	}
	if exist := utils.DirIsExist(outputPath); !exist {
		log.Fatal(errors.New("output dir is not exist"))
	}*/
}

func UploadDir(cmd *cobra.Command, args []string) {
	if exist := utils.DirIsExist(inputPath); !exist {
		log.Fatal(errors.New("input dir is not exist"))
	}
	minioSession := db.MinioClient()

	files := utils.ListShow("./")

	for _, v := range files {
		utils.Funcs = append(utils.Funcs, utils.NewFuncHandler(minioSession.PutObject, "./"+v, "videos", "go/"+v, map[string]string{}))
	}

	var wg sync.WaitGroup
	for _, v := range utils.Funcs {
		wg.Add(1)
		go func(v *utils.FuncHandler) {
			_, err := v.Call()
			if err != nil {
				return
			}
			defer wg.Done()
		}(v)
	}
	wg.Wait()
}

func Execute() error {
	return rootCMD.Execute()
}
