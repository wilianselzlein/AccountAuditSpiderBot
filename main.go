package main

import (
	"fmt"
	"AccountAuditSpiderBot/spider"
)

// This example shows how to navigate to a Account Analisys page, input a
// short program, run it, and inspect its output.
//
// If you want to actually run this spider, ensure the file paths at the config file are correct.
//
// Run:
//      cd ~/AccountAuditSpiderBot 
//      go run main.go

func main() {

	spider.ExecuteSpider()
	
	fmt.Printf("Finish\n")

}

