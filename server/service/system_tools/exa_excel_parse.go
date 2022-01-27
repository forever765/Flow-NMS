package system_tools

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system_tools"
	"github.com/xuri/excelize/v2"
	"strconv"
)

type SystemToolsService struct {
}

// 把数据导出到excel
func (exa *SystemToolsService) ParseInfoList2Excel(infoList []system_tools.IpHost, filePath string) error {
	excel := excelize.NewFile()
	excel.SetSheetRow("Sheet1", "A1", &[]string{"ID", "所属地区", "主机名", "IP地址"})
	for i, menu := range infoList {
		axis := fmt.Sprintf("A%d", i+2)
		excel.SetSheetRow("Sheet1", axis, &[]interface{}{
			menu.ID,
			menu.Area,
			menu.HostName,
			menu.IpAddr,
		})
	}
	err := excel.SaveAs(filePath)
	return err
}

// 解析excel并返回结果
func (exa *SystemToolsService) ParseExcel2Redis() error {
	skipHeader := true
	fixedHeader := []string{"ID", "所属地区", "主机名", "IP地址"}
	file, err := excelize.OpenFile(global.GVA_CONFIG.Excel.Dir + "ExcelImport.xlsx")
	if err != nil {
		return err
	}
	menus := make([]system_tools.IpHost, 0)
	rows, err := file.Rows("Sheet1")
	if err != nil {
		return err
	}
	for rows.Next() {
		row, err := rows.Columns()
		if err != nil {
			return err
		}
		if skipHeader {
			if exa.compareStrSlice(row, fixedHeader) {
				skipHeader = false
				continue
			} else {
				return errors.New("Excel格式错误")
			}
		}
		if len(row) != len(fixedHeader) {
			continue
		}
		id, _ := strconv.Atoi(row[0])
		menu := system_tools.IpHost{
			ID: 	uint(id),
			Area:      row[1],
			HostName:  row[2],
			IpAddr:  row[3],
		}
		menus = append(menus, menu)
		if err := WriteInfo2Redis(menus); err != nil{
			return err
		}
	}
	return nil
}

// 解析excel数据后写入Redis缓存
func WriteInfo2Redis(raw []system_tools.IpHost) error {
	if result, err := json.Marshal(raw); err != nil {
		return err
	} else {
		write2RedisErr := global.GVA_REDIS.Set(context.Background(), "IpHostList", result, 0).Err()
		return write2RedisErr
	}
}

// 从Redis读取数据返回
func (exa *SystemToolsService) ParseInfoList4Redis() ([]system_tools.IpHost, error) {
	var ipHost []system_tools.IpHost
	if redisR, err := global.GVA_REDIS.Get(context.Background(), "IpHostList").Result(); err!=nil {
		return nil, err
	} else {
		if err := json.Unmarshal([]byte(redisR), &ipHost); err != nil{
			return nil, err
		}
	}
	return ipHost, nil
}

func (exa *SystemToolsService) compareStrSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if (b == nil) != (a == nil) {
		return false
	}
	for key, value := range a {
		if value != b[key] {
			return false
		}
	}
	return true
}
