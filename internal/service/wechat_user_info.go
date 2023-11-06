package service

//type (
//	IWechatUserInfo interface {
//		Create(ctx context.Context, in *v1.CreateWechatUserInfoReq) (*v1.CreateWechatUserInfoRes, error)
//		GetOne(ctx context.Context, in *v1.GetOneWechatUserInfoReq) (*v1.GetOneWechatUserInfoRes, error)
//		GetList(ctx context.Context, in *v1.GetListWechatUserInfoReq) (*v1.GetListWechatUserInfoRes, error)
//		GetAll(ctx context.Context, in *v1.GetAllWechatUserInfoReq) (*v1.GetAllWechatUserInfoRes, error)
//		Modify(ctx context.Context, in *v1.ModifyWechatUserInfoReq) (*v1.ModifyWechatUserInfoRes, error)
//		Delete(ctx context.Context, in *v1.DeleteWechatUserInfoReq) (*v1.DeleteWechatUserInfoRes, error)
//		//GetUser(ctx context.Context) error
//		//DeleteById(ctx context.Context, uid uint64) error
//	}
//)
//
//var (
//	localWechatUserInfo IWechatUserInfo
//)
//
//func WechatUserInfo() IWechatUserInfo {
//	if localWechatUserInfo == nil {
//		panic("implement not found for interface IWechatUserInfo, forgot register?")
//	}
//	return localWechatUserInfo
//}
//
//func RegisterWechatUserInfo(i IWechatUserInfo) {
//	localWechatUserInfo = i
//}
