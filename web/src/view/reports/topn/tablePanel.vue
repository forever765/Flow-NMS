<template>
  <div>
    <el-row :gutter="50">
      <el-col :span="12">
        <el-table
          ref="table"
          v-loading="dataSource.loading"
          :border="dataSource.border?true:false"
          style="width: 500px; height: 30%; font-size: 10px"
          :row-style="{height:'10px'}"
          :cell-style="{padding:'0px', 'text-align':'center'}"
          :class="{ 'no-data': !dataSource.data || !dataSource.data.length }"
          :data="dataSource.src"
          @row-click="getRowData"
          @selection-change="dataSource.handleSelectionChange"
        >
          <!-- 是否显示序号 -->
          <el-table-column
            v-if="dataSource.isIndex"
            type="index"
            label="序号"
            width="55"
            align="center"
          />

          <template v-for="item in dataSource.cols">
            <!-- 需要带图标的某列，带回调事件-->
            <el-table-column
              v-if="item.isIcon"
              :key="item.prop"
              v-bind="item"
              align="center"
            >
              <template slot-scope="scope">
                <span>
                  <span>{{ item.filter(scope.row[scope.column.property]) }}</span>
                  <i v-if="item.icon" :class="[item.icon(scope.row[scope.column.property]),'icon-normal']" @click="item.handlerClick(scope.row)" />
                </span>
                <!-- 比如要输入框 显示图片等等 自己定义 -->
                <slot :name="item.prop" :scope="scope" />
              </template>
            </el-table-column>
            <!-- 大部分适用 -->
            <el-table-column
              v-if="!item.isImagePopover && !item.isTemplate && !item.isSpecial&&!item.isIcon"
              :key="item.prop"
              v-bind="item.isCodeTableFormatter ? Object.assign({ formatter: item.isCodeTableFormatter }, item) : item"
              align="center"
              show-overflow-tooltip
            />
          </template>
          <!-- 是否有操作列 -->
          <!-- 没有数据时候不固定列 -->
          <el-table-column
            v-if="dataSource.isOperation"
            :show-overflow-tooltip="dataSource.operation.overflowTooltip"
            v-bind="dataSource.data && dataSource.data.length ? { fixed: 'right' } : null"
            style="margin-right:20px"
            class-name="handle-td"
            label-class-name="tc"
            :width="dataSource.operation.width"
            :label="dataSource.operation.label"
            align="center"
          >
            <!-- UI统一一排放3个，4个以上出现更多 -->
            <!--        <template slot-scope="scope">-->
            <!--          &lt;!&ndash; 三个一排的情况，去掉隐藏的按钮后的长度 &ndash;&gt;-->
            <!--          <template v-if="dataSource.operation.data.length > 0">-->
            <!--            <div class="btn">-->
            <!--              <div v-for="(item) in dataSource.operation.data" :key="item.label">-->
            <!--                <template v-if="item.type!=='icon'">-->
            <!--                  <el-button-->
            <!--                    v-permission="item.permission"-->
            <!--                    v-bind="item"-->
            <!--                    :type="item.type?item.type:''"-->
            <!--                    size="mini"-->
            <!--                    @click.native.prevent="item.handleRow(scope.$index, scope.row, item.label)"-->
            <!--                  >-->
            <!--                    {{ item.label }}-->
            <!--                  </el-button>-->
            <!--                </template>-->
            <!--                <template v-else>-->
            <!--                  <i :class="[icon,item.icon]" v-bind="item" @click="item.handleRow(scope.$index, scope.row, item.label)" />-->
            <!--                </template>-->
            <!--              </div>-->
            <!--            </div>-->
            <!--          </template>-->
            <!--        </template>-->
          </el-table-column>
        </el-table>
      </el-col>
      <el-col :span="12">
        <el-table
          ref="table"
          v-loading="dataSource.loading"
          :border="dataSource.border?true:false"
          style="width: 500px; height: 30%; font-size: 10px"
          :row-style="{height:'10px'}"
          :cell-style="{padding:'0px', 'text-align':'center'}"
          :class="{ 'no-data': !dataSource.data || !dataSource.data.length }"
          :data="dataSource.dst"
          @row-click="getRowData"
          @selection-change="dataSource.handleSelectionChange"
        >
          <!-- 是否显示序号 -->
          <el-table-column
            v-if="dataSource.isIndex"
            type="index"
            label="序号"
            width="55"
            align="center"
          />

          <template v-for="item in dataSource.cols">
            <!-- 需要带图标的某列，带回调事件-->
            <el-table-column
              v-if="item.isIcon"
              :key="item.prop"
              v-bind="item"
              align="center"
            >
              <template slot-scope="scope">
                <span>
                  <span>{{ item.filter(scope.row[scope.column.property]) }}</span>
                  <i v-if="item.icon" :class="[item.icon(scope.row[scope.column.property]),'icon-normal']" @click="item.handlerClick(scope.row)" />
                </span>
                <!-- 比如要输入框 显示图片等等 自己定义 -->
                <slot :name="item.prop" :scope="scope" />
              </template>
            </el-table-column>
            <!-- 大部分适用 -->
            <el-table-column
              v-if="!item.isImagePopover && !item.isTemplate && !item.isSpecial&&!item.isIcon"
              :key="item.prop"
              v-bind="item.isCodeTableFormatter ? Object.assign({ formatter: item.isCodeTableFormatter }, item) : item"
              align="center"
              show-overflow-tooltip
            />
          </template>
          <!-- 是否有操作列 -->
          <!-- 没有数据时候不固定列 -->
          <el-table-column
            v-if="dataSource.isOperation"
            :show-overflow-tooltip="dataSource.operation.overflowTooltip"
            v-bind="dataSource.data && dataSource.data.length ? { fixed: 'right' } : null"
            style="margin-right:20px"
            class-name="handle-td"
            label-class-name="tc"
            :width="dataSource.operation.width"
            :label="dataSource.operation.label"
            align="center"
          >
          </el-table-column>
        </el-table>
      </el-col>

    </el-row>
    <!--  下方页脚 -->
    <div class="page">
      <el-pagination
        v-if="dataSource.pageData.total>0"
        background
        :current-page="dataSource.pageData.pageNum"
        :page-sizes="dataSource.pageData.pageSizes?dataSource.pageData.pageSizes:[20,30,50,100]"
        :page-size="dataSource.pageData.pageSize"
        layout="total, sizes, prev, pager, next, jumper"
        :total="dataSource.pageData.total"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>
  </div>
