package smartcrud

import (
	"github.com/Xuanwo/gg"
	"log"
	"os"
	"reflect"
	"strings"
)

// sgen 自动化生成器
type sgen struct {
}

type model struct {
	name string
}

// GenerateCRUD 使用
func (s *sgen) GenerateCRUD(config genConfig) error {
	// 使用反射解析用户传入的model
	models := s.analyseModel(config.srcModel)
	// 循环生成代码
	for _, m := range models {
		if err := s.generate(config.dstPath, config.srcPackage, m); err != nil {
			return err
		}
	}

	return nil
}

// analyseModel 反射分析代码
func (s *sgen) analyseModel(srts interface{}) (models []model) {
	var (
		t reflect.Type
	)

	for _, srt := range srts.([]interface{}) {
		t = reflect.TypeOf(srt)
		models = append(models, model{name: t.Name()})
	}

	return
}

// generate 代码生成
func (s *sgen) generate(dstPath, srcPackage string, model model) (err error) {
	// 拆解地址
	dstPaths := strings.Split(dstPath, "/")

	// 创造生成代码
	group := gg.NewGroup()
	group.AddPackage(dstPaths[len(dstPaths)-1])

	// imports
	group.NewImport().
		AddPath("gorm.io/gorm").
		AddPath("gorm.io/driver/mysql")

	// struct
	group.NewStruct(model.name+"Repository").AddField("db", "*gorm.DB")

	// method
	// insert
	group.NewFunction("Insert"+model.name).
		WithReceiver("r", "*"+model.name+"Repository").
		AddParameter("param", srcPackage+"."+model.name).
		AddResult("", "error").
		AddBody(gg.String(`return r.db.Create(&param).Error`))

	// Delete
	group.NewFunction("Delete"+model.name).
		WithReceiver("r", "*"+model.name+"Repository").
		AddParameter("param", srcPackage+"."+model.name).
		AddResult("", "error").
		AddBody(gg.String(`return r.db.Delete(&param).Error`))

	// Update
	group.NewFunction("Update"+model.name).
		WithReceiver("r", "*"+model.name+"Repository").
		AddParameter("param", srcPackage+"."+model.name).
		AddResult("", "error").
		AddBody(gg.String(`return r.db.Model(&` + srcPackage + "." + model.name + `{}).Updates(param).Error`))

	// Get
	group.NewFunction("Get"+model.name).
		WithReceiver("r", "*"+model.name+"Repository").
		AddResult("param", srcPackage+"."+model.name).
		AddResult("err", "error").
		AddBody(gg.String(`err = r.db.Take(&param).Error
								   return`))

	// Query
	group.NewFunction("Query"+model.name).
		WithReceiver("r", "*"+model.name+"Repository").
		AddResult("params", "[]"+srcPackage+"."+model.name).
		AddResult("err", "error").
		AddBody(gg.String(`err = r.db.Find(&params).Error
								   return params,err`))

	log.Println(group.String())

	// create file
	f, err := os.Create(dstPath + "/" + strings.ToLower(model.name) + ".go")
	if err != nil {
		return
	}
	_, err = f.WriteString(group.String())
	return
}
