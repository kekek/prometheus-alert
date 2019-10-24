package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"

	"github.com/kekek/target-go/internal/metrics"
)

var (
	programName = "target-go"
	gitVersion  = "git"
	dateVersion = "date"

	version = fmt.Sprintf("%s-%s-%s", programName, gitVersion, dateVersion)

	flagHelp bool
)

func init() {

	flag.BoolVar(&flagHelp, "v", false, "show version")
}

func main() {
	flag.Parse()

	if flagHelp {
		showVersion()
		os.Exit(0)
	}

	e := echo.New()

	router(e)

	e.Logger.Fatal(e.Start(":2112"))
}

func router(e *echo.Echo) {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/metrics", metrics.MetricHandler())

	e.GET("/version", func(c echo.Context) error {
		return c.JSON(http.StatusOK, version)
	})

}

func showVersion() {
	fmt.Println("version:", version)
}
