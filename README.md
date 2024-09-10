# Random JSON Generator

This Go application generates a file containing random, garbage JSON data of a specified size. It is useful for testing purposes where arbitrary JSON data of a specific size is needed.

## Features

- Generates random JSON objects up to the size specified by the user.
- Each JSON key is up to 16 characters long.
- Each JSON value is up to 128 characters long.
- The output is a valid JSON file.
- If the `.json` extension is not provided in the output file name, it will automatically append it.

## Usage

To generate a random JSON file, you need to specify two parameters:

- `--kb`: The size in kilobytes of the file to generate.
- `--file`: The output file path where the JSON will be saved.

### Run from Source

```bash {"id":"01J7ERZQ2E4GKG13FAMVXMA7A0"}
go run main.go --kb 10 --file fake.json
```

This will create a file named `output.json` containing approximately 10 KB of random JSON data.

### Docker

```bash {"id":"01J7ESVN977RDQT62KFF2Z7A83"}
docker build -t jsonr .
docker run -v ./:/out jsonr --kb 10 --file /out/fake.json
```

### Testing

You can run the test suite to verify that the generated JSON files match the expected sizes.

#### Running Tests

```bash {"id":"01J7ERZQ2E4GKG13FAMZ9RPX25"}
go test
```

The test creates a few files with random json data and checks that each file size is within 128 bytes of the requested size and that the file was generated correctly.