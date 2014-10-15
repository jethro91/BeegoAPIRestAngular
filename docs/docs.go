package docs

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/swagger"
)

var rootinfo string = `{"apiVersion":"1.0.0","swaggerVersion":"1.2","apis":[{"path":"/auth_user","description":"oprations for AuthUser\n"},{"path":"/barang","description":"oprations for Barang\n"}],"info":{"title":"beego Test API","description":"beego has a very cool tools to autogenerate documents for your API","contact":"astaxie@gmail.com","termsOfServiceUrl":"http://beego.me/","license":"Url http://www.apache.org/licenses/LICENSE-2.0.html"}}`
var subapi string = `{"/auth_user":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/auth_user","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/","description":"","operations":[{"httpMethod":"POST","nickname":"Post","type":"","summary":"create AuthUser","parameters":[{"paramType":"body","name":"body","description":"\"body for AuthUser content\"","dataType":"AuthUser","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.AuthUser.Id","responseModel":""},{"code":403,"message":"body is empty","responseModel":""}]}]},{"path":"/:id","description":"","operations":[{"httpMethod":"GET","nickname":"Get","type":"","summary":"get AuthUser by id","parameters":[{"paramType":"path","name":"id","description":"\"The key for staticblock\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.AuthUser","responseModel":"AuthUser"},{"code":403,"message":":id is empty","responseModel":""}]}]},{"path":"/","description":"","operations":[{"httpMethod":"GET","nickname":"Get All","type":"","summary":"get AuthUser","parameters":[{"paramType":"query","name":"query","description":"\"Filter. e.g. col1:v1,col2:v2 ...\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"fields","description":"\"Fields returned. e.g. col1,col2 ...\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"sortby","description":"\"Sorted-by fields. e.g. col1,col2 ...\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"order","description":"\"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ...\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"limit","description":"\"Limit the size of result set. Must be an integer\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"offset","description":"\"Start position of result set. Must be an integer\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.AuthUser","responseModel":"AuthUser"},{"code":403,"message":"","responseModel":""}]}]},{"path":"/:id","description":"","operations":[{"httpMethod":"PUT","nickname":"Update","type":"","summary":"update the AuthUser","parameters":[{"paramType":"path","name":"id","description":"\"The id you want to update\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"body","name":"body","description":"\"body for AuthUser content\"","dataType":"AuthUser","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.AuthUser","responseModel":"AuthUser"},{"code":403,"message":":id is not int","responseModel":""}]}]},{"path":"/:id","description":"","operations":[{"httpMethod":"DELETE","nickname":"Delete","type":"","summary":"delete the AuthUser","parameters":[{"paramType":"path","name":"id","description":"\"The id you want to delete\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"success!","responseModel":""},{"code":403,"message":"id is empty","responseModel":""}]}]}],"models":{"AuthUser":{"id":"AuthUser","properties":{"Email":{"type":"string","description":"","format":""},"First":{"type":"string","description":"","format":""},"Id":{"type":"int","description":"","format":""},"Last":{"type":"string","description":"","format":""},"Password":{"type":"string","description":"","format":""},"RegDate":{"type":"\u0026{time Time}","description":"","format":""},"RegKey":{"type":"string","description":"","format":""}}}}},"/barang":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/barang","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/","description":"","operations":[{"httpMethod":"POST","nickname":"Post","type":"","summary":"create Barang","parameters":[{"paramType":"body","name":"body","description":"\"body for Barang content\"","dataType":"Barang","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.Barang.Id","responseModel":""},{"code":403,"message":"body is empty","responseModel":""}]}]},{"path":"/:id","description":"","operations":[{"httpMethod":"GET","nickname":"Get","type":"","summary":"get Barang by id","parameters":[{"paramType":"path","name":"id","description":"\"The key for staticblock\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.Barang","responseModel":"Barang"},{"code":403,"message":":id is empty","responseModel":""}]}]},{"path":"/","description":"","operations":[{"httpMethod":"GET","nickname":"Get All","type":"","summary":"get Barang","parameters":[{"paramType":"query","name":"query","description":"\"Filter. e.g. col1:v1,col2:v2 ...\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"fields","description":"\"Fields returned. e.g. col1,col2 ...\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"sortby","description":"\"Sorted-by fields. e.g. col1,col2 ...\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"order","description":"\"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ...\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"limit","description":"\"Limit the size of result set. Must be an integer\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"offset","description":"\"Start position of result set. Must be an integer\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.Barang","responseModel":"Barang"},{"code":403,"message":"","responseModel":""}]}]},{"path":"/:id","description":"","operations":[{"httpMethod":"PUT","nickname":"Update","type":"","summary":"update the Barang","parameters":[{"paramType":"path","name":"id","description":"\"The id you want to update\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"body","name":"body","description":"\"body for Barang content\"","dataType":"Barang","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.Barang","responseModel":"Barang"},{"code":403,"message":":id is not int","responseModel":""}]}]},{"path":"/:id","description":"","operations":[{"httpMethod":"DELETE","nickname":"Delete","type":"","summary":"delete the Barang","parameters":[{"paramType":"path","name":"id","description":"\"The id you want to delete\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"success!","responseModel":""},{"code":403,"message":"id is empty","responseModel":""}]}]}],"models":{"Barang":{"id":"Barang","properties":{"Id":{"type":"int","description":"","format":""},"Keterangan":{"type":"string","description":"","format":""},"Nama":{"type":"string","description":"","format":""},"Stock":{"type":"int","description":"","format":""}}}}}}`
var rootapi swagger.ResourceListing

var apilist map[string]*swagger.ApiDeclaration

func init() {
	basepath := "/v1"
	err := json.Unmarshal([]byte(rootinfo), &rootapi)
	if err != nil {
		beego.Error(err)
	}
	err = json.Unmarshal([]byte(subapi), &apilist)
	if err != nil {
		beego.Error(err)
	}
	beego.GlobalDocApi["Root"] = rootapi
	for k, v := range apilist {
		for i, a := range v.Apis {
			a.Path = urlReplace(k + a.Path)
			v.Apis[i] = a
		}
		v.BasePath = basepath
		beego.GlobalDocApi[strings.Trim(k, "/")] = v
	}
}


func urlReplace(src string) string {
	pt := strings.Split(src, "/")
	for i, p := range pt {
		if len(p) > 0 {
			if p[0] == ':' {
				pt[i] = "{" + p[1:] + "}"
			} else if p[0] == '?' && p[1] == ':' {
				pt[i] = "{" + p[2:] + "}"
			}
		}
	}
	return strings.Join(pt, "/")
}
