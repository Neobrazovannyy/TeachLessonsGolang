package main

import(
	"fmt"
	"net/http"
	"io"
	"strings"
)

/*
Функция http.Get(url string) возвращает указатель на 
структуру http.Response (ответ сервера) и 
error (ошибку, если что-то пошло не так).
В URL-параметре указывают HTTP-протокол или HTTPS-протокол.
*/
/*
что бы получить часть заголовко достаточно задать параметр 
contentType := response.Header.Get("Content-Type")
Content-Type: [text/plain; charset=utf-8]
*/

/*
Тело ответа представлено в поле Response.Body
интерфейсом потокового чтения io.ReadCloser.
io.ReadCloser содержит два метода:
	1) Read(p []byte) (n int, err error) (интерфейс Reader);
	2) Close() error (интерфейс Closer).
! Если клиент успешно получил ответ, то нужно закрыть Body независимо от того, прочитаете вы его или нет. 
*/

/*
Метод POST
Post(url, contentType string, body io.Reader) (resp *Response, err error)
*/

func RequestGet() *http.Response {
	var res_req_get *http.Response
	var err_req_get error
	res_req_get, err_req_get = http.Get("http://127.0.0.1:8080/main_page")

	if(err_req_get!=nil){
		panic(err_req_get)
	}

	return res_req_get
}

func ResponseGetContentType(response *http.Response) {
	var content_type string = response.Header.Get("Content-Type");
	fmt.Println("================= Content-Type =================")
	fmt.Println(content_type+"\r\n")
}

func ResponsePost() *http.Response{
	var data_json_post string = `{"name":"golang", "project":"server"}`

	var res_post *http.Response
	var err_res_post error
	res_post, err_res_post = http.Post("http://127.0.0.1:8080/main_page", "application/json", strings.NewReader(data_json_post))
	defer res_post.Body.Close()

	if(err_res_post!=nil){
		panic(err_res_post)
	}

	return res_post
}

func main(){
	//---> Request Get
	fmt.Println("\n<-------------------- Response: Request Get\n")
	var response_req_get *http.Response = RequestGet()

	// Get: Header
	var format_res_get string = "================= Header =================\r\n"
	for k, v := range response_req_get.Header{
		format_res_get += fmt.Sprintf("%s: %v\r\n", k, v)
	}
	fmt.Println(format_res_get)

	// Get: Header.Content-Type
	// ResponseGetContentType(response_req_get);

	// Get: Body
	var body_response []byte
	var err_get_body error
	body_response, err_get_body = io.ReadAll(response_req_get.Body)
	response_req_get.Body.Close()

	if(err_get_body!=nil){
		panic(err_get_body)
	} else{
		fmt.Println("================= Body =================")
		fmt.Println(string(body_response))
	}

	//---> Request Post
	fmt.Println("\n<-------------------- Response: Request Post\n")
	var response_req_post *http.Response = ResponsePost()
	//Post: Header
	var format_res_post string = "================= Header =================\r\n"
	for k, v := range response_req_post.Header{
		format_res_post += fmt.Sprintf("%s: %v\r\n", k, v)
	}
	fmt.Println(format_res_post)

}