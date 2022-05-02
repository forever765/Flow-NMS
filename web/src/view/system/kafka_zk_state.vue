<template>
  <div>
    <el-row>
      <el-card :style="{'height': '100%', 'width': '100%'}">
        <el-card>
          <el-tag>Kafka集群各节点状态</el-tag>
          <el-row type="flex" justify="center" align="middle">
            <el-result v-for="(item,node) in kafkaState" :icon="item" :title="node" sub-title="Kafka节点状态" />
          </el-row>
        </el-card>
        <el-card>
          <el-tag>ZooKeeper集群各节点状态</el-tag>
          <el-row type="flex" justify="center" align="middle">
            <el-result v-for="(item,node) in zkState" :icon="item" :title="node" sub-title="Zk节点状态" />
          </el-row>
        </el-card>
      </el-card>
    </el-row>
  </div>
</template>

<script>
import { getKafkaAndZkInfo } from '@/api/system'

export default {
  name: 'KafkaAndZkInfo',
  data() {
    return {
      timer: null,
      kafkaState: this.kafkaState,
      zkState: this.zkState
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
      var result = await getKafkaAndZkInfo()
      this.kafkaState = result['data']['kafka']
      this.zkState = result['data']['zookeeper']
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
