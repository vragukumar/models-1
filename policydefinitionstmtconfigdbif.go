package models

import (
	"database/sql"
	"fmt"
	"strings"
	"utils/dbutils"
	"reflect"
)

func (obj PolicyDefinitionStmtConfig) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS PolicyDefinitionStmtConfig " +
		"( " +
		"Name TEXT, " +
		"MatchConditions TEXT, " +
		"Conditions TEXT, " +
		"Actions TEXT, " +
		"Export INTEGER, " +
		"Import INTEGER, " +
		"PRIMARY KEY(Name) " +
	")"

	_, err := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj PolicyDefinitionStmtConfig) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf("INSERT INTO PolicyDefinitionStmtConfig (Name, MatchConditions, Conditions, Actions, Export, Import) VALUES ('%v', '%v', '%v', '%v', '%v', '%v') ;",
		obj.Name, obj.MatchConditions, obj.Conditions, obj.Actions, dbutils.ConvertBoolToInt(obj.Export), dbutils.ConvertBoolToInt(obj.Import))
	fmt.Println("**** Create Object called with ", obj)

	result, err := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("**** Failed to Create table", err)
	} else {
	    objectId, err = result.LastInsertId()
	    if err != nil {
		    fmt.Println("### Failed to return last object id", err)
	    }

	}
	return objectId, err
}

func (obj PolicyDefinitionStmtConfig) DeleteObjectFromDb(objKey string, dbHdl *sql.DB) error {
	sqlKey, err := obj.GetSqlKeyStr(objKey)
	if err != nil {
		fmt.Println("GetSqlKeyStr for PolicyDefinitionStmtConfig with key", objKey, "failed with error", err)
		return err
	}

	dbCmd := "delete from PolicyDefinitionStmtConfig where " + sqlKey
	fmt.Println("### DB Deleting PolicyDefinitionStmtConfig\n")
	_, err = dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj PolicyDefinitionStmtConfig) GetObjectFromDb(objKey string, dbHdl *sql.DB) (ConfigObj, error) {
	var object PolicyDefinitionStmtConfig
	sqlKey, err := obj.GetSqlKeyStr(objKey)
	dbCmd := "select * from PolicyDefinitionStmtConfig where " + sqlKey
	var tmp2 string
	var tmp3 string
	var tmp4 string
	var tmp5 string
	err = dbHdl.QueryRow(dbCmd).Scan(&object.Name, &object.MatchConditions, &tmp2, &tmp3, &tmp4, &tmp5, )
	fmt.Println("### DB Get PolicyDefinitionStmtConfig\n", err)
convtmpConditions := strings.Split(tmp2, ",")
                        for _, x := range convtmpConditions {
                            y := strings.Replace(x, " ", "", 1)
                     object.Conditions = append(object.Conditions, string(y))
                     }
convtmpActions := strings.Split(tmp3, ",")
                        for _, x := range convtmpActions {
                            y := strings.Replace(x, " ", "", 1)
                     object.Actions = append(object.Actions, string(y))
                     }
	object.Export = dbutils.ConvertStrBoolIntToBool(tmp4)
	object.Import = dbutils.ConvertStrBoolIntToBool(tmp5)
	return object, err
}

func (obj PolicyDefinitionStmtConfig) GetKey() (string, error) {
	key := string(obj.Name)
	return key, nil
}

func (obj PolicyDefinitionStmtConfig) GetSqlKeyStr(objKey string) (string, error) {
	keys := strings.Split(objKey, "#")
	sqlKey := "Name = "+ "\"" + keys[0] + "\""
	return sqlKey, nil
}

