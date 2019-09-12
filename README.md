# go-substrate-rpc-client

The substrate RPC client for go. Some slides https://docs.google.com/presentation/d/1lvCP7wZpsl2ES6fAkHg8-LfLhgtMHcvHeITp77Bgn2w/edit#slide=id.g5d531cd2a8_0_70

## Contributing

### Installation

1. Install `golangci-lint`: https://github.com/golangci/golangci-lint#install

### Linting

`golangci-lint run --disable-all --enable=deadcode --enable=errcheck --enable=gosimple --enable=ineffassign --enable=structcheck --enable=typecheck --enable=unused --enable=varcheck --enable=dupl --enable=gochecknoinits --enable=goconst --enable=gocyclo --enable=gofmt --enable=gosec --enable=interfacer --enable=lll --enable=misspell --enable=scopelint --enable=stylecheck --enable=unconvert --enable=unparam --enable=golint --enable=goimports --enable=govet --enable=nakedret --enable=staticcheck --deadline=1m`

## Testing

## Run with centrifuge-chain

1. Install subkey command from [https://github.com/centrifuge/substrate](https://github.com/centrifuge/substrate): Clone it, `cd subkey` and run `cargo install --force --path .`
2. Run a centrifuge-chain locally:`docker run -p 9944:9944 -p 30333:30333 centrifugeio/centrifuge-chain:20190814150805-ddd3818 centrifuge-chain --ws-external --dev`
3. Now adjust the hardcoded const parameters in `test/main.go` according to your env + chain state.
4. Run `go run test/main.go`

- You need to install SubKey command from `https://github.com/centrifuge/substrate`
   - Clone it and `CD` to subkey module and run `cargo install --force --path .`
   
- Now adjust the hardcoded const parameters in `test/main.go` according to your env + chain state.
- Run `test/main.go`

## POC FLow

![Alt text](extrinsic-execution.png?raw=true "Extrinsic Execution")
