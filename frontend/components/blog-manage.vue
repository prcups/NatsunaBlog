<template>
    <b-col cols="4" style="min-width: 300px; margin-left: auto; margin-right: auto">
        <b-card :title="title">
            <b-card-text>
                {{ time }}
            </b-card-text>
            <b-button variant="light" :href=editUrl>编辑</b-button>
            <b-button variant="light" @click="deletePost">删除</b-button>
        </b-card>
    </b-col>
</template>

<script>
    import axios from "axios"

    export default {
        name: "blog-manage",
        props: ['title', 'time', 'id'],
        methods: {
            deletePost() {
                axios({
                    method: "post",
                    url: this.configVal.DeletePostUrl,
                    data: {
                        id: this.id
                    }
                })
                    .then(res => {
                        this.$router.go(0)
                    })
            }
        },
        computed:{
            editUrl: function () {
                return "/admin/modify/" + this.id
            }
        }
    }
</script>

<style scoped>

</style>