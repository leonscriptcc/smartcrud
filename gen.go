package smartcrud

// genConfig 配置
type genConfig struct {
	dstPath    string
	srcPackage string
	srcModel   []interface{}
}

// GenService 生成器行为
type GenService interface {
	// GenerateCRUD 生成基础sql代码
	GenerateCRUD(genConfig) error
}

// Generator 生成器
type Generator struct {
	Config genConfig
	GenService
}

// InitGen 初始化生成器
func InitGen(dstPath string, srcPackage string, srcModel ...interface{}) Generator {
	models := make([]interface{}, 0, len(srcModel))
	for _, s := range srcModel {
		models = append(models, s)
	}
	return Generator{
		GenService: &sgen{},
		Config: genConfig{
			dstPath:    dstPath,
			srcPackage: srcPackage,
			srcModel:   models,
		},
	}
}

func (g *Generator) GenerateCRUD() (err error) {
	return g.GenService.GenerateCRUD(g.Config)
}
