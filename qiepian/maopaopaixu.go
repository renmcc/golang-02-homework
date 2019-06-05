package qiepian

func Sort(slice []int, reverse bool) []int {
	//外层控制行
	for i:=0;i<len(slice)-1;i++ {
		//内层控制列
		for j:=0;j<len(slice)-1-i;j++{
			//对比相邻数据满足条件数据交换
			//true从小到大排序
			if reverse == true {
				if slice[j] > slice[j+1] {
					slice[j], slice[j+1] = slice[j+1], slice[j]
				}
			//false从大到小排序
			}else {
				if slice[j] < slice[j+1] {
					slice[j], slice[j+1] = slice[j+1], slice[j]
				}
			}
		}
	}
	return slice
}