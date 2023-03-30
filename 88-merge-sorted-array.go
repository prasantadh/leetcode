func merge(nums1 []int, m int, nums2 []int, n int)  {
    k := m + n - 1
    m = m - 1
    n = n - 1
    for k >= 0 {
        if m < 0 {
            nums1[k] = nums2[n]
            n--
        } else if n < 0 {
            nums1[k] = nums1[m]
            m--
        } else {
            if nums1[m] > nums2[n] {
                nums1[k] = nums1[m]
                m--
            } else {
                nums1[k] = nums2[n]
                n--
            }
        }
        k--
    }  

}
