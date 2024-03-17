<template>
    <div class="index-view">
        <div>
            <!--logo div-->

        </div>
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
    margin: 45%;
}
</style>