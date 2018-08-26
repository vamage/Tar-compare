package main

import (
	"archive/tar"
	"compress/gzip"
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
)


func main() {
	source := flag.String("source", "", "Tar file to compare from")
	destination := flag.String("destination", "", "Destination path ")
	purgeurl := flag.String("purge", "http://127.0.0.1/purge.php?file=%s", "Url to send invalidate commands to ")
	flag.Parse()
	var todelete = make([]string, 0, 40)
	realdestination,_ := filepath.Abs(* destination)
	f := filesystem(realdestination)
	realpath := filepath.Dir(realdestination)
	t, updated := readtar(source, f, realpath)
	var diff = map[string]string{}

	for k, _ := range f {
		if _, exist := t[k]; !exist {
			todelete = append(todelete, k)
		}
	}


	for k, v := range diff {
		fmt.Printf(" %s  : %s \n", k, v)
	}

	for _, file := range updated {
		realfile := fmt.Sprintf("%s/%s", realpath, file)
		err := os.Rename(fmt.Sprintf("%s.%x", realfile, t[file]), realfile)
		if err != nil {
			panic(err)
		}
	}
	purgefiles(* purgeurl,realpath,updated)
	//files need to exist to be marked for recompiling
	purgefiles(* purgeurl,realpath,todelete)
	for _, file := range todelete {
		err := os.Remove(fmt.Sprintf("%s/%s", realpath, file))
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf("%s , %s, %d", *source, *destination, len(t))
	eq := reflect.DeepEqual(t, f)
	fmt.Printf("The hashes %t\n", eq)
	if eq {
		os.Exit(0)
	}else {
		os.Exit(12)
	}
}

func readtar(intar *string, compare map[string][16]byte, folder string) (map[string][16]byte, []string) {
	var m = map[string][16]byte{}
	var updated = make([]string, 0, 40)

	r, err := os.Open(*intar)
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
			return m, updated

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
			err = os.MkdirAll(target, os.FileMode(header.Mode))
			if err != nil {
				panic(err)
			}

		// if it's a file create it
		case tar.TypeReg:
			file, _ := ioutil.ReadAll(tr)
			hash := md5.Sum(file)
			m[target] = hash
			if compare[target] != hash {
				writetempfile(file, os.FileMode(header.Mode), hash, target, folder)
				updated = append(updated, target)

			}
		}
	}
}

func filesystem(folder string) map[string][16]byte {
	var m = map[string][16]byte{}
	realpath := filepath.Dir(folder)
	if _, err := os.Stat( folder); os.IsNotExist(err) {
		return nil
	}

	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", folder, err)
			return err
		}
		//Symlinks and directories can't be hashed.
		if !info.IsDir() && info.Mode().IsRegular() {
			r, _ := os.Open(path)
			defer r.Close()
			file, _ := ioutil.ReadAll(r)
			hash := md5.Sum(file)
			m[path[len(realpath)+1:]] = hash
		}

		return nil
	})

	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", folder, err)
	}
	return m
}

func writetempfile(data []byte, mode os.FileMode, hash [16]byte, path string, desination string) {
	filename := fmt.Sprintf("%s/%s.%x", desination, path, hash)
	err := ioutil.WriteFile(filename, data, mode)
	if err != nil {
		panic(err)
	}
}

func purgefiles(prugeurl string,path string,updated[] string)  {
	for _, file := range updated {
		purgefile :=fmt.Sprintf("%s/%s", path, file)
		fmt.Printf(prugeurl,url.PathEscape(purgefile))
		fmt.Println()
	}
}

