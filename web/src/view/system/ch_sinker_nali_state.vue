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
            <el-table-column prop="name" label="名称" width="300" />
            <el-table-column prop="help" label="简介" width="600" />
            <el-table-column prop="metrics[0].value" label="当前值" />
            <el-table-column prop="metrics[0].labels.task" label="所属任务" />
          </el-table>
        </div>
      </el-card>
    </el-col>
  </el-row>
</template>

<script>
import { getChSinkerNaliInfo } from '@/api/system'

export default {
  name: 'ChSinkerNaliInfo',
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
      var result = await getChSinkerNaliInfo()
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
</style>
