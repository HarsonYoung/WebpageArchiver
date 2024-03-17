<template>
    <div class="search-box">
        <a-input-search size="large" v-model="pageData.searchKey" :style="{ width: '45em' }" placeholder="输入搜索内容" @search="getSearch" @keydown.enter.native="getSearch"
            search-button />

    </div>
</template>

<script setup lang="ts">
import { useRoute, useRouter } from 'vue-router';
import { onMounted, reactive } from 'vue';

let pageData = reactive({
    searchKey: ''
});

const route = useRoute();
const router = useRouter();

const getSearch = () => {
    router.push({ path: '/search', query: { q: pageData.searchKey } })
}

onMounted(() => {
    pageData.searchKey = route.query.q! as string;
});
</script>


<style scoped>

</style>