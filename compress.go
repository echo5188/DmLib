/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"bytes"
	"compress/zlib"
	"github.com/golang/snappy"
)

func Compress(Msg *Msg, offset int, length int, compressID int) ([]byte, error) {
	if compressID == CompressID01 {
		return snappy.Encode(nil, Msg.Dm_build_655(offset, length)), nil
	}
	return GzlibCompress(Msg, offset, length)
}

func UnCompress(srcBytes []byte, compressID int) ([]byte, error) {
	if compressID == CompressID01 {
		return snappy.Decode(nil, srcBytes)
	}
	return GzlibUncompress(srcBytes)
}

func GzlibCompress(Msg *Msg, offset int, length int) ([]byte, error) {
	var ret bytes.Buffer
	var w = zlib.NewWriter(&ret)
	w.Write(Msg.Dm_build_655(offset, length))
	w.Close()
	return ret.Bytes(), nil
}

func GzlibUncompress(srcBytes []byte) ([]byte, error) {
	var bytesBuf = new(bytes.Buffer)
	r, err := zlib.NewReader(bytes.NewReader(srcBytes))
	if err != nil {
		return nil, err
	}
	defer r.Close()
	_, err = bytesBuf.ReadFrom(r)
	if err != nil {
		return nil, err
	}
	return bytesBuf.Bytes(), nil
}
