package geometry

// カプセル化はパッケージ間でのみ有効

// 構造体の名前と大文字にすると中身のプロパティがprivateになって直接アクセスできない
type Triangle struct {
	size int
}

// メソッドの最初を小文字にするとprivateと同じになる
func (t *Triangle) doubleSize() {
	t.size *= 2
}

// メソッドの最初を大文字にするとpublicと同じになる
func (t *Triangle) SetSize(size int) {
	t.size = size
}

func (t *Triangle) Perimeter() int {
	t.doubleSize()
	return t.size * 3
}
