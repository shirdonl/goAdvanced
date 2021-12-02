package pkg

import "sync"

//创建过程数据的存储容器
type WorkProcessData struct {
	Data map[string]string
	mux  sync.RWMutex
}

//添加数据到容器
func (s *WorkProcessData) AddData(key, value string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	if s.Data == nil {
		s.Data = make(map[string]string)
	}

	s.Data[key] = value
}

//从容器获取数据
func (s *WorkProcessData) GetData() string {
	return s.Data["1"] + "," + s.Data["2"]
}
