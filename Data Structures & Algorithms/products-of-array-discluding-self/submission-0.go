func productExceptSelf(nums []int) []int {
    l := len(nums)
    prefixProducts := make([]int, l)
    sufixProducts := make([]int, l)
    prefixProducts[0]=1
    for i:=1; i<l; i++ {
        prefixProducts[i]=nums[i-1]*prefixProducts[i-1]
    }
    sufixProducts[l-1] = 1
    for j:=len(nums)-2; j>=0; j-- {
        sufixProducts[j]=nums[j+1]*sufixProducts[j+1]
    }
    fmt.Println("p",prefixProducts)
    fmt.Println("s",sufixProducts)
    res := make([]int,len(nums))
    for i:=0; i<len(nums); i++ {
        res[i] = sufixProducts[i] * prefixProducts[i] 
    }
    fmt.Println(res)
    return res
}
