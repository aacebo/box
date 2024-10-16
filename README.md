<p align="center">
	<img src="./assets/icon.png" width="120px" />
</p>
 
<p align="center">
	a zero dependency performant dependency injection library
</p>

<p align="center">
	<a href="https://opensource.org/licenses/MIT" target="_blank" alt="License">
		<img src="https://img.shields.io/badge/License-MIT-blue.svg" />
	</a>
	<a href="https://pkg.go.dev/github.com/aacebo/box" target="_blank" alt="Go Reference">
		<img src="https://pkg.go.dev/badge/github.com/aacebo/box.svg" />
	</a>
	<a href="https://goreportcard.com/report/github.com/aacebo/box" target="_blank" alt="Go Report Card">
		<img src="https://goreportcard.com/badge/github.com/aacebo/box" />
	</a>
	<a href="https://github.com/aacebo/box/actions/workflows/ci.yml" target="_blank" alt="Build">
		<img src="https://github.com/aacebo/box/actions/workflows/ci.yml/badge.svg?branch=main" />
	</a>
	<a href="https://codecov.io/gh/aacebo/box" > 
		<img src="https://codecov.io/gh/aacebo/box/graph/badge.svg?token=9XETRUUQUY" /> 
	</a>
</p>

# Install

```bash
go get github.com/aacebo/box
```

# Usage

```go
b := box.New()
b.Put(&ServiceA{}, &ServiceB{})

fn, err := b.Inject(func (a *ServiceA, b *Service B) {
	fmt.Println(a, b)
})

if err != nil {
	panic(err)
}

fn()
```

# Benchmarks

- coming soon!
