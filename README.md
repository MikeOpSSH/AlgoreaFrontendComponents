# whats-this

You can use this to retrieve the current price for an crypto currency

I used this in another project and thought it will make sense 
to publish this as an own package 

## How to use it
```go
import "https://github.com/mnlwldr/coinbase"
```

## Usage
```go
response, err := coinbase.Get("SHIB-EUR")
if err != nil {
	log.Println(err)
}
fmt.Printf("%.7f\n", response.Amount)
```

## Output
```sh
0.0000253
```

[![Go Reference](https://pkg.go.dev/badge/github.com/mnlwldr/coinbase.svg)](https://pkg.go.dev/githu