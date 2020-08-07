package service

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"log"
)

func batch() {
	opts := []selenium.ServiceOption{}
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}

	// 禁止加载图片，加快渲染速度
	imagCaps := map[string]interface{}{
		"profile.managed_default_content_settings.images": 2,
	}

	chromeCaps := chrome.Capabilities{
		Prefs: imagCaps,
		Path:  "",
		Args: []string{
			"--headless", // 设置Chrome无头模式
			"--no-sandbox",
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
	// 这是目标网站留下的坑，不加这个在linux系统中会显示手机网页，每个网站的策略不一样，需要区别处理。
	driver.AddCookie(&selenium.Cookie{
		Name:  "defaultJumpDomain",
		Value: "www",
	})
	// 导航到目标网站
	err = driver.Get("http://free.win007.com/live.aspx?Edition=1&lang=0&ad=&adurl=&color=F0F0E0&sound=0")
	if err != nil {
		panic(fmt.Sprintf("Failed to load page: %s\n", err))
	}

	//err = driver.Wait(Displayed(selenium.ByID, "#table_live"))
	//if err != nil {
	//	log.Println("错误", err)
	//	return
	//}
	trs, err := driver.FindElements(selenium.ByCSSSelector, "#table_live>tbody>tr[align]")
	if err != nil {
		log.Println("错误", err)
		return
	}
	for i := 1; i < len(trs); i++ {
		tds, err := trs[i].FindElements(selenium.ByCSSSelector, "td")
		if err != nil {
			log.Println("错误", err)
			return
		}
		league, _ := tds[1].Text()
		time, _ := tds[2].Text()
		status, _ := tds[3].Text()
		m, _ := tds[4].Text()
		score, _ := tds[5].Text()
		g, _ := tds[6].Text()

		fmt.Println(league + "|" + time + "|" + status + "|" + m + "|" + score + "|" + g)

	}

}
