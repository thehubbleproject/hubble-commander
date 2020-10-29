package core

var defaultHashes []ByteArray

// depicts the empty leaf in balance tree
var ZERO_VALUE_LEAF ByteArray

func init() {
	var err error
	ZERO_VALUE_LEAF, err = HexToByteArray("0x290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563")
	if err != nil {
		panic(err)
	}

	defaultHashes, err = GenDefaultHashes(100)
	if err != nil {
		panic(err)
	}
}

func GenDefaultHashes(depth int) ([]ByteArray, error) {
	hashes := make([]ByteArray, depth)
	hashes[0] = ZERO_VALUE_LEAF
	for i := 1; i < depth; i++ {
		parent, err := GetParent(hashes[i-1], hashes[i-1])
		if err != nil {
			return hashes, err
		}
		hashes[i] = parent
	}
	return hashes, nil
}
