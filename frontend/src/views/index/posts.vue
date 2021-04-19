<template>
    <div>
        <div style="min-height: 60vh">
            <b-row cols="12">
                <blog-post v-for="item in posts" v-bind="item" :key="item.id"></blog-post>
            </b-row>
        </div>
        <b-pagination-nav pills size="ml" :link-gen="linkGen" :number-of-pages="pages" align="center" use-router></b-pagination-nav>
    </div>
</template>

<script>
    import BlogPost from "../../components/blog-post"
    import axios from 'axios'
    import url from "../../assets/javascript/url"
    export default {
        name: "index-posts",
        components: {
            'blog-post': BlogPost
        },
        data() {
            return {
                posts: [],
                pages: 1
            }
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
        created() {
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
        watch: {
            '$route': function (newVal) {
                this.$router.go(0)
            }
        }
    }
</script>

<style scoped>
</style>
