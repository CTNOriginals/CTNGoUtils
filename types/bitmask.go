package ctntype

type TBitMask uint

func (this TBitMask) Has(flag TBitMask) bool {
	return this&flag != 0
}
func (this *TBitMask) Set(flag TBitMask) {
	*this = *this | flag
}
func (this *TBitMask) Toggle(flag TBitMask) {
	*this = *this ^ flag
}
func (this *TBitMask) Clear(flag TBitMask) {
	*this = *this &^ flag
}

// Preforms this AND flag and stores result in this.
func (this *TBitMask) Compare(flag TBitMask) {
	*this = *this & flag
}

// Set this to flag if condition is true
func (this *TBitMask) SetIf(flag TBitMask, condition bool) {
	if !condition {
		return
	}

	this.Set(flag)
}
