<template>
    <!-- <div style="display: flex; justify-content: center;align-items: center;flex-direction: column;"> -->
    <div>
        <div>
            <van-circle
                v-model:current-rate="rate"
                :rate="rate"
                :size="150"
                layer-color="#ebedf0"
                :color="gradientColor"
                :text="`${String(temp)}°C`"
            />
        </div>
        <div>
            <van-button type="primary" style="align-self: center;" @click="handleClick">Get Temperature</van-button>
        </div>
    </div>
    
    
</template>

<script setup lang="ts">
const config = useRuntimeConfig();
const temp = useCookie('temp', {maxAge: 60 * 5, default: ()=>{0}}); // 5 min
const rate = ref(50);
const gradientColor = {
    '0%': '#0000ff',
    '100%': '#ff0000',
};
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