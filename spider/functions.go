package utils

import (
	"fmt"
	"strings"
	"github.com/tebeka/selenium"
)

func checkErr(err error, wd selenium.WebDriver) {
	if err != nil {
		if wd != nil {
			GetScreenshot(wd, PRINTSCREENFILENAMEERROR)

			currentURL, _ := wd.CurrentURL()

			fmt.Println("CurrentURL:\n")
			fmt.Println(currentURL)

			//TODO Save em .html file
			fmt.Println("PageSource\n")
			fmt.Println(wd.PageSource())
		}
		panic(err.Error())
	}
}

func GetCookies(wd selenium.WebDriver) {
	cookies, _ := wd.GetCookies()
	fmt.Printf("Cookies: -------------------%#v\n", cookies)

	for _, v := range cookies {
		fmt.Println("cookie is "+v.Name)
		fmt.Println("cookie val "+v.Value)
	}
}

func GetScreenshot(wd selenium.WebDriver, nameKey string) error {
	img, err := wd.Screenshot()
	if err != nil {
		return err
	}
	file, err := os.Create(fmt.Sprintf("data/%s_%s.png", nameKey, time.Now().Format("2006-01-02 15:04:05.999999")))
	defer file.Close()

	checkErr(err, wd)
	_, err = file.Write(img)
	return err
}

// TODO: For and print N seconds
func WaitPage() {
	fmt.Printf("Sleep\n")
	fmt.Println(time.Millisecond)
	time.Sleep(time.Millisecond * 10000)
}

func ResizeWindow(wd selenium.WebDriver) {
	var handles []string
	handles, err = wd.WindowHandles()

	err = wd.ResizeWindow(handles[0], 2000, 2000)
	checkErr(err, wd)
}
