<template>
  <div>
    <div class="filter-container">
      <el-date-picker
        v-if="filterData.timeSelect"
        v-model="dateRange"
        style="width: 360px; margin:0 auto"
        type="datetimerange"
        range-separator="~"
        format="YYYY/MM/DD HH:mm:ss"
        value-format="YYYY-MM-DD HH:mm:ss"
        start-placeholder="开始时间"
        end-placeholder="结束时间"
        :shortcuts="shortcuts"
        class="filter-item"
      />
      <template v-if="filterData.elinput" :style="{'margin-buttom: 10px'}">
        <el-Radio-group v-model="protocolVersion" class="filter-item" size="small">
          <el-Radio-button label="双栈" :checked="true" name="type"></el-Radio-button>
          <el-Radio-button label="仅IPv4" name="type"></el-Radio-button>
          <el-Radio-button label="仅IPv6" name="type"></el-Radio-button>
        </el-Radio-group>
        <el-input
          v-for="(item,index) in filterData.elinput"
          :key="index"
          v-model="listQuery[item.key]"
          :placeholder="item.name"
          :style="{'width': '190px'}"
          class="filter-item"
        />
      </template>
      <template v-if="filterData.elselect">
        <el-select
          v-for="(item,index) in filterData.elselect"
          :key="index"
          v-model="listQuery[item.key]"
          :placeholder="item.name"
          clearable
          :style="{'width':item.width?item.width+'px':'50px'}"
          class="filter-item"
        >
          <el-option
            v-for="i in item.option"
            :key="i.key"
            :label="i.value"
            :value="i.key"
          />
        </el-select>
      </template>
      <div class="btn">
        <el-button class="filter-item" type="primary" @click="handleSearch">
          搜索
        </el-button>
        <el-button class="filter-item" type="warning" @click="handleReset">
          重置
        </el-button>
      </div>
    </div>
  </div>
</template>

<script>
// 搜索栏组件
// filterData:{
//   timeSelect:true,
//   elinput:[
//     {
//       name:'姓名',
//       key:'userName'
//     }
//   ],
//   elselect:[
//     {
//       name:'部门',
//       key:'department'
//       option:[{
//         key:1,
//         value:'技术部'
//       }]
//     }
//   ]
// }
import moment from 'moment'
import { ElMessage } from 'element-plus'

export default {
  name: 'FilterPane',
  props: {
    // eslint-disable-next-line vue/require-default-prop
    filterData: {
      type: Object
    }
  },
  data() {
    return {
      protocolVersion: '双栈',
      shortcuts: [
        {
          text: '1小时内',
          value: () => {
            const end = new Date()
            const start = new Date()
            start.setTime(start.getTime() - 3600 * 1000)
            this.dateRange = [start, end]
            return [start, end]
          },
        }, {
          text: '3小时内',
          value: () => {
            const end = new Date()
            const start = new Date()
            start.setTime(start.getTime() - 3600 * 1000 * 3)
            this.dateRange = [start, end]
            return [start, end]
          },
        }, {
          text: '6小时内',
          value: () => {
            const end = new Date()
            const start = new Date()
            start.setTime(start.getTime() - 3600 * 1000 * 6)
            this.dateRange = [start, end]
            return [start, end]
          },
        }, {
          text: '12小时内',
          value: () => {
            const end = new Date()
            const start = new Date()
            start.setTime(start.getTime() - 3600 * 1000 * 12)
            this.dateRange = [start, end]
            return [start, end]
          },
        }, {
          text: '24小时内',
          value: () => {
            const end = new Date()
            const start = new Date()
            start.setTime(start.getTime() - 3600 * 1000 * 24)
            this.dateRange = [start, end]
            return [start, end]
          },
        }, {
          text: '本周',
          value: () => {
            const start = this.WeekFirstDay()
            const end = new Date()
            this.dateRange = [start, end]
            return [start, end]
          },
        }, {
          text: '本月',
          value: () => {
            const end = new Date()
            const start = new Date()
            start.setDate(1)
            start.setHours(0)
            start.setSeconds(0)
            start.setMinutes(0)
            start.setTime(start)
            this.dateRange = [start, end]
            return [start, end]
          },
        }, {
          text: '上个月',
          value: () => {
            const start = new Date()
            start.setMonth(0)
            start.setDate(1)
            start.setHours(0)
            start.setMinutes(0)
            start.setSeconds(0)
            start.setTime(start)
            const end = new Date()
            end.setMonth(1)
            end.setDate(0)
            end.setHours(23)
            end.setMinutes(59)
            end.setSeconds(59)
            this.dateRange = [start, end]
            return [start, end]
          },
        }],
      dateRange: [Date.now() - 1000 * 3600, Date.now()],
      listQuery: {},
    }
  },
  watch: {
    'filterData'(val) {
      // console.log(val)
      if (val.elinput.length > 0) {
        val.elinput.map(item => {
          this.listQuery[item.key] = ''
        })
      }
      if (val.elselect.length > 0) {
        val.elinput.map(item => {
          this.listQuery[item.key] = ''
        })
      }
    },
    'filterData.rest': {
      handler: function(val) {
        if (val) {
          this.handleReset()
        }
      },
      deep: true
    }
  },
  methods: {
    WeekFirstDay() {
      var Nowdate = new Date()
      var day = Nowdate.getDay()
      if (day === 0) {
        day = 7
      }
      var WeekFirstDay = new Date(Nowdate - (day - 1) * 86400000)
      WeekFirstDay.setHours(0)
      WeekFirstDay.setMinutes(0)
      WeekFirstDay.setSeconds(0)
      return WeekFirstDay
    },
    handleSearch() {
      const data = this.listQuery
      // 把 时间 写入data，默认的时间范围传过来是unix时间戳，不是x-x-x格式的ban掉
      if (this.dateRange[0].toString().indexOf('-') !== -1) {
        data.startTime = this.dateRange[0]
        data.endTime = this.dateRange[1]
      }
      // 把 协议版本选择框 写入data
      data.protocolVersion = this.protocolVersion
      // 干掉空key
      Object.keys(data).forEach(function(key) {
        if (data[key] === '') {
          delete data[key]
        }
      })
      this.$emit('filterMsg', data)
    },
    handleReset() {
      this.listQuery['dynamicId'] = ''
      this.listQuery['class'] = ''
      this.listQuery['ipAddr'] = ''
      this.dateRange = ['', '']
      this.pageData.pageSize = 20 // 每页数量
      this.pageData.pageNum = 1 // 页码
      this.protocolVersion = '双栈'
    }
  }
}
</script>

<style scoped>
.filter-item{
  margin-left: 10px;
  display: inline-block;
}

.filter-container .filter-item:nth-of-type(1){
  margin-left: 0px;
}
.btn{
  display: inline-block;
  margin-left: 10px;
}
</style>
