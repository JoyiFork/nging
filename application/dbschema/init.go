//Do not edit this file, which is automatically generated by the generator.
package dbschema

import (
	"github.com/webx-top/db/lib/factory"
)

func init(){
	factory.Fields=map[string]map[string]*factory.FieldInfo{"ftp_user_group":map[string]*factory.FieldInfo{"id":&factory.FieldInfo{Name:"id", DataType:"int", Unsigned:true, PrimaryKey:true, AutoIncrement:true, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"", GoType:"uint", GoName:"Id"}, "created":&factory.FieldInfo{Name:"created", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"创建时间", GoType:"uint", GoName:"Created"}, "updated":&factory.FieldInfo{Name:"updated", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"修改时间", GoType:"uint", GoName:"Updated"}, "disabled":&factory.FieldInfo{Name:"disabled", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"Y", "N"}, DefaultValue:"N", Comment:"是否禁用", GoType:"string", GoName:"Disabled"}, "banned":&factory.FieldInfo{Name:"banned", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"Y", "N"}, DefaultValue:"N", Comment:"是否禁止组内用户连接", GoType:"string", GoName:"Banned"}, "name":&factory.FieldInfo{Name:"name", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:255, Options:[]string{}, DefaultValue:"", Comment:"组名称", GoType:"string", GoName:"Name"}, "directory":&factory.FieldInfo{Name:"directory", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:500, Options:[]string{}, DefaultValue:"", Comment:"授权目录", GoType:"string", GoName:"Directory"}, "ip_whitelist":&factory.FieldInfo{Name:"ip_whitelist", DataType:"", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{}, DefaultValue:"", Comment:"IP白名单(一行一个)", GoType:"string", GoName:"IpWhitelist"}, "ip_blacklist":&factory.FieldInfo{Name:"ip_blacklist", DataType:"", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{}, DefaultValue:"", Comment:"IP黑名单(一行一个)", GoType:"string", GoName:"IpBlacklist"}}, "task_group":map[string]*factory.FieldInfo{"id":&factory.FieldInfo{Name:"id", DataType:"int", Unsigned:true, PrimaryKey:true, AutoIncrement:true, Min:0, Max:0, Precision:0, MaxSize:11, Options:[]string{}, DefaultValue:"", Comment:"", GoType:"uint", GoName:"Id"}, "uid":&factory.FieldInfo{Name:"uid", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:11, Options:[]string{}, DefaultValue:"0", Comment:"用户ID", GoType:"uint", GoName:"Uid"}, "name":&factory.FieldInfo{Name:"name", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:60, Options:[]string{}, DefaultValue:"", Comment:"组名", GoType:"string", GoName:"Name"}, "description":&factory.FieldInfo{Name:"description", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:255, Options:[]string{}, DefaultValue:"", Comment:"说明", GoType:"string", GoName:"Description"}, "created":&factory.FieldInfo{Name:"created", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:11, Options:[]string{}, DefaultValue:"0", Comment:"创建时间", GoType:"uint", GoName:"Created"}}, "user_u2f":map[string]*factory.FieldInfo{"type":&factory.FieldInfo{Name:"type", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:30, Options:[]string{}, DefaultValue:"", Comment:"类型", GoType:"string", GoName:"Type"}, "extra":&factory.FieldInfo{Name:"extra", DataType:"", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{}, DefaultValue:"", Comment:"扩展设置", GoType:"string", GoName:"Extra"}, "created":&factory.FieldInfo{Name:"created", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:11, Options:[]string{}, DefaultValue:"", Comment:"绑定时间", GoType:"uint", GoName:"Created"}, "id":&factory.FieldInfo{Name:"id", DataType:"bigint", Unsigned:true, PrimaryKey:true, AutoIncrement:true, Min:0, Max:0, Precision:0, MaxSize:20, Options:[]string{}, DefaultValue:"", Comment:"", GoType:"uint64", GoName:"Id"}, "uid":&factory.FieldInfo{Name:"uid", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:11, Options:[]string{}, DefaultValue:"", Comment:"用户ID", GoType:"uint", GoName:"Uid"}, "token":&factory.FieldInfo{Name:"token", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:255, Options:[]string{}, DefaultValue:"", Comment:"签名", GoType:"string", GoName:"Token"}}, "code_invitation":map[string]*factory.FieldInfo{"disabled":&factory.FieldInfo{Name:"disabled", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"Y", "N"}, DefaultValue:"N", Comment:"是否禁用", GoType:"string", GoName:"Disabled"}, "uid":&factory.FieldInfo{Name:"uid", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:11, Options:[]string{}, DefaultValue:"0", Comment:"创建者", GoType:"uint", GoName:"Uid"}, "recv_uid":&factory.FieldInfo{Name:"recv_uid", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:11, Options:[]string{}, DefaultValue:"0", Comment:"使用者", GoType:"uint", GoName:"RecvUid"}, "used":&factory.FieldInfo{Name:"used", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:11, Options:[]string{}, DefaultValue:"0", Comment:"使用时间", GoType:"uint", GoName:"Used"}, "start":&factory.FieldInfo{Name:"start", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:11, Options:[]string{}, DefaultValue:"0", Comment:"有效时间", GoType:"uint", GoName:"Start"}, "end":&factory.FieldInfo{Name:"end", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:11, Options:[]string{}, DefaultValue:"0", Comment:"失效时间", GoType:"uint", GoName:"End"}, "id":&factory.FieldInfo{Name:"id", DataType:"int", Unsigned:true, PrimaryKey:true, AutoIncrement:true, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"ID", GoType:"uint", GoName:"Id"}, "code":&factory.FieldInfo{Name:"code", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:40, Options:[]string{}, DefaultValue:"", Comment:"邀请码", GoType:"string", GoName:"Code"}, "created":&factory.FieldInfo{Name:"created", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:11, Options:[]string{}, DefaultValue:"", Comment:"创建时间", GoType:"uint", GoName:"Created"}}, "code_verification":map[string]*factory.FieldInfo{"code":&factory.FieldInfo{Name:"code", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:40, Options:[]string{}, DefaultValue:"", Comment:"验证码", GoType:"string", GoName:"Code"}, "used":&factory.FieldInfo{Name:"used", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:11, Options:[]string{}, DefaultValue:"0", Comment:"使用时间", GoType:"uint", GoName:"Used"}, "purpose":&factory.FieldInfo{Name:"purpose", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:40, Options:[]string{}, DefaultValue:"", Comment:"目的", GoType:"string", GoName:"Purpose"}, "end":&factory.FieldInfo{Name:"end", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:11, Options:[]string{}, DefaultValue:"0", Comment:"失效时间", GoType:"uint", GoName:"End"}, "id":&factory.FieldInfo{Name:"id", DataType:"int", Unsigned:true, PrimaryKey:true, AutoIncrement:true, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"ID", GoType:"uint", GoName:"Id"}, "created":&factory.FieldInfo{Name:"created", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:11, Options:[]string{}, DefaultValue:"", Comment:"创建时间", GoType:"uint", GoName:"Created"}, "uid":&factory.FieldInfo{Name:"uid", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:11, Options:[]string{}, DefaultValue:"0", Comment:"创建者", GoType:"uint", GoName:"Uid"}, "start":&factory.FieldInfo{Name:"start", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:11, Options:[]string{}, DefaultValue:"0", Comment:"有效时间", GoType:"uint", GoName:"Start"}, "disabled":&factory.FieldInfo{Name:"disabled", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"Y", "N"}, DefaultValue:"N", Comment:"是否禁用", GoType:"string", GoName:"Disabled"}}, "ftp_user":map[string]*factory.FieldInfo{"created":&factory.FieldInfo{Name:"created", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"创建时间 ", GoType:"uint", GoName:"Created"}, "id":&factory.FieldInfo{Name:"id", DataType:"int", Unsigned:true, PrimaryKey:true, AutoIncrement:true, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"", GoType:"uint", GoName:"Id"}, "username":&factory.FieldInfo{Name:"username", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:120, Options:[]string{}, DefaultValue:"", Comment:"用户名", GoType:"string", GoName:"Username"}, "ip_blacklist":&factory.FieldInfo{Name:"ip_blacklist", DataType:"", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{}, DefaultValue:"", Comment:"IP黑名单(一行一个) ", GoType:"string", GoName:"IpBlacklist"}, "ip_whitelist":&factory.FieldInfo{Name:"ip_whitelist", DataType:"", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{}, DefaultValue:"", Comment:"IP白名单(一行一个) ", GoType:"string", GoName:"IpWhitelist"}, "updated":&factory.FieldInfo{Name:"updated", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"修改时间", GoType:"uint", GoName:"Updated"}, "group_id":&factory.FieldInfo{Name:"group_id", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"用户组", GoType:"uint", GoName:"GroupId"}, "password":&factory.FieldInfo{Name:"password", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:150, Options:[]string{}, DefaultValue:"", Comment:"密码", GoType:"string", GoName:"Password"}, "banned":&factory.FieldInfo{Name:"banned", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"Y", "N"}, DefaultValue:"N", Comment:"是否禁止连接", GoType:"string", GoName:"Banned"}, "directory":&factory.FieldInfo{Name:"directory", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:500, Options:[]string{}, DefaultValue:"", Comment:"授权目录(一行一个) ", GoType:"string", GoName:"Directory"}}, "task":map[string]*factory.FieldInfo{"id":&factory.FieldInfo{Name:"id", DataType:"int", Unsigned:true, PrimaryKey:true, AutoIncrement:true, Min:0, Max:0, Precision:0, MaxSize:11, Options:[]string{}, DefaultValue:"", Comment:"", GoType:"uint", GoName:"Id"}, "description":&factory.FieldInfo{Name:"description", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:200, Options:[]string{}, DefaultValue:"", Comment:"任务描述", GoType:"string", GoName:"Description"}, "uid":&factory.FieldInfo{Name:"uid", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:11, Options:[]string{}, DefaultValue:"0", Comment:"用户ID", GoType:"uint", GoName:"Uid"}, "type":&factory.FieldInfo{Name:"type", DataType:"tinyint", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:4, Options:[]string{}, DefaultValue:"0", Comment:"任务类型", GoType:"int", GoName:"Type"}, "cron_spec":&factory.FieldInfo{Name:"cron_spec", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:100, Options:[]string{}, DefaultValue:"", Comment:"时间表达式", GoType:"string", GoName:"CronSpec"}, "concurrent":&factory.FieldInfo{Name:"concurrent", DataType:"tinyint", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:1, Options:[]string{}, DefaultValue:"0", Comment:"是否只允许一个实例", GoType:"uint", GoName:"Concurrent"}, "command":&factory.FieldInfo{Name:"command", DataType:"", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{}, DefaultValue:"", Comment:"命令详情", GoType:"string", GoName:"Command"}, "enable_notify":&factory.FieldInfo{Name:"enable_notify", DataType:"tinyint", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:1, Options:[]string{}, DefaultValue:"0", Comment:"是否启用通知", GoType:"uint", GoName:"EnableNotify"}, "notify_email":&factory.FieldInfo{Name:"notify_email", DataType:"", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{}, DefaultValue:"", Comment:"通知人列表", GoType:"string", GoName:"NotifyEmail"}, "timeout":&factory.FieldInfo{Name:"timeout", DataType:"smallint", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:6, Options:[]string{}, DefaultValue:"0", Comment:"超时设置", GoType:"uint", GoName:"Timeout"}, "execute_times":&factory.FieldInfo{Name:"execute_times", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:11, Options:[]string{}, DefaultValue:"0", Comment:"累计执行次数", GoType:"uint", GoName:"ExecuteTimes"}, "group_id":&factory.FieldInfo{Name:"group_id", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:11, Options:[]string{}, DefaultValue:"0", Comment:"分组ID", GoType:"uint", GoName:"GroupId"}, "disabled":&factory.FieldInfo{Name:"disabled", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"Y", "N"}, DefaultValue:"N", Comment:"是否禁用", GoType:"string", GoName:"Disabled"}, "name":&factory.FieldInfo{Name:"name", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:50, Options:[]string{}, DefaultValue:"", Comment:"任务名称", GoType:"string", GoName:"Name"}, "prev_time":&factory.FieldInfo{Name:"prev_time", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"上次执行时间", GoType:"uint", GoName:"PrevTime"}, "created":&factory.FieldInfo{Name:"created", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:11, Options:[]string{}, DefaultValue:"0", Comment:"创建时间", GoType:"uint", GoName:"Created"}}, "task_log":map[string]*factory.FieldInfo{"output":&factory.FieldInfo{Name:"output", DataType:"", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{}, DefaultValue:"", Comment:"任务输出", GoType:"string", GoName:"Output"}, "error":&factory.FieldInfo{Name:"error", DataType:"", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{}, DefaultValue:"", Comment:"错误信息", GoType:"string", GoName:"Error"}, "status":&factory.FieldInfo{Name:"status", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"success", "timeout", "failure", "stop", "restart"}, DefaultValue:"success", Comment:"状态", GoType:"string", GoName:"Status"}, "elapsed":&factory.FieldInfo{Name:"elapsed", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:11, Options:[]string{}, DefaultValue:"0", Comment:"消耗时间(毫秒)", GoType:"uint", GoName:"Elapsed"}, "created":&factory.FieldInfo{Name:"created", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:11, Options:[]string{}, DefaultValue:"0", Comment:"创建时间", GoType:"uint", GoName:"Created"}, "id":&factory.FieldInfo{Name:"id", DataType:"int", Unsigned:true, PrimaryKey:true, AutoIncrement:true, Min:0, Max:0, Precision:0, MaxSize:11, Options:[]string{}, DefaultValue:"", Comment:"", GoType:"uint", GoName:"Id"}, "task_id":&factory.FieldInfo{Name:"task_id", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:11, Options:[]string{}, DefaultValue:"0", Comment:"任务ID", GoType:"uint", GoName:"TaskId"}}, "user":map[string]*factory.FieldInfo{"salt":&factory.FieldInfo{Name:"salt", DataType:"char", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:64, Options:[]string{}, DefaultValue:"", Comment:"盐值", GoType:"string", GoName:"Salt"}, "last_login":&factory.FieldInfo{Name:"last_login", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:11, Options:[]string{}, DefaultValue:"0", Comment:"最后登录时间", GoType:"uint", GoName:"LastLogin"}, "last_ip":&factory.FieldInfo{Name:"last_ip", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:150, Options:[]string{}, DefaultValue:"", Comment:"最后登录IP", GoType:"string", GoName:"LastIp"}, "disabled":&factory.FieldInfo{Name:"disabled", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"Y", "N"}, DefaultValue:"N", Comment:"状态", GoType:"string", GoName:"Disabled"}, "id":&factory.FieldInfo{Name:"id", DataType:"int", Unsigned:true, PrimaryKey:true, AutoIncrement:true, Min:0, Max:0, Precision:0, MaxSize:11, Options:[]string{}, DefaultValue:"", Comment:"", GoType:"uint", GoName:"Id"}, "username":&factory.FieldInfo{Name:"username", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:30, Options:[]string{}, DefaultValue:"", Comment:"用户名", GoType:"string", GoName:"Username"}, "email":&factory.FieldInfo{Name:"email", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:50, Options:[]string{}, DefaultValue:"", Comment:"邮箱", GoType:"string", GoName:"Email"}, "password":&factory.FieldInfo{Name:"password", DataType:"char", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:64, Options:[]string{}, DefaultValue:"", Comment:"密码", GoType:"string", GoName:"Password"}}, "vhost":map[string]*factory.FieldInfo{"id":&factory.FieldInfo{Name:"id", DataType:"int", Unsigned:true, PrimaryKey:true, AutoIncrement:true, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"ID", GoType:"uint", GoName:"Id"}, "domain":&factory.FieldInfo{Name:"domain", DataType:"", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{}, DefaultValue:"", Comment:"域名", GoType:"string", GoName:"Domain"}, "root":&factory.FieldInfo{Name:"root", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:500, Options:[]string{}, DefaultValue:"", Comment:"网站物理路径", GoType:"string", GoName:"Root"}, "created":&factory.FieldInfo{Name:"created", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"创建时间", GoType:"uint", GoName:"Created"}, "updated":&factory.FieldInfo{Name:"updated", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"更新时间", GoType:"uint", GoName:"Updated"}, "setting":&factory.FieldInfo{Name:"setting", DataType:"", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{}, DefaultValue:"", Comment:"设置", GoType:"string", GoName:"Setting"}, "disabled":&factory.FieldInfo{Name:"disabled", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"Y", "N"}, DefaultValue:"N", Comment:"是否停用", GoType:"string", GoName:"Disabled"}}}

}

