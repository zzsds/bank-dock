// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/pingan.proto

package proto_pingan

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TranFun int32

const (
	TranFun_Default TranFun = 0
	// 会员签解约维护
	TranFun_Contract TranFun = 1303
	// 入金（银行发起）
	TranFun_BankCashIn TranFun = 1310
	// 出金（银行发起）
	TranFun_BankCashOut TranFun = 1312
	// 出金（交易网发起）
	TranFun_CashOut TranFun = 1318
	// 银行发起异步出金
	TranFun_BankCashOutConfirm TranFun = 1317
	// 清算文件
	TranFun_LiquidationFile TranFun = 1000
	// 清算失败文件
	TranFun_LiquidationFailFile TranFun = 1001
	// 对账不平记录文件
	TranFun_UnevenCheckRecordFile TranFun = 1007
	// 交易网发起会员余额对账
	TranFun_MemberBalanceCheck TranFun = 1002
	// 交易网触发银行进行清算与对账
	TranFun_TriggerBankCheck TranFun = 1003
	// 交易网查询银行清算与对账文件的进度
	TranFun_TriggerBankCheckProgress TranFun = 1004
	// 银行通知交易网查看文件
	TranFun_BankNotifyViewFile TranFun = 1005
	// 交易网发起出入金流水对账及会员开销户流水匹配
	TranFun_CashFlowCheckAndSignOutMatch TranFun = 1006
	// 查银行端会员资金台帐余额【1010】
	TranFun_QueryBankMemberBalance TranFun = 1010
	// 资金汇总账号余额查询【1022】
	TranFun_QuerySummaryBalance TranFun = 1022
	// 出入金账户信息查询【1012】
	TranFun_QueryInOutAccount TranFun = 1012
	// 查时间段会员开销户信息【1016】
	TranFun_QueryBetweenMemberSignOut TranFun = 1016
	// 查交易网端会员管理账户余额【1019】
	TranFun_BankQueryMemberBalance TranFun = 1019
	// 查询时间段会员出入金流水信息【1325】
	TranFun_QueryBetweenMemberInOutDetailed TranFun = 1325
	// 查询时间段会员交易流水信息【1324】
	TranFun_QueryBeTweenMemberDetailed TranFun = 1324
	// 签到签退【1330】
	TranFun_SignUpOut TranFun = 1330
	// 子账户间支付【1332】
	TranFun_AccountBetweenPay TranFun = 1332
	// 子账户间划转【1028】
	TranFun_AccountBetweenTransfer TranFun = 1028
	// 子账户冻结解冻【1029】
	TranFun_AccountFrozenAndUnFrozen TranFun = 1029
	// 平台收费与退费【1030】
	TranFun_PlatformChargeRefund TranFun = 1030
	// 平台支付与收取【1031】
	TranFun_PlatformPayCollect TranFun = 1031
	// 查询账户回单明细【1351】
	TranFun_QueryAccountReceiptDetaild TranFun = 1351
	// 查询回单号验证码【1352】
	TranFun_QueryReceiptCode TranFun = 1352
)

var TranFun_name = map[int32]string{
	0:    "Default",
	1303: "Contract",
	1310: "BankCashIn",
	1312: "BankCashOut",
	1318: "CashOut",
	1317: "BankCashOutConfirm",
	1000: "LiquidationFile",
	1001: "LiquidationFailFile",
	1007: "UnevenCheckRecordFile",
	1002: "MemberBalanceCheck",
	1003: "TriggerBankCheck",
	1004: "TriggerBankCheckProgress",
	1005: "BankNotifyViewFile",
	1006: "CashFlowCheckAndSignOutMatch",
	1010: "QueryBankMemberBalance",
	1022: "QuerySummaryBalance",
	1012: "QueryInOutAccount",
	1016: "QueryBetweenMemberSignOut",
	1019: "BankQueryMemberBalance",
	1325: "QueryBetweenMemberInOutDetailed",
	1324: "QueryBeTweenMemberDetailed",
	1330: "SignUpOut",
	1332: "AccountBetweenPay",
	1028: "AccountBetweenTransfer",
	1029: "AccountFrozenAndUnFrozen",
	1030: "PlatformChargeRefund",
	1031: "PlatformPayCollect",
	1351: "QueryAccountReceiptDetaild",
	1352: "QueryReceiptCode",
}

