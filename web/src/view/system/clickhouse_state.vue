<template>
  <el-row>
    <el-col :span="24">
      <el-card
        class="card_item"
        :body-style="{ 'overflow-y': 'scroll' }"
      >
        <div>
          <el-table
            :data="tableData"
            height="650"
            border
            style="width: 100%"
            :row-style="{height:'20px'}"
            :cell-style="{padding:'0px', 'text-align':'left'}"
          >
            <el-table-column
              v-if="tableData"
              type="index"
              label="序号"
              width="55"
              align="center"
            />
            <el-table-column prop="name" label="名称" />
            <el-table-column prop="help" label="简介" />
            <el-table-column prop="metrics[0].value" label="当前值" width="200" />
          </el-table>
        </div>
      </el-card>
    </el-col>
  </el-row>
</template>

<script>
import { getClickhouseState } from '@/api/system'

export default {
  name: 'ClickhouseInfo',
  data() {
    return {
      timer: null,
      tableData: this.tableData
    }
  },
  created() {
    this.reload()
    this.timer = setInterval(() => {
      this.reload()
    }, 1000 * 10)
  },
  beforeDestroy() {
    clearInterval(this.timer)
    this.timer = null
  },
  methods: {
    async reload() {
      var result = await getClickhouseState()
      this.tableData = JSON.parse(result['data'])
    }
  }
}
</script>

<style>
.card_item {
  /*设置内部填充为0，几个布局元素之间没有间距*/
  padding: 0px;
  /*外部间距也是如此设置*/
  margin: 0px;
  /*统一设置高度为100%*/
  height: 100%;
}

.el-table{
  width: 100%;
  .el-table__header-wrapper table,.el-table__body-wrapper table{
    width: 100% !important;
  }
  .el-table__body, .el-table__footer, .el-table__header{
    table-layout: auto;
  }
}
</style>
