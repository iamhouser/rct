# HTTP Request Multithreaded Tester

This is a simple Go program that sends multiple HTTP GET requests to a specified URL concurrently using goroutines. It collects the HTTP response status codes from each request and provides a summary at the end, including the total execution time.

## Features

- Sends concurrent HTTP GET requests to a target URL.
- Allows specification of the number of concurrent threads.
- Collects and summarizes HTTP response status codes.
- Measures and displays total execution time.

## Requirements

- Go installed on your machine (Go 1.15 or higher recommended).

## Installation

1. **Clone the repository or copy the code into a local file** (e.g., `main.go`).

   ```bash
   git clone https://github.com/yourusername/yourrepository.git
   ```

2. **Navigate to the project directory**.

   ```bash
   cd yourrepository
   ```

3. **Build the executable**.

   ```bash
   go build main.go
   ```

## Usage

Run the program from the command line with the following flags:

```bash
./main -u <URL> -t <threads> -h <token>
```

- `-u` (required): The target URL to send requests to.
- `-t` (optional): The number of concurrent threads (default is 10).
- `-h` (optional): Authorization token (currently not used in the code).

### Example

```bash
./main -u https://example.com -t 20
```

## Output

The program will output the following information:

- Response codes from each request.
- Summary of the number of requests per status code.
- Total execution time.

### Sample Output

```
Processing...
Response code of request: 200
Response code of request: 200
...

------------------------------------
Target: https://example.com
Amount of requests: 20
Requests with code 200: 20
Program executed for: 1.234567s
```

## Notes

- The `-h` flag for the authorization token is currently accepted but not utilized in the HTTP requests.
- The program uses a `WaitGroup` to synchronize goroutines.
- **Important**: The shared slice `amountCodes` is accessed by multiple goroutines without synchronization, which may lead to race conditions. For thread-safe operations, consider using synchronization mechanisms like mutexes or channels.

## License

This project is open-source and available under the [MIT License](LICENSE).