var TranFun_value = map[string]int32{
	"Default":                         0,
	"Contract":                        1303,
	"BankCashIn":                      1310,
	"BankCashOut":                     1312,
	"CashOut":                         1318,
	"BankCashOutConfirm":              1317,
	"LiquidationFile":                 1000,
	"LiquidationFailFile":             1001,
	"UnevenCheckRecordFile":           1007,
	"MemberBalanceCheck":              1002,
	"TriggerBankCheck":                1003,
	"TriggerBankCheckProgress":        1004,
	"BankNotifyViewFile":              1005,
	"CashFlowCheckAndSignOutMatch":    1006,
	"QueryBankMemberBalance":          1010,
	"QuerySummaryBalance":             1022,
	"QueryInOutAccount":               1012,
	"QueryBetweenMemberSignOut":       1016,
	"BankQueryMemberBalance":          1019,
	"QueryBetweenMemberInOutDetailed": 1325,
	"QueryBeTweenMemberDetailed":      1324,
	"SignUpOut":                       1330,
	"AccountBetweenPay":               1332,
	"AccountBetweenTransfer":          1028,
	"AccountFrozenAndUnFrozen":        1029,
	"PlatformChargeRefund":            1030,
	"PlatformPayCollect":              1031,
	"QueryAccountReceiptDetaild":      1351,
	"QueryReceiptCode":                1352,
}

func (x TranFun) String() string {
	return proto.EnumName(TranFun_name, int32(x))
}

func (TranFun) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_83707be2c8bfa1fb, []int{0}
}

type HeaderMessage struct {
	TranFunc             string   `protobuf:"bytes,1,opt,name=tranFunc,proto3" json:"tranFunc,omitempty"`
	ServType             string   `protobuf:"bytes,2,opt,name=servType,proto3" json:"servType,omitempty"`
	MacCode              string   `protobuf:"bytes,3,opt,name=macCode,proto3" json:"macCode,omitempty"`
	TranDate             string   `protobuf:"bytes,4,opt,name=tranDate,proto3" json:"tranDate,omitempty"`
	TranTime             string   `protobuf:"bytes,5,opt,name=tranTime,proto3" json:"tranTime,omitempty"`
	RspCode              string   `protobuf:"bytes,6,opt,name=rspCode,proto3" json:"rspCode,omitempty"`
	RspMsg               string   `protobuf:"bytes,7,opt,name=rspMsg,proto3" json:"rspMsg,omitempty"`
	ConFlag              string   `protobuf:"bytes,8,opt,name=conFlag,proto3" json:"conFlag,omitempty"`
	Length               string   `protobuf:"bytes,9,opt,name=length,proto3" json:"length,omitempty"`
	CounterId            string   `protobuf:"bytes,10,opt,name=counterId,proto3" json:"counterId,omitempty"`
	ThirdLogNo           string   `protobuf:"bytes,11,opt,name=thirdLogNo,proto3" json:"thirdLogNo,omitempty"`
	Qydm                 string   `protobuf:"bytes,12,opt,name=qydm,proto3" json:"qydm,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeaderMessage) Reset()         { *m = HeaderMessage{} }
func (m *HeaderMessage) String() string { return proto.CompactTextString(m) }
func (*HeaderMessage) ProtoMessage()    {}
func (*HeaderMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_83707be2c8bfa1fb, []int{0}
}

func (m *HeaderMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeaderMessage.Unmarshal(m, b)
}
func (m *HeaderMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeaderMessage.Marshal(b, m, deterministic)
}
func (m *HeaderMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderMessage.Merge(m, src)
}
func (m *HeaderMessage) XXX_Size() int {
	return xxx_messageInfo_HeaderMessage.Size(m)
}
func (m *HeaderMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderMessage.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderMessage proto.InternalMessageInfo

func (m *HeaderMessage) GetTranFunc() string {
	if m != nil {
		return m.TranFunc
	}
	return ""
}

func (m *HeaderMessage) GetServType() string {
	if m != nil {
		return m.ServType
	}
	return ""
}

func (m *HeaderMessage) GetMacCode() string {
	if m != nil {
		return m.MacCode
	}
	return ""
}

func (m *HeaderMessage) GetTranDate() string {
	if m != nil {
		return m.TranDate
	}
	return ""
}

func (m *HeaderMessage) GetTranTime() string {
	if m != nil {
		return m.TranTime
	}
	return ""
}

func (m *HeaderMessage) GetRspCode() string {
	if m != nil {
		return m.RspCode
	}
	return ""
}

func (m *HeaderMessage) GetRspMsg() string {
	if m != nil {
		return m.RspMsg
	}
	return ""
}

func (m *HeaderMessage) GetConFlag() string {
	if m != nil {
		return m.ConFlag
	}
	return ""
}

func (m *HeaderMessage) GetLength() string {
	if m != nil {
		return m.Length
	}
	return ""
}

func (m *HeaderMessage) GetCounterId() string {
	if m != nil {
		return m.CounterId
	}
	return ""
}

func (m *HeaderMessage) GetThirdLogNo() string {
	if m != nil {
		return m.ThirdLogNo
	}
	return ""
}

func (m *HeaderMessage) GetQydm() string {
	if m != nil {
		return m.Qydm
	}
	return ""
}

type ContractRequest1303 struct {
	FuncFlag             string   `protobuf:"bytes,1,opt,name=funcFlag,proto3" json:"funcFlag,omitempty"`
	SupAcctId            string   `protobuf:"bytes,2,opt,name=supAcctId,proto3" json:"supAcctId,omitempty"`
	CustAcctId           string   `protobuf:"bytes,3,opt,name=custAcctId,proto3" json:"custAcctId,omitempty"`
	CustName             string   `protobuf:"bytes,4,opt,name=custName,proto3" json:"custName,omitempty"`
	ThirdCustId          string   `protobuf:"bytes,5,opt,name=thirdCustId,proto3" json:"thirdCustId,omitempty"`
	IdType               string   `protobuf:"bytes,6,opt,name=idType,proto3" json:"idType,omitempty"`
	IdCode               string   `protobuf:"bytes,7,opt,name=idCode,proto3" json:"idCode,omitempty"`
	RelatedAcctId        string   `protobuf:"bytes,8,opt,name=relatedAcctId,proto3" json:"relatedAcctId,omitempty"`
	AcctFlag             string   `protobuf:"bytes,9,opt,name=acctFlag,proto3" json:"acctFlag,omitempty"`
	TranType             string   `protobuf:"bytes,10,opt,name=tranType,proto3" json:"tranType,omitempty"`
	AcctName             string   `protobuf:"bytes,11,opt,name=acctName,proto3" json:"acctName,omitempty"`
	BankCode             string   `protobuf:"bytes,12,opt,name=bankCode,proto3" json:"bankCode,omitempty"`
	BankName             string   `protobuf:"bytes,13,opt,name=bankName,proto3" json:"bankName,omitempty"`
	OldRelatedAcctId     string   `protobuf:"bytes,14,opt,name=oldRelatedAcctId,proto3" json:"oldRelatedAcctId,omitempty"`
	Reserve              string   `protobuf:"bytes,15,opt,name=reserve,proto3" json:"reserve,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContractRequest1303) Reset()         { *m = ContractRequest1303{} }
func (m *ContractRequest1303) String() string { return proto.CompactTextString(m) }
func (*ContractRequest1303) ProtoMessage()    {}
func (*ContractRequest1303) Descriptor() ([]byte, []int) {
	return fileDescriptor_83707be2c8bfa1fb, []int{1}
}

func (m *ContractRequest1303) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractRequest1303.Unmarshal(m, b)
}
func (m *ContractRequest1303) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractRequest1303.Marshal(b, m, deterministic)
}
func (m *ContractRequest1303) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractRequest1303.Merge(m, src)
}
func (m *ContractRequest1303) XXX_Size() int {
	return xxx_messageInfo_ContractRequest1303.Size(m)
}
func (m *ContractRequest1303) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractRequest1303.DiscardUnknown(m)
}

