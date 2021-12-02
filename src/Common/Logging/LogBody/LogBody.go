package LogBody

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"os"
	"reflect"
	"time"
)

type Body struct {
	Data map[string]string
}
type DynamicBody Body

var StaticLog *Body

func CreateStaticBody() *Body {
	StaticLog = &Body{Data: map[string]string{}}
	return StaticLog
}

func (entry *Body) AddApplicationInfo(applicationName string) *Body {
	StaticLog.Data["ApplicationName"] = applicationName
	return StaticLog
}

func (entry *Body) AddHostName() *Body {
	hostName, _ := os.Hostname()
	StaticLog.Data["HostName"] = hostName
	return StaticLog
}

func GetDynamicLog() *DynamicBody {
	return &DynamicBody{Data: StaticLog.Data}
}

func (entry *DynamicBody) AddRequestInfo(c echo.Context) *DynamicBody {
	entry.Data["QueryString"] = c.QueryString()
	entry.Data["Scheme"] = c.Scheme()
	entry.Data["Path"] = c.Path()
	entry.Data["Method"] = c.Request().Method
	return entry
}
func (entry *DynamicBody) AddTimeInfo() *DynamicBody {
	entry.Data["Time"] = time.Now().String()
	return entry
}
func (entry *DynamicBody) AddStruct(s interface{}) *DynamicBody {
	v := reflect.ValueOf(s)
	for i := 0; i < v.NumField(); i++ {
		varName := v.Type().Field(i).Name
		varType := v.Type().Field(i).Type
		varValue := v.Field(i).Interface()
		if varType.Kind() == reflect.Struct {
			entry.AddStruct(varValue)
		} else {
			entry.Data[varName] = fmt.Sprintf("%v", varValue)
		}
	}
	return entry
}
