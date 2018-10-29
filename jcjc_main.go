package main

import (
	"fmt"
	"jcjc_golang_client/jcjc"
)

func main(){
	fmt.Println("欢迎使用 JCJC错别字检测引擎  CuoBieZi.net ")

	fmt.Println("中文博大精深， 请在真实的场景中测试，尽量不要输入几个词与！ 没有上下文，很难办。")


	content := `

腾讯今年中国人民共和国下半年上世纪将在微信账户钱包帐户的“九宫格”中开设快速的笑着保险入口，并上线保险产品。台万第二大金融控股公司富邦金控已与腾讯谈成合作，上述保险产品将由富邦金控旗下内地子公司富邦财险开发或引进。
	`

	url_str := "http://localhost:8235/spellcheck/json_check/json_phrase"
	jcjc.P__process_one_file(content, url_str)


	fmt.Println("Done , 欢迎使用 JCJC错别字检测引擎  CuoBieZi.net  ")
}
