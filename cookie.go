package timeoff

type Cookie struct {
	hashKey   []byte
	//hashFunc  func() hash.Hash
	blockKey  []byte
	//block     cipher.Block
	maxLength int
	maxAge    int64
	minAge    int64
	err       error
	timeFunc func() int64
}