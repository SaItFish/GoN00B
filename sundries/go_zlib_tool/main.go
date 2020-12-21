// Package main
// @file: main.go
// @description:
// @author: SaltFish
// @date: 2020/10/13
package main

import (
	"bufio"
	"compress/zlib"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"runtime"
	"runtime/pprof"
	"strings"
	"sync"
)

const (
	dataDir = "data"
	key     = "sf"
)

var (
	reg        = regexp.MustCompile(`.\.{1}sf$`)
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile 'file'")
	memprofile = flag.String("memprofile", "", "write memory profile to 'file'")
)

func press() {
	fInfos, err := ioutil.ReadDir(dataDir)
	var group sync.WaitGroup
	if err != nil {
		fmt.Println("Put files in data dir.")
		return
	}
	for _, fInfo := range fInfos {
		go func(fInfo os.FileInfo) {
			group.Add(1)
			defer group.Done()
			if fInfo.IsDir() || reg.MatchString(fInfo.Name()) {
				return
			}
			srcDir := dataDir + string(os.PathSeparator) + fInfo.Name()
			tarDir := dataDir + string(os.PathSeparator) + fInfo.Name() + ".sf"
			tarFile, err := os.OpenFile(tarDir, os.O_RDWR|os.O_CREATE, 0666)
			if err != nil {
				fmt.Printf("Open file error: %v\n", tarDir)
				return
			}
			defer tarFile.Close()

			data, _ := ioutil.ReadFile(srcDir)
			w := zlib.NewWriter(tarFile)
			_, _ = w.Write(data)
			_ = w.Flush() // 将缓存区的内容写入文件，不然文件会缺少最后的内容
		}(fInfo)
	}
	group.Wait()
}

func depress() {
	fInfos, err := ioutil.ReadDir(dataDir)
	var group sync.WaitGroup
	if err != nil {
		fmt.Println("Put files in data dir.")
		return
	}
	for _, fInfo := range fInfos {
		go func(fInfo os.FileInfo) {
			group.Add(1)
			defer group.Done()
			if fInfo.IsDir() || !reg.MatchString(fInfo.Name()) {
				return
			}
			srcDir := dataDir + string(os.PathSeparator) + fInfo.Name()
			tarDir := strings.TrimSuffix(srcDir, path.Ext(srcDir))

			srcFile, err := os.Open(srcDir)
			if err != nil {
				fmt.Printf("Open file error: %v\n", srcDir)
				return
			}
			defer srcFile.Close()

			tarFile, err := os.OpenFile(tarDir, os.O_RDWR|os.O_CREATE, 0666)
			if err != nil {
				fmt.Printf("Open file error: %v\n", tarDir)
				return
			}
			defer tarFile.Close()

			fz, err := zlib.NewReader(srcFile)
			r := bufio.NewReader(fz)
			w := bufio.NewWriter(tarFile)
			_, _ = io.Copy(w, r)
		}(fInfo)
	}
	group.Wait()
}

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	press()
	depress()

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // 获取最新的数据信息
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}
}
