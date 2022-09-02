# smartcrud

## Description
对于使用orm工具操作数据库的用户，smartcrud会帮你自动化生成基础的增删查改

## Installation
```
$ go get github.com/LeonScript17/smartcrud

```

## Example
```
func main() {
 g := smartcrud.InitGen("/Users/leonscript/GolandProjects/smartcrud/test", "model", model.User{})
 err := g.GenerateCRUD()
 if err != nil {
  log.Println(err)
 }
}

```