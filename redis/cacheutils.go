package redis

import "bytes"

const PUBLISHEDLIST = "PUBLISHEDLIST"

// generate key to save on redis
func Generate(string ...string) string {
	var buffer bytes.Buffer
	pre := "dousheng"
	buffer.WriteString(pre)
	for _, k := range string {
		buffer.WriteString(":")
		buffer.WriteString(k)
	}
	return buffer.String()
}