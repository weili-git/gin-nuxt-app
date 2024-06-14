<template>
    <!-- <div style="display: flex; justify-content: center;align-items: center;flex-direction: column;"> -->
    <div>
        <div>
            <el-progress type="dashboard" :percentage="rate" :size="150" :color="customColors">
                <template #default="{ percentage }">
                    <span class="percentage-value">{{ String(temp) }}°C</span>
                </template>
            </el-progress>
        </div>
        <div>
            <el-button type="primary" style="align-self: center;" @click="handleClick">Get Temperature</el-button>
        </div>
    </div>
    
    
</template>

<script setup lang="ts">
const config = useRuntimeConfig();
const temp = useCookie('temp', {maxAge: 60 * 5, default: ()=>{0}}); // 5 min
const rate = ref(50);

const customColors = [
  { color: '#1989fa', percentage: 20 },
  { color: '#5cb87a', percentage: 40 },
  { color: '#e6a23c', percentage: 60 },
  { color: '#f56c6c', percentage: 80 },
  { color: '#d32f2f', percentage: 100 },
]

async function handleClick() {
    const city = "Tokyo";
    const res = await $fetch(`https://api.openweathermap.org/data/2.5/weather?q=${city}&appid=${config.public.apiKey}`);
    const weather = ref(res);
    temp.value = weather.value.main.temp - 273.15; // temp不能改，但是它的value可以修改
    temp.value = Math.floor(temp.value * 10) / 10 // 保留一位小数
    rate.value = (temp.value + 50) *0.9;
    rate.value = Math.floor(rate.value);
}

</script>