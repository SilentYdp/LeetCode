package main

import (
	"fmt"
)

func find(residue int,start int, )  {

}

//二分查找代码实现
func erfenchazhao(a [10]int,b int)  (index int) {
//	先跟数组中间的比，比之大则在右半部继续比，否则左半部

	length := len(a)
	var left int
	var right int
	left=0
	right=length-1
	var mid int

	for left<right {


		mid=(left+right)/2

		if b==a[mid]{
			return mid
		}
		if b<a[mid]{
			right=mid-1
		}
		if b>a[mid]{
			left=mid+1
		}
	}


	return -1

}

//LeetCode283---已通过
func moveZeroes(nums []int)  {
//	在原来的数组上进行操作
//双指针，因为前面都是非0，所以从头开始遍历，遇到不是0的就往后走，是0的就表示可以覆盖


var goindex int
var exchangeindex int

goindex = 0
exchangeindex =-1
	length := len(nums)

	for goindex<length {
		if nums[goindex]!=0{
			if exchangeindex>=0 {
				nums[exchangeindex] = nums[goindex]
				exchangeindex++

			}

		}else {
			if exchangeindex<0{
			exchangeindex=goindex}
		}
		goindex++
	}
	if exchangeindex>=0{
	for exchangeindex<=length-1{
		nums[exchangeindex]=0
		exchangeindex++
	}}


}
func moveZeroes2(nums []int){
	//元素交换
	length := len(nums)
	i:=0
	k:=0
	for i <length{
		if nums[i]!=0{
			if i!=k{
				nums[i],nums[k]=nums[k],nums[i]}
			k++
		}else {

		}
		i++
	}
}
func main()  {
	fmt.Println("测试")
	var a []int=[]int{-4,-1,0,3,10}
	//ret := removeDuplicates2(a)
	sortedSquares(a)
	fmt.Println(a)

}

//插入排序
func InsertSort(a []int, n int)  {
	//针对每一个数据都去看看前面有没有比自己小的，有的话，插在比自己小的元素的前面，
//	没有的话就不动

	var j int
 for i:=1;i<n;i++{
 	//如果待插入的数据比它之前的还大，那就不用排，继续看下一个
 	if a[i]>=a[i-1]{

	}else {
		for j = i - 1; j >=0; j-- {
			if a[j] <= a[i] {
				break
			}
		}
			temp := a[i]
			for m := i - 1; m > j; m-- {
				a[m+1] = a[m]
			}
			a[j+1] = temp

		}
	}

 }
//LeetCode 88	通过
func merge(nums1 []int, m int, nums2 []int, n int)  {
	var i int
	var j int
	i=0
	j=m
	for i<n{
		nums1[j]=nums2[i]
		i++
		j++
	}
	InsertSort(nums1,m+n)


}

//LeetCode 977  通过
func sortedSquares(A []int) []int {
	var j int
	n:=len(A)
	for i:=1;i<n;i++{
		//如果待插入的数据比它之前的还大，那就不用排，继续看下一个
		if abs(A[i])>=abs(A[i-1]){

		}else {
			for j = i - 1; j >=0; j-- {
				if abs(A[j]) <= abs(A[i]) {
					break
				}
			}
			temp := A[i]
			for m := i - 1; m > j; m-- {
				A[m+1] = A[m]
			}
			A[j+1] = temp

		}
	}
	for i:=0;i<n;i++{
		A[i]=A[i]*A[i]
	}
	return A
}
func abs(a int) int  {
	if a>=0{
		return a
	}else {
		return -a
	}
}







//LeetCode75 通过 建议二刷
func sortColors(nums []int)  {
	length:=len(nums)
	if length==0 || length==1{
		return
	}

	i:=0
	k1:=0
	k2:=length-1
	for i<=k2&&k1<length{
		if nums[i]==1 {
			i++
		}else if nums[i]==0{
			if nums[i]==nums[k1]{
				i++
				k1++
			}else {
					if nums[k1]==1{
						nums[k1], nums[i] = nums[i], nums[k1]
						i++
					}else {
						nums[k1], nums[i] = nums[i], nums[k1]
					}
					k1++
			}
		}else if nums[i]==2{
				if nums[i]==nums[k2]{
					k2--
				}else {
						if nums[k2]==1{
							nums[k2], nums[i] = nums[i], nums[k2]
							i++
						}else {
							nums[k2], nums[i] = nums[i], nums[k2]
						}
					k2--
				}
		}

	}
}




//LeetCode 80 通过
func removeDuplicates2(nums []int) int {
	length:=len(nums)
	if length==0{
		return 0
	}
	if length==1{
		return 1
	}
	k:=1
	count1:=0
	temp:=nums[0]
	i:=1

	count2:=1
	for i<length{
		if nums[i]==temp{
			count2++
			if count2==2{
				nums[k]=nums[i]
				k++
			}else {
					count1++
			}
		}else {
			temp=nums[i]
			count2=1
			nums[k]=nums[i]
			k++
		}
		i++
	}
	return length-count1
}

//LeetCode 26通过
func removeDuplicates(nums []int) int {

	length:=len(nums)
	if length==1{
		return 1
	}
	count:=0
	i:=1
	k:=1
	a:=nums[0]
	for i <length{
		if nums[i]==a{
			count++
		}else {
			a=nums[i]
			if i!=k{
				nums[i],nums[k]=nums[k],nums[i]
			}
			k++
		}
		i++
	}
	return length-count

}

//LeetCode 27 通过
func removeElement(nums []int, val int) int {
	length := len(nums)
	i:=0
	var k int
	k=0
	count:=0
	for i<length{
		if nums[i]==val{
			count++
		}else {
			if i!=k{
				nums[k],nums[i]=nums[i],nums[k]
			}else {

			}
			k++
		}
		i++
	}
	return length-count
}
