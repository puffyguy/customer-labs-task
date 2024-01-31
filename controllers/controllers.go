package controllers

import (
	"fmt"
	"net/http"
	"practice/customer-labs-test/models"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	dataChannel = make(chan models.Input)
)

/*
SendData is a handler function for processing incoming JSON data from a client.
It uses the Gin framework's ShouldBindJSON to bind the request body to the 'models.Input' struct.
If the binding is unsuccessful, it returns a Bad Request response with the error details.
*/
func SendData(c *gin.Context) {
	var p models.Input

	// Bind the JSON data from the request body to the 'models.Input' struct
	bindErr := c.ShouldBindJSON(&p)
	if bindErr != nil {
		fmt.Printf("Data bind err %v\n", bindErr)
		c.JSON(http.StatusBadRequest, bindErr.Error())
		return
	}

	// Launch a goroutine to send the parsed data to the 'sendChannel'
	go func(d models.Input) {
		dataChannel <- d
	}(p)
	// defer close(sendChannel)

	data, ok := <-dataChannel
	if !ok {
		fmt.Println("Channel is closed, no data available")
		return
	}

	// Convert the received data and prepare a response
	res := dataConverter(data)

	// Send the response back to the client
	c.JSON(http.StatusOK, res)
}

/*
dataConverter is a utility function that takes an input of type 'models.Input' and converts it
into an output of type 'models.Output' with specific fields extracted from the input data.
*/
func dataConverter(d models.Input) models.Output {
	// fmt.Println("Data:", d)

	// Extract data from the 'Input' model
	inputData := d.Data

	// Create an 'Output' model with relevant fields populated
	output := models.Output{
		Event:           inputData["ev"].(string),
		EventType:       inputData["et"].(string),
		AppId:           inputData["id"].(string),
		UserId:          inputData["uid"].(string),
		MessageId:       inputData["mid"].(string),
		PageTitle:       inputData["t"].(string),
		PageURL:         inputData["p"].(string),
		BrowserLanguage: inputData["l"].(string),
		ScreenSize:      inputData["sc"].(string),
		Attributes:      createAttributes(inputData, "atr"),
		Traits:          createAttributes(inputData, "uatr"),
	}

	// Return the converted 'Output' model
	return output
}

/*
createAttributes is a utility function that takes a map of input data, a prefix, and extracts attributes
from the input data based on a specific naming convention. It returns a map of attributes with the specified prefix.
*/
func createAttributes(inputData map[string]interface{}, prefix string) map[string]interface{} {
	// Initialize an empty map to store extracted attributes
	attributes := make(map[string]interface{})

	// Define a regular expression pattern based on the prefix for attribute keys
	attrRegex := regexp.MustCompile(fmt.Sprintf(`^%sk(\d+)$`, prefix))

	// Iterate over the input data to extract attributes based on the naming convention
	for key, value := range inputData {
		// fmt.Printf("%v,%v\n", key, value)

		// Check if the current key matches the attribute naming pattern
		if matches := attrRegex.FindStringSubmatch(key); len(matches) == 2 {
			// fmt.Println("Matches:", matches)

			// Extract the index from the matched pattern
			index, _ := strconv.Atoi(matches[1])

			// Create attribute keys and type keys based on the extracted index
			attrKey := fmt.Sprintf("%sv%d", prefix, index)
			typeKey := fmt.Sprintf("%st%d", prefix, index)
			// fmt.Printf("%d,%s\n", index, attrKey)
			// fmt.Printf("%d,%s\n", index, typeKey)

			// Retrieve the attribute value and type from the input data
			attrValue, exists := inputData[attrKey]
			attrtype, existsType := inputData[typeKey]

			// Check if both the attribute value and type exist in the input data
			if exists && existsType {
				// Create a nested map with 'value' and 'type' keys and add it to the attributes map
				attributes[value.(string)] = map[string]interface{}{
					"value": attrValue,
					"type":  attrtype,
				}
			}
		}
	}

	// Return the map of extracted attributes
	return attributes
}
