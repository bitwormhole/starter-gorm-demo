package etc

import (
	"github.com/bitwormhole/starter-gorm-mysql/mysql"
	"github.com/bitwormhole/starter-gorm-sqlserver/sqlserver"
	"github.com/bitwormhole/starter-gorm/datasource"
	"github.com/bitwormhole/starter/application"
)

func _gorm_starter(inst *datasource.GormDataSource, ctx application.Context) error {

	//[component]
	//  id=db1
	//  initMethod=Open
	//  destroyMethod=Close

	err := inst.Config(&datasource.GormStarterConfig{
		DataSourceName: "demo",
		Context:        ctx,
	})
	return err
}

func _gorm_db_drivers(inst *datasource.DriverRegistrar, ctx application.Context) error {

	//[component]
	//  class= gorm-db-drivers

	inst.Register("mysql", &mysql.Driver{})
	inst.Register("sqlserver", &sqlserver.Driver{})

	return nil
}
