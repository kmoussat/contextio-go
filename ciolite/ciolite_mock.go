// Automatically generated by MockGen. DO NOT EDIT!
// Source: ciolite.go

package ciolite

import (
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	httptest "net/http/httptest"
)

// Mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *_MockInterfaceRecorder
}

// Recorder for MockInterface (not exported)
type _MockInterfaceRecorder struct {
	mock *MockInterface
}

func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &_MockInterfaceRecorder{mock}
	return mock
}

func (_m *MockInterface) EXPECT() *_MockInterfaceRecorder {
	return _m.recorder
}

func (_m *MockInterface) NewTestCioLiteServer(handler http.Handler) (CioLite, *httptest.Server) {
	ret := _m.ctrl.Call(_m, "NewTestCioLiteServer", handler)
	ret0, _ := ret[0].(CioLite)
	ret1, _ := ret[1].(*httptest.Server)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) NewTestCioLiteServer(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NewTestCioLiteServer", arg0)
}

func (_m *MockInterface) ValidateCallback(token string, signature string, timestamp int) bool {
	ret := _m.ctrl.Call(_m, "ValidateCallback", token, signature, timestamp)
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockInterfaceRecorder) ValidateCallback(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ValidateCallback", arg0, arg1, arg2)
}

func (_m *MockInterface) GetStatusCallbackURL() (GetStatusCallbackURLResponse, error) {
	ret := _m.ctrl.Call(_m, "GetStatusCallbackURL")
	ret0, _ := ret[0].(GetStatusCallbackURLResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) GetStatusCallbackURL() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetStatusCallbackURL")
}

