<template>
    <div class="index-view">
        <!-- <div>
            

        </div> -->
        <search-input></search-input>
    </div>
</template>

<script lang="ts" setup>
import { reactive } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();
const search = () => {
    router.push({ path: '/search', query: { q: "searchInput.value" } });
}

</script>

<style lang="less" scoped>
.index-view {
    display: grid;
    place-items: center;
    width: 100vw;
    height: 100vh;
}
</style>