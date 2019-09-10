package main

import (
	"bytes"
	"fmt"
	"github.com/ipfs/go-ipfs-api"
	"io/ioutil"
)

var sh *shell.Shell

func UploadIPFS(str string) string {
	//(1)创建shell方式:1>shell.NewShellWithClient(); 2>shell.NewShell
	url := "http://192.168.137.203:5001"
	sh = shell.NewShell(url)
	//若指定https信息，则使用NewShellWithClient函数里面的gohttp.Client进行指定
	//(1.1)添加节点
	//ipfs bootstrap add /ip4/192.168.1.114/tcp/4001/ipfs/QmRM8d9c2Nuwg8sqyJZPhR52TzpyYd7CpGc1FhNJQcdTkh  #添加一个ipfs连接节点
	//sh.BootstrapAdd()
	//(2)
	hash, err := sh.Add(bytes.NewBufferString(str))
	if err != nil {
		fmt.Println("上传ipfs时错误：", err)
	}
	return hash
}

func CatIPFS(hash string) string {
	sh = shell.NewShell("http://192.168.137.203:5001")

	read, err := sh.Cat(hash)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(read)

	return string(body)
}

func main() {
	//strVal := UploadIPFS("1111111111111")
	//fmt.Println("上传的字符串为:" + strVal)

	strRtn := CatIPFS("QmSXAWxGK6YwYhGVn631axQb26FzbZbndLQqbvCSJTF9u6")
	fmt.Println("下载的字符串为：" + strRtn)
}

/**
(1)备注：若在http://192.168.137.200:5001集群网络中，执行的响应结果信息为：
	上传的字符串为:QmSXAWxGK6YwYhGVn631axQb26FzbZbndLQqbvCSJTF9u6
	下载的字符串为：1111111111111
   其中通过命令进行查看:ipfs cat /ipfs/QmSXAWxGK6YwYhGVn631axQb26FzbZbndLQqbvCSJTF9u6
(2)备注:
	启动网络：docker-compose up -d
	将参数修改为：http://192.168.137.203:5001
	显示结果为：
	上传的字符串为:QmSXAWxGK6YwYhGVn631axQb26FzbZbndLQqbvCSJTF9u6
	下载的字符串为：1111111111111
*/
