# wsdlgo
Generates Go code from a WSDL file

### Install

* Download and build locally: `go get github.com/Bridgevine/wsdlgo`

### Usage
```
Usage: wsdlgo [options] service.wsdl
  -o string
        File where the generated code will be saved (default "types.go")
  -p string
        Package under which code will be generated (default "types")
It's possible to combine multiple WSDLs to produce a single output file.

./wsdlgo -cer ../pi-centurylink/certs/cert.cer -ck ../pi-centurylink/certs/cert.key ContentService.wsdl ContentOrderService.wsdl ServiceOrderStatusService.wsdl
Certificate and certificate key are optional.

```

### Influences
```
Project began as a fork of https://github.com/hooklift/gowsdl but deviated too much from original, and decided to start a brand new service borrowing a lot of its ideas.
```