package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader)Read(p []byte)(int,error){
	n,err:=reader.r.Read(p)
	if err!=nil {
		return n,err
	}
	for i:=0;i<n;i++{
		if 'a' <= p[i] && p[i] <= 'z'{
			p[i]=(p[i]-'a'+13)%26+'a'
		}else if 'A' <= p[i] && p[i] <= 'Z'{
			p[i]=(p[i]-'A'+13)%26+'A'
		}
	}
	return n,nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
