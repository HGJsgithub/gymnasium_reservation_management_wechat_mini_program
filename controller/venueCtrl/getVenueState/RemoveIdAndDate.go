package getVenueState

func RemoveIdAndDate(rawState [][]byte, timeNum int) (vst [][]bool) {
	var vsTable [][]bool
	for i := range rawState {
		//vs是每个场地不同时间的状态，也就是数据库里的一行，预约界面里的一列
		var vs []bool
		//因为id和date会各占据8和16位，后面的才是状态，因此j从24开始
		for j := 24; j < 24+timeNum; j++ {
			if rawState[i][j] == 1 {
				vs = append(vs, true)
			} else if rawState[i][j] == 0 {
				vs = append(vs, false)
			}
		}
		vsTable = append(vsTable, vs)
	}
	return vsTable
}