func (obj *PolicyDefinitionStmtConfig) GetAllObjFromDb(dbHdl *sql.DB) (objList []*PolicyDefinitionStmtConfig, e error) {
	dbCmd := "select * from PolicyDefinitionStmtConfig"
	rows, err := dbHdl.Query(dbCmd)
	if err != nil {
		fmt.Println(fmt.Sprintf("DB method Query failed for 'PolicyDefinitionStmtConfig' with error PolicyDefinitionStmtConfig", dbCmd, err))
		return objList, err
	}

	defer rows.Close()
    
	var tmp2 string
	var tmp3 string
	var tmp4 string
	var tmp5 string
	for rows.Next() {

             object := new(PolicyDefinitionStmtConfig)
             if err = rows.Scan(&object.Name, &object.MatchConditions, &object.Conditions, &object.Actions, &tmp4, &tmp5, ); err != nil {

             fmt.Println("Db method Scan failed when interating over PolicyDefinitionStmtConfig")
             }
convtmpConditions := strings.Split(tmp2, ",")
                        for _, x := range convtmpConditions {
                            y := strings.Replace(x, " ", "", 1)
                     object.Conditions = append(object.Conditions, string(y))
                     }
convtmpActions := strings.Split(tmp3, ",")
                        for _, x := range convtmpActions {
                            y := strings.Replace(x, " ", "", 1)
                     object.Actions = append(object.Actions, string(y))
                     }
	object.Export = dbutils.ConvertStrBoolIntToBool(tmp4)
	object.Import = dbutils.ConvertStrBoolIntToBool(tmp5)
	objList = append(objList, object)
    }
    return objList, nil
    }
    func (obj PolicyDefinitionStmtConfig) CompareObjectsAndDiff(updateKeys map[string]bool, dbObj ConfigObj) ([]bool, error) {
	dbV4Route := dbObj.(PolicyDefinitionStmtConfig)
	objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	dbObjVal := reflect.ValueOf(dbV4Route)
	attrIds := make([]bool, objTyp.NumField())
	idx := 0
	for i:=0; i<objTyp.NumField(); i++ {
	    fieldTyp := objTyp.Field(i)
		if fieldTyp.Anonymous {
			continue
		}

		objVal := objVal.Field(i)
		dbObjVal := dbObjVal.Field(i)
		if _, ok := updateKeys[fieldTyp.Name]; ok {
            if objVal.Kind() == reflect.Int {
                if int(objVal.Int()) != int(dbObjVal.Int()) {
                    attrIds[idx] = true
                }
            } else if objVal.Kind() == reflect.Int8 {
                if int8(objVal.Int()) != int8(dbObjVal.Int()) {
                    attrIds[idx] = true
                }
            } else if objVal.Kind() == reflect.Int16 {
                if int16(objVal.Int()) != int16(dbObjVal.Int()) {
                    attrIds[idx] = true
                }
            } else if objVal.Kind() == reflect.Int32 {
                if int32(objVal.Int()) != int32(dbObjVal.Int()) {
                    attrIds[idx] = true
                }
            } else if objVal.Kind() == reflect.Int64 {
                if int64(objVal.Int()) != int64(dbObjVal.Int()) {
                    attrIds[idx] = true
                }
            } else if objVal.Kind() == reflect.Uint {
                if uint(objVal.Uint()) != uint(dbObjVal.Uint()) {
                    attrIds[idx] = true
                }
            } else if objVal.Kind() == reflect.Uint8 {
                if uint8(objVal.Uint()) != uint8(dbObjVal.Uint()) {
                    attrIds[idx] = true
                }
            } else if objVal.Kind() == reflect.Uint16 {
                if uint16(objVal.Uint()) != uint16(dbObjVal.Uint()) {
                    attrIds[idx] = true
                }
            } else if objVal.Kind() == reflect.Uint32 {
                if uint16(objVal.Uint()) != uint16(dbObjVal.Uint()) {
                    attrIds[idx] = true
                }
            } else if objVal.Kind() == reflect.Uint64 {
                if uint16(objVal.Uint()) != uint16(dbObjVal.Uint()) {
                    attrIds[idx] = true
                }
            } else if objVal.Kind() == reflect.Bool {
                if bool(objVal.Bool()) != bool(dbObjVal.Bool()) {
                    attrIds[idx] = true
                }
            } else {
                if objVal.String() != dbObjVal.String() {
                    attrIds[idx] = true
                }
            }
            if attrIds[idx] {
				fmt.Println("attribute changed ", fieldTyp.Name)
			}
        }
		idx++

	}
	return attrIds[:idx], nil
}

    func (obj PolicyDefinitionStmtConfig) MergeDbAndConfigObj(dbObj ConfigObj, attrSet []bool) (ConfigObj, error) {
	var mergedPolicyDefinitionStmtConfig PolicyDefinitionStmtConfig
	objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	dbObjVal := reflect.ValueOf(dbObj)
	mergedObjVal := reflect.ValueOf(&mergedPolicyDefinitionStmtConfig)
	idx := 0
	for i:=0; i<objTyp.NumField(); i++ {
		if fieldTyp := objTyp.Field(i); fieldTyp.Anonymous {
			continue
		}

		objField := objVal.Field(i)
		dbObjField := dbObjVal.Field(i)
		if  attrSet[idx] {
			if dbObjField.Kind() == reflect.Int ||
			   dbObjField.Kind() == reflect.Int8 ||
			   dbObjField.Kind() == reflect.Int16 ||
			   dbObjField.Kind() == reflect.Int32 ||
			   dbObjField.Kind() == reflect.Int64 {
				mergedObjVal.Elem().Field(i).SetInt(objField.Int())
			} else if dbObjField.Kind() == reflect.Uint ||
			   dbObjField.Kind() == reflect.Uint8 ||
			   dbObjField.Kind() == reflect.Uint16 ||
			   dbObjField.Kind() == reflect.Uint32 ||
			   dbObjField.Kind() == reflect.Uint64 {
			    mergedObjVal.Elem().Field(i).SetUint(objField.Uint())
			} else if dbObjField.Kind() == reflect.Bool {
			    mergedObjVal.Elem().Field(i).SetBool(objField.Bool())
			} else {
				mergedObjVal.Elem().Field(i).SetString(objField.String())
			}
		} else {
			if dbObjField.Kind() == reflect.Int ||
			   dbObjField.Kind() == reflect.Int8 ||
			   dbObjField.Kind() == reflect.Int16 ||
			   dbObjField.Kind() == reflect.Int32 ||
			   dbObjField.Kind() == reflect.Int64 {
				mergedObjVal.Elem().Field(i).SetInt(dbObjField.Int())
			} else if dbObjField.Kind() == reflect.Uint ||
			   dbObjField.Kind() == reflect.Uint ||
			   dbObjField.Kind() == reflect.Uint8 ||
			   dbObjField.Kind() == reflect.Uint16 ||
			   dbObjField.Kind() == reflect.Uint32 {
			    mergedObjVal.Elem().Field(i).SetUint(dbObjField.Uint())
			} else if dbObjField.Kind() == reflect.Bool {
			    mergedObjVal.Elem().Field(i).SetBool(dbObjField.Bool())
			} else {
				mergedObjVal.Elem().Field(i).SetString(dbObjField.String())
			}
		}
		idx++

	}
	return mergedPolicyDefinitionStmtConfig, nil
}

    func (obj PolicyDefinitionStmtConfig) UpdateObjectInDb(dbObj ConfigObj, attrSet []bool, dbHdl *sql.DB) error {
	var fieldSqlStr string
	dbPolicyDefinitionStmtConfig := dbObj.(PolicyDefinitionStmtConfig)
	objKey, err := dbPolicyDefinitionStmtConfig.GetKey()
	objSqlKey, err := dbPolicyDefinitionStmtConfig.GetSqlKeyStr(objKey)
	dbCmd := "update " + "PolicyDefinitionStmtConfig" + " set"
objTyp := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)
	idx := 0
	for i:=0; i<objTyp.NumField(); i++ {
		if fieldTyp := objTyp.Field(i); fieldTyp.Anonymous {
			continue
		}

		if attrSet[idx] {
			fieldTyp := objTyp.Field(i)
			fieldVal := objVal.Field(i)
			if fieldVal.Kind() == reflect.Int ||
			   fieldVal.Kind() == reflect.Int8 ||
			   fieldVal.Kind() == reflect.Int16 ||
			   fieldVal.Kind() == reflect.Int32 ||
			   fieldVal.Kind() == reflect.Int64 {
				fieldSqlStr = fmt.Sprintf(" %s = '%d' ", fieldTyp.Name, int(fieldVal.Int()))
			} else if fieldVal.Kind() == reflect.Uint ||
			   fieldVal.Kind() == reflect.Uint8 ||
			   fieldVal.Kind() == reflect.Uint16 ||
			   fieldVal.Kind() == reflect.Uint32 ||
			   fieldVal.Kind() == reflect.Uint64 {
			    fieldSqlStr = fmt.Sprintf(" %s = '%d' ", fieldTyp.Name, int(fieldVal.Uint()))
			} else if objVal.Kind() == reflect.Bool {
			    fieldSqlStr = fmt.Sprintf(" %s = '%d' ", fieldTyp.Name, dbutils.ConvertBoolToInt(bool(fieldVal.Bool())))
			} else {
				fieldSqlStr = fmt.Sprintf(" %s = '%s' ", fieldTyp.Name, fieldVal.String())
			}
			dbCmd += fieldSqlStr
		}
		idx++
	}
	dbCmd += " where " + objSqlKey
	_, err = dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}