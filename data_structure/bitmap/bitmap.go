package bitmap

// BitMap represents bitmap.
type BitMap []byte

// New returns a new BitMap.
func New(length uint) BitMap {
	return make([]byte, length/8+1)
}

// Set stores value in bitmap.
func (bm BitMap) Set(value uint) {
	byteIndex := value / 8
	if byteIndex >= uint(len(bm)) {
		return
	}

	bitIndex := value % 8
	bm[byteIndex] |= 1 << bitIndex
}

// Get check if value exists in the bitmap.
func (bm BitMap) Get(value uint) bool {
	byteIndex := value / 8
	if byteIndex >= uint(len(bm)) {
		return false
	}

	bitIndex := value % 8
	return bm[byteIndex]&(1<<bitIndex) != 0
}
