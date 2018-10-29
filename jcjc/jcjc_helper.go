package jcjc

import (
	"html/template"
	"github.com/jmcvetta/napping"
	"net/http"
	"encoding/json"
	"log"
	"fmt"
	"time"
)

/**


 */
type InputParas struct {
	Content string  		`json:"content" binding:"required"  `  //
	UserName string  		`json:"username" binding:"required" `  //
	//
	CheckMode string  		  `json:"check_mode" `
	Method  string        `json:"method" `
	Content2 string    `json:"content2" `

	Content3 string    `json:"content3" `
	Content4 string    `json:"content4" `
	Url string  `json:"url"  `
	IsCheckSensorWords bool `json:"check_sensitive_word"  `   // check sensor words

	LetsKillThemAll bool ` json:"lets_kill_them_all"  `

	ReturnFormat string `json:"return_format" `

	IsBase64Content bool `json:"is_base64_content"  `

	IsReturnSentence bool `json:"is_return_sentence"  `

	IsRefineSentence bool `json:"is_refine_sentence"  `

	IsFluentCheck bool  `json:"is_fluent_check"  `

	IsShowErratum  bool `json:"is_show_erratum"  `

	// 华为云 也用到了这个功能
	IsCheckPunctuation bool   `json:"is_check_punctuation"  `

	DocId string  `json:"doc_id" `


	IsNotAllowErrOverlap bool `json:"is_not_allow_overlap" ` // 四川客户


	LongTxtAction  string  `json:"long_txt_action" `
	LongTxtSession  string `json:"long_txt_session" `
	/**
		华为项目　  校对模式，　检查非常严格
		proofreading_mode
	 */
	ProofreadingMode bool  `json:"is_proofreading_mode" `

}


type SCase struct {
	Error string
	Tips string
	Sentence string
	// 2017-08-15
	ErrInfo  string
	Pos  	 uint
	//
	MarkType int
	ErrLevel uint
	//
	WordsLen  int
	//
	ReviewWords bool

	InnerId string

	DocId string

}

type RPhraseList struct {
	MarkWords []SCase
	Cases []SCase
	MustShowMessage string
	Successed bool
	Message string
	ClientAction string
	EnterpriseExtension EnterpriseExtension `json:"extension" `
}

type EnterpriseExtension struct {
	DomainWords  []string `json:"domain_words" `
	//
	CommStr  string  `json:"result_string" `
	CommStr2  template.HTML  `json:"result_string2" `
	//

}



func Post_json_data_to_jcjc(jsonStr []byte, jcjc_url_str  string ) (string, bool) {
	s := napping.Session{}
	h := &http.Header{}
	s.Header = h
	var data map[string]json.RawMessage
	err := json.Unmarshal(jsonStr, &data)
	if err != nil {
		log.Println("error when jcjc unmarshal:", err.Error())
		fmt.Println("error when jcjc unmarshal:", err.Error())

	}
	resp, err := s.Post(jcjc_url_str, &data, nil, nil)
	if err != nil {
		log.Println("error when post data to jcjc server:", err.Error())
		fmt.Println("error when post data to jcjc server:", err.Error())

		return "{\"succeed\":false}", false
	}
	if 200 == resp.Status() {
		return resp.RawText(), true
	}

	fmt.Println("error timeout:", resp.RawText(), resp.Status())
	return "{\"succeed\":false}", false
}





func P__process_one_file(content,url_str string ) {
	{
		var one_params InputParas
		one_params.IsReturnSentence = true
		one_params.Content = content
		one_params.UserName = "请向管理员申请"  // <-------------------
		var jcjc_result RPhraseList
		params_json_bytes, err := json.Marshal(one_params)
		if nil != err {
			fmt.Println("marshal error corrigenda :", err.Error())
			log.Println("marshal error corrigenda :", err.Error())
		}
		start := time.Now()
		jcjc_result_str, ok := Post_json_data_to_jcjc(params_json_bytes, url_str)
		elapsed := time.Since(start)
		if elapsed > 3 *time.Second{
			log.Println("jcjc", " took ", elapsed)
			fmt.Println("jcjc", " took ", elapsed)
		}
		//
		if ok {
			//log.Println("jcjc_result_str:" + jcjc_result_str)
			//fmt.Println("jcjc_result_str:" + jcjc_result_str)
			if err := json.Unmarshal([]byte(jcjc_result_str), &jcjc_result); err != nil {
				jcjc_result.Successed = false
				jcjc_result.Message = err.Error()
				//return oneObj
				log.Println("error jcjc when unmarshal :", jcjc_result_str)
			} else {
				jcjc_result.Successed = true
			}
		} else {

			jcjc_result.Successed = false
			log.Println("error: return json is not valid  可能是 proxy 的原因！")
			fmt.Println("error: return json is not valid")

		}
		fmt.Println("--- --- --- --- --- --- 测试结果开始 --- --- --- --- --- --- --- --- ")
		if jcjc_result.Successed && nil != jcjc_result.Cases {
			for index, one_case := range jcjc_result.Cases {
				fmt.Println(index,"\t|\t", one_case.Error,"\t|\t", one_case.Tips,"\t|\t", one_case.Sentence)
			}
		}
		fmt.Println("--- --- --- --- --- --- 测试结果结束 --- --- --- --- --- --- --- --- ")
	}
}








