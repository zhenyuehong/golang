package main

import (
	"fmt"
	"regexp"
)

//正则表达式

func main() {
	//正则表达式匹配email
	//text  := "my email is hzycyq@gmail.com"
	//res:= regexp.MustCompile(".+@.+\\..+")
	//或者
	//res:= regexp.MustCompile(`.+@.+\..+`)
	//这样会打印出所有的字符，我们再改一下
	//text  := "my email is hzycyq@gmail.com@abc.com"
	//res:= regexp.MustCompile(`[a-zA-Z0-9]+@.+\..+`)
	//这样，我们就可以从这一串字符中获取到了我们的邮箱,但是如果字符串改成这样my email is hzycyq@gmail.com@abc.com
	//获取到的邮箱又有问题了，所以我们的后面不能直接用 . ，我们再做修改
	//text  := "my email is hzycyq@gmail.com@abc.com"
	//res:= regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9\.]+\.[a-zA-Z0-9]+`)
	//这样，就打印出正确的邮箱了

	//还有一种情况，当text是多行的时候，这样只能打印出第一个邮箱,这样我们就不能用FindString，要用
	text := `my email is hzycyq@gmail.com.cn
			my email1 is hzycyq@abc.com
			my email2 is hzycyq@qq.com
			`
	res := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)

	//findString := res.FindString(text)
	//findString := res.FindAllString(text,-1)
	//fmt.Println(findString)
	//现在，我们有个需求，想把邮箱中@ 前后的数据提取出来，我们可以这样做
	findString := res.FindAllStringSubmatch(text, -1)
	for _, val := range findString {
		fmt.Println(val)
		//fmt.Println(val[0])
	}
}