func (_m *MockInterface) CreateStatusCallbackURL(formValues CreateStatusCallbackURLParams) (CreateDeleteStatusCallbackURLResponse, error) {
	ret := _m.ctrl.Call(_m, "CreateStatusCallbackURL", formValues)
	ret0, _ := ret[0].(CreateDeleteStatusCallbackURLResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) CreateStatusCallbackURL(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateStatusCallbackURL", arg0)
}

func (_m *MockInterface) DeleteStatusCallbackURL() (CreateDeleteStatusCallbackURLResponse, error) {
	ret := _m.ctrl.Call(_m, "DeleteStatusCallbackURL")
	ret0, _ := ret[0].(CreateDeleteStatusCallbackURLResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) DeleteStatusCallbackURL() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteStatusCallbackURL")
}

func (_m *MockInterface) GetConnectTokens() ([]GetConnectTokenResponse, error) {
	ret := _m.ctrl.Call(_m, "GetConnectTokens")
	ret0, _ := ret[0].([]GetConnectTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) GetConnectTokens() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetConnectTokens")
}

func (_m *MockInterface) GetConnectToken(token string) (GetConnectTokenResponse, error) {
	ret := _m.ctrl.Call(_m, "GetConnectToken", token)
	ret0, _ := ret[0].(GetConnectTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) GetConnectToken(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetConnectToken", arg0)
}

func (_m *MockInterface) CreateConnectToken(formValues CreateConnectTokenParams) (CreateConnectTokenResponse, error) {
	ret := _m.ctrl.Call(_m, "CreateConnectToken", formValues)
	ret0, _ := ret[0].(CreateConnectTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) CreateConnectToken(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateConnectToken", arg0)
}

func (_m *MockInterface) DeleteConnectToken(token string) (DeleteConnectTokenResponse, error) {
	ret := _m.ctrl.Call(_m, "DeleteConnectToken", token)
	ret0, _ := ret[0].(DeleteConnectTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) DeleteConnectToken(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteConnectToken", arg0)
}

func (_m *MockInterface) CheckConnectToken(connectToken GetConnectTokenResponse, email string) error {
	ret := _m.ctrl.Call(_m, "CheckConnectToken", connectToken, email)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockInterfaceRecorder) CheckConnectToken(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CheckConnectToken", arg0, arg1)
}

func (_m *MockInterface) GetDiscovery(queryValues GetDiscoveryParams) (GetDiscoveryResponse, error) {
	ret := _m.ctrl.Call(_m, "GetDiscovery", queryValues)
	ret0, _ := ret[0].(GetDiscoveryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) GetDiscovery(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetDiscovery", arg0)
}

func (_m *MockInterface) GetOAuthProviders() ([]GetOAuthProvidersResponse, error) {
	ret := _m.ctrl.Call(_m, "GetOAuthProviders")
	ret0, _ := ret[0].([]GetOAuthProvidersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) GetOAuthProviders() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetOAuthProviders")
}

func (_m *MockInterface) GetOAuthProvider(key string) (GetOAuthProvidersResponse, error) {
	ret := _m.ctrl.Call(_m, "GetOAuthProvider", key)
	ret0, _ := ret[0].(GetOAuthProvidersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) GetOAuthProvider(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetOAuthProvider", arg0)
}

func (_m *MockInterface) CreateOAuthProvider(formValues CreateOAuthProviderParams) (CreateOAuthProviderResponse, error) {
	ret := _m.ctrl.Call(_m, "CreateOAuthProvider", formValues)
	ret0, _ := ret[0].(CreateOAuthProviderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) CreateOAuthProvider(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateOAuthProvider", arg0)
}

func (_m *MockInterface) DeleteOAuthProvider(key string) (DeleteOAuthProviderResponse, error) {
	ret := _m.ctrl.Call(_m, "DeleteOAuthProvider", key)
	ret0, _ := ret[0].(DeleteOAuthProviderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) DeleteOAuthProvider(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteOAuthProvider", arg0)
}

func (_m *MockInterface) GetUserConnectTokens(userID string) ([]GetConnectTokenResponse, error) {
	ret := _m.ctrl.Call(_m, "GetUserConnectTokens", userID)
	ret0, _ := ret[0].([]GetConnectTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) GetUserConnectTokens(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUserConnectTokens", arg0)
}

func (_m *MockInterface) GetUserConnectToken(userID string, token string) (GetConnectTokenResponse, error) {
	ret := _m.ctrl.Call(_m, "GetUserConnectToken", userID, token)
	ret0, _ := ret[0].(GetConnectTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) GetUserConnectToken(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUserConnectToken", arg0, arg1)
}

func (_m *MockInterface) CreateUserConnectToken(userID string, formValues CreateConnectTokenParams) (CreateConnectTokenResponse, error) {
	ret := _m.ctrl.Call(_m, "CreateUserConnectToken", userID, formValues)
	ret0, _ := ret[0].(CreateConnectTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) CreateUserConnectToken(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateUserConnectToken", arg0, arg1)
}

func (_m *MockInterface) DeleteUserConnectToken(userID string, token string) (DeleteConnectTokenResponse, error) {
	ret := _m.ctrl.Call(_m, "DeleteUserConnectToken", userID, token)
	ret0, _ := ret[0].(DeleteConnectTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) DeleteUserConnectToken(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteUserConnectToken", arg0, arg1)
}

func (_m *MockInterface) GetUserEmailAccountConnectTokens(userID string, label string) ([]GetConnectTokenResponse, error) {
	ret := _m.ctrl.Call(_m, "GetUserEmailAccountConnectTokens", userID, label)
	ret0, _ := ret[0].([]GetConnectTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) GetUserEmailAccountConnectTokens(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUserEmailAccountConnectTokens", arg0, arg1)
}

func (_m *MockInterface) GetUserEmailAccountConnectToken(userID string, label string, token string) (GetConnectTokenResponse, error) {
	ret := _m.ctrl.Call(_m, "GetUserEmailAccountConnectToken", userID, label, token)
	ret0, _ := ret[0].(GetConnectTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) GetUserEmailAccountConnectToken(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUserEmailAccountConnectToken", arg0, arg1, arg2)
}

func (_m *MockInterface) CreateUserEmailAccountConnectToken(userID string, label string, formValues CreateConnectTokenParams) (CreateConnectTokenResponse, error) {
	ret := _m.ctrl.Call(_m, "CreateUserEmailAccountConnectToken", userID, label, formValues)
	ret0, _ := ret[0].(CreateConnectTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) CreateUserEmailAccountConnectToken(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateUserEmailAccountConnectToken", arg0, arg1, arg2)
}

func (_m *MockInterface) DeleteUserEmailAccountConnectToken(userID string, label string, token string) (DeleteConnectTokenResponse, error) {
	ret := _m.ctrl.Call(_m, "DeleteUserEmailAccountConnectToken", userID, label, token)
	ret0, _ := ret[0].(DeleteConnectTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) DeleteUserEmailAccountConnectToken(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteUserEmailAccountConnectToken", arg0, arg1, arg2)
}

func (_m *MockInterface) GetUserEmailAccountsFolderMessageAttachments(userID string, label string, folder string, messageID string, queryValues EmailAccountFolderDelimiterParam) ([]GetUserEmailAccountsFolderMessageAttachmentsResponse, error) {
	ret := _m.ctrl.Call(_m, "GetUserEmailAccountsFolderMessageAttachments", userID, label, folder, messageID, queryValues)
	ret0, _ := ret[0].([]GetUserEmailAccountsFolderMessageAttachmentsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) GetUserEmailAccountsFolderMessageAttachments(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUserEmailAccountsFolderMessageAttachments", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockInterface) GetUserEmailAccountsFolderMessageAttachment(userID string, label string, folder string, messageID string, attachmentID string, queryValues EmailAccountFolderDelimiterParam) (GetUserEmailAccountsFolderMessageAttachmentsResponse, error) {
	ret := _m.ctrl.Call(_m, "GetUserEmailAccountsFolderMessageAttachment", userID, label, folder, messageID, attachmentID, queryValues)
	ret0, _ := ret[0].(GetUserEmailAccountsFolderMessageAttachmentsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) GetUserEmailAccountsFolderMessageAttachment(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUserEmailAccountsFolderMessageAttachment", arg0, arg1, arg2, arg3, arg4, arg5)
}

func (_m *MockInterface) GetUserEmailAccountsFolderMessageBody(userID string, label string, folder string, messageID string, queryValues GetUserEmailAccountsFolderMessageBodyParams) ([]GetUserEmailAccountsFolderMessageBodyResponse, error) {
	ret := _m.ctrl.Call(_m, "GetUserEmailAccountsFolderMessageBody", userID, label, folder, messageID, queryValues)
	ret0, _ := ret[0].([]GetUserEmailAccountsFolderMessageBodyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) GetUserEmailAccountsFolderMessageBody(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUserEmailAccountsFolderMessageBody", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockInterface) GetUserEmailAccountsFolderMessageFlags(userID string, label string, folder string, messageID string, queryValues EmailAccountFolderDelimiterParam) (GetUserEmailAccountsFolderMessageFlagsResponse, error) {
	ret := _m.ctrl.Call(_m, "GetUserEmailAccountsFolderMessageFlags", userID, label, folder, messageID, queryValues)
	ret0, _ := ret[0].(GetUserEmailAccountsFolderMessageFlagsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) GetUserEmailAccountsFolderMessageFlags(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUserEmailAccountsFolderMessageFlags", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockInterface) GetUserEmailAccountsFolderMessageHeaders(userID string, label string, folder string, messageID string, queryValues GetUserEmailAccountsFolderMessageHeadersParams) (GetUserEmailAccountsFolderMessageHeadersResponse, error) {
	ret := _m.ctrl.Call(_m, "GetUserEmailAccountsFolderMessageHeaders", userID, label, folder, messageID, queryValues)
	ret0, _ := ret[0].(GetUserEmailAccountsFolderMessageHeadersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) GetUserEmailAccountsFolderMessageHeaders(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUserEmailAccountsFolderMessageHeaders", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockInterface) GetUserEmailAccountsFolderMessageRaw(userID string, label string, folder string, messageID string, queryValues EmailAccountFolderDelimiterParam) (GetUserEmailAccountsFolderMessageRawResponse, error) {
	ret := _m.ctrl.Call(_m, "GetUserEmailAccountsFolderMessageRaw", userID, label, folder, messageID, queryValues)
	ret0, _ := ret[0].(GetUserEmailAccountsFolderMessageRawResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) GetUserEmailAccountsFolderMessageRaw(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUserEmailAccountsFolderMessageRaw", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockInterface) MarkUserEmailAccountsFolderMessageRead(userID string, label string, folder string, messageID string, formValues EmailAccountFolderDelimiterParam) (UserEmailAccountsFolderMessageReadResponse, error) {
	ret := _m.ctrl.Call(_m, "MarkUserEmailAccountsFolderMessageRead", userID, label, folder, messageID, formValues)
	ret0, _ := ret[0].(UserEmailAccountsFolderMessageReadResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) MarkUserEmailAccountsFolderMessageRead(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MarkUserEmailAccountsFolderMessageRead", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockInterface) MarkUserEmailAccountsFolderMessageUnRead(userID string, label string, folder string, messageID string, formValues EmailAccountFolderDelimiterParam) (UserEmailAccountsFolderMessageReadResponse, error) {
	ret := _m.ctrl.Call(_m, "MarkUserEmailAccountsFolderMessageUnRead", userID, label, folder, messageID, formValues)
	ret0, _ := ret[0].(UserEmailAccountsFolderMessageReadResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) MarkUserEmailAccountsFolderMessageUnRead(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MarkUserEmailAccountsFolderMessageUnRead", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockInterface) GetUserEmailAccountsFolderMessages(userID string, label string, folder string, queryValues GetUserEmailAccountsFolderMessageParams) ([]GetUsersEmailAccountFolderMessagesResponse, error) {
	ret := _m.ctrl.Call(_m, "GetUserEmailAccountsFolderMessages", userID, label, folder, queryValues)
	ret0, _ := ret[0].([]GetUsersEmailAccountFolderMessagesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) GetUserEmailAccountsFolderMessages(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUserEmailAccountsFolderMessages", arg0, arg1, arg2, arg3)
}

func (_m *MockInterface) GetUserEmailAccountFolderMessage(userID string, label string, folder string, messageID string, queryValues GetUserEmailAccountsFolderMessageParams) (GetUsersEmailAccountFolderMessagesResponse, error) {
	ret := _m.ctrl.Call(_m, "GetUserEmailAccountFolderMessage", userID, label, folder, messageID, queryValues)
	ret0, _ := ret[0].(GetUsersEmailAccountFolderMessagesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) GetUserEmailAccountFolderMessage(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUserEmailAccountFolderMessage", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockInterface) MoveUserEmailAccountFolderMessage(userID string, label string, folder string, messageID string, queryValues MoveUserEmailAccountFolderMessageParams) (MoveUserEmailAccountFolderMessageResponse, error) {
	ret := _m.ctrl.Call(_m, "MoveUserEmailAccountFolderMessage", userID, label, folder, messageID, queryValues)
	ret0, _ := ret[0].(MoveUserEmailAccountFolderMessageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) MoveUserEmailAccountFolderMessage(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MoveUserEmailAccountFolderMessage", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockInterface) GetUserEmailAccountsFolders(userID string, label string, queryValues GetUserEmailAccountsFoldersParams) ([]GetUsersEmailAccountFoldersResponse, error) {
	ret := _m.ctrl.Call(_m, "GetUserEmailAccountsFolders", userID, label, queryValues)
	ret0, _ := ret[0].([]GetUsersEmailAccountFoldersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) GetUserEmailAccountsFolders(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUserEmailAccountsFolders", arg0, arg1, arg2)
}

func (_m *MockInterface) GetUserEmailAccountFolder(userID string, label string, folder string, queryValues EmailAccountFolderDelimiterParam) (GetUsersEmailAccountFoldersResponse, error) {
	ret := _m.ctrl.Call(_m, "GetUserEmailAccountFolder", userID, label, folder, queryValues)
	ret0, _ := ret[0].(GetUsersEmailAccountFoldersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) GetUserEmailAccountFolder(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUserEmailAccountFolder", arg0, arg1, arg2, arg3)
}

func (_m *MockInterface) CreateUserEmailAccountFolder(userID string, label string, folder string, formValues EmailAccountFolderDelimiterParam) (CreateEmailAccountFolderResponse, error) {
	ret := _m.ctrl.Call(_m, "CreateUserEmailAccountFolder", userID, label, folder, formValues)
	ret0, _ := ret[0].(CreateEmailAccountFolderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) CreateUserEmailAccountFolder(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateUserEmailAccountFolder", arg0, arg1, arg2, arg3)
}

func (_m *MockInterface) SafeCreateUserEmailAccountFolder(userID string, label string, folder string, formValues EmailAccountFolderDelimiterParam) (bool, error) {
	ret := _m.ctrl.Call(_m, "SafeCreateUserEmailAccountFolder", userID, label, folder, formValues)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) SafeCreateUserEmailAccountFolder(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SafeCreateUserEmailAccountFolder", arg0, arg1, arg2, arg3)
}

func (_m *MockInterface) GetUserEmailAccounts(userID string, queryValues GetUserEmailAccountsParams) ([]GetUsersEmailAccountsResponse, error) {
	ret := _m.ctrl.Call(_m, "GetUserEmailAccounts", userID, queryValues)
	ret0, _ := ret[0].([]GetUsersEmailAccountsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) GetUserEmailAccounts(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUserEmailAccounts", arg0, arg1)
}

func (_m *MockInterface) GetUserEmailAccount(userID string, label string) (GetUsersEmailAccountsResponse, error) {
	ret := _m.ctrl.Call(_m, "GetUserEmailAccount", userID, label)
	ret0, _ := ret[0].(GetUsersEmailAccountsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) GetUserEmailAccount(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUserEmailAccount", arg0, arg1)
}

func (_m *MockInterface) CreateUserEmailAccount(userID string, formValues CreateUserParams) (CreateEmailAccountResponse, error) {
	ret := _m.ctrl.Call(_m, "CreateUserEmailAccount", userID, formValues)
	ret0, _ := ret[0].(CreateEmailAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) CreateUserEmailAccount(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateUserEmailAccount", arg0, arg1)
}

func (_m *MockInterface) ModifyUserEmailAccount(userID string, label string, formValues ModifyUserEmailAccountParams) (ModifyEmailAccountResponse, error) {
	ret := _m.ctrl.Call(_m, "ModifyUserEmailAccount", userID, label, formValues)
	ret0, _ := ret[0].(ModifyEmailAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) ModifyUserEmailAccount(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ModifyUserEmailAccount", arg0, arg1, arg2)
}

func (_m *MockInterface) DeleteUserEmailAccount(userID string, label string) (DeleteEmailAccountResponse, error) {
	ret := _m.ctrl.Call(_m, "DeleteUserEmailAccount", userID, label)
	ret0, _ := ret[0].(DeleteEmailAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) DeleteUserEmailAccount(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteUserEmailAccount", arg0, arg1)
}

func (_m *MockInterface) GetUserWebhooks(userID string) ([]GetUsersWebhooksResponse, error) {
	ret := _m.ctrl.Call(_m, "GetUserWebhooks", userID)
	ret0, _ := ret[0].([]GetUsersWebhooksResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) GetUserWebhooks(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUserWebhooks", arg0)
}

func (_m *MockInterface) GetUserWebhook(userID string, webhookID string) (GetUsersWebhooksResponse, error) {
	ret := _m.ctrl.Call(_m, "GetUserWebhook", userID, webhookID)
	ret0, _ := ret[0].(GetUsersWebhooksResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) GetUserWebhook(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUserWebhook", arg0, arg1)
}

func (_m *MockInterface) CreateUserWebhook(userID string, formValues CreateUserWebhookParams) (CreateUserWebhookResponse, error) {
	ret := _m.ctrl.Call(_m, "CreateUserWebhook", userID, formValues)
	ret0, _ := ret[0].(CreateUserWebhookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) CreateUserWebhook(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateUserWebhook", arg0, arg1)
}

func (_m *MockInterface) ModifyUserWebhook(userID string, webhookID string, formValues ModifyUserWebhookParams) (ModifyWebhookResponse, error) {
	ret := _m.ctrl.Call(_m, "ModifyUserWebhook", userID, webhookID, formValues)
	ret0, _ := ret[0].(ModifyWebhookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) ModifyUserWebhook(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ModifyUserWebhook", arg0, arg1, arg2)
}

func (_m *MockInterface) DeleteUserWebhookAccount(userID string, webhookID string) (DeleteWebhookResponse, error) {
	ret := _m.ctrl.Call(_m, "DeleteUserWebhookAccount", userID, webhookID)
	ret0, _ := ret[0].(DeleteWebhookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) DeleteUserWebhookAccount(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteUserWebhookAccount", arg0, arg1)
}

func (_m *MockInterface) GetUsers(queryValues GetUsersParams) ([]GetUsersResponse, error) {
	ret := _m.ctrl.Call(_m, "GetUsers", queryValues)
	ret0, _ := ret[0].([]GetUsersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) GetUsers(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUsers", arg0)
}

func (_m *MockInterface) GetUser(userID string) (GetUsersResponse, error) {
	ret := _m.ctrl.Call(_m, "GetUser", userID)
	ret0, _ := ret[0].(GetUsersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) GetUser(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUser", arg0)
}

func (_m *MockInterface) CreateUser(formValues CreateUserParams) (CreateUserResponse, error) {
	ret := _m.ctrl.Call(_m, "CreateUser", formValues)
	ret0, _ := ret[0].(CreateUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateUser", arg0)
}

func (_m *MockInterface) ModifyUser(userID string, formValues ModifyUserParams) (ModifyUserResponse, error) {
	ret := _m.ctrl.Call(_m, "ModifyUser", userID, formValues)
	ret0, _ := ret[0].(ModifyUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) ModifyUser(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ModifyUser", arg0, arg1)
}

func (_m *MockInterface) DeleteUser(userID string) (DeleteUserResponse, error) {
	ret := _m.ctrl.Call(_m, "DeleteUser", userID)
	ret0, _ := ret[0].(DeleteUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) DeleteUser(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteUser", arg0)
}

func (_m *MockInterface) GetWebhooks() ([]GetUsersWebhooksResponse, error) {
	ret := _m.ctrl.Call(_m, "GetWebhooks")
	ret0, _ := ret[0].([]GetUsersWebhooksResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) GetWebhooks() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetWebhooks")
}

func (_m *MockInterface) GetWebhook(webhookID string) (GetUsersWebhooksResponse, error) {
	ret := _m.ctrl.Call(_m, "GetWebhook", webhookID)
	ret0, _ := ret[0].(GetUsersWebhooksResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) GetWebhook(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetWebhook", arg0)
}

func (_m *MockInterface) CreateWebhook(formValues CreateUserWebhookParams) (CreateUserWebhookResponse, error) {
	ret := _m.ctrl.Call(_m, "CreateWebhook", formValues)
	ret0, _ := ret[0].(CreateUserWebhookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) CreateWebhook(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateWebhook", arg0)
}

func (_m *MockInterface) ModifyWebhook(webhookID string, formValues ModifyUserWebhookParams) (ModifyWebhookResponse, error) {
	ret := _m.ctrl.Call(_m, "ModifyWebhook", webhookID, formValues)
	ret0, _ := ret[0].(ModifyWebhookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) ModifyWebhook(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ModifyWebhook", arg0, arg1)
}

func (_m *MockInterface) DeleteWebhookAccount(webhookID string) (DeleteWebhookResponse, error) {
	ret := _m.ctrl.Call(_m, "DeleteWebhookAccount", webhookID)
	ret0, _ := ret[0].(DeleteWebhookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockInterfaceRecorder) DeleteWebhookAccount(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteWebhookAccount", arg0)
}
