package Demo

import (
	//"fmt"
	"../HttpSeries/Http"
	//"../HttpSeries/Https"
)

func ShowHttpseriesFunc () {
	url := "http://www.baidu.com"
	Http.HttpGet(url)
}
