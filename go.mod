module go_study

go 1.20

require github.com/yajium/calculator v0.0.0

// リモートではなく、ローカルディレクトリを使用するようにreplace
replace github.com/yajium/calculator => ../calculator
