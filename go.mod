module github.com/loopholelabs/scale-ext-fetch

go 1.20

replace github.com/loopholelabs/scale => ../scale/scale

require (
	github.com/loopholelabs/polyglot v1.1.1
	github.com/tetratelabs/wazero v1.2.1
)

require github.com/stretchr/testify v1.8.1 // indirect
