<template>
    <div>
        <div v-if="isChecked">
            <br>
            <h3>所有文章</h3>
            <br>
            <div style="min-height: 75vh">
                <b-row cols="12">
                    <blog-manage v-for="item in posts" v-bind="item" :key="item.id"></blog-manage>
                </b-row>
            </div>
            <b-pagination-nav :link-gen="linkGen" :number-of-pages="pages" align="center" use-router></b-pagination-nav>
        </div>
        <div v-else>
            <p>请先登录</p>
        </div>
    </div>
</template>

<script>
    import BlogManage from "../../components/blog-manage"
    import {setIfChecked} from "../../assets/javascript/check.js"
    import axios from "axios"
    import url from "../../assets/javascript/url"

    export default {
        name: "manage",
        components: {
            'blog-manage': BlogManage
        },
        data() {
            return {
                isChecked: undefined,
                pages: 1,
                posts: []
            }
        },
        created() {
            setIfChecked(this)
            axios({
                method:'get',
                url: url.GetPageNumUrl,
            })
                .then(res => {
                    this.pages = res.data.pages != 0 ? res.data.pages : 1
                })
            axios({
                method: 'get',
                url: url.GetPostsUrl,
                params: {
                    page: this.curPage()
                }
            })
                .then(res => {
                    this.posts = res.data.posts
                })
        },
        methods: {
            linkGen(pageNum) {
                return pageNum === 1 ? '?' : `?page=${pageNum}`
            },
            curPage(){
                let queryPage = this.$route.query.page
                return queryPage ? queryPage : 1
            }
        },
        watch: {
            '$route' () {
                    this.$router.go(0)
            }
        }
    }
</script>

<style scoped>
</style>