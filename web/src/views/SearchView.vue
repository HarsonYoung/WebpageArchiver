<template>
    <div class="index-view">
        <search-input></search-input>
        <span>约为 {{ TotalHits }} 条结果</span>

        <div class="error-area" v-if="errorStatus">
            <a-alert type="error">{{ erroMessage }}</a-alert>
        </div>

        <div class="item-card" v-for="item in pageData.jsonResult.result">
            <a-card :style="{ width: '600px' }" v-if="item.content != '' && errorStatus == false">
                <template #title>
                    <a-link :href=" fileLink+item.filename" target="_blank">{{ item.title }}</a-link>
                </template>
                <template #extra>
                    <a-link :href="item.link" target="_blank">原文链接</a-link>
                </template>
                <span v-html="item.content"></span>
            </a-card>
        </div>
        <div class="pagination">
            <a-pagination :total=Number(TotalHits) :page-size="10" size="large"  @change="changePage" />
        </div>
    </div>
</template>

<script lang="ts" setup>
import { useRoute, useRouter } from 'vue-router';
import { onMounted, reactive, watch } from 'vue';

// 使用接口声明类型
interface ResultItem {
    title: string
    filename: string
    link: string
    content: string
}

let erroMessage = "搜索请求出现错误"
let errorStatus = true
let TotalHits = "0"
let pageNum = "1"
let fileLink = "/archive/"

// 双向绑定
let pageData = reactive({
    searchKey: "",
    jsonResult: {
        result: {} as ResultItem[]
    },
});

const route = useRoute();
const router = useRouter();

const changePage = (currentPage: number) => {
    router.push({ path: '/search', query: { q: pageData.searchKey, p: currentPage } });
}

function queryData(keyword: string, pages : string = "1") {
    // let queryURL = `http://127.0.0.1:7845/api/search?q=${encodeURIComponent(keyword)}&p=${pages}`
    let queryURL = `/api/search?q=${encodeURIComponent(keyword)}&p=${pages}`

    fetch(queryURL)
        .then((response) => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then((data) => {
            console.log(data);
            console.log(JSON.parse(data.Result))
            TotalHits = data.TotalHits
            
            console.log(data)
            if (data.Status == "0") {
                errorStatus = true
                erroMessage = data.Message
            }
            else {
                errorStatus = false
                pageData.jsonResult = { result: JSON.parse(data.Result) }
            }

        })
        .catch((error) => {
            errorStatus = true
            console.error('There was an error:', error);
        });
}

onMounted(() => {
    pageData.searchKey = route.query.q! as string
    pageNum = route.query.p! as string
});

const scrollToTop = () => {
  window.scrollTo({
    // top: document.documentElement.offsetHeight, //回到底部
    top: 0, //回到顶部
    left: 0,
    behavior: "smooth", //smooth 平滑；auto:瞬间
  });
};

// 监听q参数
watch(() => route.query.q, (newData) => {
    // console.log(newData)
    pageData.searchKey = newData! as string

    if(pageData.searchKey != null && pageData.searchKey != ""){
        queryData(pageData.searchKey)
    }
    else{
        erroMessage = "请输入关键字"
    }

    // pageData.jsonResult = testJson
}, { immediate: true });

// 监听p参数
watch(() => route.query.p, (newData) => {
    // console.log(newData)
    pageNum = newData! as string
    queryData(pageData.searchKey, pageNum)
    scrollToTop()
}, { immediate: true });


</script>


<style lang="less" scoped>
.index-view {
    display: grid;
    place-items: center;
    width: 99vw;
    margin-top: 5vh;
    margin-bottom: 2vh;
}

.item-card {
    overflow: hidden
    
}

span.highlight {
    color: red;
}

.pagination{
    margin-top: 10px;
}
</style>