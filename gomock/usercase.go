package gomock

/**
业务比较密集的service层
*/
type Service struct {
	repository Repository
}

var (
	entry = make(map[string]interface{})
)

func (s Service) GetName() (name string) {

	if value, ok := entry["name"]; ok {
		if _, ok := value.(string); ok {
			name, _ = entry["name"].(string)
		}
	}

	if _, ok := entry["name"]; !ok {
		name = s.repository.GetName()
		entry["name"] = name
		return
	}

	return
}
