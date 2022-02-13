package maps

type Dictionary map[string]string

const (
	ErrNotFound     = DictionaryErr("could not find the word you were looking for")
	ErrWordExists   = DictionaryErr("cannot add word because it already exists")
	ErrWordNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (dictionary Dictionary) Search(key string) (string, error) {
	// 它可以返回两个值。第二个值是一个布尔值，表示是否成功找到 key。
	value, ok := dictionary[key]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}

// Add
//map是一个引用类型，所以不用指针传入
func (dictionary Dictionary) Add(key string, value string) error {
	_, found := dictionary[key]
	if found {
		return ErrWordExists
	}
	dictionary[key] = value
	return nil
}

func (dictionary Dictionary) Update(key, value string) error {
	_, found := dictionary[key]
	if !found {
		return ErrWordNotExist
	}
	dictionary[key] = value
	return nil
}

func (dictionary Dictionary) Delete(key string) {
	delete(dictionary, key)
}

/*var ap *int

func main() {
	a := 1  // define int
	b := 2  // define int
	ap = &a // set ap to address of a (&a)
	//   ap address: 0x2101f1018
	//  ap value  : 115
	*ap = 3 // change the value at address &a to 3
	//   ap address: 0x2101f1018
	//  ap value  : 3
	a = 4
	// change the value of a to 4
	//  ap address: 0x2101f1018
	//  ap value  : 4
	ap = &b
	//set ap to the address of b (&b)
	//  ap address: 0x2101f1020
	//  ap value  : 2
	ap2 := ap
	// set ap2 to the address in ap
	//  ap  address: 0x2101f1020
	//  ap  value  : 2
	//  ap2 address: 0x2101f1020
	//  ap2 value  : 2
	*ap = 5
	//change the value at the address &b to 5
	//  ap  address: 0x2101f1020
	//  ap  value  : 5
	//  ap2 address: 0x2101f1020
	//  ap2 value  : 5
	//If this was a reference ap & ap2 would now
	//have different values
	ap = &a
	//change ap to address of a (&a)
	//  ap  address: 0x2101f1018
	//  ap  value  : 4
	//  ap2 address: 0x2101f1020
	//  ap2 value  : 5
	//Since we've changed the address of ap, it now
	//has a different value then ap2
}
*/