var xxx_messageInfo_ContractRequest1303 proto.InternalMessageInfo

func (m *ContractRequest1303) GetFuncFlag() string {
	if m != nil {
		return m.FuncFlag
	}
	return ""
}

func (m *ContractRequest1303) GetSupAcctId() string {
	if m != nil {
		return m.SupAcctId
	}
	return ""
}

func (m *ContractRequest1303) GetCustAcctId() string {
	if m != nil {
		return m.CustAcctId
	}
	return ""
}

func (m *ContractRequest1303) GetCustName() string {
	if m != nil {
		return m.CustName
	}
	return ""
}

func (m *ContractRequest1303) GetThirdCustId() string {
	if m != nil {
		return m.ThirdCustId
	}
	return ""
}

func (m *ContractRequest1303) GetIdType() string {
	if m != nil {
		return m.IdType
	}
	return ""
}

func (m *ContractRequest1303) GetIdCode() string {
	if m != nil {
		return m.IdCode
	}
	return ""
}

func (m *ContractRequest1303) GetRelatedAcctId() string {
	if m != nil {
		return m.RelatedAcctId
	}
	return ""
}

func (m *ContractRequest1303) GetAcctFlag() string {
	if m != nil {
		return m.AcctFlag
	}
	return ""
}

func (m *ContractRequest1303) GetTranType() string {
	if m != nil {
		return m.TranType
	}
	return ""
}

func (m *ContractRequest1303) GetAcctName() string {
	if m != nil {
		return m.AcctName
	}
	return ""
}

func (m *ContractRequest1303) GetBankCode() string {
	if m != nil {
		return m.BankCode
	}
	return ""
}

func (m *ContractRequest1303) GetBankName() string {
	if m != nil {
		return m.BankName
	}
	return ""
}

func (m *ContractRequest1303) GetOldRelatedAcctId() string {
	if m != nil {
		return m.OldRelatedAcctId
	}
	return ""
}

func (m *ContractRequest1303) GetReserve() string {
	if m != nil {
		return m.Reserve
	}
	return ""
}

