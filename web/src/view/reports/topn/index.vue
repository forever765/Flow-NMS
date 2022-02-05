<template>
  <div className="app-container">
    <el-card :style="{'height': '80px'}">
      <filter-panel :filter-data="filterData" @filterMsg="filterMsg" />
    </el-card>
    <el-card :xs="24" :sm="18">
      <top-n-traffic-line ref="TopNTrafficLine" :traffic-line-data="trafficLineData" />
    </el-card>
    <top-n-table-panel
      :data-source="dataSource"
      @changeSize="changeSize"
      @changeNum="changeNum"
    />
  </div>
</template>

<script>
import FilterPanel from '@/view/reports/topn/filterPanel.vue'
import TopNTablePanel from '@/view/reports/topn/tablePanel.vue'
import TopNTrafficLine from '@/view/reports/topn/topnTrafficLine.vue'
import { getTopN } from '@/api/reports.js'
import { ElMessage } from 'element-plus'

export default {
  name: 'TopnIndex',
  components: { FilterPanel, TopNTablePanel, TopNTrafficLine },
  data() {
    function handleBytes(v) {
      if (v == null) {
        return 'null'
      }
      const unit = [' KB', ' MB', ' GB', ' TB', ' PB']
      let n = -1
      const s = 1024
      if (v < s) {
        return v + ' B'
      }
      do {
        v = v / s
        n++
      }
      while (Math.round(v) > s && n < unit.length)
      return Math.round(v) + unit[n]
    }

    return {
      // 流量图子组件数据
      trafficLineData: {
        Data: [],
        isLoading: true,
      },
      // 搜索栏配置
      filterData: {
        timeSelect: true,
        elinput: [
          {
            name: 'IP地址',
            width: 230,
            key: 'ipAddr'
          },
          {
            name: '类型',
            width: 230,
            key: 'class'
          },
          {
            name: 'XXX',
            width: 230,
            key: 'shareId'
          }
        ]
      },
      // 表格配置
      dataSource: {
        src: [], // 表格数据
        dst: [], // 表格数据
        cols: [
          {
            label: '主机名',
            prop: 'hostname',
          },
          {
            label: 'IP地址',
            prop: 'ipaddr'
          },
          {
            label: 'ISP信息',
            prop: 'isp'
          },
          {
            label: '数据大小',
            prop: 'mbytes',
            isCodeTableFormatter: function(val) {
              return handleBytes(val.mbytes)
            }
          },
          {
            label: '包数量',
            prop: 'packets',
            width: 100,
          },
        ], // 表格的列数据
        handleSelectionChange: () => {
        },
        isSelection: false, // 表格有多选时设置
        isOperation: false, // 表格有操作列时设置
        isIndex: true, // 列表序号
        loading: true, // loading
        pageData: {
          total: 0, // 总条数
          pageSize: 20, // 每页数量
          pageNum: 1 // 页码
        },
        operation: {
          // 表格有操作列时设置
          label: '操作', // 列名
          width: '100', // 根据实际情况给宽度
          data: [ // 功能数组
            {
              type: 'icon', // 为icon则是图标
              label: '推荐', // 功能
              icon: 'iconfont recommend-btn icon-iconkuozhan_tuijianpre',
              // permission: '3010105', // 后期这个操作的权限，用来控制权限
              handleRow: this.handleRow
            },
            {
              label: '删除', // 操作名称
              type: 'danger', // 为element btn属性则是按钮
              // permission: '2010702', // 后期这个操作的权限，用来控制权限
              handleRow: this.handleRow
            }
          ]
        }
      },
      dialogAdd: false,
      msg: {}
    }
  },
  created() {
    this.getTableData()
  },
  methods: {
    getTableData(a) {
      const data = {
        pageSize: this.dataSource.pageData.pageSize,
        pageNum: this.dataSource.pageData.pageNum
      }
      // 校验IP地址格式
      if (this.msg.ipAddr) {
        var reg = /(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}/
        if (!reg.test(this.msg.ipAddr)) {
          ElMessage.error('请输入正确的ip地址')
          return
        }
        data.ipAddr = this.msg.ipAddr
      }
      if (this.msg.startTime) {
        data.startTime = this.msg.startTime
        data.endTime = this.msg.endTime
      }
      if (this.msg.ipAddr) {
        data.ipAddr = this.msg.ipAddr
      }
      if (this.msg.class) {
        data.class = this.msg.class
      }
      data.protocolVersion = this.msg.protocolVersion
      this.dataSource.loading = true
      getTopN(data).then(res => {
        this.dataSource.loading = false
        if (res['msg'] === '获取成功') {
          if (res['data'] !== null) {
            ElMessage.success('搜索成功')
            // 少于pageSize就统计长度，否则返回10000条的total
            if (res['data'].length < this.dataSource.pageData.pageSize) {
              this.dataSource.pageData.total = res['data'].length
            } else {
              this.dataSource.pageData.total = 10000
            }
            this.dataSource.src = res['data']['src']
            this.dataSource.dst = res['data']['dst']
          } else {
            ElMessage.warning('搜索成功，无结果')
            this.dataSource.src = []
            this.dataSource.dst = []
            this.dataSource.pageData.total = 0
          }
        }
      })
    },
    filterMsg(msg) {
      this.msg = msg
      if (Object.keys(msg).length > 0) {
        this.$.refs.TopNTrafficLine.getTrafficData(msg)
        this.getTableData(msg)
      } else {
        this.getTableData()
      }
    },
    changeSize(size) {
      this.dataSource.pageData.pageSize = size
      this.getTableData()
    },
    changeNum(pageNum) {
      this.dataSource.pageData.pageNum = pageNum
      this.getTableData()
    },
  }
}
</script>

<style scoped lang='scss'>

</style>
