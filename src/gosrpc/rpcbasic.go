package gosrpc
type Rect struct{}

type Params struct {
	Width, Height int;
}
//函数必须是导出的
//必须有两个导出类型参数
//第一个参数是接收参数
//第二个参数是返回给客户端参数，必须是指针类型
//函数还要有一个返回值error
func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Width * p.Height;
	return nil;
}

func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Width + p.Height) * 2;
	return nil;
}