<template>
    <p>This is User list page.</p>
    <van-cell-group v-for="u in users.data">
        <van-cell 
            :title="u.id" 
            :value="formatDate(u.birthday)" 
            :label="`${u.first_name}* ${u.last_name}`"
            @click="()=>handleClick(u.id)"
        />
        {{ u.email }}
        {{ u.password }}
    </van-cell-group>

</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
const router = useRouter();

const res = await useFetch("http://127.0.0.1:8086/basic/");
console.table(res);
const users = ref(res);

function formatDate(date) {
    const date_ = new Date(date)
    return date_.toLocaleDateString();
}
function handleClick(id){
    router.push(`/foo/${id}`);
}
</script>