package solution

var isBadVersion func(version int) bool

func firstBadVersion(n int) int {
	l, r := 1, n
	for l < r {
		m := l + ((r - l) / 2)
		if isBadVersion(m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

/* Interpreted to C++ for varification
bool isBadVersion(int version);

class Solution {
public:
    int firstBadVersion(int n) {
        int l = 1, r = n;
        while (l < r) {
            int m = l + (r - l) / 2;
            if (isBadVersion(m)) {
                r = m;
            } else {
                l = m + 1;
            }
        }
        return l;
    }
};
*/
