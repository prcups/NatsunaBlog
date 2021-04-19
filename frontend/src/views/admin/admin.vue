<template>
    <b-container align="center">
        <b-navbar type="dark" variant="info">
            <b-navbar-brand href="/#/admin">AvaBlog控制台</b-navbar-brand>
            <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>
            <b-collapse id="nav-collapse" is-nav>
                <b-navbar-nav>
                    <b-nav-item href="/#/admin/manage">管理</b-nav-item>
                    <b-nav-item href="/#/admin/modify/-1">新建</b-nav-item>
                    <b-nav-item href="/#/">返回首页</b-nav-item>
                </b-navbar-nav>
            </b-collapse>
            <b-navbar-nav class="ml-auto">
                <b-nav-item size="sm" disabled>{{user}}</b-nav-item>
                <b-button size="sm" @click="buttonAct" class="my-2 my-sm-0">{{buttonWord}}</b-button>
            </b-navbar-nav>
        </b-navbar>
        <b-nav-text v-if="$router.currentRoute.path == '/admin'">
            欢迎来到AvaBlog控制台！<br>
            请先点击右上角登录。<br>
            选择“管理“可编辑或删除您写过的文章。”新建“会创建新文章。<br>
            本博客程序使用了Vue、BootstrapVue、GoFrame、MavonEditor、axios等框架或组件，遵循MulanPSL 2.0协议发布。
        </b-nav-text>
        <router-view></router-view>
    </b-container>
</template>

<script>
    import axios from 'axios'
    import {setIfChecked, logout, toCheck} from "../../assets/javascript/check.js"
    axios.defaults.withCredentials=true
    export default {
        name: "admin",
        data() {
            return {
                user: "",
                isChecked: undefined
            }
        },
        methods: {
            buttonAct() {
                if (this.isChecked) {
                    logout(this)
                } else {
                    toCheck(this)
                }
            }
        },
        created() {
            setIfChecked(this)
        },
        computed: {
            buttonWord: function () {
                return this.isChecked ? "退出" : "登录"
            }
        },
        watch: {
            '$route' (newVal, oldVal) {
                if (oldVal.name == "Check") {
                    this.$router.go(0)
                }
            }
        }
    }
</script>

<style scoped>

</style>