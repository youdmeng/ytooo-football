package service

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"log"
	"reflect"
)

//获取今日比赛信息 数据 今日 12:00 至次日 12:00
func GetGameInfo() {

	opts := []selenium.ServiceOption{}
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}

	// 禁止加载图片，加快渲染速度
	imagCaps := map[string]interface{}{
		"profile.managed_default_content_settings.images":                 2,
		"managed_default_content_settings.permissions.default.stylesheet": 2,
		"managed_default_content_settings.javascript":                     2,
	}

	chromeCaps := chrome.Capabilities{
		Prefs: imagCaps,
		Path:  "",
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
	//主窗口对象
	mainHander, _ := driver.CurrentWindowHandle()
	if err != nil {
		panic(err)
	}
	// 导航到目标网站
	err = driver.Get("http://free.win007.com/live.aspx?Edition=1&lang=0&ad=&adurl=&color=F0F0E0&sound=0")
	if err != nil {
		panic(fmt.Sprintf("Failed to load page: %s\n", err))
	}

	//
	////err = driver.Wait(Displayed(selenium.ByID, "#table_live"))
	trs, err := driver.FindElements(selenium.ByCSSSelector, "#table_live>tbody>tr[align]")
	if err != nil {
		log.Println("错误", err)
		return
	}
	for i := 1; i < len(trs); i++ {
		_ = driver.SwitchWindow(mainHander)
		tds, _ := trs[i].FindElements(selenium.ByCSSSelector, "td")

		league := getText(tds[0])
		time := getText(tds[1])
		status := getText(tds[2])
		m := getText(tds[3])
		score := getText(tds[4])
		g := getText(tds[5])
		halfS := getText(tds[6])

		//取得基本信息
		fmt.Println(league + "|" + time + "|" + status + "|" + m + "|" + score + "|" + g + "|" + halfS)

		//获取球队信息
		left, _ := tds[3].FindElement(selenium.ByCSSSelector, "a[title]")
		left.Click()
		handles, _ := driver.WindowHandles()
		for n := 0; n < len(handles); n++ {
			if !reflect.DeepEqual(handles[n], mainHander) {
				_ = driver.SwitchWindow(handles[n])
			}
		}
		//胜率
		leftPrEles, _ := driver.FindElements(selenium.ByCSSSelector, `table[width="560"] font`)
		var lwinPr, lbetWinPr, lBigPr, rwinPr, rbetWinPr, rBigPr string
		if nil != leftPrEles && len(leftPrEles) != 0 {
			if len(leftPrEles) > 2 {
				lwinPr = getText(leftPrEles[1])
			}
			if len(leftPrEles) > 3 {
				lbetWinPr = getText(leftPrEles[2])
			}
			if len(leftPrEles) > 4 {
				lBigPr = getText(leftPrEles[3])
			}
			if len(leftPrEles) > 7 {
				rwinPr = getText(leftPrEles[6])
			}
			if len(leftPrEles) > 8 {
				rbetWinPr = getText(leftPrEles[7])
			}
			if len(leftPrEles) > 9 {
				rBigPr = getText(leftPrEles[8])
			}
		}
		fmt.Println(lwinPr + "|" + lbetWinPr + "|" + lBigPr + "|" + rwinPr + "|" + rbetWinPr + "|" + rBigPr)
		_ = driver.Close()
		_ = driver.SwitchWindow(mainHander)
		//获取比赛时间
		_ = tds[4].Click()
		handles, _ = driver.WindowHandles()
		for n := 0; n < len(handles); n++ {
			if !reflect.DeepEqual(handles[n], mainHander) {
				_ = driver.SwitchWindow(handles[n])
			}
		}
		_ = driver.Wait(Displayed(selenium.ByCSSSelector, `#content #matchData #matchItems span`))
		dateTime, _ := driver.FindElement(selenium.ByCSSSelector, `#content #matchData #matchItems span`)
		dateTimeStr := getText(dateTime)
		if len(dateTimeStr) > 15 {
			dateTimeStr = dateTimeStr[15:]
		}

		fmt.Println(dateTimeStr)
		_ = driver.Close()

		fmt.Println("==========================================")

	}
	_ = driver.Quit()
}

func Displayed(by, elementName string) func(selenium.WebDriver) (bool, error) {
	return func(wd selenium.WebDriver) (ok bool, err error) {
		var el selenium.WebElement
		el, err = wd.FindElement(by, elementName)
		if err != nil {
			return
		}
		ok, err = el.IsDisplayed()
		return
	}
}

func getText(wb selenium.WebElement) string {
	if nil == wb {
		return ""
	}
	str, err := wb.Text()
	if nil != err {
		return ""
	}
	return str
}
