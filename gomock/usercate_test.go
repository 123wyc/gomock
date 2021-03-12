package gomock

import (
	"github.com/golang/mock/gomock"
	"testing"
)

/**
  使用gomock 完成单元测试的demo
*/

//
type MockRepositoryOption func(repository *MockRepository)

//参数结构体
type args struct {
	name string
}

func TestMockRepository_GetName(t *testing.T) {

	var argsDefault = args{
		name: "cwl",
	}
	//测试当前方法时 模拟不同情况的一个列表
	tests := []struct {
		name            string                 //测试情况的描述
		args            args                   //参数
		mockRepoOptions []MockRepositoryOption //用到的mock
		want            string                 //想要的结果
		wantErr         error                  //想要的错误
	}{
		{
			name: "获取名字-走数据库",
			args: argsDefault,
			mockRepoOptions: []MockRepositoryOption{
				func(repository *MockRepository) {
					repository.EXPECT().GetName().Return("cwl") //模拟执行业务代码中调用repository的代码
				},
			},
			want:    "cwl",
			wantErr: nil,
		},
	}

	//执行上述测试的不同情况
	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			//创建一个mock
			var mockInterface = NewMockRepository(ctrl)

			for _, v := range test.mockRepoOptions {
				v(mockInterface)
			}
			//创建一个service
			s := &Service{
				repository: mockInterface,
			}
			//调用实际要测试的usercase中的方法
			s.GetName()
		})

	}
}
