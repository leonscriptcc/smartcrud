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
	GenService
}

// InitGen 初始化生成器
func InitGen(dstPath string, srcPackage string, srcModel ...interface{}) Generator {
	return Generator{
		GenService: &sgen{
			genConfig: genConfig{
				dstPath:    dstPath,
				srcPackage: srcPackage,
				srcModel:   srcModel,
			},
		},
	}
}
