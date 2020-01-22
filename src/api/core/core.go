package core

import "fmt"

func (cc Core) UpdateDriversLocation(location *Location) (*Location, error){
	location, err := cc.db.UpdateLocation(location)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return location, nil
}

func (cc Core) InitDbTables() error{
	err := cc.db.CreateTables()
	if err != nil {
		fmt.Println(err.Error())

		return err
	}

	return nil
}

func NewCore() (*Core, error){
	db, err := NewDBBackend()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &Core{db:db}, nil
}

type Core struct {
	db *DBBackend
}