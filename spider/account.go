package spider

import (
	"fmt"
	"os"
	"time"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

const PRINTSCREENFILENAME = "PrintScreenSpider"
const PRINTSCREENFILENAMEERROR = PRINTSCREENFILENAME + "Error"

func init() {
	ftm.Printf("Yaml: %+v\n", config.load())
}

func StartService() {
	// Start a Selenium WebDriver server instance (if one is not already running).
	const (
		// These paths will be different on your system.
		seleniumPath     = y.Path[0].Selenium
		geckoDriverPath  = y.Path[0].Gecko
		chromeDriverPath = y.Path[0].Chrome
		port             = 8080
	)
	opts := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
		selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
		selenium.Output(os.Stderr),            // Output debug information to STDERR.
		selenium.ChromeDriver(chromeDriverPath),
	}
	selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	checkErr(err, nil)

	defer service.Stop()
}

func MakeLogin(wd selenium.WebDriver) {
	err = wd.Get(y.Url[0].Login)
	checkErr(err, wd)

	GetScreenshot(wd, PRINTSCREENFILENAME)
	
	elem, err := wd.FindElement(selenium.ByID, "username")
	checkErr(err, wd)

	time.Sleep(time.Millisecond * 100)

	err = elem.Clear()
	checkErr(err, wd)
	
	err = elem.SendKeys(y.User[0].Login)
	checkErr(err, wd)
	
	elem, err = wd.FindElement(selenium.ByID, "senha")
	checkErr(err, wd)

	err = elem.Clear()
	checkErr(err, wd)

	err = elem.SendKeys(y.User[0].Password)
	checkErr(err, wd)

	btn, err := wd.FindElement(selenium.ByID, "submit")
	checkErr(err, wd)
	
	err = btn.Click()
	checkErr(err, wd)
}

func ListAccounts(wd selenium.WebDriver) {
	err = wd.Get(y.Url[0].List_Accounts)
	checkErr(err, wd)

	GetCookies(wd)

	GetScreenshot(wd, PRINTSCREENFILENAME)

	var webelem selenium.WebElement

	WaitPage()

	webelem, err = wd.FindElement(selenium.ByID, "select-competencia")
	checkErr(err, wd)

	elm, err = utils.Select(webelem)
	checkErr(err, wd)

	err = elm.SelectByValue("901")
	checkErr(err, wd)
	//driver.execute_script('document.getElementById("select2-chosen-2").innerHTML = "07/2019-1";', element)

	WaitPage()

	GetScreenshot(wd, PRINTSCREENFILENAME)

	btn, err = wd.FindElement(selenium.ByID, "btn-pesquisar")
	checkErr(err, wd)

	err = btn.Click()
	checkErr(err, wd)

	WaitPage()

	// TODO: Check div exists
	// <div class="dataTables_info">Nenhum registro encontrado
	// </div>

	GetScreenshot(wd, PRINTSCREENFILENAME)
}

func ExecuteSpider() {

	StartService()

	//TODO: refact to function
	// Connect to the WebDriver instance running locally.
	//caps := selenium.Capabilities{"browserName": "firefox"}
	caps := selenium.Capabilities{"browserName": "chrome"}
	imagCaps := map[string]interface{}{}
	chromeCaps := chrome.Capabilities{
		Prefs: imagCaps,
		Path:  "",
		Args: []string{
			"--headless", // linux
			"--no-sandbox",
			"--disable-extensions", // disabling extensions
			"--disable-gpu", // applicable to windows os only
			"--disable-dev-shm-usage", // overcome limited resource problems
			"--start-maximized",
			//"--user-agent=Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1", // 模拟user-agent，防反爬
			"--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36",
		},
	}
	caps.AddChrome(chromeCaps)

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	checkErr(err, wd)
	defer wd.Quit()

	ResizeWindow(wd)

	MakeLogin(wd)

	ListAccounts(wd)

	//TODO: Save to HTML File: 
	fmt.Println(wd.PageSource())

}

