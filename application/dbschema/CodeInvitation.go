//Do not edit this file, which is automatically generated by the generator.
package dbschema

import (
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	
	"time"
)

type CodeInvitation struct {
	param   *factory.Param
	trans	*factory.Transaction
	objects []*CodeInvitation
	
	Id      	uint    	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Code    	string  	`db:"code" bson:"code" comment:"邀请码" json:"code" xml:"code"`
	Created 	uint    	`db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Uid     	uint    	`db:"uid" bson:"uid" comment:"创建者" json:"uid" xml:"uid"`
	RecvUid 	uint    	`db:"recv_uid" bson:"recv_uid" comment:"使用者" json:"recv_uid" xml:"recv_uid"`
	Used    	uint    	`db:"used" bson:"used" comment:"使用时间" json:"used" xml:"used"`
	Start   	uint    	`db:"start" bson:"start" comment:"有效时间" json:"start" xml:"start"`
	End     	uint    	`db:"end" bson:"end" comment:"失效时间" json:"end" xml:"end"`
	Disabled	string  	`db:"disabled" bson:"disabled" comment:"是否禁用" json:"disabled" xml:"disabled"`
}

func (this *CodeInvitation) Trans() *factory.Transaction {
	return this.trans
}

func (this *CodeInvitation) Use(trans *factory.Transaction) factory.Model {
	this.trans = trans
	return this
}

func (this *CodeInvitation) Objects() []*CodeInvitation {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *CodeInvitation) NewObjects() *[]*CodeInvitation {
	this.objects = []*CodeInvitation{}
	return &this.objects
}

func (this *CodeInvitation) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetTrans(this.trans).SetCollection("code_invitation").SetModel(this)
}

func (this *CodeInvitation) SetParam(param *factory.Param) factory.Model {
	this.param = param
	return this
}

func (this *CodeInvitation) Param() *factory.Param {
	if this.param == nil {
		return this.NewParam()
	}
	return this.param
}

func (this *CodeInvitation) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	return this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
}

func (this *CodeInvitation) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *CodeInvitation) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.NewObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *CodeInvitation) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
	pk, err = this.Param().SetSend(this).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint(v)
		}
	}
	return
}

func (this *CodeInvitation) Edit(mw func(db.Result) db.Result, args ...interface{}) error {
	
	return this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Update()
}

func (this *CodeInvitation) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func(){
		
	},func(){
		this.Created = uint(time.Now().Unix())
	this.Id = 0
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint(v)
		}
	}
	return 
}

func (this *CodeInvitation) Delete(mw func(db.Result) db.Result, args ...interface{}) error {
	
	return this.Param().SetArgs(args...).SetMiddleware(mw).Delete()
}

func (this *CodeInvitation) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}
