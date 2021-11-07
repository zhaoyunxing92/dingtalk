package response

type GetExtContactLabel struct {
	Response

	Results []labels `json:"results"`
}

type labels struct {
	//标签组名字
	Name string `json:"name"`

	//标签组颜色
	Color int `json:"color"`

	//标签
	Labels []label `json:"labels"`
}

type label struct {
	//标签名称
	Name string `json:"name"`

	//标签id
	Id int `json:"id"`
}
