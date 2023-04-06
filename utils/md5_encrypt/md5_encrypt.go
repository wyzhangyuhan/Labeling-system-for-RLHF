package md5_encrypt

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

func MD5(para string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(para))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}

func B64MD5(para string) string {
	return MD5(base64.StdEncoding.EncodeToString([]byte(para)))
}
