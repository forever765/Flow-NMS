<template>
  <div class="dashbord-line-box">
    <div class="dashbord-line-title">最近一小时流量</div>
    <div ref="echart" class="dashbord-line" />
  </div>
</template>
<script>
import echarts from "echarts";
import "echarts/theme/macarons";
import { getTraffic } from "@/api/charts";

const titles = ["上行速率", "下行速率"];
const type = "a";
const unit = "MBps";
export default {
  name: "Line",
  created() {
    this.getData();
  },
  mounted() {
    this.$nextTick(() => {
      this.initChart();
    });
  },
  beforeUnmount() {
    if (!this.chart) {
      return;
    }
    this.chart.dispose();
    this.chart = null;
  },
  methods: {
    async getData() {
      this.result = await getTraffic();
      return this.result;
      // for(var i = 0; i < data.length; i++){
      //   // console.log(data[i].Time + " " + data[i].in_traffic_mbps);
      // }
    },
    initChart() {
      // this.chart = echarts.init(this.$refs.echart, 'macarons')
      this.getData().then(
        (data) => (
          (this.chart = echarts.init(this.$refs.echart)), this.setOptions(data)
        )
      );
    },
    setOptions(data) {
      var result = [];
      for(var i in data['data'])
          result.push([data['data'][i]['Time'], data['data'][i]['in_traffic_mbps']]);
      console.log(result)
      // console.log(data['data'][0]['Time'])
      this.chart.setOption({
        tooltip: {
          trigger: "axis",
          axisPointer: {
            type: "cross",
            label: {
              backgroundColor: "#6a7985",
            },
          },
          backgroundColor: "#fff",
          padding: 10,
          textStyle: {
            fontSize: 12,
            color: "#152934",
            lineHeight: 24,
          },
          extraCssText:
            "box-shadow: 0 0 3px rgba(0, 0, 0, 0.3); border-radius: 0;",
          formatter: (params) => {
            var result = `${params[0].data[0]} <br/>`;
            params.map((item) => {
              result += `${item.seriesName} : ${
                isNaN(item.value[1]) ? "-" : item.value[1]
              } ${unit}</br>`;
            });
            return result;
          },
        },
        grid: {
          left: "35",
          right: "22",
          bottom: "30",
          top: "34",
        },
        yAxis: [
          {
            splitLine: {
              show: true,
              lineStyle: {
                type: "dotted",
                color: "rgba(155, 155, 155, 0.5)",
              },
            },
            axisLine: {
              show: false,
            },
            axisLabel: {
              color: "#5A6872",
              fontSize: 11,
            },
            axisTick: { show: false },
            type: "value",
          },
        ],
        xAxis: [
          {
            type: "time", // x轴为 时间轴
            splitLine: { show: false },
            axisLine: {
              lineStyle: { width: 0 },
            },
            axisLabel: {
              color: "#5A6872",
              fontSize: 11,
              interval: 5,
            },
            axisTick: { show: false },
            boundaryGap: false,
            // data: data['in'].map(function (item) {
            //   return item[0]
            // })
            data: result,
          },
        ],
        legend: { data: titles },
        color: ["#41D6C3", "#5AAAFA"],
        series: [
          {
            name: "上行速率",
            type: "line",
            symbol: "none",
            markPoint: {
              label: {
                normal: {
                  show: true,
                  backgroundColor: "#fff",
                  position: "top",
                  color: "#41D6C3",
                  borderColor: "rgba(65,214,195,0.3)",
                  borderWidth: 1,
                  padding: 8,
                  formatter: `{b}: {c} ${unit}`,
                },
              },
              symbol: "circle",
              itemStyle: {
                normal: {
                  borderColor: "rgba(65,214,195,0.3)",
                  borderWidth: 15,
                },
              },
              symbolSize: 7,
              data: [{ type: "max", name: "Max" }],
            },
            lineStyle: { normal: { color: "#41D6C3", width: 1 } },
            areaStyle: { normal: { color: "#41D6C3", opacity: 0.5 } },
            data: result,
          },
          {
            name: "下行速率",
            type: "line",
            symbol: "none",
            markPoint: {
              label: {
                normal: {
                  show: true,
                  backgroundColor: "#fff",
                  position: "top",
                  color: "#5AAAFA",
                  borderColor: "rgba(90,170,250,0.3)",
                  borderWidth: 1,
                  padding: 8,
                  formatter: `{b}: {c} ${unit}`,
                },
              },
              symbol: "circle",
              itemStyle: {
                normal: {
                  borderColor: "rgba(90,170,250,0.3)",
                  borderWidth: 15,
                },
              },
              symbolSize: 7,
              data: [{ type: "max", name: "Max" }],
            },
            lineStyle: { normal: { color: "#5AAAFA", width: 1 } },
            areaStyle: { normal: { color: "#5AAAFA", opacity: 0.5 } },
            connectNulls: true,
            // ['in_traffic_mbps']
            data: result,
          },
        ],
      });
    },
  },
};
</script>
<style lang="scss" scoped>
.dashbord-line-box {
  .dashbord-line {
    background-color: #fff;
    height: 360px;
    width: 100%;
  }
  .dashbord-line-title {
    font-weight: 600;
    margin-bottom: 12px;
  }
}
</style>
