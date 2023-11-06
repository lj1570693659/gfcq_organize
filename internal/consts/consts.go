package consts

const (
	ValidateSignatureError int = -40001
	ParseJsonError         int = -40002
	ComputeSignatureError  int = -40003
	IllegalAesKey          int = -40004
	ValidateCorpidError    int = -40005
	EncryptAESError        int = -40006
	DecryptAESError        int = -40007
	IllegalBuffer          int = -40008
	EncodeBase64Error      int = -40009
	DecodeBase64Error      int = -40010
	GenJsonError           int = -40011
	IllegalProtocolType    int = -40012
	AccessTokenKey             = "accessToken"
	AddressBookKey             = "user"
	ProductBookKey             = "product"
	PersonBookKey              = "person"
	CheckIn                    = "checkIn"

	GSHA1 = "gsha1"
	MD5   = "md5"
)
