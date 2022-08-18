package read_n_characters_given_read4

/**
 * The read4 API is already defined for you.
 *
 *     read4 := func(buf4 []byte) int
 *
 * // Below is an example of how the read4 API can be called.
 * file := File("abcdefghijk") // File is "abcdefghijk", initially file pointer (fp) points to 'a'
 * buf4 := make([]byte, 4) // Create buffer with enough space to store characters
 * read4(buf4) // read4 returns 4. Now buf = ['a','b','c','d'], fp points to 'e'
 * read4(buf4) // read4 returns 4. Now buf = ['e','f','g','h'], fp points to 'i'
 * read4(buf4) // read4 returns 3. Now buf = ['i','j','k',...], fp points to end of file
 */

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

var solution = func(read4 func([]byte) int) func([]byte, int) int {
	return func(buf []byte, n int) int {
		var (
			tmp   = make([]byte, 4)
			step  int
			index int
		)
		for index = 0; index < n; index += step {
			step = min(read4(tmp), n-index)
			if step == 0 {
				return len(buf)
			}

			for i := 0; i < step; i++ {
				buf[index+i] = tmp[i]
			}
		}
		return index
	}
}
