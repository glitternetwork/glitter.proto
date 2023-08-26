package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"testing"
)


func Test_MsgTypeURL(t *testing.T) {
	owner := sdk.AccAddress("abc")
	var msgSchema interface{}
	msgSchema = NewMsgSchema(owner,"abc",[]byte("anc"))
	sdkMsg,isOK := msgSchema.(sdk.Msg)
	assert.Equal(t, isOK,true)
	assert.Equal(t, sdk.MsgTypeURL(sdkMsg),"/blockved.glitterchain.index.MsgSchema")
}


