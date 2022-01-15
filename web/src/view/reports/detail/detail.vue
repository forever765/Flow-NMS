<template>
  <div>
    <el-table :data="tableData" height="250" fixed style="width: 100%">
<!--       :row-class-name="tableRowClassName"-->
<!--    <el-table height="250" fixed style="width: 100%">-->
      <el-table-column prop="timestamp_min" label="开始时间" width="180" />
      <el-table-column prop="timestamp_max" label="结束时间" width="180" />
      <el-table-column prop="srcip_port" label="源IP:端口" width="180" />
      <el-table-column prop="srcip_isp" label="源IP-ISP" />
      <el-table-column prop="srcip_loc" label="源IP-城市" />
      <el-table-column prop="dstip_port" label="目标IP:端口" />
      <el-table-column prop="dstip_isp" label="目标IP-ISP" />
      <el-table-column prop="dstip_loc" label="目标IP-城市" />
      <el-table-column prop="protocol" label="协议" />
      <el-table-column prop="byte" label="数据量" />
      <el-table-column prop="class" label="推测类型" />
      <el-table-column prop="packets" label="包数量" />
    </el-table>
  </div>
</template>

<!--<script lang="ts" setup>-->
<script>
// import { getNewData } from '../../../api/reports.js'
import { getNewestData } from '@/api/report.js'
export default {
  // setup(props) {
  //   return {
  //     async getData() {
  //       const { tableData } = await getNewData(20)
  //       console.log(tableData)
  //     },
  //   }
  // }
  created() {
    this.getData()
  },
  methods: {
    async getData() {
      var Result = await getNewestData()
      var tableData = Result['data']
      console.log(tableData)
      return tableData
    },
  },
}

// interface User {
//   date: string;
//   name: string;
//   address: string;
// }

// const tableData: User[] = [
//   {
//     date: '2016-05-03',
//     name: 'Tom',
//     address: 'No. 189, Grove St, Los Angeles',
//   },
// ];

// const tableRowClassName = ({
//   row,
//   rowIndex,
// }: {
//   // row: User;
//   rowIndex: number;
// }) => {
//   if (rowIndex === 1) {
//     return 'warning-row'
//   } else if (rowIndex === 3) {
//     return 'success-row'
//   } else if (rowIndex === 5) {
//     return 'danger-row'
//   }
//   return ''
// }

</script>

<style>
.el-table {
  font-size: 16px;
  /* :row-style="{height: '80px'}" ;
  padding: 2px; */
}
.el-table .warning-row {
  background-color: #e47777;
}
.el-table .success-row {
  background-color: #48d18d;
}
.el-table .danger-row {
  background-color: #f95959;
}
</style>
