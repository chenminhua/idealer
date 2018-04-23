package main

type IDService struct {
	Generators map[string]*IDGenerator
}

// todo 改成从数据库读配置
func CreateIdService() *IDService {
	materialLotGeneraotr := &IDGenerator{Category: "material_lot"}
	materialLotGeneraotr.Init()
	return &IDService{
		Generators: map[string] *IDGenerator {
			"material_lot": materialLotGeneraotr,
		},
	}
}

func (idService *IDService) GetNewId(category string) (uint64, error) {
	idGenerator := idService.Generators["material_lot"]
	return idGenerator.GetNewId()
}