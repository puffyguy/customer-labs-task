# Repository Description

This repository contains a Go application built with the Gin framework for handling incoming JSON data and processing it asynchronously. The core functionalities include parsing JSON data, converting input into structured output, and managing a communication channel between goroutines.

## SendData Function

- Handles incoming JSON data using the Gin framework.
- Parses the data and sends it to a channel for asynchronous processing.
- Retrieves processed data and responds to the client.

## dataConverter Function

- Converts input data of type 'models.Input' into structured output of type 'models.Output'.
- Extracts specific fields from the input data to create a well-defined output structure.

## createAttributes Function

- Utility function to extract attributes from a map of input data based on a naming convention.
- Generates a map of attributes with nested values and types.

The code is organized to facilitate the handling of incoming data, conversion, and concurrent processing, providing a foundation for scalable and responsive applications.