</template>

<script>
export default {
  name: 'TopNTablePane',
  // 接收父组件传递过来的值
  props: {
    //  表格数据和表格部分属性的对象
    // eslint-disable-next-line vue/require-default-prop
    dataSource: {
      type: Object
    }
  },
  data() {
    return {
      // currentPage4: 4
    }
  },
  watch: {
    'dataSource.cols': { // 监听表格列变化
      deep: true,
      handler() {
        // 解决表格列变动的抖动问题
        this.$nextTick(this.$refs.table.doLayout)
      }
    }
  },
  methods: {
    handleAdd(name) {
      console.log(name)
      this.$emit('toolMsg', name)
    },
    handleRow(index, row, lable) {
      console.log(index, row, lable)
    },
    handleSizeChange(val) {
      this.$emit('changeSize', val)
      // console.log(`每页 ${val} 条`)
    },
    handleCurrentChange(val) {
      this.$emit('changeNum', val)
      // console.log(`当前页: ${val}`)
    },
    // 点击行即可选中
    getRowData(row) {
      this.$refs.table.toggleRowSelection(row)
    }
  }
}
</script>
<style lang="scss" scoped>
.page{
  margin-top: 20px;
}
.btn{
  display: flex;
  justify-content: center;
}
.btn div{
  margin-left: 5px;
}
.reference-img{
  width: 40px;
  height: 40px;
  background-size:100% 100%;
  border-radius: 4px;
}
.image-popover{
  width: 200px;
  height: 200px;
  background-size:100% 100%;
}
.icon {
  width: 25px;
  font-size: 20px;
  font-weight: bold;
}
</style>
