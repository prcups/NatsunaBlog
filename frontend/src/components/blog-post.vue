<template>
    <b-col cols="12">
        <b-card :title="title">
            <b-card-text>
                {{ time }}
            </b-card-text>
            <b-button variant="light" :to="'post/' + id">点击阅读</b-button>
        </b-card>
    </b-col>
</template>

<script>
    export default {
        name: "blog-post",
        props: ['title', 'time', 'id']
    }
</script>

<style scoped>

</style>