type ContractResponse1303 struct {
	ThirdLogNo           string   `protobuf:"bytes,1,opt,name=thirdLogNo,proto3" json:"thirdLogNo,omitempty"`
	Teserve              string   `protobuf:"bytes,2,opt,name=teserve,proto3" json:"teserve,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContractResponse1303) Reset()         { *m = ContractResponse1303{} }
func (m *ContractResponse1303) String() string { return proto.CompactTextString(m) }
func (*ContractResponse1303) ProtoMessage()    {}
func (*ContractResponse1303) Descriptor() ([]byte, []int) {
	return fileDescriptor_83707be2c8bfa1fb, []int{2}
}

func (m *ContractResponse1303) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractResponse1303.Unmarshal(m, b)
}
func (m *ContractResponse1303) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractResponse1303.Marshal(b, m, deterministic)
}
func (m *ContractResponse1303) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractResponse1303.Merge(m, src)
}
func (m *ContractResponse1303) XXX_Size() int {
	return xxx_messageInfo_ContractResponse1303.Size(m)
}
func (m *ContractResponse1303) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractResponse1303.DiscardUnknown(m)
}

var xxx_messageInfo_ContractResponse1303 proto.InternalMessageInfo

func (m *ContractResponse1303) GetThirdLogNo() string {
	if m != nil {
		return m.ThirdLogNo
	}
	return ""
}

func (m *ContractResponse1303) GetTeserve() string {
	if m != nil {
		return m.Teserve
	}
	return ""
}

type CashOutRequest1318 struct {
	TranWebName          string   `protobuf:"bytes,1,opt,name=tranWebName,proto3" json:"tranWebName,omitempty"`
	ThirdCustId          string   `protobuf:"bytes,2,opt,name=thirdCustId,proto3" json:"thirdCustId,omitempty"`
	IdType               string   `protobuf:"bytes,3,opt,name=idType,proto3" json:"idType,omitempty"`
	IdCode               string   `protobuf:"bytes,4,opt,name=idCode,proto3" json:"idCode,omitempty"`
	TranOutType          string   `protobuf:"bytes,5,opt,name=tranOutType,proto3" json:"tranOutType,omitempty"`
	CustAcctId           string   `protobuf:"bytes,6,opt,name=custAcctId,proto3" json:"custAcctId,omitempty"`
	CustName             string   `protobuf:"bytes,7,opt,name=custName,proto3" json:"custName,omitempty"`
	SupAcctId            string   `protobuf:"bytes,8,opt,name=supAcctId,proto3" json:"supAcctId,omitempty"`
	TranType             string   `protobuf:"bytes,9,opt,name=tranType,proto3" json:"tranType,omitempty"`
	OutAcctId            string   `protobuf:"bytes,10,opt,name=outAcctId,proto3" json:"outAcctId,omitempty"`
	OutAcctIdName        string   `protobuf:"bytes,11,opt,name=outAcctIdName,proto3" json:"outAcctIdName,omitempty"`
	OutAcctIdBankName    string   `protobuf:"bytes,12,opt,name=outAcctIdBankName,proto3" json:"outAcctIdBankName,omitempty"`
	OutAcctIdBankCode    string   `protobuf:"bytes,13,opt,name=outAcctIdBankCode,proto3" json:"outAcctIdBankCode,omitempty"`
	Address              string   `protobuf:"bytes,14,opt,name=address,proto3" json:"address,omitempty"`
	CcyCode              string   `protobuf:"bytes,15,opt,name=ccyCode,proto3" json:"ccyCode,omitempty"`
	TranAmount           string   `protobuf:"bytes,16,opt,name=tranAmount,proto3" json:"tranAmount,omitempty"`
	FeeOutCustId         string   `protobuf:"bytes,17,opt,name=feeOutCustId,proto3" json:"feeOutCustId,omitempty"`
	Reserve              string   `protobuf:"bytes,18,opt,name=reserve,proto3" json:"reserve,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CashOutRequest1318) Reset()         { *m = CashOutRequest1318{} }
func (m *CashOutRequest1318) String() string { return proto.CompactTextString(m) }
func (*CashOutRequest1318) ProtoMessage()    {}
func (*CashOutRequest1318) Descriptor() ([]byte, []int) {
	return fileDescriptor_83707be2c8bfa1fb, []int{3}
}

func (m *CashOutRequest1318) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CashOutRequest1318.Unmarshal(m, b)
}
func (m *CashOutRequest1318) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CashOutRequest1318.Marshal(b, m, deterministic)
}
func (m *CashOutRequest1318) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CashOutRequest1318.Merge(m, src)
}
func (m *CashOutRequest1318) XXX_Size() int {
	return xxx_messageInfo_CashOutRequest1318.Size(m)
}
func (m *CashOutRequest1318) XXX_DiscardUnknown() {
	xxx_messageInfo_CashOutRequest1318.DiscardUnknown(m)
}

var xxx_messageInfo_CashOutRequest1318 proto.InternalMessageInfo

func (m *CashOutRequest1318) GetTranWebName() string {
	if m != nil {
		return m.TranWebName
	}
	return ""
}

func (m *CashOutRequest1318) GetThirdCustId() string {
	if m != nil {
		return m.ThirdCustId
	}
	return ""
}

func (m *CashOutRequest1318) GetIdType() string {
	if m != nil {
		return m.IdType
	}
	return ""
}

