package solution

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
// 官方方法
func isBadVersion(version int) bool {
	return false
}
func firstBadVersion(n int) int {
	l := 1
	r := n
	for l+1 < r {
		mid := l + (r-l)/2
		if isBadVersion(mid) {
			r = mid
		} else {
			l = mid
		}
	}
	if isBadVersion(l) {
		return l
	}
	if isBadVersion(r) {
		return r
	}
	return r
}
