package utils

func GetVenueStateTable(rawState [][]byte, timeNum int) (vst [][]bool) {
	var venueStateTable [][]bool
	for i, _ := range rawState {
		//vs是每个场地不同时间的状态，也就是数据库里的一行，预约界面里的一列
		var vs []bool
		//因为id会占据8位，后面的才是状态，因此j从8开始
		for j := 8; j < 8+timeNum; j++ {
			if rawState[i][j] == 1 {
				vs = append(vs, true)
			} else if rawState[i][j] == 0 {
				vs = append(vs, false)
			}
		}
		venueStateTable = append(venueStateTable, vs)
	}
	return venueStateTable
}