func (m *CashOutRequest1318) GetIdCode() string {
	if m != nil {
		return m.IdCode
	}
	return ""
}

func (m *CashOutRequest1318) GetTranOutType() string {
	if m != nil {
		return m.TranOutType
	}
	return ""
}

func (m *CashOutRequest1318) GetCustAcctId() string {
	if m != nil {
		return m.CustAcctId
	}
	return ""
}

func (m *CashOutRequest1318) GetCustName() string {
	if m != nil {
		return m.CustName
	}
	return ""
}

func (m *CashOutRequest1318) GetSupAcctId() string {
	if m != nil {
		return m.SupAcctId
	}
	return ""
}

func (m *CashOutRequest1318) GetTranType() string {
	if m != nil {
		return m.TranType
	}
	return ""
}

func (m *CashOutRequest1318) GetOutAcctId() string {
	if m != nil {
		return m.OutAcctId
	}
	return ""
}

func (m *CashOutRequest1318) GetOutAcctIdName() string {
	if m != nil {
		return m.OutAcctIdName
	}
	return ""
}

func (m *CashOutRequest1318) GetOutAcctIdBankName() string {
	if m != nil {
		return m.OutAcctIdBankName
	}
	return ""
}

func (m *CashOutRequest1318) GetOutAcctIdBankCode() string {
	if m != nil {
		return m.OutAcctIdBankCode
	}
	return ""
}

func (m *CashOutRequest1318) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *CashOutRequest1318) GetCcyCode() string {
	if m != nil {
		return m.CcyCode
	}
	return ""
}

func (m *CashOutRequest1318) GetTranAmount() string {
	if m != nil {
		return m.TranAmount
	}
	return ""
}

func (m *CashOutRequest1318) GetFeeOutCustId() string {
	if m != nil {
		return m.FeeOutCustId
	}
	return ""
}

func (m *CashOutRequest1318) GetReserve() string {
	if m != nil {
		return m.Reserve
	}
	return ""
}

