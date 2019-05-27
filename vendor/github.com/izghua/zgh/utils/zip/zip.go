/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-22
 * Time: 00:47
 */
package zip

import (
	"archive/zip"
	"github.com/izghua/zgh"
	"io"
	"os"
)

// compress the files
// it is not designed by my
// i copy it what from https://studygolang.com/articles/7471
//files file arrays，it can be the diff file or directory
//dest be used to the addr
func Compress(files []*os.File, dest string) error {
	d, _ := os.Create(dest)
	defer d.Close()
	w := zip.NewWriter(d)
	defer w.Close()
	for _, file := range files {
		err := compress(file, "", w)
		if err != nil {
			zgh.ZLog().Error("message","compress the files","error",err.Error())
			return err
		}
	}
	return nil
}

func compress(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		zgh.ZLog().Error("message","compress the files","error",err.Error())
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				zgh.ZLog().Error("message","compress the files","error",err.Error())
				return err
			}
			err = compress(f, prefix, zw)
			if err != nil {
				zgh.ZLog().Error("message","compress the files","error",err.Error())
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		header.Name = prefix + "/" + header.Name
		if err != nil {
			zgh.ZLog().Error("message","compress the files","error",err.Error())
			return err
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			zgh.ZLog().Error("message","compress the files","error",err.Error())
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			zgh.ZLog().Error("message","compress the files","error",err.Error())
			return err
		}
	}
	return nil
}

