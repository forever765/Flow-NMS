<template>
  <div class="upload">
    <div class="gva-table-box">

      <div class="gva-btn-list">
        <el-upload
          class="excel-btn"
          :action="`${path}/iphost/importExcel`"
          :headers="{'x-token':token}"
          :on-success="loadExcel"
          :show-file-list="false"
        >
          <el-button size="mini" type="primary" icon="el-icon-upload2">导入</el-button>
        </el-upload>
        <el-button class="excel-btn" size="mini" type="primary" icon="el-icon-download" @click="handleExcelExport('ExcelExport.xlsx')">导出</el-button>
        <el-button class="excel-btn" size="mini" type="success" icon="el-icon-download" @click="downloadExcelTemplate()">下载模板</el-button>
      </div>
      <el-table :data="tableData">
        <el-table-column type="index" label="序号" width="55" align="center" />
        <el-table-column align="left" show-overflow-tooltip label="地区" min-width="160" prop="area" />
        <el-table-column align="left" show-overflow-tooltip label="类型" min-width="160" prop="type" />
        <el-table-column align="left" show-overflow-tooltip label="主机名" min-width="160" prop="hostname" />
        <el-table-column align="left" label="IP地址" min-width="90" prop="ipaddr" />
      </el-table>
    </div>
  </div>
</template>

<script>
import { ElMessage } from 'element-plus'

const path = import.meta.env.VITE_BASE_API
import { mapGetters } from 'vuex'
import infoList from '@/mixins/infoList'
import { exportExcel, loadExcelData, downloadTemplate } from '@/api/excel'
export default {
  name: 'Excel',
  mixins: [infoList],
  data() {
    return {
      listApi: loadExcelData,
      path: path
    }
  },
  computed: {
    ...mapGetters('user', ['userInfo', 'token'])
  },
  created() {
    this.pageSize = 999
    this.getTableData()
  },
  methods: {
    handleExcelExport(fileName) {
      if (!fileName || typeof fileName !== 'string') {
        fileName = 'ExcelExport.xlsx'
      }
      exportExcel(this.tableData, fileName)
    },
    loadExcel() {
      this.listApi = loadExcelData
      ElMessage.success('导入成功')
      this.getTableData()
    },
    downloadExcelTemplate() {
      downloadTemplate('ExcelTemplate.xlsx')
    }
  }
}
</script>

<style lang="scss" scoped>
.btn-list{
  display: flex;
  margin-bottom: 12px;
  justify-content: flex-end;

}
.excel-btn+.excel-btn{
  margin-left: 10px;
}
</style>
