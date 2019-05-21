package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	integ "github.com/zokypesch/flyweight/integration"
)

var listVendor = map[string]interface{}{
	"kudo": integ.NewKudoService,
	"mba":  integ.NewMbaService,
}

func posting(c *gin.Context) {

	var responseGetAll map[string]interface{}
	errMessage := ""
	raw, _ := c.GetRawData()

	json.Unmarshal([]byte(raw), &responseGetAll)

	source := responseGetAll["source"].(string)
	productCode := responseGetAll["product_code"].(string)
	_, exist := listVendor[source]

	if !exist {
		c.JSON(400, gin.H{
			"status":  "your request terminated",
			"message": "uknow vendor",
		})
	}

	// log.Println(responseGetAll)

	st := listVendor[source].(func(code string, data interface{}) integ.Vendor)(productCode, nil)

	if err := st.Do(); err != nil {
		errMessage = err.Error()
	}

	c.JSON(200, gin.H{
		"message": "executed",
		"err":     errMessage,
	})
}

func main() {

	router := gin.Default()
	router.POST("/test", posting)
	router.Run()
	fmt.Println("Hello world")
}
