package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
	"time"
)

func AuthTypeC(authKey, originalUrl string, t int64) (string, error) {
	oUrl, err := url.Parse(originalUrl)
	if err != nil {
		return "", err
	}
	ts := time.Now().Unix()
	if t != 0 {
		ts = t
	}
	tsStr := fmt.Sprintf("%X", ts)
	hash := md5Hash(authKey + oUrl.Path + tsStr)
	newUrl := fmt.Sprintf("%s://%s/%s/%s%s", oUrl.Scheme, oUrl.Host, hash, tsStr, oUrl.Path)
	return newUrl, nil
}

// md5 hash
func md5Hash(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	hashSum := h.Sum(nil)
	hexString := hex.EncodeToString(hashSum)
	return hexString
}
