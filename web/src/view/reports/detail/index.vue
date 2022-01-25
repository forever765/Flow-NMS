<template>
  <div className="app-container">
    <el-card>
      <filter-pane :filter-data="filterData" @filterMsg="filterMsg" />
    </el-card>
    <table-pane
      :data-source="dataSource"
      @changeSize="changeSize"
      @changeNum="changeNum"
    />
  </div>
</template>

<script>
import FilterPane from '@/view/reports/detail/filterPane.vue'
import TablePane from '@/view/reports/detail/tablePane.vue'
import { getNewestData } from '@/api/reports.js'
import { ElMessage } from 'element-plus'

export default {
  name: 'GetNewestData',
  components: { FilterPane, TablePane },
  data() {
    function handleBytes(v) {
      if (v == null) {
        return ''
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
        data: [], // 表格数据
        cols: [
          {
            label: '开始时间',
            prop: 'timestamp_min',
            // isImagePopover: true,
            width: 180
          },
          {
            label: '结束时间',
            prop: 'timestamp_max',
            width: 180
          },
          {
            label: '源IP:端口',
            prop: 'src_ip_port',
            width: 180,
          },
          {
            label: '源IP-ISP',
            prop: 'src_loc_isp'
          },
          {
            label: '目标IP:端口',
            prop: 'dst_ip_port',
            width: 180,
          },
          {
            label: '目标IP-ISP',
            prop: 'dst_loc_isp',
          },
          {
            label: '数据大小',
            prop: 'bytes',
            isCodeTableFormatter: function(val) {
              return handleBytes(val.bytes)
            }
          },
          {
            label: '类型',
            prop: 'class',
            width: 100,
          },
          {
            label: '协议类型',
            prop: 'protocol_etype'
          }
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
    this.getList()
  },
  methods: {
    getList(a) {
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
      if (this.msg.beginDate) {
        data.beginDate = this.msg.beginDate
        data.endDate = this.msg.endDate
      }
      if (this.msg.ipAddr) {
        data.ipAddr = this.msg.ipAddr
      }
      if (this.msg.class) {
        data.class = this.msg.class
      }
      data.protocolVersion = this.msg.protocolVersion
      this.dataSource.loading = true
      getNewestData(data).then(res => {
        this.dataSource.loading = false
        if (res['msg'] === '获取成功') {
          if (res['data'].length > 0) {
            this.dataSource.pageData.total = 1000
            this.dataSource.data = res['data']
          } else {
            this.dataSource.data = []
            this.dataSource.pageData.total = 0
          }
        }
      })
    },
    filterMsg(msg) {
      this.msg = msg
      if (Object.keys(msg).length > 0) {
        this.getList(msg)
      } else {
        this.getList()
      }
    },
    changeSize(size) {
      this.dataSource.pageData.pageSize = size
      this.getList()
    },
    changeNum(pageNum) {
      this.dataSource.pageData.pageNum = pageNum
      this.getList()
    },
  }
}
</script>

<style scoped lang='scss'>

</style>
