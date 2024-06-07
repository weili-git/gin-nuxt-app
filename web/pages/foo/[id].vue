<template>
    <h1>Hello</h1>
    {{ data }}
</template>

<script setup>
const route = useRoute();
const id = computed(() => +route.params.id);
const res = await useFetch(`http://localhost:8086/basic/${id.value}/`)
const data = ref(res);
</script>