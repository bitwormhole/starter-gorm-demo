// 这个文件是由 starter-configen 工具生成的配置代码，千万不要手工修改里面的任何内容。
package etc

import(
	datasource_877e7cce "github.com/bitwormhole/starter-gorm/datasource"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
)

func Config(cb application.ConfigBuilder) error {

    // _gorm_db_drivers
    cb.AddComponent(&config.ComInfo{
		ID: "_gorm_db_drivers",
		Class: "gorm-db-drivers",
		Scope: application.ScopeSingleton,
		Aliases: []string{},
		OnNew: func() lang.Object {
		    return &datasource_877e7cce.DriverRegistrar{}
		},
		OnInject: func(obj lang.Object,context application.Context) error {
		    target := obj.(*datasource_877e7cce.DriverRegistrar)
		    return _gorm_db_drivers(target,context)
		},
    })

    // _gorm_starter
    cb.AddComponent(&config.ComInfo{
		ID: "db1",
		Class: "",
		Scope: application.ScopeSingleton,
		Aliases: []string{},
		OnNew: func() lang.Object {
		    return &datasource_877e7cce.GormDataSource{}
		},
		OnInit: func(obj lang.Object) error {
		    target := obj.(*datasource_877e7cce.GormDataSource)
		    return target.Open()
		},
		OnDestroy: func(obj lang.Object) error {
		    target := obj.(*datasource_877e7cce.GormDataSource)
		    return target.Close()
		},
		OnInject: func(obj lang.Object,context application.Context) error {
		    target := obj.(*datasource_877e7cce.GormDataSource)
		    return _gorm_starter(target,context)
		},
    })

    return nil
}

