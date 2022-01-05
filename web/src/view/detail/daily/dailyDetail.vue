<template>
  <div>
    <el-table :data="tableData" height="600" fixed style="width: 100%" :row-class-name="tableRowClassName">
<!--    <el-table :data="tableData" height="250" fixed style="width: 100%" :row-class-name="tableRowClassName">-->
      <el-table-column prop="date" label="日期" width="180" />
      <el-table-column prop="name" label="名称" width="180" />
      <el-table-column prop="address" label="地址" />
    </el-table>
  </div>
</template>

<script lang="ts" setup>
interface User {
  date: string,
  name: string,
  address: string
}

const tableRowClassName = ({
  row,
  rowIndex,
}: {
  row: User,
  rowIndex: number
}) => {
  if (rowIndex === 1) {
    return 'warning-row'
  } else if (rowIndex === 3) {
    return 'success-row'
  } else if (rowIndex === 5) {
    return 'danger-row'
  }
  return ''
}

const tableData: User[] = [
  {
    date: '2016-05-03',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
  {
    date: '2016-05-02',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
  {
    date: '2016-05-04',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
  {
    date: '2016-05-01',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
  {
    date: '2016-05-01',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
  {
    date: '2016-05-01',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
  {
    date: '2016-05-01',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
  {
    date: '2016-05-01',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
  {
    date: '2016-05-01',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
  {
    date: '2016-05-01',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
  {
    date: '2016-05-01',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
  {
    date: '2016-05-01',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
  {
    date: '2016-05-01',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  }
]

</script>

<style>
  .el-table .warning-row {
    background-color: #e47777;
  }
  .el-table .success-row {
    background-color: #48D18D;
  }
  .el-table .danger-row {
    background-color: #F95959;
  }
</style>
