# Cross-Platform Command Execution Application

### Description

This Go application allows for basic command execution on macOS and Windows. It supports ping and system information retrieval, with an HTTP interface for sending commands.

Build Instructions
To compile the application, make sure you have Go installed.

#### Compile for macOS
```
GOOS=darwin GOARCH=amd64 go build -o crossplatform-app
```
#### Compile for Windows
```
GOOS=windows GOARCH=amd64 go build -o crossplatform-app.exe
```

### API Documentation

The application provides an HTTP endpoint for executing commands and returning results in JSON format.

#### Endpoint: `/execute`
#### Method: POST
#### Request Example
````
{
  "type": "ping", // or "sysinfo"
  "payload": "localhost" // only for "ping" type
}
````

- `type`: Can be `"ping"` for network connectivity check or `"sysinfo"` for system information.
- `payload`: Hostname or IP address for the `ping` command.

#### Response Example
```
{
  "success": true,
  "data": {
    // Command-specific data
  },
  "error": "string"
}
```

- `success`: Indicates if the command was successful.
- `data`: Contains the command result (ping response time, system hostname, IP, etc.).
- `error`: Contains an error message, if any.

### Installation Guide
1. macOS Installation:

    - Run the installer package .pkg you created.
    - The application will be installed in /usr/local/bin/ and will start automatically on boot using a Launch Daemon.

2. Windows Installation (Pending):

    - Currently, the Windows installer is not provided. The app can be run directly by executing the compiled .exe file.

# Testing

To run the unit tests:
```
go test -v
```

Tests include:

- Ping: Checks connectivity to a specified host.
- GetSystemInfo: Verifies that system hostname and IP address are correctly retrieved.

# Demonstration
The short video or GIF demonstrating the application in action was not included because GitHub does not support direct commits of video files.

To run the app:

1. Start the application.
2. Send a POST request to http://localhost:8080/execute with either a "ping" or "sysinfo" command.
3. Observe the JSON response with command results.

