package perf_tracker

import (
	"bytes"
	"fmt"
	"github.com/aman-bansal/go_perf_tracker/perf_tracker/internal"
	"go/format"
	"io/ioutil"
	"time"
)

func CalculateFunctionPerf(startTime time.Time, funcName string) {
	endTime := time.Now().Unix()
	timeElapsed := endTime - startTime.Unix()

	fmt.Printf("time elapsed in %s is %v \n", funcName, timeElapsed)
}

func GenerateCode(filePathVsFunc map[string][]string) error{
	for filePath, funcs := range filePathVsFunc {
		fileset, file, err := internal.ParseFile(filePath)
		if err != nil {
			return err
		}

		internal.FormatCode(fileset, file, funcs)
		buf := new(bytes.Buffer)
		err = format.Node(buf, fileset, file)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return nil
		}
		byt := buf.Bytes()
		err = ioutil.WriteFile(filePath, byt, 0664)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return nil
		}
	}
	return nil
}