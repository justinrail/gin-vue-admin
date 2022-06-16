<template>
  <div>
    <el-row>
      <el-cascader
        placeholder="请选择网关下设备"
        :props="cascaderProps"
        @change="handleChange"
      ></el-cascader>
      <el-button type="primary" round @click="refeshPointData()">刷新</el-button>
    </el-row>
    <el-row>
      <el-table style="width: 100%" ref="multipleTable" :data="tableData">
        <el-table-column prop="deviceID" label="设备ID"> </el-table-column>
        <el-table-column prop="deviceName" label="设备名称"> </el-table-column>
        <el-table-column prop="pointID" label="信号ID"> </el-table-column>
        <el-table-column prop="pointName" label="信号名称"> </el-table-column>
        <el-table-column prop="liveValue" label="值"> </el-table-column>
        <el-table-column prop="updateTime" label="更新时间"> </el-table-column>
      </el-table>
    </el-row>
  </div>
</template>

<script setup>
import { getGatewayDeviceTree } from "@/api/hub/hub_gatewayDeviceTree";
import { getDeviceLivePoints } from "@/api/hub/hub_livePoint";
import { ref } from 'vue'

var tableData = ref([])
var cascade_value = [];
var cascaderProps = {
  expandTrigger: "hover",
  lazy: true,
  async lazyLoad(node, resolve) {
    let tree = await getGatewayDeviceTree();
    if (tree.code === 0) {
      let arr = [];
      tree.data.forEach((element) => {
        let nod = {
          value: element.gatewayID,
          label: element.gatewayName,
          children: [],
        };

        element.Devices.forEach((device) => {
          let dev = {
            value: device.deviceID,
            label: device.deviceName,
            leaf: true,
          };
          nod.children.push(dev);
        });

        arr.push(nod);
      });
      resolve(arr);
    }
  },
};

const handleChange = (value) => {
  cascade_value = value
};

const refeshPointData = async (value) => {
  if (cascade_value.length == 2) {    
    let resp = await getDeviceLivePoints({ gatewayID: cascade_value[0],deviceID: cascade_value[1]});
   
    if (resp.code === 0) {
      tableData.value = resp.data;
      console.log(tableData);
    }
  }
};

</script>

<style>
</style>