type CashOutResponse1318 struct {
	FrontLogNo           string   `protobuf:"bytes,1,opt,name=frontLogNo,proto3" json:"frontLogNo,omitempty"`
	HandFee              string   `protobuf:"bytes,2,opt,name=handFee,proto3" json:"handFee,omitempty"`
	FeeOutCustId         string   `protobuf:"bytes,3,opt,name=feeOutCustId,proto3" json:"feeOutCustId,omitempty"`
	Reserve              string   `protobuf:"bytes,4,opt,name=reserve,proto3" json:"reserve,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CashOutResponse1318) Reset()         { *m = CashOutResponse1318{} }
func (m *CashOutResponse1318) String() string { return proto.CompactTextString(m) }
func (*CashOutResponse1318) ProtoMessage()    {}
func (*CashOutResponse1318) Descriptor() ([]byte, []int) {
	return fileDescriptor_83707be2c8bfa1fb, []int{4}
}

func (m *CashOutResponse1318) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CashOutResponse1318.Unmarshal(m, b)
}
func (m *CashOutResponse1318) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CashOutResponse1318.Marshal(b, m, deterministic)
}
func (m *CashOutResponse1318) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CashOutResponse1318.Merge(m, src)
}
func (m *CashOutResponse1318) XXX_Size() int {
	return xxx_messageInfo_CashOutResponse1318.Size(m)
}
func (m *CashOutResponse1318) XXX_DiscardUnknown() {
	xxx_messageInfo_CashOutResponse1318.DiscardUnknown(m)
}

var xxx_messageInfo_CashOutResponse1318 proto.InternalMessageInfo

func (m *CashOutResponse1318) GetFrontLogNo() string {
	if m != nil {
		return m.FrontLogNo
	}
	return ""
}

func (m *CashOutResponse1318) GetHandFee() string {
	if m != nil {
		return m.HandFee
	}
	return ""
}

func (m *CashOutResponse1318) GetFeeOutCustId() string {
	if m != nil {
		return m.FeeOutCustId
	}
	return ""
}

func (m *CashOutResponse1318) GetReserve() string {
	if m != nil {
		return m.Reserve
	}
	return ""
}

func init() {
	proto.RegisterEnum("proto.pingan.TranFun", TranFun_name, TranFun_value)
	proto.RegisterType((*HeaderMessage)(nil), "proto.pingan.HeaderMessage")
	proto.RegisterType((*ContractRequest1303)(nil), "proto.pingan.ContractRequest1303")
	proto.RegisterType((*ContractResponse1303)(nil), "proto.pingan.ContractResponse1303")
	proto.RegisterType((*CashOutRequest1318)(nil), "proto.pingan.CashOutRequest1318")
	proto.RegisterType((*CashOutResponse1318)(nil), "proto.pingan.CashOutResponse1318")
}

func init() { proto.RegisterFile("proto/pingan.proto", fileDescriptor_83707be2c8bfa1fb) }

var fileDescriptor_83707be2c8bfa1fb = []byte{
	// 1068 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x56, 0x4b, 0x6f, 0x23, 0x45,
	0x10, 0x26, 0x9b, 0xc7, 0xd8, 0x95, 0x64, 0xd3, 0xe9, 0x64, 0xc3, 0x6c, 0x08, 0xbb, 0x4b, 0xb4,
	0x07, 0xb4, 0x42, 0x0b, 0x51, 0x2e, 0x7b, 0xcd, 0x3a, 0xb2, 0x88, 0xb4, 0x79, 0xe0, 0x4d, 0xe0,
	0xdc, 0x99, 0x29, 0xdb, 0xa3, 0x8c, 0xbb, 0x9d, 0x9e, 0x9e, 0x8d, 0xcc, 0x99, 0xc7, 0x89, 0x33,
	0x37, 0xc4, 0x05, 0x4e, 0x70, 0x41, 0xfc, 0x07, 0xb8, 0x21, 0x21, 0xf1, 0x1f, 0x58, 0x9e, 0x12,
	0xe2, 0x80, 0x84, 0x84, 0x50, 0x57, 0xf7, 0xcc, 0x78, 0x6c, 0xc8, 0x29, 0xfe, 0xbe, 0xaf, 0xeb,
	0xd1, 0x55, 0x5f, 0xdb, 0x01, 0x3e, 0xd4, 0xca, 0xa8, 0xd7, 0x87, 0x89, 0xec, 0x09, 0xf9, 0x90,
	0x00, 0x5f, 0xa2, 0x3f, 0x0f, 0x1d, 0xb7, 0xfd, 0xdd, 0x0d, 0x58, 0x7e, 0x13, 0x45, 0x8c, 0xfa,
	0x10, 0xb3, 0x4c, 0xf4, 0x90, 0x6f, 0x42, 0xc3, 0x68, 0x21, 0xdb, 0xb9, 0x8c, 0xc2, 0x99, 0x7b,
	0x33, 0xaf, 0x36, 0x3b, 0x25, 0xb6, 0x5a, 0x86, 0xfa, 0xd9, 0xe9, 0x68, 0x88, 0xe1, 0x0d, 0xa7,
	0x15, 0x98, 0x87, 0x10, 0x0c, 0x44, 0xd4, 0x52, 0x31, 0x86, 0xb3, 0x24, 0x15, 0xb0, 0xc8, 0xb8,
	0x2f, 0x0c, 0x86, 0x73, 0x55, 0x46, 0x8b, 0x0b, 0xed, 0x34, 0x19, 0x60, 0x38, 0x5f, 0x69, 0x16,
	0xdb, 0x8c, 0x3a, 0x1b, 0x52, 0xc6, 0x05, 0x97, 0xd1, 0x43, 0xbe, 0x01, 0x0b, 0x3a, 0x1b, 0x1e,
	0x66, 0xbd, 0x30, 0x20, 0xc1, 0x23, 0x1b, 0x11, 0x29, 0xd9, 0x4e, 0x45, 0x2f, 0x6c, 0xb8, 0x08,
	0x0f, 0x6d, 0x44, 0x8a, 0xb2, 0x67, 0xfa, 0x61, 0xd3, 0x45, 0x38, 0xc4, 0xb7, 0xa0, 0x19, 0xa9,
	0x5c, 0x1a, 0xd4, 0x07, 0x71, 0x08, 0x24, 0x55, 0x04, 0xbf, 0x03, 0x60, 0xfa, 0x89, 0x8e, 0x9f,
	0xa8, 0xde, 0x91, 0x0a, 0x17, 0x49, 0x1e, 0x63, 0x38, 0x87, 0xb9, 0xcb, 0x51, 0x3c, 0x08, 0x97,
	0x48, 0xa1, 0xcf, 0xdb, 0x3f, 0xcc, 0xc2, 0x5a, 0x4b, 0x49, 0xa3, 0x45, 0x64, 0x3a, 0x78, 0x99,
	0x63, 0x66, 0x76, 0x76, 0xdf, 0xd8, 0xb5, 0x37, 0xed, 0xe6, 0x32, 0xa2, 0xe6, 0xfc, 0x5c, 0x0b,
	0x6c, 0xbb, 0xc8, 0xf2, 0xe1, 0x5e, 0x14, 0x99, 0x83, 0xd8, 0x0f, 0xb6, 0x22, 0x6c, 0x17, 0x51,
	0x9e, 0x19, 0x2f, 0xbb, 0xe1, 0x8e, 0x31, 0x36, 0xb3, 0x45, 0x47, 0x62, 0x50, 0xce, 0xb7, 0xc0,
	0xfc, 0x1e, 0x2c, 0x52, 0xbf, 0xad, 0x3c, 0xb3, 0xc1, 0x6e, 0xc4, 0xe3, 0x94, 0x9d, 0x4c, 0x12,
	0xd3, 0x46, 0xdd, 0x90, 0x3d, 0x72, 0x3c, 0x0d, 0x3f, 0x28, 0x78, 0x9a, 0xfd, 0x7d, 0x58, 0xd6,
	0x98, 0x0a, 0x83, 0xb1, 0x6f, 0xc8, 0x4d, 0xba, 0x4e, 0xda, 0x9e, 0x44, 0x14, 0x19, 0xba, 0xad,
	0x9b, 0x78, 0x89, 0xcb, 0x9d, 0xdb, 0x9a, 0x30, 0xb6, 0x73, 0x5b, 0xd5, 0xc7, 0xd1, 0x5d, 0x16,
	0xab, 0x38, 0xba, 0xcb, 0x26, 0x34, 0xce, 0x85, 0xbc, 0xa0, 0x9e, 0xdc, 0xc4, 0x4b, 0x5c, 0x68,
	0x14, 0xb7, 0x5c, 0x69, 0x14, 0xf7, 0x00, 0x98, 0x4a, 0xe3, 0x4e, 0xad, 0xe9, 0x9b, 0x74, 0x66,
	0x8a, 0x27, 0xcf, 0xa1, 0xf5, 0x34, 0x86, 0x2b, 0xde, 0x73, 0x0e, 0x6e, 0x9f, 0xc0, 0x7a, 0xb5,
	0xd6, 0x6c, 0xa8, 0x64, 0x86, 0xb4, 0xd7, 0xba, 0x47, 0x66, 0xa6, 0x3c, 0x12, 0x42, 0x60, 0x7c,
	0x46, 0xb7, 0xd9, 0x02, 0x6e, 0x7f, 0x3f, 0x07, 0xbc, 0x25, 0xb2, 0xfe, 0x71, 0x5e, 0x19, 0x65,
	0xe7, 0x11, 0xad, 0x4c, 0x0b, 0xf9, 0x0e, 0x9e, 0xd3, 0x6d, 0x66, 0xfc, 0xca, 0x2a, 0x6a, 0x72,
	0xa9, 0x37, 0xae, 0x5b, 0xea, 0xec, 0xff, 0x2c, 0x75, 0xae, 0xb6, 0x54, 0x5f, 0xf3, 0x38, 0x37,
	0x14, 0x34, 0x5f, 0xd5, 0xf4, 0xd4, 0x84, 0x09, 0x17, 0xae, 0x35, 0x61, 0x30, 0x61, 0xc2, 0x9a,
	0xbd, 0x1b, 0x93, 0xf6, 0x1e, 0xb7, 0x43, 0x73, 0xc2, 0x0e, 0x5b, 0xd0, 0x54, 0x79, 0x51, 0xd4,
	0x3f, 0xcf, 0x92, 0xb0, 0x56, 0x2c, 0xc1, 0x98, 0x63, 0xea, 0x24, 0x7f, 0x0d, 0x56, 0x4b, 0xe2,
	0x71, 0xe1, 0x11, 0xe7, 0x9f, 0x69, 0x61, 0xea, 0x34, 0x0d, 0x6b, 0xf9, 0x3f, 0x4e, 0xd3, 0xdc,
	0x42, 0x08, 0x44, 0x1c, 0x6b, 0xcc, 0x32, 0xef, 0xa8, 0x02, 0xd2, 0x57, 0x51, 0x34, 0xa2, 0x68,
	0x6f, 0x24, 0x0f, 0xc9, 0x30, 0x5a, 0xc8, 0xbd, 0x81, 0xfd, 0x9a, 0x09, 0x99, 0x37, 0x4c, 0xc9,
	0xf0, 0x6d, 0x58, 0xea, 0x22, 0x1e, 0xe7, 0xc6, 0xaf, 0x77, 0x95, 0x4e, 0xd4, 0xb8, 0x71, 0x9b,
	0xf2, 0xba, 0x4d, 0x3f, 0x9a, 0x81, 0xb5, 0xd2, 0x54, 0x85, 0x4d, 0x77, 0x1e, 0xd9, 0xaa, 0x5d,
	0xad, 0xa4, 0xa9, 0xd9, 0xb4, 0x62, 0x6c, 0xc6, 0xbe, 0x90, 0x71, 0x1b, 0x4b, 0x9b, 0x7a, 0x38,
	0xd5, 0xcf, 0xec, 0xf5, 0xfd, 0xcc, 0xd5, 0xfa, 0x79, 0xf0, 0x7c, 0x1e, 0x82, 0x53, 0xf7, 0xfb,
	0xc1, 0x17, 0x21, 0xd8, 0xc7, 0xae, 0xc8, 0x53, 0xc3, 0x5e, 0xe0, 0xcb, 0xd0, 0x28, 0xde, 0x13,
	0xfb, 0x18, 0xf8, 0x0a, 0x00, 0x4d, 0x55, 0x64, 0xfd, 0x03, 0xc9, 0x3e, 0x01, 0xce, 0x60, 0xb1,
	0x20, 0x8e, 0x73, 0xc3, 0x3e, 0x05, 0xbe, 0x04, 0x41, 0x81, 0x3e, 0x07, 0xfe, 0x22, 0xf0, 0x31,
	0xbd, 0xa5, 0x64, 0x37, 0xd1, 0x03, 0xf6, 0x19, 0xf0, 0x75, 0x58, 0x79, 0x92, 0x5c, 0xe6, 0x49,
	0x2c, 0x4c, 0xa2, 0x64, 0x3b, 0x49, 0x91, 0xfd, 0x18, 0xf0, 0x10, 0xd6, 0xc6, 0x59, 0x91, 0xa4,
	0xa4, 0x3c, 0x0f, 0xf8, 0x26, 0xdc, 0x3a, 0x93, 0xf8, 0x0c, 0x65, 0xab, 0x8f, 0xd1, 0x45, 0x07,
	0x23, 0xa5, 0x63, 0xd2, 0x7e, 0x0f, 0x6c, 0x91, 0x43, 0x1c, 0x9c, 0xa3, 0x7e, 0x2c, 0x52, 0x21,
	0x23, 0xa4, 0x23, 0xec, 0xa7, 0x80, 0xdf, 0x02, 0x76, 0xaa, 0x93, 0x5e, 0xcf, 0x2a, 0xf2, 0xc2,
	0xd1, 0x3f, 0x07, 0xfc, 0x65, 0x08, 0x27, 0xe9, 0x13, 0xad, 0x7a, 0xd6, 0x11, 0xec, 0x97, 0xa0,
	0xe8, 0xf9, 0x48, 0x99, 0xa4, 0x3b, 0x7a, 0x3b, 0xc1, 0x2b, 0xaa, 0xf3, 0x6b, 0xc0, 0x5f, 0x81,
	0x2d, 0x7b, 0x91, 0x76, 0xaa, 0xae, 0x28, 0x68, 0x4f, 0xc6, 0x4f, 0x93, 0x9e, 0x7d, 0x7c, 0x87,
	0xc2, 0x44, 0x7d, 0xf6, 0x5b, 0xc0, 0x5f, 0x82, 0x8d, 0xb7, 0x72, 0xd4, 0x23, 0x9b, 0xa0, 0xd6,
	0x13, 0xfb, 0x83, 0x6e, 0x47, 0xe2, 0xd3, 0x7c, 0x30, 0x10, 0xf6, 0x8c, 0x53, 0xfe, 0x09, 0xf8,
	0x06, 0xac, 0x92, 0x72, 0x60, 0x93, 0xed, 0x45, 0xf4, 0xdb, 0xc6, 0xfe, 0x0c, 0xf8, 0x1d, 0xb8,
	0xed, 0xd2, 0xa1, 0xb9, 0x42, 0x94, 0x2e, 0xa3, 0xaf, 0xc9, 0xfe, 0xa2, 0x72, 0xb6, 0x12, 0x9d,
	0xa9, 0x97, 0xfb, 0x3b, 0xe0, 0xf7, 0xe1, 0xee, 0x74, 0x30, 0x55, 0xd8, 0x47, 0x23, 0x92, 0x14,
	0x63, 0xf6, 0x25, 0xf0, 0xbb, 0xb0, 0xe9, 0x4f, 0x9d, 0x56, 0xa7, 0xca, 0x03, 0x5f, 0x00, 0xbf,
	0x09, 0x4d, 0x5b, 0xf1, 0x6c, 0x68, 0x6b, 0x7e, 0x05, 0xb6, 0x57, 0xdf, 0xa1, 0x4f, 0x7c, 0x22,
	0x46, 0xec, 0x6b, 0xb0, 0xbd, 0xd4, 0x79, 0x6b, 0xa8, 0xac, 0x8b, 0x9a, 0xbd, 0xd7, 0xb0, 0x23,
	0xf7, 0x62, 0x5b, 0xab, 0x77, 0x51, 0xee, 0xc9, 0xf8, 0x4c, 0xba, 0x8f, 0xec, 0xfd, 0x06, 0xbf,
	0x0d, 0xeb, 0x27, 0xa9, 0x30, 0x5d, 0xa5, 0x07, 0xad, 0xbe, 0xd0, 0x3d, 0xec, 0x60, 0x37, 0x97,
	0x31, 0xfb, 0xa0, 0x61, 0xb7, 0x51, 0x48, 0x27, 0x62, 0xd4, 0x52, 0x69, 0x8a, 0x91, 0x61, 0x1f,
	0x36, 0xca, 0xc6, 0x7d, 0xde, 0x0e, 0x46, 0x98, 0x0c, 0xfd, 0xd5, 0x62, 0xf6, 0x0d, 0xd8, 0xed,
	0xd3, 0x01, 0xaf, 0xd8, 0x67, 0xcd, 0xbe, 0x85, 0xf3, 0x05, 0xfa, 0xd7, 0x6a, 0xf7, 0xdf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x18, 0x48, 0x37, 0x25, 0x77, 0x09, 0x00, 0x00,
}
