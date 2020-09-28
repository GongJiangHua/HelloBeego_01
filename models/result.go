package models
/*公共的用于返回结构体的类型定义，Data表示任意类型

 */
type Result struct {
	Code int//接口返回状态类型
	Message string//接口返回状态对应的描述信息
	Data interface{}//返回的数据
}
