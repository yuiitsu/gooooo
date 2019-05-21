package source

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

//
var registerInstance *register

type App struct {
	W http.ResponseWriter
	Params url.Values
}

type controller struct {
	requestType string
	Path string
	ControllerType reflect.Type
	ControllerValue reflect.Value
}

type register struct {
	Controllers []*controller
}

type jsonResult struct {
	Code int
	Msg string
	Data interface{}
}

func GetRegisterInstance() *register {
	if registerInstance == nil {
		registerInstance = &register{}
	}
	return registerInstance
}

func (this App) GetParams(key string) string {
	if _, ok := this.Params[key]; ok {
		return this.Params[key][0]
	} else {
		return ""
	}
}

func Router(requestType string, path string, c interface{}) {
	r := GetRegisterInstance()
	rValue := reflect.ValueOf(c)
	rType := reflect.TypeOf(c)
	r.Controllers = append(r.Controllers, &controller{
		requestType, path, rType, rValue})
}

func result() {

}

func (this App) Write(data interface{}) {
	this.W.Header().Set("Content-type", "application/json")
	//
	result := jsonResult{0, "ok", data}
	resultJson, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	this.W.Write(resultJson)
}

func Run() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		//
		r := GetRegisterInstance()
		//
		requestType := request.Method
		path := request.URL.Path
		paths := strings.Split(path, "/")
		method := paths[len(paths) - 1]
		//
		u, err := url.ParseQuery(request.URL.RawQuery)
		if err == nil {
		}
		params := u
		//
		controller := &controller{}
		for _, value := range r.Controllers {
			if path == value.Path {
				controller.Path = value.Path
				controller.ControllerValue = value.ControllerValue
				controller.ControllerType = value.ControllerType
				controller.requestType = value.requestType
				break
			}
			item := value.Path
			fmt.Println(item)

		}

		if controller.Path != "" {
			if requestType != controller.requestType {
				writer.WriteHeader(405)
			} else {
				//
				reciver := controller.ControllerValue.Elem().FieldByName("App")
				reciver.Set(reflect.ValueOf(App{writer, params}))
				m, exist := controller.ControllerType.MethodByName(method)
				if exist {
					args := []reflect.Value{controller.ControllerValue}
					m.Func.Call(args)
				} else {
					writer.WriteHeader(403)
				}
			}
		} else {
			writer.WriteHeader(404)
		}
	})
	http.ListenAndServe(":8080", nil)
}
