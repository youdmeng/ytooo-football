package service

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"log"
	"strings"
	"time"
)

// 获取所有联赛 一天执行一次就可以
func GetLeagueInfo() {
	var opts []selenium.ServiceOption
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	// 禁止加载图片，加快渲染速度
	//imagCaps := map[string]interface{}{
	//	"profile.managed_default_content_settings.images": 2,
	//}
	chromeCaps := chrome.Capabilities{
		//Prefs: imagCaps,
		Path: "",
		Args: []string{
			"--headless", // 设置Chrome无头模式
			"--no-sandbox",
			"--disable-gpu",
			"--disable-dev-shm-usage",
			//"--proxy-server=socks5://127.0.0.1:10808",
			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7", // 模拟user-agent，防反爬
		},
	}
	caps.AddChrome(chromeCaps)
	// 启动chromedriver，端口号可自定义
	_, err := selenium.NewChromeDriverService("Q:\\chromedriver.exe", 9515, opts...)
	if err != nil {
		log.Printf("Error starting the ChromeDriver server: %v", err)
	}
	// 调起chrome浏览器
	driver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9515))
	if err != nil {
		panic(err)
	}
	_ = driver.ResizeWindow("", 1920, 1080)
	// 导航到目标网站
	err = driver.Get("http://free.win007.com/live.aspx?Edition=1&lang=0&ad=&adurl=&color=F0F0E0&sound=0")
	if err != nil {
		panic(fmt.Sprintf("Failed to load page: %s\n", err))
	}
	time.Sleep(1000)
	ele, _ := driver.FindElement(selenium.ByCSSSelector, `img[alt="赛事选择"]`)
	_ = ele.Click()
	labels, _ := driver.FindElements(selenium.ByCSSSelector, "#myleague tr label")

	//获取联赛信息
	var ls []string
	for i := 0; i < len(labels); i++ {
		league, _ := labels[i].Text()
		if strings.Contains(league, "中超") {
			league = "中超"
		}
		ls = append(ls, league)
	}

}
