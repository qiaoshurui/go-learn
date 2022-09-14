// description:
// @author renshiwei
// Date: 2022/9/14 19:01

package generics

type CommonResponse struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type TestResponse struct {
	Code string   `json:"code"`
	Msg  string   `json:"msg"`
	Data []string `json:"data"`
}
