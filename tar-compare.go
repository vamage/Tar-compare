package main

import (
	"archive/tar"
	"compress/gzip"
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)


func main() {
	source := flag.String("source", "", "Tar file to compare from")
	cache := flag.String("cache", "", "cache path ")
	deploy     := flag.String("deploy", "1", "deploy path ")
	flag.Parse()
	err := os.MkdirAll(*cache, 0755)
	if err != nil {
		panic(err)
	}
	cachepath,_  :=  filepath.Abs(* cache)
	deploypath :=  * deploy
	fmt.Printf("1: %s \n 2:%s \n", cachepath,deploypath)
	_ = readtar(* source,cachepath,deploypath)

}

func readtar(intar string, cachepath string,deploypath string ) (map[string][16]byte) {
	var m = map[string][16]byte{}

	r, err := os.Open(intar)
	gzr, err := gzip.NewReader(r)
	if err != nil {
		fmt.Printf("couldn't open %s", intar)
		panic("failed")
	}
	defer gzr.Close()
	tr := tar.NewReader(gzr)
	for {
		header, err := tr.Next()

		switch {

		// if no more files are found return
		case err == io.EOF:
			return m

			// return any other error
		case err != nil:
			panic(err)

			// if the header is nil, just skip it (not sure how this happens)
		case header == nil:
			continue
		}

		// the target location where the dir/file should be created
		target := header.Name

		// the following switch could also be done using fi.Mode(), not sure if there
		// a benefit of using one vs. the other.
		// fi := header.FileInfo()

		// check the file type
		switch header.Typeflag {
		case tar.TypeDir:
			//ensure directory exist
			cachefolder := fmt.Sprintf("%s/%s", cachepath, target)
			deployfolder  := fmt.Sprintf("%s/%s", deploypath, target)
			err = os.MkdirAll(cachefolder, os.FileMode(header.Mode))
			if err != nil {
				panic(err)
			}
			err = os.MkdirAll(deployfolder, os.FileMode(header.Mode))
			if err != nil {
				panic(err)
			}
		// if it's a file create it
		case tar.TypeReg:
			file, _ := ioutil.ReadAll(tr)
			hash := md5.Sum(file)
			m[target] = hash
				writecachefile(file, os.FileMode(header.Mode), hash, target, cachepath,deploypath)

		}
	}
}



func writecachefile(data []byte, mode os.FileMode, hash [16]byte, path string, cachepath string, destinationpath string) {
	tempfilename := fmt.Sprintf("%s/%s.%x", cachepath, path, hash)
	destinationfile := fmt.Sprintf("%s/%s.%x", destinationpath, path, hash)
	if _, err := os.Stat(tempfilename); os.IsNotExist(err) {

		err := ioutil.WriteFile(tempfilename, data, mode)
		if err != nil {
			panic(err)
		}
	}
	if realpath,_ := os.Readlink(destinationfile); realpath!= tempfilename {
		os.Remove(destinationfile)

		err := os.Symlink(tempfilename, destinationfile)
		if err != nil {
			panic(err)
		}
	}
}



