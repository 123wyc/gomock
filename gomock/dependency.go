package gomock

/**
将需要mock的方法 接口化
*/

//go:generate mockgen -source=dependency.go -destination=mock_dependency.go -package=gomock
type Repository interface {
	GetName() string
}
