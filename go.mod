module go_study

go 1.20

require github.com/yajium/calculator v0.0.0

require (
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.18 // indirect
	github.com/rs/zerolog v1.29.1 // indirect
	golang.org/x/sys v0.7.0 // indirect
)

// リモートではなく、ローカルディレクトリを使用するようにreplace
replace github.com/yajium/calculator => ../calculator
