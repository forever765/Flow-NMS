<template>
  <div>
    <el-row :gutter="10">
      <el-col :span="24">
        <daily-detail />
      </el-col>
    </el-row>
  </div>
</template>

<script>
import dailyDetail from '@/view/detail/daily/dailyDetail.vue'
export default {
  name: 'Detail',
  components: {
    dailyDetail,
  },
}
</script>

<style scoped>
.load-more {
  margin-left: 120px;
}

.avatar-img {
  float: left;
  height: 40px;
  width: 40px;
  border-radius: 50%;
  -webkit-border-radius: 50%;
  -moz-border-radius: 50%;
  margin-top: 15px;
}

.org-img {
  height: 150px;
  width: 150px;
}

.author-name {
  float: left;
  line-height: 65px !important;
  margin-left: 10px;
  color: darkblue;
  line-height: 100px;
  font-family: "Lucida Sans", "Lucida Sans Regular", "Lucida Grande",
    "Lucida Sans Unicode", Geneva, Verdana, sans-serif;
}

.dom-center {
  margin-left: 50%;
  transform: translateX(-50%);
}
</style>
